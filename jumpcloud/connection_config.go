package jumpcloud

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type jumpcloudConfig struct {
	APIKey *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &jumpcloudConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) jumpcloudConfig {
	if connection == nil || connection.Config == nil {
		return jumpcloudConfig{}
	}
	config, _ := connection.Config.(jumpcloudConfig)
	return config
}
