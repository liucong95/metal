package article

import (
	"context"
	"fmt"
	"time"
	"encoding/json"

	"github.com/astaxie/beego"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"github.com/astaxie/beego/logs"

	pb "metal/protocol"
)

//gZoneHttpClientProxy 对外
var gZoneHttpClientProxy pb.GZoneHttpClientProxy
//gZoneHttpInnerClientProxy 对内
var gZoneHttpInnerClientProxy pb.GZoneHttpInnerClientProxy

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
	gZoneHttpInnerClientProxy = pb.NewGZoneHttpInnerClientProxy(opts...)

	//对外接口
	opts = append(opts, client.WithTarget("ip://127.0.0.1:8000"))
	gZoneHttpClientProxy = pb.NewGZoneHttpClientProxy(opts...)
}

//GetMomentList 心情列表
func GetMomentList(reqType int32, roleID uint64, targetID uint64, topicID uint64, start, end uint32) ([]*pb.MomentDataSet, uint32, error){
	//查看个人空间心情
	req := &pb.GetMomentListReq{
			DwReqType: reqType,
			UllRoleID: roleID,
			UllTargetID: targetID,
			UllTopicID: topicID,
			DwStart: start,
			DwEnd: end,
		}

	rsp, err := gZoneHttpClientProxy.GetMomentList(context.TODO(), req)
	if err != nil{
		logs.Error("err = %v, req:%s", err, req.String())
		return nil,0,err
	}
	
	if rsp.ErrNo != 0 {
		logs.Error("req = %s\n, rsp = %s", req, rsp.String())
		return nil,0,fmt.Errorf("getmoment error:%s", rsp.ErrMsg)
	}

	if rsp.AstMomentDataSetList == nil{
		return nil,0,nil
	}
	
	_,err = json.Marshal(rsp.AstMomentDataSetList)
	if err != nil{
		logs.Error("json encoding err:%s",err)
		return nil,0,err
	}

	return rsp.AstMomentDataSetList,rsp.DwTotal,nil
}