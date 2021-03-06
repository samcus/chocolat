package api

import (
	"encoding/json"
	"github.com/angdev/chocolat/lib/query"
)

type inspectResult struct {
	Name       string      `json:"name"`
	Properties interface{} `json:"properties"`
	Url        string      `json:"url"`
}

type queryResult struct {
	Result interface{} `json:"result"`
}

type queryGroupResult struct {
	Result interface{}
	Groups RawResult
}

func (this *queryGroupResult) MarshalJSON() ([]byte, error) {
	result := make(RawResult)
	result["result"] = this.Result

	collapsed := make(RawResult)
	query.CollapseField(this.Groups, &collapsed)

	for k, v := range collapsed {
		result[k] = v
	}

	return json.Marshal(result)
}

type queryGroupResultArray []queryGroupResult

type queryIntervalResult struct {
	Result    interface{} `json:"value"`
	TimeFrame TimeFrame   `json:"timeframe"`
}

func (this *queryIntervalResult) MarshalJSON() ([]byte, error) {
	result := make(RawResult)

	switch this.Result.(type) {
	case queryResult:
		result["value"] = this.Result.(queryResult).Result
	default:
		result["value"] = this.Result
	}

	result["timeframe"] = this.TimeFrame
	return json.Marshal(result)
}
