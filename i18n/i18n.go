package i18n

import (
	"strings"
)

const MissingTranslation = "ERROR: Missing translation"

type i18n struct {
	FallbackLocaleCode []string
	locales            map[string]*Locale
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

func NewI18n() *i18n {
	return &i18n{
		locales: make(map[string]*Locale),
	}
}
