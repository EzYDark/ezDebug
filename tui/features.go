package tui

// Feature defines a toggleable feature with its state and action.
type Feature struct {
	Name        string
	Description string
	Enabled     bool
	Action      func(enabled bool)
}

type FeatureList []Feature

// featList is the global feature list.
var featList FeatureList

// GetFeatureList returns a pointer to the global feature list.
func GetFeatureList() *FeatureList {
	return &featList
}

// AddFeature adds a feature to the global feature list.
func (list *FeatureList) Add(feature Feature) *FeatureList {
	*list = append(*list, feature)
	return list
}

// SetFeatureList sets the global feature list to a new feature list.
func (list *FeatureList) Set(newList []Feature) *FeatureList {
	*list = newList
	return list
}
