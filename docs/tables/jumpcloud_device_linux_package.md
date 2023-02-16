# Table: jumpcloud_device_linux_package

The `jumpcloud_device_linux_package` table can be used to query information about all the packages installed on a Linux device.

To query all applications installed in a MacOS or a Windows device, use the `jumpcloud_device_macos_app` and `jumpcloud_device_windows_program` tables respectively.

## Examples

### Basic info

```sql
select
  name,
  version,
  install_time,
  size,
  device_id
from
  jumpcloud_device_linux_package;
```

### Get the device information

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as package_name,
  a.version as package_version,
  a.install_time
from
  jumpcloud_device_linux_package as a
  join jumpcloud_device as d on d.id = a.device_id;
```

### List devices with Tailscale app installed

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as package_name,
  a.version as package_version,
  a.install_time
from
  jumpcloud_device_linux_package as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name ilike 'tailscale%';
```

### List computers with an older version of Zoom app (< 5.12)

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as package_name,
  a.version as package_version,
  a.install_time
from
  jumpcloud_device_linux_package as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name ilike 'zoom%'
  and string_to_array(split_part(a.version, ' ', 1), '.')::int[] < string_to_array('5.12', '.')::int[];
```

### List all programs installed in last 24 hours

```sql
select
  name,
  version,
  install_time,
  device_id
from
  jumpcloud_device_linux_package
where
  install_time >= (current_timestamp - interval '1 day')
order by
  install_time desc;
```
