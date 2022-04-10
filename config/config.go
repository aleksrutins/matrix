package config

type MatrixConfig struct {
	ProjectName string
	ConfigPath  string
	ConfigRegex struct {
		Find    string
		Replace string
	}
	Targets        map[string][]string
	Configurations map[string]string
}
