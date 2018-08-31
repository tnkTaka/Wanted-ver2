package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Wanted-ver2/controller"
)

func studentRouter(student *gin.RouterGroup) {
	//生徒プロフィール
	student.GET("/profile/:student_id", controller.GetByStudentId)
	//生徒プロフィール更新
	student.PUT("/profile/update/:student_id", nil)
	//生徒ログイン
	student.POST("/login", controller.LoginStudent)
}
