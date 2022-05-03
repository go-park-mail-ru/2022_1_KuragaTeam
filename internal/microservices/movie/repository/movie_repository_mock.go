// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"myapp/internal/microservices/movie"
	"myapp/internal/microservices/movie/proto"
	"sync"
)

// Ensure, that MockMovieStorage does implement movie.Storage.
// If this is not the case, regenerate this file with moq.
var _ movie.Storage = &MockMovieStorage{}

// MockMovieStorage is a mock implementation of movie.Storage.
//
// 	func TestSomethingThatUsesStorage(t *testing.T) {
//
// 		// make and configure a mocked movie.Storage
// 		mockedStorage := &MockMovieStorage{
// 			GetAllMoviesFunc: func(limit int, offset int) ([]*proto.Movie, error) {
// 				panic("mock out the GetAllMovies method")
// 			},
// 			GetOneFunc: func(id int) (*proto.Movie, error) {
// 				panic("mock out the GetOne method")
// 			},
// 			GetRandomMovieFunc: func() (*proto.MainMovie, error) {
// 				panic("mock out the GetRandomMovie method")
// 			},
// 			GetSeasonsAndEpisodesFunc: func(seriesId int) ([]*proto.Season, error) {
// 				panic("mock out the GetSeasonsAndEpisodes method")
// 			},
// 		}
//
// 		// use mockedStorage in code that requires movie.Storage
// 		// and then make assertions.
//
// 	}
type MockMovieStorage struct {
	// GetAllMoviesFunc mocks the GetAllMovies method.
	GetAllMoviesFunc func(limit int, offset int) ([]*proto.Movie, error)

	// GetOneFunc mocks the GetOne method.
	GetOneFunc func(id int) (*proto.Movie, error)

	// GetRandomMovieFunc mocks the GetRandomMovie method.
	GetRandomMovieFunc func() (*proto.MainMovie, error)

	// GetSeasonsAndEpisodesFunc mocks the GetSeasonsAndEpisodes method.
	GetSeasonsAndEpisodesFunc func(seriesId int) ([]*proto.Season, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetAllMovies holds details about calls to the GetAllMovies method.
		GetAllMovies []struct {
			// Limit is the limit argument value.
			Limit int
			// Offset is the offset argument value.
			Offset int
		}
		// GetOne holds details about calls to the GetOne method.
		GetOne []struct {
			// ID is the id argument value.
			ID int
		}
		// GetRandomMovie holds details about calls to the GetRandomMovie method.
		GetRandomMovie []struct {
		}
		// GetSeasonsAndEpisodes holds details about calls to the GetSeasonsAndEpisodes method.
		GetSeasonsAndEpisodes []struct {
			// SeriesId is the seriesId argument value.
			SeriesId int
		}
	}
	lockGetAllMovies          sync.RWMutex
	lockGetOne                sync.RWMutex
	lockGetRandomMovie        sync.RWMutex
	lockGetSeasonsAndEpisodes sync.RWMutex
}

// GetAllMovies calls GetAllMoviesFunc.
func (mock *MockMovieStorage) GetAllMovies(limit int, offset int) ([]*proto.Movie, error) {
	if mock.GetAllMoviesFunc == nil {
		panic("MockMovieStorage.GetAllMoviesFunc: method is nil but Storage.GetAllMovies was just called")
	}
	callInfo := struct {
		Limit  int
		Offset int
	}{
		Limit:  limit,
		Offset: offset,
	}
	mock.lockGetAllMovies.Lock()
	mock.calls.GetAllMovies = append(mock.calls.GetAllMovies, callInfo)
	mock.lockGetAllMovies.Unlock()
	return mock.GetAllMoviesFunc(limit, offset)
}

// GetAllMoviesCalls gets all the calls that were made to GetAllMovies.
// Check the length with:
//     len(mockedStorage.GetAllMoviesCalls())
func (mock *MockMovieStorage) GetAllMoviesCalls() []struct {
	Limit  int
	Offset int
} {
	var calls []struct {
		Limit  int
		Offset int
	}
	mock.lockGetAllMovies.RLock()
	calls = mock.calls.GetAllMovies
	mock.lockGetAllMovies.RUnlock()
	return calls
}

// GetOne calls GetOneFunc.
func (mock *MockMovieStorage) GetOne(id int) (*proto.Movie, error) {
	if mock.GetOneFunc == nil {
		panic("MockMovieStorage.GetOneFunc: method is nil but Storage.GetOne was just called")
	}
	callInfo := struct {
		ID int
	}{
		ID: id,
	}
	mock.lockGetOne.Lock()
	mock.calls.GetOne = append(mock.calls.GetOne, callInfo)
	mock.lockGetOne.Unlock()
	return mock.GetOneFunc(id)
}

// GetOneCalls gets all the calls that were made to GetOne.
// Check the length with:
//     len(mockedStorage.GetOneCalls())
func (mock *MockMovieStorage) GetOneCalls() []struct {
	ID int
} {
	var calls []struct {
		ID int
	}
	mock.lockGetOne.RLock()
	calls = mock.calls.GetOne
	mock.lockGetOne.RUnlock()
	return calls
}

// GetRandomMovie calls GetRandomMovieFunc.
func (mock *MockMovieStorage) GetRandomMovie() (*proto.MainMovie, error) {
	if mock.GetRandomMovieFunc == nil {
		panic("MockMovieStorage.GetRandomMovieFunc: method is nil but Storage.GetRandomMovie was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetRandomMovie.Lock()
	mock.calls.GetRandomMovie = append(mock.calls.GetRandomMovie, callInfo)
	mock.lockGetRandomMovie.Unlock()
	return mock.GetRandomMovieFunc()
}

// GetRandomMovieCalls gets all the calls that were made to GetRandomMovie.
// Check the length with:
//     len(mockedStorage.GetRandomMovieCalls())
func (mock *MockMovieStorage) GetRandomMovieCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetRandomMovie.RLock()
	calls = mock.calls.GetRandomMovie
	mock.lockGetRandomMovie.RUnlock()
	return calls
}

// GetSeasonsAndEpisodes calls GetSeasonsAndEpisodesFunc.
func (mock *MockMovieStorage) GetSeasonsAndEpisodes(seriesId int) ([]*proto.Season, error) {
	if mock.GetSeasonsAndEpisodesFunc == nil {
		panic("MockMovieStorage.GetSeasonsAndEpisodesFunc: method is nil but Storage.GetSeasonsAndEpisodes was just called")
	}
	callInfo := struct {
		SeriesId int
	}{
		SeriesId: seriesId,
	}
	mock.lockGetSeasonsAndEpisodes.Lock()
	mock.calls.GetSeasonsAndEpisodes = append(mock.calls.GetSeasonsAndEpisodes, callInfo)
	mock.lockGetSeasonsAndEpisodes.Unlock()
	return mock.GetSeasonsAndEpisodesFunc(seriesId)
}

// GetSeasonsAndEpisodesCalls gets all the calls that were made to GetSeasonsAndEpisodes.
// Check the length with:
//     len(mockedStorage.GetSeasonsAndEpisodesCalls())
func (mock *MockMovieStorage) GetSeasonsAndEpisodesCalls() []struct {
	SeriesId int
} {
	var calls []struct {
		SeriesId int
	}
	mock.lockGetSeasonsAndEpisodes.RLock()
	calls = mock.calls.GetSeasonsAndEpisodes
	mock.lockGetSeasonsAndEpisodes.RUnlock()
	return calls
}