// Code generated by make generate-models DO NOT EDIT.

package private

type GetUserTradesByCurrencyAndTimeRequest struct {
	Count          int64  `json:"count" mapstructure:"count"`
	Currency       string `json:"currency" mapstructure:"currency"`
	EndTimestamp   int64  `json:"end_timestamp" mapstructure:"end_timestamp"`
	StartTimestamp int64  `json:"start_timestamp" mapstructure:"start_timestamp"`
}
