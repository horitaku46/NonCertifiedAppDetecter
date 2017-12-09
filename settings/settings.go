package settings

// TargetAppName 解析対象の名前
var TargetAppName = "qiita_non_certified"

// Hosts 解析対象のアプリ内で使用されているホスト名
var Hosts = []string{"qiita.com", ""}

// AnalyticsItems - 解析する項目と正規表現とキーワード
var AnalyticsItems = map[string][]string{
	"GENDER":       {`(?i)sex`, `(?i)gender`},
	"PHONENUMBER":  {`(?i)phone_number`, `(?i)phonenumber`, `(?i)phnumber`, `\d{2,5}[-()]\d{1,4}[-)]\d{4}`},
	"MAILADDRESS":  {`(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`},
	"PASSWORD":     {`(?i)password`, `(?i)psword`, `(?i)pswd`},
	"AGE":          {`(?i)["]age`, `(?i)year`},
	"LOCATION":     {`(?i)location`, `(?i)locate`, `(?i)latitude`, `(?i)["]lat`, `(?i)longitude`, `(?i)["]lon`, `(?i)[\s]geo`},
	"AUTHORIZE":    {`(?i)_id`, `(?i)scope`, `(?i)user`, `(?i)token`, `(?i)`},
	"USERINFO":     {`(?i)screen_name`, `(?i)_id`, `(?i)[^o]id[\s|:]`},
	"TAG_TIMELINE": {`(?i)[\s]tl`, `(?i)timeline`, `(?i)tag`, `(?i)follow`},
}

// // iOS - Twitter👇
// // TargetAppName 解析対象の名前
// var TargetAppName = "twitter_non_certified_name"
//
// // Hosts 解析対象のアプリ内で使用されているホスト名
// var Hosts = []string{"api.twitter.com", "userstream.twitter.com", "outlook.office365.com", "calendar.google.com"}
//
// // AnalyticsItems - 解析する項目と正規表現とキーワード
// var AnalyticsItems = map[string][]string{
// 	"GENDER":      {`(?i)sex`, `(?i)gender`},
// 	"PHONENUMBER": {`(?i)phone_number`, `(?i)phonenumber`, `(?i)phnumber`, `\d{2,5}[-()]\d{1,4}[-)]\d{4}`},
// 	"MAILADDRESS": {`(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`},
// 	"PASSWORD":    {`(?i)password`, `(?i)psword`, `(?i)pswd`},
// 	"AGE":         {`(?i)["]age`, `(?i)year`},
// 	"LOCATION":    {`(?i)location`, `(?i)locate`, `(?i)latitude`, `(?i)["]lat`, `(?i)longitude`, `(?i)["]lon`, `(?i)[\s]geo`},
// 	"USERINFO":    {`(?i)screen_name`, `(?i)_id`, `(?i)id[\s|:]`},
// 	"DM":          {`(?i)[\s]dm`, `(?i)direct`, `(?i)direct_messeage`, `(?i)directmesseage`},
// 	"BLOCK":       {`(?i)block`, `(?i)block_user`, `(?i)blockuser`, `(?i)blocklist`, `(?i)block_list`},
// 	"MUTE":        {`(?i)mute`, `(?i)mute_user`, `(?i)muteuser`, `(?i)block_list`, `(?i)block_list`},
// 	"FOLLOW":      {`(?i)follow`, `(?i)friend`, `(?i)friendship`, `(?i)follower`},
// 	"TIMELINE":    {`(?i)[\s]tl`, `(?i)timeline`},
// }

// // Android Twitter👇
// // TargetAppName 解析対象の名前
// var TargetAppName = "twitter_non_certified_name"
//
// // Hosts 解析対象のアプリ内で使用されているホスト名
// var Hosts = []string{"twitter.com", "mobile.twitter.com", "api.twitter.com", "userstream.twitter.com",
// 	"analytics.twitter.com", "android.clients.google.com"}
//
// // AnalyticsItems - 解析する項目と正規表現とキーワード
// var AnalyticsItems = map[string][]string{
// 	"GENDER":      {`(?i)sex`, `(?i)gender`},
// 	"PHONENUMBER": {`(?i)phone_number`, `(?i)phonenumber`, `(?i)phnumber`, `\d{2,5}[-()]\d{1,4}[-)]\d{4}`},
// 	"MAILADDRESS": {`(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`},
// 	"PASSWORD":    {`(?i)password`, `(?i)psword`, `(?i)pswd`},
// 	"AGE":         {`(?i)["]age`, `(?i)year`},
// 	"LOCATION":    {`(?i)location`, `(?i)locate`, `(?i)latitude`, `(?i)["]lat`, `(?i)longitude`, `(?i)["]lon`, `(?i)[\s]geo`},
// 	"USERINFO":    {`(?i)screen_name`, `(?i)_id`, `(?i)[^o]id[\s|:]`},
// 	"DM":          {`(?i)[\s]dm`, `(?i)direct`, `(?i)direct_messeage`, `(?i)directmesseage`},
// 	"BLOCK":       {`(?i)block`, `(?i)block_user`, `(?i)blockuser`, `(?i)blocklist`, `(?i)block_list`},
// 	"MUTE":        {`(?i)mute`, `(?i)mute_user`, `(?i)muteuser`, `(?i)block_list`, `(?i)block_list`},
// 	"FOLLOW":      {`(?i)follow`, `(?i)friend`, `(?i)friendship`},
// 	"TIMELINE":    {`(?i)[\s]tl`, `(?i)timeline`},
// }
