package models

import (
	"skyshi/features/activity"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Email string
	Title string
}

func (a *Activity) ToCore() activity.Activity {
	return activity.Activity{
		ID:         a.ID,
		Email:      a.Email,
		Title:      a.Title,
		Created_at: a.CreatedAt,
		Updated_at: a.UpdatedAt,
		Deleted_at: a.DeletedAt.Time,
	}
}

func ToCoreList(data []Activity) []activity.Activity {
	var result []activity.Activity
	for _, v := range data {
		result = append(result, v.ToCore())
	}

	return result
}

func CoreToModel(data activity.Activity) Activity {
	return Activity{
		Email: data.Email,
		Title: data.Title,
	}
}
