package parser

type Parser interface {
	Parse(string) error
	GetString(string) string
}
