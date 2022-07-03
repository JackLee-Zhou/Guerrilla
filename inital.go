package main

import (
	"fmt"
	"time"
)
var loadTime time.Time

// TODO 加载配置，进行装载的初始化
func init(){
	fmt.Println(loadTime.UTC().String())
}