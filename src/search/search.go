package search

import (
	"encoding/json"
	"honestbee_exam/src/schema"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/glog"
)

func EsQuery(keyWord string) (*QueryData, error) {
	url := schema.Config.ElasticSearch.GetURL()

	dsl := query{match{matchDsl{keyWord}}}
	dslJsonEncode, _ := json.Marshal(dsl)
	payload := strings.NewReader(string(dslJsonEncode))
	req, _ := http.NewRequest("GET", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		glog.Errorf("EsQuery set level error => %+v\n  KW => %+v", err, keyWord)
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var result *QueryData
	err = json.Unmarshal(body, &result)
	if err != nil {
		glog.Errorf("EsQuery set level error => %+v\n  KW => %+v", err, keyWord)
		return nil, err
	}
	return result, nil
}
