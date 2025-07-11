package i18n

import (
	"context"
	"net/http"
	"strings"
)

func parseAcceptLanguage(header string) []string {
	segments := strings.Split(header, ",")
	localeCodes := make([]string, 0, len(segments))

	for _, segment := range segments {
		code := strings.TrimSpace(strings.Split(segment, ";")[0])
		if code == "*" {
			code = ""
		}
		localeCodes = append(localeCodes, code)
	}
	return localeCodes
}

func (i *I18n) LocaleDetector(r *http.Request) (context.Context, error) {
    header := r.Header.Get("Accept-Language")
    locales := parseAcceptLanguage(header)
    preferredLocale := locales[0]
    
    for _, locale := range locales {
        if i.HasLocale(locale) {
            preferredLocale = locale
            break
        }
    }
    
    return i.WithLocale(r.Context(), preferredLocale)
}
