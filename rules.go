package vphong

// Onset represents all possible onsets in Vietnamese
var Onset = []string{
	"b", "d", "h", "l", "m", "n", "p", "r", "s", "t", "v", "x", "đ", "p",
	"tr", "th", "ch", "ph", "nh", "kh", "gi", "qu",
	"ngh", "ng", "gh", "g", "k", "c",
}

// CusOnsets maps Vietnamese onsets to their IPA equivalents
var CusOnsets = map[string]string{
	"p":   "p",
	"b":   "ɓ",
	"m":   "m",
	"ph":  "f",
	"v":   "v",
	"th":  "tʰ",
	"t":   "t",
	"đ":   "d",
	"n":   "n",
	"x":   "s",
	"d":   "z",
	"gi":  "z",
	"l":   "l",
	"tr":  "ʈ",
	"s":   "ʂ",
	"r":   "ʐ",
	"ch":  "c",
	"nh":  "ɲ",
	"k":   "k",
	"c":   "k",
	"qu":  "kw",
	"ng":  "ŋ",
	"ngh": "ŋ",
	"kh":  "x",
	"g":   "ɣ",
	"gh":  "ɣ",
	"h":   "h",
}

// CusGi maps "gi" and its tones to IPA
var CusGi = map[string]string{
	"gi": "zi", "gí": "zi", "gì": "zi", "gỉ": "zi", "gĩ": "zi", "gị": "zi",
}

// CusQu maps "quy" and its tones to IPA
var CusQu = map[string]string{
	"quy": "kwi", "qúy": "kwi", "qùy": "kwi", "qủy": "kwi", "qũy": "kwi", "qụy": "kwi",
}

// CusNuclei maps Vietnamese nuclei (vowels) to their IPA equivalents
var CusNuclei = map[string]string{
	"a": "a", "á": "a", "à": "a", "ả": "a", "ã": "a", "ạ": "a",
	"â": "ɤ̆", "ấ": "ɤ̆", "ầ": "ɤ̆", "ẩ": "ɤ̆", "ẫ": "ɤ̆", "ậ": "ɤ̆",
	"ă": "ă", "ắ": "ă", "ằ": "ă", "ẳ": "ă", "ẵ": "ă", "ặ": "ă",
	"e": "ɛ", "é": "ɛ", "è": "ɛ", "ẻ": "ɛ", "ẽ": "ɛ", "ẹ": "ɛ",
	"ê": "e", "ế": "e", "ề": "e", "ể": "e", "ễ": "e", "ệ": "e",
	"i": "i", "í": "i", "ì": "i", "ỉ": "i", "ĩ": "i", "ị": "i",
	"o": "ɔ", "ó": "ɔ", "ò": "ɔ", "ỏ": "ɔ", "õ": "ɔ", "ọ": "ɔ",
	"ô": "o", "ố": "o", "ồ": "o", "ổ": "o", "ỗ": "o", "ộ": "o",
	"ơ": "ɤ", "ớ": "ɤ", "ờ": "ɤ", "ở": "ɤ", "ỡ": "ɤ", "ợ": "ɤ",
	"u": "u", "ú": "u", "ù": "u", "ủ": "u", "ũ": "u", "ụ": "u",
	"ư": "ɯ", "ứ": "ɯ", "ừ": "ɯ", "ử": "ɯ", "ữ": "ɯ", "ự": "ɯ",
	"y": "i", "ý": "i", "ỳ": "i", "ỷ": "i", "ỹ": "i", "ỵ": "i",
	"eo": "ɛw", "éo": "ɛw", "èo": "ɛw", "ẻo": "ɛw", "ẽo": "ɛw", "ẹo": "ɛw",
	"êu": "ew", "ếu": "ew", "ều": "ew", "ểu": "ew", "ễu": "ew", "ệu": "ew",
	"ia": "iɛ", "ía": "iɛ", "ìa": "iɛ", "ỉa": "iɛ", "ĩa": "iɛ", "ịa": "iɛ",
	"iê": "ie", "iế": "ie", "iể": "ie", "iề": "ie", "iễ": "ie", "iệ": "ie",
	"oo": "ɔ", "óo": "ɔ", "òo": "ɔ", "ỏo": "ɔ", "õo": "ɔ", "ọo": "ɔ",
	"oó": "ɔ", "oò": "ɔ", "oỏ": "ɔ", "oõ": "ɔ", "oọ": "ɔ",
	"ôô": "o", "ốô": "o", "ồô": "o", "ổô": "o", "ỗô": "o", "ộô": "o",
	"ôố": "o", "ôồ": "o", "ôổ": "o", "ôỗ": "o", "ôộ": "o",
	"ua": "uo", "úa": "uo", "ùa": "uo", "ủa": "uo", "ũa": "uo", "ụa": "uo",
	"uô": "uo", "uố": "uo", "uồ": "uo", "uổ": "uo", "uỗ": "uo", "uộ": "uo",
	"ưa": "ɯə", "ứa": "ɯə", "ừa": "ɯə", "ửa": "ɯə", "ữa": "ɯə", "ựa": "ɯə",
	"ươ": "ɯə", "ướ": "ɯə", "ườ": "ɯə", "ưở": "ɯə", "ưỡ": "ɯə", "ượ": "ɯə",
	"yê": "ie", "yế": "ie", "yề": "ie", "yể": "ie", "yễ": "ie", "yệ": "ie",
	"uơ": "uə", "uớ": "uə", "uờ": "uə", "uở": "uə", "uỡ": "uə", "uợ": "uə",
}

