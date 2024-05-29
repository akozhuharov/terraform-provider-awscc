---
page_title: "awscc_qbusiness_application Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::QBusiness::Application Resource Type
---

# awscc_qbusiness_application (Resource)

Definition of AWS::QBusiness::Application Resource Type

## Example Usage

### Create a Q business application with an existing Identity center instance

```terraform
resource "awscc_qbusiness_application" "example" {
  description                  = "Example QBusiness Application"
  display_name                 = "example_q_app"
  identity_center_instance_arn = data.aws_ssoadmin_instances.example.arns[0]
  attachments_configuration = {
    attachments_control_mode = "ENABLED"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}

data "aws_ssoadmin_instances" "example" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String)

### Optional

- `attachments_configuration` (Attributes) (see [below for nested schema](#nestedatt--attachments_configuration))
- `description` (String)
- `encryption_configuration` (Attributes) (see [below for nested schema](#nestedatt--encryption_configuration))
- `identity_center_instance_arn` (String)
- `role_arn` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `application_arn` (String)
- `application_id` (String)
- `created_at` (String)
- `id` (String) Uniquely identifies the resource.
- `identity_center_application_arn` (String)
- `status` (String)
- `updated_at` (String)

<a id="nestedatt--attachments_configuration"></a>
### Nested Schema for `attachments_configuration`

Required:

- `attachments_control_mode` (String)


<a id="nestedatt--encryption_configuration"></a>
### Nested Schema for `encryption_configuration`

Optional:

- `kms_key_id` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_qbusiness_application.example <resource ID>
```