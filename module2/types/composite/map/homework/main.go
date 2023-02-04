package main

import "fmt"

type Project struct {
	Name  string
	Stars int
}

func main() {
	projects := []Project{
		{
			Name:  "https://github.com/docker/compose",
			Stars: 27600,
		},
		{
			Name:  "https://github.com/docker/compose1",
			Stars: 657600,
		},
		{
			Name:  "https://github.com/docker/compose2",
			Stars: 223100,
		},
		{
			Name:  "https://github.com/docker/compose3",
			Stars: 435600,
		},
		{
			Name:  "https://github.com/docker/compose4",
			Stars: 12345678,
		},
		{
			Name:  "https://github.com/docker/compose5",
			Stars: 33334444,
		},
		{
			Name:  "https://github.com/docker/compose6",
			Stars: 66661111,
		},
		{
			Name:  "https://github.com/docker/compose7",
			Stars: 55553433,
		},
		{
			Name:  "https://github.com/docker/compose8",
			Stars: 2325455,
		},
		{
			Name:  "https://github.com/docker/compose9",
			Stars: 111111,
		},
		{
			Name:  "https://github.com/docker/compose10",
			Stars: 957356346,
		},
		{
			Name:  "https://github.com/docker/compose11",
			Stars: 5675673,
		},
		{
			Name:  "https://github.com/docker/compose12",
			Stars: 211111,
		},
		// сюда впишите ваши остальные 12 структур
	}
	test := make(map[string]int)
	_ = test

	// в цикле запишите в map
	for _, i := range projects {
		test[i.Name] = i.Stars
	}
	// в цикле пройдитесь по мапе и выведите значения в консоль

	for k, v := range test {
		fmt.Println(k, v)
	}
}
