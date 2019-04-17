package main

import (
	"body1/scripts/lib"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("iis.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	Sites := []lib.Site{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := regexp.MustCompile(`SITE \"(?P<name>[\w\s\-\.\d]+)\" \(id:(?P<id>\d+),bindings:(?P<bindings>[http\/\*\:\d\:\w\.\,]+),state:\w+`)
		match := r.FindStringSubmatch(strings.TrimSpace(scanner.Text()))
		// SITE "Medtech1" (id:135545492,bindings:http/:80:medtech1.com,http/*:80:web5.medtech1.com,state:Started)
		if len(match) != 0 {
			site := lib.Site{}
			site.Name = match[1]
			i, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			site.ID = i
			for _, bind := range strings.Split(match[3], ",") {
				reg := regexp.MustCompile(`http\/\*?\:(?P<port>\d+)\:(?P<domain>[\w\d\-\.]+)`)
				m := reg.FindStringSubmatch(bind)

				if len(m) == 0 {
					site.Bindings = append(site.Bindings, "localhost")
				} else {
					// fmt.Println(m[0])
					site.Bindings = append(site.Bindings, m[2])
				}
			}
			// fmt.Println(site)
			Sites = append(Sites, site)
		}
	}

	f, err := os.Open("iis.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		// fmt.Println
		id, err := strconv.Atoi(record[1])
		if err != nil {
			continue
		}
		for i := range Sites {
			if Sites[i].ID == id {
				Sites[i].Path = record[14]
				Sites[i].Proto = record[13]
			}
		}
	}
	fmt.Println("Location,Site Name,Aliases")
	for _, s := range Sites {
		fmt.Println(s)
	}
}
