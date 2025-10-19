package main

import (
	"encoding/csv"
	"os"
)



func loadRecipient(filepath string, ch chan Receiver) error{
	defer close(ch)
	f, err := os.Open(filepath)
	if err != nil{
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err :=r.ReadAll()
	if err != nil{
		return err
	}

	for _ ,record :=range records[1:]{
		ch <- Receiver{
			Name : record[0],
			Email: record[1],
		}	
	}
	return nil
}
