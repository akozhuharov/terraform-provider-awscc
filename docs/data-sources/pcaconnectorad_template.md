---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_pcaconnectorad_template Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::PCAConnectorAD::Template
---

# awscc_pcaconnectorad_template (Data Source)

Data Source schema for AWS::PCAConnectorAD::Template



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `connector_arn` (String)
- `definition` (Attributes) (see [below for nested schema](#nestedatt--definition))
- `name` (String)
- `reenroll_all_certificate_holders` (Boolean)
- `tags` (Map of String)
- `template_arn` (String)

<a id="nestedatt--definition"></a>
### Nested Schema for `definition`

Read-Only:

- `template_v2` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2))
- `template_v3` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3))
- `template_v4` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4))

<a id="nestedatt--definition--template_v2"></a>
### Nested Schema for `definition.template_v2`

Read-Only:

- `certificate_validity` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--certificate_validity))
- `enrollment_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--enrollment_flags))
- `extensions` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--extensions))
- `general_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--general_flags))
- `private_key_attributes` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--private_key_attributes))
- `private_key_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--private_key_flags))
- `subject_name_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--subject_name_flags))
- `superseded_templates` (List of String)

<a id="nestedatt--definition--template_v2--certificate_validity"></a>
### Nested Schema for `definition.template_v2.certificate_validity`

Read-Only:

- `renewal_period` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--certificate_validity--renewal_period))
- `validity_period` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--certificate_validity--validity_period))

<a id="nestedatt--definition--template_v2--certificate_validity--renewal_period"></a>
### Nested Schema for `definition.template_v2.certificate_validity.validity_period`

Read-Only:

- `period` (Number)
- `period_type` (String)


<a id="nestedatt--definition--template_v2--certificate_validity--validity_period"></a>
### Nested Schema for `definition.template_v2.certificate_validity.validity_period`

Read-Only:

- `period` (Number)
- `period_type` (String)



<a id="nestedatt--definition--template_v2--enrollment_flags"></a>
### Nested Schema for `definition.template_v2.enrollment_flags`

Read-Only:

- `enable_key_reuse_on_nt_token_keyset_storage_full` (Boolean)
- `include_symmetric_algorithms` (Boolean)
- `no_security_extension` (Boolean)
- `remove_invalid_certificate_from_personal_store` (Boolean)
- `user_interaction_required` (Boolean)


<a id="nestedatt--definition--template_v2--extensions"></a>
### Nested Schema for `definition.template_v2.extensions`

Read-Only:

- `application_policies` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--extensions--application_policies))
- `key_usage` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--extensions--key_usage))

<a id="nestedatt--definition--template_v2--extensions--application_policies"></a>
### Nested Schema for `definition.template_v2.extensions.key_usage`

Read-Only:

