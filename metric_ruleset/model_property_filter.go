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

// PropertyFilter Filtering rule to include or exclude metrics specified in the Simple Metric Matcher
type PropertyFilter struct {
	// Dimension key you want to match for the selected metric name
	Property *string `json:"property,omitempty"`
	// Dimension values you want to match for the selected metric name
	PropertyValue []string `json:"propertyValue,omitempty"`
	// Specify whether you want the result of the Simple Metric Matcher to include or exclude the results from the ingest ruleset
	NOT *bool `json:"NOT,omitempty"`
}

func (o PropertyFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	if o.PropertyValue != nil {
		toSerialize["propertyValue"] = o.PropertyValue
	}
	if o.NOT != nil {
		toSerialize["NOT"] = o.NOT
	}
	return json.Marshal(toSerialize)
}
