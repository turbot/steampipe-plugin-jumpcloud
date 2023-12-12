---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/jumpcloud.svg"
brand_color: "#14A19C"
display_name: "JumpCloud"
short_name: "jumpcloud"
description: "Steampipe plugin to query servers, applications, user groups, and more from your JumpCloud organization."
og_description: "Query JumpCloud with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/jumpcloud-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# JumpCloud + Steampipe

[JumpCloud](https://jumpcloud.com) provides an open directory platform that helps to unify the technology stack across identity, access, and device management, cost-effectively that doesn't sacrifice security or functionality.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

List JumpCloud user details:

```sql
select
  username,
  created,
  email,
  mfa
from
  jumpcloud_user;
```

```
+------------+---------------------------+------------------------+-----------------------------------------------+
| username   | created                   | email                  | mfa                                           |
+------------+---------------------------+------------------------+-----------------------------------------------+
| johnweb    | 2022-12-16T15:42:32+05:30 | johnweb@example.com    | <null>                                        |
| cookiesowl | 2022-12-19T15:10:02+05:30 | cookiesowl@example.com | {"exclusionUntil":"2022-12-27T02:30:24.498Z"} |
+------------+---------------------------+------------------------+-----------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/jumpcloud/tables)**

## Get started

### Install

Download and install the latest JumpCloud plugin:

```bash
steampipe plugin install jumpcloud
```

### Credentials

| Item        | Description                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | JumpCloud requires an [API token](https://docs.jumpcloud.com/api/2.0/index.html#section/API-Key/Access-Your-API-Key) for all requests.                                                                |
| Permissions | API tokens have the same permissions as the user who creates them, and if the user permissions change, the API token permissions also change.                                                         |
| Radius      | Each connection represents a single JumpCloud Installation.                                                                                                                                           |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/jumpcloud.spc`)<br />2. Credentials specified in environment variables, e.g., `JUMPCLOUD_API_KEY`, `JUMPCLOUD_ORG_ID`. |

### Configuration

Installing the latest jumpcloud plugin will create a config file (`~/.steampipe/config/jumpcloud.spc`) with a single connection named `jumpcloud`:

```hcl
connection "jumpcloud" {
  plugin = "jumpcloud"

  # The admin API key to access JumpCloud resources.
  # This can also be set via the `JUMPCLOUD_API_KEY` environment variable.
  # api_key = "1b234ac9de5f5gh67i89j10k9l366mnop6q965r6"

  # The JumpCloud organization ID to which you would like to make the request.
  # It is required for all multi-tenant admins when making API requests to JumpCloud.
  # This can also be set via the `JUMPCLOUD_ORG_ID` environment variable.
  # org_id = "123a45b6c78d8e9f6gh0769i"
}
```

### Credentials from Environment Variables

The JumpCloud plugin will use the standard JumpCloud environment variables to obtain credentials **only if other arguments (`api_key` and `org_id`) are not specified** in the connection:

```sh
export JUMPCLOUD_API_KEY=1b234ac9de5f5gh67i89j10k9l366mnop6q965r6
export JUMPCLOUD_ORG_ID=123a45b6c78d8e9f6gh0769i
```


