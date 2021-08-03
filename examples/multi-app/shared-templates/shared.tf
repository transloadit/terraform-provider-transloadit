terraform {
  required_version = ">= 1.0.0"
  required_providers {
    transloadit = {
      source  = "transloadit/transloadit"
      version = "0.4.0"
    }
  }
}

resource "transloadit_template" "import-file" {
  name        = "import-file"
  template    = <<EOT
{
  "steps": {
    "imported": {
      "robot": "/http/import",
      "url": "https://example.com/logo.gif"
    }
  }
}
EOT
}