package router

import (
	controller "chat/Src/Endpoint/Class/Infrestructure/Controller"

	"github.com/gin-gonic/gin"
)

type ClassRouter struct {
	getAllClassesController         *controller.GetAllClassesController
	getClassByIDController          *controller.GetClassByIDController
	getClassesByTeacherIDController *controller.GetClassesByTeacherIDController
	createClassController           *controller.CreateClassController
	updateClassController           *controller.UpdateClassController
	deleteClassController           *controller.DeleteClassController
}

func NewClassRouter(
	getAllClassesController *controller.GetAllClassesController,
	getClassByIDController *controller.GetClassByIDController,
	getClassesByTeacherIDController *controller.GetClassesByTeacherIDController,
	createClassController *controller.CreateClassController,
	updateClassController *controller.UpdateClassController,
	deleteClassController *controller.DeleteClassController,
) *ClassRouter {
	return &ClassRouter{
		getAllClassesController:         getAllClassesController,
		getClassByIDController:          getClassByIDController,
		getClassesByTeacherIDController: getClassesByTeacherIDController,
		createClassController:           createClassController,
		updateClassController:           updateClassController,
		deleteClassController:           deleteClassController,
	}
}

func (r *ClassRouter) RegisterRoutes(g *gin.Engine) {
	classGroup := g.Group("/api/v1/classes")
	{
		classGroup.GET("/getAll", r.getAllClassesController.GetAllClasses)
		classGroup.GET("/:classId", r.getClassByIDController.GetClassByID)
		classGroup.GET("/teacher/:teacherId", r.getClassesByTeacherIDController.GetClassesByTeacherID)
		classGroup.POST("/create", r.createClassController.CreateClass)
		classGroup.PUT("/update", r.updateClassController.UpdateClass)
		classGroup.DELETE("/:classId", r.deleteClassController.DeleteClass)
	}
}
