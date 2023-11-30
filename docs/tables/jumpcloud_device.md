---
title: "Steampipe Table: jumpcloud_device - Query JumpCloud Devices using SQL"
description: "Allows users to query Devices in JumpCloud, specifically device details such as hostname, operating system, and associated users, providing insights into device management and user access."
---

# Table: jumpcloud_device - Query JumpCloud Devices using SQL

JumpCloud Devices are the individual computing assets within the JumpCloud platform. These devices can be any type of computing asset, such as servers, desktops, or laptops, and can run various operating systems, including Windows, Mac, and Linux. The JumpCloud platform allows for centralized device management, enabling administrators to track and control user access, apply security policies, and monitor device status.

## Table Usage Guide

The `jumpcloud_device` table provides insights into individual devices within the JumpCloud platform. As a system administrator, you can explore device-specific details through this table, including hostname, operating system, and associated user access. Utilize it to uncover information about devices, such as those with specific operating systems, the users associated with each device, and the overall status of each device.

## Examples

### Basic info
Explore active devices within your network, identifying their operating system, version, and creation date. This query is useful for maintaining an up-to-date inventory and ensuring all devices are running the correct OS versions.

```sql
select
  display_name,
  serial_number,
  os,
  version,
  active,
  created
from
  jumpcloud_device;
```

### List devices with MFA disabled
Explore which devices in your network have multi-factor authentication disabled. This is essential for identifying potential security risks and ensuring all devices comply with security policies.

```sql
select
  display_name,
  serial_number,
  os,
  version,
  active,
  created
from
  jumpcloud_device
where
  not allow_multi_factor_authentication;
```

### List inactive devices
Explore which devices are inactive in your system. This is useful for identifying unused resources and potentially improving system efficiency.

```sql
select
  display_name,
  serial_number,
  os,
  version,
  created
from
  jumpcloud_device
where
  not active;
```

### Get hardware related info
Determine the areas in which hardware-related information is needed to gain insights into device specifications such as CPU type, number of logical and physical cores, hardware model, and version. This can be useful in managing resources, optimizing performance, and planning upgrades.

```sql
select
  display_name,
  serial_number,
  device_info ->> 'cpu_type' as cpu_type,
  device_info ->> 'cpu_logical_cores' as cpu_logical_cores,
  device_info ->> 'cpu_physical_cores' as cpu_physical_cores,
  device_info ->> 'hardware_model' as hardware_model,
  device_info ->> 'hardware_version' as hardware_version
from
  jumpcloud_device;
```

### List devices not allowing SSH password authentication
Explore which devices in your network are configured to disallow SSH password authentication. This is useful for enhancing security by identifying devices that rely on more secure authentication methods.

```sql
select
  display_name,
  serial_number,
  os,
  version,
  created
from
  jumpcloud_device
where
  not allow_ssh_password_authentication;
```

### List devices with Full Disk Encryption (FDE) disabled
Determine the areas in which devices do not have Full Disk Encryption (FDE) enabled. This can be useful in identifying potential security risks and ensuring that all devices comply with company encryption policies.

```sql
select
  display_name,
  serial_number,
  os,
  version,
  active,
  created
from
  jumpcloud_device
where
  not (fde -> 'active')::boolean;
```

### List the user details of devices
Explore active devices by identifying the user details associated with each one. This can be useful for administrators to monitor device usage and track login activity.

```sql
select
  display_name,
  serial_number,
  active,
  u ->> 'userName' as username,
  u ->> 'lastLogin' as last_login_at,
  u ->> 'admin' as is_admin
from
  jumpcloud_device,
  jsonb_array_elements(user_metrics) as u;
```