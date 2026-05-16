SHELL := /usr/bin/env bash
ROOT_DIR := $(shell pwd)
SCRIPT1 := $(ROOT_DIR)/scripts/install_via_docker.sh
SCRIPT2 := $(ROOT_DIR)/scripts/install_via_kube.sh

.PHONY: all install build build-images clean

all: build

install:
	@bash "$(SCRIPT1)" install

build:
	@bash "$(SCRIPT1)" build

clean:
	@bash "$(SCRIPT1)" clean

build-images:
	@bash "$(SCRIPT1)" build-images

k8s-install:
	@bash "$(SCRIPT2)" create-cluster
	@bash "$(SCRIPT2)" import-images
	@bash "$(SCRIPT2)" apply

k8s-deploy:
	@bash "$(SCRIPT2)" deploy

delete_cluster:
	@bash "$(SCRIPT2)" delete-cluster