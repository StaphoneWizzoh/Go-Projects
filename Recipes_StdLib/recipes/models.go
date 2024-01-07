package recipes

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name string `json:"name"`
}

type RecipeStore interface {
	Add(name string, recipe Recipe) error
	Get(name string) (Recipe, error)
	Update(name string, recipe Recipe) error
	List() (map[string]Recipe, error)
	Remove(name string) error
}
