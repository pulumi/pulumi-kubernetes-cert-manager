# Override during CI using `make [TARGET] PROVIDER_VERSION=""` or by setting a PROVIDER_VERSION environment variable
# Local & branch builds will just used this fixed default version unless specified
PROVIDER_VERSION ?= 1.0.0-alpha.0+dev
# Use this normalised version everywhere rather than the raw input to ensure consistency.
VERSION_GENERIC = $(shell pulumictl convert-version --language generic --version "$(PROVIDER_VERSION)")

PACK            := kubernetes-cert-manager
PROJECT         := github.com/pulumi/pulumi-${PACK}

PROVIDER        := pulumi-resource-${PACK}
VERSION_PATH    := provider/pkg/version.Version

WORKING_DIR     := $(shell pwd)
SCHEMA_PATH     := ${WORKING_DIR}/provider/cmd/${PROVIDER}/schema.json

GOPATH          := $(shell go env GOPATH)

export PULUMI_IGNORE_AMBIENT_PLUGINS = true

# Ensure the codegen file is present so that the hard-coded "Tar Provider Binaries" step doesn't fail
codegen: .pulumi/bin/pulumi # Required by CI
	mkdir -p bin && touch bin/pulumi-gen-kubernetes-cert-manager

provider: build_provider # Required by CI
test_provider: # Required by CI
generate_schema: # Required by CI
local_generate: generate # Required by CI

generate:: gen_go_sdk gen_dotnet_sdk gen_nodejs_sdk gen_python_sdk generate_java

build:: build_provider build_dotnet_sdk build_nodejs_sdk build_python_sdk

install:: install_provider install_dotnet_sdk install_nodejs_sdk


# Provider

build_provider::
	rm -rf ${WORKING_DIR}/bin/${PROVIDER}
	cd provider/cmd/${PROVIDER} && go build -o ${WORKING_DIR}/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION_GENERIC}" .

install_provider:: build_provider
	cp ${WORKING_DIR}/bin/${PROVIDER} ${GOPATH}/bin


# Go SDK

gen_go_sdk: .pulumi/bin/pulumi
	rm -rf sdk/go
	.pulumi/bin/pulumi package gen-sdk ${SCHEMA_PATH} --language go --version ${VERSION_GENERIC}
build_go_sdk::
generate_go: gen_go_sdk # Required by CI
build_go: # Required by CI
install_go_sdk:: # Required by CI

# .NET SDK

gen_dotnet_sdk: .pulumi/bin/pulumi
	rm -rf sdk/dotnet
	.pulumi/bin/pulumi package gen-sdk ${SCHEMA_PATH} --language dotnet --version ${VERSION_GENERIC}

build_dotnet_sdk:: gen_dotnet_sdk
	cd sdk/dotnet/ && \
		dotnet build

install_dotnet_sdk:: # Required by CI
	rm -rf ${WORKING_DIR}/nuget
	mkdir -p ${WORKING_DIR}/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

generate_dotnet: gen_dotnet_sdk # Required by CI
build_dotnet: build_dotnet_sdk # Required by CI

# Node.js SDK

gen_nodejs_sdk: .pulumi/bin/pulumi
	rm -rf sdk/nodejs
	.pulumi/bin/pulumi package gen-sdk ${SCHEMA_PATH} --language nodejs --version ${VERSION_GENERIC}
	# HACKHACK: work around https://github.com/pulumi/pulumi/issues/7979:
	find sdk/nodejs -name "*.ts" -exec sed -i.bak \
		's/pulumiKubernetes\.types\.input\.\([a-zA-Z0-9]*\)\.\([a-zA-Z0-9]*\)\.\([a-zA-Z]*\)Args/pulumiKubernetes.types.input.\1.\2.\3/g' \
			'{}' \;

build_nodejs_sdk:: gen_nodejs_sdk
	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc --version && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/

generate_nodejs: gen_nodejs_sdk # Required by CI
build_nodejs: build_nodejs_sdk # Required by CI
install_nodejs_sdk:: # Required by CI
	yarn unlink ${PACK} || true
	yarn link --cwd ${WORKING_DIR}/sdk/nodejs/bin

