package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// getOneMovie gets a single movie from the database
func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, movie, "movie"); err != nil {
		app.errorJSON(w, err)
		return
	}
}

// getAllMovies gets all movies from the database
func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, movies, "movies"); err != nil {
		app.errorJSON(w, err)
		return
	}
}

// getAllGenres gets all of the genres from the database
func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.GenresAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, genres, "genres"); err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	movies, err := app.models.DB.All(genreID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, movies, "movies"); err != nil {
		app.errorJSON(w, err)
		return
	}
}

// deleteMovie deletes a movie from the database
func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

}

// updateMovie updates a movie in the database
func (app *application) editMovie(w http.ResponseWriter, r *http.Request) {
	type jsonResponse struct {
		OK bool `json:"ok"`
	}

	ok := jsonResponse{
		OK: true,
	}

	err := app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// searchMovies searches for a movie in the database
func (app *application) searchMovies(w http.ResponseWriter, r *http.Request) {

}