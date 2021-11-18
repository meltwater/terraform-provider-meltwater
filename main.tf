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
