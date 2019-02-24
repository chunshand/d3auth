package d3auth

import (

	//	"encoding/json"
	"encoding/json"
	"errors"

	//	"net/url"
	"regexp"
	//	"strconv"
)

//获取登录地址
func (e *Auth_gitee) Get_Rurl(state string) string {
	return "https://gitee.com/oauth/authorize?client_id=" + e.Conf.Appid + "&redirect_uri=" + e.Conf.Rurl + "&response_type=code"
}

//获取token
func (e *Auth_gitee) Get_Token(code string) (*Auth_gitee_succ_res, error) {

	str, err := HttpPost("https://gitee.com/oauth/token?grant_type=authorization_code&code=" + code + "&client_id=" + e.Conf.Appid + "&redirect_uri=" + e.Conf.Rurl + "&client_secret=" + e.Conf.Appkey)
	if err != nil {
		return nil, err
	}

	ismatch, _ := regexp.MatchString("error", str)
	if ismatch {

		return nil, errors.New(str)

	} else {
		p := &Auth_gitee_succ_res{}
		err := json.Unmarshal([]byte(str), p)
		if err != nil {
			return nil, err
		}
		return p, nil
	}

}

//获取第三方用户信息
func (e *Auth_gitee) Get_User_Info(access_token string) (string, error) {

	str, err := HttpGet("https://gitee.com/api/v5/user?access_token=" + access_token)
	if err != nil {
		return "", err
	}
	return string(str), nil

}

//构造方法
func NewAuth_gitee(config *Auth_conf) *Auth_gitee {
	return &Auth_gitee{
		Conf: config,
	}
}
