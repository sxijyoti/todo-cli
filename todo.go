package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	ID            int
	Title         string
	DueDate       string
	Completed     bool
	CompletedDate string
}

type Todos []Todo

func (t *Todos) Add(title string, dueDate string) {
	todo := Todo{
		ID:            len(*t) + 1,
		Title:         title,
		DueDate:       dueDate,
		Completed:     false,
		CompletedDate: "",
	}
	*t = append(*t, todo)
}

func (t *Todos) Completed(id int) {
	if err := t.indexValidate((id)); err != nil {
		fmt.Println(err)
		return
	}
	(*t)[id-1].Completed = true
	(*t)[id-1].CompletedDate = time.Now().Format("23-12-2006 15:04")
}

func (t *Todos) indexValidate(id int) error {
	if id < 1 || id > len(*t) {
		e := errors.New("Invalid ID")
		fmt.Println("Invalid ID")
		return e
	}
	return nil
}

func (t *Todos) Delete(id int) {
	if err := t.indexValidate(id); err != nil {
		fmt.Println(err)
		return
	}
	*t = append((*t)[:id-1], (*t)[id:]...)
}
