package sess

import (
	"log"
	"time"

	"git.martianoids.com/queru/retroserver/internal/cfg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

func NewStore() {
	log.Println("Session new")
	var domain string
	if cfg.IsDev() {
		domain = "localhost"
	} else {
		domain = "updater.retrowiki.es"
	}
	if store == nil {
		store = session.New(session.Config{
			Expiration:     8640 * time.Hour,
			CookieName:     "rw_sess_id",
			CookieDomain:   domain,
			CookieSameSite: "Strict",
			CookieSecure:   true,
		})
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
