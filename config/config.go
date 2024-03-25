package config

type IConfig interface {
	SetConfig(cfg any)
	ReadFile(file string) error
	Decode() error
	Defaults() error
	Validate() error
}

var Config IConfig = &YAMLConfig{}

func Load(file string, in any) error {
	var err error
	Config.SetConfig(in)
	err = Config.ReadFile(file)
	if err != nil {
		return err
	}
	err = Config.Decode()
	if err != nil {
		return err
	}
	err = Config.Defaults()
	if err != nil {
		return err
	}
	err = Config.Validate()
	if err != nil {
		return err
	}
	return nil
}
