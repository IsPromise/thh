package spidercmd

type TUserInfo struct {
	Data struct {
		User struct {
			Result struct {
				Typename                   string `json:"__typename"`
				ID                         string `json:"id"`
				RestID                     string `json:"rest_id"`
				AffiliatesHighlightedLabel struct {
				} `json:"affiliates_highlighted_label"`
				HasNftAvatar bool `json:"has_nft_avatar"`
				Legacy       struct {
					BlockedBy           bool   `json:"blocked_by"`
					Blocking            bool   `json:"blocking"`
					CanDm               bool   `json:"can_dm"`
					CanMediaTag         bool   `json:"can_media_tag"`
					CreatedAt           string `json:"created_at"`
					DefaultProfile      bool   `json:"default_profile"`
					DefaultProfileImage bool   `json:"default_profile_image"`
					Description         string `json:"description"`
					Entities            struct {
						Description struct {
							Urls []interface{} `json:"urls"`
						} `json:"description"`
						URL struct {
							Urls []struct {
								DisplayURL  string `json:"display_url"`
								ExpandedURL string `json:"expanded_url"`
								URL         string `json:"url"`
								Indices     []int  `json:"indices"`
							} `json:"urls"`
						} `json:"url"`
					} `json:"entities"`
					FastFollowersCount      int      `json:"fast_followers_count"`
					FavouritesCount         int      `json:"favourites_count"`
					FollowRequestSent       bool     `json:"follow_request_sent"`
					FollowedBy              bool     `json:"followed_by"`
					FollowersCount          int      `json:"followers_count"`
					Following               bool     `json:"following"`
					FriendsCount            int      `json:"friends_count"`
					HasCustomTimelines      bool     `json:"has_custom_timelines"`
					IsTranslator            bool     `json:"is_translator"`
					ListedCount             int      `json:"listed_count"`
					Location                string   `json:"location"`
					MediaCount              int      `json:"media_count"`
					Muting                  bool     `json:"muting"`
					Name                    string   `json:"name"`
					NormalFollowersCount    int      `json:"normal_followers_count"`
					Notifications           bool     `json:"notifications"`
					PinnedTweetIdsStr       []string `json:"pinned_tweet_ids_str"`
					ProfileBannerExtensions struct {
						MediaColor struct {
							R struct {
								Ok struct {
									Palette []struct {
										Percentage float64 `json:"percentage"`
										Rgb        struct {
											Blue  int `json:"blue"`
											Green int `json:"green"`
											Red   int `json:"red"`
										} `json:"rgb"`
									} `json:"palette"`
								} `json:"ok"`
							} `json:"r"`
						} `json:"mediaColor"`
					} `json:"profile_banner_extensions"`
					ProfileBannerURL       string `json:"profile_banner_url"`
					ProfileImageExtensions struct {
						MediaColor struct {
							R struct {
								Ok struct {
									Palette []struct {
										Percentage float64 `json:"percentage"`
										Rgb        struct {
											Blue  int `json:"blue"`
											Green int `json:"green"`
											Red   int `json:"red"`
										} `json:"rgb"`
									} `json:"palette"`
								} `json:"ok"`
							} `json:"r"`
						} `json:"mediaColor"`
					} `json:"profile_image_extensions"`
					ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
					ProfileInterstitialType string        `json:"profile_interstitial_type"`
					Protected               bool          `json:"protected"`
					ScreenName              string        `json:"screen_name"`
					StatusesCount           int           `json:"statuses_count"`
					TranslatorType          string        `json:"translator_type"`
					URL                     string        `json:"url"`
					Verified                bool          `json:"verified"`
					WantRetweets            bool          `json:"want_retweets"`
					WithheldInCountries     []interface{} `json:"withheld_in_countries"`
				} `json:"legacy"`
				SmartBlockedBy        bool `json:"smart_blocked_by"`
				SmartBlocking         bool `json:"smart_blocking"`
				SuperFollowEligible   bool `json:"super_follow_eligible"`
				SuperFollowedBy       bool `json:"super_followed_by"`
				SuperFollowing        bool `json:"super_following"`
				LegacyExtendedProfile struct {
				} `json:"legacy_extended_profile"`
				IsProfileTranslatable bool `json:"is_profile_translatable"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}
