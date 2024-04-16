package entities

type Product struct {
	Identifier int `json:"identifier"`
	FullName string `json:"full_name"`
	StateName string `json:"state_name"`
}
