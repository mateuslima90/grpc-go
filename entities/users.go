package entities

type User struct {
	ObjectID string `bson:"_id" json:"_id"`

	Username string `json:"Username,omitempty"`

	Name string `json:"Name,omitempty"`

	Email string `json:"Email,omitempty"`
}

type UserDTO struct {
	Username string `json:"Username,omitempty"`

	Name string `json:"Name,omitempty"`

	Email string `json:"Email,omitempty"`
}