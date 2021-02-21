package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
	"github.com/wlwanpan/mcdata"
)

func main() {
	p := argparse.NewParser("mcdata", "Minecraft data go generator")

	genCmd := p.NewCommand("gen", "Generate the go struct for the provided version")
	e := genCmd.Selector("e", "edition", []string{mcdata.EditionPC, mcdata.EditionPE}, &argparse.Options{
		Required: false,
		Default:  mcdata.EditionPC,
		Help:     "Minecraft edition",
	})
	v := genCmd.String("v", "version", &argparse.Options{
		Required: true,
		Help:     "Minecraft version",
	})
	o := genCmd.String("o", "output", &argparse.Options{
		Required: true,
		Help:     "Destination dir of generated files.",
	})

	if err := p.Parse(os.Args); err != nil {
		log.Fatal(p.Usage(err))
	}

	switch {
	case genCmd.Happened():
		if err := mcdata.GenerateGoStructs("minecraft-data/data", *e, *v, *o); err != nil {
			fmt.Println(p.Usage(err))
		}
		fmt.Println("Successfully generated Minecraft data structs!")
	default:
		fmt.Println(p.Usage(nil))
	}
}
