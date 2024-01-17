export GOFLAGS=-mod=vendor
export GOPRIVATE=github.com/wallester

mod-vendor:
	go mod tidy && go mod vendor

validate:
	openapi-generator validate -i openapi.yaml

generate:
	openapi-generator generate \
	-i openapi.yaml \
	-g go \
	--additional-properties packageName=onfido_openapi,useTags=true \
	--git-user-id wallester \
  	--git-repo-id go-onfido-openapi
