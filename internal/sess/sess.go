package sess

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

func NewStore() {
	log.Println("Session new")
	if store == nil {
		store = session.New(session.Config{Expiration: 8640 * time.Hour})
	}
}

func Store() *session.Store {
	return store
}

func Get(c *fiber.Ctx) (*session.Session, error) {
	sess, err := store.Get(c)
	if err != nil {
		log.Println("ERROR getting session:", err)
	}

	return sess, err
}
