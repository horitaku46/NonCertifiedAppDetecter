package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/elazarl/goproxy"
	"github.com/horitaku46/NonCertifiedAppDetecter/analytics"
	"github.com/horitaku46/NonCertifiedAppDetecter/models"
	"github.com/horitaku46/NonCertifiedAppDetecter/settings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	session, error := mgo.Dial("mongodb://localhost/" + settings.TargetAppName)
	check(error)
	defer session.Close()
	db := session.DB(settings.TargetAppName)
	packetsCol := db.C("packets")
	detectPacketsCol := db.C("detectPackets")

	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()

	setCA(caCert, caKey)
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {

			packet, error := httputil.DumpRequestOut(req, true)
			check(error)

			decodeURL, error := url.QueryUnescape(req.URL.String())
			check(error)

			// 解析
			var packetInfoArray = analytics.AnalyzePacket(decodeURL, req.Host, string(packet))

			// DBに保存
			p := &models.Packet{
				ID:         bson.NewObjectId(),
				Timestamp:  time.Now(),
				URL:        decodeURL,
				HTTPMethod: req.Method,
				Host:       req.Host,
				Packet:     string(packet),
			}
			packetsCol.Insert(p)

			// 解析結果表示
			for index, packetInfo := range packetInfoArray {
				if index == 0 {
					// DBに保存
					dp := &models.DetectPacket{
						Packet:     *p,
						DetectInfo: packetInfoArray,
					}
					detectPacketsCol.Insert(dp)

					if packetInfo.IsContainHost {
						fmt.Println("ホスト判定: " + "想定しているホストです.")
					} else {
						fmt.Println("ホスト判定: " + "想定していないホストです.")
					}
				}
				fmt.Println("検知情報: " + packetInfo.AnalyticsItemName + "が含まれている可能性があります.")
				fmt.Println("使用した正規表現やキーワード: " + packetInfo.RegexpKeyWord)
				fmt.Println("検知項目: " + "\x1b[31;43m" + strings.Join(packetInfo.DetectParts, ", ") + "\x1b[0m")
				fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------------------------------")
				if index == len(packetInfoArray)-1 {
					printColorPacket(string(packet), packetInfo.DetectParts)
					fmt.Println("================================================================================================================================================================")
				}
			}
			return req, nil
		})

	fmt.Println("---------------------------------------------------------------------------------------------------------")
	fmt.Println("|----------------------------------NonCertifiedAppDetector Start----------------------------------------|")
	fmt.Println("---------------------------------------------------------------------------------------------------------")

	proxy.Verbose = *verbose
	log.Fatal(http.ListenAndServe(*addr, proxy))
}

var remainPacket string

// TODO: - ロジックがゴミってるので再考
func printColorPacket(packet string, detectParts []string) {
	for index, detectPart := range detectParts {
		var firstDetectPartIndex int
		var packetLastIndex int
		var nonDetectPart string

		if index == 0 {
			firstDetectPartIndex = strings.Index(packet, detectPart)
			packetLastIndex = utf8.RuneCountInString(packet)
			nonDetectPart = packet[0:firstDetectPartIndex]
		} else {
			firstDetectPartIndex = strings.Index(remainPacket, detectPart)
			packetLastIndex = utf8.RuneCountInString(remainPacket)
			nonDetectPart = remainPacket[0:firstDetectPartIndex]
		}

		lastDetectPartIndex := firstDetectPartIndex + utf8.RuneCountInString(detectPart)
		if index == 0 {
			remainPacket = packet[lastDetectPartIndex:packetLastIndex]
		} else {
			remainPacket = remainPacket[lastDetectPartIndex:packetLastIndex]
		}

		fmt.Print(nonDetectPart)
		fmt.Print("\x1b[31;43m" + detectPart + "\x1b[0m")

		if index == len(detectParts)-1 {
			fmt.Println(remainPacket)
		}
	}
}
