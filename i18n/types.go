package i18n

import (
	"io/fs"

	"github.com/go-flocky/i18n/internal/kvTree"
)

type I18n struct {
	locales    map[string]*Locale
	localeFS   fs.FS
	Dictionary map[LocaleCode]*kvtree.KeyValueTree[string]
	config     i18nConfig
}

type LocaleCode = string

type Locale struct {
	Code LocaleCode
	Name string
}

type LocaleContextKey = string
type i18nConfig struct {
	FallbackLocaleCode []string         `yaml:"fallback"`
	DefaultLocaleCode  string           `yaml:"default"`
	KeySeparator       string           `yaml:"separator"`
	LocaleContextKey   LocaleContextKey `yaml:"ctxKey"`
}

type localeConfig struct {
	Code string `yaml:"code"`
	Name string `yaml:"name"`
}
