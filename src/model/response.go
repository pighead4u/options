package model

/**
 * http响应类
 **/
type OptionResponse struct {
	Data      interface{} `json:"data"`
	IsSuccess bool        `json:"isSuccess"`
	Msg       string      `json:"msg"`
}
