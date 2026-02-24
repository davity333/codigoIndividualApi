package router

import (
	controller "chat/Src/Endpoint/Enrollment/Infrestructure/Controller"

	"github.com/gin-gonic/gin"
)

type EnrollmentRouter struct {
	getAllEnrollmentsController         *controller.GetAllEnrollmentsController
	getEnrollmentsByClassIDController   *controller.GetEnrollmentsByClassIDController
	getEnrollmentsByStudentIDController *controller.GetEnrollmentsByStudentIDController
	createEnrollmentController          *controller.CreateEnrollmentController
	cancelEnrollmentController          *controller.CancelEnrollmentController
	completeEnrollmentController        *controller.CompleteEnrollmentController
}

func NewEnrollmentRouter(
	getAllEnrollmentsController *controller.GetAllEnrollmentsController,
	getEnrollmentsByClassIDController *controller.GetEnrollmentsByClassIDController,
	getEnrollmentsByStudentIDController *controller.GetEnrollmentsByStudentIDController,
	createEnrollmentController *controller.CreateEnrollmentController,
	cancelEnrollmentController *controller.CancelEnrollmentController,
	completeEnrollmentController *controller.CompleteEnrollmentController,
) *EnrollmentRouter {
	return &EnrollmentRouter{
		getAllEnrollmentsController:         getAllEnrollmentsController,
		getEnrollmentsByClassIDController:   getEnrollmentsByClassIDController,
		getEnrollmentsByStudentIDController: getEnrollmentsByStudentIDController,
		createEnrollmentController:          createEnrollmentController,
		cancelEnrollmentController:          cancelEnrollmentController,
		completeEnrollmentController:        completeEnrollmentController,
	}
}

func (r *EnrollmentRouter) RegisterRoutes(g *gin.Engine) {
	enrollmentGroup := g.Group("/api/v1/enrollments")
	{
		enrollmentGroup.GET("/getAll", r.getAllEnrollmentsController.GetAllEnrollments)
		enrollmentGroup.GET("/class/:classId", r.getEnrollmentsByClassIDController.GetEnrollmentsByClassID)
		enrollmentGroup.GET("/student/:studentId", r.getEnrollmentsByStudentIDController.GetEnrollmentsByStudentID)
		enrollmentGroup.POST("/create", r.createEnrollmentController.CreateEnrollment)
		enrollmentGroup.PUT("/cancel/:enrollmentId", r.cancelEnrollmentController.CancelEnrollment)
		enrollmentGroup.PUT("/complete/:enrollmentId", r.completeEnrollmentController.CompleteEnrollment)
	}
}
