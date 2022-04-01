package api

import (
	"todo_list/pkg/utils"
	"todo_list/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createTaskService service.CreateTaskService //声明user服务对象
	//校验用户身份
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//绑定服务对象
	if err := c.ShouldBind(&createTaskService); err == nil {
		res := createTaskService.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(500, err)
	}
}

func ShowTask(c *gin.Context) {
	var showTaskService service.ShowTaskService //声明user服务对象
	//校验用户身份
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//绑定服务对象
	if err := c.ShouldBind(&showTaskService); err == nil {
		res := showTaskService.Show(claim.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logrus.Error(err)
		c.JSON(500, err)
	}
}