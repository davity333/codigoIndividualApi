package di

import (
	classController "chat/Src/Endpoint/Class/Infrestructure/Controller"
	enrollmentController "chat/Src/Endpoint/Enrollment/Infrestructure/Controller"
	messageController "chat/Src/Endpoint/Message/Infrestructure/Controller"
	reservationController "chat/Src/Endpoint/Reservations/Infrestructure/Controller"
	userController "chat/Src/Endpoint/User/Infrestructure/Controller"
	contactController "chat/Src/Endpoint/Contacts/Infrestructure/Controller"
)

type ContactDependencies struct {
	GetAllContactsController    *contactController.GetAllContactsController
	CreateContactController    *contactController.CreateContactController
	DeleteContactController    *contactController.DeleteContactController
}

type UserDependencies struct {
	GetAllUserController    *userController.GetAllUserController
	GetUserByNameController *userController.GetUserByNameController
	CreateUserController    *userController.CreateUserController
	LoginUserController     *userController.LoginUserUseCase
	GetTeacherByIDController *userController.GetTeacherByIDController
}

type MessageDependencies struct {
	GetAllMessageByUserController *messageController.GetAllMessageController
	SendMessageByUserController   *messageController.SendMessageController
	DeleteMessageByUserController *messageController.DeleteMessageController
	SubscribeMessageController    *messageController.SubscribeMessageController
}

type ReservationDependencies struct {
	GetAllReservationsController *reservationController.GetAllReservationsController
	GetReservationByIDController *reservationController.GetReservationByIDController
	CreateReservationController  *reservationController.CreateReservationController
	UpdateReservationController  *reservationController.UpdateReservationController
	DeleteReservationController  *reservationController.DeleteReservationController
}

type ClassDependencies struct {
	GetAllClassesController         *classController.GetAllClassesController
	GetClassByIDController          *classController.GetClassByIDController
	GetClassesByTeacherIDController *classController.GetClassesByTeacherIDController
	CreateClassController           *classController.CreateClassController
	UpdateClassController           *classController.UpdateClassController
	DeleteClassController           *classController.DeleteClassController
}

type EnrollmentDependencies struct {
	GetAllEnrollmentsController         *enrollmentController.GetAllEnrollmentsController
	GetEnrollmentsByClassIDController   *enrollmentController.GetEnrollmentsByClassIDController
	GetEnrollmentsByStudentIDController *enrollmentController.GetEnrollmentsByStudentIDController
	CreateEnrollmentController          *enrollmentController.CreateEnrollmentController
	CancelEnrollmentController          *enrollmentController.CancelEnrollmentController
	CompleteEnrollmentController        *enrollmentController.CompleteEnrollmentController
}
