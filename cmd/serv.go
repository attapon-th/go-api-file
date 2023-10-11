/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/spf13/cobra"
)

var (
	prefix  = "/"
	host    = "0.0.0.0"
	port    = 8888
	dirpath = "./assets"
	maxAge  = 3600
	browse  = true
)

// servCmd represents the serv command
var servCmd = &cobra.Command{
	Use:   "serv",
	Short: "start fiber app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		serv()
	},
}

func init() {
	rootCmd.AddCommand(servCmd)
	servCmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0", "Host")
	servCmd.Flags().IntVarP(&port, "port", "P", 80, "LitenerPort")
	servCmd.Flags().StringVarP(&prefix, "prefix", "p", "/", "URL PathPrefix")
	servCmd.Flags().StringVarP(&dirpath, "dirpath", "d", "./assets", "DirPath")
	servCmd.Flags().IntVarP(&maxAge, "max-age", "a", 3600, "Cashe Max-Age")
	servCmd.Flags().BoolVarP(&browse, "browse", "b", false, "Browse")
}

func serv() {
	os.MkdirAll(dirpath, os.ModePerm)
	app := fiber.New()
	app.Use(prefix, filesystem.New(filesystem.Config{
		Root:   http.Dir(dirpath),
		Browse: browse,
		MaxAge: maxAge,
		// PathPrefix: prefix,
	}))

	slog.Info("Configs", "host", host)
	slog.Info("Configs", "port", port)
	slog.Info("Configs", "prefix", prefix)
	slog.Info("Configs", "dirpath", dirpath)
	slog.Info("Configs", "maxAge", maxAge)
	slog.Info("Configs", "browse", browse)

	// slog.Info("Listen", "Serv", host+":"+fmt.Sprint(port))
	if err := app.Listen(host + ":" + fmt.Sprint(port)); err != nil {
		slog.Error("Listen failed", "Error", err.Error())
	}
}
