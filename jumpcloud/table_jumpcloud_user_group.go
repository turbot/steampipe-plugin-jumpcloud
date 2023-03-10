package jumpcloud

import (
	"context"

	v2 "github.com/Subhajit97/jcapi-go/v2"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJumpCloudUserGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jumpcloud_user_group",
		Description: "JumpCloud User Group",
		List: &plugin.ListConfig{
			Hydrate: listJumpCloudUserGroups,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getJumpCloudUserGroup,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "Display name of a user group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "A unique identifier for the user group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Type_"),
			},
			{
				Name:        "samba_enabled",
				Description: "If true, samba authentication is configured for the group.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getJumpCloudUserGroup,
				Transform:   transform.FromField("Attributes.SambaEnabled"),
			},
			{
				Name:        "members",
				Description: "A list of the users associated with the group.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getJumpCloudUserGroupMemberships,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "posix_groups",
				Description: "A list of POSIX groups.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getJumpCloudUserGroup,
				Transform:   transform.FromField("Attributes.PosixGroups"),
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

func listJumpCloudUserGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV2Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_user_group.listJumpCloudUserGroups", "connection_error", err)
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
		userGroups, _, err := client.UserGroupsApi.GroupsUserList(ctx, "application/json", "application/json", localVarOptionals)
		if err != nil {
			plugin.Logger(ctx).Error("jumpcloud_user_group.listJumpCloudUserGroups", "query_error", err)
			return nil, err
		}

		for _, userGroup := range userGroups {
			// Increase the resource count by 1
			resourceCount++

			d.StreamListItem(ctx, userGroup)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if no data
		if len(userGroups) == 0 {
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

func getJumpCloudUserGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV2Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_user_group.getJumpCloudUserGroup", "connection_error", err)
		return nil, err
	}

	var userGroupID string
	if h.Item != nil {
		userGroupID = h.Item.(v2.UserGroup).Id
	} else {
		userGroupID = d.EqualsQualString("id")
	}

	// Required quals cannot be empty
	if userGroupID == "" {
		return nil, nil
	}

	data, resp, err := client.UserGroupsApi.GroupsUserGet(ctx, userGroupID, "application/json", "application/json", nil)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_user_group.getJumpCloudUserGroup", "query_error", err)

		// Ignore if resource not found error
		if resp.StatusCode == 404 {
			return nil, nil
		}

		// Else return the error
		return nil, err
	}

	return data, nil
}

func getJumpCloudUserGroupMemberships(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	var userGroupID string
	if h.Item != nil {
		userGroupID = h.Item.(v2.UserGroup).Id
	} else {
		userGroupID = d.EqualsQualString("id")
	}

	// Create client
	client, err := getV2Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_user_group.getJumpCloudUserGroupMemberships", "connection_error", err)
		return nil, err
	}

	var members []v2.GraphObject

	localVarOptionals := map[string]interface{}{}

	// Count the number of resources returned by the API.
	// Set the value to 0.
	resourceCount := 0

	for {
		data, _, err := client.UserGroupMembersMembershipApi.GraphUserGroupMembersList(ctx, userGroupID, "application/json", "application/json", localVarOptionals)
		if err != nil {
			plugin.Logger(ctx).Error("jumpcloud_user_group.getJumpCloudUserGroupMemberships", "query_error", err)
			return nil, err
		}

		for _, i := range data {
			// Increase the resource count by 1
			resourceCount++

			// append associated user details
			members = append(members, *i.To)
		}

		// Return if no data
		if len(data) == 0 {
			break
		}

		// Else, set the skip param to list remaining resources.
		// The attribute skip will skip the first n resources that are already listed,
		// and start from the immediate next to list the remaining resources.
		localVarOptionals["skip"] = int32(resourceCount)
	}

	return members, nil
}