// CusOffglides maps Vietnamese offglides to their IPA equivalents
var CusOffglides = map[string]string{
	"ai": "aj", "ái": "aj", "ài": "aj", "ải": "aj", "ãi": "aj", "ại": "aj",
	"ay": "ăj", "áy": "ăj", "ày": "ăj", "ảy": "ăj", "ãy": "ăj", "ạy": "ăj",
	"ao": "aw", "áo": "aw", "ào": "aw", "ảo": "aw", "ão": "aw", "ạo": "aw",
	"au": "ăw", "áu": "ăw", "àu": "ăw", "ảu": "ăw", "ãu": "ăw", "ạu": "ăw",
	"ây": "ɤ̆j", "ấy": "ɤ̆j", "ầy": "ɤ̆j", "ẩy": "ɤ̆j", "ẫy": "ɤ̆j", "ậy": "ɤ̆j",
	"âu": "ɤ̆w", "ấu": "ɤ̆w", "ầu": "ɤ̆w", "ẩu": "ɤ̆w", "ẫu": "ɤ̆w", "ậu": "ɤ̆w",
	"eo": "ew", "éo": "ew", "èo": "ew", "ẻo": "ew", "ẽo": "ew", "ẹo": "ew",
	"iu": "iw", "íu": "iw", "ìu": "iw", "ỉu": "iw", "ĩu": "iw", "ịu": "iw",
	"oi": "ɔj", "ói": "ɔj", "òi": "ɔj", "ỏi": "ɔj", "õi": "ɔj", "ọi": "ɔj",
	"ôi": "oj", "ối": "oj", "ồi": "oj", "ổi": "oj", "ỗi": "oj", "ội": "oj",
	"ui": "uj", "úi": "uj", "ùi": "uj", "ủi": "uj", "ũi": "uj", "ụi": "uj",
	"uy": "ʷi", "úy": "uj", "ùy": "uj", "ủy": "uj", "ũy": "uj", "ụy": "uj",
	"uý": "ʷi", "uỳ": "ʷi", "uỷ": "ʷi", "uỹ": "ʷi", "uỵ": "ʷi",
	"ơi": "ɤj", "ới": "ɤj", "ời": "ɤj", "ởi": "ɤj", "ỡi": "ɤj", "ợi": "ɤj",
	"ưi": "ɯj", "ứi": "ɯj", "ừi": "ɯj", "ửi": "ɯj", "ữi": "ɯj", "ựi": "ɯj",
	"ưu": "ɯw", "ứu": "ɯw", "ừu": "ɯw", "ửu": "ɯw", "ữu": "ɯw", "ựu": "ɯw",
	"iêu": "iew", "iếu": "iew", "iều": "iew", "iểu": "iew", "iễu": "iew", "iệu": "iew",
	"yêu": "iew", "yếu": "iew", "yều": "iew", "yểu": "iew", "yễu": "iew", "yệu": "iew",
	"uôi": "uəj", "uối": "uəj", "uồi": "uəj", "uổi": "uəj", "uỗi": "uəj", "uội": "uəj",
	"ươi": "ɯəj", "ưới": "ɯəj", "ười": "ɯəj", "ưởi": "ɯəj", "ưỡi": "ɯəj", "ượi": "ɯəj",
	"ươu": "ɯəw", "ướu": "ɯəw", "ườu": "ɯəw", "ưởu": "ɯəw", "ưỡu": "ɯəw", "ượu": "ɯəw",
}

