package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		var user SignUserInfo

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	router.Run(":8080")
}

/*
常用驗證器:

	required 必填欄位 (不能為空，且不能沒有這個欄位)，例: binding:"required"

針對字串的長度:
	min 最小長度
	max 最大長度
	len 固定長度

針對數字的大小:
	eq 等於
	ne 不等於
	gt 大於
	gte 大於等於
	lt 小於
	lte 小於等於

針對同級欄位的:
	eqfield 等於其他欄位的值
	nefield 不等於其他欄位的值

列舉:
	oneof=red green	，只能是 red 或 green

字串:
	contains	包含某字串
	excludes	不包含某字串
	startswith	字串前綴
	endswith	字串後綴

陣列:
	dive	dive後面的驗證就是針對陣列中的每一個元素

網路驗證:
	ip
	ipv4
	ipv6
	uri
	url

日期驗證:
	datetime=2006-01-02 15:04:05

其他:
	- 忽略欄位，例: binding:"-"

*/

type SignUserInfo struct {
	Name       string   `json:"name" binding:"required"`
	Age        int      `json:"age" binding:"lt=30,gt=18"`
	Password   string   `json:"password"`
	RePassword string   `json:"re_password" binding:"eqfield=Password"`
	Sex        string   `json:"sex" binding:"oneof=man woman"`
	LikeList   []string `json:"like_list" binding:"required,dive,startswith=like"`
	IP         string   `json:"ip" binding:"ip"`
	Url        string   `json:"url" binding:"url"`
	Uri        string   `json:"uri" binding:"uri"`
	Date       string   `json:"date" binding:"datetime=2006-01-02 15:04:05"`
}
