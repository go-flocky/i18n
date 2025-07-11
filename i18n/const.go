package i18n

import "fmt"

const (
	ConfigFileName        = "config.yaml"
)

func MissingTranslation(localeCode, key string) string {
    return fmt.Sprintf("!(MISSING: key=%s.%s)", localeCode, key)
}