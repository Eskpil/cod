package main

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	machineController "github.com/eskpil/salmon/services/api/controllers/machines"
	"github.com/eskpil/salmon/services/api/internal"
)

func main() {
	e := gin.Default()

	ctx, err := internal.NewContext()

	if err != nil {
		log.Fatalf("Failed to initialize a new mycontext: %v\n", err)
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		ticker := time.NewTicker(20 * time.Second)
		quit := make(chan struct{})

		log.Info("Performing initial routine")

		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		go ctx.PerformRoutine(c)

		log.Info("Starting main loop")
		for {
			select {
			case <-ticker.C:
				go func() {
					c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
					defer cancel()

					ctx.PerformRoutine(c)
				}()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}(wg)

	gin.SetMode(gin.ReleaseMode)

	e.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	e.GET("/api/machines/", machineController.GetAll())
	e.GET("/api/machines/:id/", machineController.GetById())

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		e.Run("0.0.0.0:8090")
	}(wg)

	wg.Wait()
}
