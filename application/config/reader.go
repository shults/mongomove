package config

type Reader interface {
	Read(path string) (*Config, error)
}
