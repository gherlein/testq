/*
*    Test Question for Interviews
*    Copyright 2018 Gregory C. Herlein
*
*    Released under the MIT License:  https://gherlein.mit-license.org/
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
{"Title":"Love","Year":"2015","Rated":"NOT RATED","Released":"30 Oct 2015","Runtime":"135 min","Genre":"Drama, Romance","Director":"Gaspar Noé","Writer":"Gaspar Noé","Actors":"Aomi Muyock, Karl Glusman, Klara Kristin, Ugo Fox","Plot":"Murphy is an American living in Paris who enters a highly sexually and emotionally charged relationship with the unstable Electra. Unaware of the effect it will have on their relationship, they invite their pretty neighbor into their bed.","Language":"English, French","Country":"France, Belgium","Awards":"2 wins & 1 nomination.","Poster":"https://m.media-amazon.com/images/M/MV5BMTQzNDUwODk5NF5BMl5BanBnXkFtZTgwNzA0MDQ2NTE@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"6.0/10"},{"Source":"Rotten Tomatoes","Value":"39%"},{"Source":"Metacritic","Value":"51/100"}],"Metascore":"51","imdbRating":"6.0","imdbVotes":"32,916","imdbID":"tt3774694","Type":"movie","DVD":"15 Mar 2016","BoxOffice":"$176,061","Production":"Alchemy","Website":"N/A","Response":"True"}
*/

type MovieT struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Metascore  string `json:"Metascore"`
	imdbRating string `json:"imdbRating"`
	imdbVotes  string `json:"imdbVotes"`
	imdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
	Ratings    []Rating
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

var (
	movie MovieT
	title string
)

func main() {

	title := flag.String("title", "", "movie title to search for (required)")
	debug := flag.Bool("debug", false, "prints the API query for debugging (optional)")

	flag.Parse()
	if *title == "" {
		fmt.Printf("movie title not specified - usage:  -title=\"<Movie Title>\"\n")
		os.Exit(1)
	}

	key := os.Getenv("OMDAPI_KEY")
	if key == "" {
		fmt.Printf("environment variable OMDAPI_KEY not found - aborting\n")
		os.Exit(1)
	}

	query := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=\"%s\"", key, *title)

	if *debug == true {
		fmt.Println(query)
	}

	response, err := http.Get(query)
	if err != nil {
		fmt.Printf("query failed: %s\n", err)
		os.Exit(1)
	}

	data, _ := ioutil.ReadAll(response.Body)
	//	fmt.Println(string(data))

	if err := json.Unmarshal([]byte(data), &movie); err != nil {
		log.Fatal(err)
	}

	var rating string = "NO-RECORD"
	for _, s := range movie.Ratings {
		if s.Source == "Rotten Tomatoes" {
			rating = s.Value
		}
	}
	fmt.Printf("%s: %s\n", movie.Title, rating)

}
