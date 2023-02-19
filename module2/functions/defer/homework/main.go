package main

import (
	"errors"
	"fmt"
)

// MergeDictsJob is a job to merge dictionaries into a single dictionary.
type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

// errors
var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

// BEGIN (write your solution here)
func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {

	job.IsFinished = true
	job.Merged = make(map[string]string)
	// проверка длины Dicts
	if len(job.Dicts) < 2 {
		return job, errNotEnoughDicts
	}

	// проверка на пустую мапу в Dicts
	for i := range job.Dicts {

		if job.Dicts[i] == nil {
			return job, errNilDict
		}

		// если все нормально добавляем результаты в Merged
		for k, v := range job.Dicts[i] {
			job.Merged[k] = v
		}
	}

	return job, nil
}

// END

// Пример работы
func main() {
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{}))                                                   // &MergeDictsJob{IsFinished: true}, "at least 2 dictionaries are required"
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"}, nil}}))        // &MergeDictsJob{IsFinished: true, Dicts: []map[string]string{{"a": "b"},nil}}, "nil dictionary"
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"}, {"b": "c"}}})) // &MergeDictsJob{IsFinished: true, Dicts: []map[string]string{{"a": "b", "b": "c"}}}, nil

}
