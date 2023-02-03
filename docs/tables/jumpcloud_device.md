# Table: jumpcloud_device

Mac, Windows, or Linux devices are bound to JumpCloud by installing JumpCloud's system agent. Once installed and bound, you can remotely and securely manage these devices and their user accounts, set policies, execute commands, enable MFA, and more.

## Examples

### Basic info

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

### List devices with full disk encryption (FDE) disabled

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

### List all device users

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
