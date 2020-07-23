name=plugin-test
version=0.0.1
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

export GOPATH=$(mkfile_path)/../../../

default: build

.PHONY:build build_plugin
build:
	go build -trimpath -o plugin-test .

build_plugin:
	go build -trimpath --buildmode=plugin -o plugins/example.so ./example/.