package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ================= 配置区域 =================
const (
	// 1. 填写你的企业ID
	CorpID = "wwe281396b16304b53"
	// 2. 填写拥有客户联系权限的应用 Secret
	CorpSecret = "kX9ikjomTEoYPm3IxQB4E1DLdoqPVzB6V-r71CAqQRg"
	// 3. 填写你要查询的外部联系人ID (以wm开头)
	TargetExternalUserID = "wmxxxxxxxxxxxxxxxxxxxxxx"
)

// ================= 数据结构定义 =================

// AccessTokenResponse 获取Token的响应
type AccessTokenResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// ExternalContactResponse 获取客户详情的响应
type ExternalContactResponse struct {
	ErrCode         int             `json:"errcode"`
	ErrMsg          string          `json:"errmsg"`
	ExternalContact ExternalProfile `json:"external_contact"` // 客户基础信息
	FollowUser      []FollowUser    `json:"follow_user"`      // 跟进该客户的企业成员列表
}

// ExternalProfile 客户基础信息
type ExternalProfile struct {
	ExternalUserid string `json:"external_userid"`
	Name           string `json:"name"`
	Avatar         string `json:"avatar"`
	Type           int    `json:"type"`   // 1表示该外部联系人是微信用户，2表示是企业微信用户
	Gender         int    `json:"gender"` // 0-未知 1-男 2-女
	Unionid        string `json:"unionid"`
}

// FollowUser 跟进成员信息（包含标签）
type FollowUser struct {
	Userid     string `json:"userid"`
	Remark     string `json:"remark"`
	Createtime int64  `json:"createtime"`
	Tags       []Tag  `json:"tags"` // 该成员给客户打的标签
}

type Tag struct {
	GroupName string `json:"group_name"`
	TagName   string `json:"tag_name"`
	Type      int    `json:"type"`
}

// ================= 核心逻辑 =================

func main() {
	// 第一步：获取 Access Token
	token, err := getAccessToken(CorpID, CorpSecret)
	if err != nil {
		fmt.Printf("获取Token失败: %v\n", err)
		return
	}
	fmt.Println("✅ 成功获取 Access Token")

	// 第二步：获取客户详情
	customerData, err := getCustomerDetail(token, TargetExternalUserID)
	if err != nil {
		fmt.Printf("获取客户详情失败: %v\n", err)
		return
	}

	// 第三步：解析并展示数据
	printCustomerInfo(customerData)
}

// getAccessToken 调用微信API获取凭证
func getAccessToken(corpID, secret string) (string, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpID, secret)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result AccessTokenResponse
	json.Unmarshal(body, &result)

	if result.ErrCode != 0 {
		return "", fmt.Errorf("API错误: %d - %s", result.ErrCode, result.ErrMsg)
	}
	return result.AccessToken, nil
}

// getCustomerDetail 调用获取客户详情接口
func getCustomerDetail(token, externalUserID string) (*ExternalContactResponse, error) {
	// 注意：这里没有处理 cursor 分页，默认拉取全部或第一页
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=%s&external_userid=%s", token, externalUserID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// 用于调试，打印原始JSON
	// fmt.Println("API Raw Response:", string(body))

	var result ExternalContactResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, fmt.Errorf("API错误: %d - %s", result.ErrCode, result.ErrMsg)
	}

	return &result, nil
}

// printCustomerInfo 格式化输出你需要的数据
func printCustomerInfo(data *ExternalContactResponse) {
	contact := data.ExternalContact

	fmt.Println("\n================ 客户档案 ================")
	fmt.Printf("姓名: %s\n", contact.Name)
	genderStr := "未知"
	if contact.Gender == 1 {
		genderStr = "男"
	}
	if contact.Gender == 2 {
		genderStr = "女"
	}
	fmt.Printf("性别: %s\n", genderStr)
	fmt.Printf("头像URL: %s\n", contact.Avatar)
	fmt.Printf("UnionID (用于关联游戏账号): %s\n", contact.Unionid)

	fmt.Println("\n---------------- 标签聚合 ----------------")
	// 聚合去重逻辑
	uniqueTags := make(map[string]string) // key="GroupName-TagName" 用于去重

	for _, user := range data.FollowUser {
		// 打印每个跟进人的备注
		t := time.Unix(user.Createtime, 0)
		fmt.Printf("> 跟进员工: %s (添加时间: %s, 备注: %s)\n", user.Userid, t.Format("2006-01-02 15:04:05"), user.Remark)

		// 收集标签
		for _, tag := range user.Tags {
			key := fmt.Sprintf("%s-%s", tag.GroupName, tag.TagName)
			if _, exists := uniqueTags[key]; !exists {
				uniqueTags[key] = fmt.Sprintf("[%s] %s", tag.GroupName, tag.TagName)
			}
		}
	}

	fmt.Println("\n>>> 最终去重后的标签列表:")
	if len(uniqueTags) == 0 {
		fmt.Println("暂无标签")
	} else {
		for _, tagStr := range uniqueTags {
			fmt.Println(tagStr)
		}
	}
	fmt.Println("==========================================")
}
