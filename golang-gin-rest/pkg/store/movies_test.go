package store

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dkapanidis/app-starters/golang-gin-rest/pkg/models"
)

// TestCreateMovieOK creates a movie correctly
func TestCreateMovieOK(t *testing.T) {
	// Prepare
	movie := models.Movie{
		Name: "godfather",
	}
	movies = []models.Movie{}

	// Command
	movie, err := CreateMovie(movie)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, movie.Name, movies[0].Name, "the movie is not stored properly to db")
}

// TestCreateMovieErrorInvalidID creates a movie with invalid id
func TestCreateMovieErrorInvalidID(t *testing.T) {
	// Prepare
	movie := models.Movie{
		ID:   1,
		Name: "godfather",
	}
	movies = []models.Movie{}

	// Command
	movie, err := CreateMovie(movie)

	// Assert
	assert.Error(t, err, "invalid id")
}

// TestGetMovieOK gets a movie correctly
func TestGetMovieOK(t *testing.T) {
	// Prepare
	movie := models.Movie{
		ID:   1234,
		Name: "godfather",
	}
	movies = []models.Movie{movie}

	// Command
	response, err := GetMovie(movie.ID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, movie.Name, response.Name, "the movie is not retrieved properly from db")
}

// TestGetMovieErrorNotFound gets a movie that does not exist
func TestGetMovieErrorNotFound(t *testing.T) {
	// Prepare
	movies = []models.Movie{}

	// Command
	_, err := GetMovie(1)

	// Assert
	assert.Error(t, err, "record not found")
}

// TestUpdateMovieOK updates a movie correctly
func TestUpdateMovieOK(t *testing.T) {
	// Prepare
	movie := models.Movie{
		ID:   1234,
		Name: "godfather",
	}
	movies = []models.Movie{movie}

	// Command
	movie.Name = "godfather 2"
	_, err := UpdateMovie(movie.ID, movie)

	// Assert
	assert.NoError(t, err)
}

// TestUpdateMovieErrorNotFound updates a movie that does not exist
func TestUpdateMovieErrorNotFound(t *testing.T) {
	// Prepare
	movie := models.Movie{
		ID:   1,
		Name: "godfather 2",
	}
	movies = []models.Movie{}

	// Command
	_, err := UpdateMovie(movie.ID, movie)

	// Assert
	assert.Error(t, err, "record not found")
}

// TestUpdateMovieErrorInvalidID updates a movie with invalid id
func TestUpdateMovieErrorInvalidID(t *testing.T) {
	// Prepare
	movie := models.Movie{
		ID:   1234,
		Name: "godfather",
	}
	movies = []models.Movie{movie}

	// Command
	_, err := UpdateMovie(1, movie)

	// Assert
	assert.Error(t, err, "invalid id")
}

// TestListMoviesOK lists movies
func TestListMoviesOK(t *testing.T) {
	// setup
	movies = []models.Movie{}
	movie := models.Movie{
		Name: "godfather",
	}

	// movies list is empty
	movies := ListMovies()
	assert.Len(t, movies, 0)

	// create a movie
	movie, err := CreateMovie(movie)
	assert.NoError(t, err)

	// movies list has a movie
	movies = ListMovies()
	assert.Len(t, movies, 1)
	assert.Equal(t, movie.Name, movies[0].Name)
}

// TestDeleteMovieOK deletes a movie
func TestDeleteMovieOK(t *testing.T) {
	// setup
	movie := models.Movie{
		ID:   1234,
		Name: "godfather",
	}
	movies = []models.Movie{movie}

	// execute
	DeleteMovie(movie.ID)

	// assert
	_, err := GetMovie(movie.ID)
	assert.Error(t, err, "record not found")
}