- `critical` (Boolean)
- `policies` (Attributes List) (see [below for nested schema](#nestedatt--definition--template_v2--extensions--key_usage--policies))

<a id="nestedatt--definition--template_v2--extensions--key_usage--policies"></a>
### Nested Schema for `definition.template_v2.extensions.key_usage.policies`

Read-Only:

- `policy_object_identifier` (String)
- `policy_type` (String)



<a id="nestedatt--definition--template_v2--extensions--key_usage"></a>
### Nested Schema for `definition.template_v2.extensions.key_usage`

Read-Only:

- `critical` (Boolean)
- `usage_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v2--extensions--key_usage--usage_flags))

<a id="nestedatt--definition--template_v2--extensions--key_usage--usage_flags"></a>
### Nested Schema for `definition.template_v2.extensions.key_usage.usage_flags`

Read-Only:

- `data_encipherment` (Boolean)
- `digital_signature` (Boolean)
- `key_agreement` (Boolean)
- `key_encipherment` (Boolean)
- `non_repudiation` (Boolean)




<a id="nestedatt--definition--template_v2--general_flags"></a>
### Nested Schema for `definition.template_v2.general_flags`

Read-Only:

- `auto_enrollment` (Boolean)
- `machine_type` (Boolean)


<a id="nestedatt--definition--template_v2--private_key_attributes"></a>
### Nested Schema for `definition.template_v2.private_key_attributes`

Read-Only:

- `crypto_providers` (List of String)
- `key_spec` (String)
- `minimal_key_length` (Number)


<a id="nestedatt--definition--template_v2--private_key_flags"></a>
### Nested Schema for `definition.template_v2.private_key_flags`

Read-Only:

- `client_version` (String)
- `exportable_key` (Boolean)
- `strong_key_protection_required` (Boolean)


<a id="nestedatt--definition--template_v2--subject_name_flags"></a>
### Nested Schema for `definition.template_v2.subject_name_flags`

Read-Only:

- `require_common_name` (Boolean)
- `require_directory_path` (Boolean)
- `require_dns_as_cn` (Boolean)
- `require_email` (Boolean)
- `san_require_directory_guid` (Boolean)
- `san_require_dns` (Boolean)
- `san_require_domain_dns` (Boolean)
- `san_require_email` (Boolean)
- `san_require_spn` (Boolean)
- `san_require_upn` (Boolean)



<a id="nestedatt--definition--template_v3"></a>
### Nested Schema for `definition.template_v3`

Read-Only:

- `certificate_validity` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--certificate_validity))
- `enrollment_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--enrollment_flags))
- `extensions` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--extensions))
- `general_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--general_flags))
- `hash_algorithm` (String)
- `private_key_attributes` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--private_key_attributes))
- `private_key_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--private_key_flags))
- `subject_name_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--subject_name_flags))
- `superseded_templates` (List of String)

<a id="nestedatt--definition--template_v3--certificate_validity"></a>
### Nested Schema for `definition.template_v3.certificate_validity`

Read-Only:

- `renewal_period` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--certificate_validity--renewal_period))
- `validity_period` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--certificate_validity--validity_period))

<a id="nestedatt--definition--template_v3--certificate_validity--renewal_period"></a>
### Nested Schema for `definition.template_v3.certificate_validity.validity_period`

Read-Only:

- `period` (Number)
- `period_type` (String)


<a id="nestedatt--definition--template_v3--certificate_validity--validity_period"></a>
### Nested Schema for `definition.template_v3.certificate_validity.validity_period`

Read-Only:

- `period` (Number)
- `period_type` (String)



<a id="nestedatt--definition--template_v3--enrollment_flags"></a>
### Nested Schema for `definition.template_v3.enrollment_flags`

Read-Only:

- `enable_key_reuse_on_nt_token_keyset_storage_full` (Boolean)
- `include_symmetric_algorithms` (Boolean)
- `no_security_extension` (Boolean)
- `remove_invalid_certificate_from_personal_store` (Boolean)
- `user_interaction_required` (Boolean)


<a id="nestedatt--definition--template_v3--extensions"></a>
### Nested Schema for `definition.template_v3.extensions`

Read-Only:

- `application_policies` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--extensions--application_policies))
- `key_usage` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--extensions--key_usage))

<a id="nestedatt--definition--template_v3--extensions--application_policies"></a>
### Nested Schema for `definition.template_v3.extensions.key_usage`

Read-Only:

- `critical` (Boolean)
- `policies` (Attributes List) (see [below for nested schema](#nestedatt--definition--template_v3--extensions--key_usage--policies))

<a id="nestedatt--definition--template_v3--extensions--key_usage--policies"></a>
### Nested Schema for `definition.template_v3.extensions.key_usage.policies`

Read-Only:

- `policy_object_identifier` (String)
- `policy_type` (String)



<a id="nestedatt--definition--template_v3--extensions--key_usage"></a>
### Nested Schema for `definition.template_v3.extensions.key_usage`

Read-Only:

- `critical` (Boolean)
- `usage_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--extensions--key_usage--usage_flags))

