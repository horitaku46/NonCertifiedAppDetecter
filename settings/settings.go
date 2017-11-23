package settings

// TargetAppName 解析対象の名前
var TargetAppName = "oriless_non_certified"

// Hosts 解析対象のアプリ内で使用されているホスト名
var Hosts = []string{"api.twitter.com"}

// AnalyticsItems - 解析する項目と正規表現とキーワード
var AnalyticsItems = map[string][]string{
	"GENDER":      {`(?i)sex`, `(?i)gender`},
	"PHONENUMBER": {`(?i)phone_number`, `(?i)phone`, `\d{2,5}[-()]\d{1,4}[-)]\d{4}`},
	"MAILADDRESS": {`(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`},
	"PASSWORD":    {`(?i)password`, `(?i)psword`, `(?i)pswd`},
	"AGE":         {`(?i)[\s]age`, `(?i)years`},
	"LOCATION":    {`(?i)location`, `(?i)locate`, `(?i)latitude`, `(?i)lat`, `(?i)longitude`, `(?i)lon`},
	"DM":          {`(?i)dm`, `(?i)direct`},
	"USER":        {`(?i)screen_name`, `(?i)user_id`},
	"BLOCK":       {`(?i)block`},
	"MUTE":        {`(?i)mute`},
	"FOLLOW":      {`(?i)friends`, `(?i)friendships`, `(?i)followers`},
	"TIMELINE":    {`(?i)user_timeline`, `(?i)tl`, `(?i)timeline`},
	"MEDIA":       {`(?i)friends`, `(?i)friendships`, `(?i)followers`},
}
