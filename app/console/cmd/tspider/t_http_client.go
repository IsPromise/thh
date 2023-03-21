package tspider

import (
	"fmt"
	"thh/arms"
	"thh/bundles/config"

	"github.com/go-resty/resty/v2"
)

type tClient struct {
	httpClient *resty.Client
}

// https://wejson.cn/header2json/
var headersMap = map[string]string{}

func newTClient() tClient {
	client := resty.New()
	// Setting a Proxy URL and Port
	client.SetProxy(config.GetString("T_PROXY"))
	client.SetBaseURL("https://api.twitter.com/")
	client.SetHeaders(headersMap)
	//client.SetAuthToken("")
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
		"variables": arms.JsonEncode(variables{
			UserId:                      userId,
			Count:                       count,
			Cursor:                      cursor,
			IncludePromotedContent:      false,
			WithSuperFollowsUserFields:  true,
			WithDownvotePerspective:     true,
			WithReactionsMetadata:       false,
			WithReactionsPerspective:    false,
			WithSuperFollowsTweetFields: true,
		}),
		"features": arms.JsonEncode(map[string]any{
			"responsive_web_twitter_blue_verified_badge_is_enabled":             true,
			"responsive_web_graphql_exclude_directive_enabled":                  false,
			"verified_phone_label_enabled":                                      false,
			"responsive_web_graphql_timeline_navigation_enabled":                true,
			"responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
			"tweetypie_unmention_optimization_enabled":                          true,
			"vibe_api_enabled":                                                        true,
			"responsive_web_edit_tweet_api_enabled":                                   true,
			"graphql_is_translatable_rweb_tweet_is_translatable_enabled":              true,
			"view_counts_everywhere_api_enabled":                                      true,
			"longform_notetweets_consumption_enabled":                                 true,
			"tweet_awards_web_tipping_enabled":                                        false,
			"freedom_of_speech_not_reach_fetch_enabled":                               false,
			"standardized_nudges_misinfo":                                             true,
			"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": false,
			"interactive_text_enabled":                                                true,
			"responsive_web_text_conversations_enabled":                               false,
			"responsive_web_enhance_cards_enabled":                                    false,
		}),
	}).Get("/graphql/fzE3zNMTkr-CJufrDwjC4A/Following")
}

// getTList 远程方法
func (itself tClient) getTList(userId string, count int, cursor string) (*resty.Response, error) {
	type variables struct {
		UserID                                 string `json:"userId"`
		Count                                  int    `json:"count"`
		Cursor                                 string `json:"cursor,omitempty"`
		IncludePromotedContent                 bool   `json:"includePromotedContent"`
		WithQuickPromoteEligibilityTweetFields bool   `json:"withQuickPromoteEligibilityTweetFields"`
		WithSuperFollowsUserFields             bool   `json:"withSuperFollowsUserFields"`
		WithDownvotePerspective                bool   `json:"withDownvotePerspective"`
		WithReactionsMetadata                  bool   `json:"withReactionsMetadata"`
		WithReactionsPerspective               bool   `json:"withReactionsPerspective"`
		WithSuperFollowsTweetFields            bool   `json:"withSuperFollowsTweetFields"`
		WithVoice                              bool   `json:"withVoice"`
		WithV2Timeline                         bool   `json:"withV2Timeline"`
	}
	return itself.httpClient.R().SetQueryParams(map[string]string{
		"variables": arms.JsonEncode(variables{
			UserID:                                 userId,
			Count:                                  count,
			Cursor:                                 cursor,
			IncludePromotedContent:                 true,
			WithQuickPromoteEligibilityTweetFields: true,
			WithSuperFollowsUserFields:             true,
			WithDownvotePerspective:                true,
			WithReactionsMetadata:                  false,
			WithReactionsPerspective:               false,
			WithSuperFollowsTweetFields:            true,
			WithVoice:                              true,
			WithV2Timeline:                         true,
		}),
		"features": arms.JsonEncode(map[string]any{
			"responsive_web_twitter_blue_verified_badge_is_enabled":             true,
			"responsive_web_graphql_exclude_directive_enabled":                  false,
			"verified_phone_label_enabled":                                      false,
			"responsive_web_graphql_timeline_navigation_enabled":                true,
			"responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
			"tweetypie_unmention_optimization_enabled":                          true,
			"vibe_api_enabled":                                                        true,
			"responsive_web_edit_tweet_api_enabled":                                   true,
			"graphql_is_translatable_rweb_tweet_is_translatable_enabled":              true,
			"view_counts_everywhere_api_enabled":                                      true,
			"longform_notetweets_consumption_enabled":                                 true,
			"tweet_awards_web_tipping_enabled":                                        false,
			"freedom_of_speech_not_reach_fetch_enabled":                               false,
			"standardized_nudges_misinfo":                                             true,
			"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": false,
			"interactive_text_enabled":                                                true,
			"responsive_web_text_conversations_enabled":                               false,
			"responsive_web_enhance_cards_enabled":                                    false,
		}),
	}).Get("/graphql/OXXUyHfKYZ-xLx4NcL9-_Q/UserTweets")
}

// getUserInfo 远程方法
func (itself tClient) getUserInfo(ScreenName string) (*resty.Response, error) {
	type variables struct {
		ScreenName                 string `json:"screen_name"`
		WithSafetyModeUserFields   bool   `json:"withSafetyModeUserFields"`
		WithSuperFollowsUserFields bool   `json:"withSuperFollowsUserFields"`
	}
	type features struct {
		ResponsiveWebTwitterBlueVerifiedBadgeIsEnabled            bool `json:"responsive_web_twitter_blue_verified_badge_is_enabled"`
		ResponsiveWebGraphqlExcludeDirectiveEnabled               bool `json:"responsive_web_graphql_exclude_directive_enabled"`
		VerifiedPhoneLabelEnabled                                 bool `json:"verified_phone_label_enabled"`
		ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled bool `json:"responsive_web_graphql_skip_user_profile_image_extensions_enabled"`
		ResponsiveWebGraphqlTimelineNavigationEnabled             bool `json:"responsive_web_graphql_timeline_navigation_enabled"`
	}

	return itself.httpClient.R().SetQueryParams(map[string]string{
		"variables": arms.JsonEncode(variables{
			ScreenName:                 ScreenName,
			WithSafetyModeUserFields:   true,
			WithSuperFollowsUserFields: true,
		}),
		"features": arms.JsonEncode(features{
			ResponsiveWebTwitterBlueVerifiedBadgeIsEnabled: true,
			VerifiedPhoneLabelEnabled:                      false,
			ResponsiveWebGraphqlTimelineNavigationEnabled:  true,
		}),
	}).Get("/graphql/rePnxwe9LZ51nQ7Sn_xN_A/UserByScreenName")
}

type toolClient struct {
	client *resty.Client
}

var stdToolClient toolClient

func newToolClient() toolClient {
	client := resty.New()
	// Setting a Proxy URL and Port
	proxyPath := config.GetString("T_PROXY")
	client.SetProxy(proxyPath)
	//client.SetOutputDirectory("")
	return toolClient{client: client}
}

func (itself toolClient) downMedia(url string, filename string) {
	downMedia := config.GetBool("T_DOWNMEDIA", false)
	if !downMedia {
		return
	}

	_, err := itself.client.R().SetOutput(filename).Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(url, "下载成功", filename)
}
