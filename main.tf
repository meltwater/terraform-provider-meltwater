variable "meltwater_api_key" {
    type = string
}

provider "meltwater" {
  # NOTE: This is populated from the `TF_VAR_meltwater_api_key` environment variable.
  api_key = var.meltwater_api_key
}

#Add search creation here
resource "meltwater_search" "my_awesome_search" {
  type = "social"
  category = "explore"
  name = "Golang - terraform"
  query {
    keyword {
      case_sensitivity = "yes"
      all_keywords = ["golanga"]
      any_keywords = []
      not_keywords = []
    }
  }
}

resource "meltwater_recurring_export" "my_awesome_recurring_export" {
  search_id = meltwater_search.my_awesome_search.id
  timezone = "Europe/Paris"
  window_time_unit = "week"
  window_time = "00:00:00"
  window_size = 1
}
