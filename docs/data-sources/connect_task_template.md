---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_task_template Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Connect::TaskTemplate
---

# awscc_connect_task_template (Data Source)

Data Source schema for AWS::Connect::TaskTemplate



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) The identifier (arn) of the task template.
- `client_token` (String) the client token string in uuid format
- `constraints` (Attributes) The constraints for the task template (see [below for nested schema](#nestedatt--constraints))
- `contact_flow_arn` (String) The identifier of the contact flow.
- `defaults` (Attributes List) (see [below for nested schema](#nestedatt--defaults))
- `description` (String) The description of the task template.
- `fields` (Attributes List) The list of task template's fields (see [below for nested schema](#nestedatt--fields))
- `instance_arn` (String) The identifier (arn) of the instance.
- `name` (String) The name of the task template.
- `status` (String) The status of the task template
- `tags` (Attributes Set) One or more tags. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--constraints"></a>
### Nested Schema for `constraints`

Read-Only:

- `invisible_fields` (Attributes List) The list of the task template's invisible fields (see [below for nested schema](#nestedatt--constraints--invisible_fields))
- `read_only_fields` (Attributes List) The list of the task template's read only fields (see [below for nested schema](#nestedatt--constraints--read_only_fields))
- `required_fields` (Attributes List) The list of the task template's required fields (see [below for nested schema](#nestedatt--constraints--required_fields))

<a id="nestedatt--constraints--invisible_fields"></a>
### Nested Schema for `constraints.invisible_fields`

Read-Only:

- `id` (Attributes) the identifier (name) for the task template field (see [below for nested schema](#nestedatt--constraints--invisible_fields--id))

<a id="nestedatt--constraints--invisible_fields--id"></a>
### Nested Schema for `constraints.invisible_fields.id`

Read-Only:

- `name` (String) The name of the task template field



<a id="nestedatt--constraints--read_only_fields"></a>
### Nested Schema for `constraints.read_only_fields`

Read-Only:

- `id` (Attributes) the identifier (name) for the task template field (see [below for nested schema](#nestedatt--constraints--read_only_fields--id))

<a id="nestedatt--constraints--read_only_fields--id"></a>
### Nested Schema for `constraints.read_only_fields.id`

Read-Only:

- `name` (String) The name of the task template field



<a id="nestedatt--constraints--required_fields"></a>
### Nested Schema for `constraints.required_fields`

Read-Only:

- `id` (Attributes) the identifier (name) for the task template field (see [below for nested schema](#nestedatt--constraints--required_fields--id))

<a id="nestedatt--constraints--required_fields--id"></a>
### Nested Schema for `constraints.required_fields.id`

Read-Only:

- `name` (String) The name of the task template field




<a id="nestedatt--defaults"></a>
### Nested Schema for `defaults`

Read-Only:

- `default_value` (String) the default value for the task template's field
- `id` (Attributes) the identifier (name) for the task template field (see [below for nested schema](#nestedatt--defaults--id))

<a id="nestedatt--defaults--id"></a>
### Nested Schema for `defaults.id`

Read-Only:

- `name` (String) The name of the task template field



<a id="nestedatt--fields"></a>
### Nested Schema for `fields`

Read-Only:

- `description` (String) The description of the task template's field
- `id` (Attributes) the identifier (name) for the task template field (see [below for nested schema](#nestedatt--fields--id))
- `single_select_options` (List of String) list of field options to be used with single select
- `type` (String) The type of the task template's field

<a id="nestedatt--fields--id"></a>
### Nested Schema for `fields.id`

Read-Only:

- `name` (String) The name of the task template field



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

