package main

import (
	"log"

	"github.com/mehtadhruv2104/searchEngine/backend/handler"
	"github.com/mehtadhruv2104/searchEngine/backend/parquet"
	"github.com/mehtadhruv2104/searchEngine/backend/routes"
)

func main() {

	path := "data"
	records,index, err := parquet.Load_Files(path)
	if err!=nil {
		log.Printf("Error in loading Files and Initializing Index",err)
	}
	
	h := handler.NewHandler(records, index)
	router:= routes.StartEngine(h)
	router.Run(":8100")

 }
 