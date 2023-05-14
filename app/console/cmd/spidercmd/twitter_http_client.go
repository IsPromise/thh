package spidercmd

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/leancodebox/goose/jsonopt"
)

type tClient struct {
	httpClient *resty.Client
}

// https://wejson.cn/header2json/
var headersMap = map[string]string{}

func parseHeaders(headers string) map[string]string {
	headerMap := make(map[string]string)
	splitHeaders := strings.Split(headers, "\n")

	for _, headerLine := range splitHeaders {
		headerLine = strings.TrimRight(headerLine, "\r")
		keyVal := strings.SplitN(headerLine, ":", 2)

		if len(keyVal) == 2 {
			headerMap[strings.TrimSpace(keyVal[0])] = strings.TrimSpace(keyVal[1])
		}
	}

	return headerMap
}

func newTClient() tClient {
	client := resty.New()
	// Setting a Proxy URL and Port
	if useProxy {
		client.SetProxy(proxy)
	}
	client.SetBaseURL("https://twitter.com/")
	client.SetHeaders(getHeader())
	return tClient{client}
}

// getFollowList 远程方法
func (itself tClient) getFollowList(userId string, count int, cursor string) (*resty.Response, error) {
	type variables struct {
		UserId                      string `json:"userId"`
		Count                       int    `json:"count"`
		Cursor                      string `json:"cursor,omitempty"`
		IncludePromotedContent      bool   `json:"includePromotedContent"`
		WithSuperFollowsUserFields  bool   `json:"withSuperFollowsUserFields"`
		WithDownvotePerspective     bool   `json:"withDownvotePerspective"`
		WithReactionsMetadata       bool   `json:"withReactionsMetadata"`
		WithReactionsPerspective    bool   `json:"withReactionsPerspective"`
		WithSuperFollowsTweetFields bool   `json:"withSuperFollowsTweetFields"`
	}

	return itself.httpClient.R().SetQueryParams(map[string]string{
		"variables": jsonopt.Encode(variables{
			UserId:                 userId,
			Count:                  count,
			Cursor:                 cursor,
			IncludePromotedContent: false,
		}),
		"features": jsonopt.Encode(map[string]any{
			"rweb_lists_timeline_redesign_enabled":                              false,
			"blue_business_profile_image_shape_enabled":                         true,
			"responsive_web_graphql_exclude_directive_enabled":                  true,
			"verified_phone_label_enabled":                                      false,
			"creator_subscriptions_tweet_preview_api_enabled":                   false,
			"responsive_web_graphql_timeline_navigation_enabled":                true,
			"responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
			"tweetypie_unmention_optimization_enabled":                          true,
			"vibe_api_enabled":                                                        true,
			"responsive_web_edit_tweet_api_enabled":                                   true,
			"graphql_is_translatable_rweb_tweet_is_translatable_enabled":              true,
			"view_counts_everywhere_api_enabled":                                      true,
			"longform_notetweets_consumption_enabled":                                 true,
			"tweet_awards_web_tipping_enabled":                                        false,
			"freedom_of_speech_not_reach_fetch_enabled":                               true,
			"standardized_nudges_misinfo":                                             true,
			"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": false,
			"interactive_text_enabled":                                                true,
			"responsive_web_text_conversations_enabled":                               false,
			"longform_notetweets_rich_text_read_enabled":                              true,
			"longform_notetweets_inline_media_enabled":                                false,
			"responsive_web_enhance_cards_enabled":                                    false,
		}),
	}).Get("i/api/graphql/q4cKckK0lNxWkHfAXXXzJQ/Following")
}

// getTList 远程方法
func (itself tClient) getTList(userId string, count int, cursor string) (*resty.Response, error) {
	type variables struct {
		UserId                                 string `json:"userId"`
		Count                                  int    `json:"count"`
		Cursor                                 string `json:"cursor"`
		IncludePromotedContent                 bool   `json:"includePromotedContent"`
		WithQuickPromoteEligibilityTweetFields bool   `json:"withQuickPromoteEligibilityTweetFields"`
		WithVoice                              bool   `json:"withVoice"`
		WithV2Timeline                         bool   `json:"withV2Timeline"`
	}

	return itself.httpClient.R().SetQueryParams(map[string]string{
		"variables": jsonopt.Encode(variables{
			UserId:                                 userId,
			Count:                                  count,
			Cursor:                                 cursor,
			IncludePromotedContent:                 true,
			WithQuickPromoteEligibilityTweetFields: true,
			WithVoice:                              true,
			WithV2Timeline:                         true,
		}),
		"features": jsonopt.Encode(map[string]any{
			"rweb_lists_timeline_redesign_enabled":                              false,
			"blue_business_profile_image_shape_enabled":                         true,
			"responsive_web_graphql_exclude_directive_enabled":                  true,
			"verified_phone_label_enabled":                                      false,
			"creator_subscriptions_tweet_preview_api_enabled":                   false,
			"responsive_web_graphql_timeline_navigation_enabled":                true,
			"responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
			"tweetypie_unmention_optimization_enabled":                          true,
			"vibe_api_enabled":                                                        true,
			"responsive_web_edit_tweet_api_enabled":                                   true,
			"graphql_is_translatable_rweb_tweet_is_translatable_enabled":              true,
			"view_counts_everywhere_api_enabled":                                      true,
			"longform_notetweets_consumption_enabled":                                 true,
			"tweet_awards_web_tipping_enabled":                                        false,
			"freedom_of_speech_not_reach_fetch_enabled":                               true,
			"standardized_nudges_misinfo":                                             true,
			"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": false,
			"interactive_text_enabled":                                                true,
			"responsive_web_text_conversations_enabled":                               false,
			"longform_notetweets_rich_text_read_enabled":                              true,
			"longform_notetweets_inline_media_enabled":                                false,
			"responsive_web_enhance_cards_enabled":                                    false,
		}),
	}).Get("i/api/graphql/WzJjibAcDa-oCjCcLOotcg/UserTweets")
}

// getUserInfo 远程方法
func (itself tClient) getUserInfo(ScreenName string) (*resty.Response, error) {
	return itself.httpClient.R().SetQueryParams(map[string]string{
		"variables": jsonopt.Encode(map[string]any{
			"screen_name":              ScreenName,
			"withSafetyModeUserFields": true,
		}),
		"features": jsonopt.Encode(map[string]any{
			"blue_business_profile_image_shape_enabled":                         true,
			"responsive_web_graphql_exclude_directive_enabled":                  true,
			"verified_phone_label_enabled":                                      false,
			"highlights_tweets_tab_ui_enabled":                                  true,
			"creator_subscriptions_tweet_preview_api_enabled":                   false,
			"responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
			"responsive_web_graphql_timeline_navigation_enabled":                true,
		}),
	}).Get("i/api/graphql/pVrmNaXcxPjisIvKtLDMEA/UserByScreenName")
}

type toolClient struct {
	client *resty.Client
}

var stdToolClient toolClient

func newToolClient() toolClient {
	client := resty.New()
	// Setting a Proxy URL and Port
	if useProxy {
		client.SetProxy(proxy)
	}
	return toolClient{client: client}
}

func (itself toolClient) downMedia(url string, filename string) {
	if !needDownMedia() {
		return
	}

	_, err := itself.client.R().SetOutput(filename).Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(url, "下载成功", filename)
}
