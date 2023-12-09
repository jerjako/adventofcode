package main

import (
	"bytes"
	"fmt"
	"log"
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

		if err := cmd.Run(); err != nil {
			panic(err)
		}
		if cmdErr.String() != "" {
			panic(cmdErr.String())
		}

		log.Print(cmdOut.String())

		utils.LogColor(partColor[part], fmt.Sprintf("End part %d after: %s", part, time.Since(t)))
	}
}
