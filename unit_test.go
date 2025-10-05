package AVmerge

import (
	"testing"
)

// go test -v -timeout 10h -run TestClient
func TestClient(t *testing.T) {
	mc := new(MergeConfig)
	mc.MysqlHost = "192.168.5.2"
	mc.MysqlPort = "3306"
	mc.MysqlUser = "root"
	mc.MysqlPassword = "163453"
	Client(mc)
}
