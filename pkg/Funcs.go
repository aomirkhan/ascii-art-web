package pkg

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strings"
)

func Start(s string) []string {
	List := strings.Split(s, "\n")
	b := 0
	var alph []string
	for i := 1; len(List) > i; i++ {
		if i%9 == 0 {
			alph = append(alph, strings.Join(List[b:i+1], "\n"))
			b = i + 1
		}
	}
	alph[0] = "      \n      \n      \n      \n      \n      \n      \n"
	return alph
}

func Print(s string, alph []string) string {
	var obj [][]string
	var buf bytes.Buffer
	for i := range s {
		for j := 32; 126 > j; j++ {
			if rune(s[i]) == rune(j) {
				obj = append(obj, strings.Split(alph[j-32], "\n"))
			}
		}
	}
	ind := 0
	for ind < 8 {
		var line string
		for i := range obj {
			line = line + strings.ReplaceAll(obj[i][ind], "\n", "")
		}
		fmt.Fprintln(&buf, line)
		ind++
	}

	return buf.String()
}

func ReadFile(fileName string) string {
	body, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(body)
}

func Md5(s string) string {
	h := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", h)
}
