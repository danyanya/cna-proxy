// cna-proxy: server for bypass Apple CNA

// Web server handle only path "/hotspot-*.*"
// and returns redirect HTML on first client
// income, on second and other returns empty HTML

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

// environment contains defautl config from environment
type environment struct {
	// serve address (host:port)
	ServeAddr string `envconfig:"SERVE_ADDR" default:":9998"`
	// redirect html path on disc
	RedirectHTML string `envconfig:"REDIRECT_HTML" default:""`
	// cache timeout to store client data
	CacheTimeout time.Duration `envconfig:"CACHE_TIMEOUT" default:"10s"`
}

// defaultRedirectHTML return redirection (with delayed for 500 ms) to gowifi.ru
var defaultRedirectHTML = []byte(`
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

	redirectHTMLbyte, err := ioutil.ReadFile(env.RedirectHTML)
	if err != nil || len(redirectHTMLbyte) == 0 {
		redirectHTMLbyte = defaultRedirectHTML
	}

	var redirectHTML = string(redirectHTMLbyte)

	// cache initialization
	var ch = cache.New(env.CacheTimeout, env.CacheTimeout)

	// router initialization
	var e = echo.New()
	e.HideBanner = true

	// include standard recover and logger for router
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// set handler for Apple CNA common internet detection URI
	e.GET("/hotspot-*.*",
		func(c echo.Context) error {
			var clientID, _, _ = net.SplitHostPort(c.Request().RemoteAddr)
			if _, found := ch.Get(clientID); found {
				return c.HTML(200, string(redirectHTML))
			}
			ch.Set(clientID, true, cache.DefaultExpiration)
			return c.HTML(200, "<html></html>")
		},
	)

	// start server
	if err = e.Start(env.ServeAddr); err != nil {
		panic(err.Error())
	}
}
