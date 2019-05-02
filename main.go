package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type HealthData struct {
	XMLName xml.Name `xml:"HealthData"`
	Record  []Record
}

type Record struct {
	Type          string `xml:"type,attr" csv:"type"`
	SourceName    string `xml:"sourceName,attr" csv:"-"`
	SourceVersion string `xml:"sourceVersion,attr" csv:"-"`
	Unit          string `xml:"unit,attr" csv:"unit"`
	CreationDate  string `xml:"creationDate,attr" csv:"-"`
	StartDate     string `xml:"startDate,attr" csv:"startDate"`
	EndDate       string `xml:"endDate,attr" csv:"endDate"`
	Value         string `xml:"value,attr" csv:"value"`
}

func main() {

	// open xml file.
	xmlFile, err := os.Open("./apple_health_export/export.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	// read xml data.
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
	}

	// parse xml and export to csv.
	csvFile, err := os.Create("./export.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	var healthData HealthData
	xml.Unmarshal(xmlData, &healthData)
	records := healthData.Record
	gocsv.MarshalFile(&records, csvFile)

}
