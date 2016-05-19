package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBotEventHandler ...
type LineBotEventHandler struct {
	botClient *linebot.Client
}

// SetLineBotClient to assign linebot client handler
func (lbe *LineBotEventHandler) SetLineBotClient(bc *linebot.Client) {
	lbe.botClient = bc
}

// OnAddedAsFriendOperation ...
func (lbe *LineBotEventHandler) OnAddedAsFriendOperation(mids []string) {
	lbe.botClient.SendText(mids, "感謝你加入....！")
}

// OnBlockedAccountOperation ...
func (lbe *LineBotEventHandler) OnBlockedAccountOperation(mids []string) {
	lbe.botClient.SendText(mids, "被封鎖了")
}

func (lbe *LineBotEventHandler) matchString(pattern, s string) (result bool) {
	result, _ = regexp.MatchString(pattern, s)
	return
}

// OnTextMessage ...
func (lbe *LineBotEventHandler) OnTextMessage(from, text string) {
	strAfterCut := strings.ToUpper(text)
	strResult := "目前系統功能如下\n" +
		"利率(預設為台幣),外幣利率\n" +
		"匯率(預設為總表),美元,日圓與人民幣匯率\n" +
		"歷史匯率\n" +
		"台外幣各類存款餘額"

	switch {
	case lbe.matchString("外幣+.*利率+.*", strAfterCut):
		strResult = "常用外幣利率表\n 美元 定存 2.3% 活存 1.8% \n 日圓 定存 0.1% 活存 0.1%"
		break
	case lbe.matchString("台幣+.*利率+.*", strAfterCut):
	case lbe.matchString("利率.*", strAfterCut):
		strResult = "台幣活存利率表 \n 活存 0.5% 活儲 0.6% \n 定存\n 三個月 0.76% 六個月 0.78% 一年 0.80% 三年 0.80%\n https://www.skbank.com.tw/RAT/RAT2_TWSaving.aspx"
		break
	case lbe.matchString("(美元|美金|USD)+.*歷史+.*匯率+.*", strAfterCut):
		strResult = "美元歷史匯率參考 http://tw.exchange-rates.org/history/TWD/USD/T"
		break
	case lbe.matchString("(日圓|日元|日幣|JPY)+.*歷史+.*匯率+.*", strAfterCut):
		strResult = "日元歷史匯率參考 http://tw.exchange-rates.org/history/TWD/JPY/T"
		break
	case lbe.matchString("(人民幣|RMB)+.*歷史+.*匯率+.*", strAfterCut):
		strResult = "人民幣歷史匯率參考 http://tw.exchange-rates.org/history/TWD/CNY/T"
		break
	case lbe.matchString("歷史+.*匯率+.*", strAfterCut):
		strResult = "歷史匯率參考 http://tw.exchange-rates.org/history/TWD/USD/T"
		break
	case lbe.matchString("(美元|美金|USD)+.*匯率+.*", strAfterCut):
		strResult = "美元匯率\n" +
			"現金買入 32.32000\n" +
			"現金賣出 32.86200\n" +
			"即期買入 32.62000\n" +
			"即期賣出 32.72000"
		break
	case lbe.matchString("(日圓|日元|日幣|JPY)+.*匯率+.*", strAfterCut):
		strResult = "日圓匯率\n" +
			"現金買入 0.29160\n" +
			"現金賣出 0.30260\n" +
			"即期買入 0.29800\n" +
			"即期賣出 0.30200"
		break
	case lbe.matchString("(人民幣|RMB)+.*匯率+.*", strAfterCut):
		strResult = "人民幣匯率\n" +
			"現金買入 4.89000\n" +
			"現金賣出 5.05200\n" +
			"即期買入 4.96200\n" +
			"即期賣出 5.01200"
		break
	case lbe.matchString("匯率+.*", strAfterCut):
		strResult = ""
		lbe.botClient.SendImage([]string{from},
			"https://linebot.gaze.tw/exrate.png",
			"https://linebot.gaze.tw/exrate.png")
		break

	case lbe.matchString("(美元|美金|USD)+.*(活存|存款)+.*(餘額)?.*", strAfterCut):
		strResult = "您的美元活存帳戶餘額為: 233,188.66 美元"
		break
	case lbe.matchString("(日圓|日元|日幣|JPY)+.*(活存|存款)+.*(餘額)?.*", strAfterCut):
		strResult = "您的日元活存帳戶餘額為: 233,188.66 日元"
		break
	case lbe.matchString("(人民幣|RMB)+.*(活存|存款)+.*餘額?.*", strAfterCut):
		strResult = "您沒有人民幣帳戶，若要開立請點連結 https://virtual.bank"
		break
	case lbe.matchString("(美元|美金|USD)+.*(定存|存單)+.*(餘額)?.*", strAfterCut):
		strResult = "您的美元定存帳戶餘額為: 1,000.00 美元"
		break
	case lbe.matchString("(日圓|日元|日幣|JPY)+.*(定存|存單)+.*(餘額)?.*", strAfterCut):
		strResult = "您沒有日元定存帳戶，若要開立請點連結 https://virtual.bank"
		break
	case lbe.matchString("(人民幣|RMB)+.*(定存|存單)+.*(餘額)?.*", strAfterCut):
		strResult = "您沒有人民幣帳戶，若要開立請點連結 https://virtual.bank"
		break
	case lbe.matchString("(台幣)?.*(存款|活存|帳戶)+.*(餘額)?.*", strAfterCut):
		strResult = "您的台幣活存帳戶餘額為: 233,188.66 元\n "
		break
	case lbe.matchString("(台幣)?.*(定存|存單)+.*(餘額)?.*", strAfterCut):
		strResult = "您的台幣定存帳戶餘額為: 1,000,000.00 元\n"
		break
	case lbe.matchString("餘額.*", strAfterCut):
		strResult = "您的台幣活存帳戶餘額為: 233,188.66 元\n "
		break

	}
	if strResult != "" {
		lbe.botClient.SendText([]string{from}, strResult)
	}
	log.Printf("Received text \"%s\" from %s", text, from)
}

// OnImageMessage ...
func (lbe *LineBotEventHandler) OnImageMessage(from string, rc *linebot.ReceivedContent) {
	lbe.botClient.SendText([]string{from}, "收到一張照片")
	log.Print("=== Received Image ===")
}

// OnVideoMessage ...
func (lbe *LineBotEventHandler) OnVideoMessage(from string, rc *linebot.ReceivedContent) {
	lbe.botClient.SendText([]string{from}, "收到一段錄影")
	log.Print("=== Received Video ===")
}

// OnAudioMessage ...
func (lbe *LineBotEventHandler) OnAudioMessage(from string, duration int) {
	lbe.botClient.SendText([]string{from}, "收到一段錄音")
	log.Print("=== Received Audio ===")
}

// OnLocationMessage ...
func (lbe *LineBotEventHandler) OnLocationMessage(from, title, address string, latitude, longitude float64) {
	lbe.botClient.SendText([]string{from}, "收到地點資訊")
	log.Print("=== Received Location ===")
}

// OnStickerMessage ...
func (lbe *LineBotEventHandler) OnStickerMessage(from string, stickerPackageID, stickerID, stickerVersion int) {
	lbe.botClient.SendText([]string{from}, "收到一張貼紙")
	log.Print("=== Received Sticker ===")
}

// OnContactMessage ...
func (lbe *LineBotEventHandler) OnContactMessage(from, MID, displayName string) {
	lbe.botClient.SendText([]string{from}, "收到聯絡人資料")
	log.Print("=== Received Contact ===")
}