package controllers

import (
	"censusV/database"
	"censusV/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCensusData(c *gin.Context) {
	var data []models.CensusData
	db := database.DB

	// Filtering
	query := db.Model(&data)
	if income := c.Query("income"); income != "" {
		query = query.Where("income LIKE ?", "%"+income+"%")
	}
	if education := c.Query("education"); education != "" {
		query = query.Where("LOWER(education) LIKE ?", "%"+strings.ToLower(education)+"%")
	}
	if maritalStatus := c.Query("marital_status"); maritalStatus != "" {
		query = query.Where("LOWER(marital_status) LIKE ?", "%"+strings.ToLower(maritalStatus)+"%")
	}
	if occupation := c.Query("occupation"); occupation != "" {
		query = query.Where("LOWER(occupation) LIKE ?", "%"+strings.ToLower(occupation)+"%")
	}
	if minAge := c.Query("minAge"); minAge != "" {
		minAgeInt, err := strconv.Atoi(minAge)
		if err == nil {
			query = query.Where("age >= ?", minAgeInt)
		}
	}
	if maxAge := c.Query("maxAge"); maxAge != "" {
		maxAgeInt, err := strconv.Atoi(maxAge)
		if err == nil {
			query = query.Where("age <= ?", maxAgeInt)
		}
	}
	if ageRange := c.Query("range"); ageRange != "" {
		ages := strings.Split(ageRange, "-")
		if len(ages) == 2 {
			minAge, err1 := strconv.Atoi(ages[0])
			maxAge, err2 := strconv.Atoi(ages[1])
			if err1 == nil && err2 == nil {
				query = query.Where("age >= ? AND age <= ?", minAge, maxAge)
			}
		} else if strings.HasSuffix(ageRange, "+") {
			minAge, err := strconv.Atoi(strings.TrimSuffix(ageRange, "+"))
			if err == nil {
				query = query.Where("age >= ?", minAge)
			}
		}
	}
	if ageRange := c.Query("age"); ageRange != "" {
		ages := strings.Split(ageRange, "-")
		if len(ages) == 2 {
			minAge, err1 := strconv.Atoi(ages[0])
			maxAge, err2 := strconv.Atoi(ages[1])
			if err1 == nil && err2 == nil {
				query = query.Where("age >= ? AND age <= ?", minAge, maxAge)
			}
		} else if strings.HasSuffix(ageRange, "+") {
			minAge, err := strconv.Atoi(strings.TrimSuffix(ageRange, "+"))
			if err == nil {
				query = query.Where("age >= ?", minAge)
			}
		}
	}

	// Sorting
	if order := c.Query("order_by"); order != "" {
		if directionParam := c.Query("direction"); directionParam != "" {
			direction := c.DefaultQuery("direction", directionParam)
			query = query.Order(order + " " + direction)
		}
	}

	// Pagination
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'limit' parameter"})
		return
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'offset' parameter"})
		return
	}

	var totalElements int64
	countQuery := query // Duplicate the query
	if err := countQuery.Count(&totalElements).Error; err != nil {
		fmt.Println("Error counting elements:", err)
	}
	// fmt.Println("total elements", totalElements)
	query = query.Limit(limit).Offset(offset)

	// Fetch results
	if err := query.Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching census data"})
		return
	}

	response := struct {
		Total int64    `json:"records"`
		Data  []models.CensusData `json:"paginated_data"`
	}{
		Total: totalElements,
		Data:  data,
	}

	c.JSON(http.StatusOK, response)
}

func GetSummary(c *gin.Context) {
	var summary []struct {
		Education string `json:"education"`
		Count     int    `json:"count"`
	}

	if err := database.DB.Model(&models.CensusData{}).
		Select("education, COUNT(*) as count").
		Group("education").
		Scan(&summary).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data summary"})
		return
	}

	c.JSON(http.StatusOK, summary)
}
