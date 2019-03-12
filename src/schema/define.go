package schema

import (
	"fmt"
)

var Config Conf

type Conf struct {
	ElasticSearch ElasticSearchConf
	Server        ServerConf
}

type ServerConf struct {
	bind string
	Port int
}

func (this ServerConf) GetBind() string {
	return fmt.Sprintf(
		"%v:%v",
		this.bind,
		this.Port)
}

type ElasticSearchConf struct {
	Host  string
	Port  int
	Index string
	Type  string
}

func (this ElasticSearchConf) GetURL() string {
	return fmt.Sprintf(
		"http://%v:%v/%v/%v/_search",
		this.Host,
		this.Port,
		this.Index,
		this.Type)
}
