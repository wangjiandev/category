package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/wangjiandev/category/common"
	"github.com/wangjiandev/category/domain/repository"
	myService "github.com/wangjiandev/category/domain/service"
	"github.com/wangjiandev/category/handler"
	category "github.com/wangjiandev/category/proto/category"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		// 设置服务地址和端口
		micro.Address("127.0.0.1:8082"),
		// 添加consul作为注册中心
		micro.Registry(consulRegistry),
	)
	fmt.Println(consulConfig)
	// 获取mysql配置
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	fmt.Println(mysqlConfig)
	db, err := gorm.Open("mysql", mysqlConfig.User+":"+mysqlConfig.Password+"@tcp(127.0.0.1:3306)/"+mysqlConfig.Database+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// 禁止复数表
	db.SingularTable(true)

	// Initialise service
	service.Init()

	// Initialise table
	rp := repository.NewCategoryRepository(db)
	rp.InitTable()

	categoryDataService := myService.NewCategoryDataService(repository.NewCategoryRepository(db))
	// Register Handler
	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		log.Error(err)
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
