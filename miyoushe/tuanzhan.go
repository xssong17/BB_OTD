package miyoushe

import (
	"github.com/tidwall/gjson"
)

/*
@Project ：崩崩查水表
@File    ：tuanzhan.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2023/12/17 20:54
*/
var (
	//  点位映射
	PointMap = map[string]string{
		"1437": "A", // 诡雷迷城（七图）
		"1438": "B",
		"1439": "C",
		"1440": "D",
		"1441": "E",
		"1442": "F",
		"1443": "G",
		"1444": "H",
		"1445": "I",

		"1446": "A", // 诡雷迷城（八图）
		"1447": "B",
		"1448": "C",
		"1449": "D",
		"1450": "E",
		"1451": "F",
		"1452": "G",
		"1453": "H",
		"1454": "I",
		"1455": "J",

		"1401": "A", // 普通模式（七图）
		"1402": "B",
		"1403": "C",
		"1404": "D",
		"1405": "E",
		"1406": "F",
		"1407": "G",
		"1408": "H",
		"1409": "I",

		"1411": "A", //普通模式（九图）
		"1412": "B",
		"1413": "C",
		"1414": "D",
		"1415": "E",
		"1416": "F",
		"1417": "G",
		"1418": "H",
		"1419": "I",
		"1420": "J",
		"1421": "K",

		"1422": "A", // 资源争夺战（七图）
		"1423": "B",
		"1424": "C",
		"1425": "D",
		"1426": "E",
		"1427": "F",
		"1428": "G",
	}

	//  胜负结果映射
	OutcomeMap = map[string]string{
		"1": "胜利",
		"2": "失败",
	}
)

type Action struct {
	EnterTime      int64  //	进图时间
	Point          string //	点位
	PassTime       int64  //	耗时
	ContributeRate int64  //	贡献值
}

type TuanZhanItem struct {
	Name1        string   //	我方社团名
	Name2        string   //	敌方社团名
	MapID        int64    //	地图ID
	Status       string   //	胜负
	History      []Action //	行动信息
	MapName      string   //	地图名称
	MapFigureUrl string   //	地图地址
}

// GetTuanZhanInfo
//
//	@Description: 获取团战信息
//	@param cookie 登录cookie
//	@param server 服务器信息
//	@param roleId 角色uid
//	@return TuanZhanItem 团战信息item
//	@return error 错误信息
func GetTuanZhanInfo(cookie, server, roleId string) (TuanZhanItem, error) {

	//  请求api获取团战信息
	resp, err := GetInfoFactory(TuanZhanURl, cookie, server, roleId)
	if err != nil {
		return TuanZhanItem{}, nil
	}
	//  解析数据
	//ClubName := gjson.GetBytes(resp, "data.profile.name").String() //  社团名
	//ClubID := gjson.GetBytes(resp, "data.profile.id").Int()        //  社团ID
	//fmt.Printf("%-20s%-20s\n%-20s%-20d\n\n", "社团名", "社团ID", ClubName, ClubID)

	history := gjson.GetBytes(resp, "data.profile.battle_info.history").Array() //  获取进攻行为列表
	if len(history) == 0 {
		return TuanZhanItem{}, nil
	}

	//  解析最近一场对狙
	lastHistoryObj := history[0]
	lastHistoryArray := []Action{}
	historyArray := lastHistoryObj.Get("history").Array()

	//  格式化信息
	for _, each := range historyArray {
		action := Action{
			EnterTime:      each.Get("enter_time").Int(),
			Point:          PointMap[each.Get("level_id").String()],
			PassTime:       each.Get("pass_time").Int(),
			ContributeRate: each.Get("contribute_rate").Int() / 10,
		}
		lastHistoryArray = append(lastHistoryArray, action)
	}
	tuanZhanItem := TuanZhanItem{
		Name1:        lastHistoryObj.Get("name1").String(),
		Name2:        lastHistoryObj.Get("name2").String(),
		MapID:        lastHistoryObj.Get("map_id").Int(),
		Status:       OutcomeMap[lastHistoryObj.Get("status").String()],
		History:      lastHistoryArray,
		MapName:      lastHistoryObj.Get("map_name").String(),
		MapFigureUrl: lastHistoryObj.Get("map_figure_url").String(),
	}
	return tuanZhanItem, nil
}
