ROOT_PACKAGE:=github.com/chronojam/terraform-operator
ABSOLUTE_PATH:=$(GOPATH)/src/$(ROOT_PACKAGE)
PWD:=$(shell pwd)

install-struct-generator:
	go build -o $(PWD)/.build/generator $(ABSOLUTE_PATH)/cmd/generator/*.go
install-k8s-code-generator:
	go get -u k8s.io/code-generator/..

generate-struct-types: install-struct-generator
	mkdir -p $(ABSOLUTE_PATH)/pkg/apis/aws/v1alpha1
	$(PWD)/.build/generator
	go fmt $(ABSOLUTE_PATH)/pkg/apis/aws/v1alpha1/*.go

generate-k8s-funcs:
	cd $(GOPATH)/src/k8s.io/code-generator && ./generate-groups.sh all "$(ROOT_PACKAGE)/pkg/client" "$(ROOT_PACKAGE)/pkg/apis" "aws:v1alpha1"

generate-all: generate-struct-types generate-k8s-funcs
