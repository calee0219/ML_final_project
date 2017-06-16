package main

import (
	"encoding/csv"
	"os"
	"strings"
	"strconv"
	"fmt"
	"github.com/montanaflynn/stats"
)

func main() {
	in, _ := os.Open("loans_data.csv")
	out, _ := os.OpenFile("loans_data2.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer out.Close()
	r := csv.NewReader(in)
	w := csv.NewWriter(out)
	// build title
	data, _ := r.ReadAll()
	table := make([][]string, 0)
	row := make([]string, 0)
	title := make(map[string]int)
	// -> 0
	value := make(map[string]bool)
	tmpTitle := data[0][0]
	for _, v := range data[1:] {
		if !value[v[0]] {
			value[v[0]] = true
			row = append(row, tmpTitle+"_"+strings.Replace(v[0], " ", "_", -1))
			title[v[0]] = len(row) - 1
		}
	}
	// -> 4
	tmpTitle = data[0][4]
	for _, v := range data[1:] {
		if !value[v[4]] {
			value[v[4]] = true
			row = append(row, tmpTitle+"_"+strings.Replace(v[4], " ", "_", -1))
			title[v[4]] = len(row) - 1
		}
	}
	// -> 6
	tmpTitle = data[0][6]
	for _, v := range data[1:] {
		if !value[v[6]] {
			value[v[6]] = true
			row = append(row, tmpTitle+"_"+strings.Replace(v[6], " ", "_", -1))
			title[v[6]] = len(row) - 1
		}
	}
	// other
	for k, v := range data[0] {
		if k != 0 && k != 4 && k != 6 {
			row = append(row, v)
			title[v] = len(row) - 1
		}
	}
	table = append(table, row)
	// build table
	for _, rowData := range data[1:] {
		row = make([]string, len(title))
		for k := range row {
			row[k] = "0"
		}
		for k, v := range rowData {
			if k == 0 || k == 4 || k == 6 {
				row[title[v]] = "1"
			} else {
				row[title[data[0][k]]] = v
			}
		}
		table = append(table, row)
	}
	// normalize
        for k := range table[0][:len(table[0])-1]{
          col := make([]float64, 0)
          for i := 1 ; i<len(table) ; i++ {
            tmpF, err := strconv.ParseFloat(table[i][k], 64)
            if err != nil {
              fmt.Println(err)
              os.Exit(1)
            }
            col = append(col, tmpF)
          }
          mean, err := stats.Mean(col)
          if err != nil {
            fmt.Println(err)
            os.Exit(1)
          }
          stderv, err := stats.StandardDeviationPopulation(col)
          if err != nil {
            fmt.Println(err)
            os.Exit(1)
          }
          for i := 1 ; i<len(table) ; i++ {
            table[i][k] = fmt.Sprintf("%f", (col[i-1]-mean)/stderv)
          }
        }
	w.WriteAll(table)
	w.Flush()
}

// 0 4 6
