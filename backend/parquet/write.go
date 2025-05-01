package parquet

import (
	"fmt"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)



type Person struct {
    Name  string `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    Age   int32  `parquet:"name=age, type=INT32"`
    Email string `parquet:"name=email, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

func WriteToParquet() error {
    fw, err := local.NewLocalFileWriter("example.parquet")
    if err != nil {
        return err
    }
	defer fw.Close()
	fmt.Println("I reached writer")

    pw, err := writer.NewParquetWriter(fw, new(Person), 4)
    if err != nil {
        return err
    }
	defer pw.WriteStop()

    persons := []Person{
        {"Alice", 25, "alice@example.com"},
        {"Bob", 30, "bob@example.com"},
		{"Alice", 25, "alice@example.com"},
        {"Bob", 30, "bob@example.com"},
		{"Alice", 25, "alice@example.com"},
        {"Bob", 30, "bob@example.com"},
		{"Alice", 25, "alice@example.com"},
        {"Bob", 30, "bob@example.com"},
		{"Alice", 25, "alice@example.com"},
        {"Bob", 30, "bob@example.com"},
    }

    for _, person := range persons {
        if err = pw.Write(person); err != nil {
            return err
        }
    }

    return nil
}