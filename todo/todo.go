package todo

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	db *sql.DB
}

func New(dsn string) *Client {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("open db failed: %s", err)
	}
	if _, err := db.Exec(`create table if not exists tasks(id integer not null primary key, name text, done integer)`); err != nil {
		log.Fatalf("create table failed %s", err)
	}
	return &Client{db}
}

// TaskはTODOリストのタスク
type Task struct {
	ID   int
	Name string
	// Doneは完了状態を示す。trueなら終わったこととする
	Done bool
}

func (c *Client) ListTasks() ([]Task, error) {
	rows, err := c.db.Query("select * from tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var id int
		var name string
		var done bool
		if err := rows.Scan(&id, &name, &done); err != nil {
			return nil, err
		}
		tasks = append(tasks, Task{id, name, done})
	}
	return tasks, nil
}

func (c *Client) ShowTasks() (string, error) {
	ts, err := c.ListTasks()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	for _, t := range ts {
		if t.Done {
			buf.WriteString(fmt.Sprintf("%s: (done)\n", t.Name))
		} else {
			buf.WriteString(fmt.Sprintf("%s\n", t.Name))
		}
	}
	return buf.String(), nil
}

// Createはタスクを作成する
func (c *Client) Create(name string) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into tasks(id, name, done) values(null, ?, 0)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(name); err != nil {
		return err
	}
	return tx.Commit()
}

// Completeはタスクを完了する
func (c *Client) Complete(id string) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update tasks set done = 1 where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(id); err != nil {
		return err
	}
	return tx.Commit()
}

// CompleteByNameはタスクを完了する
func (c *Client) CompleteByName(name string) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update tasks set done = 1 where name = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(name); err != nil {
		return err
	}
	return tx.Commit()
}
