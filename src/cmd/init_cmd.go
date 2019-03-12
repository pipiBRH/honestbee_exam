package cmd

import (
	"flag"
	"honestbee_exam/src/schema"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
)

func init() {
	flag.Set("logtostderr", "true")
	configPath := flag.String("config", "conf/dev.toml", "specific config file")
	flag.Parse()
	if _, err := toml.DecodeFile(*configPath, &schema.Config); err != nil {
		glog.Fatalf("Parser config error : %+v", err)
	}
}
