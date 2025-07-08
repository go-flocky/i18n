package locales

import "embed"

//go:embed config.yaml
//go:embed de en fr
var LocaleFS embed.FS
