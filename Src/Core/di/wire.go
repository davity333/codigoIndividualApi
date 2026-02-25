//go:build wireinject
// +build wireinject

package di

import (
	"chat/Src/Core/sse"
	classApplication "chat/Src/Endpoint/Class/Application"
	classRepository "chat/Src/Endpoint/Class/Domain/Repository"
	classController "chat/Src/Endpoint/Class/Infrestructure/Controller"
	classSql "chat/Src/Endpoint/Class/Infrestructure/Sql"
	enrollmentApplication "chat/Src/Endpoint/Enrollment/Application"
	enrollmentRepository "chat/Src/Endpoint/Enrollment/Domain/Repository"
	enrollmentController "chat/Src/Endpoint/Enrollment/Infrestructure/Controller"
	enrollmentSql "chat/Src/Endpoint/Enrollment/Infrestructure/Sql"
	messageApplication "chat/Src/Endpoint/Message/Application"
	messageRepository "chat/Src/Endpoint/Message/Domain/Repository"
	messageController "chat/Src/Endpoint/Message/Infrestructure/Controller"
	messageSql "chat/Src/Endpoint/Message/Infrestructure/Sql"
	reservationApplication "chat/Src/Endpoint/Reservations/Application"
	reservationRepository "chat/Src/Endpoint/Reservations/Domain/Repository"
	reservationController "chat/Src/Endpoint/Reservations/Infrestructure/Controller"
	reservationSql "chat/Src/Endpoint/Reservations/Infrestructure/Sql"
	userApplication "chat/Src/Endpoint/User/Application"
	userRepository "chat/Src/Endpoint/User/Domain/Repository"
	userController "chat/Src/Endpoint/User/Infrestructure/Controller"
	userSql "chat/Src/Endpoint/User/Infrestructure/Sql"

	contactApplication "chat/Src/Endpoint/Contacts/Application"
	contactController "chat/Src/Endpoint/Contacts/Infrestructure/Controller"
	contactSql "chat/Src/Endpoint/Contacts/Infrestructure/Sql"

	"github.com/google/wire"
)

func InitializeContactDependencies() (*ContactDependencies, error) {
	wire.Build(
		contactSql.NewMySQL,
		wire.Bind(new(contactApplication.IContact), new(*contactSql.Mysql)),
		contactApplication.NewGetAllContactsUseCase,
		contactApplication.NewCreateContactUseCase,
		contactApplication.NewDeleteContactUseCase,
		contactController.NewGetAllContactsController,
		contactController.NewCreateContactController,
		contactController.NewDeleteContactController,
		wire.Struct(new(ContactDependencies), "*"),
	)

	return nil, nil
}

func InitializeUserDependencies() (*UserDependencies, error) {
	wire.Build(
		userSql.NewMySQL,
		wire.Bind(new(userRepository.IUser), new(*userSql.Mysql)),
		userApplication.NewGetAllUsersUseCase,
		userApplication.NewGetUserByNameUseCase,
		userApplication.NewCreateUserUseCase,
		userApplication.NewLoginUserUseCase,
		userController.NewGetAllUserController,
		userController.NewGetUserByNameController,
		userController.NewCreateUserController,
		userController.NewLoginUserUseCase,
		wire.Struct(new(UserDependencies), "*"),
	)

	return nil, nil
}

func InitializeBroadcaster() *sse.Broadcaster {
	return sse.NewBroadcaster()
}

// ProvideSendMessageUseCase creates SendMessageUseCase with broadcaster configured
func ProvideSendMessageUseCase(repository messageRepository.IMessage, broadcaster *sse.Broadcaster) *messageApplication.SendMessageUseCase {
	usecase := messageApplication.NewSendMessageUseCase(repository)
	usecase.SetBroadcaster(broadcaster)
	return usecase
}

func InitializeMessageDependencies() (*MessageDependencies, error) {
	wire.Build(
		InitializeBroadcaster,
		messageSql.NewMySQL,
		wire.Bind(new(messageRepository.IMessage), new(*messageSql.Mysql)),
		messageApplication.NewGetMessageByUserIdUseCase,
		ProvideSendMessageUseCase,
		messageApplication.NewDeleteMessageUseCase,
		messageController.NewGetAllMessageController,
		messageController.NewSendMessageController,
		messageController.NewDeleteMessageController,
		messageController.NewSubscribeMessageController,
		wire.Struct(new(MessageDependencies), "*"),
	)

	return nil, nil
}

func InitializeReservationDependencies() (*ReservationDependencies, error) {
	wire.Build(
		reservationSql.NewMySQL,
		wire.Bind(new(reservationRepository.IReservation), new(*reservationSql.Mysql)),
		reservationApplication.NewGetAllReservationsUseCase,
		reservationApplication.NewGetReservationByIDUseCase,
		reservationApplication.NewCreateReservationUseCase,
		reservationApplication.NewUpdateReservationUseCase,
		reservationApplication.NewDeleteReservationUseCase,
		reservationController.NewGetAllReservationsController,
		reservationController.NewGetReservationByIDController,
		reservationController.NewCreateReservationController,
		reservationController.NewUpdateReservationController,
		reservationController.NewDeleteReservationController,
		wire.Struct(new(ReservationDependencies), "*"),
	)

	return nil, nil
}

func InitializeClassDependencies() (*ClassDependencies, error) {
	wire.Build(
		classSql.NewClassSQL,
		wire.Bind(new(classRepository.IClass), new(*classSql.ClassSQL)),
		classApplication.NewGetAllClassesUseCase,
		classApplication.NewGetClassByIDUseCase,
		classApplication.NewGetClassesByTeacherIDUseCase,
		classApplication.NewCreateClassUseCase,
		classApplication.NewUpdateClassUseCase,
		classApplication.NewDeleteClassUseCase,
		classController.NewGetAllClassesController,
		classController.NewGetClassByIDController,
		classController.NewGetClassesByTeacherIDController,
		classController.NewCreateClassController,
		classController.NewUpdateClassController,
		classController.NewDeleteClassController,
		wire.Struct(new(ClassDependencies), "*"),
	)

	return nil, nil
}

func InitializeEnrollmentDependencies() (*EnrollmentDependencies, error) {
	wire.Build(
		enrollmentSql.NewEnrollmentSQL,
		wire.Bind(new(enrollmentRepository.IEnrollment), new(*enrollmentSql.EnrollmentSQL)),
		enrollmentApplication.NewGetAllEnrollmentsUseCase,
		enrollmentApplication.NewGetEnrollmentsByClassIDUseCase,
		enrollmentApplication.NewGetEnrollmentsByStudentIDUseCase,
		enrollmentApplication.NewCreateEnrollmentUseCase,
		enrollmentApplication.NewCancelEnrollmentUseCase,
		enrollmentApplication.NewCompleteEnrollmentUseCase,
		enrollmentController.NewGetAllEnrollmentsController,
		enrollmentController.NewGetEnrollmentsByClassIDController,
		enrollmentController.NewGetEnrollmentsByStudentIDController,
		enrollmentController.NewCreateEnrollmentController,
		enrollmentController.NewCancelEnrollmentController,
		enrollmentController.NewCompleteEnrollmentController,
		wire.Struct(new(EnrollmentDependencies), "*"),
	)

	return nil, nil
}
