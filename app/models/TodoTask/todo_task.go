package TodoTask

import (
	"time"
)

const tableName = "todo_task"
const pid = "task_id"
const fieldTaskName = "task_name"
const fieldTaskDescription = "task_description"
const fieldStatus = "status"
const fieldCreateTime = "create_time"
const fieldDeadline = "deadline"
const fieldWeight = "weight"
const fieldPaused = "paused"

type Entity struct {
	TaskId          uint64    `gorm:"primaryKey;column:task_id;autoIncrement;not null;" json:"taskId"`                        // 主键
	TaskName        string    `gorm:"column:task_name;type:varchar(50);not null;" json:"taskName"`                            // 任务名
	TaskDescription string    `gorm:"column:task_description;type:varchar(1024);not null;" json:"taskDescription"`            //
	Status          int       `gorm:"column:status;type:int;not null;" json:"status"`                                         // 任务状态（0：未完成，1：已完成）
	CreateTime      time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"createTime"` //
	Deadline        time.Time `gorm:"column:deadline;type:datetime;not null;" json:"deadline"`                                //
	Weight          int       `gorm:"column:weight;type:int;not null;" json:"weight"`                                         // 任务权重（用于优先级排序）
	Paused          int       `gorm:"column:paused;type:int;not null;" json:"paused"`                                         // 任务暂停状态（0：未暂停，1：已暂停）
}

// func (itself *Entity) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *Entity) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *Entity) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *Entity) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *Entity) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *Entity) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *Entity) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *Entity) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *Entity) AfterFind(tx *gorm.DB) (err error) {}

func (itself *Entity) TableName() string {
	return tableName
}

type status int

const (
	NotFinished status = iota
	Finish
)

type paused int

const (
	Running paused = iota
	Stop
)

type EntityDto struct {
	TaskId          uint64 `json:"taskId"`          // 主键
	TaskName        string `json:"taskName"`        // 任务名
	TaskDescription string `json:"taskDescription"` //
	Status          int    `json:"status"`          // 任务状态（0：未完成，1：已完成）
	CreateTime      string `json:"createTime"`      //
	Deadline        string `json:"deadline"`        //
	Weight          int    `json:"weight"`          // 任务权重（用于优先级排序）
	Paused          int    `json:"paused"`          // 任务暂停状态（0：未暂停，1：已暂停）
}

func (itself *Entity) ToDto() EntityDto {
	return EntityDto{
		TaskId:          itself.TaskId,
		TaskName:        itself.TaskName,
		TaskDescription: itself.TaskDescription,
		Status:          itself.Status,
		CreateTime:      itself.CreateTime.Format(time.DateTime),
		Deadline:        itself.Deadline.Format(time.DateTime),
		Weight:          itself.Weight,
		Paused:          itself.Paused,
	}

}
