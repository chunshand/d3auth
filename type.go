package d3outh

//	"encoding/json"

//基本配置
type Outh_conf struct {
	Appid  string
	Appkey string
	Rurl   string
}

//@ qq 结构 ------------------------------------------------- start

type Outh_qq struct {
	Conf *Outh_conf
}
type Outh_qq_err_res struct {
	Error             int    `json:"error"`
	Error_description string `json:"error_description"`
}
type Outh_qq_me struct {
	Client_ID string `json:"client_id"`
	OpenID    string `json:"openid"`
}

//@ qq 结构 ------------------------------------------------- end

//@ weibo 结构 ------------------------------------------------- start

type Outh_wb struct {
	Conf *Outh_conf
}

type Outh_wb_err_res struct {
	Error             int    `json:"error_code"`
	Error_description string `json:"error"`
}

type Outh_wb_succ_res struct {
	Access_Token string `json:"access_token"`
	Openid       string `json:"uid"`
}

//@ weibo 结构 ------------------------------------------------- end

//@ weixin 结构 ------------------------------------------------- start

type Outh_wx struct {
	Conf *Outh_conf
}

type Outh_wx_err_res struct {
	Error             int    `json:"errcode"`
	Error_description string `json:"errmsg"`
}

type Outh_wx_succ_res struct {
	Access_Token string `json:"access_token"`
	Openid       string `json:"openid"`
}

//@ weixin 结构 ------------------------------------------------- end

//@ github 结构 ------------------------------------------------- start

type Outh_github struct {
	Conf *Outh_conf
}

type Outh_github_err_res struct {
	Error             int    `json:"errcode"`
	Error_description string `json:"errmsg"`
}

type Outh_github_succ_res struct {
	Access_Token string `json:"access_token"`
	Openid       string `json:"openid"`
}

//@ github 结构 ------------------------------------------------- end
