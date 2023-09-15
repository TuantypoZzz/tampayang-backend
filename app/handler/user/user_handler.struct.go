package user_handler

import "time"

type GetAllUserHandlerStructWithTime struct {
	Id           int64      `json:"id"`
	Name         string     `json:"name"`
	Rating       float64    `json:"rating"`
	Booleandesu  int64      `json:"booleandesu"`
	Created      string     `json:"created"`
	Created_date time.Time  `json:"created_date"`
}