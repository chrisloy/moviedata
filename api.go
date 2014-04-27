package moviedata

import (
	"os"
	"net/http"
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"encoding/json"
)

const api = "https://api.themoviedb.org/3"
var key = "api_key=" + getKey()

// Get a movie by ID
func GetMovie(id int) Movie {
	url := fmt.Sprintf("%s/movie/%d?%s", api, id, key)
	resp, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	var movie Movie
	json.Unmarshal(body, &movie)
	return movie
}

func getKey() string {
	key, err := ioutil.ReadFile(os.Getenv("HOME") + "/.apikeys/moviedatabase.com")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(key))
}