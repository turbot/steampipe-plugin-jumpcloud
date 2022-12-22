package jumpcloud

import (
	"context"

	v1 "github.com/Subhajit97/jcapi-go/v1"
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
		Get: &plugin.GetConfig{
			Hydrate:    getJumpcloudOrganization,
			KeyColumns: plugin.SingleColumn("id"),
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
			{
				Name:        "created",
				Description: "The date and time when the organization was created.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getJumpcloudOrganization,
			},
			{
				Name:        "has_credit_card",
				Description: "True, if credit card details has been provided for billing.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getJumpcloudOrganization,
			},
			{
				Name:        "has_stripe_customer_id",
				Description: "True, if a Stripe customer ID has been provided..",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getJumpcloudOrganization,
			},
			{
				Name:        "total_billing_estimate",
				Description: "Indicates the estimated billing for the organization.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getJumpcloudOrganization,
			},
			{
				Name:        "entitlement",
				Description: "Specifies the billing entitlement.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getJumpcloudOrganization,
			},
			{
				Name:        "settings",
				Description: "Specifies the organization settings.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getJumpcloudOrganization,
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

		for _, org := range orgs.Results {
			// Increase the resource count by 1
			resourceCount++

			d.StreamListItem(ctx, org)

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

//// HYDRATE FUNCTIONS

func getJumpcloudOrganization(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_organization.getJumpcloudOrganization", "connection_error", err)
		return nil, err
	}

	var organizationID string
	if h.Item != nil {
		organizationID = h.Item.(v1.OrganizationslistResults).Id
	} else {
		organizationID = d.EqualsQualString("id")
	}

	// Required quals cannot be empty
	if organizationID == "" {
		return nil, nil
	}

	data, resp, err := client.OrganizationsApi.OrganizationGet(ctx, organizationID, "application/json", "application/json", nil)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_user.getJumpcloudUser", "query_error", err)

		// Ignore if resource not found error
		if resp.StatusCode == 404 {
			return nil, nil
		}

		// Else return the error
		return nil, err
	}

	return data, nil
}
