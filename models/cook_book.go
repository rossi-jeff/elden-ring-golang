package models

import "gorm.io/gorm"

type CookBook struct {
	gorm.Model
	Name        string   `gorm:"size:100;index:cook_book_name;"`
	Description string   `gorm:"size:500"`
	Recipes     []Recipe `gorm:"many2many:cook_book_recipes"`
}

type CookBookRecipe struct {
	CookBookID uint `gorm:"primaryKey"`
	RecipeID   uint `gorm:"primaryKey"`
}
