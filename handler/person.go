package handler

import (
	"fga/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// CreatePerson AddAccount godoc
// @Summary      Add an person
// @Tags         person
// @Produce      json
// @Param        first_name formData   string false "first person name"
// @Param        last_name  formData   string false "last person name"
// @Success      200      {object}  model.Person
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /person [post]
func CreatePerson(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var person model.Person

		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")

		person.FirstName = firstName
		person.LastName = lastName

		db.Create(&person)

		result := gin.H{
			"result": person,
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetPerson(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var person model.Person
		id := c.Param("id")

		err := db.Where("id = ?", id).First(&person).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err.Error(),
				"count":  0,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": person,
		})
	}
}

func GetPersons(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		persons := make([]*model.Person, 0)

		db.Find(&persons)
		if len(persons) <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"result": persons,
				"count":  0,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	}

}

func UpdatePerson(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")

		var existedPerson model.Person
		err := db.Where("id = ?", id).First(&existedPerson).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": "data not found",
			})

			return
		}

		var newPerson model.Person
		newPerson.FirstName = firstName
		newPerson.LastName = lastName

		err = db.Model(&existedPerson).Updates(newPerson).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": "update failed",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "successfuly update data",
		})
	}
}

func DeletePerson(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var person model.Person

		err := db.Where("id = ?", id).First(&person).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": "data not found",
			})

			return
		}

		err = db.Delete(&person).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": "delete failed",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "data deleted successfully",
		})
	}
}
