---
title: "Steampipe Table: jumpcloud_device_windows_program - Query JumpCloud Windows Programs using SQL"
description: "Allows users to query Windows Programs in JumpCloud, specifically the details of each program installed on a Windows device, providing insights into device software inventory."
---

# Table: jumpcloud_device_windows_program - Query JumpCloud Windows Programs using SQL

JumpCloud Windows Programs are an integral part of the JumpCloud Directory-as-a-Service platform. They provide detailed information about every program installed on a Windows device managed by JumpCloud. This includes data about the program's name, version, publisher, and installation date, among other attributes.

## Table Usage Guide

The `jumpcloud_device_windows_program` table provides insights into Windows Programs within JumpCloud Directory-as-a-Service platform. As a system administrator or IT manager, explore program-specific details through this table, including the program's name, version, publisher, and installation date. Utilize it to manage and audit your software inventory, such as identifying outdated software versions, verifying software publishers, and tracking software installation dates.

**Important Notes**
- To query all applications installed in a MacOS or a Linux device, use the `jumpcloud_device_macos_app` and `jumpcloud_device_linux_package` tables respectively.

## Examples

### Basic info
Discover the segments that provide information about installed Windows programs on JumpCloud devices, such as their names, versions, and installation details. This can be useful in managing software inventory and tracking device-specific installations.

```sql
select
  name,
  version,
  install_date,
  install_location,
  device_id
from
  jumpcloud_device_windows_program;
```

### Get the device information
Determine the areas in which specific program versions are installed across devices. This can help in maintaining software consistency and managing updates.

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as program,
  a.version as program_version,
  a.install_date
from
  jumpcloud_device_windows_program as a
  join jumpcloud_device as d on d.id = a.device_id;
```

### List devices with tailscale app installed
Identify devices that have the Tailscale app installed to monitor software usage and version control. This can assist in ensuring all devices are running the latest, most secure version of the application.

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as program,
  a.version as program_version,
  a.install_date
from
  jumpcloud_device_windows_program as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name like 'Tailscale%';
```

### List computers with an older version of zoom app (< 5.12)
This query is used to identify computers that are running an outdated version of the Zoom application. It's beneficial for IT administrators who need to ensure all devices are using the most recent software for security and functionality purposes.

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as program,
  a.version as program_version,
  a.install_date
from
  jumpcloud_device_windows_program as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name ilike 'zoom%'
  and string_to_array(split_part(a.version, ' ', 1), '.')::int[] < string_to_array('5.12', '.')::int[];
```

### List all programs installed in last 24 hours
Explore the recently installed programs on your device within the past day. This can help maintain device security by identifying any potentially harmful or unwanted installations.

```sql
select
  name,
  version,
  install_date,
  device_id
from
  jumpcloud_device_windows_program
where
  install_date >= (current_timestamp - interval '1 day')
order by
  install_date desc;
```