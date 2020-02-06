//
/*
Package cna-proxy is stands for bypass Apple Captive Network Assistant (CNA)
and open Captive Portal in Safari or other browser.

Simple example:
	cna-proxy

By default it runs on default 9998 port and 30 second
interval to store clients.

Usage example:
	SERVE_ADDR=:80 CACHE_TIMEOUT=60m cna-proxy

Web server handle only GET on "/hotspot-*.*":

* return redirect HTML on first client income to trigger CNA

* on second and other returns empty HTML to pretend about there is
open Intenet access

*/
package main // import "github.com/danyanya/cna-proxy"
