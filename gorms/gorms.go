package gorms

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Generic helper wrappers for our differents models
type GormDB interface {
	Migrate(ctx context.Context, models any) (bool, error)
	Create(ctx context.Context, models any) (any, error)
	GetRows(ctx context.Context, models any) (any, error)
	Get(ctx context.Context, models any, fields map[string]any) (any, error)
	GetAll(ctx context.Context, models any) (any, error)
	First(ctx context.Context, models any, id string) (any, error)
	FindAll(ctx context.Context, models any, query string) (any, error)
	Update(ctx context.Context, models any, id string, fields map[string]any) (bool, error)
	Updates(ctx context.Context, model, updaded any) error
	Delete(ctx context.Context, models any, id string) (bool, error)
}

type gormDB struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) GormDB {
	return &gormDB{
		db: db,
	}
}

func (g *gormDB) Migrate(ctx context.Context, models any) (bool, error) {
	if err := g.db.AutoMigrate(models); err != nil {
		return false, err
	}
	return true, nil
}

// Create data from any given models
func (g *gormDB) Create(ctx context.Context, models any) (any, error) {
	if err := g.db.Model(models).Error; err != nil {
		return "", err
	}

	tx := g.db.Create(models)
	if tx.Error != nil {
		return "", tx.Error
	}
	return models, nil
}

func (g *gormDB) GetRows(ctx context.Context, models any) (any, error) {
	rows, err := g.db.Model(models).Preload(clause.Associations).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		g.db.ScanRows(rows, models)
	}
	return models, nil
}

// Get data from different query fields
func (g *gormDB) Get(
	ctx context.Context, models any, fields map[string]any) (any, error) {
	if err := g.db.Where(fields).Preload(clause.Associations).Find(models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (g *gormDB) First(ctx context.Context, models any, id string) (any, error) {
	if err := g.db.Preload(clause.Associations).First(models, id).Error; err != nil {
		return nil, err
	}
	return models, nil
}

// Get all model with association (nested model)
func (g *gormDB) GetAll(ctx context.Context, models any) (any, error) {
	if err := g.db.Preload(clause.Associations).Find(models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

// find any given model with a query
func (g *gormDB) FindAll(ctx context.Context, models any, query string) (any, error) {
	if err := g.db.Where(query).Find(models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

// update any given models with their column and values that need to be change
func (g *gormDB) Update(
	ctx context.Context, models any, id string, fields map[string]any) (bool, error) {
	for index, value := range fields {
		if err := g.db.Debug().Model(models).Where("id = ?", id).Update(index, value).Error; err != nil {
			return false, err
		}
	}
	return true, nil
}

func (g *gormDB) Updates(ctx context.Context, model, updaded any) error {
	if err := g.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Debug().Model(model).Updates(updaded); err != nil {
		return err.Error
	}
	return nil
}

// delete any given data from models with id
func (g *gormDB) Delete(ctx context.Context, models any, id string) (bool, error) {
	if err := g.db.Debug().Delete(models, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
