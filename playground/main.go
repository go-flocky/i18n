package main

import (
	"fmt"
	"log/slog"

	flockyI18n "github.com/go-flocky/i18n/i18n"
	"github.com/go-flocky/i18n/playground/locales"
)

func main() {
	i18n, err := flockyI18n.NewI18n(locales.LocaleFS)
	if err != nil {
		slog.Error("Error creating i18n instance:", "err", err)
		return
	}
	if err := i18n.LoadLocales(); err != nil {
		slog.Error("Error loading locales:", "err", err)
		return
	}
	fmt.Println(i18n.T("de", "hello"))
	fmt.Println(i18n.T("de", "chicken", 2))
}
