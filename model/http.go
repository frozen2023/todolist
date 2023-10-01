package model

type CommonResult struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	STATUS_OK   = 200
	STATUS_FAIL = 404
	MSG_OK      = "操作成功"
	MSG_FAIL    = "操作失败"
)

func OkResult(data interface{}) CommonResult {
	return CommonResult{Code: STATUS_OK, Msg: MSG_OK, Data: data}

}

func FailResult() CommonResult {
	return CommonResult{Code: STATUS_FAIL, Msg: MSG_FAIL, Data: nil}
}
