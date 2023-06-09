package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/xv-anx/dtextra/modules"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var judge bool
var cheat *modules.CheatData

func main() {
	cheats := modules.NewCheats()

	dir := "./cheats"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	filePath := "./cheats/code.txt"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer file.Close()
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer file.Close()

		sc := bufio.NewScanner(file)
		for sc.Scan() {
			lines := sc.Text()
			if lines == "" {
				continue
			}
			line := strings.Split(lines, "\n")

			for _, l := range line {
				if strings.HasPrefix(l, "[") && strings.HasSuffix(l, "]") {
					funcName := strings.TrimPrefix(strings.TrimSuffix(l, "]"), "[")
					funcName = strings.Map(func(r rune) rune {
						if unicode.IsSpace(r) || unicode.IsPunct(r) {
							return '_'
						}
						return r
					}, funcName)
					words := strings.Split(funcName, "_")
					for i, word := range words {
						if len(word) > 0 {
							if i == 0 {
								words[i] = strings.ToLower(word)
							} else {
								words[i] = cases.Title(language.Und).String(word)
							}
						}
					}
					funcName = strings.Join(words, "")
					cheat = cheats.AddCheatData(funcName)
					continue
				}

				if l == "!%" {
					judge = true
					continue
				}
				if judge && l != "%!" {
					parts := strings.Fields(l)
					if len(parts) == 2 {
						address := "0x" + strings.TrimPrefix(parts[0], "0")
						data := "0x" + strings.ReplaceAll(parts[1], ",", "")
						if cheat == nil {
							fmt.Fprintln(os.Stderr, "Cheat data not found")
							return
						}
						cheats.AddReadonlyData(address, data)
					}

					if judge && strings.HasPrefix(l, "{") {
						note := strings.TrimPrefix(strings.TrimSuffix(l, "}"), "{")
						if cheat != nil {
							cheat.Note = note
						}
					}

				} else if !judge && l != "!%" {
					parts := strings.Fields(l)
					if len(parts) == 2 {
						address := strings.TrimPrefix(parts[0], "0x")
						if len(address) < 8 {
							address = "0" + address
						}
						if cheat == nil {
							fmt.Fprintln(os.Stderr, "Cheat data not found")
							return
						}
						if cheat.StartAddr == "" {
							cheat.StartAddr = "0x" + address
						}
						dataParts := strings.Fields(parts[1])
						var data string
						if len(dataParts) == 1 {
							data = "0x" + dataParts[0]
						} else if len(dataParts) >= 2 {
							data = "0x" + dataParts[1]
						}
						cheat.Values = append(cheat.Values, data)
					}

				} else if l == "%!" {
					judge = false
					continue
				}
				if cheat == nil {
					fmt.Fprintln(os.Stderr, "invalid file format: Cheats not found")
					return
				}
			}
		}

		if err := sc.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error while reading file: ", err)
			return
		}

		for _, cheat := range cheats.Data {
			cheat.PrintCheatData()
		}
	}
}
