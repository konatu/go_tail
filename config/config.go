package config

import (
	"fmt"
	"io/ioutil"

	validator "gopkg.in/go-playground/validator.v9"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	LoadConfigFromYAML("config.yaml")
}

type Config struct {
	File []File `yaml:"config"`
}

type File struct {
	Name      *string    `yaml:"name" validate:"required"`
	LogConfig *LogConfig `yaml:"log_config"`
}

type LogConfig struct {
	Path   *string `yaml:"path"`
	format *string `yaml:"format"`
}

func LoadConfigFromYAML(path string) (*File, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data: %+v\n", string(data))

	var con Config
	err = yaml.Unmarshal(data, &con)
	if err != nil {
		panic(err)
	}
	fmt.Printf("d: %+v", con)
	return nil, nil
}

func (f *File) Validate() error {
	v := validator.New()
	err := v.Struct(f)
	return err
}
