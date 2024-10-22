## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#48](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/48))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#48](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/48))

## v0.4.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#43](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/43))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#43](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/43))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-jumpcloud/blob/main/docs/LICENSE). ([#43](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/43))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#42](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/42))

## v0.3.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#32](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/32))

## v0.3.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#29](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/29))
- Recompiled plugin with Go version `1.21`. ([#29](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/29))

## v0.2.1 [2023-02-25]

_Bug fixes_

- Fixed the `install_date` column in `jumpcloud_device_windows_program` table to correctly convert the install date into UTC format to avoid returning errors on being queried. ([#21](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/21))

## v0.2.0 [2023-02-16]

_What's new?_

- New tables added
  - [jumpcloud_device_linux_package](https://hub.steampipe.io/plugins/turbot/jumpcloud/tables/jumpcloud_device_linux_package) ([#15](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/15))
  - [jumpcloud_device_macos_app](https://hub.steampipe.io/plugins/turbot/jumpcloud/tables/jumpcloud_device_macos_app) ([#13](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/13))
  - [jumpcloud_device_windows_program](https://hub.steampipe.io/plugins/turbot/jumpcloud/tables/jumpcloud_device_windows_program) ([#14](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/14))

_Enhancements_

- Added column `device_info` to `jumpcloud_device` table. ([#18](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/18))

## v0.1.0 [2023-02-04]

_What's new?_

- New tables added
  - [jumpcloud_device](https://hub.steampipe.io/plugins/turbot/jumpcloud/tables/jumpcloud_device`) ([#6](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/6))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.1.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v500-2022-11-16) with many cache improvements. ([#8](https://github.com/turbot/steampipe-plugin-jumpcloud/pull/8))

## v0.0.1 [2023-01-13]

_What's new?_

- New tables added

  - [jumpcloud_application](https://hub.steampipe.io/plugins/turbot/jumpcloud/tables/jumpcloud_application)
  - [jumpcloud_organization](https://hub.steampipe.io/plugins/turbot/jumpcloud/tables/jumpcloud_organization)
  - [jumpcloud_radius_server](https://hub.steampipe.io/plugins/turbot/jumpcloud/tables/jumpcloud_radius_server)
  - [jumpcloud_user](https://hub.steampipe.io/plugins/turbot/jumpcloud/tables/jumpcloud_user)
  - [jumpcloud_user_group](https://hub.steampipe.io/plugins/turbot/jumpcloud/tables/jumpcloud_user_group)
