package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"git.martianoids.com/queru/retroserver/internal/banner"
	"git.martianoids.com/queru/retroserver/internal/build"
	"git.martianoids.com/queru/retroserver/internal/cfg"
	"git.martianoids.com/queru/retroserver/internal/controller"
	"git.martianoids.com/queru/retroserver/internal/matomo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/ace"
	"gopkg.in/alecthomas/kingpin.v2"
)

//go:embed asset
var asset embed.FS

func main() {
	// command line flags and params
	cfg.Main.Env = kingpin.Flag("environment", "dev or prod mode").
		Short('e').Default("prod").String()
	cfg.Main.Server.IP = kingpin.Flag("ip", "IP to listen").
		Short('i').Default("127.0.0.1").String()
	cfg.Main.Server.Port = kingpin.Flag("port", "Port to listen").
		Short('p').Default("8080").String()
	cfg.Main.Server.BodyLimitMb = kingpin.Flag("body-limit", "Body limit in MiB").
		Default("4").Int()
	cfg.Main.Server.RTimeout = kingpin.Flag("read-timeout", "Read timeout").
		Short('r').Default("10s").Duration()
	cfg.Main.Server.WTimeout = kingpin.Flag("write-timeout", "Write timeout").
		Short('w').Default("10s").Duration()
	cfg.Main.Server.Concurrency = kingpin.Flag("concurrency",
		"Maximum number of concurrent connections in MiB").Default("256").Int()

	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(cfg.Version).Author(cfg.Author)
	kingpin.CommandLine.Help = "Web Application Server"
	kingpin.Parse()

	// root path
	cfg.Main.Root = filepath.Dir(".")

	// template engine
	engine := ace.New("./views", ".ace")

	// app and configuration
	app := fiber.New(
		fiber.Config{
			ReadTimeout:           *cfg.Main.Server.RTimeout,
			WriteTimeout:          *cfg.Main.Server.WTimeout,
			BodyLimit:             *cfg.Main.Server.BodyLimitMb * 1024 * 1024,
			Concurrency:           *cfg.Main.Server.Concurrency * 1024,
			ServerHeader:          "RetroServer_" + cfg.Version,
			DisableStartupMessage: true,
			Views:                 engine,
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError
				txt := banner.Error500
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}
				if code == 404 {
					txt = banner.Error404
				}
				err = ctx.Status(code).SendString(
					banner.Title + "\n" + txt + banner.Separator + err.Error() + banner.Separator)
				if err != nil {
					return ctx.Status(500).SendString("Internal Server Error")
				}

				return nil
			},
		},
	)

	// session
	store := session.New(session.Config{Expiration: 8640 * time.Hour})
	// stats
	if cfg.IsProd() {
		app.Use(func(c *fiber.Ctx) error {
			matomo.NewVisit(c, store)
			return c.Next()
		})
	}

	// compression
	if *cfg.Main.Env == "prod" {
		app.Use(compress.New(compress.Config{Level: 0}))
	}

	// logger
	app.Use(logger.New())

	// cors
	app.Use(cors.New())

	// routes
	controller.Setup(app)

	// static assets
	if cfg.IsDev() {
		app.Static("/asset", "asset")
	} else {
		subFS, _ := fs.Sub(asset, "asset")
		app.Use("/asset", filesystem.New(filesystem.Config{
			Root:         http.FS(subFS),
			NotFoundFile: "Static file not found",
		}))
	}

	// recover from panic
	if !cfg.IsDev() {
		app.Use(recover.New())
	}

	// startup banner and info
	log.Println(banner.Console)
	if *cfg.Main.Env == "dev" {
		log.Println("SERVER/ENV", "Development mode ON")
	} else {
		log.Println("SERVER/ENV", "Production mode ON")
	}
	log.Println(build.Version())
	log.Println("Listening in", *cfg.Main.Server.IP+":"+*cfg.Main.Server.Port)

	// server UP
	app.Listen(*cfg.Main.Server.IP + ":" + *cfg.Main.Server.Port)
}
