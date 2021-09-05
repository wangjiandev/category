package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func GetMysqlConfigFromConsul(consulConfig config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	err := consulConfig.Get(path...).Scan(mysqlConfig)
	if err != nil {
		return nil
	}
	return mysqlConfig

}
