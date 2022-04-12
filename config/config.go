package config

type MatrixConfig struct {
	ProjectName string
	ConfigPath  string
	ConfigRegex struct {
		Find    string
		Replace string
	}
	Move []struct {
		Old string
		New string
	}
	Targets        map[string][]string
	Configurations map[string]string
}
