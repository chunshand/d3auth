package d3auth

import (
	//	"encoding/json"
	"errors" //	"fmt"
	//	"net/url"
	"regexp"
	//	"strconv"
)

//获取登录地址
func (e *Auth_github) Get_Rurl(state string) string {
	return "https://github.com/login/oauth/authorize?client_id=" + e.Conf.Appid + "&redirect_uri=" + e.Conf.Rurl + "&state=" + state
}

//获取token
func (e *Auth_github) Get_Token(code string) (string, error) {
	str, err := HttpGet("https://github.com/login/oauth/access_token?client_id=" + e.Conf.Appid + "&client_secret=" + e.Conf.Appkey + "&code=" + code + "&redirect_uri=" + e.Conf.Rurl)
	if err != nil {
		return "", err
	}

	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {

		return "", errors.New(str)

	} else {
		re, _ := regexp.Compile("access_token=(.*)&scope")
		newres := re.FindStringSubmatch(str)
		if len(newres) >= 2 {
			return newres[1], nil
		}
		return "", nil
	}

}

//获取第三方用户信息
func (e *Auth_github) Get_User_Info(access_token string) (string, error) {

	str, err := HttpGet("https://api.github.com/user?access_token=" + access_token)
	if err != nil {
		return "", err
	}
	return string(str), nil

}

//构造方法
func NewAuth_github(config *Auth_conf) *Auth_github {
	return &Auth_github{
		Conf: config,
	}
}
