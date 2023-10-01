package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name        string           `gorm:"size:100;index:recipe_name;"`
	Description string           `gorm:"size:500"`
	Materials   []RecipeMaterial `gorm:"foreignKey:RecipeID"`
	CookBooks   []CookBook       `gorm:"many2many:cook_book_recipes"`
}

type RecipeMaterial struct {
	RecipeID   uint `gorm:"primaryKey"`
	MaterialID uint `gorm:"primaryKey"`
	Quantity   uint16
	Material   Material `gorm:"foreignKey:MaterialID"`
}
