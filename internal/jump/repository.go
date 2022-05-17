package jump

import (
	"time"

	postgresGORM "chalkan.github.com/internal/database"
)

type Repository interface {
	Insert(jumpEvent JumpEvent)
	Delete()
	Update()
	Get()
	List()
}

type repository struct {
	gorm *postgresGORM.PostgresGORM
}

func NewRepository(gorm *postgresGORM.PostgresGORM) Repository {
	return &repository{
		gorm: gorm,
	}
}

func (repo *repository) Insert(event JumpEvent) {
	repo.gorm.GetGORM().Create(NewJump(event.Name, true, time.Now()))
}
func (repo *repository) Delete() {

}
func (repo *repository) Update() {

}
func (repo *repository) Get() {

}
func (repo *repository) List() {

}
