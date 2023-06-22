package Editor

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

const DateIta = "02-01-2006"

var now = time.Now()

func UpgradeOldFileCSV(Ofile [][]string) [][]string {
	Nfile := make([][]string, 0)
	for i := range Ofile {
		Nfile = append(Nfile, UpdateLine(Ofile[i]))
	}
	sort.Slice(Nfile,
		func(i, j int) bool {
			if Nfile[i][5] > Nfile[j][5] {
				Nfile[i][0] = strconv.Itoa(i)
				Nfile[j][0] = strconv.Itoa(i + 1)
				return true
			} else {
				Nfile[j][0] = strconv.Itoa(i)
				Nfile[i][0] = strconv.Itoa(i + 1)
				return false
			}
		},
	)
	return Nfile
}

func UpdateLine(str []string) []string {
	nline := make([]string, 0)
	nline = append(nline, "-1")
	nline = append(nline, str...)

	date, err := time.Parse(DateIta, nline[3])
	if err != nil {
		log.Fatal(err)
	}

	nline = append(nline, strconv.Itoa(GetAge(date)))
	nline = append(nline, date.Format(time.DateOnly))

	return nline
}

func GetAge(data time.Time) int {
	diff := now.Sub(data).Hours()
	return int(diff/24) / 365
}

func PrintMatrix(mtx [][]string) {
	str := ""
	for i := range mtx {
		str += fmt.Sprintf("\n %s, %s, %s, %s, %s", mtx[i][0], mtx[i][1], mtx[i][2], mtx[i][3], mtx[i][4])
	}
	fmt.Println(str)
}
