package application

import (
    "errors"

    entities "chat/Src/Endpoint/Class/Domain/Entities"
    repository "chat/Src/Endpoint/Class/Domain/Repository"
)

type CreateClassUseCase struct {
    repository repository.IClass
}

func NewCreateClassUseCase(repo repository.IClass) *CreateClassUseCase {
    return &CreateClassUseCase{
        repository: repo,
    }
}

func (u *CreateClassUseCase) Execute(class *entities.Class) (*entities.Class, error) {

    // Validar que endTime > startTime
    if class.EndTime <= class.StartTime {
        return nil, errors.New("La hora de fin debe ser mayor que la hora de inicio")
    }

    // Validar conflicto de horario
    conflict, err := u.repository.HasScheduleConflict(
        class.TeacherID,
        class.StartTime,
        class.EndTime,
        class.ClassDate,
    )
    if err != nil {
        return nil, err
    }

    if conflict {
        return nil, errors.New("El maestro ya tiene una clase asignada en ese horario")
    }

    // Crear clase
    return u.repository.CreateClass(class)
}
