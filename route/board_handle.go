package route

import (
	"github.com/gin-gonic/gin"
	"leadboard/model"
	"net/http"
)

func HandleGetBoard(g *gin.Context) {
	g.JSON(200, model.GetLeaderBoard())
}

//TODO:在这里完成返回一个用户提交历史的Handle function
func HandleUserHistory(g *gin.Context) {
	username := g.Param("user")
	err, user := model.GetUserByName(username)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
		})
	} else {
		g.JSON(http.StatusAccepted, gin.H{
			"code": 0,
			"data": user,
		})
	}
}

//TODO:在这里完成接受提交内容，进行评判的handle function
func HandleSubmit(g *gin.Context) {
	type SubmitForm struct {
		UserName string `json:"user"`
		Avatar   string `json:"avatar"`
		Content  string `json:"content"`
	}
	var submit SubmitForm
	err := g.ShouldBindJSON(&submit)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"code": -3,
			"msg":  "提交内容非法呜呜",
		})
	} else if submit.UserName == "" || submit.Content == "" || submit.Avatar == "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "参数不全啊",
		})
	} else if len(submit.UserName) > 255 {
		g.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "用户名太长了",
		})
	} else if len(submit.Content) > 100000 {
		g.JSON(http.StatusBadRequest, gin.H{
			"code": -2,
			"msg":  "图像太大了",
		})
	} else if er, _ := model.CreateSubmission(submit.UserName, submit.Avatar, submit.Content); er != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"code": -3,
			"msg":  "非法内容呜呜呜",
		})
	} else {
		g.JSON(http.StatusAccepted, gin.H{
			"code": 0,
			"msg":  "提交成功",
			"data": model.GetLeaderBoard(),
		})
	}
}
