package views

import (
	"Recipes_StdLib/recipes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	RecipeRe       = regexp.MustCompile(`^/recipes/*$`)
	RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+([?:-[a-z0-9]+)+)$`)
)

type RecipesHandler struct {
	Store recipes.RecipeStore
}

type HomeHandler struct{}

func NewRecipeHandler(s recipes.RecipeStore) *RecipesHandler {
	return &RecipesHandler{
		Store: s,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

func generateRandomSlug(input string) string {
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)
	base64String := base64.URLEncoding.EncodeToString(randomBytes)
	slug := strings.ReplaceAll(base64String, "=", "")
	return slug
}

func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
		h.CreateRecipe(w, r)
		return
	case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
		h.ListRecipes(w, r)
		return
	case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
		h.GetRecipe(w, r)
		return
	case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
		h.UpdateRecipe(w, r)
		return
	case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
		h.DeleteRecipe(w, r)
		return
	default:
		return
	}
}

func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside create recipes function handler.")

	var recipe recipes.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		log.Fatal(err)
		return
	}
	resourceID := generateRandomSlug(recipe.Name)
	if err := h.Store.Add(resourceID, recipe); err != nil {
		InternalServerErrorHandler(w, r)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h *RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside list recipes function handler.")
	allRecipes, err := h.Store.List()
	jsonBytes, err := json.Marshal(allRecipes)

	if err != nil {
		InternalServerErrorHandler(w, r)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)

	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	recipe, err := h.Store.Get(matches[1])
	if err != nil {
		if errors.Is(err, recipes.NotFoundError) {
			NotFoundErrorHandler(w, r)
			return
		}

		InternalServerErrorHandler(w, r)
		log.Fatal(err)
		return
	}

	jsonBytes, err := json.Marshal(recipe)
	if err != nil {
		InternalServerErrorHandler(w, r)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	if err != nil {
		InternalServerErrorHandler(w, r)
		log.Fatal(err)
		return
	}
}

func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	var recipe recipes.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		log.Fatal(err)
		return
	}

	if err := h.Store.Update(matches[1], recipe); err != nil {
		if errors.Is(err, recipes.NotFoundError) {
			NotFoundErrorHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	matches := RecipeReWithID.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.Store.Remove(matches[1]); err != nil {
		InternalServerErrorHandler(w, r)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
