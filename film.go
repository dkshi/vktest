package vktest

type Film struct {
	FilmID      int     `json:"film_id" db:"film_id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	ReleaseDate string  `json:"release_date" db:"relese_date"`
	Rating      float64 `json:"rating" db:"rating"`
	Actors      []Actor `json:"actors" db:"-"`
}
