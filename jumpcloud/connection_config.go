package jumpcloud

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
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
	if connection == nil || connection.GetConfig() == nil {
		return jumpCloudConfig{}
	}
	config, _ := connection.GetConfig().(jumpCloudConfig)
	return config
}
