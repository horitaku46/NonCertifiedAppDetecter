package models

// DetectInfo -
type DetectInfo struct {
	IsContainHost     bool     `bson:"is_contain_host"`
	AnalyticsItemName string   `bson:"analytics_item_name"`
	RegexpKeyWord     string   `bson:"regexp_keyword"`
	DetectParts       []string `bson:"detext_parts"`
}
