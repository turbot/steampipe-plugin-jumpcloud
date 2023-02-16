# Table: jumpcloud_device_macos_app

The `jumpcloud_device_macos_app` table can be used to query information about all the applications installed in a MacOS device.

To query all applications installed in a Windows or a Linux device, use the `jumpcloud_device_windows_program` and `jumpcloud_device_linux_package` tables respectively.

## Examples

### Basic info

```sql
select
  name,
  version,
  last_opened_time,
  path,
  device_id
from
  jumpcloud_device_macos_app;
```

### Get the device information

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as app,
  a.version as app_version,
  a.last_opened_time
from
  jumpcloud_device_macos_app as a
  join jumpcloud_device as d on d.id = a.device_id;
```

### List devices with tailscale app installed

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as app,
  a.version as app_version,
  a.last_opened_time
from
  jumpcloud_device_macos_app as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name = 'Tailscale.app';
```

### List computers with an older version of zoom app (< 5.12)

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as app,
  a.version as app_version,
  a.last_opened_time
from
  jumpcloud_device_macos_app as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name = 'zoom.us.app'
  and string_to_array(split_part(a.version, ' ', 1), '.')::int[] < string_to_array('5.12', '.')::int[];
```

### List all apps used in last 24 hours

```sql
select
  name,
  version,
  last_opened_time,
  device_id
from
  jumpcloud_device_macos_app
where
  last_opened_time >= (current_timestamp - interval '1 day')
order by
  last_opened_time desc;
```
