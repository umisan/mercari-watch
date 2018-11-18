package config

import "os"

var CRAWL_DURATION int = 1 //分
var DATABASE_URL string = os.Getenv("DATABASE_URL")
