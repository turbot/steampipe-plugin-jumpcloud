# Table: jumpcloud_device_windows_program

The `jumpcloud_device_windows_program` table can be used to query information about all the programs installed in a Windows device.

To query all applications installed in a MacOS or a Linux device, use the `jumpcloud_device_macos_app` and `jumpcloud_device_linux_package` tables respectively.

## Examples

### Basic info

```sql
select
  name,
  version,
  install_date,
  install_location,
  device_id
from
  jumpcloud_device_windows_program;
```

### Get the device information

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as program,
  a.version as program_version,
  a.install_date
from
  jumpcloud_device_windows_program as a
  join jumpcloud_device as d on d.id = a.device_id;
```

### List devices with tailscale app installed

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as program,
  a.version as program_version,
  a.install_date
from
  jumpcloud_device_windows_program as a
  join jumpcloud_device as d on d.id = a.device_id
where
  a.name like 'Tailscale%';
```

### List computers with an older version of zoom app (< 5.12)

```sql
select
  d.display_name as device_name,
  d.serial_number,
  a.name as program,
  a.version as program_version,
  a.install_date
from
  jumpcloud_device_windows_program as a
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
  install_date,
  device_id
from
  jumpcloud_device_windows_program
where
  install_date >= (current_timestamp - interval '1 day')
order by
  install_date desc;
```
