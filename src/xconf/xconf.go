
package xconf

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
)

type Configure struct {
	LogLevel	string
	LogDir		string
}

func Load(path string, config any) error {
	_, err := toml.DecodeFile(path, config)
	return err
}

func (s *Configure)String(config any) string {
	info, _ := json.MarshalIndent(s, "", "    ")
	return string(info)
}