variable "url" {
  type    = string
}

variable "revisions_schema" {
  type    = string
	default = "public"
}

env "local" {
	url = var.url
  dev = "docker://postgres/14-alpine/kucingmenangis?search_path=public&sslmode=disable"
  src = [
		"file://schema/codebase.hcl",
	]

	migration {
    dir = "file://migration"
		revisions_schema = var.revisions_schema
    format = atlas
  }
}
