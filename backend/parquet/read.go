package parquet

import (
	"github.com/mehtadhruv2104/searchEngine/backend/models"
	"github.com/mehtadhruv2104/searchEngine/backend/search"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)



func Read_file(path string) (error, []models.LogEntry,search.Index ){

	fr,err := local.NewLocalFileReader(path)
	var records []models.LogEntry
	var index search.Index
	if err != nil{
		return err, records,index 
	}
	defer fr.Close()
	
	pr,err := reader.NewParquetReader(fr,&models.LogEntry{},4)
	if err != nil{
		return err,records,index
	}
	defer pr.ReadStop()

	numofrows := int(pr.GetNumRows())
	
	for numofrows > 0 {
		readCount := 10
		if numofrows < 10{
			readCount = numofrows
		}
		logs := make([]models.LogEntry, readCount)
		if err = pr.Read(&logs); err != nil {
			return err,records,index
		}
		records = append(records, logs...)
		numofrows -= readCount
	}

	index,err = search.BuildIndex(records)
	if err != nil{
		return err,records,index
	}

	return nil,records,index

}