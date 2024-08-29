package author

import "time"

type Author struct {
	ID           int
	Name         string
	Surname      string
	BirthCountry Country
	DateOfBirth  time.Time
}

type Country struct {
	Name string
}
