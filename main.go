package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/a-collins/meal-planner/storage"
	"github.com/a-collins/meal-planner/types"
	log "github.com/sirupsen/logrus"
	// "github.com/aws/aws-sdk-go/aws"
)

func main() {
	log.SetLevel(log.InfoLevel)
	filePath := "meals.json"
	sd := storage.StorageDriver{
		Storer: storage.FileStorer{
			FilePath: filePath,
		},
	}

	initialMeals := []types.Meal{
		{
			Name:        "fajitas",
			Ingredients: []string{"wraps", "peppers", "onion", "spices", "cheese", "chicken"},
			Tags:        []string{"easy", "mexican", "weeknight"},
		},
		{
			Name: "some other meal",
		},
		{
			Name:        "chicken",
			Ingredients: []string{"just chicken"},
			Tags:        []string{"plain", "bland", "weekend"},
		},
	}

	mealObj, err := json.Marshal(initialMeals)
	if err != nil {
		panic(fmt.Errorf("failed to marshal initial meals - %v", err))
	}
	err = os.WriteFile(filePath, mealObj, 0777)
	if err != nil {
		panic(fmt.Errorf("failed writing meal initial meals - %v", err))
	}

	err = sd.AddMeal(&types.Meal{Name: "pasta"})
	if err != nil {
		log.Error(err)
	}
	err = sd.DeleteMeal("chicken")
	if err != nil {
		log.Error(err)
	}
}
