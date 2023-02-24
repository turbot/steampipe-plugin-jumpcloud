package jumpcloud

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/Subhajit97/jcapi-go/v1"
	v2 "github.com/Subhajit97/jcapi-go/v2"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJumpCloudDeviceWindowsProgram(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jumpcloud_device_windows_program",
		Description: "JumpCloud Windows Device Program",
		List: &plugin.ListConfig{
			ParentHydrate: listJumpCloudDevices,
			Hydrate:       listJumpCloudDeviceWindowsPrograms,
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
				Description: "The name of the program.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "version",
				Description: "The installed version of the program.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "collection_time",
				Description: "The time when the data was collected by the JumpCloud agent.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "identifying_number",
				Description: "A system generated unique identifying number for the program.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "install_date",
				Description: "The time when the program was installed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.From(formatInstallDate).Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "install_location",
				Description: "Specifies the path where the program was installed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "install_source",
				Description: "The source of the program.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "publisher",
				Description: "The publisher of the program.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "uninstall_string",
				Description: "The uninstall string of the program.",
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

func listJumpCloudDeviceWindowsPrograms(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	device := h.Item.(v1.System)

	// Return nil, if the device_id quals is passed with an empty value.
	if d.EqualsQuals["device_id"] != nil && d.EqualsQualString("device_id") == "" {
		return nil, nil
	}

	// This table lists programs of all the windows devices available in the JumpCloud.
	// Restrict the table by passing a device ID to list programs for a specific windows device.
	if d.EqualsQualString("device_id") != "" && device.Id != d.EqualsQualString("device_id") {
		return nil, nil
	}

	// Create client
	client, err := getV2Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_device_windows_program.listJumpCloudDeviceWindowsPrograms", "connection_error", err)
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
		programs, _, err := client.SystemInsightsApi.SysteminsightsListPrograms(ctx, "application/json", "application/json", localVarOptionals)
		if err != nil {
			plugin.Logger(ctx).Error("jumpcloud_device_windows_program.listJumpCloudDeviceWindowsPrograms", "query_error", err)
			return nil, err
		}

		for _, program := range programs {
			// Increase the resource count by 1
			resourceCount++

			d.StreamListItem(ctx, program)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if all the data has been processed
		if len(programs) == 0 {
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

func formatInstallDate(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	program := d.HydrateItem.(v2.SystemInsightsPrograms)

	// Return nil, if the install_date is empty
	if program.InstallDate == "" {
		return nil, nil
	}

	// There are some cases where the API returns invalid date range, i.e. 20202130
	// Return nil, if the date range is invalid
	_, err := types.ToTime(program.InstallDate)
	if err != nil {
		return nil, nil
	}

	// As per the API documentation, install_date returns a date string.
	// But the format of the date is not consistent.
	// It can be yyyymmdd, or, yyyy-mm-dd, or mm/dd/yyyy.
	// By default the Steampipe SDK transforms the yyyymmdd and yyyy-mm-dd formats,
	// but for 01/02/2006 format, it returns an error.
	// Below mentioned code snippet will parse the mm/dd/yyyy to a valid UTC format.
	// Parse the date string in mm/dd/yyyy format
	t, err := parseAndConvertToUTC(program.InstallDate)
	if err != nil {
		return nil, err
	}

	return t.Format(time.RFC3339), nil
}
