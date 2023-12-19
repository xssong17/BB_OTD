package miyoushe

import (
	"BB_OTD/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
@Project ：崩崩查水表
@File    ：factory.go
@IDE     ：GoLand
@Author  ：xssong
@Date    ：2023/12/17 20:54
*/

const (
	Server     = "gf01" // 崩坏2服务器代号，官服 gf01
	UserAgent  = "Mozilla/5.0 (Linux; Android 7.1.2; vmos Build/NZH54D; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/81.0.4044.117 Mobile Safari/537.36 miHoYoBBS"
	AppVersion = "2.42.1"
	ClientType = "5"
)

// getDS
//
//	@Description: 计算ds参数
//	@param params
//	@return string
func getDS(params map[string]string) string {
	generateDs := func(salt string) (string, error) {

		t := strconv.FormatInt(time.Now().Unix(), 10)
		r := strconv.Itoa(rand.Intn(100000) + 100000)
		b := ""
		q := utils.EncodeParams(params)
		s := fmt.Sprintf("salt=%s&t=%s&r=%s&b=%s&q=%s", salt, t, r, b, q)
		h := md5.New()
		_, err := h.Write([]byte(s))
		if err != nil {
			return "", err
		}
		c := hex.EncodeToString(h.Sum(nil))
		return fmt.Sprintf("%s,%s,%s", t, r, c), nil
	}

	salt := "xV8v4Qu54lUKrEYFZkJhB8cuOh9Asafs"
	ds, err := generateDs(salt)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return ds
}

// GetInfoFactory
//
//	@Description: 从米游社接口获取信息-工厂
//	@param url 请求api
//	@param cookie 登录cookie信息
//	@param server 服务器信息
//	@param roleId 角色uid
//	@return []byte 响应信息
//	@return error 错误信息
func GetInfoFactory(url, cookie, server, roleId string) ([]byte, error) {

	params := map[string]string{
		"role_id": roleId,
		"server":  server,
	}

	ds := getDS(params) //  获取ds参数
	if ds == "" {
		return nil, nil
	}

	headers := map[string]string{
		"DS":                ds,
		"x-rpc-app_version": AppVersion,
		"x-rpc-page":        "3.1.3_#/bh2",
		"User-Agent":        fmt.Sprintf("%s/%s", UserAgent, ClientType),
		"x-rpc-client_type": ClientType,
		"Cookie":            cookie,
	}
	return utils.Get(url, params, headers) //  请求接口

}
