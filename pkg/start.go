package pkg

import (
	"fmt"
	"log"
	"strings"
)

func Converter(banner string, text string) string {
	// err1 := "Needed 1 argument, instead we have  "
	// err2 := "Needed characters related to ASCII"
	err3 := "Wrong file!"
	var final string
	k := ReadFile("pkg/banners/" + banner + ".txt")
	k2 := Md5(k)
	hashShadow := "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hashThinkertoy := "86d9947457f6a41a18cb98427e314ff8"
	if !(k2 == hashShadow || k2 == hashStandard || k2 == hashThinkertoy) {
		log.Fatal(err3)
	}
	fmt.Println(text)
	x := strings.Split(text, "\r\n")

	k = strings.ReplaceAll(string(k), "\r", "")
	mapa := Start(k)

	for i := range x {
		final = final + Print(x[i], mapa)
	}

	return final
	// Funcs.PrintEverything(args, mapa)
}
