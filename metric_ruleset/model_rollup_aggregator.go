/*
Metric Ingest Ruleset API

Metric ingest ruleset API

API version: 3.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metric_ruleset

import (
	"encoding/json"
)

// RollupAggregator An aggregation rule configured in your ruleset
type RollupAggregator struct {
	// Dimensions you want to keep in the aggregation rule
	DimensionsToKeep []string `json:"dimensionsToKeep,omitempty"`
	// New aggregated metric name
	OutputName string `json:"outputName"`
	// Aggregation rule type
	Type string `json:"type"`
}

func (o RollupAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DimensionsToKeep != nil {
		toSerialize["dimensionsToKeep"] = o.DimensionsToKeep
	}
	if true {
		toSerialize["outputName"] = o.OutputName
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}
