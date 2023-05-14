// Package verify 校验
package verify

import (
	"encoding/json"
	"fmt"

	"github.com/link1st/go-stress-testing/model"
)

// WebSocketResponseJSON 返回数据结构体，返回值为json
type WebSocketResponseJSON struct {
	Seq      string `json:"seq"`
	Cmd      string `json:"cmd"`
	Response struct {
		Code    int         `json:"code"`
		CodeMsg string      `json:"codeMsg"`
		Data    interface{} `json:"data"`
	} `json:"response"`
}

// WebSocketJSON 通过返回的Body 判断
// 返回示例: {"seq":"1566276523281-585638","cmd":"heartbeat","response":{"code":200,"codeMsg":"Success","data":null}}
// code 取body中的返回code
func WebSocketJSON(request *model.Request, seq string, msg []byte) (code int, isSucceed bool) {
	responseJSON := &WebSocketResponseJSON{}
	err := json.Unmarshal(msg, responseJSON)
	if err != nil {
		code = model.ParseError
		fmt.Printf("请求结果 json.Unmarshal msg:%s err:%v", string(msg), err)
	} else {
		  isSucceed = true
		}
	// 开启调试模式
	if request.GetDebug() {
		fmt.Printf("请求结果 seq:%s body:%s \n", seq, string(msg))
	}
	return
}
