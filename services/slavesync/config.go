package slavesync

import (
	"github.com/influxdata/influxdb/monitor/diagnostics"
)

// Config represents the meta configuration.
type Config struct {
	Enabled       		bool  		`toml:"enabled"`
	SlaveUrl      		string 		`toml:"slave-url"`
	SlaveUsername 		string 		`toml:"slave-username"`
	SlavePassword 		string 		`toml:"slave-password"`
}

// Diagnostics returns a diagnostics representation of a subset of the Config.
func (c Config) Diagnostics() (*diagnostics.Diagnostics, error) {
	if !c.Enabled {
		return diagnostics.RowFromMap(map[string]interface{}{
			"enabled": false,
		}), nil
	}

	return diagnostics.RowFromMap(map[string]interface{}{
		"enabled":               false,
		"slave-url":         	 c.SlaveUrl,
		"slave-username":        c.SlaveUsername,
		"slave-password":        c.SlavePassword,
	}), nil
}
