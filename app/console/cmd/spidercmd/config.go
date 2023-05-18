package spidercmd

import (
	"github.com/leancodebox/goose/preferences"
	"path/filepath"
	"time"
)

const (
	QueueKey             = "twitter:screenName:list"
)

var (
	useProxy             = preferences.GetBool("spider.twitter.useProxy", true)
	proxy                = preferences.GetString("spider.twitter.proxy")
	rootPrefix           = preferences.GetString("spider.twitter.output", "./storage/tmp/")
	outputPrefix         = filepath.Join(rootPrefix, time.Now().Format("20060102_150405"))
	spiderDeep           = preferences.GetInt("spider.twitter.deep", 0)
	allUsePush           = preferences.GetBool("spider.twitter.allusepush", false)
	spiderTwitterMaxPage = preferences.GetInt("spider.twitter.maxPage", "")
)

func getScreenNameSlice() []string {
	return preferences.GetStringSlice("spider.twitter.screenNameList")
}

func needDownMedia() bool {
	return preferences.GetBool("spider.twitter.downmedia", false)
}

func getHeader() map[string]string {
	return parseHeaders(preferences.Get("spider.twitter.header"))
}
