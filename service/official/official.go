package official

import (
	"context"
	"fmt"

	"github.com/astaxie/beego/logs"

	pb "metal/protocol"
	"metal/service"
)

//CreateOfficialAcount 创建官方账号
func CreateOfficialAcount(name , head, signture string) (uint64, error){
	req := &pb.CreateOfficialAcountReq{
		SzRoleName: name,
		SzHeadFrame: head,
		SzSignature: signture,
	}

	rsp, err := service.ZoneInnerClientProxy.CreateOfficialAcount(context.TODO(), req)
	if err != nil{
		logs.Error("err = %v, req:%s", err, req.String())
		return 0,err
	}
	
	if rsp.ErrNo != 0 {
		logs.Error("req = %s\n, rsp = %s", req, rsp.String())
		return 0,fmt.Errorf("getmoment error:%s", rsp.ErrMsg)
	}

	return rsp.UllRoleID,nil
}
//GetOfficialAcountList 心情列表
func GetOfficialAcountList() ([]*pb.ZoneProfileSet, int, error){
	//查看个人空间心情
	req := &pb.GetOfficialAcountListReq{}

	rsp, err := service.ZoneInnerClientProxy.GetOfficialAcountList(context.TODO(), req)
	if err != nil{
		logs.Error("err = %v, req:%s", err, req.String())
		return nil,0,err
	}
	
	if rsp.ErrNo != 0 {
		logs.Error("req = %s\n, rsp = %s", req, rsp.String())
		return nil,0,fmt.Errorf("getmoment error:%s", rsp.ErrMsg)
	}

	if rsp.AstAcountList == nil{
		return nil,0,nil
	}

	return rsp.AstAcountList,len(rsp.AstAcountList),nil
}