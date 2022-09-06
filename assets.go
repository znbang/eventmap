package assets

import (
	"embed"
)

//go:embed frontend/dist/spa
var FS embed.FS
