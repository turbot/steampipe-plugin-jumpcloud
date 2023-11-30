---
title: "Steampipe Table: jumpcloud_radius_server - Query JumpCloud Radius Servers using SQL"
description: "Allows users to query Radius Servers in JumpCloud, specifically providing details about each server's ID, name, networks, and other related information."
---

# Table: jumpcloud_radius_server - Query JumpCloud Radius Servers using SQL

JumpCloud Radius Server is a feature within JumpCloud's Directory-as-a-Service platform that allows for secure, centralized authentication and authorization for network access. It provides a way to manage and control user access to networks, ensuring only authorized individuals can access specific network resources. JumpCloud Radius Server helps maintain network security by enforcing access policies and providing detailed logging of access attempts.

## Table Usage Guide

The `jumpcloud_radius_server` table provides insights into Radius Servers within JumpCloud's Directory-as-a-Service platform. As a network administrator, explore server-specific details through this table, including server ID, name, networks, and other related information. Utilize it to monitor and control user access to networks, enforce access policies, and maintain network security.

## Examples

### Basic info
Discover the segments that are linked to your JumpCloud Radius Server. This query can be used to analyze the name, ID, and organization associated with each network source IP, providing insights into your server's connections.

```sql
select
  name,
  id,
  organization,
  network_source_ip
from
  jumpcloud_radius_server;
```

### List servers with MFA disabled
Identify instances where Multi-Factor Authentication (MFA) is disabled on servers to enhance security measures by promptly addressing potential vulnerabilities.

```sql
select
  name,
  id,
  organization,
  network_source_ip
from
  jumpcloud_radius_server
where
  mfa = 'DISABLED';
```

### List all users that can access the server
Discover the segments that can access a specific server. This is useful in maintaining security and managing user access by identifying who has access to your server.

```sql
select
  s.name,
  s.id,
  s.organization,
  s.network_source_ip,
  ug.display_name as user_name
from
  jumpcloud_radius_server as s,
  jsonb_array_elements(users) as u
  left join jumpcloud_user as ug on u ->> 'id' = ug.id;
```