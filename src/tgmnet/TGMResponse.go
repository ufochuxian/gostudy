package tgmnet

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseOld struct {
	Errno   int                    `json:"errno"`
	ErrMsg  string                 `json:"errmsg"`
	TestIds []int                  `json:"test_ids,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

func SuccessResponseOld(c *gin.Context, testIDs []int, params map[string]interface{}) {
	c.JSON(http.StatusOK, ResponseOld{TestIds: testIDs, Params: params})
}

func ErrorResponseOld(c *gin.Context, testIDs []int, params map[string]interface{}) {
	c.JSON(http.StatusOK, ResponseOld{TestIds: testIDs, Params: params})
}

type Callback func(resp ResponseOld)
