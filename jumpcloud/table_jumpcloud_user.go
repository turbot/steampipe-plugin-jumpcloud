package jumpcloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJumpcloudUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jumpcloud_user",
		Description: "JumpCloud User",
		List: &plugin.ListConfig{
			Hydrate: listJumpcloudUsers,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getJumpcloudUser,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "display_name",
				Description: "Specifies the user’s preferred full name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Displayname"),
			},
			{
				Name:        "id",
				Description: "An unique identifier for the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "username",
				Description: "Specifies the username.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "Specifies the description provided by the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "The comapny email of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created",
				Description: "Specifies the timestamp when the user is created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "account_locked",
				Description: "True, if the user account is locked.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "activated",
				Description: "True, if the user account is active.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "allow_public_key",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "company",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cost_center",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "department",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},

			{
				Name:        "employee_identifier",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "employee_type",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enable_manage_uid",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "enable_user_portal_multifactor",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "external_dn",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "external_source_type",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "externally_managed",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "firstname",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "job_title",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "lastname",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "location",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "middlename",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "password_expiration_date",
				Description: "",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "password_expired",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "password_never_expires",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "passwordless_sudo",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Displayname"),
			},
		},
	}
}

//// LIST FUNCTION

func listJumpcloudUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_user.listJumpcloudUserGroups", "connection_error", err)
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
		users, _, err := client.SystemusersApi.SystemusersList(ctx, "application/json", "application/json", localVarOptionals)
		if err != nil {
			plugin.Logger(ctx).Error("jumpcloud_user.listJumpcloudUserGroups", "query_error", err)
			return nil, err
		}

		for _, user := range users.Results {
			// Increase the resource count by 1
			resourceCount++

			d.StreamListItem(ctx, user)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if no data
		if len(users.Results) == 0 {
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

func getJumpcloudUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_user.getJumpcloudUser", "connection_error", err)
		return nil, err
	}
	userID := d.EqualsQualString("id")

	// Required quals cannot be empty
	if userID == "" {
		return nil, nil
	}

	data, resp, err := client.SystemusersApi.SystemusersGet(ctx, userID, "application/json", "application/json", nil)
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
