package FTwitterUser

import querybuild "thh/arms/querymaker"

func Create(entity *FTwitterUser) int64 {
	result := builder().Create(entity)
	return result.RowsAffected
}

func Save(entity *FTwitterUser) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func SaveAll(entities *[]FTwitterUser) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}

func Update(entity *FTwitterUser) {
	builder().Save(entity)
}

func UpdateAll(entities *[]FTwitterUser) {
	builder().Save(entities)
}

func Delete(entity FTwitterUser) int64 {
	result := builder().Delete(&entity)
	return result.RowsAffected
}

func Get(id any) (entity FTwitterUser) {
	builder().Where(pid, id).First(&entity)
	return
}

func GetByName(name string) (entity FTwitterUser) {
	builder().Where(querybuild.Eq(fieldScreenName, name)).First(&entity)
	return
}

func GetByRestId(restId string) (entity FTwitterUser) {
	builder().Where(querybuild.Eq(fieldRestId, restId)).First(&entity)
	return
}

func GetBy(field, value string) (entity FTwitterUser) {
	builder().Where(field+" = ?", value).First(&entity)
	return
}

func All() (entities []FTwitterUser) {
	builder().Find(&entities)
	return
}

func GetByDesc(desc string) (entities []FTwitterUser) {
	builder().Where(querybuild.Like(fieldDesc, desc)).Order(querybuild.Desc(pid)).Limit(1000).Find(&entities)
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
	Data     []FTwitterUser
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
	Data     []FTwitterUser
} {
	var list []FTwitterUser
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
		b.Where(querybuild.Like(fieldDesc, q.Search))
	}
	b.Limit(q.PageSize).Offset(q.PageSize * q.Page).Order("id desc").Find(&list)

	var total int64
	if q.Search != "" {
		builder().Where(querybuild.Like(fieldDesc, q.Search)).Count(&total)
	} else {
		builder().Count(&total)
	}
	return struct {
		Page     int
		PageSize int
		Total    int64
		Data     []FTwitterUser
	}{Page: q.Page, PageSize: q.PageSize, Data: list, Total: total}
}