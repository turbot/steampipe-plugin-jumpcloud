package jumpcloud

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	jcapiv1 "github.com/Subhajit97/jcapi-go/v1"
	jcapiv2 "github.com/Subhajit97/jcapi-go/v2"
)

// Create service client for JumpCloud's V2 API
func getV2Client(ctx context.Context, d *plugin.QueryData) (*jcapiv2.APIClient, error) {
	// Load clientOptions from cache
	sessionCacheKey := "jumpcloud.apiclient_v2"
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

// Create service client for JumpCloud's V1 API
func getV1Client(ctx context.Context, d *plugin.QueryData) (*jcapiv1.APIClient, error) {
	// Load clientOptions from cache
	sessionCacheKey := "jumpcloud.apiclient_v1"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*jcapiv1.APIClient), nil
	}

	// Get jumpcloud config
	jumpcloudConfig := GetConfig(d.Connection)

	// No creds
	if jumpcloudConfig.APIKey == nil {
		return nil, fmt.Errorf("api_key must be passed in the config")
	}

	config := jcapiv1.NewConfiguration()
	config.AddDefaultHeader("x-api-key", *jumpcloudConfig.APIKey)

	// Create client
	client := jcapiv1.NewAPIClient(config)

	// save clientOptions in cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil
}
