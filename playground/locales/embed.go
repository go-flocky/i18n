package locales

import "embed"

//go:embed config.yaml
//go:embed de en fr en-US
var LocaleFS embed.FS
