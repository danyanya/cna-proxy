package main

import (
	"io/ioutil"
	"net"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/patrickmn/go-cache"
)

type environment struct {
	ServeAddr    string        `envconfig:"SERVE_ADDR" default:":9998"`
	RedirectHtml string        `envconfig:"REDIRECT_HTML" default:""`
	CacheTimeout time.Duration `envconfig:"CACHE_TIMEOUT" default:"10s"`
}

var defaultRedirectHtml = []byte(`
<HTML>
<HEAD>
<meta charset="UTF-8">
<TITLE>Success</TITLE>
</HEAD>
<body>
<script>
     function addElement(absoluteurl)
     {
        var anchor = document.createElement('a');
        anchor.href = absoluteurl;
        window.document.body.insertBefore(anchor, window.document.body.firstChild);
        setTimeout(function(){ anchor.click(); }, 500);
        // document.body.appendChild(elemDiv); // appends last of that element
    }
    addElement("http://gowifi.ru");
</script>
</body>
</HTML>
`)

func main() {
	var env environment
	var err = envconfig.Process("", &env)
	if err != nil {
		panic(err)
	}

	redirectHtml, err := ioutil.ReadFile(env.RedirectHtml)
	if err != nil || len(redirectHtml) == 0 {
		redirectHtml = defaultRedirectHtml
	}

	var ch = cache.New(env.CacheTimeout, env.CacheTimeout)
	var e = echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Apple iOS common internet detection URI
	e.GET("/hotspot-*.*",
		func(c echo.Context) error {
			var clientID, _, _ = net.SplitHostPort(c.Request().RemoteAddr)
			if _, found := ch.Get(clientID); found {
				return c.HTML(200, string(redirectHtml))
			}
			ch.Set(clientID, true, cache.DefaultExpiration)
			return c.HTML(200, "<html></html>")
		},
	)

	if err = e.Start(env.ServeAddr); err != nil {
		panic(err.Error())
	}
}
