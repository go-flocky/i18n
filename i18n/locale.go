package i18n

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/go-flocky/i18n/internal/kvTree"
	"gopkg.in/yaml.v3"
)

func (i *I18n) LoadLocale(localeDir fs.FS, code string) error {
	locale := &Locale{Code: code}

	separator := i.config.KeySeparator

	dict := kvtree.NewTree[string](separator)

	err := fs.WalkDir(localeDir, ".", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		// TODO: nested locale files
		if d.IsDir() || !strings.HasSuffix(path, ".yaml") {
			return nil
		}

		data, err := fs.ReadFile(localeDir, path)
		if err != nil {
			return err
		}

		if path == fmt.Sprintf("%s.yaml", code) {
			var cfg localeConfig
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return fmt.Errorf("failed to parse locale config %s: %w", path, err)
			}

			if locale.Code == "" {
				locale.Code = cfg.Code
			}
			if locale.Name == "" {
				locale.Name = cfg.Name
			}
			return nil
		}

		var raw map[string]any
		if err := yaml.Unmarshal(data, &raw); err != nil {
			return fmt.Errorf("failed to parse translation file %s: %w", path, err)
		}

		flattenMap(raw, "", separator, dict)

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to load locale %s: %w", code, err)
	}

	i.Dictionary[code] = dict
	i.locales[locale.Code] = locale

	return nil
}

func flattenMap(data map[string]any, prefix, separator string, dict *kvtree.KeyValueTree[string]) {
	for k, v := range data {
		var fullKey string
		if prefix != "" {
			fullKey = prefix + separator + k
		} else {
			fullKey = k
		}

		switch val := v.(type) {
		case string:
			dict.Set(fullKey, val)
		case map[string]any:
			flattenMap(val, fullKey, separator, dict)
		default:
			dict.Set(fullKey, fmt.Sprint(val))
		}
	}
}

func (i *I18n) LoadLocales() error {
	potentialLocales, err := fs.ReadDir(i.localeFS, ".")
	if err != nil {
		return fmt.Errorf("failed to read locale directory: %w", err)
	}

	var locales []fs.DirEntry
	for _, entry := range potentialLocales {
		if !entry.IsDir() {
			continue
		}
		locales = append(locales, entry)
	}

	if len(locales) == 0 {
		return fmt.Errorf("no locale directories found")
	}

	for _, entry := range locales {
		entryName := entry.Name()
		subFS, err := fs.Sub(i.localeFS, entryName)
		if err != nil {
			return fmt.Errorf("failed to create filesystem for %s: %w", entryName, err)
		}

		if err := i.LoadLocale(subFS, entryName); err != nil {
			return err
		}
	}

	return nil
}
