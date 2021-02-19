package types

// Cause struct is used to store the case info of each user
type Cause struct {
	Name        string   `json:"name" bson:"name" validate:"required"`
	Description string   `json:"description" bson:"description"`
	To          []string `json:"to" bson:"to" validate:"required"`
	Subjects    []string `json:"subjects" bson:"subjects" validate:"required"`
	Bodies      []string `json:"bodies" bson:"bodies" validate:"required"`
}
