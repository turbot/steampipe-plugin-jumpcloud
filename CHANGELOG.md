## v0.3.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters.
- Recompiled plugin with Go version `1.21`.

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
