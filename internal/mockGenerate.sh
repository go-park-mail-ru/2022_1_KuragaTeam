#!/bin/bash

moq -out mock/movie_usecase_mock.go -pkg mock movie Service:MockMovieService