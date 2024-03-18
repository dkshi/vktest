package vktest

type Actor struct {
	ActorID     int    `json:"actor_id" db:"actor_id"`
	Name        string `json:"name" db:"name"`
	Gender      string `json:"gender" db:"gender"`
	DateOfBirth string `json:"birth_date" db:"birth_date"`
	Films       []Film `json:"films" db:"-"`
}
