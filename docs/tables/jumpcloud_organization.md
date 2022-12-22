# Table: jumpcloud_organization

List all the JumpCloud organizations along with the settings and billing information.

## Examples

### Basic info

```sql
select
  display_name,
  id,
  logo_url,
  created
from
  jumpcloud_organization;
```

### List organizations with no payment options configured

```sql
select
  display_name,
  id,
  logo_url,
  created
from
  jumpcloud_organization
where
  not has_credit_card;
```

### Check if password requires minimum length of 14 or greater

```sql
select
  display_name,
  id,
  settings -> 'passwordPolicy' ->> 'minLength' as password_min_length
from
  jumpcloud_organization;
```

### Check if password expires within 90 days or less

```sql
select
  display_name,
  id,
  settings -> 'passwordPolicy' -> 'passwordExpirationInDays' as password_min_length
from
  jumpcloud_organization;
```
