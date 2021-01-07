package controller

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

// Core - Core item
type Core struct {
	File    string `json:"file"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Date    string `json:"date"`
}

// Cores - Core list
type Cores struct {
	Cores []Core `json:"cores"`
}

// Game - game dir
type Game struct {
	Dir  string `json:"file"`
	Date string `json:"date"`
}

// Games - Games dir list
type Games struct {
	Games []Game `json:"games"`
}

func files(app *fiber.App) {
	f := app.Group("/files")

	// cores
	f.Get("/cores", func(c *fiber.Ctx) error {
		dir, err := ioutil.ReadDir("./files/cores")
		if err != nil {
			return err
		}
		coreList := Cores{}

		r := regexp.MustCompile(`\A([A-Za-z0-9\-]+)\_([0-9a-z]+)\.rbf\z`)
		for _, file := range dir {
			if r.MatchString(file.Name()) {
				fdata := r.FindSubmatch([]byte(file.Name()))
				coreList.Cores = append(coreList.Cores, Core{
					string(fdata[0]), string(fdata[1]), string(fdata[2]), file.ModTime().String(),
				})
			}
		}

		return c.JSON(coreList)
	})

	// core
	f.Get("/core/:name/:version", func(c *fiber.Ctx) error {
		dir, err := ioutil.ReadDir("./files/cores")
		if err != nil {
			return err
		}

		r := regexp.MustCompile(`\A(` + c.Params("name") + `)\_([0-9a-z]+)\.rbf\z`)
		for _, file := range dir {
			if r.MatchString(file.Name()) {
				fdata := r.FindSubmatch([]byte(file.Name()))
				if string(fdata[2]) == c.Params("version") {
					return c.SendStatus(fiber.StatusNotModified)
				}

				c.Set("Content-Disposition", "inline; filename=\""+file.Name()+"\"")
				return c.SendFile("./files/cores/"+file.Name(), true)
			}
		}

		return fiber.ErrNotFound
	})

	// games
	f.Get("/games", func(c *fiber.Ctx) error {
		dir, err := ioutil.ReadDir("./files/games")
		if err != nil {
			return err
		}

		gameList := Games{}
		for _, d := range dir {
			gameList.Games = append(gameList.Games, Game{d.Name(), d.ModTime().String()})
		}

		return c.JSON(gameList)
	})

	// games/game
	f.Get("/games/:game", func(c *fiber.Ctx) error {
		dir, err := ioutil.ReadDir("./files/games/" + c.Params("game"))
		if err != nil {
			return err
		}

		gameList := Games{}
		for _, d := range dir {
			gameList.Games = append(gameList.Games, Game{d.Name(), d.ModTime().String()})
		}

		return c.JSON(gameList)
	})

	// games/game/file
	f.Get("/games/:game/:file", func(c *fiber.Ctx) error {
		path := fmt.Sprintf("./files/games/%s/%s", c.Params("game"), c.Params("file"))
		if _, err := os.Stat(path); err == nil {
			c.Set("Content-Disposition", "inline; filename=\""+c.Params("file")+"\"")
			return c.SendFile(path, true)
		}

		return fiber.ErrNotFound
	})
}
