package handlers

import (
	"mentalarts_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMusician godoc
// @Summary Create a new musician
// @Description Adds a new musician to the database.
// @Tags Musician
// @Accept json
// @Produce json
// @Param musician body models.Musician true "Musician Info"
// @Success 201 {object} models.Musician
// @Failure 400 {object} string "Error message"
// @Router /musicians [post]
func (h Handler) CreateMusician(c *gin.Context) {
	var musician models.Musician

	err := c.BindJSON(&musician)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.Create(&musician).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, musician)
}

// GetMusicians godoc
// @Summary Retrieve all musicians
// @Description Returns a list of all musicians.
// @Tags Musician
// @Produce json
// @Success 200 {array} models.Musician
// @Failure 400 {object} string "Error message"
// @Router /musicians [get]
func (h Handler) GetMusicians(c *gin.Context) {
	var musicians []models.Musician

	if err := h.db.Find(&musicians).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, musicians)
}

// UpdateMusician godoc
// @Summary Update a musician
// @Description Updates an existing musician's information.
// @Tags Musician
// @Accept json
// @Produce json
// @Param id path int true "Musician ID"
// @Param musician body models.Musician true "Updated Musician Info"
// @Success 200 {object} models.Musician
// @Failure 400 {object} string "Error message"
// @Router /musicians/{id} [put]
func (h Handler) UpdateMusician(c *gin.Context) {
	var musician models.Musician
	id := c.Params.ByName("id")

	if err := h.db.Where("id = ?", id).First(&musician).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := c.BindJSON(&musician)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.Save(&musician).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, musician)
}

// DeleteMusicianById godoc
// @Summary Delete a musician by ID
// @Description Deletes the musician with the specified ID.
// @Tags Musician
// @Produce json
// @Param id path int true "Musician ID"
// @Success 200 {string} string "Musician deleted"
// @Failure 400 {object} string "Error message"
// @Router /musicians/{id} [delete]
func (h Handler) DeleteMusicianById(c *gin.Context) {
	var musician models.Musician
	id := c.Params.ByName("id")

	if err := h.db.Where("id = ?", id).First(&musician).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.Delete(&musician).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Musician deleted")
}

// DeleteAllMusicians godoc
// @Summary Delete all musicians
// @Description Deletes all musicians from the database.
// @Tags Musician
// @Produce json
// @Success 200 {string} string "All musicians deleted"
// @Failure 400 {object} string "Error message"
// @Router /musicians [delete]
func (h Handler) DeleteAllMusicians(c *gin.Context) {
	if err := h.db.Delete(models.Musician{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "All musicians deleted")
}

// GetMusicianById godoc
// @Summary Retrieve a musician by ID
// @Description Returns a single musician by ID.
// @Tags Musician
// @Produce json
// @Param id path int true "Musician ID"
// @Success 200 {object} models.Musician
// @Failure 400 {object} string "Error message"
// @Router /musicians/{id} [get]
func (h Handler) GetMusicianById(c *gin.Context) {
	id := c.Params.ByName("id")
	var musician models.Musician
	if err := h.db.Where("id = ?", id).First(&musician).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, musician)
}
