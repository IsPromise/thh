package FTwitterSpiderHis

import "github.com/spf13/cast"

func Create(entity *FTwitterSpiderHis) int64 {
	result := builder().Create(entity)
	return result.RowsAffected
}

func Save(entity *FTwitterSpiderHis) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func saveAll(entities *[]FTwitterSpiderHis) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}

func deleteEntity(entity *FTwitterSpiderHis) int64 {
	result := builder().Delete(&entity)
	return result.RowsAffected
}

func get(id any) (entity FTwitterSpiderHis) {
	builder().First(&entity, id)
	return
}

func all() (entities []FTwitterSpiderHis) {
	builder().Find(&entities)
	return
}

type PageQuery struct {
	Page, PageSize int
	Search         string
}

func Page(q PageQuery) struct {
	Page     int
	PageSize int
	Total    int64
	Data     []FTwitterSpiderHis
} {
	var list []FTwitterSpiderHis
	if q.Page > 0 {
		q.Page -= 1
	} else {
		q.Page = 0
	}
	if q.PageSize < 1 {
		q.PageSize = 1
	}
	b := builder()
	b.Limit(q.PageSize).Offset(q.PageSize * q.Page).Order("id desc").Find(&list)

	latest := GetLatest()

	return struct {
		Page     int
		PageSize int
		Total    int64
		Data     []FTwitterSpiderHis
	}{Page: q.Page, PageSize: q.PageSize, Data: list, Total: cast.ToInt64(latest.Id)}
}

func GetLatest() (entity FTwitterSpiderHis) {
	builder().Order("id desc").Limit(1).First(&entity)
	return
}
