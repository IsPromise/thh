package controllers

import (
	"github.com/leancodebox/goose/array"
	"thh/app/http/controllers/component"
	"thh/app/models/TodoTask"
	"time"
)

type IdValue struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

func TodoStatusList() component.Response {
	return component.SuccessResponse([]IdValue{
		{0, "创建"},
		{1, "完成"},
		{2, "暂停"},
		{3, "作废"},
	})
}

type CreateTaskRequest struct {
	TaskName    string   `json:"taskName" validate:"required"`
	Description string   `json:"description"`
	Deadline    WrapTime `json:"deadline"`
	Weight      int      `json:"weight"`
}

func CreateTask(request CreateTaskRequest) component.Response {
	todoTaskEntity := TodoTask.Entity{
		//TaskId:          0,
		TaskName:        request.TaskName,
		TaskDescription: request.Description,
		Status:          0,
		CreateTime:      time.Now(),
		Deadline:        request.Deadline.Time,
		Weight:          request.Weight,
	}
	TodoTask.Create(&todoTaskEntity)
	return component.SuccessResponse(todoTaskEntity.ToDto())
}

type UpdateTaskRequest struct {
	TaskId      int      `json:"taskId"`
	TaskName    string   `json:"taskName"`
	Description string   `json:"description"`
	Status      int      `json:"status"`
	Deadline    WrapTime `json:"deadline"`
	Weight      int      `json:"weight"`
	Paused      int      `json:"paused"`
}

func UpdateTask(request UpdateTaskRequest) component.Response {
	entity := TodoTask.Get(request.TaskId)
	if entity.TaskId == 0 {
		return component.FailResponse("任务不存在")
	}
	entity.TaskName = request.TaskName
	entity.TaskDescription = request.Description
	entity.Status = request.Status
	entity.Deadline = request.Deadline.Time
	entity.Weight = request.Weight
	entity.Paused = request.Paused
	TodoTask.Save(&entity)
	return component.SuccessResponse(entity.ToDto())
}

type FindTodoListRequest struct {
	Status  []int `json:"status"`
	NeedAll bool  `json:"needAll"`
}

func FindTodoList(request FindTodoListRequest) component.Response {
	tasks := TodoTask.QueryAll(request.NeedAll,request.Status)
	return component.SuccessResponse(array.ArrayMap(func(t *TodoTask.Entity) TodoTask.EntityDto {
		return t.ToDto()
	}, tasks))
}

const (
	YyyyMmDd       = "2006-01-02"
	YyyyMmDdHhMmSs = "2006-01-02 15:04:05"
)

type WrapTime NullTime

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (t *WrapTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		t.Valid = false
		return nil
	}

	var now time.Time
	if len(string(data)) == len(YyyyMmDd)+2 {
		now, err = time.ParseInLocation(`"`+YyyyMmDd+`"`, string(data), time.Local)
		t.Valid = true
		t.Time = now
	} else {
		now, err = time.ParseInLocation(`"`+YyyyMmDdHhMmSs+`"`, string(data), time.Local)
		t.Valid = true
		t.Time = now
	}
	return
}

func (t *WrapTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(YyyyMmDdHhMmSs)+2)
	b = append(b, '"')
	b = t.Time.AppendFormat(b, YyyyMmDdHhMmSs)
	b = append(b, '"')
	return b, nil
}

func (t *WrapTime) String() string {
	if !t.Valid {
		return "null"
	}
	return t.Time.Format(YyyyMmDdHhMmSs)
}
