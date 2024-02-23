package util

import (
	"fmt"
	"time"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/timeout"
)

const (
	SendTimeout = time.Second * 10
	DialTimeout = time.Second * 10
)

func ApiPost(urlPath string, header map[string]string, req interface{}, apiRsp interface{}) error {
	rsp, err := gentleman.
		New().
		Post().
		URL(urlPath).
		SetHeaders(header).
		Use(body.JSON(req)).
		Use(timeout.Request(SendTimeout)).
		Send()
	if err != nil {
		return err
	}
	if !rsp.Ok {
		return fmt.Errorf("code:%d,body:%s", rsp.StatusCode, string(rsp.Bytes()))
	}
	return rsp.JSON(apiRsp)
}

func ApiGet(urlPath string, header map[string]string, params map[string]string, apiRsp interface{}) error {
	rsp, err := gentleman.New().
		Get().
		URL(urlPath).
		SetQueryParams(params).
		SetHeaders(header).
		Use(timeout.Request(SendTimeout)).
		Send()
	if err != nil {
		return err
	}
	if !rsp.Ok {

		return fmt.Errorf("code:%d,body:%s", rsp.StatusCode, string(rsp.Bytes()))
	}
	return rsp.JSON(apiRsp)
}
