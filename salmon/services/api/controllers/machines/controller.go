package machines

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	machineService "github.com/eskpil/salmon/services/api/internal/machines"
	log "github.com/sirupsen/logrus"
)

func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		machines, err := machineService.GetAll(ctx)

		if err != nil {
			log.Error(err)
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, machines)
	}
}

func GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id := c.Param("id")
		machine, err := machineService.GetById(ctx, id)

		if err != nil {
			log.Error(err)
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, machine)
	}
}
