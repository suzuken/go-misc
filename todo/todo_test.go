package todo

import (
	"testing"
)

func TestCreate(t *testing.T) {
	c := New("test.db")
	if err := c.Create("cooking"); err != nil {
		t.Fatalf("cooking failed %s", err)
	}
	tasks, err := c.ListTasks()
	if err != nil {
		t.Fatalf("get task list failed %s", err)
	}
	if len(tasks) != 1 {
		t.Fatalf("task list too long. got %d, want %d", len(tasks), 1)
	}
}

func TestShowTasks(t *testing.T) {
}
