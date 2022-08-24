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

// UpdateMetricRulesetResponse Response body returned by Update Ingest Ruleset
type UpdateMetricRulesetResponse struct {
	// Aggregation rules in the ruleset
	Aggregators []RollupAggregator `json:"aggregators,omitempty"`
	// User ID of the user who created this ingest ruleset
	Creator *string `json:"creator,omitempty"`
	// Date and time when this ruleset was created
	Created *int64 `json:"created,omitempty"`
	// Metric ingest ruleset description
	Description *string      `json:"description,omitempty"`
	Destination *Destination `json:"destination,omitempty"`
	// Ruleset status
	Enabled *bool `json:"enabled,omitempty"`
	// Ruleset ID
	Id *string `json:"id,omitempty"`
	// ID of the user who last updated the ruleset
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// Time at which this ruleset was last updated
	LastUpdated   *int64         `json:"lastUpdated,omitempty"`
	MetricMatcher *MetricMatcher `json:"metricMatcher,omitempty"`
	// Name of the ruleset
	Name *string `json:"name,omitempty"`
	// Version of the ruleset
	Version *int64 `json:"version,omitempty"`
}

func (o UpdateMetricRulesetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Aggregators != nil {
		toSerialize["aggregators"] = o.Aggregators
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedBy != nil {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.MetricMatcher != nil {
		toSerialize["metricMatcher"] = o.MetricMatcher
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}
