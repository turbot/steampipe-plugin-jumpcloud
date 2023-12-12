package jumpcloud

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type jumpCloudConfig struct {
	APIKey *string `hcl:"api_key"`
	OrgID  *string `hcl:"org_id"`
}

func ConfigInstance() interface{} {
	return &jumpCloudConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) jumpCloudConfig {
	if connection == nil || connection.Config == nil {
		return jumpCloudConfig{}
	}
	config, _ := connection.Config.(jumpCloudConfig)
	return config
}
