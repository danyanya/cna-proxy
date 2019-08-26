package main

import (
	"fmt"
	"net"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/patrickmn/go-cache"
)

type environment struct {
	ServeAddr string `envconfig:"SERVE_ADDR" default:":9998"`
}

func main() {

	var env environment
	var err = envconfig.Process("", &env)
	if err != nil {
		panic(err)
	}

	var ch = cache.New(5*time.Minute, 10*time.Minute)

	var e = echo.New()
	e.Use(middleware.Recover())

	e.GET("/hotspot-*.",
		func(c echo.Context) error {

			var clientID, _, _ = net.SplitHostPort(c.Request().RemoteAddr)
			fmt.Printf("addr: %s\n", clientID)
			if _, found := ch.Get(clientID); found {
				return c.HTML(200,
					`
				<HTML>
         		 <HEAD><TITLE>Success</TITLE></HEAD>
         		 <BODY>
       		     <a href="http://ya.ru">This link will open in Safari</a>
       			   </BODY>
     		   </HTML> 
				`,
				)
			}
			ch.Set(clientID, true, cache.DefaultExpiration)
			return c.HTML(200, "<html></html>")
		},
	)

	if err = e.Start(env.ServeAddr); err != nil {
		panic(err.Error())
	}
}
