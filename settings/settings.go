package settings

// TargetAppName 解析対象の名前
var TargetAppName = "netflix"

// Hosts 解析対象のアプリ内で使用されているホスト名
var Hosts = []string{}

// AnalyticsItems - 解析する項目と正規表現とキーワード
var AnalyticsItems = map[string][]string{
	"GENDER":      {`(?i)[\s]sex[\s|:]`, `(?i)[\s]gender[\s|:]`},
	"PHONENUMBER": {`(?i)[\s]phone_number[\s|:]`, `(?i)[\s]phone[\s|:]`, `\d{2,5}[-()]\d{1,4}[-)]\d{4}`},
	"MAILADDRESS": {`(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)`},
	"PASSWORD":    {`(?i)[\s]password[\s|:]`, `(?i)[\s]psword[\s|:]`, `(?i)[\s]pswd[\s|:]`},
	"AGE":         {`(?i)[\s]age[\s|:]`, `(?i)[\s]years[\s|:]`},
	"LOCATION":    {`(?i)[\s]location[\s|:]`, `(?i)[\s]locate[\s|:]`, `(?i)[\s]latitude[\s|:]`, `(?i)[\s]lat[\s|:]`, `(?i)[\s]longitude[\s|:]`, `(?i)[\s]lon[\s|:]`},
}