// CusOnglides maps Vietnamese onglides to their IPA equivalents
var CusOnglides = map[string]string{
	"oa": "ʷa", "oá": "ʷa", "oà": "ʷa", "oả": "ʷa", "oã": "ʷa", "oạ": "ʷa",
	"óa": "ʷa", "òa": "ʷa", "ỏa": "ʷa", "õa": "ʷa", "ọa": "ʷa",
	"oă": "ʷă", "oắ": "ʷă", "oằ": "ʷă", "oẳ": "ʷă", "oẵ": "ʷă", "oặ": "ʷă",
	"oe": "ʷɛ", "oé": "ʷɛ", "oè": "ʷɛ", "oẻ": "ʷɛ", "oẽ": "ʷɛ", "oẹ": "ʷɛ",
	"óe": "ʷɛ", "òe": "ʷɛ", "ỏe": "ʷɛ", "õe": "ʷɛ", "ọe": "ʷɛ",
	"ua": "ʷa", "uá": "ʷa", "uà": "ʷa", "uả": "ʷa", "uã": "ʷa", "uạ": "ʷa",
	"uă": "ʷă", "uắ": "ʷă", "uằ": "ʷă", "uẳ": "ʷă", "uẵ": "ʷă", "uặ": "ʷă",
	"uâ": "ʷɤ̆", "uấ": "ʷɤ̆", "uầ": "ʷɤ̆", "uẩ": "ʷɤ̆", "uẫ": "ʷɤ̆", "uậ": "ʷɤ̆",
	"ue": "ʷɛ", "ué": "ʷɛ", "uè": "ʷɛ", "uẻ": "ʷɛ", "uẽ": "ʷɛ", "uẹ": "ʷɛ",
	"uê": "ʷe", "uế": "ʷe", "uề": "ʷe", "uể": "ʷe", "uễ": "ʷe", "uệ": "ʷe",
	"uơ": "ʷɤ", "uớ": "ʷɤ", "uờ": "ʷɤ", "uở": "ʷɤ", "uỡ": "ʷɤ", "uợ": "ʷɤ",
	"uy": "ʷi", "uý": "ʷi", "uỳ": "ʷi", "uỷ": "ʷi", "uỹ": "ʷi", "uỵ": "ʷi",
	"uya": "ʷiə", "uyá": "ʷiə", "uyà": "ʷiə", "uyả": "ʷiə", "uyã": "ʷiə", "uyạ": "ʷiə",
	"uyê": "ʷiə", "uyế": "ʷiə", "uyề": "ʷiə", "uyể": "ʷiə", "uyễ": "ʷiə", "uyệ": "ʷiə",
	"uyu": "ʷiu", "uyú": "ʷiu", "uyù": "ʷiu", "uyủ": "ʷiu", "uyũ": "ʷiu", "uyụ": "ʷiu",
	"uýu": "ʷiu", "uỳu": "ʷiu", "uỷu": "ʷiu", "uỹu": "ʷiu", "uỵu": "ʷiu",
	"oen": "ʷen", "oén": "ʷen", "oèn": "ʷen", "oẻn": "ʷen", "oẽn": "ʷen", "oẹn": "ʷen",
	"oet": "ʷet", "oét": "ʷet", "oèt": "ʷet", "oẻt": "ʷet", "oẽt": "ʷet", "oẹt": "ʷet",
}

// CusOnoffglides maps Vietnamese on-offglides to their IPA equivalents
var CusOnoffglides = map[string]string{
	"oe": "ɛj", "oé": "ɛj", "oè": "ɛj", "oẻ": "ɛj", "oẽ": "ɛj", "oẹ": "ɛj",
	"oai": "aj", "oái": "aj", "oài": "aj", "oải": "aj", "oãi": "aj", "oại": "aj",
	"oay": "ăj", "oáy": "ăj", "oày": "ăj", "oảy": "ăj", "oãy": "ăj", "oạy": "ăj",
	"oao": "aw", "oáo": "aw", "oào": "aw", "oảo": "aw", "oão": "aw", "oạo": "aw",
	"oeo": "ew", "oéo": "ew", "oèo": "ew", "oẻo": "ew", "oẽo": "ew", "oẹo": "ew",
	"óeo": "ew", "òeo": "ew", "ỏeo": "ew", "õeo": "ew", "ọeo": "ew",
	"ueo": "ew", "uéo": "ew", "uèo": "ew", "uẻo": "ew", "uẽo": "ew", "uẹo": "ew",
	"uai": "aj", "uái": "aj", "uài": "aj", "uải": "aj", "uãi": "aj", "uại": "aj",
	"uay": "ăj", "uáy": "ăj", "uày": "ăj", "uảy": "ăj", "uãy": "ăj", "uạy": "ăj",
	"uây": "ɤ̆j", "uấy": "ɤ̆j", "uầy": "ɤ̆j", "uẩy": "ɤ̆j", "uẫy": "ɤ̆j", "uậy": "ɤ̆j",
}

