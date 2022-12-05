# tfadd

Generate valid Terraform configuration from state.

## Install

```
go install github.com/magodo/tfadd@latest
```

## Intro

The goal of this tool is to improve the [import experience](https://learn.hashicorp.com/tutorials/terraform/state-import?in=terraform/state&utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) of Terraform, that rather than constructing the configurations from scratch, `tfadd` (try its best to) provide users a **valid** configuration automatically.

> The **valid** here means the generated configuration should raise no error and show no diff during `terraform plan`. 

Currently to generate the state, the tool supports *full mode* (with `-full`) or *partial mode* (by default).

- In *full mode*, `tfadd` outputs all non-computed properties in the generated config. The generated config might be **invalid** for kinds of reasons, where manual modification is needed. But the benefit is that it works for any Terraform provider.
- In *partial mode*, `tfadd` only outputs properties without `Optional+Computed` properties, with cross property constraints taken into consideration. This mode aims aims to generate a **valid** Terraform config. Currently, this mode can only works for the following providers:

    |Name|Version|
    |-|-|
    |registry.terraform.io/hashicorp/aws|v4.45.0|
    |registry.terraform.io/hashicorp/azurerm|v3.34.0|
    |registry.terraform.io/hashicorp/google|v4.44.0|

## Usage

The typical usage is to use `tfadd` together with `terraform import`:

1. Prepare an empty [workspace](https://www.terraform.io/language/state/workspaces) (e.g. an empty directory for *local* backend)
1. Identify the existing resources to import, write down the empty resource block
1. (*partial mode* only) Run `tfadd init [providers...] > terraform.tf` to populate the Terraform setting to pin the provider version
1. Run `terraform init` to initialize the providers
1. Import the resources via `terraform import`
1. Run `tfadd state` or `tfadd state -full` to generate the configuration


## Limitation

- Only the managed resources of the root module in the state file will get the config generated, any child module will be skipped. 
- No inter-resource dependency generated.
