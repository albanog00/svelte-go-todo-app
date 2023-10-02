package models

import (
	"errors"
)

type Task struct {
	Id          string `json:"id" gorm:"primary_key"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Time        string `json:"time"`
}

func GetTasks() ([]*Task, error) {
	var tasks []*Task
	res := db.Find(&tasks)
	if res.Error != nil {
		return nil, errors.New("no tasks found")
	}
	return tasks, nil
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
