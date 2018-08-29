package main

import "fmt"
import "io/ioutil"
import "net/http"
import (
	"strconv"
	"encoding/json"
	"sort"
)

type MovieData struct {
	Poster     string `json:"Poster,omitempty"`
	Title string `json:"title,omitempty"`
	Type string `json:"type,omitempty"`
	Year int `json:"year,omitempty"`
	ImdbID string `json:"imdbID,omitempty"`
}

type MoviePage struct {
	Page     string `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
	Total int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
	Data []MovieData `json:"data,omitempty"`
}

const URL = "https://jsonmock.hackerrank.com/api/movies/search"

func getMovieTitles(substr string) []string {
	curPage := 1
	movies := []string{}
	for {
		movieURL := URL + "/?Title=" + substr + "&page=" + strconv.Itoa(curPage)
		resp, err := http.Get(movieURL)
		if err != nil {
			// handle error
			return movies
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		var moviePage MoviePage
		err = json.Unmarshal(body, &moviePage)
		if err != nil {
			return movies
		}
		for i := 0; i < len(moviePage.Data); i++ {
			movies = append(movies, moviePage.Data[i].Title)
		}
		if curPage == moviePage.TotalPages {
			break
		}
		curPage++
	}
	sort.Strings(movies)
	return movies
}

func main() {
	fmt.Println(getMovieTitles("man"))
}