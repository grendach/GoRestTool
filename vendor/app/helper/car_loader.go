package helper

import (
	"app/model"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// LoadCarsCsv function that reads conftent of CSV file
func LoadCarsCsv() []model.Car {
	var cars []model.Car

	// Producent, Model, Year, Colour
	pwd, _ := os.Getwd()
	csvFile, _ := os.Open(pwd + "/db/car-list.csv")
	println(pwd + "/../../db/car-list.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		year, err := strconv.Atoi(line[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		cars = append(cars, model.Car{
			Producent: line[0],
			Model:     line[1],
			Colour:    line[2],
			Year:      year,
		})
	}

	return cars
}
