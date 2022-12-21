package service

import (
	"skyshi/features/activity"
)

type Service struct {
	do activity.DataInterface
}

func New(data activity.DataInterface) activity.ServiceInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) GetActivities() ([]activity.Activity, string, error) {
	data, msg, err := service.do.GetActivities()
	return data, msg, err
}

func (service *Service) GetActivity(id uint) (activity.Activity, string, error) {
	data, msg, err := service.do.GetActivity(id)
	return data, msg, err
}

func (service *Service) Create(core activity.Activity) (activity.Activity, string, error) {
	data, msg, err := service.do.Create(core)
	return data, msg, err
}

func (service *Service) Delete(id uint) (string, error) {
	msg, err := service.do.Delete(id)
	return msg, err
}

func (service *Service) Update(id uint, core activity.Activity) (activity.Activity, string, error) {
	msg, err := service.do.Update(id, core)
	if err != nil {
		return activity.Activity{}, msg, err
	}

	data, msg2, err2 := service.do.GetActivity(id)
	return data, msg2, err2
}
