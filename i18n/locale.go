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

type localeConfig struct {
	Code string `yaml:"code"`
	Name string `yaml:"name"`
}

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

		var translations map[string]any
		if err := yaml.Unmarshal(data, &translations); err != nil {
			return err
		}

		mergeDict(locale.Dictionary, buildDictTree(translations))

		fmt.Println("Dict: ", locale.Dictionary.PrintTree())

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
