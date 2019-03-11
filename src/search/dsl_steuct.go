package search

type query struct {
	Query match `json:"query"`
}

type match struct {
	Match matchDsl `json:"match"`
}

type matchDsl struct {
	Pname string `json:"pname"`
}

type QueryData struct {
	Took    int   `json:"took"`
	Timeout bool  `json:"timed_out"`
	Shards  shard `json:"_shards"`
	Hits    hit   `json:"hits"`
}

type hit struct {
	Total    int       `json:"total"`
	MaxScore float32   `json:"max_score"`
	HitList  []hitList `json:"hits"`
}

type shard struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type hitList struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float32 `json:"_score"`
	Source goods   `json:"_source"`
}

type goods struct {
	Gdid  int    `json:"pid"`
	Pname string `json:"pname"`
	Url   string `json:"url"`
	Image string `json:"image"`
	Price string `json:"price"`
}
