package requests_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rjhoppe/firekeeper/requests"
)

type RandomRecipesResponse struct {
	RecipeOne   string `json:"recipeOne"`
	RecipeTwo   string `json:"recipeTwo"`
	RecipeThree string `json:"recipeThree"`
}

type RandomDrinkResponse struct {
	Message      string `json:"message"`
	ExternalId   string `json:"idDrink"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	Glass        string `json:"glass"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
}

func TestGETDrinkRequest(t *testing.T) {
	// Prepare expected data
	expected := RandomDrinkResponse{
		Message:      "Drink of the Day",
		ExternalId:   "1001",
		Name:         "Cranberry Cordial",
		Category:     "Homemade Liqueur",
		Glass:        "Collins",
		Ingredients:  "1/2 kg chopped  Cranberries, 3/4 L  Sugar, 1/2 L  Light rum, ",
		Instructions: "Place the chopped cranberries in a 2 liter jar that has a tight-fitting lid. Add the sugar and rum. Adjust the lid securely and place the jar in a cool, dark place. Invert the jar and shake it every day for six weeks. Strain the cordial into bottles and seal with corks.",
	}

	// Start a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(expected)
	}))

	defer server.Close()

	// Use the test server's URL
	req := requests.GETRequest{Url: server.URL}
	resp, _ := req.Send()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var actual RandomDrinkResponse
	err := json.Unmarshal(resp.Body, &actual)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetDinnerRequest(t *testing.T) {
	// Prepare expected data
	expected := RandomRecipesResponse{
		RecipeOne:   "1: Spaghetti",
		RecipeTwo:   "2: Tacos",
		RecipeThree: "3: Curry",
	}

	// Start a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(expected)
	}))

	defer server.Close()

	// Use the test server's URL
	req := requests.GETRequest{Url: server.URL}
	resp, _ := req.Send()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var actual RandomRecipesResponse
	err := json.Unmarshal(resp.Body, &actual)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
