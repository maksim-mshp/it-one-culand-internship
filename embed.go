package embed

import "embed"

//go:embed api/openapi.json api/openapi.yml
var OpenAPIFS embed.FS
