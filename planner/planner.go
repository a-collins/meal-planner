package planner

import (
	"github.com/a-collins/meal-planner/storage"
)

type planner struct {
	storage storage.StorageDriver
}

// TODO:
// - Type that describes desired tags -> days preferences
// - Function which produces plan
// 	- Given plan preferences, produces plan for the week
//  - No duplicates setting?
//  - Batch meals setting? e.g. one meal can be used for multiple nights
