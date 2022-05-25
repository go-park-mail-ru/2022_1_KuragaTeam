package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"myapp/internal"
	genre "myapp/internal/genre/repository"
	"myapp/internal/microservices/compilations/proto"
	"myapp/internal/microservices/compilations/repository"
	"myapp/internal/microservices/compilations/utils/images"
	persons "myapp/internal/persons/repository"
	"testing"
)

func TestMovieCompilationsUsecase_GetByPerson(t *testing.T) {
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie1 := proto.MovieInfo{
		ID:   1,
		Name: "Movie1",
		Genre: []*proto.Genre{
			{
				ID:   1,
				Name: "Боевик",
			},
			{
				ID:   2,
				Name: "Триллер",
			},
		},
		Picture: "picture_name.webp",
	}
	movie1FromStorage := proto.MovieInfo{
		ID:      movie1.ID,
		Name:    movie1.Name,
		Genre:   nil,
		Picture: movie1.Picture,
	}
	movie1.Picture, _ = images.GenerateFileURL(movie1.Picture, "posters")

	movie2 := proto.MovieInfo{
		ID:   2,
		Name: "Movie2",
		Genre: []*proto.Genre{
			{
				ID:   1,
				Name: "Фантастика",
			},
			{
				ID:   2,
				Name: "Семейный",
			},
		},
		Picture: "picture_name2.webp",
	}
	movie2FromStorage := proto.MovieInfo{
		ID:      movie2.ID,
		Name:    movie2.Name,
		Genre:   nil,
		Picture: movie2.Picture,
	}
	movie2.Picture, _ = images.GenerateFileURL(movie2.Picture, "posters")

	MC1Movie1 := movie1FromStorage
	MC1Movie2 := movie2FromStorage
	MC1 := proto.MovieCompilation{
		Name: "Test MC1",
		Movies: []*proto.MovieInfo{
			&MC1Movie1,
			&MC1Movie2,
		},
	}

	tests := []struct {
		name               string
		MCstorageMock      *repository.MockMovieCompilationStorage
		genreStorageMock   *genre.MockGenreStorage
		personsStorageMock *persons.MockPersonsStorage
		expected           *proto.MovieCompilation
		expectedError      bool
	}{
		{
			name: "Get compilation by person",
			MCstorageMock: &repository.MockMovieCompilationStorage{
				GetByPersonFunc: func(id int) (*proto.MovieCompilation, error) {
					return &MC1, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					switch id {
					case 1:
						return []internal.Genre{
							{
								ID:   1,
								Name: "Боевик",
							},
							{
								ID:   2,
								Name: "Триллер",
							},
						}, nil
					case 2:
						return []internal.Genre{
							{
								ID:   1,
								Name: "Фантастика",
							},
							{
								ID:   2,
								Name: "Семейный",
							},
						}, nil
					case 3:
						return []internal.Genre{
							{
								ID:   1,
								Name: "Детектив",
							},
							{
								ID:   2,
								Name: "Криминал",
							},
						}, nil
					}
					return nil, nil
				},
			},
			expected: &proto.MovieCompilation{
				Name:   MC1.Name,
				Movies: []*proto.MovieInfo{&movie1, &movie2},
			},
			expectedError: false,
		},
		{
			name: "MC storage returns error",
			MCstorageMock: &repository.MockMovieCompilationStorage{
				GetByPersonFunc: func(id int) (*proto.MovieCompilation, error) {
					return nil, errors.New(testError)
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					switch id {
					case 1:
						return []internal.Genre{
							{
								ID:   1,
								Name: "Боевик",
							},
							{
								ID:   2,
								Name: "Триллер",
							},
						}, nil
					case 2:
						return []internal.Genre{
							{
								ID:   1,
								Name: "Фантастика",
							},
							{
								ID:   2,
								Name: "Семейный",
							},
						}, nil
					case 3:
						return []internal.Genre{
							{
								ID:   1,
								Name: "Детектив",
							},
							{
								ID:   2,
								Name: "Криминал",
							},
						}, nil
					}
					return nil, nil
				},
			},
			expectedError: true,
		},
		{
			name: "MC storage returns error",
			MCstorageMock: &repository.MockMovieCompilationStorage{
				GetByPersonFunc: func(id int) (*proto.MovieCompilation, error) {
					return &MC1, nil
				},
			},
			genreStorageMock: &genre.MockGenreStorage{
				GetByMovieIDFunc: func(id int) ([]internal.Genre, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r := NewService(test.MCstorageMock, test.genreStorageMock, test.personsStorageMock)
			mainMC, err := r.GetByPerson(context.Background(), &proto.GetByIDOptions{ID: 1})

			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				for i := 0; i < len(test.expected.Movies); i++ {
					assert.Equal(t, test.expected.Movies[i].ID, mainMC.Movies[i].ID)
					assert.Equal(t, test.expected.Movies[i].Name, mainMC.Movies[i].Name)
					assert.Equal(t, test.expected.Movies[i].Picture, mainMC.Movies[i].Picture)
					for j := 0; j < len(test.expected.Movies[i].Genre); j++ {
						assert.Equal(t, test.expected.Movies[i].Genre[j].ID, mainMC.Movies[i].Genre[j].ID)
						assert.Equal(t, test.expected.Movies[i].Genre[j].Name, mainMC.Movies[i].Genre[j].Name)
					}
				}
			}
		})
	}
}
