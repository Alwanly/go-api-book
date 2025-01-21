package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

func LoadConfig(configName string) (GlobalConfig, error) {
	// Load default config
	loadDefaults()

	// Load config from file
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(configName)

	// load from environment variables
	viper.AutomaticEnv()

	// -- binding environment variables as config source
	emptyConfig := &GlobalConfig{}

	// count the number of fields (config) on this struct
	fieldCount := reflect.TypeOf(emptyConfig).Elem().NumField()

	for i := 0; i < fieldCount; i++ {
		// get the field tag name, for example `mapstructure:"ENV"`
		field := string(reflect.TypeOf(emptyConfig).Elem().Field(i).Tag)

		// trim the tag and just return the tag name, for example `ENV`
		// then bind it as env to load to viper config
		if err := viper.BindEnv(field[14 : len(field)-1]); err != nil {
			fmt.Println("Error bind env.", err)
		}
	}

	// read config from sources (config file and environment variables)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Error reading config file:", err)
			return GlobalConfig{}, err
		}
	}

	// unmarshal config
	c := GlobalConfig{}
	if err := viper.Unmarshal(&c); err != nil {
		return GlobalConfig{}, err
	}

	return c, nil
}
