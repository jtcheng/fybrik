#!/usr/bin/env bash
# Copyright 2020 The Kubernetes Authors.
# SPDX-License-Identifier: Apache-2.0

source ./common.sh

header_text "Checking for bin/etcd"
[[ -f bin/etcd ]] && exit 0

header_text "Installing bin/etcd"
mkdir -p ./bin

K8S_VERSION=1.22.1
curl -sSLo envtest-bins.tar.gz "https://storage.googleapis.com/kubebuilder-tools/kubebuilder-tools-${K8S_VERSION}-$(go env GOOS)-$(go env GOARCH).tar.gz"

tar -zvxf envtest-bins.tar.gz
mv kubebuilder/bin/* bin
rm envtest-bins.tar.gz
rm -r kubebuilder