export GOFLAGS=-mod=vendor
export GOPRIVATE=github.com/wallester
export GO_POST_PROCESS_FILE=gofmt -w

mod-vendor:
	go mod tidy && go mod vendor

validate:
	openapi-generator validate -i openapi.yaml

generate:
	openapi-generator generate \
	-i openapi.yaml \
	-g go \
	--enable-post-process-file \
	--additional-properties packageName=onfido_openapi,useTags=true \
	--git-user-id wallester \
	--git-repo-id go-onfido-openapi
	@echo "Patching model_report.go ..."
	go run -mod=mod scripts/patch_model_report.go
	go mod tidy && go mod vendor
