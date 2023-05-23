package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/fogleman/gg"
	"math/rand"
	"strconv"
	"thh/app/http/controllers/component"
	"thh/app/models/DataReps"
	"thh/app/models/Users"
	"thh/bundles/logging"
	"time"

	"github.com/leancodebox/goose/jwt"

	"github.com/spf13/cast"
)

const (
	expireTime = time.Second * 86400 * 7
)

type RegReq struct {
	Email          string `json:"email" validate:"required"`
	Username       string `json:"userName"  validate:"required"`
	Password       string `json:"passWord"  validate:"required"`
	NickName       string `json:"nickName" gorm:"default:'QMPlusUser'"`
	InvitationCode string `json:"invitationCode"`
}

// Register
// @todo user表增加验证字段
// 创建后验证码存入redis，发认证送邮件。
// 邮件 附有 url?code=xxx
// 验证后更新验证字段
// 清除验证码
func Register(r RegReq) component.Response {
	userEntity := Users.MakeUser(r.Username, r.Password, r.Email)
	err := Users.Create(userEntity)

	if err != nil {
		return component.FailResponse(cast.ToString(err))
	}

	token, err := jwt.CreateNewToken(userEntity.Id, expireTime)
	if err != nil {
		return component.FailResponse(cast.ToString(err))
	}
	return component.SuccessResponse(component.DataMap{
		"message": "ok",
		"token":   token,
	})
}

type LoginReq struct {
	Username    string `json:"userName"   validate:"required"`
	Password    string `json:"password"   validate:"required"`
	CaptchaId   string `json:"captchaId"`
	CaptchaCode string `json:"captchaCode"`
}

func Login(r LoginReq) component.Response {
	//if VerifyCaptcha(r.CaptchaId, r.CaptchaCode) {
	//	return component.SuccessResponse("ok")
	//}
	userEntity, err := Users.Verify(r.Username, r.Password)
	if err != nil {
		logging.Info(err)
		return component.FailResponse("验证失败")
	}
	token, err := jwt.CreateNewToken(userEntity.Id, expireTime)
	if err != nil {
		logging.Info(err)
		return component.FailResponse("验证失败")
	}
	return component.SuccessResponse(component.DataMap{
		"message": "ok",
		"token":   token,
	})
}

func GetCaptcha() component.Response {
	randString, captchaImg := GenerateCaptcha()
	key := cast.ToString(time.Now().Nanosecond()) + cast.ToString(rand.Float64())
	DataReps.Set(key, randString)
	return component.SuccessResponse(map[string]any{
		"captchaKey": key,
		"captchaImg": captchaImg,
	})
}

func VerifyCaptcha(captchaKey, captchaCode string) bool {
	if DataReps.Get(captchaKey) == captchaCode {
		return true
	}
	return false
}

func GenerateCaptcha() (string, string) {
	// 生成随机字符串
	code := RandString(4)

	width := 120
	height := 40
	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// 将字符串绘制到画布上
	for i, c := range code {
		dc.SetRGB(rand.Float64(), rand.Float64(), rand.Float64())
		x := float64(i)*float64(width)/4 + rand.Float64()*10 - 5
		y := rand.Float64()*20 + 10
		dc.DrawStringAnchored(string(c), x, y, 0.5, 0.5)
	}

	// 添加干扰线
	for i := 0; i < 9; i++ {
		dc.SetRGBA(rand.Float64(), rand.Float64(), rand.Float64(), 0.2)
		x1 := rand.Float64() * float64(width)
		y1 := rand.Float64() * float64(height)
		x2 := rand.Float64() * float64(width)
		y2 := rand.Float64() * float64(height)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}

	buffer := bytes.NewBuffer(nil)
	err := dc.EncodePNG(buffer)
	if err != nil {
		return "", ""
	}
	enc := `data:image/png;base64,` + base64.StdEncoding.EncodeToString(buffer.Bytes())
	return code, enc
}

// RandString 生成指定长度的随机字符串
func RandString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type null struct {
}

// UserInfo 用户信息
func UserInfo(req component.BetterRequest[null]) component.Response {
	userEntity, err := req.GetUser()
	if err != nil {
		return component.FailResponse("账号异常" + err.Error())
	}
	return component.SuccessResponse(userEntity)
}

type EditUserInfoReq struct {
	Nickname string `json:"nickname"`
}

// EditUserInfo 编辑用户
func EditUserInfo(req component.BetterRequest[EditUserInfoReq]) component.Response {
	return component.SuccessResponse("success")
}

// Invitation 邀请码
func Invitation(req component.BetterRequest[null]) component.Response {
	base36 := strconv.FormatInt(int64(req.UserId), 36)
	return component.SuccessResponse(map[string]any{
		"invitation": base36,
	})
}
