package service

import (
	"time"

	"github.com/astaxie/beego"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"github.com/astaxie/beego/logs"

	pb "metal/protocol"
)

//gZoneHttpClientProxy 对外
var ZoneClientProxy pb.GZoneHttpClientProxy
//gZoneHttpInnerClientProxy 对内
var ZoneInnerClientProxy pb.GZoneHttpInnerClientProxy

func init() {
	protol := beego.AppConfig.String("troc_protocol")
	netwrok := beego.AppConfig.String("trpc_netwrok")
	target := beego.AppConfig.String("trpc_server")
	namespace := beego.AppConfig.String("trpc_server")
	serverName := beego.AppConfig.String("trpc_name")
	timeout, err := beego.AppConfig.Int("trpc_timeout")
	if err != nil{
		logs.Error("tiemout config err:%s",err)
		timeout = 2
	}

	opts := []client.Option{
		client.WithProtocol(protol),
		client.WithNetwork(netwrok),
		client.WithTarget(target),
		client.WithNamespace(namespace),
		client.WithServiceName(serverName),
		client.WithTimeout(time.Second * time.Duration(timeout)),
	}

	//对内接口
	ZoneInnerClientProxy = pb.NewGZoneHttpInnerClientProxy(opts...)

	//对外接口
	opts = append(opts, client.WithTarget("ip://127.0.0.1:8000"))
	ZoneClientProxy = pb.NewGZoneHttpClientProxy(opts...)
}