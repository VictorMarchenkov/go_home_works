package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

type customError struct {
	Timestamp string
	FileName  string
	Err       error
}

func NewCustomError(name string, t string, e error) *customError {
	return &customError{
		Timestamp: t,
		FileName:  name,
		Err:       e,
	}
}

func (ve *customError) Error() string {
	return fmt.Sprintf("%s at moment: %s with file: %s", ve.Err, ve.Timestamp, ve.FileName)
}

func main() {
	var (
		directory = "test"
		nFiles    = 10 //1000000
	)

	err := os.Mkdir(directory, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	clearDirectory(directory)
	createFiles(directory, nFiles)
}

// clearDirectory cleaning directory
func clearDirectory(directory string) {
	dir, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Printf("%v:", err)
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{directory, d.Name()}...))
	}
}

// createFiles generate new files.
func createFiles(directory string, n int) {
	var fName, path string
	_, err := os.ReadDir(path)
	if err != nil {
		err = os.Mkdir(path, 0750)
		if err != nil && !os.IsExist(err) {
			//	log.Fatal(err)
			fmt.Printf("сan't create a directory %v", err)
		}
	}

	fakePanicSlice := [2]int{0, 0}
	for i := 0; i < n; i++ {
		t := time.Now().Format("2006-01-02 15:04:05,000")
		fName = fmt.Sprintf("file%d.txt", i)
		path = fmt.Sprintf("%s/%s", directory, fName)
		_, err := os.Create(path)
		if err != nil {
			fmt.Printf("\nerror creating file %s err:%s\n", fName, err)
		}
		go func() {
			defer func() {
				if v := recover(); v != nil {
					//fmt.Printf("recover after panic - : %s\n", v)
					err = NewCustomError(fName, t, fmt.Errorf("there was an error"))
					fmt.Println(err)
				}
			}()
			if i == 3 { // create conditions for implicit panic
				fmt.Println(fakePanicSlice[i])
			}

		}()
	}
}
