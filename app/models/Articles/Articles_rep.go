package Articles

import (
	"context"
	querybuild "thh/arms/querymaker"
)

type Rep struct {
	ctx *context.Context
}

func NewRep(ctx *context.Context) Rep {
	return Rep{
		ctx: ctx,
	}
}

func (itself Rep) Save(entity *Articles) int64 {
	result := builder().WithContext(*itself.ctx).Save(entity)
	return result.RowsAffected
}

func Create(entity *Articles) int64 {
	result := builder().Create(entity)
	return result.RowsAffected
}

func Save(entity *Articles) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func SaveAll(entities *[]Articles) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}

func Delete(entity *Articles) int64 {
	result := builder().Delete(entity)
	return result.RowsAffected
}

func Get(id any) (entity Articles) {
	builder().Where(querybuild.Eq(pid, id)).First(&entity)
	return
}

func GetBy(field, value string) (entity Articles) {
	builder().Where(field+" = ?", value).First(&entity)
	return
}

func All() (entities []Articles) {
	builder().Find(&entities)
	return
}

func IsExist(field, value string) bool {
	var count int64
	builder().Where(field+" = ?", value).Count(&count)
	return count > 0
}

func GetByMaxIdPage(id uint64, pageSize int) (entities []Articles) {
	builder().Where(querybuild.Gt(pid, id)).Order(querybuild.Desc(fieldUpdateTime)).Limit(pageSize).Find(&entities)
	return
}
