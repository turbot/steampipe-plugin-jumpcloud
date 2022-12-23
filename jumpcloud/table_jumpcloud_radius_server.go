package jumpcloud

import (
	"context"

	v1 "github.com/Subhajit97/jcapi-go/v1"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableJumpcloudRadiusServer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "jumpcloud_radius_server",
		Description: "JumpCloud RADIUS Server",
		List: &plugin.ListConfig{
			Hydrate: listJumpcloudRadiusServers,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the RADIUS server.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The ID of the server.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "organization",
				Description: "The name of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "network_source_ip",
				Description: "The  IP address from which your organization's traffic will originate",
				Type:        proto.ColumnType_IPADDR,
			},
			{
				Name:        "mfa",
				Description: "The MFA status of the server.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "shared_secret",
				Description: "The character string that is configured on both the client hardware and on the RADIUS server.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_lockout_action",
				Description: "Specifies the action to be performed when the user gets lockout.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_password_expiration_action",
				Description: "Specifies the action to be performed when the user's password got expires.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_groups",
				Description: "A list of user groups associated with the server.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getJumpcloudRadiusServerGroupAssociations,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "users",
				Description: "A list of users associated with the server.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getJumpcloudRadiusServerUserAssociations,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "tag_names",
				Description: "Specifies a list of tag names attached with the server.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags",
				Description: "A list of tags attached with the server.",
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

func listJumpcloudRadiusServers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV1Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_radius_server.listJumpcloudRadiusServers", "connection_error", err)
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
		servers, _, err := client.RadiusServersApi.RadiusServersList(ctx, "application/json", "application/json", localVarOptionals)
		if err != nil {
			plugin.Logger(ctx).Error("jumpcloud_radius_server.listJumpcloudRadiusServers", "query_error", err)
			return nil, err
		}

		for _, user := range servers.Results {
			// Increase the resource count by 1
			resourceCount++

			d.StreamListItem(ctx, user)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// Return if no data
		if len(servers.Results) == 0 {
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

func getJumpcloudRadiusServerGroupAssociations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV2Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_radius_server.getJumpcloudRadiusServerAssociations", "connection_error", err)
		return nil, err
	}
	serverID := h.Item.(v1.Radiusserver).Id

	// Required quals cannot be empty
	if serverID == "" {
		return nil, nil
	}

	data, _, err := client.GraphApi.GraphRadiusServerTraverseUserGroup(ctx, serverID, "application/json", "application/json", nil)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_radius_server.getJumpcloudRadiusServerAssociations", "query_error", err)
		return nil, err
	}

	return data, nil
}

func getJumpcloudRadiusServerUserAssociations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getV2Client(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_radius_server.getJumpcloudRadiusServerUserAssociations", "connection_error", err)
		return nil, err
	}

	var serverID string
	if h.Item != nil {
		serverID = h.Item.(v1.Radiusserver).Id
	} else {
		serverID = d.EqualsQualString("id")
	}

	// Required quals cannot be empty
	if serverID == "" {
		return nil, nil
	}

	data, _, err := client.GraphApi.GraphRadiusServerTraverseUser(ctx, serverID, "application/json", "application/json", nil)
	if err != nil {
		plugin.Logger(ctx).Error("jumpcloud_radius_server.getJumpcloudRadiusServerUserAssociations", "query_error", err)
		return nil, err
	}

	return data, nil
}
