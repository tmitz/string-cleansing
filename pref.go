package stringcleansing

import "fmt"
import "strings"

type Pref struct {
	Address, Code string
}

func PrefixPref(address string, pref string) string {
	res := address

	if !strings.Contains(address, pref) {
		res = pref + address
	}

	return res
}

func PrefCodeString(n string) (Pref, error) {
	pref := Pref{}
	switch n {
	case "01":
		pref.Code = "hokkaido"
		pref.Address = "北海道"
	case "02":
		pref.Code = "aomori"
		pref.Address = "青森県"
	case "03":
		pref.Code = "iwate"
		pref.Address = "岩手県"
	case "04":
		pref.Code = "miyagi"
		pref.Address = "宮城県"
	case "05":
		pref.Code = "akita"
		pref.Address = "秋田県"
	case "06":
		pref.Code = "yamagata"
		pref.Address = "山形県"
	case "07":
		pref.Code = "fukushima"
		pref.Address = "福島県"
	case "08":
		pref.Code = "ibaraki"
		pref.Address = "茨城県"
	case "09":
		pref.Code = "tochigi"
		pref.Address = "栃木県"
	case "10":
		pref.Code = "gunma"
		pref.Address = "群馬県"
	case "11":
		pref.Code = "saitama"
		pref.Address = "埼玉県"
	case "12":
		pref.Code = "chiba"
		pref.Address = "千葉県"
	case "13":
		pref.Code = "tokyo"
		pref.Address = "東京都"
	case "14":
		pref.Code = "kanagawa"
		pref.Address = "神奈川県"
	case "15":
		pref.Code = "niigata"
		pref.Address = "新潟県"
	case "16":
		pref.Code = "toyama"
		pref.Address = "富山県"
	case "17":
		pref.Code = "ishikawa"
		pref.Address = "石川県"
	case "18":
		pref.Code = "fukui"
		pref.Address = "福井県"
	case "19":
		pref.Code = "yamanashi"
		pref.Address = "山梨県"
	case "20":
		pref.Code = "nagano"
		pref.Address = "長野県"
	case "21":
		pref.Code = "gifu"
		pref.Address = "岐阜県"
	case "22":
		pref.Code = "shizuoka"
		pref.Address = "静岡県"
	case "23":
		pref.Code = "aichi"
		pref.Address = "愛知県"
	case "24":
		pref.Code = "mie"
		pref.Address = "三重県"
	case "25":
		pref.Code = "shiga"
		pref.Address = "滋賀県"
	case "26":
		pref.Code = "kyoto"
		pref.Address = "京都府"
	case "27":
		pref.Code = "osaka"
		pref.Address = "大阪府"
	case "28":
		pref.Code = "hyogo"
		pref.Address = "兵庫県"
	case "29":
		pref.Code = "nara"
		pref.Address = "奈良県"
	case "30":
		pref.Code = "wakayama"
		pref.Address = "和歌山県"
	case "31":
		pref.Code = "tottori"
		pref.Address = "鳥取県"
	case "32":
		pref.Code = "shimane"
		pref.Address = "島根県"
	case "33":
		pref.Code = "okayama"
		pref.Address = "岡山県"
	case "34":
		pref.Code = "hiroshima"
		pref.Address = "広島県"
	case "35":
		pref.Code = "yamaguchi"
		pref.Address = "山口県"
	case "36":
		pref.Code = "tokushima"
		pref.Address = "徳島県"
	case "37":
		pref.Code = "kagawa"
		pref.Address = "香川県"
	case "38":
		pref.Code = "ehime"
		pref.Address = "愛媛県"
	case "39":
		pref.Code = "kochi"
		pref.Address = "高知県"
	case "40":
		pref.Code = "fukuoka"
		pref.Address = "福岡県"
	case "41":
		pref.Code = "saga"
		pref.Address = "佐賀県"
	case "42":
		pref.Code = "nagasaki"
		pref.Address = "長崎県"
	case "43":
		pref.Code = "kumamoto"
		pref.Address = "熊本県"
	case "44":
		pref.Code = "oita"
		pref.Address = "大分県"
	case "45":
		pref.Code = "miyazaki"
		pref.Address = "宮崎県"
	case "46":
		pref.Code = "kagoshima"
		pref.Address = "鹿児島県"
	case "47":
		pref.Code = "okinawa"
		pref.Address = "沖縄県"
	default:
		return pref, fmt.Errorf("Not found prefecture. code: %s", n)
	}
	return pref, nil
}
