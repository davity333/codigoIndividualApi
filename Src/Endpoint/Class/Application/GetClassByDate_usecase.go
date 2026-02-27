package application

import (
    entities "chat/Src/Endpoint/Class/Domain/Entities"
    repository "chat/Src/Endpoint/Class/Domain/Repository"
)

type GetClassesByDateUseCase struct {
    repo repository.IClass
}

func NewGetClassesByDateUseCase(repo repository.IClass) *GetClassesByDateUseCase {
    return &GetClassesByDateUseCase{repo: repo}
}

func (uc *GetClassesByDateUseCase) Execute(date string) ([]entities.ClassWithTeacher, error) {
    return uc.repo.GetClassesByDate(date)
}
