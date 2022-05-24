package repository

import (
	"database/sql"
	"golang.org/x/net/context"
	"myapp/internal/microservices/movie"
	"myapp/internal/microservices/movie/proto"
)

type movieStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *movieStorage {
	return &movieStorage{db: db}
}

func (ms *movieStorage) GetOne(id int) (*proto.Movie, error) {
	sqlScript := "SELECT id, name, is_movie, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, " +
		"picture, video, trailer FROM movies WHERE id=$1"

	var selectedMovie proto.Movie
	err := ms.db.QueryRow(sqlScript, id).Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.IsMovie,
		&selectedMovie.NamePicture, &selectedMovie.Year, &selectedMovie.Duration, &selectedMovie.AgeLimit,
		&selectedMovie.Description, &selectedMovie.KinopoiskRating, &selectedMovie.Tagline, &selectedMovie.Picture,
		&selectedMovie.Video, &selectedMovie.Trailer)
	if err != nil {
		return nil, err
	}

	return &selectedMovie, nil
}

func (ms *movieStorage) GetSeasonsAndEpisodes(seriesID int) ([]*proto.Season, error) {
	sqlScript := "SELECT id, number FROM seasons WHERE movie_id = $1 ORDER BY number;"

	selectedSeasons := make([]*proto.Season, 0)

	rows, err := ms.db.Query(sqlScript, seriesID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var singleSeason proto.Season
		if err = rows.Scan(&singleSeason.ID, &singleSeason.Number); err != nil {
			return nil, err
		}
		selectedSeasons = append(selectedSeasons, &singleSeason)
	}

	sqlScript = "SELECT e.id, e.name, e.number, e.description, e.video, e.photo, s.id, s.number FROM episode AS e JOIN seasons s on e.season_id = s.id WHERE s.movie_id = $1 ORDER BY s.number, e.number;"

	rows, err = ms.db.Query(sqlScript, seriesID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var singleEpisode proto.Episode
		var seasonId, seasonNumber int32
		if err = rows.Scan(&singleEpisode.ID, &singleEpisode.Name, &singleEpisode.Number, &singleEpisode.Description,
			&singleEpisode.Video, &singleEpisode.Picture, &seasonId, &seasonNumber); err != nil {
			return nil, err
		}
		selectedSeasons[seasonNumber-1].Episodes = append(selectedSeasons[seasonNumber-1].Episodes, &singleEpisode)
	}

	return selectedSeasons, nil
}

func (ms *movieStorage) GetAllMovies(limit, offset int) ([]*proto.Movie, error) {
	sqlScript := "SELECT id, name, is_movie, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, " +
		"picture, video, trailer FROM movies LIMIT $1 OFFSET $2"

	selectedMovies := make([]*proto.Movie, 0, limit)

	rows, err := ms.db.Query(sqlScript, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var singleMovie proto.Movie
		if err = rows.Scan(&singleMovie.ID, &singleMovie.Name, &singleMovie.IsMovie, &singleMovie.NamePicture, &singleMovie.Year,
			&singleMovie.Duration, &singleMovie.AgeLimit, &singleMovie.Description, &singleMovie.KinopoiskRating,
			&singleMovie.Tagline, &singleMovie.Picture, &singleMovie.Video, &singleMovie.Trailer); err != nil {
			return nil, err
		}
		selectedMovies = append(selectedMovies, &singleMovie)
	}

	return selectedMovies, nil
}

func (ms *movieStorage) GetRandomMovie() (*proto.MainMovie, error) {
	sqlScript := "SELECT id, name_picture, tagline, picture FROM movies ORDER BY RANDOM() LIMIT 1"

	var mainMovie proto.MainMovie
	err := ms.db.QueryRow(sqlScript).Scan(&(mainMovie.ID), &(mainMovie.NamePicture),
		&(mainMovie.Tagline), &(mainMovie.Picture))
	if err != nil {
		return nil, err
	}

	return &mainMovie, nil
}

func (ms *movieStorage) GetMovieRating(movieID int) (*movie.GetMovieRatingAnswer, error) {
	sqlScript := "SELECT rating_sum, rating_count FROM movies WHERE id = $1"

	returnValue := movie.GetMovieRatingAnswer{
		RatingSum:   0,
		RatingCount: 0,
	}
	err := ms.db.QueryRow(sqlScript, movieID).Scan(&returnValue.RatingSum, &returnValue.RatingCount)

	return &returnValue, err
}

func (ms *movieStorage) AddMovieRating(options *proto.AddRatingOptions) error {
	ctx := context.Background()
	transaction, err := ms.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = transaction.ExecContext(ctx, "INSERT INTO rating(movie_id, user_id, rating) VALUES ($1, $2, $3);", options.MovieID, options.UserID, options.Rating)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	_, err = transaction.ExecContext(ctx, "UPDATE movies SET rating_count=((SELECT m1.rating_count FROM movies AS m1 WHERE m1.id=$1)+1) WHERE id=$1;", options.MovieID)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	_, err = transaction.ExecContext(ctx, "UPDATE movies SET rating_sum=((SELECT m1.rating_sum FROM movies AS m1 WHERE m1.id=$1)+$2) WHERE id=$1;", options.MovieID, options.Rating)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}
	return err

	//_, err := ms.db.Exec(sqlScript, options.MovieID, options.UserID, options.Rating)
	//return err
}

func (ms *movieStorage) ChangeMovieRating(options *proto.AddRatingOptions) error {
	ctx := context.Background()
	transaction, err := ms.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = transaction.ExecContext(ctx, "UPDATE movies SET rating_sum=( "+
		"(SELECT m1.rating_sum FROM movies AS m1 WHERE m1.id=$1) "+
		"+ $3 "+
		"- (SELECT rating FROM rating WHERE movie_id = $1 AND user_id=$2)"+
		") "+
		"WHERE id=$1; ", options.MovieID, options.UserID, options.Rating)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	_, err = transaction.ExecContext(ctx, "UPDATE rating SET rating=$3 WHERE movie_id = $1 AND user_id=$2;", options.MovieID, options.UserID, options.Rating)
	if err != nil {
		err := transaction.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}
	return err
}
func (ms *movieStorage) CheckRatingExists(options *proto.AddRatingOptions) (*movie.CheckRatingExistsAnswer, error) {
	sqlScript := "SELECT rating FROM rating WHERE user_id=$1 AND movie_id = $2"

	returnValue := movie.CheckRatingExistsAnswer{
		Exists: false,
		Rating: 0,
	}
	err := ms.db.QueryRow(sqlScript, options.UserID, options.MovieID).Scan(&returnValue.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			returnValue.Exists = false
			return &returnValue, nil
		}
		return nil, err
	}
	returnValue.Exists = true
	return &returnValue, nil
}
