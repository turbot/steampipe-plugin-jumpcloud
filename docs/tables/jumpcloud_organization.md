# Table: jumpcloud_organization

JumpCloud organization table lists all the organizations along with the settings and billing information.

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
  not has_credit_card
  and not has_stripe_customer_id;
```

### Check if password requires minimum length of 14

```sql
select
  display_name,
  id,
  settings -> 'passwordPolicy' ->> 'minLength' as password_min_length,
  (settings -> 'passwordPolicy' ->> 'minLength')::int >= 14 as password_min_length_14_or_greater
from
  jumpcloud_organization;
```

### Check if password expires within 90 days

```sql
select
  display_name,
  id,
  settings -> 'passwordPolicy' ->> 'passwordExpirationInDays' as password_expiration,
  (settings -> 'passwordPolicy' ->> 'passwordExpirationInDays')::int <= 90 as password_expiration_within_90
from
  jumpcloud_organization;
```
