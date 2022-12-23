package jumpcloud

import (
	"context"

	v1 "github.com/Subhajit97/jcapi-go/v1"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJumpCloudApplication(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jumpcloud_application",
		Description: "JumpCloud Application",
		List: &plugin.ListConfig{
			Hydrate: listJumpCloudApplications,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getJumpCloudApplication,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the application.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An uniquely identifier for the application.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "display_name",
				Description: "The display name of the application.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "display_label",
				Description: "The name of the application to display.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization",
				Description: "The name of the JumpCloud organization where the application is created.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "beta",
				Description: "If true, the application is in beta.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "learn_more",
				Description: "Specifies the link where you can find more information related to the application.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sso_url",
				Description: "The SSO URL suffix to use.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "config",
				Description: "Specifies the application configuration.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "user_groups",
				Description: "Specifies the application configuration.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getJumpCloudApplicationGroupAssociation,
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

func listJumpCloudApplications(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_application.listJumpCloudApplications", "connection_error", err)
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

	// Count the number of resource returned by the API.
	// Set the value to 0.
	resourceCount := 0

	for {
		applicationList, _, err := client.ApplicationsApi.ApplicationsList(ctx, "application/json", "application/json", localVarOptionals)
		if err != nil {
			plugin.Logger(ctx).Error("jumpcloud_application.listJumpCloudApplications", "query_error", err)
			return nil, err
		}

		for _, application := range applicationList.Results {
			// Increase the resource count by 1
			resourceCount++

			d.StreamListItem(ctx, application)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if no data
		if len(applicationList.Results) == 0 {
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

func getJumpCloudApplication(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_application.getJumpCloudApplication", "connection_error", err)
		return nil, err
	}
	applicationID := d.EqualsQualString("id")

	// Required quals cannot be empty
	if applicationID == "" {
		return nil, nil
	}

	data, resp, err := client.ApplicationsApi.ApplicationsGet(ctx, applicationID, nil)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_application.getJumpCloudApplication", "query_error", err)

		// Ignore if resource not found error
		if resp.StatusCode == 404 {
			return nil, nil
		}

		// Else return the error
		return nil, err
	}

	return data, nil
}

func getJumpCloudApplicationGroupAssociation(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV2Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_application.getJumpCloudApplicationGroupAssociation", "connection_error", err)
		return nil, err
	}

	var applicationID string
	if h.Item != nil {
		applicationID = h.Item.(v1.Application).Id
	} else {
		applicationID = d.EqualsQualString("id")
	}

	// Required quals cannot be empty
	if applicationID == "" {
		return nil, nil
	}

	data, _, err := client.ApplicationsApi.GraphApplicationTraverseUserGroup(ctx, applicationID, "application/json", "application/json", nil)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_application.getJumpCloudApplicationGroupAssociation", "query_error", err)
		return nil, err
	}

	return data, nil
}
