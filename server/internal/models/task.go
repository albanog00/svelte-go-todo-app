package models

import (
	"errors"
	"time"
)

type Task struct {
	Id          string    `json:"id" gorm:"primary_key"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"createdAt" gorm:"default:current_timestamp(3)"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"default:NULL ON UPDATE current_timestamp(3)"`
	DeletedAt   time.Time `json:"deletedAt" gorm:"default:NULL"`
}

const LIMIT int = 5

func GetTasks(pageNum int64) ([]*Task, int64, error) {
	var tasks []*Task
	var count int64 = 0

	res := db.Model(&Task{}).
		Count(&count).
		Offset(LIMIT * int(pageNum)).
		Limit(LIMIT).
		Find(&tasks)

	if res.Error != nil {
		return nil, 0, errors.New("no tasks found")
	}
	return tasks, count, nil
}

func CreateTask(task *Task) (*Task, error) {
	res := db.Create(task)
	if res.RowsAffected == 0 {
		return &Task{}, errors.New("can't create task")
	}
	return task, nil
}

func DeleteTask(id string) error {
	res := db.Delete(&Task{}, "id LIKE ?", id)
	if res.RowsAffected == 0 {
		return errors.New("can't delete task")
	}

	return nil
}
