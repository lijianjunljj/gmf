package config

type Parser interface {
	Parse(string) error
	Get(string) interface{}
}
