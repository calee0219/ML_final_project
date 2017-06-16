package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("loans_data.csv")
	out, _ := os.OpenFile("loans_data2.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer out.Close()
	r := csv.NewReader(in)
	w := csv.NewWriter(out)
	tmp, _ := r.Read()
	w.Write(tmp)
	tmpA, _ := r.ReadAll()
	index0 := 1
	m0 := make(map[string]int)
	index4 := 1
	m4 := make(map[string]int)
	index6 := 1
	m6 := make(map[string]int)
	for _, tmp = range tmpA {
		if _, ok := m0[tmp[0]]; !ok {
			m0[tmp[0]] = index0 * 10
			index0++
		}
		tmp[0] = fmt.Sprintf("%d", m0[tmp[0]])

		if _, ok := m4[tmp[4]]; !ok {
			m4[tmp[4]] = index4 * 10
			index4++
		}
		tmp[4] = fmt.Sprintf("%d", m4[tmp[4]])

		if _, ok := m6[tmp[6]]; !ok {
			m6[tmp[6]] = index6 * 10
			index6++
		}
		tmp[6] = fmt.Sprintf("%d", m6[tmp[6]])
		w.Write(tmp)
	}
	w.Flush()
}
