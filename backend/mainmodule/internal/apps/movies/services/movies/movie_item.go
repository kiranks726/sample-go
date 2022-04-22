package movies

type Movie struct {
	Id               string  `json:"Id,omitempty"`
	MovieId          int     `json:"MovieId" csv:"id"`
	ImdbId           string  `json:"ImdbId" csv:"imdb_id"`
	Title            string  `json:"Title" csv:"title"`
	Tagline          string  `json:"Tagline" csv:"tagline"`
	Overview         string  `json:"Overview" csv:"overview"`
	PosterPath       string  `json:"PosterPath" csv:"poster_path"`
	Video            string  `json:"Video" csv:"video"`
	Runtime          float32 `json:"Runtime" csv:"runtime"`
	OriginalLanguage string  `json:"OriginalLanguage" csv:"original_language"`
	SpokenLanguage   string  `json:"SpokenLanguage" csv:"spoken_languages"`
	Status           string  `json:"Status" csv:"status"`
	ReleaseDate      string  `json:"ReleaseDate" csv:"release_date"`
}
