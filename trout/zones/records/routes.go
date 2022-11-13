package records

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/eskpil/cod/trout/database"
	"github.com/google/uuid"

	recordService "github.com/eskpil/cod/trout/internal/records"

	"github.com/gin-gonic/gin"
)

func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		zoneId := c.Param("zoneId")

		records, err := recordService.GetAll(ctx, zoneId)
		if err != nil {
			log.Error(err)
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, records)
	}
}

func CreateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		zoneId := c.Param("zoneId")

		var body struct {
			Fqdn  string      `json:"fqdn"`
			Type  uint16      `json:"type"`
			Value interface{} `json:"value"`
		}

		if err := c.BindJSON(&body); err != nil {
			log.Error(err)
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return

		}

		var record database.Record

		record.Id = uuid.New().String()
		record.Fqdn = body.Fqdn
		record.Type = body.Type
		record.Value = body.Value
		record.Ttl = uint32((30 * time.Minute).Seconds())
		record.ZoneId = zoneId

		if err := recordService.Create(ctx, record); err != nil {
			log.Error(err)
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

	}
}
