package router

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hewenyu/gin-simple/Models/response"
	"github.com/hewenyu/toolspackage/release/zlog"
)

func Notify(c *gin.Context) {

	var wxn WX_NOTIFY

	zlog.Zap().Info(c.Request.Method)

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		zlog.Zap().Info("read http body failed！error msg:" + err.Error())

		response.FailWithMessage(err.Error(), c)
		return
	}

	zlog.Zap().Info(string(body))

	dec := json.NewDecoder(strings.NewReader(string(body)))

	if err := dec.Decode(&wxn); err == io.EOF {
		response.FailWithMessage(err.Error(), c)
		return
	} else if err != nil {
		zlog.Zap().Info("read http body failed！error msg:" + err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	if wxn.EventType == "TRANSACTION.SUCCESS" {

		res, err := wxn.UnSignature()

		if err != nil {
			zlog.Zap().Info("read http body failed！error msg:" + err.Error())
			response.FailWithMessage(err.Error(), c)
			return
		}
		// 用户订单更新
		zlog.Zap().Info("get orfer:" + res.OutTradeNo)
	}

	response.OkWithDetailed("", "获取成功", c)
}
