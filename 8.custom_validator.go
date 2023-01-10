package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}

	router.POST("/", func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": _GetValidMsg(err, &user)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	router.Run(":8080")
}

type User struct {
	Name string `json:"name" binding:"required, sign" msg:"用戶名驗證失敗"`
	Age  int    `json:"age" binding:"required" msg:"請輸入年齡"`
}

func signValid(fl validator.FieldLevel) bool {
	nameList := []string{"123", "456", "789"}
	for _, nameStr := range nameList {
		name := fl.Field().Interface().(string)
		if name == nameStr {
			return false
		}
	}
	return true
}

// GetValidMsg 返回結構體中 tag 的 msg 參數; 使用時，需要傳 obj 的指針
func _GetValidMsg(err error, obj any) string {
	// 將 err 接口斷言為指定類型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 斷言成功
		getObj := reflect.TypeOf(obj)
		for _, e := range errs {
			// 循環每一個錯誤訊息
			// 根據報錯欄位名，取得結構體的具體屬性
			if f, ok := getObj.Elem().FieldByName(e.Field()); ok {
				return f.Tag.Get("msg")
			}
		}
	}
	return err.Error()
}
