package main

import (
	"fmt"
	"time"
)

var loadTime time.Time

// TODO load the config
func init() {
	fmt.Println(loadTime.UTC().String())
}
