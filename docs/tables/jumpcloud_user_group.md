# Table: jumpcloud_user_group

An user group is a collection of users to perform group-based assignments to resources. User Groups provide your users access to resources.

## Examples

### Basic info

```sql
select
  name,
  id,
  type
from
  jumpcloud_user_group;
```

### List groups with Samba authentication enabled

```sql
select
  name,
  id,
  type
from
  jumpcloud_user_group
where
  samba_enabled;
```

### List unused groups

```sql
select
  name,
  id,
  type
from
  jumpcloud_user_group
where
  members is null;
```
