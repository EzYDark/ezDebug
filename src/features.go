package src

// Feature defines a toggleable feature with its state and action.
type Feature struct {
	Name        string
	Description string
	Enabled     bool
	Action      func(enabled bool)
}

type FeatureList []Feature

var featList FeatureList

func GetFeatureList() *FeatureList {
	return &featList
}

func (list *FeatureList) Add(feature Feature) *FeatureList {
	*list = append(*list, feature)
	return list
}

func (list *FeatureList) Set(newList []Feature) *FeatureList {
	*list = newList
	return list
}
