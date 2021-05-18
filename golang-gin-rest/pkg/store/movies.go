package store

import (
	"errors"
	"math/rand"

	"github.com/dkapanidis/app-starters/golang-gin-rest/pkg/models"
)

// ListMovies lists movies
func ListMovies() []models.Movie {
	return movies
}

// CreateMovie creates a movie
func CreateMovie(movie models.Movie) (models.Movie, error) {
	// make sure payload does not contain id
	if movie.ID != 0 {
		return movie, errors.New("invalid id")
	}
	movie.ID = uint(rand.Uint32())
	movies = append(movies, movie)
	return movie, nil
}

// GetMovie gets a movie
func GetMovie(id uint) (*models.Movie, error) {
	for _, movie := range movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, errors.New("movie not found")
}

// UpdateMovie updates a movie
func UpdateMovie(id uint, movie models.Movie) (*models.Movie, error) {
	// make sure payload contains correct id
	if id != movie.ID {
		return nil, errors.New("invalid id")
	}
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(append(movies[:index], movie), movies[index+1:]...)
			return &movie, nil
		}
	}
	return nil, errors.New("movie not found")
}

// DeleteMovie deletes a movies
func DeleteMovie(id uint) {
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
