package miyoushe

import (
	"github.com/tidwall/gjson"
)

/*
@Project ：崩崩查水表
@File    ：roleInfo.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2023/12/17 17:16
*/

type Role struct {
	UID      string //	角色uid
	Name     string //	角色名
	Level    int64  //	角色等级
	Coins    int64  //	水晶
	Club     string //	社团信息
	VIPLevel int64  //	高贵的vip等级（
}

func GetRoleInfo(cookie, server, roleId string) (Role, error) {
	//  请求接口获取人物信息
	rsp, err := GetInfoFactory(RoleInfoUrl, cookie, server, roleId)
	if err != nil {
		return Role{}, nil
	}
	//  格式化信息
	role := Role{
		UID:      roleId,
		Name:     gjson.GetBytes(rsp, "data.role.nickname").String(),
		Level:    gjson.GetBytes(rsp, "data.role.level").Int(),
		Coins:    gjson.GetBytes(rsp, "data.stats.h_coin").Int(),
		VIPLevel: gjson.GetBytes(rsp, "data.stats.vip_level").Int(),
		Club:     "xxxx", //  TODO 结合团本或团战接口补充社团信息
	}

	return role, nil
}
