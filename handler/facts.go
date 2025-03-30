package handler

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Home route handler!")
}

func GetFacts(c *fiber.Ctx) error {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM facts")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer rows.Close()

	var facts []Fact

	for rows.Next() {
		var fact Fact
		err := rows.Scan(&fact.ID, &fact.Question, &fact.Answer)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		facts = append(facts, fact)
	}

	return c.JSON(facts)
}

type Fact struct {
	ID      int    `json:"id"`
	Question string `json:"question"`
	Answer  string `json:"answer"`
}
