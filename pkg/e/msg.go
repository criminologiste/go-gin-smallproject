package e

var MsgFlag = map[int]string{
	SUCCESS:                        "OK",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "TOKEN授权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "TOKEN超时",
	ERROR_AUTH_TOKEN:               "TOKEN生成失败",
	ERROR_AUTH:                     "TOKEN错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[ERROR]
}
