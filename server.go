package main

import (
	"fmt"
	"log"
	"path/filepath"

	"git.martianoids.com/queru/retroserver/internal/banner"
	"git.martianoids.com/queru/retroserver/internal/build"
	"git.martianoids.com/queru/retroserver/internal/cfg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gopkg.in/alecthomas/kingpin.v2"
)

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
	cfg.Main.Server.Concurrency = kingpin.Flag("concurrency", "Maximum number of concurrent connections in MiB").
		Default("256").Int()

	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version(cfg.Version).Author(cfg.Author)
	kingpin.CommandLine.Help = "Web Application Server"
	kingpin.Parse()

	if *cfg.Main.Env == "dev" {
		log.Println("SERVER/ENV", "Development mode ON")
	} else {
		log.Println("SERVER/ENV", "Production mode ON")
	}
	log.Println(build.Version())

	// root path
	cfg.Main.Root = filepath.Dir(".")

	// app and configuration
	app := fiber.New(
		fiber.Config{
			ReadTimeout:  *cfg.Main.Server.RTimeout,
			WriteTimeout: *cfg.Main.Server.WTimeout,
			BodyLimit:    *cfg.Main.Server.BodyLimitMb * 1024 * 1024,
			Concurrency:  *cfg.Main.Server.Concurrency * 1024,
			ServerHeader: "RetroServer_" + cfg.Version,
		},
	)

	// compression
	if *cfg.Main.Env == "prod" {
		app.Use(compress.New())
	}

	// logger
	app.Use(logger.New())

	// cors
	app.Use(cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(banner.Title)
	})

	app.Get("/server/version", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf(
			"+ SERVER VERSION:\n\n- %s\n- %s\n- %s\n- %s\n",
			build.Version(),
			build.VersionShort(),
			build.BinVersion(),
			build.CompileTime(),
		))
	})

	// recover from panic
	app.Use(recover.New())

	// 404 Not found
	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString(banner.Error404)
	})

	// server UP
	app.Listen(*cfg.Main.Server.IP + ":" + *cfg.Main.Server.Port)
}
