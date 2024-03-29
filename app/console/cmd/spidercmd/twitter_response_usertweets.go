package spidercmd

type UserTweetsResponse struct {
	Data struct {
		User struct {
			Result struct {
				Typename   string `json:"__typename"`
				TimelineV2 struct {
					Timeline struct {
						Instructions []struct {
							Type    string `json:"type"`
							Entries []struct {
								EntryId   string `json:"entryId"`
								SortIndex string `json:"sortIndex"`
								Content   struct {
									EntryType   string `json:"entryType"`
									ItemContent struct {
										ItemType     string `json:"itemType"`
										TweetResults struct {
											Result struct {
												Typename string `json:"__typename"`
												RestId   string `json:"rest_id"`
												Core     struct {
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
																	Url struct {
																		Urls []struct {
																			DisplayUrl  string `json:"display_url"`
																			ExpandedUrl string `json:"expanded_url"`
																			Url         string `json:"url"`
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
																ProfileBannerUrl       string `json:"profile_banner_url"`
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
																Url                     string        `json:"url"`
																Verified                bool          `json:"verified"`
																WantRetweets            bool          `json:"want_retweets"`
																WithheldInCountries     []interface{} `json:"withheld_in_countries"`
															} `json:"legacy"`
															SuperFollowEligible bool `json:"super_follow_eligible"`
															SuperFollowedBy     bool `json:"super_followed_by"`
															SuperFollowing      bool `json:"super_following"`
														} `json:"result"`
													} `json:"user_results"`
												} `json:"core"`
												DownvotePerspective struct {
													IsDownvoted bool `json:"isDownvoted"`
												} `json:"downvotePerspective"`
												UnmentionInfo struct {
												} `json:"unmention_info"`
												Legacy struct {
													CreatedAt         string `json:"created_at"`
													ConversationIdStr string `json:"conversation_id_str"`
													DisplayTextRange  []int  `json:"display_text_range"`
													Entities          struct {
														Media []struct {
															DisplayUrl        string `json:"display_url"`
															ExpandedUrl       string `json:"expanded_url"`
															IdStr             string `json:"id_str"`
															Indices           []int  `json:"indices"`
															MediaUrlHttps     string `json:"media_url_https"`
															SourceStatusIdStr string `json:"source_status_id_str"`
															SourceUserIdStr   string `json:"source_user_id_str"`
															Type              string `json:"type"`
															Url               string `json:"url"`
															Features          struct {
																Large struct {
																	Faces []struct {
																		X int `json:"x"`
																		Y int `json:"y"`
																		H int `json:"h"`
																		W int `json:"w"`
																	} `json:"faces"`
																} `json:"large,omitempty"`
																Medium struct {
																	Faces []struct {
																		X int `json:"x"`
																		Y int `json:"y"`
																		H int `json:"h"`
																		W int `json:"w"`
																	} `json:"faces"`
																} `json:"medium,omitempty"`
																Small struct {
																	Faces []struct {
																		X int `json:"x"`
																		Y int `json:"y"`
																		H int `json:"h"`
																		W int `json:"w"`
																	} `json:"faces"`
																} `json:"small,omitempty"`
																Orig struct {
																	Faces []struct {
																		X int `json:"x"`
																		Y int `json:"y"`
																		H int `json:"h"`
																		W int `json:"w"`
																	} `json:"faces"`
																} `json:"orig,omitempty"`
															} `json:"features"`
															Sizes struct {
																Large struct {
																	H      int    `json:"h"`
																	W      int    `json:"w"`
																	Resize string `json:"resize"`
																} `json:"large"`
																Medium struct {
																	H      int    `json:"h"`
																	W      int    `json:"w"`
																	Resize string `json:"resize"`
																} `json:"medium"`
																Small struct {
																	H      int    `json:"h"`
																	W      int    `json:"w"`
																	Resize string `json:"resize"`
																} `json:"small"`
																Thumb struct {
																	H      int    `json:"h"`
																	W      int    `json:"w"`
																	Resize string `json:"resize"`
																} `json:"thumb"`
															} `json:"sizes"`
															OriginalInfo struct {
																Height     int `json:"height"`
																Width      int `json:"width"`
																FocusRects []struct {
																	X int `json:"x"`
																	Y int `json:"y"`
																	W int `json:"w"`
																	H int `json:"h"`
																} `json:"focus_rects,omitempty"`
															} `json:"original_info"`
														} `json:"media"`
														UserMentions []struct {
															IdStr      string `json:"id_str"`
															Name       string `json:"name"`
															ScreenName string `json:"screen_name"`
															Indices    []int  `json:"indices"`
														} `json:"user_mentions"`
														Urls     []interface{} `json:"urls"`
														Hashtags []struct {
															Indices []int  `json:"indices"`
															Text    string `json:"text"`
														} `json:"hashtags"`
														Symbols []interface{} `json:"symbols"`
													} `json:"entities"`
													ExtendedEntities struct {
														Media []struct {
															DisplayUrl          string `json:"display_url"`
															ExpandedUrl         string `json:"expanded_url"`
															IdStr               string `json:"id_str"`
															Indices             []int  `json:"indices"`
															MediaKey            string `json:"media_key"`
															MediaUrlHttps       string `json:"media_url_https"`
															SourceStatusIdStr   string `json:"source_status_id_str"`
															SourceUserIdStr     string `json:"source_user_id_str"`
															Type                string `json:"type"`
															Url                 string `json:"url"`
															AdditionalMediaInfo struct {
																Monetizable bool `json:"monetizable"`
																SourceUser  struct {
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
																						Urls []struct {
																							DisplayUrl  string `json:"display_url"`
																							ExpandedUrl string `json:"expanded_url"`
																							Url         string `json:"url"`
																							Indices     []int  `json:"indices"`
																						} `json:"urls"`
																					} `json:"description"`
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
																				} `json:"profile_banner_extensions,omitempty"`
																				ProfileBannerUrl       string `json:"profile_banner_url,omitempty"`
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
																} `json:"source_user"`
															} `json:"additional_media_info,omitempty"`
															ExtMediaColor struct {
																Palette []struct {
																	Percentage float64 `json:"percentage"`
																	Rgb        struct {
																		Blue  int `json:"blue"`
																		Green int `json:"green"`
																		Red   int `json:"red"`
																	} `json:"rgb"`
																} `json:"palette"`
															} `json:"ext_media_color"`
															MediaStats struct {
																ViewCount int `json:"viewCount"`
															} `json:"mediaStats,omitempty"`
															ExtMediaAvailability struct {
																Status string `json:"status"`
															} `json:"ext_media_availability"`
															Features struct {
																Large struct {
																	Faces []struct {
																		X int `json:"x"`
																		Y int `json:"y"`
																		H int `json:"h"`
																		W int `json:"w"`
																	} `json:"faces"`
																} `json:"large,omitempty"`
																Medium struct {
																	Faces []struct {
																		X int `json:"x"`
																		Y int `json:"y"`
																		H int `json:"h"`
																		W int `json:"w"`
																	} `json:"faces"`
																} `json:"medium,omitempty"`
																Small struct {
																	Faces []struct {
																		X int `json:"x"`
																		Y int `json:"y"`
																		H int `json:"h"`
																		W int `json:"w"`
																	} `json:"faces"`
																} `json:"small,omitempty"`
																Orig struct {
																	Faces []struct {
																		X int `json:"x"`
																		Y int `json:"y"`
																		H int `json:"h"`
																		W int `json:"w"`
																	} `json:"faces"`
																} `json:"orig,omitempty"`
															} `json:"features"`
															Sizes struct {
																Large struct {
																	H      int    `json:"h"`
																	W      int    `json:"w"`
																	Resize string `json:"resize"`
																} `json:"large"`
																Medium struct {
																	H      int    `json:"h"`
																	W      int    `json:"w"`
																	Resize string `json:"resize"`
																} `json:"medium"`
																Small struct {
																	H      int    `json:"h"`
																	W      int    `json:"w"`
																	Resize string `json:"resize"`
																} `json:"small"`
																Thumb struct {
																	H      int    `json:"h"`
																	W      int    `json:"w"`
																	Resize string `json:"resize"`
																} `json:"thumb"`
															} `json:"sizes"`
															OriginalInfo struct {
																Height     int `json:"height"`
																Width      int `json:"width"`
																FocusRects []struct {
																	X int `json:"x"`
																	Y int `json:"y"`
																	W int `json:"w"`
																	H int `json:"h"`
																} `json:"focus_rects,omitempty"`
															} `json:"original_info"`
															VideoInfo struct {
																AspectRatio    []int `json:"aspect_ratio"`
																DurationMillis int   `json:"duration_millis"`
																Variants       []struct {
																	Bitrate     int    `json:"bitrate,omitempty"`
																	ContentType string `json:"content_type"`
																	Url         string `json:"url"`
																} `json:"variants"`
															} `json:"video_info,omitempty"`
														} `json:"media"`
													} `json:"extended_entities"`
													FavoriteCount             int    `json:"favorite_count"`
													Favorited                 bool   `json:"favorited"`
													FullText                  string `json:"full_text"`
													IsQuoteStatus             bool   `json:"is_quote_status"`
													Lang                      string `json:"lang"`
													PossiblySensitive         bool   `json:"possibly_sensitive"`
													PossiblySensitiveEditable bool   `json:"possibly_sensitive_editable"`
													QuoteCount                int    `json:"quote_count"`
													ReplyCount                int    `json:"reply_count"`
													RetweetCount              int    `json:"retweet_count"`
													Retweeted                 bool   `json:"retweeted"`
													Source                    string `json:"source"`
													UserIdStr                 string `json:"user_id_str"`
													IdStr                     string `json:"id_str"`
													RetweetedStatusResult     struct {
														Result struct {
															Typename string `json:"__typename"`
															RestId   string `json:"rest_id"`
															Core     struct {
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
																					Urls []struct {
																						DisplayUrl  string `json:"display_url"`
																						ExpandedUrl string `json:"expanded_url"`
																						Url         string `json:"url"`
																						Indices     []int  `json:"indices"`
																					} `json:"urls"`
																				} `json:"description"`
																				Url struct {
																					Urls []struct {
																						DisplayUrl  string `json:"display_url"`
																						ExpandedUrl string `json:"expanded_url"`
																						Url         string `json:"url"`
																						Indices     []int  `json:"indices"`
																					} `json:"urls"`
																				} `json:"url,omitempty"`
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
																			} `json:"profile_banner_extensions,omitempty"`
																			ProfileBannerUrl       string `json:"profile_banner_url,omitempty"`
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
																			Url                     string        `json:"url,omitempty"`
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
															} `json:"core"`
															DownvotePerspective struct {
																IsDownvoted bool `json:"isDownvoted"`
															} `json:"downvotePerspective"`
															UnmentionInfo struct {
															} `json:"unmention_info"`
															Legacy struct {
																CreatedAt         string `json:"created_at"`
																ConversationIdStr string `json:"conversation_id_str"`
																DisplayTextRange  []int  `json:"display_text_range"`
																Entities          struct {
																	Media []struct {
																		DisplayUrl    string `json:"display_url"`
																		ExpandedUrl   string `json:"expanded_url"`
																		IdStr         string `json:"id_str"`
																		Indices       []int  `json:"indices"`
																		MediaUrlHttps string `json:"media_url_https"`
																		Type          string `json:"type"`
																		Url           string `json:"url"`
																		Features      struct {
																			Large struct {
																				Faces []struct {
																					X int `json:"x"`
																					Y int `json:"y"`
																					H int `json:"h"`
																					W int `json:"w"`
																				} `json:"faces"`
																			} `json:"large,omitempty"`
																			Medium struct {
																				Faces []struct {
																					X int `json:"x"`
																					Y int `json:"y"`
																					H int `json:"h"`
																					W int `json:"w"`
																				} `json:"faces"`
																			} `json:"medium,omitempty"`
																			Small struct {
																				Faces []struct {
																					X int `json:"x"`
																					Y int `json:"y"`
																					H int `json:"h"`
																					W int `json:"w"`
																				} `json:"faces"`
																			} `json:"small,omitempty"`
																			Orig struct {
																				Faces []struct {
																					X int `json:"x"`
																					Y int `json:"y"`
																					H int `json:"h"`
																					W int `json:"w"`
																				} `json:"faces"`
																			} `json:"orig,omitempty"`
																		} `json:"features"`
																		Sizes struct {
																			Large struct {
																				H      int    `json:"h"`
																				W      int    `json:"w"`
																				Resize string `json:"resize"`
																			} `json:"large"`
																			Medium struct {
																				H      int    `json:"h"`
																				W      int    `json:"w"`
																				Resize string `json:"resize"`
																			} `json:"medium"`
																			Small struct {
																				H      int    `json:"h"`
																				W      int    `json:"w"`
																				Resize string `json:"resize"`
																			} `json:"small"`
																			Thumb struct {
																				H      int    `json:"h"`
																				W      int    `json:"w"`
																				Resize string `json:"resize"`
																			} `json:"thumb"`
																		} `json:"sizes"`
																		OriginalInfo struct {
																			Height     int `json:"height"`
																			Width      int `json:"width"`
																			FocusRects []struct {
																				X int `json:"x"`
																				Y int `json:"y"`
																				W int `json:"w"`
																				H int `json:"h"`
																			} `json:"focus_rects,omitempty"`
																		} `json:"original_info"`
																	} `json:"media"`
																	UserMentions []struct {
																		IdStr      string `json:"id_str"`
																		Name       string `json:"name"`
																		ScreenName string `json:"screen_name"`
																		Indices    []int  `json:"indices"`
																	} `json:"user_mentions"`
																	Urls     []interface{} `json:"urls"`
																	Hashtags []struct {
																		Indices []int  `json:"indices"`
																		Text    string `json:"text"`
																	} `json:"hashtags"`
																	Symbols []interface{} `json:"symbols"`
																} `json:"entities"`
																ExtendedEntities struct {
																	Media []struct {
																		DisplayUrl          string `json:"display_url"`
																		ExpandedUrl         string `json:"expanded_url"`
																		IdStr               string `json:"id_str"`
																		Indices             []int  `json:"indices"`
																		MediaKey            string `json:"media_key"`
																		MediaUrlHttps       string `json:"media_url_https"`
																		Type                string `json:"type"`
																		Url                 string `json:"url"`
																		AdditionalMediaInfo struct {
																			Monetizable bool `json:"monetizable"`
																		} `json:"additional_media_info,omitempty"`
																		ExtMediaColor struct {
																			Palette []struct {
																				Percentage float64 `json:"percentage"`
																				Rgb        struct {
																					Blue  int `json:"blue"`
																					Green int `json:"green"`
																					Red   int `json:"red"`
																				} `json:"rgb"`
																			} `json:"palette"`
																		} `json:"ext_media_color"`
																		MediaStats struct {
																			ViewCount int `json:"viewCount"`
																		} `json:"mediaStats,omitempty"`
																		ExtMediaAvailability struct {
																			Status string `json:"status"`
																		} `json:"ext_media_availability"`
																		Features struct {
																			Large struct {
																				Faces []struct {
																					X int `json:"x"`
																					Y int `json:"y"`
																					H int `json:"h"`
																					W int `json:"w"`
																				} `json:"faces"`
																			} `json:"large,omitempty"`
																			Medium struct {
																				Faces []struct {
																					X int `json:"x"`
																					Y int `json:"y"`
																					H int `json:"h"`
																					W int `json:"w"`
																				} `json:"faces"`
																			} `json:"medium,omitempty"`
																			Small struct {
																				Faces []struct {
																					X int `json:"x"`
																					Y int `json:"y"`
																					H int `json:"h"`
																					W int `json:"w"`
																				} `json:"faces"`
																			} `json:"small,omitempty"`
																			Orig struct {
																				Faces []struct {
																					X int `json:"x"`
																					Y int `json:"y"`
																					H int `json:"h"`
																					W int `json:"w"`
																				} `json:"faces"`
																			} `json:"orig,omitempty"`
																		} `json:"features"`
																		Sizes struct {
																			Large struct {
																				H      int    `json:"h"`
																				W      int    `json:"w"`
																				Resize string `json:"resize"`
																			} `json:"large"`
																			Medium struct {
																				H      int    `json:"h"`
																				W      int    `json:"w"`
																				Resize string `json:"resize"`
																			} `json:"medium"`
																			Small struct {
																				H      int    `json:"h"`
																				W      int    `json:"w"`
																				Resize string `json:"resize"`
																			} `json:"small"`
																			Thumb struct {
																				H      int    `json:"h"`
																				W      int    `json:"w"`
																				Resize string `json:"resize"`
																			} `json:"thumb"`
																		} `json:"sizes"`
																		OriginalInfo struct {
																			Height     int `json:"height"`
																			Width      int `json:"width"`
																			FocusRects []struct {
																				X int `json:"x"`
																				Y int `json:"y"`
																				W int `json:"w"`
																				H int `json:"h"`
																			} `json:"focus_rects,omitempty"`
																		} `json:"original_info"`
																		VideoInfo struct {
																			AspectRatio    []int `json:"aspect_ratio"`
																			DurationMillis int   `json:"duration_millis"`
																			Variants       []struct {
																				Bitrate     int    `json:"bitrate,omitempty"`
																				ContentType string `json:"content_type"`
																				Url         string `json:"url"`
																			} `json:"variants"`
																		} `json:"video_info,omitempty"`
																	} `json:"media"`
																} `json:"extended_entities"`
																FavoriteCount             int    `json:"favorite_count"`
																Favorited                 bool   `json:"favorited"`
																FullText                  string `json:"full_text"`
																IsQuoteStatus             bool   `json:"is_quote_status"`
																Lang                      string `json:"lang"`
																PossiblySensitive         bool   `json:"possibly_sensitive"`
																PossiblySensitiveEditable bool   `json:"possibly_sensitive_editable"`
																QuoteCount                int    `json:"quote_count"`
																ReplyCount                int    `json:"reply_count"`
																RetweetCount              int    `json:"retweet_count"`
																Retweeted                 bool   `json:"retweeted"`
																Source                    string `json:"source"`
																UserIdStr                 string `json:"user_id_str"`
																IdStr                     string `json:"id_str"`
																Place                     struct {
																	Attributes struct {
																	} `json:"attributes"`
																	BoundingBox struct {
																		Coordinates [][][]float64 `json:"coordinates"`
																		Type        string        `json:"type"`
																	} `json:"bounding_box"`
																	ContainedWithin []interface{} `json:"contained_within"`
																	Country         string        `json:"country"`
																	CountryCode     string        `json:"country_code"`
																	FullName        string        `json:"full_name"`
																	Name            string        `json:"name"`
																	Id              string        `json:"id"`
																	PlaceType       string        `json:"place_type"`
																	Url             string        `json:"url"`
																} `json:"place,omitempty"`
																InReplyToScreenName string `json:"in_reply_to_screen_name,omitempty"`
																InReplyToUserIdStr  string `json:"in_reply_to_user_id_str,omitempty"`
															} `json:"legacy"`
														} `json:"result"`
													} `json:"retweeted_status_result"`
												} `json:"legacy"`
												QuickPromoteEligibility struct {
													Eligibility string `json:"eligibility"`
												} `json:"quick_promote_eligibility"`
											} `json:"result"`
										} `json:"tweet_results"`
										TweetDisplayType string `json:"tweetDisplayType"`
									} `json:"itemContent,omitempty"`
									Value               string `json:"value,omitempty"`
									CursorType          string `json:"cursorType,omitempty"`
									StopOnEmptyResponse bool   `json:"stopOnEmptyResponse,omitempty"`
								} `json:"content"`
							} `json:"entries,omitempty"`
							Entry struct {
								EntryId   string `json:"entryId"`
								SortIndex string `json:"sortIndex"`
								Content   struct {
									EntryType   string `json:"entryType"`
									ItemContent struct {
										ItemType     string `json:"itemType"`
										TweetResults struct {
											Result struct {
												Typename string `json:"__typename"`
												RestId   string `json:"rest_id"`
												Core     struct {
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
																	Url struct {
																		Urls []struct {
																			DisplayUrl  string `json:"display_url"`
																			ExpandedUrl string `json:"expanded_url"`
																			Url         string `json:"url"`
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
																ProfileBannerUrl       string `json:"profile_banner_url"`
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
																Url                     string        `json:"url"`
																Verified                bool          `json:"verified"`
																WantRetweets            bool          `json:"want_retweets"`
																WithheldInCountries     []interface{} `json:"withheld_in_countries"`
															} `json:"legacy"`
															SuperFollowEligible bool `json:"super_follow_eligible"`
															SuperFollowedBy     bool `json:"super_followed_by"`
															SuperFollowing      bool `json:"super_following"`
														} `json:"result"`
													} `json:"user_results"`
												} `json:"core"`
												DownvotePerspective struct {
													IsDownvoted bool `json:"isDownvoted"`
												} `json:"downvotePerspective"`
												UnmentionInfo struct {
												} `json:"unmention_info"`
												Legacy struct {
													CreatedAt           string `json:"created_at"`
													ConversationControl struct {
														Policy            string `json:"policy"`
														ConversationOwner struct {
															Legacy struct {
																ScreenName string `json:"screen_name"`
															} `json:"legacy"`
														} `json:"conversation_owner"`
													} `json:"conversation_control"`
													ConversationIdStr string `json:"conversation_id_str"`
													DisplayTextRange  []int  `json:"display_text_range"`
													Entities          struct {
														UserMentions []interface{} `json:"user_mentions"`
														Urls         []struct {
															DisplayUrl  string `json:"display_url"`
															ExpandedUrl string `json:"expanded_url"`
															Url         string `json:"url"`
															Indices     []int  `json:"indices"`
														} `json:"urls"`
														Hashtags []struct {
															Indices []int  `json:"indices"`
															Text    string `json:"text"`
														} `json:"hashtags"`
														Symbols []interface{} `json:"symbols"`
													} `json:"entities"`
													FavoriteCount             int    `json:"favorite_count"`
													Favorited                 bool   `json:"favorited"`
													FullText                  string `json:"full_text"`
													IsQuoteStatus             bool   `json:"is_quote_status"`
													Lang                      string `json:"lang"`
													LimitedActions            string `json:"limited_actions"`
													PossiblySensitive         bool   `json:"possibly_sensitive"`
													PossiblySensitiveEditable bool   `json:"possibly_sensitive_editable"`
													QuoteCount                int    `json:"quote_count"`
													ReplyCount                int    `json:"reply_count"`
													RetweetCount              int    `json:"retweet_count"`
													Retweeted                 bool   `json:"retweeted"`
													Source                    string `json:"source"`
													UserIdStr                 string `json:"user_id_str"`
													IdStr                     string `json:"id_str"`
													SelfThread                struct {
														IdStr string `json:"id_str"`
													} `json:"self_thread"`
												} `json:"legacy"`
												QuickPromoteEligibility struct {
													Eligibility string `json:"eligibility"`
												} `json:"quick_promote_eligibility"`
											} `json:"result"`
										} `json:"tweet_results"`
										TweetDisplayType string `json:"tweetDisplayType"`
										SocialContext    struct {
											Type        string `json:"type"`
											ContextType string `json:"contextType"`
											Text        string `json:"text"`
										} `json:"socialContext"`
									} `json:"itemContent"`
									ClientEventInfo struct {
										Component string `json:"component"`
										Details   struct {
											TimelinesDetails struct {
												InjectionType string `json:"injectionType"`
											} `json:"timelinesDetails"`
										} `json:"details"`
									} `json:"clientEventInfo"`
								} `json:"content"`
							} `json:"entry,omitempty"`
						} `json:"instructions"`
						ResponseObjects struct {
							FeedbackActions    []interface{} `json:"feedbackActions"`
							ImmediateReactions []interface{} `json:"immediateReactions"`
						} `json:"responseObjects"`
					} `json:"timeline"`
				} `json:"timeline_v2"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}
