package i18n

import (
	"context"
	"fmt"
)

func (i *I18n) WithLocale(ctx context.Context, locale LocaleCode) (context.Context, error) {
	l := locale
	if !i.HasLocale(locale) {
		if !i.HasLocale(i.config.DefaultLocaleCode) {
			return nil, fmt.Errorf("cannot use non existant locale in context")
		}
	}

	return context.WithValue(ctx, i.config.LocaleContextKey, l), nil
}

func (i *I18n) GetLocaleFromCtx(ctx context.Context) LocaleCode {
	localeCode := ctx.Value(i.config.LocaleContextKey)

	switch localeCode := localeCode.(type) {
	case string:
		if localeCode != "" {
			return localeCode
		}
		return i.config.DefaultLocaleCode
	default:
		return i.config.DefaultLocaleCode
	}

}
