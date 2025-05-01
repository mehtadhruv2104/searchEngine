package parquet

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/mehtadhruv2104/searchEngine/backend/models"
	"github.com/mehtadhruv2104/searchEngine/backend/search"
)


func Load_Files(path string)([]models.LogEntry, search.Index, error){

	start := time.Now()
	var allRecords []models.LogEntry
	index := make(search.Index)
	var mu sync.Mutex 
	var wg sync.WaitGroup


	files, err := os.ReadDir(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read directory %s: %v", path, err)
	}

	parquetFiles := []string{}
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".parquet" {
			parquetFiles = append(parquetFiles, filepath.Join(path, file.Name()))
		}
	}
	log.Printf("Found %d Parquet files in %s", len(parquetFiles), path)

	for _, filePath := range parquetFiles {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()

			err,records, fileIndex := Read_file(filePath)
			if err != nil {
				log.Printf("Error reading %s: %v", filePath, err)
				return
			}
			log.Printf("Read %d records from %s", len(records), filePath)

			mu.Lock()
            offset := len(allRecords)
            for word, ids := range fileIndex {
                adjustedIDs := make([]int, len(ids))
                for i, id := range ids {
                    adjustedIDs[i] = id + offset
                }
                index[word] = append(index[word], adjustedIDs...)
                index[word] = search.UniqueSorted(index[word])
            }
            allRecords = append(allRecords, records...)
            mu.Unlock()
		}(filePath)
	}

	wg.Wait()

	duration := time.Since(start)
	log.Printf("Loaded %d records and built index in %v", len(allRecords), duration)

	return allRecords, index, nil
}