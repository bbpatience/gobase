package authorization

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	response "gobase/api/common"
	"gobase/global/my_errors"
	"gobase/global/variables"
)

type HeaderParams struct {
	key string `header:"appid"`
	nonce string `header:"nonce"`
	timestamp string `header:"timestamp"`
	sign string `header:"sign"`
}

func CheckAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		headerParams := HeaderParams{}
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			variables.ZapLog.Error(my_errors.ErrorsAuthHeaderFail, zap.Error(err))
			context.Abort()
		}

		secret := "haha" //  TODO  load from db with app key.
		toSign := headerParams.nonce + "|" + secret + "|" + headerParams.timestamp
		if toSign == headerParams.sign {
			context.Next()
		} else {
			response.ErrorAuthFail(context)
		}
	}
}
