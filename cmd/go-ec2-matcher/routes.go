package main

import "net/http"

type Route struct {
	Method      string
	Path        string
	Name        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"GET", "/", "Index", Index},
	Route{"GET", "/ec2/matcher/", "EC2MatchFinder", EC2MatchHandler},
	Route{"GET", "/ec2/prices/", "EC2PriceFinder", EC2PriceHandler},
}
