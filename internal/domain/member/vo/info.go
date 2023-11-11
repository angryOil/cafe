package vo

import "time"

type Info struct {
	Id        int
	UserId    int
	Nickname  string
	CreatedAt time.Time
}
