# Table: jumpcloud_user

User identities are at the core of JumpCloud. As a core directory, JumpCloud provides centralized, authoritative versions of each employee's identities so they can use a single set of credentials across all their resources.

## Examples

### Basic info

```sql
select
  display_name,
  username,
  email,
  activated,
  created
from
  jumpcloud_user;
```

### List all suspended users

```sql
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

### List all the users for whom MFA is not enabled

```sql
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

### List users not associated with any group

```sql
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
)
```
