![image](https://hub.steampipe.io/images/plugins/turbot/jumpcloud-social-graphic.png)

# JumpCloud Plugin for Steampipe

Use SQL to query infrastructure servers, applications, user groups, and more from your JumpCloud organization.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/jumpcloud)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/jumpcloud/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-jumpcloud/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install jumpcloud
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/jumpcloud#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/jumpcloud#configuration).

Run a query:

```sql
select
  username,
  created,
  email,
  mfa
from
  jumpcloud_user;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-jumpcloud.git
cd steampipe-plugin-jumpcloud
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/jumpcloud.spc
```

Try it!

```shell
steampipe query
> .inspect jumpcloud
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-jumpcloud/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [JumpCloud Plugin](https://github.com/turbot/steampipe-plugin-jumpcloud/labels/help%20wanted)
