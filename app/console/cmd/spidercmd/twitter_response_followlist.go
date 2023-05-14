package spidercmd

type TFollowList struct {
	Data struct {
		User struct {
			Result struct {
				Typename string `json:"__typename"`
				Timeline struct {
					Timeline struct {
						Instructions []struct {
							Type    string `json:"type"`
							Entries []struct {
								EntryId   string `json:"entryId"`
								SortIndex string `json:"sortIndex"`
								Content   struct {
									EntryType   string `json:"entryType"`
									ItemContent struct {
										ItemType    string `json:"itemType"`
										UserResults struct {
											Result struct {
												Typename                   string `json:"__typename"`
												Id                         string `json:"id"`
												RestId                     string `json:"rest_id"`
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
													} `json:"entities"`
													FastFollowersCount     int      `json:"fast_followers_count"`
													FavouritesCount        int      `json:"favourites_count"`
													FollowRequestSent      bool     `json:"follow_request_sent"`
													FollowedBy             bool     `json:"followed_by"`
													FollowersCount         int      `json:"followers_count"`
													Following              bool     `json:"following"`
													FriendsCount           int      `json:"friends_count"`
													HasCustomTimelines     bool     `json:"has_custom_timelines"`
													IsTranslator           bool     `json:"is_translator"`
													ListedCount            int      `json:"listed_count"`
													Location               string   `json:"location"`
													MediaCount             int      `json:"media_count"`
													Muting                 bool     `json:"muting"`
													Name                   string   `json:"name"`
													NormalFollowersCount   int      `json:"normal_followers_count"`
													Notifications          bool     `json:"notifications"`
													PinnedTweetIdsStr      []string `json:"pinned_tweet_ids_str"`
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
													ProfileImageUrlHttps    string        `json:"profile_image_url_https"`
													ProfileInterstitialType string        `json:"profile_interstitial_type"`
													Protected               bool          `json:"protected"`
													ScreenName              string        `json:"screen_name"`
													StatusesCount           int           `json:"statuses_count"`
													TranslatorType          string        `json:"translator_type"`
													Verified                bool          `json:"verified"`
													WantRetweets            bool          `json:"want_retweets"`
													WithheldInCountries     []interface{} `json:"withheld_in_countries"`
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
													} `json:"profile_banner_extensions,omitempty"`
													ProfileBannerUrl string `json:"profile_banner_url,omitempty"`
												} `json:"legacy"`
												SuperFollowEligible bool `json:"super_follow_eligible"`
												SuperFollowedBy     bool `json:"super_followed_by"`
												SuperFollowing      bool `json:"super_following"`
												Professional        struct {
													RestId           string `json:"rest_id"`
													ProfessionalType string `json:"professional_type"`
													Category         []struct {
														Id   int    `json:"id"`
														Name string `json:"name"`
													} `json:"category"`
												} `json:"professional,omitempty"`
											} `json:"result"`
										} `json:"user_results"`
										UserDisplayType string `json:"userDisplayType"`
									} `json:"itemContent,omitempty"`
									ClientEventInfo struct {
										Component string `json:"component"`
										Element   string `json:"element"`
									} `json:"clientEventInfo,omitempty"`
									Value      string `json:"value,omitempty"`
									CursorType string `json:"cursorType,omitempty"`
								} `json:"content"`
							} `json:"entries"`
						} `json:"instructions"`
					} `json:"timeline"`
				} `json:"timeline"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}
