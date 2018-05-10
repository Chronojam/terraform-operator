ROOT_PACKAGE:=github.com/chronojam/terraform-operator
ABSOLUTE_PATH:=$(GOPATH)/src/$(ROOT_PACKAGE)
PWD:=$(shell pwd)

install-pluralizer:
	go build -o $(PWD)/.build/plural $(ABSOLUTE_PATH)/cmd/plural/*.go

install-struct-generator:
	go build -o $(PWD)/.build/generator $(ABSOLUTE_PATH)/cmd/generator/*.go

clean-generation:
	rm -rf $(ABSOLUTE_PATH)/pkg/apis/aws/v1alpha1/*.go

generate-struct-types: install-struct-generator
	mkdir -p $(ABSOLUTE_PATH)/pkg/apis/aws/v1alpha1
	$(PWD)/.build/generator
	go fmt $(ABSOLUTE_PATH)/pkg/apis/aws/v1alpha1/*.go

generate-k8s-funcs:
	# https://github.com/kubernetes/code-generator/issues/21
	cd vendor/k8s.io/code-generator && ./generate-groups.sh all "$(ROOT_PACKAGE)/pkg/client" "$(ROOT_PACKAGE)/pkg/apis" "aws:v1alpha1"

generate-stubs: install-pluralizer
	hack/scaffolder.sh

generate-all: clean-generation generate-struct-types generate-k8s-funcs
