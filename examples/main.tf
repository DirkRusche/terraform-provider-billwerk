terraform {
  required_providers {
    billwerk = {
      source = "hashicorp.com/dirkrusche/billwerk"
    }
  }
}

variable "billwerk_provider" {
  type = string
}

provider "billwerk" {
  token = var.billwerk_provider
  url = "https://sandbox.billwerk.com"
}

resource "billwerk_mail_template" "test" {
  internal_name = "test"
  external_id = "foo_bar"
  event_type = "ActivePhaseChanged"

  language {
    language = "_c"
    subject = "hello"
    text = "world"
  }

  language {
    language = "en"
    subject = "foo"
    text = "bar"
  }
}
