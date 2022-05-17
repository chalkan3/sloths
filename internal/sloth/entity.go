package sloth

type Sloth struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Family string `json:"family,omitempty"`
}

func NewSloth(id int, name, family string) *Sloth {
	return &Sloth{
		ID:     id,
		Name:   name,
		Family: family,
	}
}

func (sloth *Sloth) SetID(id int) *Sloth {
	sloth.ID = id
	return sloth
}

func (sloth *Sloth) GetID() int {
	return sloth.ID
}
