package slavesync

import "github.com/influxdata/influxdb/client/v2"

type SlaveSyn struct {
	Enabled     bool
	url         string
	username    string
	password    string
	Config      *Config
	slaveClient client.Client
}

func NewSlaveSyn(c Config) *SlaveSyn {
	ss := &SlaveSyn{
		Enabled:  c.Enabled,
		url:      c.SlaveUrl,
		username: c.SlaveUsername,
		password: c.SlavePassword,
	}
	return ss
}

func (ss SlaveSyn) initSlaveClient() (*SlaveSyn, error) {
	ss.slaveClient.Write()// TODO 写入
}