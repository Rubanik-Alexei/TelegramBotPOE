package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

func SearchRecipes(item string) []string {
	result := make([]string, 0)
	file, err := os.Open("recipes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	startInd := 0
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "-") {
			startInd = i
		}
		if !strings.Contains(lines[i], item) {
			continue
		} else {
			tmp := ""
			for j := startInd + 1; j <= i+1; j++ {
				tmp += lines[j] + "\n"
				// fmt.Println(lines[j])
			}
			fmt.Println(tmp)
			result = append(result, tmp)
		}
	}
	return result
}

func GetInfo() {
	response := make([]string, 0)
	//resp, err := soup.Get("https://pathofexile.fandom.com/ru/wiki/%D0%A0%D0%B5%D1%86%D0%B5%D0%BF%D1%82%D1%8B_%D1%82%D0%BE%D1%80%D0%B3%D0%BE%D0%B2%D1%86%D0%B5%D0%B2")
	resp, err := soup.Get("https://pathofexile.fandom.com/ru/wiki/%D0%9F%D1%80%D0%BE%D1%80%D0%BE%D1%87%D0%B5%D1%81%D1%82%D0%B2%D0%BE")
	if err != nil {
		response = append(response, "POESite is not respond")
		//return
	}
	doc := soup.HTMLParse(resp)
	tt := fmt.Sprint(doc.FullText())
	// s := strings.Replace(tt, "\n\n", "\n", -1)
	// regex, err := regexp.Compile("\n\n")
	// if err != nil {
	// 	return
	// }
	// for regex.MatchString(tt) {
	// 	tt = regex.ReplaceAllString(tt, "\n")
	// }
	err = os.WriteFile("output.txt", []byte(tt), 0644)
	if err != nil {
		panic(err)
	}
	//fmt.Println(tt)
}
