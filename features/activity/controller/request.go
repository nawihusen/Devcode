package controller

import "skyshi/features/activity"

type Request struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

func (r *Request) ToCore() activity.Activity {
	return activity.Activity{
		Title: r.Title,
		Email: r.Email,
	}
}
