---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_lambda_function Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Lambda::Function
---

# awscc_lambda_function (Resource)

Resource Type definition for AWS::Lambda::Function



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **code** (Attributes) (see [below for nested schema](#nestedatt--code))
- **role** (String) The Amazon Resource Name (ARN) of the function's execution role.

### Optional

- **code_signing_config_arn** (String) A unique Arn for CodeSigningConfig resource
- **dead_letter_config** (Attributes) The dead-letter queue for failed asynchronous invocations. (see [below for nested schema](#nestedatt--dead_letter_config))
- **description** (String) A description of the function.
- **environment** (Attributes) A function's environment variable settings. (see [below for nested schema](#nestedatt--environment))
- **file_system_configs** (Attributes List, Max: 1) Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an AWS::EFS::MountTarget resource, you must also specify a DependsOn attribute to ensure that the mount target is created or updated before the function. (see [below for nested schema](#nestedatt--file_system_configs))
- **function_name** (String) The name of the Lambda function, up to 64 characters in length. If you don't specify a name, AWS CloudFormation generates one.
- **handler** (String) The name of the method within your code that Lambda calls to execute your function. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime
- **image_config** (Attributes) (see [below for nested schema](#nestedatt--image_config))
- **kms_key_arn** (String) The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
- **layers** (List of String) A list of function layers to add to the function's execution environment. Specify each layer by its ARN, including the version.
- **memory_size** (Number) The amount of memory that your function has access to. Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.
- **package_type** (String) PackageType.
- **reserved_concurrent_executions** (Number) The number of simultaneous executions to reserve for the function.
- **runtime** (String) The identifier of the function's runtime.
- **tags** (Attributes Set) A list of tags to apply to the function. (see [below for nested schema](#nestedatt--tags))
- **timeout** (Number) The amount of time that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.
- **tracing_config** (Attributes) The function's AWS X-Ray tracing configuration. To sample and record incoming requests, set Mode to Active. (see [below for nested schema](#nestedatt--tracing_config))
- **vpc_config** (Attributes) The VPC security groups and subnets that are attached to a Lambda function. When you connect a function to a VPC, Lambda creates an elastic network interface for each combination of security group and subnet in the function's VPC configuration. The function can only access resources and the internet through that VPC. (see [below for nested schema](#nestedatt--vpc_config))

### Read-Only

- **arn** (String) Unique identifier for function resources
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--code"></a>
### Nested Schema for `code`

Required:

- **image_uri** (String) ImageUri.
- **s3_bucket** (String) An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.
- **s3_key** (String) The Amazon S3 key of the deployment package.
- **s3_object_version** (String) For versioned objects, the version of the deployment package object to use.
- **zip_file** (String) The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..


<a id="nestedatt--dead_letter_config"></a>
### Nested Schema for `dead_letter_config`

Optional:

- **target_arn** (String) The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.


<a id="nestedatt--environment"></a>
### Nested Schema for `environment`

Optional:

- **variables** (Map of String) Environment variable key-value pairs.


<a id="nestedatt--file_system_configs"></a>
### Nested Schema for `file_system_configs`

Optional:

- **arn** (String) The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.
- **local_mount_path** (String) The path where the function can access the file system, starting with /mnt/.


<a id="nestedatt--image_config"></a>
### Nested Schema for `image_config`

Optional:

- **command** (List of String) Command.
- **entry_point** (List of String) EntryPoint.
- **working_directory** (String) WorkingDirectory.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- **value** (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--tracing_config"></a>
### Nested Schema for `tracing_config`

Optional:

- **mode** (String) The tracing mode.


<a id="nestedatt--vpc_config"></a>
### Nested Schema for `vpc_config`

Optional:

- **security_group_ids** (List of String) A list of VPC security groups IDs.
- **subnet_ids** (List of String) A list of VPC subnet IDs.

