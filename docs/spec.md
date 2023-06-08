# Specification

The release notes contain entries designed to replicate the structure of a Terraform provider.

 - `provider` -- describes the provider level changes.

 - `resource` -- describes the resource level changes, one per resource.

 - `data_source` -- describes the data source level changes, one per resource.

The file for release notes has the extension ***.cn***, which stands for "Change Note".

## Provider

```hcl
provider "provider_name" {}
```

## Resource

```hcl
resource "resource_name" {
  bug_fix      = <STRING>
  doc          = <STRING>
  enhancement  = <STRING>
  feature      = <STRING>
  new_resource = <BOOL>
}
```

## Data Source

```hcl
data_source "data_source_name" {}
```
