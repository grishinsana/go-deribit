// Code generated by make generate-models DO NOT EDIT.

package notifications

type DeribitPriceRankingIndexNameRepeated []struct {
	Enabled    bool    `json:"enabled" mapstructure:"enabled"`
	Identifier string  `json:"identifier" mapstructure:"identifier"`
	Price      float64 `json:"price" mapstructure:"price"`
	Timestamp  int64   `json:"timestamp" mapstructure:"timestamp"`
	Weight     int64   `json:"weight" mapstructure:"weight"`
}
