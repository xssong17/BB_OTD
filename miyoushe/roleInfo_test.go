package miyoushe

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"strconv"
	"testing"
)

/*
@Project ：崩崩查水表
@File    ：roleInfo_test.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2023/12/19 23:20
*/

func TestGetRole(t *testing.T) {
	//  以下cookie和uid请自行更改，cookie获取方法将在后续进行补充
	cookie := ""
	server := ServerGF
	uid := ""
	role, err := GetRoleInfo(cookie, server, uid)
	if err != nil {
		t.Fatalf("%v", err)
	}
	//  格式化输出表格展示
	table, err := gotable.Create("角色名", "uid", "等级", "水晶数量", "vip")
	if err != nil {
		t.Fatalf("%v", err)
	}
	rowList := []string{role.Name, role.UID, strconv.Itoa(int(role.Level)), strconv.Itoa(int(role.Coins)), strconv.Itoa(int(role.VIPLevel))}
	err = table.AddRow(rowList)
	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Println(table)

}
