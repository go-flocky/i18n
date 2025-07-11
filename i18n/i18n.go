package i18n

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/go-flocky/i18n/internal/kvTree"
)

func (i *I18n) GetLocale(code string) *Locale {
	return i.locales[code]
}

func (i *I18n) RegisterLocale(locale *Locale) {
	i.locales[locale.Code] = locale
}

func (i *I18n) HasLocale(locale LocaleCode) bool {
	_, ok := i.Dictionary[locale]
	return ok
}

func (i *I18n) GetDictionaryForCode(localeCode LocaleCode) (*kvtree.KeyValueTree[string], bool) {
	dict, ok := i.Dictionary[localeCode]
	return dict, ok
}

func (i *I18n) T(ctx context.Context, key string, args ...any) string {
	if translated := i.translate(i.GetLocaleFromCtx(ctx), key, args...); translated != "" {
		return translated
	}

	for _, fallbackCode := range i.config.FallbackLocaleCode {
		if translated := i.translate(fallbackCode, key, args...); translated != "" {
			return translated
		}
	}

	return MissingTranslation(i.GetLocaleFromCtx(ctx), key)
}

func (i *I18n) translate(localeCode LocaleCode, key string, args ...any) string {
	dict, ok := i.GetDictionaryForCode(localeCode)
	if !ok {
		return ""
	}

	translated, ok := dict.Get(key)
	if !ok {
		return ""
	}

	if len(args) > 0 {
		return fmt.Sprintf(*translated, args...)
	}

	return *translated
}

func NewI18n(filesystem fs.FS) (*I18n, error) {
	i18nConfig, err := parseConfig(filesystem)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if i18nConfig.KeySeparator == "" {
		i18nConfig.KeySeparator = "."
	}

	instance := &I18n{
		locales:    make(map[string]*Locale),
		config:     *i18nConfig,
		localeFS:   filesystem,
		Dictionary: make(map[string]*kvtree.KeyValueTree[string]),
	}

	return instance, nil
}
