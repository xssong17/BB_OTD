package miyoushe

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"strconv"
	"testing"
)

/*
@Project ：崩崩查水表
@File    ：tuanzhan_test.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2023/12/19 23:25
*/

func TestGetTuanZhanInfo(t *testing.T) {
	//  以下cookie和uid请自行更改，cookie获取方法将在后续进行补充
	cookie := ""
	server := ServerGF
	uid := ""
	tuanZhanInfo, err := GetTuanZhanInfo(cookie, server, uid)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//  格式化输出表格展示
	table, err := gotable.Create("时间点", "位置", "耗时", "贡献率")
	if err != nil {
		t.Fatalf("%v", err)
	}

	for _, action := range tuanZhanInfo.History {
		minutes := action.EnterTime / 60
		remainingSeconds := action.EnterTime % 60
		enterTime := fmt.Sprintf("%02d:%02d", minutes, remainingSeconds)
		rowList := []string{enterTime, action.Point, strconv.Itoa(int(action.PassTime)), strconv.Itoa(int(action.ContributeRate))}
		err = table.AddRow(rowList)
		if err != nil {
			t.Fatalf("%v", err)
		}
	}
	fmt.Println(table)
}
