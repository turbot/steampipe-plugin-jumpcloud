package jumpcloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJumpCloudDevice(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jumpcloud_device",
		Description: "JumpCloud Device",
		List: &plugin.ListConfig{
			Hydrate: listJumpCloudDevices,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getJumpCloudDevice,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "display_name",
				Description: "Display name of the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "A unique identifier JumpCloud generated for the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "os",
				Description: "The operating system installed in the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "active",
				Description: "If true, the device is active.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "created",
				Description: "The time when the device was enrolled.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "allow_multi_factor_authentication",
				Description: "If true, MFA is enabled for the device.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "last_contact",
				Description: "The time when the device was scanned.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "agent_version",
				Description: "The JumCloud agent version installed on the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "allow_public_key_authentication",
				Description: "If true, the device allows public key authentication.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "allow_ssh_password_authentication",
				Description: "If true, the device allows SSH password authentication.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "allow_ssh_root_login",
				Description: "If true, the device allows root login using SSH.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "amazon_instance_id",
				Description: "The amazon instance ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AmazonInstanceID").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "arch",
				Description: "The CPU that a Linux distribution runs on.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "file_system",
				Description: "The device's file system.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "has_service_account",
				Description: "If true, device has service accounts.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "hostname",
				Description: "The hostname of the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "modify_sshd_config",
				Description: "If true, device allows to modify the SSHD config file.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("ModifySSHDConfig").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "os_family",
				Description: "Specifies the OS family.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "remote_ip",
				Description: "The remote IP address of the device.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("RemoteIP").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "serial_number",
				Description: "The serial number of the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ssh_root_enabled",
				Description: "If true, device allowed to perform SSH on root.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "system_timezone",
				Description: "Specifies the area timezone of the device.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "system_token",
				Description: "Specifies the system token used to enroll the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "template_name",
				Description: "Specifies the device template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "version",
				Description: "The current OS version installed on the device.",
				Type:        proto.ColumnType_STRING,
			},

			// JSON columns
			{
				Name:        "connection_history",
				Description: "The device connection history.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "fde",
				Description: "Indicates if the full disk encryption is active in the system.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "mdm",
				Description: "Specifies the mobile device management (MDM) configuration where the device is enrolled.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "network_interfaces",
				Description: "Specifies the list of network interface.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "service_account_state",
				Description: "Specifies the service account state information.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "sshd_params",
				Description: "Specifies the list of SSHD config params of the device.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "system_insights",
				Description: "Specifies the system_insights.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "A list of tags assigned to the device.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "username_hashes",
				Description: "Specifies the username hashes of the device.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "user_metrics",
				Description: "Specifies a list of user metrics.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "organization_id",
				Description: "Specifies the ID of the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Organization"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DisplayName"),
			},
		},
	}
}

//// LIST FUNCTION

func listJumpCloudDevices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_device.listJumpCloudDevices", "connection_error", err)
		return nil, err
	}

	localVarOptionals := map[string]interface{}{}

	// Limit indicates the number of records to return at once.
	// By default the limit is set to 10 by the SDK.
	// If the required limit is less than the default value,
	// update the default value to use it.
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < 10 {
			l := int32(*limit)
			localVarOptionals["limit"] = l
		}
	}

	// Count the number of resources returned by the API.
	// Set the value to 0.
	resourceCount := 0

	for {
		systems, _, err := client.SystemsApi.SystemsList(ctx, "application/json", "application/json", localVarOptionals)
		if err != nil {
			plugin.Logger(ctx).Error("jumpcloud_device.listJumpCloudDevices", "query_error", err)
			return nil, err
		}

		for _, system := range systems.Results {
			// Increase the resource count by 1
			resourceCount++

			d.StreamListItem(ctx, system)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all the data has been processed
		if resourceCount >= int(systems.TotalCount) {
			break
		}

		// Else, set the skip param to list remaining resources.
		// The attribute skip will skip the first n resources that are already listed,
		// and start from the immediate next to list the remaining resources.
		localVarOptionals["skip"] = int32(resourceCount)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getJumpCloudDevice(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_device.getJumpCloudDevice", "connection_error", err)
		return nil, err
	}
	userID := d.EqualsQualString("id")

	// Required quals cannot be empty
	if userID == "" {
		return nil, nil
	}

	data, resp, err := client.SystemsApi.SystemsGet(ctx, userID, "application/json", "application/json", nil)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_device.getJumpCloudDevice", "query_error", err)

		// Ignore if resource not found error
		if resp.StatusCode == 404 {
			return nil, nil
		}

		// Else return the error
		return nil, err
	}

	return data, nil
}
