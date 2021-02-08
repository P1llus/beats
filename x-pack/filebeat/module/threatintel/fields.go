// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package threatintel

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "threatintel", asset.ModuleFieldsPri, AssetThreatintel); err != nil {
		panic(err)
	}
}

// AssetThreatintel returns asset data.
// This is the base64 encoded gzipped contents of module/threatintel.
func AssetThreatintel() string {
	return "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"
}
