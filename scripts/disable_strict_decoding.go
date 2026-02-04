//go:build ignore

// This script disables strict JSON decoding in generated model files to allow unknown fields from Onfido API responses.
// Run after `make generate`: go run scripts/disable_strict_decoding.go
package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files, err := filepath.Glob("model_*.go")
	if err != nil {
		panic(err)
	}

	files = append(files, "utils.go")

	count := 0
	for _, file := range files {
		if patchFile(file) {
			count++
		}
	}

	println("✓ Disabled strict JSON decoding in", count, "files")
}

func patchFile(filename string) bool {
	content, err := os.ReadFile(filename)
	if err != nil {
		return false
	}

	s := string(content)

	// Replace decoder.DisallowUnknownFields() in model files
	oldDecoder := "decoder.DisallowUnknownFields()"
	newDecoder := "// decoder.DisallowUnknownFields() - disabled to allow unknown fields from Onfido API"

	// Replace dec.DisallowUnknownFields() in utils.go
	oldDec := "dec.DisallowUnknownFields()"
	newDec := "// dec.DisallowUnknownFields() - disabled to allow unknown fields from Onfido API"

	modified := false

	if strings.Contains(s, oldDecoder) {
		s = strings.ReplaceAll(s, oldDecoder, newDecoder)
		modified = true
	}

	if strings.Contains(s, oldDec) {
		s = strings.ReplaceAll(s, oldDec, newDec)
		modified = true
	}

	if modified {
		err = os.WriteFile(filename, []byte(s), 0644)
		if err != nil {
			panic(err)
		}
	}

	return modified
}
