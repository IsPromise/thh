package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"thh/app/models/Articles"
	"thh/app/models/Comment"
	"thh/app/models/Users"
	"thh/arms"
	"time"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "tool:articles_make",
		Short: "articles_make",
		Run:   runArticlesMake,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
	appendCommand(&cobra.Command{
		Use:   "tool:createAndDeleted",
		Short: "createAndDeleted",
		Run:   createAndDeleted,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})

	appendCommand(&cobra.Command{
		Use:   "tool:createAndUpdate",
		Short: "createAndDeleted",
		Run:   createAndUpdate,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func createAndUpdate(_ *cobra.Command, _ []string) {
	art := Articles.Articles{UserId: 1, Content: `
你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好
`}
	Articles.Save(&art)

	art.Content = "haohaohaohaohao"

	time.Sleep(time.Second * 3)

	Articles.Save(&art)

	fmt.Println(art)
}

func createAndDeleted(_ *cobra.Command, _ []string) {
	art := Articles.Articles{UserId: 1, Content: `
你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好
`}
	Articles.Save(&art)

	Articles.Delete(&art)

	fmt.Println(art)
}

func runArticlesMake(_ *cobra.Command, _ []string) {
	userEntity := Users.MakeUser(cast.ToString(time.Now().UnixMilli()), "123456", cast.ToString(time.Now())+"@qq.com")
	err := Users.Create(userEntity)
	if err != nil {
		fmt.Println("用户创建失败", err)
	}

	userList := Users.All()
	fmt.Print(userList)
	ctx := context.WithValue(context.Background(), "traceId", arms.GetTrace())
	fmt.Println(ctx.Value("traceId"))

	ArticlesRep := Articles.NewRep(&ctx)
	CommentRep := Comment.NewRep(&ctx)
	for _, user := range userList {
		for i := 0; i < 10; i++ {

			art := Articles.Articles{UserId: user.Id, Content: `
你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好
`}
			ArticlesRep.Save(&art)
			for _, cUser := range userList {
				comment := Comment.Comment{UserId: cUser.Id, ArticleId: art.Id, Content: cUser.Username + "觉得不错"}
				CommentRep.Save(&comment)
				comment = Comment.Comment{UserId: cUser.Id, ArticleId: art.Id, Content: cUser.Username + "觉得不错"}
				CommentRep.Save(&comment)
				comment = Comment.Comment{UserId: cUser.Id, ArticleId: art.Id, Content: cUser.Username + "觉得不错"}
				CommentRep.Save(&comment)
			}
		}
	}
}