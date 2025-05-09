package signalfx

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/signalfx/signalfx-go/alertmuting"
)

// AlertMutingRuleAPIURL is the base URL for interacting with alert muting rules.
const AlertMutingRuleAPIURL = "/v2/alertmuting"

// CreateAlertMutingRule creates an alert muting rule.
func (c *Client) CreateAlertMutingRule(ctx context.Context, muteRequest *alertmuting.CreateUpdateAlertMutingRuleRequest) (*alertmuting.AlertMutingRule, error) {
	payload, err := json.Marshal(muteRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(ctx, "POST", AlertMutingRuleAPIURL, nil, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = newResponseError(resp, http.StatusCreated); err != nil {
		return nil, err
	}

	finalRule := &alertmuting.AlertMutingRule{}

	err = json.NewDecoder(resp.Body).Decode(finalRule)
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return finalRule, err
}

// DeleteAlertMutingRule deletes an alert muting rule.
func (c *Client) DeleteAlertMutingRule(ctx context.Context, name string) error {
	resp, err := c.doRequest(ctx, "DELETE", AlertMutingRuleAPIURL+"/"+name, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = newResponseError(resp, http.StatusNoContent); err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return nil
}

// GetAlertMutingRule gets an alert muting rule.
func (c *Client) GetAlertMutingRule(ctx context.Context, id string) (*alertmuting.AlertMutingRule, error) {
	resp, err := c.doRequest(ctx, "GET", AlertMutingRuleAPIURL+"/"+id, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = newResponseError(resp, http.StatusOK); err != nil {
		return nil, err
	}

	finalRule := &alertmuting.AlertMutingRule{}

	err = json.NewDecoder(resp.Body).Decode(finalRule)
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return finalRule, err
}

// UpdateAlertMutingRule updates an alert muting rule.
func (c *Client) UpdateAlertMutingRule(ctx context.Context, id string, muteRequest *alertmuting.CreateUpdateAlertMutingRuleRequest) (*alertmuting.AlertMutingRule, error) {
	payload, err := json.Marshal(muteRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(ctx, "PUT", AlertMutingRuleAPIURL+"/"+id, nil, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = newResponseError(resp, http.StatusOK); err != nil {
		return nil, err
	}

	finalRule := &alertmuting.AlertMutingRule{}

	err = json.NewDecoder(resp.Body).Decode(finalRule)
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return finalRule, err
}

// SearchAlertMutingRules searches for alert muting rules given a query string in `name`.
func (c *Client) SearchAlertMutingRules(ctx context.Context, include string, limit int, query string, offset int) (*alertmuting.SearchResult, error) {
	params := url.Values{}
	params.Add("include", include)
	params.Add("limit", strconv.Itoa(limit))
	params.Add("query", query)
	params.Add("offset", strconv.Itoa(offset))

	resp, err := c.doRequest(ctx, "GET", AlertMutingRuleAPIURL, params, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = newResponseError(resp, http.StatusOK); err != nil {
		return nil, err
	}

	finalRules := &alertmuting.SearchResult{}

	err = json.NewDecoder(resp.Body).Decode(finalRules)
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return finalRules, err
}
