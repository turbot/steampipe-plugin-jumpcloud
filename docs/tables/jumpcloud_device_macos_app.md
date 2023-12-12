---
title: "Steampipe Table: jumpcloud_device_macos_app - Query JumpCloud macOS Apps using SQL"
description: "Allows users to query macOS Apps in JumpCloud, specifically to retrieve details about each app installed on macOS devices managed by JumpCloud."
---

# Table: jumpcloud_device_macos_app - Query JumpCloud macOS Apps using SQL

JumpCloud is a cloud-based directory service that connects users to their workstations, applications, files, and networks. macOS Apps in JumpCloud refers to the applications installed on macOS devices managed by JumpCloud. This offers an overview of the software landscape across all macOS devices in an organization.

## Table Usage Guide

The `jumpcloud_device_macos_app` table provides insights into macOS Apps within JumpCloud's directory service. As a system administrator, explore app-specific details through this table, including app names, versions, and associated macOS devices. Utilize it to uncover information about apps, such as their distribution across devices, outdated versions, and the need for updates or replacements.

**Important Notes**
- To query all applications installed in a Windows or a Linux device, use the `jumpcloud_device_windows_program` and `jumpcloud_device_linux_package` tables respectively.

## Examples

### Basic info
Explore the details of macOS applications across devices to gain insights into application usage patterns, such as the last time an application was opened. This can be particularly useful for IT administrators to understand software usage and manage resources efficiently.

```sql+postgres
select
  name,
  version,
  last_opened_time,
  path,
  device_id
from
  jumpcloud_device_macos_app;
```

```sql+sqlite
select
  name,
  version,
  last_opened_time,
  path,
  device_id
from
  jumpcloud_device_macos_app;
```

### Get the device information
Explore which applications are installed on specific devices, including their versions and last accessed times. This can help in maintaining up-to-date software across all devices.

```sql+postgres
select
  d.display_name as device_name,
  d.serial_number,
  a.name as app,
  a.version as app_version,
  a.last_opened_time
from
  jumpcloud_device_macos_app as a
  join jumpcloud_device as d on d.id = a.device_id;
```

```sql+sqlite
select
  d.display_name as device_name,
  d.serial_number,
  a.name as app,
  a.version as app_version,
  a.last_opened_time
from
  jumpcloud_device_macos_app as a
  join jumpcloud_device as d on d.id = a.device_id;
```

### List devices with tailscale app installed
Determine the devices that have the Tailscale app installed to manage and monitor software versions and usage. This is beneficial for ensuring software compliance and identifying potential security risks.

```sql+postgres
select
  d.display_name as device_name,
  d.serial_number,
  a.name as app,
  a.version as app_version,
  a.last_opened_time
from
  jumpcloud_device_macos_app as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name = 'Tailscale.app';
```

```sql+sqlite
select
  d.display_name as device_name,
  d.serial_number,
  a.name as app,
  a.version as app_version,
  a.last_opened_time
from
  jumpcloud_device_macos_app as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name = 'Tailscale.app';
```

### List computers with an older version of zoom app (< 5.12)
This query is useful to identify which computers are running an outdated version of the Zoom app, specifically versions older than 5.12. This information can help in maintaining software updates and ensuring all devices are running the most secure and efficient version of the application.

```sql+postgres
select
  d.display_name as device_name,
  d.serial_number,
  a.name as app,
  a.version as app_version,
  a.last_opened_time
from
  jumpcloud_device_macos_app as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name = 'zoom.us.app'
  and string_to_array(split_part(a.version, ' ', 1), '.')::int[] < string_to_array('5.12', '.')::int[];
```

```sql+sqlite
Error: SQLite does not support string_to_array and split functions.
```

### List all apps used in last 24 hours
Discover the applications that have been accessed in the last 24 hours. This can help in monitoring user activity and ensuring software usage compliance.

```sql+postgres
select
  name,
  version,
  last_opened_time,
  device_id
from
  jumpcloud_device_macos_app
where
  last_opened_time >= (current_timestamp - interval '1 day')
order by
  last_opened_time desc;
```

```sql+sqlite
select
  name,
  version,
  last_opened_time,
  device_id
from
  jumpcloud_device_macos_app
where
  last_opened_time >= datetime('now', '-1 day')
order by
  last_opened_time desc;
```