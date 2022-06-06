# tfadd

Generate valid Terraform configuration from state.

## Install

```
go install github.com/magodo/tfadd@latest
```

## Usage

The goal of this tool is to improve the [import experience](https://learn.hashicorp.com/tutorials/terraform/state-import?in=terraform/state&utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) of Terraform, that rather than constructing the configurations from scratch, `tfadd` (try its best to) provide users a **valid** configuration automatically.

> The **valid** here means the generated configuration should raise no error and show no diff during `terraform plan`. 

The typical usage is to use `tfadd` together with `terraform import`:

1. Prepare an empty [workspace](https://www.terraform.io/language/state/workspaces) (e.g. an empty directory for *local* backend)
1. (`tfadd` only) Run `tfadd init [providers...] > terraform.tf` to populate the Terraform setting to pin the provider version
1. Run `terraform init` to initialize the providers
1. Identify the existing resources to be managed via `terraform`, write down the empty resource block and import them via `terraform import`
1. (`tfadd` only) Run `tfadd state` to generate the configuration

Currently, the tool supports the following providers:

|Name|Version|
|-|-|
|registry.terraform.io/hashicorp/aws|v4.17.1|
|registry.terraform.io/hashicorp/azurerm|v3.9.0|
|registry.terraform.io/hashicorp/google|v4.23.0|

## Limitation

- Only the managed resources of the root module in the state file will get the config generated, any child module will be skipped. 
- No inter-resource dependency generated.
