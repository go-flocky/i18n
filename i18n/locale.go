package i18n

import (
	"fmt"
	"io/fs"
	"strings"

	"gopkg.in/yaml.v3"
)

func (i *I18n) LoadLocale(localeDir fs.FS, code string) error {
	locale := &Locale{
		Code: code,
		Dictionary: &Dictionary{
			ChildDict: make(map[string]*Dictionary),
		},
	}

	err := fs.WalkDir(localeDir, ".", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil || d.IsDir() || !strings.HasSuffix(path, ".yaml") {
			return walkErr
		}

		data, err := fs.ReadFile(localeDir, path)
		if err != nil {
			return err
		}

		var cfg localeConfig
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return err
		}

		if path == fmt.Sprintf("%s.yaml", cfg.Code) {
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
			return err
		}

		translations := make(map[string]any)

		for key, value := range raw {
			switch val := value.(type) {
			case map[string]any:
				b, err := yaml.Marshal(val)
				if err != nil {
					return err
				}
				var dict DictionaryValue
				if err := yaml.Unmarshal(b, &dict); err == nil {
					return err
				} else {
					translations[key] = val
					translations[key] = dict
				}
			default:
				translations[key] = val
			}
		}

		mergeDictionarys(locale.Dictionary, buildDictionaryTree(translations))

		return nil
	})

	if err != nil {
		return err
	}

	i.RegisterLocale(locale)
	return nil
}

func (i *I18n) LoadLocales() error {
	base := *i.localeFS
	entries, err := fs.ReadDir(base, ".")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		subFS, err := fs.Sub(base, entry.Name())
		if err != nil {
			return err
		}

		if err := i.LoadLocale(subFS, entry.Name()); err != nil {
			return err
		}
	}

	return nil
}
