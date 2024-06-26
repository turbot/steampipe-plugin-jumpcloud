package jumpcloud

import (
	"context"
	"fmt"

	v1 "github.com/Subhajit97/jcapi-go/v1"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJumpCloudDeviceLinuxPackage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jumpcloud_device_linux_package",
		Description: "JumpCloud Linux Device Package",
		List: &plugin.ListConfig{
			ParentHydrate: listJumpCloudDevices,
			Hydrate:       listJumpCloudDeviceLinuxPackages,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "device_id", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "device_id",
				Description: "A JumpCloud generated unique identifier for the device.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SystemId"),
			},
			{
				Name:        "name",
				Description: "The name of the package.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "version",
				Description: "The installed version of the package.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "size",
				Description: "Specifies the size of the package.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "arch",
				Description: "Specifies the package architecture.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "collection_time",
				Description: "The time when the data was collected by the JumpCloud agent.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "install_time",
				Description: "The time when the package was installed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("InstallTime").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "maintainer_or_vendor",
				Description: "The name of the maintainer or vendor of the package.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "mount_namespace_id",
				Description: "The mount name space ID of the package.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "package_format",
				Description: "The format of the package.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "package_group_or_section",
				Description: "Specifies the package group or section.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "pid_with_namespace",
				Description: "Specifies the PID with namespace.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "release_or_revision",
				Description: "Specifies the release or revision of the package.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization_id",
				Description: "Specifies the ID of the organization.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getOrganization,
				Transform:   transform.FromValue(),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listJumpCloudDeviceLinuxPackages(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	device := h.Item.(v1.System)

	// Return nil, if the device_id quals is passed with an empty value.
	if d.EqualsQuals["device_id"] != nil && d.EqualsQualString("device_id") == "" {
		return nil, nil
	}

	// This table lists packages of all the linux devices available in the JumpCloud.
	// Restrict the table by passing a device ID to list programs for a specific linux device.
	if d.EqualsQualString("device_id") != "" && device.Id != d.EqualsQualString("device_id") {
		return nil, nil
	}

	// Create client
	client, err := getV2Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_device_linux_package.listJumpCloudDeviceLinuxPackages", "connection_error", err)
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

	// Get the required field
	filters := []string{fmt.Sprintf("system_id:eq:%s", device.Id)}
	localVarOptionals["filter"] = filters

	// Count the number of resources returned by the API.
	// Set the value to 0.
	resourceCount := 0

	for {
		packages, _, err := client.SystemInsightsApi.SysteminsightsListLinuxPackages(ctx, "application/json", "application/json", localVarOptionals)
		if err != nil {
			plugin.Logger(ctx).Error("jumpcloud_device_linux_package.listJumpCloudDeviceLinuxPackages", "query_error", err)
			return nil, err
		}

		for _, i := range packages {
			// Increase the resource count by 1
			resourceCount++

			d.StreamListItem(ctx, i)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all the data has been processed
		if len(packages) == 0 {
			break
		}

		// Else, set the skip param to list remaining resources.
		// The attribute skip will skip the first n resources that are already listed,
		// and start from the immediate next to list the remaining resources.
		localVarOptionals["skip"] = int32(resourceCount)
	}

	return nil, nil
}
