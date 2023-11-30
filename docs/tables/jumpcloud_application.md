---
title: "Steampipe Table: jumpcloud_application - Query JumpCloud Applications using SQL"
description: "Allows users to query JumpCloud Applications, providing details about each application including its ID, name, display label, description, and other related information."
---

# Table: jumpcloud_application - Query JumpCloud Applications using SQL

JumpCloud Applications is a part of JumpCloud's Directory-as-a-Service platform that allows the management of user access to various IT resources, including applications. It provides a simplified way to manage user access across multiple applications, ensuring secure and efficient user authentication. It is a key component in implementing single sign-on (SSO) capabilities across an organization's IT environment.

## Table Usage Guide

The `jumpcloud_application` table provides insights into applications within JumpCloud's Directory-as-a-Service platform. As an IT administrator, you can explore application-specific details through this table, including application ID, name, display label, description, and more. Utilize it to manage and monitor user access across multiple applications, ensuring secure and efficient user authentication.

## Examples

### Basic info
Explore the details of your applications managed through JumpCloud, such as their names, IDs, display labels, and SSO URLs. This can be beneficial for auditing purposes, ensuring correct configurations, and maintaining an organized inventory of your applications.

```sql
select
  name,
  id,
  display_label,
  sso_url
from
  jumpcloud_application;
```

### List of users who can can access the application
Identify the users who have permissions to access a specific application. This is useful to manage and monitor user access for security and administrative purposes.

```sql
with application_group_association as (
  select
    name,
    g ->> 'id' as group_id
  from
    jumpcloud_application,
    jsonb_array_elements(user_groups) as g
),
group_user_association as (
  select
    a.name as app_name,
    g.id as group_id,
    g.members
  from
    application_group_association as a
    left join jumpcloud_user_group as g on g.id = a.group_id
)
select
  ga.app_name,
  ga.group_id,
  u.display_name as user_name
from
  group_user_association as ga,
  jsonb_array_elements(members) as m
  left join jumpcloud_user as u on m ->> 'id' = u.id;
```