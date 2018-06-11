package slave_sync

import (
	"github.com/influxdata/influxdb/monitor/diagnostics"
	"errors"
)

const (
	// 默认关闭主从同步
	DefaultEnabled = false
)

// 主从同步配置
type Config struct {
	Enabled       		bool  		`toml:"enabled"`
	SlaveUrl      		string 		`toml:"slave-url"`
	SlaveUsername 		string 		`toml:"slave-username"`
	SlavePassword 		string 		`toml:"slave-password"`
}

// 默认配置
func NewConfig() Config {
	return Config{
		Enabled: DefaultEnabled,
	}
}

func (c Config) Diagnostics() (*diagnostics.Diagnostics, error) {
	if !c.Enabled {
		return diagnostics.RowFromMap(map[string]interface{}{
			"enabled": DefaultEnabled,
		}), nil
	}

	return diagnostics.RowFromMap(map[string]interface{}{
		"enabled":               c.Enabled,
		"slave-url":         	 c.SlaveUrl,
		"slave-username":        c.SlaveUsername,
		"slave-password":        c.SlavePassword,
	}), nil
}

func (c Config) Validate() error {
	if c.Enabled && c.SlaveUrl == "" {
		return errors.New("enable slave sync, slave url is must")
	}
	return nil
}