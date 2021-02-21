package main

import (
	"database/sql"
	"fmt"
)

type course struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Units        string `json:"units"`
	Description  string `json:"description"`
	UpperDivOnly bool `json:"upper_div_only"`
	GradOnly     bool `json:"grad_only"`
}

func (c *course) getCourse(db *sql.DB) error {
    return db.QueryRow(
		"SELECT c.title, c.units, c.description, c.upper_only, c.grad_only FROM courses c WHERE c.ID = $1",
		c.ID,
	).Scan(&c.Title, &c.Units, &c.Description, &c.UpperDivOnly, &c.GradOnly)
}

func (c *course) getPreqs(db *sql.DB) ([]course, error) {
	rows, err := db.Query(`
		SELECT c.id, c.title, c.units, c.description, c.upper_only, c.grad_only
		FROM course_prerequisites cp
		JOIN prerequisites p ON cp.prerequisite_id = p.group_id
		JOIN courses c ON p.courses = c.id
		WHERE cp.id = $1
	`, c.ID)

	fmt.Println(rows)

	if err != nil {
        return nil, err
    }

	defer rows.Close()

	courses := []course{}

	for rows.Next() {
		var c course
		if err := rows.Scan(&c.ID, &c.Units, &c.Description, &c.UpperDivOnly, &c.GradOnly); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}

	return courses, nil
}