# Override during CI using `make [TARGET] PROVIDER_VERSION=""` or by setting a PROVIDER_VERSION environment variable
# Local & branch builds will just used this fixed default version unless specified
PROVIDER_VERSION ?= 1.0.0-alpha.0+dev
# Use this normalised version everywhere rather than the raw input to ensure consistency.
VERSION_GENERIC = $(shell pulumictl convert-version --language generic --version "$(PROVIDER_VERSION)")

PACK            := kubernetes-cert-manager
PROJECT         := github.com/pulumi/pulumi-${PACK}

PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION_PATH    := pkg/version.Version

WORKING_DIR     := $(shell pwd)
SCHEMA_PATH     := ${WORKING_DIR}/provider/cmd/${PROVIDER}/schema.json

GOPATH          := $(shell go env GOPATH)


codegen: # Required by CI
	rm -rf ${WORKING_DIR}/bin/${CODEGEN}
	cd provider/cmd/${CODEGEN} && go build -o ${WORKING_DIR}/bin/${CODEGEN}

provider: build_provider # Required by CI
test_provider: # Required by CI
generate_schema: # Required by CI

generate:: gen_go_sdk gen_dotnet_sdk gen_nodejs_sdk gen_python_sdk generate_java

build:: build_provider build_dotnet_sdk build_nodejs_sdk build_python_sdk

install:: install_provider install_dotnet_sdk install_nodejs_sdk


# Provider

build_provider::
	rm -rf ${WORKING_DIR}/bin/${PROVIDER}
	cd provider/cmd/${PROVIDER} && VERSION=${VERSION_GENERIC} SCHEMA=${SCHEMA_PATH} go generate main.go
	cd provider/cmd/${PROVIDER} && go build -o ${WORKING_DIR}/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION_GENERIC}" .

install_provider:: build_provider
	cp ${WORKING_DIR}/bin/${PROVIDER} ${GOPATH}/bin


# Go SDK

gen_go_sdk::
	rm -rf sdk/go
	bin/${CODEGEN} go sdk/go ${SCHEMA_PATH}
build_go_sdk::
generate_go: gen_go_sdk # Required by CI
build_go: # Required by CI
install_go_sdk:: # Required by CI


# .NET SDK

gen_dotnet_sdk::
	rm -rf sdk/dotnet
	bin/${CODEGEN} dotnet sdk/dotnet ${SCHEMA_PATH}

build_dotnet_sdk:: DOTNET_VERSION := $(shell pulumictl convert-version --language dotnet -v "$(VERSION_GENERIC)")
build_dotnet_sdk:: gen_dotnet_sdk
	cd sdk/dotnet/ && \
		echo "${DOTNET_VERSION}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}

install_dotnet_sdk:: # Required by CI
	rm -rf ${WORKING_DIR}/nuget
	mkdir -p ${WORKING_DIR}/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

generate_dotnet: gen_dotnet_sdk # Required by CI
build_dotnet: build_dotnet_sdk # Required by CI

# Node.js SDK

gen_nodejs_sdk::
	rm -rf sdk/nodejs
	bin/${CODEGEN} nodejs sdk/nodejs ${SCHEMA_PATH}
	# HACKHACK: work around https://github.com/pulumi/pulumi/issues/7979:
	find sdk/nodejs -name "*.ts" -exec sed -i.bak \
		's/pulumiKubernetes\.types\.input\.\([a-zA-Z0-9]*\)\.\([a-zA-Z0-9]*\)\.\([a-zA-Z]*\)Args/pulumiKubernetes.types.input.\1.\2.\3/g' \
			'{}' \;

build_nodejs_sdk:: NODE_VERSION := $(shell pulumictl convert-version --language javascript -v "$(VERSION_GENERIC)")
build_nodejs_sdk:: gen_nodejs_sdk
	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc --version && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak -e "s/\$${VERSION}/$(NODE_VERSION)/g" ./bin/package.json && \
		rm ./bin/package.json.bak

generate_nodejs: gen_nodejs_sdk # Required by CI
build_nodejs: build_nodejs_sdk # Required by CI
install_nodejs_sdk:: # Required by CI
	yarn unlink ${PACK} || true
	yarn link --cwd ${WORKING_DIR}/sdk/nodejs/bin


# Python SDK

gen_python_sdk::
	rm -rf sdk/python
	bin/${CODEGEN} python sdk/python ${SCHEMA_PATH}
	cp ${WORKING_DIR}/README.md sdk/python

build_python_sdk:: PYPI_VERSION := $(shell pulumictl convert-version --language python -v "$(VERSION_GENERIC)")
build_python_sdk:: gen_python_sdk
	cd sdk/python/ && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION_GENERIC)"/g' ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && python3 setup.py build sdist

generate_python: gen_python_sdk # Required by CI
build_python: build_python_sdk # Required by CI
install_python_sdk:: # Required by CI

# Java SDK

generate_java: # Required by CI
	pulumi package gen-sdk ${SCHEMA_PATH} -o sdk --language java
	cp ${WORKING_DIR}/README.md sdk/java
build_java: # Required by CI
	cd sdk/java && gradle --console=plain build

install_java_sdk: # Required by CI
