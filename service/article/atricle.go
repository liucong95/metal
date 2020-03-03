package article

import (
	"context"
	"fmt"

	"github.com/astaxie/beego/logs"

	pb "metal/protocol"
	"metal/service"
)

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

	rsp, err := service.ZoneClientProxy.GetMomentList(context.TODO(), req)
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

	return rsp.AstMomentDataSetList,rsp.DwTotal,nil
}