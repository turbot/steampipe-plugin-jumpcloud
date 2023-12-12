---
title: "Steampipe Table: jumpcloud_device_linux_package - Query JumpCloud Linux Packages using SQL"
description: "Allows users to query JumpCloud Linux Packages, specifically the details of Linux packages installed on a device, providing insights into device software inventory and potential vulnerabilities."
---

# Table: jumpcloud_device_linux_package - Query JumpCloud Linux Packages using SQL

JumpCloud is a cloud-based directory service that allows IT admins to control user identities and resource access. The Linux Packages in JumpCloud provide details about the software packages installed on a Linux device. It helps in maintaining an updated software inventory and identifying potential software vulnerabilities.

## Table Usage Guide

The `jumpcloud_device_linux_package` table provides insights into Linux packages installed on devices managed by JumpCloud. As a system administrator, explore package-specific details through this table, including package names, versions, and installation status. Utilize it to maintain an updated software inventory, identify outdated packages, and uncover potential vulnerabilities due to unpatched or deprecated software.

**Important Notes**
- To query all applications installed in a MacOS or a Windows device, use the `jumpcloud_device_macos_app` and `jumpcloud_device_windows_program` tables respectively.

## Examples

### Basic info
Explore which Linux packages have been installed on your Jumpcloud devices, along with their versions and installation times. This can help in managing device software and identifying any outdated or unnecessary packages.

```sql+postgres
select
  name,
  version,
  install_time,
  size,
  device_id
from
  jumpcloud_device_linux_package;
```

```sql+sqlite
select
  name,
  version,
  install_time,
  size,
  device_id
from
  jumpcloud_device_linux_package;
```

### Get the device information
Explore the installed software packages on your devices. This allows you to understand what applications are installed on each device, their versions, and when they were installed, which can be crucial for managing software updates and ensuring device security.

```sql+postgres
select
  d.display_name as device_name,
  d.serial_number,
  a.name as package_name,
  a.version as package_version,
  a.install_time
from
  jumpcloud_device_linux_package as a
  join jumpcloud_device as d on d.id = a.device_id;
```

```sql+sqlite
select
  d.display_name as device_name,
  d.serial_number,
  a.name as package_name,
  a.version as package_version,
  a.install_time
from
  jumpcloud_device_linux_package as a
  join jumpcloud_device as d on d.id = a.device_id;
```

### List devices with tailscale app installed
Discover the devices that have the Tailscale app installed. This can be useful to assess the spread and usage of the app within your network.

```sql+postgres
select
  d.display_name as device_name,
  d.serial_number,
  a.name as package_name,
  a.version as package_version,
  a.install_time
from
  jumpcloud_device_linux_package as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name ilike 'tailscale%';
```

```sql+sqlite
select
  d.display_name as device_name,
  d.serial_number,
  a.name as package_name,
  a.version as package_version,
  a.install_time
from
  jumpcloud_device_linux_package as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name like 'tailscale%';
```

### List computers with an older version of zoom app (< 5.12)
Determine the areas in which devices are running an outdated version of the Zoom application. This can help in identifying devices that need to be updated for better security and improved features.

```sql+postgres
select
  d.display_name as device_name,
  d.serial_number,
  a.name as package_name,
  a.version as package_version,
  a.install_time
from
  jumpcloud_device_linux_package as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name ilike 'zoom%'
  and string_to_array(split_part(a.version, ' ', 1), '.')::int[] < string_to_array('5.12', '.')::int[];
```

```sql+sqlite
Error: SQLite does not support string_to_array and split functions.
```

### List all packages installed in last 24 hours
Explore the recent system updates by identifying all software packages installed within the last day. This can help in tracking system changes and troubleshooting any issues that may arise due to the new installations.

```sql+postgres
select
  name,
  version,
  install_time,
  device_id
from
  jumpcloud_device_linux_package
where
  install_time >= (current_timestamp - interval '1 day')
order by
  install_time desc;
```

```sql+sqlite
select
  name,
  version,
  install_time,
  device_id
from
  jumpcloud_device_linux_package
where
  install_time >= datetime('now','-1 day')
order by
  install_time desc;
```