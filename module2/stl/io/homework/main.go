// You can edit this code!
// Click here and start typing.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	data := []string{
		"there is 3pm, but im still alive to write this code snippet\n",
		"чистый код лучше, чем заумный код\n",
		"ваш код станет наследием будущих программистов\n",
		"задумайтесь об этом\n",
	}
	// здесь расположите буфер
	buff := new(bytes.Buffer)
	// запишите данные в буфер
	for _, i := range data {
		buff.WriteString(i)
	}
	// создайте файл
	file, er := os.Create("./test.txt")

	defer func() { file.Close() }()

	if er != nil {
		panic(er)
	}
	// запишите данные в файл
	io.Copy(file, buff)
	// прочтите данные в новый буфе
	f, _ := os.Open("./test.txt")
	// p := make([]byte, 4)
	buff1 := new(bytes.Buffer)

	io.Copy(buff1, f)
	fmt.Print(buff1.String())
	// у вас все получится!
}
