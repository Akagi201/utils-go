package conflag

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

func parse(file string, positions ...string) ([]string, error) {
	var conf config

	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	switch filepath.Ext(file) {
	case ".toml":
		conf, err = parseAsToml(r)
	case ".json":
		conf, err = parseAsJSON(r)
	case ".yaml", ".yml":
		conf, err = parseAsYaml(r)
	}
	if err != nil {
		return nil, err
	}
	return conf.toArgs(option{
		boolValue:  BoolValue,
		longHyphen: LongHyphen,
	}, positions...), nil
}

func parseAsToml(r io.Reader) (config, error) {
	var conf config
	_, err := toml.DecodeReader(r, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func parseAsJSON(r io.Reader) (config, error) {
	var conf config
	if err := json.NewDecoder(r).Decode(&conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func parseAsYaml(r io.Reader) (config, error) {
	var conf config

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
