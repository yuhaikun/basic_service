package main

import (
	"basic_service/config"
	"basic_service/dao/mysql"
	basic "basic_service/kitex_gen/basic/foundation"
	"flag"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	// "os"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "f", "./config/config.yaml", "配置文件的路径")
	flag.Parse()

	if err := config.Init(filePath); err != nil {
		klog.Errorf("init config failed,err:%v\n", err)
		return
	}
	if err := mysql.InitDB(); err != nil {
		klog.Errorf("init mysql failed,err:%v\n", err)
		return
	}

	//svr := basic.NewServer(new(FoundationImpl))

	r, err := etcd.NewEtcdRegistry(config.GlobalServerConfig.EtcdInfo.Endpoints)
	if err != nil {
		log.Fatal(err)
	}
	// ip := os.Getenv("PUBLIC_IP")

	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%s", config.GlobalServerConfig.Port))

	server := basic.NewServer(new(FoundationImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "Foundation",
	}), server.WithServiceAddr(addr))

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
