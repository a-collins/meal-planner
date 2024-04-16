package storage

import (
	"encoding/json"
	"fmt"
	"os"

	types "github.com/a-collins/meal-planner/types"
	log "github.com/sirupsen/logrus"
)

type FileStorer struct {
	FilePath string
}

// AddMeal adds a meal to the data store
func (fs FileStorer) AddMeal(m *types.Meal) error {
	log.Debugf("Adding meal %s", m.Name)

	meals, err := loadMeals(fs.FilePath)
	if err != nil {
		log.Error(fmt.Errorf("failed to add meal '%v' - %v", m.Name, err))
		return err
	}

	if mealExists(m.Name, &meals) {
		return fmt.Errorf("meal '%v' already exists in the database", m.Name)
	}
	log.Debugf("Meals before: %v", meals)
	// Add meal to meal list
	log.Info(fmt.Sprintf("Adding meal '%v'", m.Name))
	meals = append(meals, *m)
	log.Debugf("Meals after: %v", meals)

	// Write json back into file
	err = writeMeals(fs.FilePath, &meals)
	if err != nil {
		return fmt.Errorf("failed writing meal '%v' - %v", m.Name, err)
	}
	return nil
}

// GetMeal gets a meal from the data store. Returns nil meal and nil error if meal was not found.
func (fs FileStorer) GetMeal(name string) (*types.Meal, error) {
	meals, err := loadMeals(fs.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to GetMeal '%v' - %v", name, err)
	}
	for _, m := range meals {
		if m.Name == name {
			return &m, nil
		}
	}
	return nil, nil
}

// DeleteMeal deletes a meal from the data store by name
func (fs FileStorer) DeleteMeal(name string) error {
	meals, err := loadMeals(fs.FilePath)
	if err != nil {
		return fmt.Errorf("failed to DeleteMeal '%v' - %v", name, err)
	}
	index := -1
	for i, m := range meals {
		if m.Name == name {
			index = i
		}
	}
	if index != -1 {
		newMeals := append(meals[:index], meals[(index+1):]...)
		err = writeMeals(fs.FilePath, &newMeals)
		if err != nil {
			return fmt.Errorf("failed deleting meal '%v' - %v", name, err)
		}
	} else {
		log.Infof("Meal %v not found, nothing deleted", name)
	}
	return nil
}

// loadMeals loads meals from the given filePath into the provided meals object
// TODO: should this just return a meals object rather than take one in to load stuff into?
func loadMeals(filePath string) ([]types.Meal, error) {
	meals := []types.Meal{}
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		log.Error(fmt.Errorf("failed to load meals object from '%v'. Err=%v", filePath, err))
		return nil, err
	}
	err = json.Unmarshal(fileContents, &meals)
	if err != nil {
		log.Error(fmt.Errorf("failed to unmarshal meals object from file '%v'. Err=%v", filePath, err))
		return nil, err
	}
	log.Debugf("Loaded meals from %s", filePath)

	return meals, nil
}

func writeMeals(filePath string, meals *[]types.Meal) error {
	payload, err := json.Marshal(meals)
	if err != nil {
		return fmt.Errorf("failed to marshal meals into json object - %v", err)
	}
	err = os.WriteFile(filePath, payload, 0777)
	if err != nil {
		return fmt.Errorf("failed to write meals to file - %v", err)
	}
	return nil
}

func mealExists(mealName string, meals *[]types.Meal) bool {
	for _, meal := range *meals {
		if meal.Name == mealName {
			return true
		}
	}
	return false
}
