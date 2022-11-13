package main

import (
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"

	dnsService "github.com/eskpil/cod/trout/internal/dns"
	"github.com/eskpil/cod/trout/zones"
	"github.com/gin-gonic/gin"
	"github.com/miekg/dns"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	wg := &sync.WaitGroup{}
	{
		srv := &dns.Server{Addr: "192.168.0.72:" + strconv.Itoa(53), Net: "udp"}
		srv.Handler = dnsService.New()

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			if err := srv.ListenAndServe(); err != nil {
				log.Fatalf("Failed to set udp listener %s\n", err.Error())
			}
		}(wg)
	}
	{
		e := gin.Default()

		api := e.Group("/api/resources/")
		api.Use(func(c *gin.Context) {
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

		zones.SubRoutes(api.Group("/zones"))

		api.GET("zones/", zones.GetAll())
		api.POST("zones/", zones.CreateOne())

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			e.Run("0.0.0.0:8090")
		}(wg)
	}

	wg.Wait()
}
