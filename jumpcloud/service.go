package jumpcloud

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
)

func getV2Client(ctx context.Context, d *plugin.QueryData) (*jcapiv2.APIClient, error) {
	// Load clientOptions from cache
	sessionCacheKey := "jumpcloud.apiclient"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*jcapiv2.APIClient), nil
	}

	// Get jumpcloud config
	jumpcloudConfig := GetConfig(d.Connection)

	// No creds
	if jumpcloudConfig.APIKey == nil {
		return nil, fmt.Errorf("api_key must be passed in the config")
	}

	config := jcapiv2.NewConfiguration()
	config.AddDefaultHeader("x-api-key", *jumpcloudConfig.APIKey)

	// Create client
	client := jcapiv2.NewAPIClient(config)

	// save clientOptions in cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil
}
