package storage

import (
	"fmt"

	types "github.com/a-collins/meal-planner/types"
	log "github.com/sirupsen/logrus"
)

type StorageDriver struct {
	Storer types.MealStorer
}

// AddMeal adds a meal to the data store
func (s StorageDriver) AddMeal(m *types.Meal) error {
	if m.Name == "" {
		return fmt.Errorf("input meal has empty name")
	}
	err := s.Storer.AddMeal(m)
	if err != nil {
		// TODO: Handle different error types here
		log.Errorf("Failed to add meal %s. Error: %s", m.Name, err)
		return err
	}
	log.Infof("Successfully added meal %s", m.Name)
	return nil
}

// GetMeal gets a meal from the data store
func (s StorageDriver) GetMeal(name string) (*types.Meal, error) {
	meal, err := s.Storer.GetMeal(name)
	if err != nil {
		// TODO: Handle different error types here
		log.Errorf("Failed to delete meal %s. Error: %s", name, err)
		return nil, err
	}
	log.Infof("Successfully retrieved meal %s", name)
	return meal, nil
}

// DeleteMeal deletes a meal from the data store by name
func (s StorageDriver) DeleteMeal(name string) error {
	err := s.Storer.DeleteMeal(name)
	if err != nil {
		// TODO: Handle different error types here
		log.Errorf("Failed to delete meal %s. Error: %s", name, err)
		return err
	}
	log.Infof("Successfully deleted meal %s", name)
	return nil
}
