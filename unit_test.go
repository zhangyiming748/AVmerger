package AVmerge

import (
	"testing"
)

// go test -v -run  TestRemoveSrc
func TestAndroid2PC(t *testing.T) {
	mc := new(MergeConfig)
	mc.MysqlUser = "root"
	mc.MysqlPassword = "163453"
	mc.MysqlHost = "192.168.5.2"
	mc.MysqlPort = "3306"
	Android2PC(mc)
}
