package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 1. 这里填你的企业ID (我的企业 -> 企业信息 -> 最底下)
const CorpID = "wwe281396b16304b53"

// 2. 这里填你刚刚获取到的 Secret (应用管理 -> 自建应用 -> Secret)
const Secret = "kX9ikjomTEoYPm3IxQB4E1DLdoqPVzB6V-r71CAqQRg"

// 3. 这里填你刚刚在通讯录里看到的账号 (你的员工ID)
const MyUserID = "LuXueCheng"

func main() {
	// === 第一步：拿 Token ===
	tokenUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", CorpID, Secret)
	resp, err := http.Get(tokenUrl)
	if err != nil {
		fmt.Println("获取Token失败:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var tokenResult map[string]interface{}
	json.Unmarshal(body, &tokenResult)
	accessToken := tokenResult["access_token"].(string)
	fmt.Println(">>> 拿到 Token 成功")

	// === 第二步：查你名下的客户列表 ===
	// 接口文档：获取客户列表
	listUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?access_token=%s&userid=%s", accessToken, MyUserID)

	respList, err := http.Get(listUrl)
	if err != nil {
		fmt.Println("获取客户列表失败:", err)
		return
	}
	defer respList.Body.Close()
	listBody, _ := ioutil.ReadAll(respList.Body)

	// === 第三步：打印结果 ===
	fmt.Println(">>> 客户列表数据如下 (外面的就是 ExternalUserID):")
	fmt.Println(string(listBody))
}
