package jumpcloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Define cached version of getJumpCloudOrganization
// by default, WithCache cached the data per connection
// if no argument is passed in WithCache, the cache key will be in the format of <function_name>-<connection_name>
var getOrganization = plugin.HydrateFunc(getOrganizationUncached).WithCache()

// returns details about the JumpCloud Organization
func getOrganizationUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// Trace logging to debug cache and execution flows
	plugin.Logger(ctx).Trace("getOrganizationUncached", "status", "starting", "connection_name", d.Connection.Name)

	// Get jumpcloud config
	jumpCloudConfig := GetConfig(d.Connection)
	if jumpCloudConfig.OrgID != nil {
		return *jumpCloudConfig.OrgID, nil
	}

	return nil, nil
}
