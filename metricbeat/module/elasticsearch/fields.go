// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package elasticsearch

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "elasticsearch", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}