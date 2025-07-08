package i18n

import (
	"io"
	"io/fs"
	"strings"

	"gopkg.in/yaml.v3"
)

const MissingTranslation = "ERROR: Missing translation"

type i18nConfig struct {
	FallbackLocaleCode []string `yaml:"fallback"`
	DefaultLocaleCode  string   `yaml:"default"`
}

type i18n struct {
	FallbackLocaleCode []string
	DefaultLocaleCode  string
	locales            map[string]*Locale
	localeFS           *fs.FS
}

func (i *i18n) GetLocale(code string) *Locale {
	return i.locales[code]
}

func (i *i18n) RegisterLocale(locale *Locale) {
	i.locales[locale.Code] = locale
}

func (i *i18n) T(localeCode, key string) string {
	result := i.translate(localeCode, key)
	if result != MissingTranslation {
		return result
	}

	for _, fallbackCode := range i.FallbackLocaleCode {
		if fallbackCode == localeCode {
			continue
		}
		result = i.translate(fallbackCode, key)
		if result != MissingTranslation {
			return result
		}
	}

	return MissingTranslation
}

func (i *i18n) translate(localeCode, key string) string {
	locale := i.locales[localeCode]
	if locale == nil {
		return MissingTranslation
	}

	keys := strings.Split(key, ".")
	dict := locale.Dictionary

	for _, part := range keys {
		if dict == nil || dict.ChildDict == nil {
			return MissingTranslation
		}
		dict = dict.ChildDict[part]
	}

	if dict != nil && dict.Value != "" {
		return dict.Value
	}

	return MissingTranslation
}

func NewI18n(fs fs.FS) (*i18n, error) {
	configFile, err := fs.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	cfg, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	var i18nConfig i18nConfig
	err = yaml.Unmarshal(cfg, &i18nConfig)
	if err != nil {
		return nil, err
	}

	return &i18n{
		locales:            make(map[string]*Locale),
		FallbackLocaleCode: i18nConfig.FallbackLocaleCode,
		DefaultLocaleCode:  i18nConfig.DefaultLocaleCode,
		localeFS:           &fs,
	}, nil
}
