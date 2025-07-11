package main

import (
	"context"
	"fmt"
	"log"

	internalI18n"github.com/go-flocky/i18n/i18n"
	"github.com/go-flocky/i18n/playground/locales"
)

func main() {
	i18n, err := internalI18n.NewI18n(locales.LocaleFS)
	if err != nil {
		log.Println("Error creating i18n instance: ", err)
		return
	}
	if err := i18n.LoadLocales(); err != nil {
		log.Println("Error loading locales: ", err)
		return
	}

	// en
	fmt.Println("En: ")
	ctx, err := i18n.WithLocale(context.Background(), "en")
	if err != nil {
		log.Println("Error set locale in context: ", err)
	}

	fmt.Println(i18n.T(ctx, "anrsietnai"))
	fmt.Println(i18n.T(ctx, "hello", "You"))
	fmt.Println(i18n.T(ctx, "chicken.zero", 1))

	//de
	fmt.Println("De: ")
	ctx, err = i18n.WithLocale(ctx, "fr")
	if err != nil {
		log.Println("Error set locale in context: ", err)
	}

	fmt.Println(i18n.T(ctx, "anrsietnai"))
	fmt.Println(i18n.T(ctx, "hello", "You"))
	fmt.Println(i18n.T(ctx, "chicken.zero", 1))

	// fr
	fmt.Println("Fr: ")
	ctx, err = i18n.WithLocale(ctx, "fr")
	if err != nil {
		log.Println("Error set locale in context: ", err)
	}

	fmt.Println(i18n.T(ctx, "anrsietnai"))
	fmt.Println(i18n.T(ctx, "hello", "You"))
	fmt.Println(i18n.T(ctx, "chicken.zero", 1))
}
