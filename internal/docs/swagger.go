package docs

import (
	"embed"
)

//go:embed swagger_doc.json
var swaggerFS embed.FS

func SwaggerJSON() ([]byte, error) {
	return swaggerFS.ReadFile("swagger_doc.json")
}
