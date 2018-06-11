package slave_sync

import (
	"fmt"
	"testing"
	"runtime"
	"github.com/influxdata/influxdb/models"
)

type ConfigTest struct {
	Enabled       		bool  		`toml:"enabled"`
	SlaveUrl      		string 		`toml:"slave-url"`
	SlaveUsername 		string 		`toml:"slave-username"`
	SlavePassword 		string 		`toml:"slave-password"`
}

func NewConfigTest() *ConfigTest {
	return &ConfigTest{

	}
}

func TestNewConfig(t *testing.T) {
	ct := ConfigTest{}
	fmt.Println(ct)
	ct1 := &ConfigTest{}
	fmt.Println(ct1)
}

func TestReadMem(t *testing.T) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Println(stats.Alloc / 1024)
}


func TestEnum(t *testing.T) {
	fmt.Println(models.ConsistencyLevelOne)
}