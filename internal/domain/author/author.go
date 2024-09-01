package author

import "time"

type Author struct {
	ID           int
	Name         string
	Surname      string
	BirthCountry string
	DateOfBirth  time.Time
}
