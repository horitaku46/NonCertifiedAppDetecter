package analytics

import (
	"regexp"

	"github.com/horitaku46/NonCertifiedAppDetecter/models"
	"github.com/horitaku46/NonCertifiedAppDetecter/settings"
)

func isContainHost(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// AnalyzePacket 正規表現を用いてパケットを解析
func AnalyzePacket(url string, host string, packet string) []models.DetectInfo {
	packetInfoArray := []models.DetectInfo{}

	for analyticsItemName, regexpKeywords := range settings.AnalyticsItems {

		for _, regexpKeyword := range regexpKeywords {
			rep := regexp.MustCompile(regexpKeyword)
			detectParts := rep.FindAllString(packet, -1)

			if detectParts != nil {
				packetInfo := models.DetectInfo{}
				packetInfo.IsContainHost = isContainHost(settings.Hosts, host)
				packetInfo.AnalyticsItemName = analyticsItemName
				packetInfo.RegexpKeyWord = regexpKeyword
				packetInfo.DetectParts = detectParts
				packetInfoArray = append(packetInfoArray, packetInfo)
			}
		}
	}
	return packetInfoArray
}
