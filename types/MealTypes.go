package types

type MealStorer interface {
	AddMeal(m *Meal) error
	GetMeal(name string) (*Meal, error)
	DeleteMeal(name string) error
}

type Meal struct {
	Name        string
	Tags        []string
	Ingredients []string
}
