# Table: jumpcloud_radius_server

JumpCloud's cloud-based RADIUS service extends your organization's user JumpCloud credentials to your WiFi and other resources that support the RADIUS protocol.

## Examples

### Basic info

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

### List all groups that can access the server

```sql
select
  s.name,
  s.id,
  s.organization,
  s.network_source_ip,
  ug.name as group_name
from
  jumpcloud_radius_server as s,
  jsonb_array_elements(groups) as g
  left join jumpcloud_user_group as ug on g ->> 'id' = ug.id;
```
