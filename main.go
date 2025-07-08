package main

import (
	"fmt"

	flockyI18n "github.com/go-flocky/i18n/i18n"
)

func main() {
	i18n := flockyI18n.NewI18n()
	i18n.FallbackLocaleCode = []string{"en", "de"}

	i18n.RegisterLocale(&flockyI18n.Locale{
		Code: "en",
		Dictionary: &flockyI18n.Dictionary{
			ChildDict: map[string]*flockyI18n.Dictionary{
				"hello": {
					Value: "Hello",
					ChildDict: map[string]*flockyI18n.Dictionary{
						"world":    {Value: "World"},
						"fallback": {Value: "Fallback"},
					},
				},
			},
		},
	})

	i18n.RegisterLocale(&flockyI18n.Locale{
		Code: "fr",
		Dictionary: &flockyI18n.Dictionary{
			ChildDict: map[string]*flockyI18n.Dictionary{
				"hello": {
					Value: "Bonjour",
					ChildDict: map[string]*flockyI18n.Dictionary{
						"world":    {Value: "Monde"},
						"fallback": {Value: "Retour"},
					},
				},
			},
		},
	})

	i18n.RegisterLocale(&flockyI18n.Locale{
		Code: "de",
		Dictionary: &flockyI18n.Dictionary{
			ChildDict: map[string]*flockyI18n.Dictionary{
				"hello": {
					Value: "Hallo",
					ChildDict: map[string]*flockyI18n.Dictionary{
						"world": {Value: "Welt"},
					},
				},
			},
		},
	})

	fmt.Println("En: ")
	fmt.Println(i18n.T("en", ""))
	fmt.Println(i18n.T("en", "hello"))
	fmt.Println(i18n.T("en", "hello.world"))
	fmt.Println(i18n.T("en", "hello.fallback"))
	fmt.Println("")

	fmt.Println("De: ")
	fmt.Println(i18n.T("de", ""))
	fmt.Println(i18n.T("de", "hello"))
	fmt.Println(i18n.T("en", "hello.world"))
	fmt.Println(i18n.T("de", "hello.fallback"))
	fmt.Println("")

	fmt.Println("Fr: ")
	fmt.Println(i18n.T("fr", ""))
	fmt.Println(i18n.T("fr", "hello"))
	fmt.Println(i18n.T("fr", "hello.world"))
	fmt.Println(i18n.T("fr", "hello.fallback"))
}
