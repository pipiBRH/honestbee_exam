package main

import (
	"fmt"
	"github.com/golang/glog"
	_ "honestbee_exam/src/cmd"
	"honestbee_exam/src/search"
)

func main() {
	defer glog.Flush()
	glog.Info("Spider GO!")

	a, _ := search.EsQuery("iphone")
	fmt.Println(a.Hits.HitList)
}
