package controllers

import (
	"fmt"
	"thh/app/http/controllers/component"
	"thh/app/models/Users"
	"thh/arms/jwt"
	"thh/arms/logger"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/spf13/cast"
)

const (
	expireTime = time.Second * 86400 * 7
)

type RegReq struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"userName"  validate:"required"`
	Password string `json:"passWord"  validate:"required"`
	NickName string `json:"nickName" gorm:"default:'QMPlusUser'"`
}

// Register
// @todo user表增加验证字段
// 创建后验证码存入redis，发认证送邮件。
// 邮件 附有 url?code=xxx
// 验证后更新验证字段
// 清除验证码
func Register(r RegReq) component.Response {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

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
	Username string `json:"userName"   validate:"required"`
	Password string `json:"password"   validate:"required"`
}

func Login(r LoginReq) component.Response {
	userEntity, err := Users.Verify(r.Username, r.Password)
	if err != nil {
		logger.Info(err)
		return component.FailResponse("验证失败")
	}
	token, err := jwt.CreateNewToken(userEntity.Id, expireTime)
	if err != nil {
		logger.Info(err)
		return component.FailResponse("验证失败")
	}
	return component.SuccessResponse(component.DataMap{
		"message": "ok",
		"token":   token,
	})
}

type null struct {
}

func UserInfoV4(req component.BetterRequest[null]) component.Response {
	userEntity, err := req.GetUser()
	if err != nil {
		return component.FailResponse("账号异常" + err.Error())
	}
	return component.SuccessResponse(userEntity)
}

func GetCaptcha() component.Response {
	captchaId, captchaImg := GenerateCaptcha()
	return component.SuccessResponse(map[string]any{
		"captcha_id":  captchaId,
		"captcha_img": captchaImg,
	})
}

func VerifyCaptchaApi(c *gin.Context) component.Response {
	captchaId := c.PostForm("captcha_id")
	captchaCode := c.PostForm("captcha_code")

	if VerifyCaptcha(captchaId, captchaCode) {
		return component.SuccessResponse("ok")
	} else {
		return component.FailResponse("invalid captcha")
	}
}

func GenerateCaptcha() (string, string) {
	return "", ""
}

func VerifyCaptcha(captchaId, captchaCode string) bool {
	return true
}
