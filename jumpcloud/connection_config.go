package jumpcloud

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type jumpCloudConfig struct {
	APIKey *string `cty:"api_key"`
	OrgID  *string `cty:"org_id"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"org_id": {
		Type: schema.TypeString,
	},
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
