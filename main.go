package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

var tvShows = []string{"I'm sure you gonna enjoy watching: "}

var listTvShows = []string{"Suits", "Hometown Cha Cha Cha", "The Boys", "The King's Affection", "Squid Game"}
var listAnimes = []string{"Attack on Titan", "Demon Slayer", "Jujutsu Kaisen", "Mushoku Tensei", "Kaguya Sama Love is War"}

func info() {
	app.Name = "TV shows / Anime CLI"
	app.Usage = "CLI to manage your TV shows and Anime"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		{
			Name:  "Felipe Lin",
			Email: "felipejtlin@gmail.com",
		},
	}
}

func flags() {
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "tvshows, t",
			Usage: "List all the tv shows",
		},
		cli.BoolFlag{
			Name:  "anime, a",
			Usage: "List all the anime",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("tvshows") {
			for _, show := range listTvShows {
				fmt.Println(show)
			}
		}
		if c.Bool("anime") {
			for _, anime := range listAnimes {
				fmt.Println(anime)
			}
		}
		return nil
	}

}

func switchCommand() {
	app.Commands = []cli.Command{
		{
			Name:    "tv",
			Aliases: []string{"t"},
			Usage:   "Watch tv shows",
			Subcommands: []cli.Command{
				{
					Name:    "Suits",
					Aliases: []string{"s"},
					Usage:   "Watch 'Suits' tv show",
					Action: func(c *cli.Context) {
						suits := "Suits"
						suitsShow := append(tvShows, suits)
						str := strings.Join(suitsShow, " ")
						fmt.Println(str)
					},
				},
				{
					Name:    "Hometown_Cha_Cha_Cha",
					Aliases: []string{"h"},
					Usage:   "Watch 'Hometown Cha Cha Cha' tv show",
					Action: func(c *cli.Context) {
						hw := "Hometown Cha Cha Cha"
						hwShow := append(tvShows, hw)
						str := strings.Join(hwShow, " ")
						fmt.Println(str)
					},
				},
				{
					Name:    "The_Boys",
					Aliases: []string{"t"},
					Usage:   "Watch 'The Boys' tv show",
					Action: func(c *cli.Context) {
						tb := "The Boys"
						tbShow := append(tvShows, tb)
						str := strings.Join(tbShow, " ")
						fmt.Println(str)
					},
				},
				{
					Name:    "The_King's_Affection",
					Aliases: []string{"k"},
					Usage:   "Watch 'The King's Affection' tv show",
					Action: func(c *cli.Context) {
						tka := "The King's Affection"
						tkaShow := append(tvShows, tka)
						str := strings.Join(tkaShow, " ")
						fmt.Println(str)
					},
				},
				{
					Name:    "Squid_Game",
					Aliases: []string{"sq"},
					Usage:   "Watch 'Squid Game' tv show",
					Action: func(c *cli.Context) {
						sq := "Squid Game"
						sqShow := append(tvShows, sq)
						str := strings.Join(sqShow, " ")
						fmt.Println(str)
					},
				},
			},
		},
		{
			Name:    "anime",
			Aliases: []string{"a"},
			Usage:   "Watch anime",
			Subcommands: []cli.Command{
				{
					Name:    "Attack_on_Titan",
					Aliases: []string{"at"},
					Usage:   "Watch 'Attack on Titan' anime",
					Action: func(c *cli.Context) {
						aot := "Attack on Titan"
						aotAnime := append(tvShows, aot)
						str := strings.Join(aotAnime, " ")
						fmt.Println(str)
					},
				},
				{
					Name:    "Demon_Slayer",
					Aliases: []string{"ds"},
					Usage:   "Watch 'Demon Slayer' anime",
					Action: func(c *cli.Context) {
						ds := "Demon Slayer"
						dsAnime := append(tvShows, ds)
						str := strings.Join(dsAnime, " ")
						fmt.Println(str)
					},
				},
				{
					Name:    "Jujutsu_Kaisen",
					Aliases: []string{"jk"},
					Usage:   "Watch 'Jujutsu Kaisen' anime",
					Action: func(c *cli.Context) {
						jk := "Jujutsu Kaisen"
						jkAnime := append(tvShows, jk)
						str := strings.Join(jkAnime, " ")
						fmt.Println(str)
					},
				},
				{
					Name:    "Mushoku_Tensei",
					Aliases: []string{"mt"},
					Usage:   "Watch 'Mushoku Tensei' anime",
					Action: func(c *cli.Context) {
						mt := "Mushoku Tensei"
						mtAnime := append(tvShows, mt)
						str := strings.Join(mtAnime, " ")
						fmt.Println(str)
					},
				},
				{
					Name:    "Kaguya_Sama_Love_is_War",
					Aliases: []string{"k"},
					Usage:   "Watch 'Kaguya Sama Love is War' anime",
					Action: func(c *cli.Context) {
						k := "Kaguya Sama Love is War"
						kAnime := append(tvShows, k)
						str := strings.Join(kAnime, " ")
						fmt.Println(str)
					},
				},
			},
		},
	}
}

func main() {
	info()
	switchCommand()
	flags()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
