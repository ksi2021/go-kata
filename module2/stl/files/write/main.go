package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeToFile(data, filePATH string) {
	file, er := os.Create(filePATH)
	check(er)
	defer file.Close()
	file.WriteString(data)
	// file.Write([]byte(data))
}

var RuTransiltMap = map[rune]string{
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
	'д': "d",
	'е': "e",
	'ё': "yo",
	'ж': "zh",
	'з': "z",
	'и': "i",
	'й': "j",
	'к': "k",
	'л': "l",
	'м': "m",
	'н': "n",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'у': "u",
	'ф': "f",
	'х': "h",
	'ц': "c",
	'ч': "ch",
	'ш': "sh",
	'щ': "sch",
	'ъ': "'",
	'ы': "y",
	'ь': "",
	'э': "e",
	'ю': "ju",
	'я': "ja",
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s\n", name)

	// запись имени в файл
	writeToFile(name, "./tmp/name.txt")

	file, err := os.Open("./tmp/example.txt")
	check(err)
	data := make([]byte, 64)
	var text string
	for {
		_, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		text += string(data) // сохранение прочитаного в переменную
	}

	for k, v := range RuTransiltMap {
		text = strings.ReplaceAll(text, string(k), v)
		text = strings.ReplaceAll(text, strings.ToUpper(string(k)), strings.ToTitle(v))
	}

	writeToFile(text, "./tmp/example.processed.txt")

}
