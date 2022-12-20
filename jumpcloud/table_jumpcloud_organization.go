package jumpcloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJumpcloudOrganization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jumpcloud_organization",
		Description: "JumpCloud Organization",
		List: &plugin.ListConfig{
			Hydrate: listJumpcloudOrganizations,
		},
		Columns: []*plugin.Column{
			{
				Name:        "display_name",
				Description: "The name of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The ID of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "logo_url",
				Description: "The organization logo image URL.",
				Type:        proto.ColumnType_STRING,
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

func listJumpcloudOrganizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_organization.listJumpcloudOrganizations", "connection_error", err)
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
		orgs, _, err := client.OrganizationsApi.OrganizationList(ctx, "application/json", "application/json", localVarOptionals)
		if err != nil {
			plugin.Logger(ctx).Error("jumpcloud_organization.listJumpcloudOrganizations", "query_error", err)
			return nil, err
		}

		for _, user := range orgs.Results {
			// Increase the resource count by 1
			resourceCount++

			d.StreamListItem(ctx, user)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if no data
		if len(orgs.Results) == 0 {
			break
		}

		// Else, set the skip param to list remaining resources.
		// The attribute skip will skip the first n resources that are already listed,
		// and start from the immediate next to list the remaining resources.
		localVarOptionals["skip"] = int32(resourceCount)
	}

	return nil, nil
}
