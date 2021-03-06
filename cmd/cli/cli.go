package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Network Lookup CLI"
	app.Usage = "Query IP's, CNAME's, MX records and Name-Servers"

	// Uses the same flag for all commands so can define it up here
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "ajb.app",
			// TODO: Change this to Maykind once domain transfer is done.
		},
	}

	// COMMANDS
	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks Up the NameServers for a Particular Host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("url"))
				if err != nil {
					return err
				}
				// Log the results to the console
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},

		{
			// net.LookupIP() returns a slice of IP addresses and as
			// such we iterate over these in order to print them out
			// in a nicer fashion
			Name:  "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},

		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(cname)
				return nil
			},
		},

		{
			Name:  "mx",
			Usage: "Looks up the MX records for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}

	// Run the app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
