package i18n

type Dictionary struct {
	Value     string
	ChildDict map[string]*Dictionary
}
