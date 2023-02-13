package jumpcloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-jumpcloud"

// Plugin creates this (jumpcloud) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultGetConfig: &plugin.GetConfig{},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"jumpcloud_application":     tableJumpCloudApplication(ctx),
			"jumpcloud_device":          tableJumpCloudDevice(ctx),
			"jumpcloud_organization":    tableJumpCloudOrganization(ctx),
			"jumpcloud_radius_server":   tableJumpCloudRadiusServer(ctx),
			"jumpcloud_user":            tableJumpCloudUser(ctx),
			"jumpcloud_user_group":      tableJumpCloudUserGroup(ctx),
			"jumpcloud_windows_program": tableJumpCloudWindowsProgram(ctx),
		},
	}

	return p
}
