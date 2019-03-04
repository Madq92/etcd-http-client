package config

var Config = struct {
	APPName string `default:"app name"`

	Etcd struct {
		hosts    []string
		Username string `default:"root"`
		Password string `default:"root"`
	}

	MySql struct {
		Url      string
		User     string `default:"root"`
		Password string `default:"123456"`
		Port     uint   `default:"3306"`
	}

	Log struct {
		Level string
	}
}{}
