package zones

import (
	"context"
	"net/http"
	"time"

	"github.com/eskpil/cod/trout/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	zoneService "github.com/eskpil/cod/trout/internal/zones"
	recordRoutes "github.com/eskpil/cod/trout/zones/records"

	log "github.com/sirupsen/logrus"
)

func SubRoutes(group *gin.RouterGroup) {
	records := group.Group("/:zoneId/records/")
	records.GET("/", recordRoutes.GetAll())
	records.POST("/", recordRoutes.CreateOne())
}

func GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		zones, err := zoneService.GetAll(ctx)
		if err != nil {
			log.Error(err)
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, zones)
	}
}

func CreateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var body struct {
			Name string `json:"name"`
			Fqdn string `json:"fqdn"`
		}

		if err := c.BindJSON(&body); err != nil {
			log.Error(err)
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return

		}

		var zone database.Zone

		zone.Id = uuid.New().String()
		zone.Name = body.Name
		zone.Fqdn = body.Fqdn

		if err := zoneService.Create(ctx, zone); err != nil {
			log.Error(err)
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, zone)
	}
}
