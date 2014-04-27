package moviedata

import "fmt"

// Movie type
type Movie struct {
	Adult bool
	Id int
	Original_title string
}

func (m Movie) String() string {
	return fmt.Sprintf("Adult: %t\nId: %d\nTitle: %s", m.Adult, m.Id, m.Original_title)
}