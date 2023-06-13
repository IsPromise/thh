package FTwitterTweet

import "github.com/leancodebox/goose/querymaker"

func Create(entity *FTwitterTweet) int64 {
	result := builder().Create(entity)
	return result.RowsAffected
}

func Save(entity *FTwitterTweet) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func SaveAll(entities *[]FTwitterTweet) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}

func DeleteEntity(entity *FTwitterTweet) int64 {
	result := builder().Delete(&entity)
	return result.RowsAffected
}

func Get(id any) (entity FTwitterTweet) {
	builder().First(&entity, id)
	return
}
func GetUserTweet(screenName, conversationIdStr string) (entity FTwitterTweet) {
	builder().Where(querymaker.Eq(fieldScreenName, screenName)).
		Where(querymaker.Eq(fieldConversationId, conversationIdStr)).First(&entity)
	return
}

func All() (entities []FTwitterTweet) {
	builder().Find(&entities)
	return
}
func GetByContent(desc string) (entities []FTwitterTweet) {
	builder().Where(querymaker.Like(fieldContext, desc)).Order(querymaker.Desc(pid)).Limit(1000).Find(&entities)
	return
}

func DefaultPage(page int) PageResult[*FTwitterTweet] {
	return Page(PageQuery{
		Page:     page,
		PageSize: 10,
	})
}

type PageQuery struct {
	Page, PageSize int
	Search         string
	UserFilter     []string
}

type PageResult[T any] struct {
	Page     int
	PageSize int
	Total    int64
	Data     []T
}

func Page(q PageQuery) PageResult[*FTwitterTweet] {
	var list []*FTwitterTweet
	if q.Page > 0 {
		q.Page -= 1
	} else {
		q.Page = 0
	}
	if q.PageSize < 1 {
		q.PageSize = 1
	}
	b := builder()
	if q.Search != "" {
		b.Where(querymaker.Like(fieldContext, q.Search))
	}
	if len(q.UserFilter) > 0 {
		b.Where(querymaker.NotIn(fieldOriginScreenName, q.UserFilter))
	}
	var total int64
	b.Count(&total)
	b.Limit(q.PageSize).Offset(q.PageSize * q.Page).Order("id desc").Find(&list)

	return struct {
		Page     int
		PageSize int
		Total    int64
		Data     []*FTwitterTweet
	}{Page: q.Page, PageSize: q.PageSize, Data: list, Total: total}
}
