package controller

import (
	"altman/common/response"
	"altman/global"
	"altman/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Index(c *gin.Context) {
	p := c.DefaultQuery("p", "1")
	size := c.DefaultQuery("size", "20")

	page, _ := strconv.Atoi(p)
	limit, _ := strconv.Atoi(size)
	offset := (page - 1) * limit

	var adminList []models.Admin

	global.Wg.Add(1)
	go func() {
		defer global.Wg.Done()
		//global.DB.Limit(200).Find(&adminList)
		global.DB.Limit(limit).Offset(offset).Find(&adminList)
	}()
	global.Wg.Wait()

	response.Success(c, adminList)
}
