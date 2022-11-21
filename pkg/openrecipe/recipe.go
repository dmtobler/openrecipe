package openrecipe

type Recipe struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Servings int    `json:"servings"`
}
