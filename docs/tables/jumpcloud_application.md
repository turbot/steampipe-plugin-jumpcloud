# Table: jumpcloud_application

JumpCloud lets you streamline user provisioning, management, and access to applications with features like Single Sign On (SSO) and Just-In-Time (JIT) provisioning. The table `jumpcloud_application` will list all the applications integrated with the JumpCloud.

## Examples

### Basic info

```sql
select
  name,
  id,
  display_label,
  sso_url
from
  jumpcloud_application;
```

### List of users who can can access the applications

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
