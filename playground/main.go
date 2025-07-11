package main

import (
	"context"
	"fmt"
	internalI18n "github.com/go-flocky/i18n/i18n"
	"github.com/go-flocky/i18n/playground/locales"
	"log"
	"net/http"
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

	fmt.Println(i18n.ListLocaleCodes())

	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(i18n.T(r.Context(), "hello", r.URL.Query().Get("user"))))
	})

	rs := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, err := i18n.LocaleDetector(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	log.Println("Server on port :3000?user=changeme")
	log.Fatalln(http.ListenAndServe(":3000", rs(r)))
}
