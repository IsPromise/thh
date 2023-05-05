package twservice

import (
	"thh/app/models/FTwitter/FTwitterSpiderHis"
	"thh/app/models/FTwitter/FTwitterUser"
	"thh/app/models/FTwitter/FTwitterUserHis"
	"thh/bundles/restyopt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
)

func SaveUserEntity(restId, screenName, desc, name string) FTwitterUser.FTwitterUser {
	userEntity := FTwitterUser.GetByRestId(restId)
	// 内容变了 或者是新创建的
	if userEntity.Id == 0 || !(userEntity.Desc == desc && userEntity.ScreenName == screenName && userEntity.Name == name) {
		FTwitterUserHis.Save(&FTwitterUserHis.FTwitterUserHis{RestId: restId, ScreenName: screenName, Desc: desc, Name: name})
	}
	if userEntity.Id == 0 {
		userEntity.ScreenName = screenName
		userEntity.RestId = restId
		userEntity.CreateTime = time.Now()
	}
	userEntity.Name = name
	userEntity.Desc = desc
	FTwitterUser.Save(&userEntity)
	return userEntity
}

func SaveTSpiderHis(typeId int, target string, r *resty.Response, err error) int64 {
	curlBuildStr := restyopt.GetCurlByR(*r)
	successStr := "0"
	if err != nil {
		successStr = cast.ToString(err.Error())
	}
	entity := FTwitterSpiderHis.FTwitterSpiderHis{Type: typeId, Target: target, Curl: curlBuildStr, Success: successStr, Content: r.String(), CreateTime: time.Now()}
	return FTwitterSpiderHis.Save(&entity)
}
