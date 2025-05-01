package parquet

import (
	"fmt"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)



func GetSchema() error{
	fr, err := local.NewLocalFileReader("File1.parquet") 
	if err != nil{
		return err
	}
	defer fr.Close()
	
	pr,err := reader.NewParquetReader(fr, nil,1)
	if err != nil{
		return err
	}
	defer pr.ReadStop()

	schema := pr.SchemaHandler
	fmt.Println("Parquet Schema:")
    for _, info := range schema.SchemaElements {
        fmt.Printf("Path: %v, Type: %v, ConvertedType: %v\n",
            info.GetName(), info.Type, info.ConvertedType)
    }

	return nil
}

