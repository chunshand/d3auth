# d3outh
golang写的简易版第三方登录方法


# 安装

`go get github.com/zcshan/d3outh`

# emmm

	Appid   appid

	Appkey  secret

	Rurl    回调地址
 


# 使用

## qq

qqconf := &d3outh.Outh_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://www.change.tm"}

qqouth := d3outh.NewOuth_qq(qqconf)

fmt.Print(qqouth.Get_Rurl("state")) //获取第三方登录地址

token, err := qqouth.Get_Token("code")  //回调页收的code 获取token


me, err := qqouth.Get_Me(token)  //获取第三方id

Client_ID string `json:"client_id"`
OpenID    string `json:"openid"`

userinfo, _ := wbouth.Get_User_Info(token, me.OpenID)  //获取用户信息 userinfo 是一个json字符串返回

## weibo

wbconf := &d3outh.Outh_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://www.change.tm"}

wbouth := outh.NewOuth_wb(wbconf)

fmt.Print(wbouth.Get_Rurl("state")) //获取第三方登录地址


tokenobj, err := wbouth.Get_Token("code")

Access_Token string `json:"access_token"`
Openid       string `json:"uid"`

userinfo, _ := wbouth.Get_User_Info(tokenobj.Access_Token, tokenobj.Openid)//获取用户信息 userinfo 是一个json字符串返回

## wechat

wxconf := &d3outh.Outh_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://www.change.tm"}

wxouth := outh.NewOuth_wx(wxconf)

fmt.Print(wxouth.Get_Rurl("sate")） //获取第三方登录地址

wxres, err := wxouth.Get_Token("code")

userinfo, _ := wxouth.Get_User_Info(wxres.Access_Token, wxres.Openid) //获取用户信息 userinfo 是一个json字符串返回