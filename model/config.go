package model

type Config struct {
	Path        string   `koanf:"path"`
	ExcludeDirs []string `koanf:"exclude-dirs"`
	Recursive   bool     `koanf:"recursive"`
	OutputFile  string   `koanf:"output"`
}