// CusCodas maps Vietnamese codas to their IPA equivalents
var CusCodasMap = map[string]string{
	"p": "p", "t": "t", "c": "k", "m": "m", "n": "n", "ng": "ŋ", "nh": "ŋ", "ch": "k",
}

var SpecialRhyme = []string{"anh", "ành", "ãnh", "ảnh", "ánh", "ạnh", "ách", "ạch"}

// CusTonesP maps Vietnamese tones to their numerical representations
var CusTonesP = map[string]int{
	"á": 5, "à": 2, "ả": 4, "ã": 3, "ạ": 6,
	"ấ": 5, "ầ": 2, "ẩ": 4, "ẫ": 3, "ậ": 6,
	"ắ": 5, "ằ": 2, "ẳ": 4, "ẵ": 3, "ặ": 6,
	"é": 5, "è": 2, "ẻ": 4, "ẽ": 3, "ẹ": 6,
	"ế": 5, "ề": 2, "ể": 4, "ễ": 3, "ệ": 6,
	"í": 5, "ì": 2, "ỉ": 4, "ĩ": 3, "ị": 6,
	"ó": 5, "ò": 2, "ỏ": 4, "õ": 3, "ọ": 6,
	"ố": 5, "ồ": 2, "ổ": 4, "ỗ": 3, "ộ": 6,
	"ớ": 5, "ờ": 2, "ở": 4, "ỡ": 3, "ợ": 6,
	"ú": 5, "ù": 2, "ủ": 4, "ũ": 3, "ụ": 6,
	"ứ": 5, "ừ": 2, "ử": 4, "ữ": 3, "ự": 6,
	"ý": 5, "ỳ": 2, "ỷ": 4, "ỹ": 3, "ỵ": 6,
}

// RimeTone represents all possible rime tones in Vietnamese
var RimeTone = []string{
	"a", "ă", "â", "e", "ê", "i", "o", "ô", "ơ", "u", "ư", "y", "iê", "oa", "oă", "oe", "oo", "uâ", "uê", "uô", "uơ", "uy", "ươ", "uyê", "yê", // blank
	"á", "ắ", "ấ", "é", "ế", "í", "ó", "ố", "ớ", "ú", "ứ", "ý", "iế", "óa", "oắ", "óe", "oó", "uấ", "uế", "uố", "ướ", "úy", "ướ", "uyế", "yế", // grave
	"oá", "oé", "óo", "uý",
	"à", "ằ", "ầ", "è", "ề", "ì", "ò", "ồ", "ờ", "ù", "ừ", "ỳ", "iề", "òa", "oằ", "òe", "oò", "uầ", "uề", "uồ", "ườ", "ùy", "ườ", "uyề", "yề", // acute
	"oà", "oè", "òo", "uỳ",
	"ả", "ẳ", "ẩ", "ẻ", "ể", "ỉ", "ỏ", "ổ", "ở", "ủ", "ử", "ỷ", "iểu", "ỏa", "oẳ", "ỏe", "oỏ", "uẩ", "uể", "uổ", "ưở", "ủy", "ưở", "uyể", "yể", // hook
	"oả", "oẻ", "ỏo", "uỷ",
	"ã", "ẵ", "ẫ", "ẽ", "ễ", "ĩ", "õ", "ỗ", "ỡ", "ũ", "ữ", "ỹ", "iễ", "õa", "oẵ", "õe", "oõ", "uẫ", "uễ", "uỗ", "ưỡ", "ũy", "ưỡ", "uyễ", "yễ", // tilde
	"oã", "oẽ", "õo", "uỹ",
	"ạ", "ặ", "ậ", "ẹ", "ệ", "ị", "ọ", "ộ", "ợ", "ụ", "ự", "ỵ", "iệ", "ọa", "oặ", "ọe", "oọ", "uậ", "uệ", "uệ", "ượ", "ụy", "ượ", "uyệ", "yệ", // dot
	"oạ", "oẹ", "ọo", "uỵ",
}

// Based on: https://github.com/v-nhandt21/Viphoneme/blob/master/viphoneme/T2IPA.py
