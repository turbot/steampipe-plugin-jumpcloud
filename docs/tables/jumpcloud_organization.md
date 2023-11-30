---
title: "Steampipe Table: jumpcloud_organization - Query JumpCloud Organizations using SQL"
description: "Allows users to query JumpCloud Organizations, providing insights into the organization's details including ID, name, and created timestamp."
---

# Table: jumpcloud_organization - Query JumpCloud Organizations using SQL

JumpCloud is a cloud-based directory service that enables IT admins to control user identities and resource access. It provides a centralized way to manage users, systems, and IT resources across a business's entire environment, both on-premises and in the cloud. JumpCloud simplifies user and system management for IT admins in businesses of all sizes, from small startups to large enterprises.

## Table Usage Guide

The `jumpcloud_organization` table provides insights into organizations within JumpCloud. As an IT administrator, explore organization-specific details through this table, including the organization's ID, name, and the timestamp of when it was created. Utilize it to uncover information about organizations, such as the number of users in each organization, the systems they have access to, and the overall structure of your JumpCloud environment.

## Examples

### Basic info
Gain insights into the basic details of your organization, such as its name, unique identifier, logo, and creation date. This can be useful for a general overview or initial audit of your organization's information.

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
Explore which organizations have not set up any payment options yet. This is useful for identifying potential billing issues and ensuring all organizations have a proper payment method configured.

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
Determine whether your organization's password policy meets security standards by ensuring it requires a minimum length of 14 characters. This is beneficial for maintaining robust security and preventing unauthorized access.

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
Explore which user accounts have password policies set to expire within 90 days. This is useful for ensuring adherence to security best practices and mitigating potential vulnerabilities.

```sql
select
  display_name,
  id,
  settings -> 'passwordPolicy' ->> 'passwordExpirationInDays' as password_expiration,
  (settings -> 'passwordPolicy' ->> 'passwordExpirationInDays')::int <= 90 as password_expiration_within_90
from
  jumpcloud_organization;
```