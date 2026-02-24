package embed

import "embed"

//go:embed api/openapi.yml
var OpenAPIFS embed.FS
