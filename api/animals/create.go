package animalsapi

import (
	"net/http"
	"strconv"
	"time"

	"livestock-saas/server/internal/animals"
	"livestock-saas/server/internal/database"

	"github.com/gin-gonic/gin"
)

type createAnimalRequest struct {
	EarTag      string      `json:"ear_tag"`
	ElectronicID string     `json:"electronic_id"`
	Name        string      `json:"name"`
	Breed       string      `json:"breed"`
	Sex         string      `json:"sex"`
	Status      string      `json:"status"`
	BirthDate   string      `json:"birth_date"`
	BirthWeight interface{} `json:"birth_weight"`
}

func CreateAnimal(c *gin.Context) {
	var req createAnimalRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// basic validation
	errs := make(map[string]string)
	if req.EarTag == "" {
		errs["ear_tag"] = "ear_tag is required"
	}
	if req.Name == "" {
		errs["name"] = "name is required"
	}
	if req.Breed == "" {
		errs["breed"] = "breed is required"
	}
	if req.Sex == "" {
		errs["sex"] = "sex is required"
	}
	if req.Status == "" {
		errs["status"] = "status is required"
	}
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	// map to model
	animal := animals.Animal{
		EarTag:       req.EarTag,
		ElectronicID: req.ElectronicID,
		Name:         req.Name,
		Breed:        req.Breed,
		Sex:          req.Sex,
		Status:       req.Status,
	}

	// parse birth date if provided (expecting YYYY-MM-DD)
	if req.BirthDate != "" {
		if t, err := time.Parse("2006-01-02", req.BirthDate); err == nil {
			animal.BirthDate = t
		}
	}

	// parse birth weight (can be number or string)
	if req.BirthWeight != nil {
		switch v := req.BirthWeight.(type) {
		case float64:
			animal.BirthWeight = v
		case int:
			animal.BirthWeight = float64(v)
		case string:
			if v != "" {
				if parsed, err := parseFloat(v); err == nil {
					animal.BirthWeight = parsed
				}
			}
		}
	}

	if err := database.DB.Create(&animal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, animal)
}

// parseFloat attempts to parse a string to float64
func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
