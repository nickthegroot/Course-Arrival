package main

import (
	"database/sql"
	"log"
	"math/rand"
)

type course struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Units        string `json:"units"`
	Description  string `json:"description"`
	UpperDivOnly bool `json:"upper_div_only"`
	GradOnly     bool `json:"grad_only"`
	PrereqID string `json:"-"`
}

type department string

type coursePlan struct {
	Courses []course `json:"course_plan"`
	Dot	string `json:"dot"`
}

func (d *department) getDeptCourses(db *sql.DB) ([]course, error) {
	rows, err := db.Query(`
		SELECT c.id, c.title, c.units, c.description, c.upper_only, c.grad_only, d.group
		FROM department d
		JOIN department_prerequisites dp ON d.group = dp.group
		JOIN courses c on dp.course = c.id
		WHERE d.code = $1
	`, d)

	if err != nil {
		log.Fatal(err)
        return nil, err
    }

	defer rows.Close()

	groups := make(map[string][]course)

	for rows.Next() {
		var c course
		if err := rows.Scan(&c.ID, &c.Title, &c.Units, &c.Description, &c.UpperDivOnly, &c.GradOnly, &c.PrereqID); err != nil {
			log.Fatal(err)
			return nil, err
		}
		groups[c.PrereqID] = append(groups[c.PrereqID], c)
	}

	picks := []course{}
	for _, courses := range groups {
		randIdx := rand.Intn(len(courses))
		pick := courses[randIdx]
		picks = append(picks, pick)
	}

	return picks, nil
}

func (c *course) getCourse(db *sql.DB) error {
    return db.QueryRow(
		"SELECT c.title, c.units, c.description, c.upper_only, c.grad_only FROM courses c WHERE c.ID = $1",
		c.ID,
	).Scan(&c.Title, &c.Units, &c.Description, &c.UpperDivOnly, &c.GradOnly)
}

func (c *course) getPreqs(db *sql.DB) ([]course, error) {
	rows, err := db.Query(`
		SELECT c.id, c.title, c.units, c.description, c.upper_only, c.grad_only, p.group_id
		FROM courses_prerequisites cp
		JOIN prerequisites p ON cp.prerequisites_id = p.group_id
		JOIN courses c ON p.courses = c.id
		WHERE cp.id = $1
	`, c.ID)

	if err != nil {
		log.Fatal(err)
        return nil, err
    }

	defer rows.Close()

	groups := make(map[string][]course)

	for rows.Next() {
		var c course
		if err := rows.Scan(&c.ID, &c.Title, &c.Units, &c.Description, &c.UpperDivOnly, &c.GradOnly, &c.PrereqID); err != nil {
			log.Fatal(err)
			return nil, err
		}
		groups[c.PrereqID] = append(groups[c.PrereqID], c)
	}

	picks := []course{}
	for _, courses := range groups {
		randIdx := rand.Intn(len(courses))
		pick := courses[randIdx]
		picks = append(picks, pick)
	}

	return picks, nil
}