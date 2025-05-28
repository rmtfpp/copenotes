package config

type (
	Config struct {
		//HTTP     HTTPConfig
		//App      AppConfig
		//Cache    CacheConfig
		Database DatabaseConfig
		//Files    FilesConfig
		//Tasks    TasksConfig
		//Mail     MailConfig
	}

	DatabaseConfig struct {
		Driver         string
		Connection     string
		TestConnection string
	}
)

func GetConfig() Config {

	c := Config{
		Database: DatabaseConfig{
			Driver:         "sqlite3",
			Connection:     "dbs/copenotes.db",
			TestConnection: "",
		},
	}

	return c

}
