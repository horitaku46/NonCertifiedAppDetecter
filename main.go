package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sort"
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

			packet, _ := httputil.DumpRequestOut(req, true)
			// check(error)

			decodeURL, _ := url.QueryUnescape(req.URL.String())
			// check(error)

			decodePacket, _ := url.QueryUnescape(string(packet))
			// check(error)

			// 解析
			var packetInfoArray = analytics.AnalyzePacket(decodeURL, req.Host, decodePacket)

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

			var detectAllParts []string

			// 解析結果表示
			for index, packetInfo := range packetInfoArray {
				if index == 0 {
					// DBに保存
					dp := &models.DetectPacket{
						Packet:     *p,
						DetectInfo: packetInfoArray,
					}
					detectPacketsCol.Insert(dp)

					u, error := url.Parse(decodeURL)
					check(error)
					fmt.Println("Scheme: " + u.Scheme)
					fmt.Println("Host: " + u.Host)
					fmt.Printf("User: %s\n", u.User)
					fmt.Println("Path: " + u.Path)
					fmt.Println("RawPath: " + u.RawPath)
					fmt.Println("↓Quries↓")
					for key, values := range u.Query() {
						fmt.Printf("\x1b[36m"+"%s:\n"+"\x1b[0m", key)
						for _, v := range values {
							fmt.Printf("\x1b[32m"+"%s\n"+"\x1b[0m", v)
						}
					}

					fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------------------------------")

					if packetInfo.IsContainHost {
						fmt.Println("ホスト判定: " + "想定しているホストです.")
					} else {
						fmt.Println("ホスト判定: " + "想定していないホストです.")
					}
				}
				fmt.Println("検知情報: " + packetInfo.AnalyticsItemName + "が含まれている可能性があります.")
				fmt.Println("使用した正規表現やキーワード: " + packetInfo.RegexpKeyWord)
				fmt.Println("検知部分: " + "\x1b[31;43m" + strings.Join(packetInfo.DetectParts, ", ") + "\x1b[0m")
				fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------------------------------")
				detectAllParts = append(detectAllParts, packetInfo.DetectParts...)
				if index == len(packetInfoArray)-1 {
					printColorPacket(decodePacket, detectAllParts)
					fmt.Println("\x1b[31m================================================================================================================================================================\x1b[0m")
					r, _ := http.NewRequest("GET", "", nil)
					return r, nil
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

// TODO: - ロジックがゴミってるので再考
func printColorPacket(packet string, detectParts []string) {

	var uniqDetectParts = uniqStrArray(detectParts)
	var detects Detects = []Detect{}

	var indexs []int
	for _, part := range uniqDetectParts {
		detect := Detect{}
		detect.Part = part
		index := strings.Index(packet, part)
		if !isContainIndex(indexs, index) {
			detect.Index = index
			detects = append(detects, detect)
			indexs = append(indexs, index)
		}
	}
	// Index順に並び替え
	sort.Sort(detects)

	var remainPacket string

	for index, detect := range detects {
		var firstDetectPartIndex int
		var packetLastIndex int
		var nonDetectPart string

		if index == 0 {
			firstDetectPartIndex = strings.Index(packet, detect.Part)
			packetLastIndex = utf8.RuneCountInString(packet)
			nonDetectPart = packet[0:firstDetectPartIndex]
		} else {
			firstDetectPartIndex = strings.Index(remainPacket, detect.Part)
			packetLastIndex = utf8.RuneCountInString(remainPacket)
			nonDetectPart = remainPacket[0:firstDetectPartIndex]
		}

		lastDetectPartIndex := firstDetectPartIndex + utf8.RuneCountInString(detect.Part)
		if index == 0 {
			remainPacket = packet[lastDetectPartIndex:packetLastIndex]
		} else {
			remainPacket = remainPacket[lastDetectPartIndex:packetLastIndex]
		}

		fmt.Print(nonDetectPart)
		fmt.Print("\x1b[31;43m" + detect.Part + "\x1b[0m")

		if index == len(detects)-1 {
			fmt.Println(remainPacket)
		}
	}
}

// Detect -
type Detect struct {
	Part  string
	Index int
}

// Detects -
type Detects []Detect

func (d Detects) Len() int {
	return len(d)
}

func (d Detects) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Detects) Less(i, j int) bool {
	return d[i].Index < d[j].Index
}

func uniqStrArray(strs []string) []string {
	m := make(map[string]bool)
	uniq := []string{}
	for _, ele := range strs {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	return uniq
}

func isContainIndex(arr []int, index int) bool {
	for _, v := range arr {
		if v == index {
			return true
		}
	}
	return false
}
