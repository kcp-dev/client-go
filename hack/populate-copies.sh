#!/usr/bin/env bash

# Copyright 2022 The KCP Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

# This script populates our third_party directory, along with any other files we're
# wholesale copying from the upstream k8s.io/client-go repository. All files retain
# their original author's copyright.

source_dir="$( go list -m -json k8s.io/client-go | jq --raw-output .Dir )"

sink_dir="./third_party/k8s.io/client-go/dynamic/"
mkdir -p "${sink_dir}"
for file in scheme simple; do
  cp "${source_dir}/dynamic/${file}.go" "${sink_dir}"
done

for extension in $( find "${source_dir}/listers" -type f -name '*_expansion.go' ); do
  sink="./clients/${extension##"${source_dir}/"}"
  mkdir -p "$( dirname "${sink}" )"
  cp "${extension}" "${sink}"
done