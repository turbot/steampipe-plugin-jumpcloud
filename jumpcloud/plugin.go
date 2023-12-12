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
		},
		TableMap: map[string]*plugin.Table{
			"jumpcloud_application":            tableJumpCloudApplication(ctx),
			"jumpcloud_device":                 tableJumpCloudDevice(ctx),
			"jumpcloud_device_linux_package":   tableJumpCloudDeviceLinuxPackage(ctx),
			"jumpcloud_device_macos_app":       tableJumpCloudDeviceMacOSApp(ctx),
			"jumpcloud_device_windows_program": tableJumpCloudDeviceWindowsProgram(ctx),
			"jumpcloud_organization":           tableJumpCloudOrganization(ctx),
			"jumpcloud_radius_server":          tableJumpCloudRadiusServer(ctx),
			"jumpcloud_user":                   tableJumpCloudUser(ctx),
			"jumpcloud_user_group":             tableJumpCloudUserGroup(ctx),
		},
	}

	return p
}
