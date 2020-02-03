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
	Route{"GET", "/ec2/find_match/", "FindEC2Match", EC2MatchHandler},
}
