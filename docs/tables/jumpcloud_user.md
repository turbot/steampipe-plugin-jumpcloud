---
title: "Steampipe Table: jumpcloud_user - Query JumpCloud Users using SQL"
description: "Allows users to query JumpCloud Users, providing details about each user's profile, status, and associated metadata."
---

# Table: jumpcloud_user - Query JumpCloud Users using SQL

JumpCloud is a cloud-based directory service that connects users to their workstations, applications, files, and networks. It is designed to control and manage user access to both internal and external IT resources such as WiFi and VPN networks, servers, and web applications. JumpCloud supports various platforms including Mac, Windows, and Linux, and offers features such as LDAP-as-a-Service, RADIUS-as-a-Service, device management, and single sign-on.

## Table Usage Guide

The `jumpcloud_user` table provides insights into user profiles within JumpCloud. As a system administrator, explore user-specific details through this table, including profile information, status, and associated metadata. Utilize it to manage and monitor user access to IT resources, ensuring the security and efficiency of your IT environment.

## Examples

### Basic info
Explore which JumpCloud users are activated and when they were created. This can be used to manage user accounts and track their activity.

```sql+postgres
select
  display_name,
  username,
  email,
  activated,
  created
from
  jumpcloud_user;
```

```sql+sqlite
select
  display_name,
  username,
  email,
  activated,
  created
from
  jumpcloud_user;
```

### List suspended users
Discover the segments that contain suspended users to manage system access and maintain security. This helps in identifying potential threats and ensuring only authorized users have access.

```sql+postgres
select
  display_name,
  username,
  email,
  activated,
  created
from
  jumpcloud_user
where
  suspended;
```

```sql+sqlite
select
  display_name,
  username,
  email,
  activated,
  created
from
  jumpcloud_user
where
  suspended = 1;
```

### List users with MFA disabled
Explore which users have not enabled multi-factor authentication (MFA) to identify potential security risks and enforce stronger access controls.

```sql+postgres
select
  display_name,
  username,
  email,
  activated,
  created
from
  jumpcloud_user
where
  mfa -> 'configured' is null
  or not (mfa -> 'configured')::boolean;
```

```sql+sqlite
select
  display_name,
  username,
  email,
  activated,
  created
from
  jumpcloud_user
where
  json_extract(mfa, '$.configured') is null
  or not json_extract(mfa, '$.configured');
```

### List users not associated with any group
Determine the areas in which users are not linked to any group. This is useful to identify potential issues with user management and ensure all users are properly grouped for access control and permissions management.

```sql+postgres
with user_associated_with_groups as (
  select
    distinct member ->> 'id' as user_id
  from
    jumpcloud_user_group,
    jsonb_array_elements(members) as member
)
select
  display_name,
  username,
  email,
  activated,
  created
from
  jumpcloud_user
where id not in (
  select user_id from user_associated_with_groups
);
```

```sql+sqlite
with user_associated_with_groups as (
  select
    distinct json_extract(member.value, '$.id') as user_id
  from
    jumpcloud_user_group,
    json_each(members) as member
)
select
  display_name,
  username,
  email,
  activated,
  created
from
  jumpcloud_user
where jumpcloud_user.id not in (
  select user_id from user_associated_with_groups
);
```