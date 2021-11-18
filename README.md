# Meltwater Terraform Provider
A Terraform provider for Meltwater based on the persistent/infrastructure parts of the Meltwater API. Such as setting up webhooks and recurring exports.

## Requirements

*	[Terraform](https://www.terraform.io/downloads.html) 0.11.x to 0.13.x
*	[Go](https://golang.org/doc/install) 1.14 to 1.17 (to build the provider plugin)

## Build

Clone repository anywhere:

```sh
$ git clone https://github.com/meltwater/terraform-provider-meltwater.git
```

Enter the provider directory and build the provider

```sh
$ cd terraform-provider-meltwater
$ make compile
```

Or alternatively, to install it as a plugin, run

```sh
$ cd terraform-provider-meltwater
$ make install
```

## Using the provider

If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

### Basic Usage

Set an environment variable, `TF_VAR_meltwater_api_key` to store your Meltwater API key. This is the recommended way to not commit an access token into your version control system.

    export TF_VAR_meltwater_api_key=<your meltwater api key>

Your token is now accessible in your Terraform configuration as
`var.meltwater_api_key`, and can be used to configure the provider.

The example below demonstrates the following operations:

  * Create a recurring export
  * TODO: Create a search
  * TODO: Create a hook to receive matching documents

```hcl
terraform {
  required_providers {
    meltwater = {
      source = "meltwater/meltwater"
      version = "<VERSION_TBC>"
    }
  }
}

variable "meltwater_api_key" {
    type = string
}

provider "meltwater" {
  # NOTE: This is populated from the `TF_VAR_meltwater_api_key` environment variable.
  api_key = var.meltwater_api_key
}

resource "meltwater_recurring_export" "my_awesome_recurring_export" {
  search_id = 16058498
  timezone = "Europe/London"
  window_time_unit = "week"
  window_time = "00:00:00"
  window_size = 1
}
```
