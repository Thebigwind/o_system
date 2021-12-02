package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func LoginWXSmall(code string) (wxInfo RespWXSmall, err error) {
	//https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code

	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	resp, err := http.Get(fmt.Sprintf(url, AppId, AppSecret, code))
	if err != nil {
		return wxInfo, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&wxInfo)
	if err != nil {
		return wxInfo, err
	}
	//err = tools.BindJson(resp.Body, &wxInfo)
	if err != nil {
	}
	if wxInfo.Errcode != 0 {
		return wxInfo, errors.New(fmt.Sprintf("code: %d, errmsg: %s", wxInfo.Errcode, wxInfo.ErrMsg))
	}
	return wxInfo, nil
}
