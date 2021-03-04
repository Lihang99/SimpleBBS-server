package main

import (
	"SimpleBBS-server/controller"
	"SimpleBBS-server/dao/mysql"
	"SimpleBBS-server/dao/redis"
	"SimpleBBS-server/logger"
	"SimpleBBS-server/router"
	"SimpleBBS-server/settings"
	snowflake "SimpleBBS-server/utils/snowflake"
	"fmt"
)

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("need config file.eg: SimpleBBS config.yaml")
	//	return
	//}
	//// 加载配置
	//if err := settings.Init(os.Args[1]); err != nil {
	//	fmt.Printf("load config failed, err:%v\n", err)
	//	return
	//}
	// 加载配置
	if err := settings.Init("./conf/dev.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	//初始化id生成器
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("Init id gennerator failed,err:%v\n", err)
		return
	}
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	//注册路由
	r := router.SetupRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("Run server failed , err : %v\n", err)
		return
	}
}
