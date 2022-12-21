package service

import "skyshi/features/todo"

type Service struct {
	do todo.DataInterface
}

func New(data todo.DataInterface) todo.ServiceInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) GetAllTodo() ([]todo.Todo, string, error) {
	data, msg, err := service.do.GetAllTodo()
	return data, msg, err
}

func (service *Service) GetATodo(id uint) (todo.Todo, string, error) {
	data, msg, err := service.do.GetATodo(id)
	return data, msg, err
}

func (service *Service) Create(data todo.Todo) (todo.Todo, string, error) {
	data, msg, err := service.do.Create(data)
	return data, msg, err
}

func (service *Service) Delete(id uint) (string, error) {
	msg, err := service.do.Delete(id)
	return msg, err
}

func (service *Service) Update(id uint, core todo.Todo) (todo.Todo, string, error) {
	msg, err := service.do.Update(id, core)
	if err != nil {
		return todo.Todo{}, msg, err
	}

	data, msg2, err2 := service.do.GetATodo(id)
	return data, msg2, err2
}
