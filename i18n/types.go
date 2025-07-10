package i18n

import "io/fs"

type i18nConfig struct {
	FallbackLocaleCode []string `yaml:"fallback"`
	DefaultLocaleCode  string   `yaml:"default"`
}

type I18n struct {
	FallbackLocaleCodes []string
	DefaultLocaleCode   string
	locales             map[string]*Locale
	localeFS            *fs.FS
}

type Dictionary struct {
	Value     DictionaryValue
	ChildDict map[string]*Dictionary
}

type DictionaryValue struct {
	Zero  string `yaml:"zero"`
	One   string `yaml:"one"`
	Two   string `yaml:"two"`
	Few   string `yaml:"few"`
	Many  string `yaml:"many"`
	Other string `yaml:"other"`
}

type Locale struct {
	Code       string
	Name       string
	Dictionary *Dictionary
}

type localeConfig struct {
	Code string `yaml:"code"`
	Name string `yaml:"name"`
}
