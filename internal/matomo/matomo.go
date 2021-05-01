package matomo

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

const MURL = "https://stats.martianoids.com/matomo.php"

// Visit
type Visit struct {
	IDSite     uint8  `json:"idsite"`
	Rec        uint8  `json:"rec"`
	ActionName string `json:"action_name"`
	URL        string `json:"url"`
	ID         string `json:"_id"`
	Rand       uint64 `json:"rand"`
	Version    uint8  `json:"apiv"`
	Ref        string `json:"urlref"`
	Agent      string `json:"ua"`
	Lang       string `json:"lang"`
}

// NewVisit - Sends a visit to statistics server
func NewVisit(c *fiber.Ctx, store *session.Store) {
	v := new(Visit)
	v.IDSite = 25
	v.Rec = 1
	v.ActionName = ""
	v.URL = string(c.Request().URI().FullURI())

	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	defer sess.Save()

	if sess.Get("uuid") == nil {
		v.ID = utils.UUID()
		sess.Set("uuid", v.ID)
	} else {
		v.ID = sess.Get("uuid").(string)
	}

	v.Rand = rand.Uint64()
	v.Version = 1
	v.Ref = string(c.Context().Referer())
	v.Agent = c.Get("User-Agent")
	v.Lang = c.Get("Accept-Language")

	go v.send()
}

func (v *Visit) send() {
	query := fmt.Sprintf("idsite=%v&rec=%v&action_name=%s&url=%s&_id=%s&rand=%v&apiv=%v"+
		"&urlref=%s&ua=%s&lang=%s",
		v.IDSite, v.Rec, v.ActionName, v.URL, v.ID, v.Rand, v.Version,
		v.Ref, v.Agent, v.Lang,
	)

	_, err := http.Get(MURL + "?" + query)
	if err != nil {
		log.Println("VISIT REC ERROR:", err)
	}
}
