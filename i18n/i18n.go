package i18n

import (
	"fmt"
	"io"
	"io/fs"
	"strings"

	"gopkg.in/yaml.v3"
)

func (i *I18n) GetLocale(code string) *Locale {
	return i.locales[code]
}

func (i *I18n) RegisterLocale(locale *Locale) {
	i.locales[locale.Code] = locale
}

func (i *I18n) T(localeCode, key string, args ...any) string {
	var value string

	if len(args) >= 1 {
		if count, ok := args[0].(int); ok {
			value = i.translatePlural(localeCode, key, count)
			if value != MissingTranslation {
				return value
			}
		} else {
			fmt.Printf("Invalid argument: %v\n", args[0])
		}
	}

	value = i.translate(localeCode, key)
	if value != MissingTranslation {
		return value
	}

	for _, fallbackCode := range i.FallbackLocaleCodes {
		if fallbackCode == localeCode {
			continue
		}
		value = i.translate(fallbackCode, key)
		if value != MissingTranslation {
			return value
		}
	}

	return MissingTranslation
}

func (i *I18n) translatePlural(localeCode, key string, n int) string {
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


	if dict != nil {
		return dict.Value.One
	}

	return MissingTranslation
}

func (i *I18n) translate(localeCode, key string) string {
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

	if dict != nil && dict.Value.One != "" {
		return dict.Value.One
	}

	return MissingTranslation
}

func parseConfig(fs fs.FS) (*i18nConfig, error) {
	configFile, err := fs.Open(ConfigFileName)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	cfg, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	i18nConfig := &i18nConfig{} 
	err = yaml.Unmarshal(cfg, &i18nConfig)
	if err != nil {
		return nil, err
	}
	return i18nConfig, nil
}

func NewI18n(fs fs.FS) (*I18n, error) {
	i18nConfig,err := parseConfig(fs)
	if err != nil {
		return nil, err
	}

	instance := &I18n{
		locales:             make(map[string]*Locale),
		FallbackLocaleCodes: i18nConfig.FallbackLocaleCode,
		DefaultLocaleCode:   i18nConfig.DefaultLocaleCode,
		localeFS:            &fs,
	}

	return instance, nil
}
