package signalfx

import (
	"context"
	"net/http"
	"testing"

	"github.com/signalfx/signalfx-go/metric_ruleset"
	"github.com/stretchr/testify/assert"
)

func TestCreateMetricRuleset(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(MetricRulesetApiURL, verifyRequest(t, http.MethodPost, true, http.StatusOK, nil, "metric_ruleset/create_ruleset_success.json"))

	dest := metric_ruleset.FULL_FIDELITY
	result, err := client.CreateMetricRuleset(context.Background(), &metric_ruleset.CreateMetricRulesetRequest{
		Name:        "container cpu utilization by realm and service",
		Version:     1,
		Enabled:     true,
		Destination: &dest,
		MetricMatcher: metric_ruleset.MetricMatcher{
			SimpleMetricMatcher: &metric_ruleset.SimpleMetricMatcher{
				MetricName: "container_cpu_utilization",
				Type:       "simple",
			},
		},
		Aggregators: []metric_ruleset.RollupAggregator{
			{
				OutputName:       "cpu_by_realm_service",
				DimensionsToKeep: []string{"sfx_realm", "sfx_service"},
				Type:             "rollup",
			},
		},
	})

	assert.NoError(t, err, "Unexpected error creating metric ruleset")
	assert.Equal(t, "container cpu utilization by realm and service", *result.Name, "Name does not match")
	assert.Equal(t, 1, len(result.Aggregators), "Unexpected length of aggregators array")
	assert.Equal(t, "simple", result.MetricMatcher.SimpleMetricMatcher.Type, "Matcher type does not match expected")
}

func TestGetMetricRuleset(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(MetricRulesetApiURL+"/metric_ruleset_id", verifyRequest(t, http.MethodGet, true, http.StatusOK, nil, "metric_ruleset/get_ruleset_success.json"))

	result, err := client.GetMetricRuleset(context.Background(), "metric_ruleset_id")
	assert.NoError(t, err, "Unexpected error getting metric ruleset")
	assert.Equal(t, "container cpu utilization by realm and service", *result.Name, "Name does not match")
}

func TestUpdateMetricRuleset(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(MetricRulesetApiURL+"/metric_ruleset_id", verifyRequest(t, http.MethodPut, true, http.StatusOK, nil, "metric_ruleset/create_ruleset_success.json"))

	name := "newName"
	dest := metric_ruleset.DROP
	enabled := true
	version := int64(2)
	result, err := client.UpdateMetricRuleset(context.Background(), "metric_ruleset_id", &metric_ruleset.UpdateMetricRulesetRequest{
		Name:        &name,
		Version:     &version,
		Enabled:     &enabled,
		Destination: &dest,
		MetricMatcher: &metric_ruleset.MetricMatcher{
			SimpleMetricMatcher: &metric_ruleset.SimpleMetricMatcher{
				MetricName: "container_cpu_utilization",
				Type:       "simple",
			},
		},
	})
	assert.NoError(t, err, "Unexpected error creating integration")
	assert.Equal(t, "container cpu utilization by realm and service", *result.Name, "Name does not match")
}

func TestDeleteMetricRuleset(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc(MetricRulesetApiURL+"/metric_ruleset_id", verifyRequest(t, http.MethodDelete, true, http.StatusNoContent, nil, ""))

	err := client.DeleteMetricRuleset(context.Background(), "metric_ruleset_id")
	assert.NoError(t, err, "Unexpected error getting metric ruleset")
}
