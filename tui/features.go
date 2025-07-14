package tui

// Feature defines a toggleable feature with its state and action.
type Feature struct {
	Name           string
	Description    string
	StartOnStartup bool
	Enabled        bool

	// OnStart enables the feature and runs its logic.
	OnStart func(self *Feature)
	// OnStop disables the feature and runs its cleanup.
	OnStop func(self *Feature)
}

// Toggle flips the feature's state by calling OnStart or OnStop.
func (f *Feature) Toggle() {
	if f.Enabled {
		if f.OnStop != nil {
			f.OnStop(f)
		}
	} else {
		if f.OnStart != nil {
			f.OnStart(f)
		}
	}
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
