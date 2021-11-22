variable "meltwater_api_key" {
    type = string
}

provider "meltwater" {
  # NOTE: This is populated from the `TF_VAR_meltwater_api_key` environment variable.
  api_key = var.meltwater_api_key
}

resource "meltwater_search" "my_awesome_search" {
  type = "social"
  category = "explore"
  name = "Golang - terraform"
  query {
    /*keyword {
      case_sensitivity = "yes"
      all_keywords = ["golang"]
      any_keywords = []
      not_keywords = ["java"]
    }*/
    /*combined {
      all_searches = [16058498]
      not_searches = [7912335]
    }*/
    /*boolean {
      case_sensitivity = "hybrid"
      boolean = "(SourceName: /r/ProgrammingHumour OR SourceName: /r/Golang) AND metaData.discussionType:\"og\" AND language:\"fr\""
    }/
  }
}

/*resource "meltwater_recurring_export" "my_awesome_recurring_export" {
  search_id = meltwater_search.my_awesome_search.id
  timezone = "Europe/Paris"
  window_time_unit = "week"
  window_time = "00:00:00"
  window_size = 1
}*/
