package conf

import (
	"ecosmos-api/common/logger"
	"testing"
)

// go test -test.run Test_yaml -v
func Test_yaml(t *testing.T) {
	InitConf("../../api/conf/")

	logger.Print(Conf.Default.Jwt.Expire)
	logger.Print(Conf.Default.Jwt.Secret)

	logger.Print(GetString("jwt::secret"))
	logger.Print(GetInt64("jwt::expire"))
}
