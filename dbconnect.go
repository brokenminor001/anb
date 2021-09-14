package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type tables struct {
	team_name string
	points    int
}

func main() {

	connStr := "user=asmolin password=oadnpvia dbname=anb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select team_name,points from groupe_one")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	table := []tables{}

	for rows.Next() {
		p := tables{}
		err := rows.Scan(&p.team_name, &p.points)
		if err != nil {
			fmt.Println(err)
			continue
		}
		table = append(table, p)
	}
	for _, p := range table {
		fmt.Println(p.team_name, p.points)

	}
}
