package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mehtadhruv2104/searchEngine/backend/models"
	"github.com/mehtadhruv2104/searchEngine/backend/search"
)

type Handler struct {
	Records []models.LogEntry
	Index   search.Index
}

func NewHandler(records []models.LogEntry, index search.Index) *Handler {
	return &Handler{
		Records: records,
		Index:   index,
	}
}

func checkEmptyString(s *string) bool {
    if s == nil {
        return true
    }else if(*s == ""){
		return true
	}
    return false
}

func validateRequest(req models.SearchRequest)(bool){
	if(req.Query == "" &&  checkEmptyString(req.AppName) && checkEmptyString(req.MsgID) && checkEmptyString(req.Namespace) && checkEmptyString(req.Severity)){
		return false
	}
	return true
}


func (h Handler)HandleSearch(c *gin.Context){

	var req models.SearchRequest
	resp := models.SearchResponse{}
	err := c.BindJSON(&req)
	if err != nil{
		log.Printf("Error in the API request", err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if !validateRequest(req){
		log.Printf("Error in the API request", err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	fmt.Println("request", req)
	filter := search.Filter{
		SeverityString: req.Severity,
		AppName: req.AppName,
		MsgId: req.MsgID,
		Namespace: req.Namespace,
	}
	start := time.Now()
	filteredRecords,err := search.Search(h.Records,h.Index,req.Query,filter)
	duration := time.Since(start).Milliseconds()
	if err != nil{
		log.Printf("Search has failed", err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	fmt.Println("Search Results", filteredRecords)
	resp = models.SearchResponse{
		Records:   filteredRecords,
		Count:     len(filteredRecords),
		TimeTaken: duration,
	}
	c.JSON(http.StatusOK, resp)

}
	