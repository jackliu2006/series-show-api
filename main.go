package main

import (
	//"fmt"
	"SeriesShowWeb/model"
	"SeriesShowWeb/services"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	xmlbody := model.XMLBody{}
	var jsonSeries []model.Series
	var jsonEpics []model.Episode

	files, err := ioutil.ReadDir("./xml")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		
		byteValue, err := ioutil.ReadFile("./xml/" + f.Name())
		if err != nil {
			log.Println("error while reading file", err)
		}
		//log.Println("1: ",string(byteValue))

		xml.Unmarshal(byteValue, &xmlbody)
		//log.Println("epic before loop=", xmlbody.Episode)

		jsonSeries = append(jsonSeries, xmlbody.Series)

		for _, e := range xmlbody.Episode {
			//log.Println("epic in loop=",e)
			jsonEpics = append(jsonEpics, e)
		}

		log.Println("2: ", len(jsonSeries), len(jsonEpics))

	}

	//save to json file
	jsonSeriesByte, err := json.Marshal(jsonSeries)
	if err != nil {
		log.Println("error while marshal", err)
	}
	jsonEpicsByte, err := json.Marshal(jsonEpics)
	if err != nil {
		log.Println("error while marshal", err)
	}
	

	os.Remove("series.json")
	os.Remove("epics.json")
	
	ioutil.WriteFile("series.json", jsonSeriesByte, 0644)
	ioutil.WriteFile("epics.json", jsonEpicsByte, 0644)

	//start server

	router := gin.New()

	router.GET("/series", services.GetSeriesList)
	router.GET("/series/:seriesID/epics", services.GetEpicsBySeriesID)

	router.Run("localhost:8080")

}
