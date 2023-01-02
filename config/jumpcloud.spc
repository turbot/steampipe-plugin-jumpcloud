connection "jumpcloud" {
  plugin = "jumpcloud"

  # The admin API key to access JumpCloud resources.
  api_key = "YOUR_API_KEY"

  # The JumpCloud organization ID to which you would like to make the request.
  # It is required for all multi-tenant admins when making API requests to JumpCloud.
  org_id = "ORGANIZATION_ID"
}