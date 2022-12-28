package records

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/miekg/dns"
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

		if !strings.HasSuffix(body.Fqdn, ".") {
			c.String(http.StatusBadRequest, fmt.Errorf("\"%s\" is not a valid fqdn.", body.Fqdn).Error())
			c.Abort()
			return
		}

		if ok, err := regexp.Match(`^(([a-z0-9][a-z0-9\-]*[a-z0-9])|[a-z0-9]+\.)*([a-z]+|xn\-\-[a-z0-9]+)\.?$`, []byte(body.Fqdn)); err != nil || !ok {
			log.Error(err)
			c.String(http.StatusBadRequest, fmt.Errorf("\"%s\" is not a valid fqdn.", body.Fqdn).Error())
			c.Abort()
			return
		}

		var record database.Record

		switch body.Type {
		case dns.TypeA, dns.TypeAAAA:
			if fmt.Sprintf("%T", body.Value) != "string" {
				err := fmt.Errorf("\"value\" must be a string")
				c.String(http.StatusBadRequest, err.Error())
				c.Abort()
				return
			}
		case dns.TypeTXT:
			if fmt.Sprintf("%T", body.Value) != "[]string" {
				err := fmt.Errorf("\"value\" must be a string array")
				c.String(http.StatusBadRequest, err.Error())
				c.Abort()
				return

			}
		}

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
