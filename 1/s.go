package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)


  type Record struct {
	  Name string
	  Surname string
	  Number int
	  LastAccess string
  }

var myData = []Record{}


  func readCSVfile(filepath string) ([]string , error) {
	  _, err := os.Stat(filepath)
	  if err != nil {
		  return nil, err
	  }
	  file, err := os.Open(filepath)
	  if err != nil {
		  return nil, err
	  }
	  defer file.Close()
	  reader , err := csv.NewReader(file).ReadAll()
	  if err != nil {
		  return nil, err
	  }
	  return [][]string{}

  }


func main() {


}
