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

package diagnostics

import (
	"encoding/json"
	"time"
)

// TODO, Make this function better, can't really be marshalling and unmarshalling everywhere.
func (d *Diagnostics) getMetrics() {
	metrics := Metrics{}
	response := d.apiRequest("/stats")
	m, _ := json.Marshal(response)
	json.Unmarshal(m, &metrics)
	metrics.Timestamp = time.Now().Format("20060102150405")
	metricsall, _ := json.Marshal(&d.Metrics)
	d.writeToFile(d.DiagFolder, "metrics.json", metricsall)
}