<a id="nestedatt--definition--template_v3--extensions--key_usage--usage_flags"></a>
### Nested Schema for `definition.template_v3.extensions.key_usage.usage_flags`

Read-Only:

- `data_encipherment` (Boolean)
- `digital_signature` (Boolean)
- `key_agreement` (Boolean)
- `key_encipherment` (Boolean)
- `non_repudiation` (Boolean)




<a id="nestedatt--definition--template_v3--general_flags"></a>
### Nested Schema for `definition.template_v3.general_flags`

Read-Only:

- `auto_enrollment` (Boolean)
- `machine_type` (Boolean)


<a id="nestedatt--definition--template_v3--private_key_attributes"></a>
### Nested Schema for `definition.template_v3.private_key_attributes`

Read-Only:

- `algorithm` (String)
- `crypto_providers` (List of String)
- `key_spec` (String)
- `key_usage_property` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--private_key_attributes--key_usage_property))
- `minimal_key_length` (Number)

<a id="nestedatt--definition--template_v3--private_key_attributes--key_usage_property"></a>
### Nested Schema for `definition.template_v3.private_key_attributes.minimal_key_length`

Read-Only:

- `property_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v3--private_key_attributes--minimal_key_length--property_flags))
- `property_type` (String)

<a id="nestedatt--definition--template_v3--private_key_attributes--minimal_key_length--property_flags"></a>
### Nested Schema for `definition.template_v3.private_key_attributes.minimal_key_length.property_flags`

Read-Only:

- `decrypt` (Boolean)
- `key_agreement` (Boolean)
- `sign` (Boolean)




<a id="nestedatt--definition--template_v3--private_key_flags"></a>
### Nested Schema for `definition.template_v3.private_key_flags`

Read-Only:

- `client_version` (String)
- `exportable_key` (Boolean)
- `require_alternate_signature_algorithm` (Boolean)
- `strong_key_protection_required` (Boolean)


<a id="nestedatt--definition--template_v3--subject_name_flags"></a>
### Nested Schema for `definition.template_v3.subject_name_flags`

Read-Only:

- `require_common_name` (Boolean)
- `require_directory_path` (Boolean)
- `require_dns_as_cn` (Boolean)
- `require_email` (Boolean)
- `san_require_directory_guid` (Boolean)
- `san_require_dns` (Boolean)
- `san_require_domain_dns` (Boolean)
- `san_require_email` (Boolean)
- `san_require_spn` (Boolean)
- `san_require_upn` (Boolean)



<a id="nestedatt--definition--template_v4"></a>
### Nested Schema for `definition.template_v4`

Read-Only:

- `certificate_validity` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--certificate_validity))
- `enrollment_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--enrollment_flags))
- `extensions` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--extensions))
- `general_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--general_flags))
- `hash_algorithm` (String)
- `private_key_attributes` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--private_key_attributes))
- `private_key_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--private_key_flags))
- `subject_name_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--subject_name_flags))
- `superseded_templates` (List of String)

<a id="nestedatt--definition--template_v4--certificate_validity"></a>
### Nested Schema for `definition.template_v4.certificate_validity`

Read-Only:

- `renewal_period` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--certificate_validity--renewal_period))
- `validity_period` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--certificate_validity--validity_period))

<a id="nestedatt--definition--template_v4--certificate_validity--renewal_period"></a>
### Nested Schema for `definition.template_v4.certificate_validity.validity_period`

Read-Only:

- `period` (Number)
- `period_type` (String)


<a id="nestedatt--definition--template_v4--certificate_validity--validity_period"></a>
### Nested Schema for `definition.template_v4.certificate_validity.validity_period`

Read-Only:

- `period` (Number)
- `period_type` (String)



<a id="nestedatt--definition--template_v4--enrollment_flags"></a>
### Nested Schema for `definition.template_v4.enrollment_flags`

Read-Only:

- `enable_key_reuse_on_nt_token_keyset_storage_full` (Boolean)
- `include_symmetric_algorithms` (Boolean)
- `no_security_extension` (Boolean)
- `remove_invalid_certificate_from_personal_store` (Boolean)
- `user_interaction_required` (Boolean)


<a id="nestedatt--definition--template_v4--extensions"></a>
### Nested Schema for `definition.template_v4.extensions`

Read-Only:

- `application_policies` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--extensions--application_policies))
- `key_usage` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--extensions--key_usage))

<a id="nestedatt--definition--template_v4--extensions--application_policies"></a>
### Nested Schema for `definition.template_v4.extensions.key_usage`

Read-Only:

- `critical` (Boolean)
- `policies` (Attributes List) (see [below for nested schema](#nestedatt--definition--template_v4--extensions--key_usage--policies))

<a id="nestedatt--definition--template_v4--extensions--key_usage--policies"></a>
### Nested Schema for `definition.template_v4.extensions.key_usage.policies`

Read-Only:

- `policy_object_identifier` (String)
- `policy_type` (String)



<a id="nestedatt--definition--template_v4--extensions--key_usage"></a>
### Nested Schema for `definition.template_v4.extensions.key_usage`

Read-Only:

- `critical` (Boolean)
- `usage_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--extensions--key_usage--usage_flags))

