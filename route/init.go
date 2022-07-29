package route

import (
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	//TODO:register your route here
	//for example:
	//r.POST("/create-user",HandleCreateUser)，这个接口不是要求中的，仅仅作为示例
	//r.GET("/leaderboard",HandleGetBoard)
	r.GET("/leaderboard", HandleGetBoard)
	r.GET("/history/:user", HandleUserHistory)
	r.POST("/submit", HandleSubmit)
	v := r.Group("/vote", CheckUserAgent)
	{
		v.POST("", HandleVote)
	}

	return r
}
