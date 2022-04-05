#!/bin/bash

moq -out mock/movie_usecase_mock.go -pkg mock movie Service:MockMovieService
moq -out mock/movie_repository_mock.go -pkg mock movie Storage:MockMovieStorage

moq -out mock/genre_repository_mock.go -pkg mock genre Storage:MockGenreStorage
moq -out mock/country_repository_mock.go -pkg mock country Storage:MockCountryStorage
moq -out mock/persons_repository_mock.go -pkg mock persons Storage:MockPersonsStorage