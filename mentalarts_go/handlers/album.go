package handlers

import (
	"mentalarts_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAlbums godoc
// @Summary Get all albums
// @Description Returns a list of all albums as JSON.
// @Tags Albums
// @Accept json
// @Produce json
// @Success 200 {array} models.Album
// @Router /albums [get]
func (h Handler) GetAlbums(c *gin.Context) {
	var albums []models.Album

	if err := h.db.Preload("Musician").Find(&albums).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumByID godoc
// @Summary Get an album by ID
// @Description Returns the album with the specified ID as JSON.
// @Tags Albums
// @Accept json
// @Produce json
// @Param id path int true "Album ID"
// @Success 200 {object} models.Album
// @Router /album/{id} [get]
func (h Handler) GetAlbumByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var album models.Album
	if err := h.db.Where("id = ?", id).First(&album).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

// CreateAlbum godoc
// @Summary Create a new album
// @Description Adds a new album to the database.
// @Tags Albums
// @Accept json
// @Produce json
// @Param album body models.Album true "New Album Info"
// @Success 201 {object} models.Album
// @Router /album [post]
func (h Handler) CreateAlbum(c *gin.Context) {
	var album models.Album

	err := c.BindJSON(&album)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.db.Create(&album).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, album)
}

// UpdateAlbum godoc
// @Summary Update an album
// @Description Updates the details of an existing album.
// @Tags Albums
// @Accept json
// @Produce json
// @Param id path int true "Album ID"
// @Param album body models.Album true "Updated Album Info"
// @Success 200 {object} models.Album
// @Router /album/{id} [put]
func (h Handler) UpdateAlbum(c *gin.Context) {
	id := c.Params.ByName("id")
	var album models.Album
	if err := h.db.Where("id = ?", id).First(&album).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	err := c.BindJSON(&album)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.db.Save(&album).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

// DeleteAlbumByID godoc
// @Summary Delete an album by ID
// @Description Deletes the album with the specified ID.
// @Tags Albums
// @Accept json
// @Produce json
// @Param id path int true "Album ID"
// @Success 200 {string} string "Album deleted"
// @Router /album/{id} [delete]
func (h Handler) DeleteAlbumByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var album models.Album
	if err := h.db.Where("id = ?", id).First(&album).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.db.Delete(&album).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, "Album deleted")
}

// DeleteAllAlbums godoc
// @Summary Delete all albums
// @Description Deletes all albums from the database.
// @Tags Albums
// @Accept json
// @Produce json
// @Success 200 {string} string "All albums deleted"
// @Router /albums [delete]
func (h Handler) DeleteAllAlbums(c *gin.Context) {
	if err := h.db.Delete(models.Album{}).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, "All albums deleted")
}
