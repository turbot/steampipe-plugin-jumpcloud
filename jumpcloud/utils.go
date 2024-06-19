package jumpcloud

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Define cached version of getJumpCloudOrganization
// by default, WithCache cached the data per connection
// if no argument is passed in WithCache, the cache key will be in the format of <function_name>-<connection_name>
var getOrganizationMemoized = plugin.HydrateFunc(getOrganizationUncached).Memoize(memoize.WithCacheKeyFunction(getOrganizationCacheKey))

func getOrganization(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getOrganizationMemoized(ctx, d, h)
}

// Build a cache key for the call to getOrganizationIdCacheKey.
func getOrganizationCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getOrganization"
	return key, nil
}

// returns details about the JumpCloud Organization
func getOrganizationUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// Trace logging to debug cache and execution flows
	plugin.Logger(ctx).Debug("getOrganizationUncached", "status", "starting", "connection_name", d.Connection.Name)

	// Get jumpcloud config
	jumpCloudConfig := GetConfig(d.Connection)
	if jumpCloudConfig.OrgID != nil {
		return *jumpCloudConfig.OrgID, nil
	}

	return nil, nil
}

// Parse the input time string in the specified format
func parseAndConvertToUTC(inputTime string) (time.Time, error) {
	var t time.Time
	var err error

	// Try to parse the input time string in different formats
	formats := []string{"20060102", "2006-01-02", "01/02/2006", "1/02/2006", "01/2/2006", "1/2/2006"}
	for _, format := range formats {
		t, err = time.Parse(format, inputTime)
		if err == nil {
			break
		}
	}

	if err != nil {
		return time.Time{}, err
	}

	// Convert the parsed time to UTC format
	utcTime := t.UTC()

	return utcTime, nil
}
