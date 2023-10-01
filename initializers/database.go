package initializers

import (
	"encoding/json"
	"fmt"
	"go-elden-ring-api/config"
	"go-elden-ring-api/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnect() {
	conf := config.New()
	dsn := ConnectionString(conf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Default().Print("Database Connected")
	}
	Migrate(db)
	Seed(db)
	DB = db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Material{})
	db.AutoMigrate(&models.Recipe{})
	db.AutoMigrate(&models.CookBook{})
	db.AutoMigrate(&models.RecipeMaterial{})
}

func Seed(db *gorm.DB) {
	if db.Migrator().HasTable(&models.Material{}) {
		SeedMaterials(db)
	}
	if db.Migrator().HasTable(&models.Recipe{}) {
		SeedRecipes(db)
	}
	if db.Migrator().HasTable(&models.CookBook{}) {
		SeedCookBooks(db)
	}

}

type MaterialJson struct {
	Name        string
	Description string
}

func SeedMaterials(db *gorm.DB) {
	err := db.First(&models.Material{}).Error
	if err != nil {
		log.Default().Print("Seeding Materials")
		var data []MaterialJson
		bytes, _ := os.ReadFile("data/materials.json")
		err := json.Unmarshal(bytes, &data)
		if err == nil {
			for _, item := range data {
				material := models.Material{Name: item.Name, Description: item.Description}
				result := db.Create(&material)
				if result.Error != nil {
					log.Default().Print("Error saving: " + item.Name)
				}
			}
		}
	}

}

type RecipeMaterialJson struct {
	Quantity int16
	Name     string
}

type RecipeJson struct {
	Name        string
	Description string
	Materials   []RecipeMaterialJson
}

func SeedRecipes(db *gorm.DB) {
	err := db.First(&models.Recipe{}).Error
	if err != nil {
		log.Default().Print("Seeding Recipes")
		var data []RecipeJson
		bytes, _ := os.ReadFile("data/recipes.json")
		err := json.Unmarshal(bytes, &data)
		if err == nil {
			for _, item := range data {
				recipe := models.Recipe{Name: item.Name, Description: item.Description}
				db.Create(&recipe)
				for _, mat := range item.Materials {
					material := models.Material{}
					db.Model(&models.Material{}).Where("Name = ?", mat.Name).First(&material)
					recipeMaterial := models.RecipeMaterial{RecipeID: recipe.ID, MaterialID: material.ID, Quantity: uint16(mat.Quantity)}
					db.Create(&recipeMaterial)
				}
			}
		}
	}
}

type CookBookJson struct {
	Name        string
	Description string
	Recipes     []string
}

func SeedCookBooks(db *gorm.DB) {
	err := db.First(&models.CookBook{}).Error
	if err != nil {
		log.Default().Print("Seeding Cook Books")
		var data []CookBookJson
		bytes, _ := os.ReadFile("data/cook_books.json")
		err := json.Unmarshal(bytes, &data)
		if err == nil {
			for _, item := range data {
				cookBook := models.CookBook{Name: item.Name, Description: item.Description}
				db.Create(&cookBook)
				for _, name := range item.Recipes {
					recipe := models.Recipe{}
					db.Model(&models.Recipe{}).Where("Name = ?", name).First(&recipe)
					cookBookRecipe := models.CookBookRecipe{CookBookID: cookBook.ID, RecipeID: recipe.ID}
					db.Create(&cookBookRecipe)
				}
			}
		}
	}
}

func ConnectionString(d config.DbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", d.DbUser, d.DbPass, d.DbHost, d.DbPort, d.DbName)
}
