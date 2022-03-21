package moviesCompilations

type Movie struct {
	ID      int      `json:"id" form:"id"`
	Name    string   `json:"name" form:"name"`
	Genre   []string `json:"genre" form:"genre"`
	Picture string   `json:"picture" form:"picture"`
}

type MovieCompilation struct {
	Name   string  `json:"compilation_name"`
	Movies []Movie `json:"movies"`
}
