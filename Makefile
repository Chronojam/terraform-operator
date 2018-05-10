ROOT_PACKAGE:=github.com/chronojam/terraform-operator
ABSOLUTE_PATH:=$(GOPATH)/src/$(ROOT_PACKAGE)
PWD:=$(shell pwd)

install-struct-generator:
	go build -o $(PWD)/.build/generator $(ABSOLUTE_PATH)/cmd/generator/*.go
install-k8s-code-generator:
	# https://github.com/kubernetes/code-generator/issues/21
	go get -u k8s.io/apimachinery || true
	go get -u k8s.io/code-generator || true

generate-struct-types: install-struct-generator
	mkdir -p $(ABSOLUTE_PATH)/pkg/apis/aws/v1alpha1
	$(PWD)/.build/generator
	go fmt $(ABSOLUTE_PATH)/pkg/apis/aws/v1alpha1/*.go


generate-k8s-funcs: install-k8s-code-generator
	cd $(GOPATH)/src/k8s.io/code-generator && ./generate-groups.sh all "$(ROOT_PACKAGE)/pkg/client" "$(ROOT_PACKAGE)/pkg/apis" "aws:v1alpha1"

generate-all: generate-struct-types generate-k8s-funcs
