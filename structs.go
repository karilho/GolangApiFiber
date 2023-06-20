package main

type People struct {
	ID    int    `ksql:"id" json:"id"`
	Name  string `sql:"name" json:"name"`
	Email string `ksql:"email" json:"email"`
	Phone int    `ksql:"phone" json:"phone"`
}
