---
title: "Steampipe Table: jumpcloud_user_group - Query JumpCloud User Groups using SQL"
description: "Allows users to query User Groups in JumpCloud, providing detailed information about group IDs, names, and types."
---

# Table: jumpcloud_user_group - Query JumpCloud User Groups using SQL

JumpCloud User Groups is a feature within the JumpCloud Directory-as-a-Service that allows you to manage and organize users. It provides a way to group users based on certain criteria, such as job function, department, or location. JumpCloud User Groups help you manage access control, enforce policies, and streamline user management tasks.

## Table Usage Guide

The `jumpcloud_user_group` table provides insights into User Groups within JumpCloud Directory-as-a-Service. As an IT administrator, explore group-specific details through this table, including group IDs, names, and types. Utilize it to manage access control, enforce policies, and streamline user management tasks.

## Examples

### Basic info
Explore which user groups exist within your system, identifying them by their unique attributes such as name and type. This can be useful for understanding the structure and organization of your user groups.

```sql+postgres
select
  name,
  id,
  type
from
  jumpcloud_user_group;
```

```sql+sqlite
select
  name,
  id,
  type
from
  jumpcloud_user_group;
```

### List groups with samba authentication enabled
Determine the areas in which user groups are using Samba authentication. This can be useful for understanding which groups have this specific feature enabled, providing insights into your network's security protocols.

```sql+postgres
select
  name,
  id,
  type
from
  jumpcloud_user_group
where
  samba_enabled;
```

```sql+sqlite
select
  name,
  id,
  type
from
  jumpcloud_user_group
where
  samba_enabled = 1;
```

### List unused groups
Identify instances where user groups within JumpCloud are not being utilized. This can help streamline your system management by removing or repurposing these unused groups.

```sql+postgres
select
  name,
  id,
  type
from
  jumpcloud_user_group
where
  members is null;
```

```sql+sqlite
select
  name,
  id,
  type
from
  jumpcloud_user_group
where
  members is null;
```