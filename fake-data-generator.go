package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"

	"github.com/jaswdr/faker"
)

func main() {
	fake := faker.New()

	reportingDates := []string{"2022-01-01", "2022-02-01", "2022-03-01", "2022-04-01", "2022-05-01", "2022-06-01", "2022-07-01", "2022-08-01", "2022-09-01", "2022-10-01", "2022-11-01", "2022-12-01", "2023-01-01"}
	headers := []string{"unique_key", "reporting_date", "country", "plan", "organisation_id", "mrr"}
	pricings := []string{"30", "35", "50"}
	organisationIds := []string{}
	for i := 0; i < 100; i++ {
		organisationIds = append(organisationIds, fake.UUID().V4())
	}

	empData := [][]string{}

	for _, date := range reportingDates {
		for i := 0; i < 10; i++ {
			organisationIds = append(organisationIds, fake.UUID().V4())
		}
		for _, organisationId := range organisationIds {
			randomIndex := rand.Intn(len(pricings))
			pricing := pricings[randomIndex]

			empData = append(empData, []string{date + "__" + organisationId, date, "FR", "PREMIUM1", organisationId, pricing})
		}
	}

	sort.SliceStable(empData, func(i, j int) bool {
		return strings.Compare(empData[i][0], empData[j][0]) < 0
	})

	if _, err := os.Stat("mart/organisation-mrr.csv"); os.IsNotExist(err) {
		os.MkdirAll("mart", 0700)
	}

	csvFile, err := os.Create("mart/organisation-mrr.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)

	_ = csvwriter.Write(headers)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}

	csvwriter.Flush()
	csvFile.Close()
}
