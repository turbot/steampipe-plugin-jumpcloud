package jumpcloud

import (
	"context"
	"fmt"
	"os"

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

	// Get the credentials
	apiKey, orgID := getCredentialsByPrecedence(d)

	// No creds
	if apiKey == "" {
		return nil, fmt.Errorf("api_key must be configured")
	}

	if orgID == "" {
		return nil, fmt.Errorf("org_id must be configured")
	}

	config := jcapiv2.NewConfiguration()
	config.AddDefaultHeader("x-api-key", apiKey)
	config.AddDefaultHeader("x-org-id", orgID)

	// Create client
	client := jcapiv2.NewAPIClient(config)

	// save clientOptions in cache
	// data will be cached per connection basis
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

	// Get the credentials
	apiKey, orgID := getCredentialsByPrecedence(d)

	// No creds
	if apiKey == "" {
		return nil, fmt.Errorf("api_key must be configured")
	}

	if orgID == "" {
		return nil, fmt.Errorf("org_id must be configured")
	}

	config := jcapiv1.NewConfiguration()
	config.AddDefaultHeader("x-api-key", apiKey)
	config.AddDefaultHeader("x-org-id", orgID)

	// Create client
	client := jcapiv1.NewAPIClient(config)

	// save clientOptions in cache
	// data will be cached per connection basis
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil
}

/*
Create client for JumpCloud's Organization service API

For a multi-tenant admin, when making API requests to JumpCloud
x-org-id must be passed in the header with a valid organization ID
to which the client would like to make the request.
But Organization APIs doesn't allow organization selection via header, and
returns an error:

Status: 403 Forbidden, Body: {"message":"Forbidden: organization selection not allowed via header","error":"Forbidden: organization selection not allowed via header"}

getOrganizationAPIClient function will only use to create the Organization API client
which will not take any 'x-org-id' header.
*/
func getOrganizationAPIClient(ctx context.Context, d *plugin.QueryData) (*jcapiv1.APIClient, error) {
	// Load clientOptions from cache
	sessionCacheKey := "jumpcloud.apiclient_orgv1"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*jcapiv1.APIClient), nil
	}

	// Get the credentials
	apiKey, _ := getCredentialsByPrecedence(d)

	// No creds
	if apiKey == "" {
		return nil, fmt.Errorf("api_key must be configured")
	}

	config := jcapiv1.NewConfiguration()
	config.AddDefaultHeader("x-api-key", apiKey)

	// Create client
	client := jcapiv1.NewAPIClient(config)

	// save clientOptions in cache
	// data will be cached per connection basis
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil
}

/*
Returns credentials by precedence.

Precedence of credentials:
  - Credentials set in config
  - Value set using JUMPCLOUD_API_KEY, and JUMPCLOUD_ORG_ID env var
*/
func getCredentialsByPrecedence(d *plugin.QueryData) (apiKey string, orgID string) {
	// Get jumpcloud config
	jumpCloudConfig := GetConfig(d.Connection)

	// Check for env vars
	apiKey = os.Getenv("JUMPCLOUD_API_KEY")
	orgID = os.Getenv("JUMPCLOUD_ORG_ID")

	// If credentials set in the config, override it
	if jumpCloudConfig.APIKey != nil {
		apiKey = *jumpCloudConfig.APIKey
	}

	if jumpCloudConfig.OrgID != nil {
		orgID = *jumpCloudConfig.OrgID
	}

	return
}
