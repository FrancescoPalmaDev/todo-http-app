package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Id    int
	Title string
}

type ToDoList struct {
	Tasks    []Task
	FilePath string
}

func (t *ToDoList) Load() error {
	data, err := os.ReadFile(t.FilePath)
	if err != nil {
		return nil
	}
	return json.Unmarshal(data, &t.Tasks)
}

func (t *ToDoList) Save() error {
	data, err := json.Marshal(t.Tasks)
	if err == nil {
		err = os.WriteFile(t.FilePath, data, 0644)
	}
	return err
}

func (t *ToDoList) Add(title string) error {
	newTask := Task{len(t.Tasks) + 1, title}
	t.Tasks = append(t.Tasks, newTask)
	return t.Save()
}

func (t *ToDoList) Remove(id int) error {
	for i, task := range t.Tasks {
		if task.Id == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			break
		}
	}
	return t.Save()
}

func (t *ToDoList) PrintList() {
	for i := 0; i < len(t.Tasks); i++ {
		fmt.Printf("%d - %s\n", t.Tasks[i].Id, t.Tasks[i].Title)
	}
}
