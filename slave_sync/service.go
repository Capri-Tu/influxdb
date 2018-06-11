package slave_sync

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/influxdata/influxdb/models"
)

// 向从机写入数据的服务，根据配置文件设置从机URL，用户名和密码
func WritePoints(bp client.BatchPoints) error {
	if c, e := client.NewHTTPClient(client.HTTPConfig{});e == nil {
		return c.Write(bp)
	} else {
		return e
	}
}

// 将参数转换为client.Write(bp BatchPoints)的参数格式
func ToBatchPoints(database string, retentionPolicy string, consistency string, precision string, points models.Points) client.BatchPoints {
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision:        precision,
		Database:         database,
		RetentionPolicy:  retentionPolicy,
		WriteConsistency: consistency,
	})
	for _, point := range points {
		bp.AddPoint(client.NewPointFrom(point))
	}
	return bp
}

