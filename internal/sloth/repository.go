package sloth

import (
	postgresGORM "chalkan.github.com/internal/database"
)

var DB = map[int]*Sloth{1: NewSloth(1, "first", "none")}

type Repository interface {
	Add(sloth *Sloth) *Sloth
	Update(sloth *Sloth) *Sloth
	Delete(sloth *Sloth)
	Get(sloth *Sloth) *Sloth
	List() []*Sloth
}

type repository struct {
	gorm *postgresGORM.PostgresGORM
}

func NewRepository(gorm *postgresGORM.PostgresGORM) Repository {
	return &repository{
		gorm: gorm,
	}
}

func (r *repository) Add(sloth *Sloth) *Sloth {
	r.gorm.GetGORM().Create(NewSlothModel(sloth.Name, sloth.Family))
	return sloth
}

func (r *repository) Delete(sloth *Sloth) {
}

func (r *repository) Update(sloth *Sloth) *Sloth {
	return sloth
}

func (r *repository) Get(sloth *Sloth) *Sloth {

	return DB[sloth.GetID()]
}

func (r *repository) List() []*Sloth {

	return make([]*Sloth, 1)
}
