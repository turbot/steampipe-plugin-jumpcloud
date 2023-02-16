package jumpcloud

import (
	"context"
	"fmt"
	"strconv"

	v1 "github.com/Subhajit97/jcapi-go/v1"
	v2 "github.com/Subhajit97/jcapi-go/v2"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJumpCloudDeviceMacOSApp(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jumpcloud_device_macos_app",
		Description: "JumpCloud MacOS Device App",
		List: &plugin.ListConfig{
			ParentHydrate: listJumpCloudDevices,
			Hydrate:       listJumpCloudDeviceMacOSApps,
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
				Description: "The name of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "version",
				Description: "The installed version of the app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("BundleShortVersion"),
			},
			{
				Name:        "last_opened_time",
				Description: "The time when the app was last opened.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.From(validateLastOpenedTimeValue).Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "display_name",
				Description: "The display name of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "path",
				Description: "Specifies the device path where the app was installed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "applescript_enabled",
				Description: "True if the applescript is enabled for the app.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.From(convertStringToBool),
			},
			{
				Name:        "bundle_executable",
				Description: "Specifies the bundle executable.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bundle_identifier",
				Description: "The bundle identifier of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bundle_name",
				Description: "The bundle version of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bundle_package_type",
				Description: "The bundle package type of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bundle_version",
				Description: "The bundle version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "category",
				Description: "The app category.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "collection_time",
				Description: "The time when the data was collected by the JumpCloud agent.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "compiler",
				Description: "The app complier.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "copyright",
				Description: "The copyright information of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "development_region",
				Description: "The development region of the app.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "element",
				Description: "Specifies the app element.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "environment",
				Description: "The app environment.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "info_string",
				Description: "The app information.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "minimum_system_version",
				Description: "The minimum OS version required to install the app.",
				Type:        proto.ColumnType_STRING,
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

func listJumpCloudDeviceMacOSApps(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	device := h.Item.(v1.System)

	// Return nil, if the device_id quals is passed with an empty value.
	if d.EqualsQualString("device_id") == "" {
		return nil, nil
	}

	// This table lists applications of all the devices available in the JumpCloud.
	// Restrict the table by passing a device ID to list applications for a specific device.
	if d.EqualsQualString("device_id") != "" && device.Id != d.EqualsQualString("device_id") {
		return nil, nil
	}

	// Create client
	client, err := getV2Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_device_macos_app.listJumpCloudDeviceMacOSApps", "connection_error", err)
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
		apps, _, err := client.SystemInsightsApi.SysteminsightsListApps(ctx, "application/json", "application/json", localVarOptionals)
		if err != nil {
			plugin.Logger(ctx).Error("jumpcloud_device_macos_app.listJumpCloudDeviceMacOSApps", "query_error", err)
			return nil, err
		}

		for _, app := range apps {
			// Increase the resource count by 1
			resourceCount++

			d.StreamListItem(ctx, app)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all the data has been processed
		if len(apps) == 0 {
			break
		}

		// Else, set the skip param to list remaining resources.
		// The attribute skip will skip the first n resources that are already listed,
		// and start from the immediate next to list the remaining resources.
		localVarOptionals["skip"] = int32(resourceCount)
	}

	return nil, nil
}

//// TRANSFORM FUNCTIONS

func convertStringToBool(_ context.Context, d *transform.TransformData) (interface{}, error) {
	app := d.HydrateItem.(v2.SystemInsightsApps)
	if app.ApplescriptEnabled != "" {
		formattedValue, err := strconv.Atoi(app.ApplescriptEnabled)
		if err != nil {
			return nil, fmt.Errorf("failed to convert applescript_enabled value to integer: %v", err)
		}
		return formattedValue != 0, nil
	}

	return nil, nil
}

func validateLastOpenedTimeValue(_ context.Context, d *transform.TransformData) (interface{}, error) {
	app := d.HydrateItem.(v2.SystemInsightsApps)

	// If the app is never opened, the API returns the lastOpenedTime as -1.
	// Return nil, if the value is -1.
	if app.LastOpenedTime == -1 {
		return nil, nil
	}
	return app.LastOpenedTime, nil
}
