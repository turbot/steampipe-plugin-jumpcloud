package jumpcloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJumpCloudUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jumpcloud_user",
		Description: "JumpCloud User",
		List: &plugin.ListConfig{
			Hydrate: listJumpCloudUsers,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getJumpCloudUser,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "display_name",
				Description: "Specifies the user's preferred full name.",
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
				Description: "The technical user name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "Specifies the description provided by the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "The users e-mail address, which is also used for log ins. E-mail addresses have to be unique across all JumpCloud accounts, there cannot be two users with the same e-mail address.",
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
				Name:        "suspended",
				Description: "True, if the user account is suspended.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "allow_public_key",
				Description: "If true, public keys are allowed for the user.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "company",
				Description: "The name of the company.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cost_center",
				Description: "Specifies the cost center.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "department",
				Description: "Specifies the department the employee is part of.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "employee_identifier",
				Description: "A unique identifier of the user inside an organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "employee_type",
				Description: "The employment type of the employee.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enable_manage_uid",
				Description: "If true, a managed UID is generated for the user.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "enable_user_portal_multifactor",
				Description: "If true, MFA is enabled while logging in to the user portal.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "external_dn",
				Description: "The external DN provided for the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "external_source_type",
				Description: "Specifies the external source type of the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "externally_managed",
				Description: "Specifies whether the user is externally managed.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "first_name",
				Description: "The user's first name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Firstname"),
			},
			{
				Name:        "job_title",
				Description: "The user's job title.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_name",
				Description: "The user's last name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Lastname"),
			},
			{
				Name:        "location",
				Description: "The user's location.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "middle_name",
				Description: "The user's middle name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Middlename"),
			},
			{
				Name:        "organization",
				Description: "The name of the organization the user is working with.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "password_expiration_date",
				Description: "Specifies the timestamp when the password will expire.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "password_expired",
				Description: "True if the password was expired.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "password_never_expires",
				Description: "If true, the password never gets expired.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "passwordless_sudo",
				Description: "If true, password is not required while using sudo.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "public_key",
				Description: "The public key for the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "totp_enabled",
				Description: "If true, TOTP is enabled for the user.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "mfa",
				Description: "Specifies the MFA configuration for the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "attributes",
				Description: "A list of attributes for the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "ssh_keys",
				Description: "A list of SSH public keys for the user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "A list of tags attached with the user.",
				Type:        proto.ColumnType_JSON,
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

func listJumpCloudUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_user.listJumpCloudUsers", "connection_error", err)
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
			plugin.Logger(ctx).Error("jumpcloud_user.listJumpCloudUsers", "query_error", err)
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

func getJumpCloudUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_user.getJumpCloudUser", "connection_error", err)
		return nil, err
	}
	userID := d.EqualsQualString("id")

	// Required quals cannot be empty
	if userID == "" {
		return nil, nil
	}

	data, resp, err := client.SystemusersApi.SystemusersGet(ctx, userID, "application/json", "application/json", nil)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_user.getJumpCloudUser", "query_error", err)

		// Ignore if resource not found error
		if resp.StatusCode == 404 {
			return nil, nil
		}

		// Else return the error
		return nil, err
	}

	return data, nil
}
