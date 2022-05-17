package sloth

import (
	postgresGORM "chalkan.github.com/internal/database"
)

type Migrations struct {
	gorm *postgresGORM.PostgresGORM
}

func NewMigrations(gorm *postgresGORM.PostgresGORM) *Migrations {
	return &Migrations{
		gorm: gorm,
	}
}

func (m *Migrations) Migrate() {
	if ok := m.gorm.GetGORM().Migrator().HasTable(SlothModel{}); !ok {
		m.gorm.GetGORM().Migrator().CreateTable(SlothModel{})
	}
}
