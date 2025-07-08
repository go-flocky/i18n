package i18n

import (
	"fmt"
	"io/fs"
	"strings"

	"gopkg.in/yaml.v3"
)

type Locale struct {
	Code       string
	Name       string
	Dictionary *Dictionary
}

type rawLocale struct {
	Code string         `yaml:"code"`
	Name string         `yaml:"name"`
	Dict map[string]any `yaml:"dict"`
}

func (i *i18n) LoadLocale(localeDir fs.FS) error {
	locale := &Locale{
		Dictionary: &Dictionary{
			ChildDict: make(map[string]*Dictionary),
		},
	}

	err := fs.WalkDir(localeDir, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".yaml") {
			return err
		}

		data, readErr := fs.ReadFile(localeDir, path)
		if readErr != nil {
			return readErr
		}

		var raw rawLocale
		if unmarshalErr := yaml.Unmarshal(data, &raw); unmarshalErr != nil {
			return unmarshalErr
		}

		if locale.Code == "" {
			locale.Code = raw.Code
		}

		if locale.Name == "" {
			locale.Name = raw.Name
		}

		mergeDict(locale.Dictionary, buildDictTree(raw.Dict))
		return nil
	})

	if err != nil {
		return err
	}

	if locale.Code == "" {
		return fmt.Errorf("no locale code found in directory")
	}

	i.RegisterLocale(locale)
	return nil
}

func (i *i18n) LoadLocales() error {
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

		if err := i.LoadLocale(subFS); err != nil {
			return err
		}
	}

	return nil
}
