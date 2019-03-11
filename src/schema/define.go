package schema

import (
	"fmt"
)

var Config Conf

type Conf struct {
	ElasticSearch ElasticSearchConf
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
