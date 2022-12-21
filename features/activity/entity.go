package activity

import "time"

type Activity struct {
	ID         uint
	Email      string
	Title      string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

type ServiceInterface interface {
	GetActivities() ([]Activity, string, error)
	GetActivity(id uint) (Activity, string, error)
	Create(Activity) (Activity, string, error)
	Delete(id uint) (string, error)
	Update(id uint, core Activity) (Activity, string, error)
}

type DataInterface interface {
	GetActivities() ([]Activity, string, error)
	GetActivity(id uint) (Activity, string, error)
	Create(Activity) (Activity, string, error)
	Delete(id uint) (string, error)
	Update(id uint, core Activity) (string, error)
}
