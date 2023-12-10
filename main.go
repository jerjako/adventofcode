package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/jerjako/adventofcode/utils"
)

var partColor = map[int]utils.Color{
	1: utils.Green,
	2: utils.Cyan,
}

func main() {
	now := time.Now().UTC()
	config := utils.ParseConfig()

	if config.Init {
		utils.LogColor(utils.Red, fmt.Sprintf("Create %d/%d", config.Year, config.Day))

		_, err := os.Stat(fmt.Sprintf("./%d/%d/", config.Year, config.Day))
		if err == nil {
			panic("Folder already exist")
		} else if !os.IsNotExist(err) {
			panic(err)
		}

		cmd := exec.Command("cp", "--recursive", "./day_example", fmt.Sprintf("./%d/%d", config.Year, config.Day))

		var cmdOut, cmdErr bytes.Buffer
		cmd.Stdout = &cmdOut
		cmd.Stderr = &cmdErr
		err = cmd.Run()
		if cmdErr.String() != "" {
			panic(cmdErr.String())
		}
		if err != nil {
			panic(err)
		}

		utils.LogColor(utils.Red, fmt.Sprintf("Created %s", cmdOut.String()))

		os.Exit(1)
	}

	utils.LogColor(utils.Purple, fmt.Sprintf("Running day %d of year %d", config.Day, config.Year))
	defer func() { utils.LogColor(utils.Purple, fmt.Sprintf("End day %d after: %s", config.Day, time.Since(now))) }()

	parts := []int{}

	if config.Part1 {
		parts = append(parts, 1)
	}
	if config.Part2 {
		parts = append(parts, 2)
	}

	for _, part := range parts {
		utils.LogColor(partColor[part], fmt.Sprintf("Running part %d", part))
		t := time.Now().UTC()

		cmd := exec.Command("go", "run", fmt.Sprintf("./%d/%d/main.go", config.Year, config.Day))
		cmd.Args = append(cmd.Args, []string{
			fmt.Sprintf("-year=%d", config.Year),
			fmt.Sprintf("-day=%d", config.Day),
			fmt.Sprintf("-part%d", part),
		}...)
		if config.TestInput {
			cmd.Args = append(cmd.Args, "-test")
		}
		var cmdOut, cmdErr bytes.Buffer
		cmd.Stdout = &cmdOut
		cmd.Stderr = &cmdErr
		err := cmd.Run()
		if cmdErr.String() != "" {
			panic(cmdErr.String())
		}
		if err != nil {
			panic(err)
		}

		log.Print(cmdOut.String())

		utils.LogColor(partColor[part], fmt.Sprintf("End part %d after: %s", part, time.Since(t)))
	}
}
