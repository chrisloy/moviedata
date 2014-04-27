package moviedata

import (
	"testing"
)

func TestGetMovie(t *testing.T) {
    movie := GetMovie(550)
    if id := movie.Id; id != 550 { t.Errorf("Wrong ID [%d]", id) }
    if title := movie.Original_title; title != "Fight Club" { t.Errorf("Wrong Title [%s]", title) }
    if adult := movie.Adult; adult { t.Errorf("Wrong Adult [true]") }
}
