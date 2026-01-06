//go:build ignore

// This script patches model_report.go to use discriminator-based parsing
// Run after `make generate`: go run scripts/patch_model_report.go
package main

import (
	"os"
	"strings"
)

const newUnmarshalJSON = `// Unmarshal JSON data into one of the pointers in the struct using the 'name' discriminator
func (dst *Report) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		Name string ` + "`json:\"name\"`" + `
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return fmt.Errorf("failed to extract discriminator 'name': %w", err)
	}

	switch discriminator.Name {
	case "device_intelligence":
		dst.DeviceIntelligenceReport = &DeviceIntelligenceReport{}
		return json.Unmarshal(data, dst.DeviceIntelligenceReport)
	case "document":
		dst.DocumentReport = &DocumentReport{}
		return json.Unmarshal(data, dst.DocumentReport)
	case "document_with_address_information":
		dst.DocumentWithAddressInformationReport = &DocumentWithAddressInformationReport{}
		return json.Unmarshal(data, dst.DocumentWithAddressInformationReport)
	case "document_with_driver_verification":
		dst.DocumentWithDriverVerificationReport = &DocumentWithDriverVerificationReport{}
		return json.Unmarshal(data, dst.DocumentWithDriverVerificationReport)
	case "document_with_driving_licence_information":
		dst.DocumentWithDrivingLicenceInformationReport = &DocumentWithDrivingLicenceInformationReport{}
		return json.Unmarshal(data, dst.DocumentWithDrivingLicenceInformationReport)
	case "facial_similarity_photo":
		dst.FacialSimilarityPhotoReport = &FacialSimilarityPhotoReport{}
		return json.Unmarshal(data, dst.FacialSimilarityPhotoReport)
	case "facial_similarity_video":
		dst.FacialSimilarityVideoReport = &FacialSimilarityVideoReport{}
		return json.Unmarshal(data, dst.FacialSimilarityVideoReport)
	case "identity_enhanced":
		dst.IdentityEnhancedReport = &IdentityEnhancedReport{}
		return json.Unmarshal(data, dst.IdentityEnhancedReport)
	case "known_faces":
		dst.KnownFacesReport = &KnownFacesReport{}
		return json.Unmarshal(data, dst.KnownFacesReport)
	case "photo_fully_auto":
		dst.PhotoFullyAutoReport = &PhotoFullyAutoReport{}
		return json.Unmarshal(data, dst.PhotoFullyAutoReport)
	case "proof_of_address":
		dst.ProofOfAddressReport = &ProofOfAddressReport{}
		return json.Unmarshal(data, dst.ProofOfAddressReport)
	case "right_to_work":
		dst.RightToWorkReport = &RightToWorkReport{}
		return json.Unmarshal(data, dst.RightToWorkReport)
	case "watchlist_enhanced":
		dst.WatchlistEnhancedReport = &WatchlistEnhancedReport{}
		return json.Unmarshal(data, dst.WatchlistEnhancedReport)
	case "watchlist_peps_only":
		dst.WatchlistPepsOnlyReport = &WatchlistPepsOnlyReport{}
		return json.Unmarshal(data, dst.WatchlistPepsOnlyReport)
	case "watchlist_sanctions_only":
		dst.WatchlistSanctionsOnlyReport = &WatchlistSanctionsOnlyReport{}
		return json.Unmarshal(data, dst.WatchlistSanctionsOnlyReport)
	case "watchlist_standard":
		dst.WatchlistStandardReport = &WatchlistStandardReport{}
		return json.Unmarshal(data, dst.WatchlistStandardReport)
	default:
		return fmt.Errorf("unknown report type: %s", discriminator.Name)
	}
}`

func main() {
	content, err := os.ReadFile("model_report.go")
	if err != nil {
		panic(err)
	}

	s := string(content)

	s = strings.Replace(s, "\t\"gopkg.in/validator.v2\"\n", "", 1)

	startMarker := "// Unmarshal JSON data into one of the pointers in the struct\nfunc (dst *Report) UnmarshalJSON(data []byte) error {"
	endMarker := "\n\n// Marshal data from the first non-nil pointers in the struct to JSON"

	startIdx := strings.Index(s, startMarker)
	if startIdx == -1 {
		startMarker = "func (dst *Report) UnmarshalJSON(data []byte) error {"
		startIdx = strings.Index(s, startMarker)

		if startIdx == -1 {
			println("Could not find UnmarshalJSON function start")
			return
		}

		commentStart := strings.LastIndex(s[:startIdx], "// ")
		if commentStart != -1 && startIdx-commentStart < 100 {
			startIdx = commentStart
		}
	}

	endIdx := strings.Index(s[startIdx:], endMarker)
	if endIdx == -1 {
		println("Could not find UnmarshalJSON function end")
		return
	}
	endIdx += startIdx

	newContent := s[:startIdx] + newUnmarshalJSON + s[endIdx:]

	err = os.WriteFile("model_report.go", []byte(newContent), 0644)
	if err != nil {
		panic(err)
	}

	println("✓ Patched model_report.go with discriminator-based UnmarshalJSON")
}
