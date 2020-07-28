package services

import (
	//"io/ioutil"
	"net/http"

	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gojsonq"
	"log"
)

func GetSeriesList(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")


	jq := gojsonq.New().File("./series.json")

	//log.Println("jquery serach result in getEpcicbyid",jq.Get() )

	c.JSON(http.StatusOK, gin.H{"Series-list":jq.SortBy("SeriesID").Get()})


}

func GetEpicsBySeriesID(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	seriesID := c.Param("seriesID")

	log.Print("input seriesID = ", seriesID)

	jq := gojsonq.New().File("./epics.json").Where("seriesid", "=", seriesID)

	//log.Println("jquery serach result in getEpcicbyid",jq.Get() )

	c.JSON(http.StatusOK, gin.H{"Epics-list":jq.SortBy("EpisodeNumber").Get()})
}
