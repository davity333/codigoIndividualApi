package di

import (
	classRouter "chat/Src/Endpoint/Class/Infrestructure/Router"
	enrollmentRouter "chat/Src/Endpoint/Enrollment/Infrestructure/Router"
	messageRouter "chat/Src/Endpoint/Message/Infrestructure/Router"
	reservationRouter "chat/Src/Endpoint/Reservations/Infrestructure/Router"
	userRouter "chat/Src/Endpoint/User/Infrestructure/Router"
	contactRouter "chat/Src/Endpoint/Contacts/Infrestructure/Router"

	"github.com/gin-gonic/gin"
)

func InitRoutes(g *gin.Engine) error {
	userDeps, err := InitializeUserDependencies()
	if err != nil {
		return err
	}

	messageDeps, err := InitializeMessageDependencies()
	if err != nil {
		return err
	}

	reservationDeps, err := InitializeReservationDependencies()
	if err != nil {
		return err
	}

	classDeps, err := InitializeClassDependencies()
	if err != nil {
		return err
	}

	enrollmentDeps, err := InitializeEnrollmentDependencies()
	if err != nil {
		return err
	}

	contactDeps, err := InitializeContactDependencies()
	if err != nil {
		return err
	}

	contactRouter.RouterContact(
		g,
		contactDeps.GetAllContactsController,
		contactDeps.CreateContactController,
		contactDeps.DeleteContactController,
	)

	userRouter.RouterUser(
		g,
		userDeps.GetAllUserController,
		userDeps.GetUserByNameController,
		userDeps.CreateUserController,
		userDeps.LoginUserController,
		userDeps.GetTeacherByIDController,
	)

	messageRouter.MessageRouter(
		g,
		messageDeps.GetAllMessageByUserController,
		messageDeps.SendMessageByUserController,
		messageDeps.DeleteMessageByUserController,
		messageDeps.SubscribeMessageController,
	)

	reservationRouter.ReservationRouter(
		g,
		reservationDeps.GetAllReservationsController,
		reservationDeps.GetReservationByIDController,
		reservationDeps.CreateReservationController,
		reservationDeps.UpdateReservationController,
		reservationDeps.DeleteReservationController,
	)

	classRouterInstance := classRouter.NewClassRouter(
		classDeps.GetAllClassesController,
		classDeps.GetClassByIDController,
		classDeps.GetClassesByTeacherIDController,
		classDeps.CreateClassController,
		classDeps.UpdateClassController,
		classDeps.DeleteClassController,
	)
	classRouterInstance.RegisterRoutes(g)

	enrollmentRouterInstance := enrollmentRouter.NewEnrollmentRouter(
		enrollmentDeps.GetAllEnrollmentsController,
		enrollmentDeps.GetEnrollmentsByClassIDController,
		enrollmentDeps.GetEnrollmentsByStudentIDController,
		enrollmentDeps.CreateEnrollmentController,
		enrollmentDeps.CancelEnrollmentController,
		enrollmentDeps.CompleteEnrollmentController,
	)
	enrollmentRouterInstance.RegisterRoutes(g)

	return nil
}
