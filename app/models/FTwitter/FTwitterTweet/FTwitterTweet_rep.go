package FTwitterTweet

import "thh/arms/querymaker"

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

func Update(entity *FTwitterTweet) {
	builder().Save(entity)
}

func UpdateAll(entities *[]FTwitterTweet) {
	builder().Save(entities)
}

func Delete(entity FTwitterTweet) int64 {
	result := builder().Delete(&entity)
	return result.RowsAffected
}

func Get(id any) (entity FTwitterTweet) {
	builder().Where(pid, id).First(&entity)
	return
}
func GetUserTweet(screenName, conversationIdStr string) (entity FTwitterTweet) {
	builder().Where(querymaker.Eq(fieldScreenName, screenName)).
		Where(querymaker.Eq(fieldConversationId, conversationIdStr)).First(&entity)
	return
}

func GetBy(field, value string) (entity FTwitterTweet) {
	builder().Where(field+" = ?", value).First(&entity)
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

func IsExist(field, value string) bool {
	var count int64
	builder().Where(field+" = ?", value).Count(&count)
	return count > 0
}

func DefaultPage(page int) struct {
	Page     int
	PageSize int
	Total    int64
	Data     []FTwitterTweet
} {
	return Page(PageQuery{
		Page:     page,
		PageSize: 10,
	})
}

type PageQuery struct {
	Page, PageSize int
	Search         string
}

func Page(q PageQuery) struct {
	Page     int
	PageSize int
	Total    int64
	Data     []FTwitterTweet
} {
	var list []FTwitterTweet
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
	b.Limit(q.PageSize).Offset(q.PageSize * q.Page).Order("id desc").Find(&list)

	var total int64
	if q.Search != "" {
		builder().Where(querymaker.Like(fieldContext, q.Search)).Count(&total)
	} else {
		builder().Count(&total)
	}
	return struct {
		Page     int
		PageSize int
		Total    int64
		Data     []FTwitterTweet
	}{Page: q.Page, PageSize: q.PageSize, Data: list, Total: total}
}
