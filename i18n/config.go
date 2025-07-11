package i18n

import (
	"io"
	"io/fs"

	"gopkg.in/yaml.v3"
)

func parseConfig(fs fs.FS) (*i18nConfig, error) {
	configFile, err := fs.Open(ConfigFileName)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	cfg, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	i18nConfig := new(i18nConfig)
	err = yaml.Unmarshal(cfg, &i18nConfig)
	if err != nil {
		return nil, err
	}
	return i18nConfig, nil
}