# Python SDK

gen_python_sdk: .pulumi/bin/pulumi
	rm -rf sdk/python
	.pulumi/bin/pulumi package gen-sdk ${SCHEMA_PATH} --language python --version ${VERSION_GENERIC}
	cp ${WORKING_DIR}/README.md sdk/python

build_python_sdk:: gen_python_sdk
	cd sdk/python/ && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build && \
		cd ./bin && ../venv/bin/python -m build .

generate_python: gen_python_sdk # Required by CI
build_python: build_python_sdk # Required by CI
install_python_sdk:: # Required by CI

# Java SDK

generate_java: .pulumi/bin/pulumi # Required by CI
	.pulumi/bin/pulumi package gen-sdk ${SCHEMA_PATH} -o sdk --language java
	cp ${WORKING_DIR}/README.md sdk/java
build_java: generate_java # Required by CI
	cd sdk/java && gradle --console=plain build

install_java_sdk: # Required by CI

# Pulumi for codegen

.pulumi/bin/pulumi: PULUMI_VERSION := $(shell cat .pulumi.version)
.pulumi/bin/pulumi: HOME := $(WORKING_DIR)
.pulumi/bin/pulumi: .pulumi.version
	curl -fsSL https://get.pulumi.com | sh -s -- --version "$(PULUMI_VERSION)"

# Set these variables to enable signing of the windows binary
AZURE_SIGNING_CLIENT_ID ?=
AZURE_SIGNING_CLIENT_SECRET ?=
AZURE_SIGNING_TENANT_ID ?=
AZURE_SIGNING_KEY_VAULT_URI ?=
SKIP_SIGNING ?=

bin/jsign-6.0.jar:
	wget https://github.com/ebourg/jsign/releases/download/6.0/jsign-6.0.jar --output-document=bin/jsign-6.0.jar

sign-windows-exe-amd64: GORELEASER_ARCH := amd64_v1
sign-windows-exe-arm64: GORELEASER_ARCH := arm64

sign-windows-exe-%: bin/jsign-6.0.jar
	@# Only sign windows binary if fully configured.
	@# Test variables set by joining with | between and looking for || showing at least one variable is empty.
	@# Move the binary to a temporary location and sign it there to avoid the target being up-to-date if signing fails.
	@set -e; \
	if [[ "${SKIP_SIGNING}" != "true" ]]; then \
		if [[ "|${AZURE_SIGNING_CLIENT_ID}|${AZURE_SIGNING_CLIENT_SECRET}|${AZURE_SIGNING_TENANT_ID}|${AZURE_SIGNING_KEY_VAULT_URI}|" == *"||"* ]]; then \
			echo "Can't sign windows binaries as required configuration not set: AZURE_SIGNING_CLIENT_ID, AZURE_SIGNING_CLIENT_SECRET, AZURE_SIGNING_TENANT_ID, AZURE_SIGNING_KEY_VAULT_URI"; \
			echo "To rebuild with signing delete the unsigned windows exe file and rebuild with the fixed configuration"; \
			if [[ "${CI}" == "true" ]]; then exit 1; fi; \
		else \
			file=dist/pulumi-kubernetes-cert-manager_windows_${GORELEASER_ARCH}/pulumi-resource-kubernetes-cert-manager.exe; \
			mv $${file} $${file}.unsigned; \
			az login --service-principal \
				--username "${AZURE_SIGNING_CLIENT_ID}" \
				--password "${AZURE_SIGNING_CLIENT_SECRET}" \
				--tenant "${AZURE_SIGNING_TENANT_ID}" \
				--output none; \
			ACCESS_TOKEN=$$(az account get-access-token --resource "https://vault.azure.net" | jq -r .accessToken); \
			java -jar bin/jsign-6.0.jar \
				--storetype AZUREKEYVAULT \
				--keystore "PulumiCodeSigning" \
				--url "${AZURE_SIGNING_KEY_VAULT_URI}" \
				--storepass "$${ACCESS_TOKEN}" \
				$${file}.unsigned; \
			mv $${file}.unsigned $${file}; \
			az logout; \
		fi; \
	fi
