package config

type Config struct {
	Source      DatabaseConnection `yaml:"source"`
	Destination DatabaseConnection `yaml:"destination"`
	BulkSize    uint               `yaml:"bulkSize"`
	Query       string             `yaml:"query"`
}

type DatabaseConnection struct {
	Database   string
	Collection string
}
