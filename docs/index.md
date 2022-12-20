---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/jumpcloud.svg"
brand_color: "#99CCCC"
display_name: "JumpCloud"
short_name: "jumpcloud"
description: "Steampipe plugin to query servers, applications, user groups, and more from your JumpCloud organization."
og_description: "Query JumpCloud with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/jumpcloud-social-graphic.png"
---

# JumpCloud + Steampipe

[JumpCloud](https://jumpcloud.com) provides an open directory platform that helps to unify the technology stack across identity, access, and device management, cost-effectively that doesn't sacrifice security or functionality.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List users with MFA disabled in your JumpCloud organization:

```sql
select
  display_name,
  username,
  email,
  created
from
  jumpcloud_user
where
  mfa is null
  or not (mfa -> 'configured')::boolean;
```

```
+--------------+------------+-----------------+---------------------------+
| display_name | username   | email           | created                   |
+--------------+------------+-----------------+---------------------------+
| John         | johnweb    | john@domain.com | 2022-12-16T15:42:32+05:30 |
| Adam         | cookiesowl | adam@domain.com | 2022-12-16T21:32:45+05:30 |
+--------------+------------+-----------------+---------------------------+
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

All the tables authenticates the requests via a user API key. To locate your API Key:

- Log into the [JumpCloud Admin Console](https://console.jumpcloud.com).
- Go to the username drop down located in the top-right of the Console.
- Retrieve your API key from `My API Key`.

### Configuration

Installing the latest jumpcloud plugin will create a config file (`~/.steampipe/config/jumpcloud.spc`) with a single connection named `jumpcloud`:

```hcl
connection "jumpcloud" {
  plugin = "jumpcloud"

  # The admin API key to access JumpCloud resources.
  api_key = "YOUR_API_KEY"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-jumpcloud
- Community: [Slack Channel](https://steampipe.io/community/join)
