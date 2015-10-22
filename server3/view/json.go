package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
)

type resultType struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func Json(handler func(*gin.Context)(interface{},error) )(func (*gin.Context)){
	return func (c *gin.Context){
		result := &resultType{};
		handlerResult,error := handler(c);
		if error != nil{
			result.Code = 1;
			result.Msg = error.Error();
		}else{
			result.Code = 0;
			result.Msg = "";
			result.Data = handlerResult;
		}
		resultString,error := json.Marshal(&result);
		if error != nil{
			c.String(http.StatusOK,"{code:1,msg:'json序列化失败',data:''}");
		}else{
			c.String(http.StatusOK,string(resultString));
		}
	}
}