<a id="nestedatt--definition--template_v4--extensions--key_usage--usage_flags"></a>
### Nested Schema for `definition.template_v4.extensions.key_usage.usage_flags`

Read-Only:

- `data_encipherment` (Boolean)
- `digital_signature` (Boolean)
- `key_agreement` (Boolean)
- `key_encipherment` (Boolean)
- `non_repudiation` (Boolean)




<a id="nestedatt--definition--template_v4--general_flags"></a>
### Nested Schema for `definition.template_v4.general_flags`

Read-Only:

- `auto_enrollment` (Boolean)
- `machine_type` (Boolean)


<a id="nestedatt--definition--template_v4--private_key_attributes"></a>
### Nested Schema for `definition.template_v4.private_key_attributes`

Read-Only:

- `algorithm` (String)
- `crypto_providers` (List of String)
- `key_spec` (String)
- `key_usage_property` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--private_key_attributes--key_usage_property))
- `minimal_key_length` (Number)

<a id="nestedatt--definition--template_v4--private_key_attributes--key_usage_property"></a>
### Nested Schema for `definition.template_v4.private_key_attributes.minimal_key_length`

Read-Only:

- `property_flags` (Attributes) (see [below for nested schema](#nestedatt--definition--template_v4--private_key_attributes--minimal_key_length--property_flags))
- `property_type` (String)

<a id="nestedatt--definition--template_v4--private_key_attributes--minimal_key_length--property_flags"></a>
### Nested Schema for `definition.template_v4.private_key_attributes.minimal_key_length.property_flags`

Read-Only:

- `decrypt` (Boolean)
- `key_agreement` (Boolean)
- `sign` (Boolean)




<a id="nestedatt--definition--template_v4--private_key_flags"></a>
### Nested Schema for `definition.template_v4.private_key_flags`

Read-Only:

- `client_version` (String)
- `exportable_key` (Boolean)
- `require_alternate_signature_algorithm` (Boolean)
- `require_same_key_renewal` (Boolean)
- `strong_key_protection_required` (Boolean)
- `use_legacy_provider` (Boolean)


<a id="nestedatt--definition--template_v4--subject_name_flags"></a>
### Nested Schema for `definition.template_v4.subject_name_flags`

Read-Only:

- `require_common_name` (Boolean)
- `require_directory_path` (Boolean)
- `require_dns_as_cn` (Boolean)
- `require_email` (Boolean)
- `san_require_directory_guid` (Boolean)
- `san_require_dns` (Boolean)
- `san_require_domain_dns` (Boolean)
- `san_require_email` (Boolean)
- `san_require_spn` (Boolean)
- `san_require_upn` (Boolean)