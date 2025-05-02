package search

import (
	"fmt"
	"strings"

	"github.com/mehtadhruv2104/searchEngine/backend/models"
)

type Filter struct {
    SeverityString string
    AppName        string
    Namespace      string
	MsgId			string
}

func Search(records []models.LogEntry, index Index, query string, filter Filter) ([]models.LogEntry,error) {

    queries := parseQuery(query)
    var recordIDs []int
    if len(queries) > 0 {
		if len(queries[""]) == 0 {
			recordIDs = make([]int, len(records))
			for i := range records {
				recordIDs[i] = i
			}
		} else {
			words := queries[""]
			recordIDs = index[words[0]]
			
			for _, word := range words[1:] {
				recordIDs = intersect(recordIDs, index[word])
			}
		}
		for field, words := range queries {
			if field == "" {
				continue
			}
			for _, word := range words {
				fieldIDs := index[field+"."+word]
				fmt.Println("Field", field, "Word", word, "FieldIDs", fieldIDs)
				recordIDs = intersect(recordIDs, fieldIDs)
			}
		}
    } else {
        recordIDs = make([]int, len(records))
        for i := range records {
            recordIDs[i] = i
        }
    }

    var results []models.LogEntry
    for _, id := range recordIDs {
        record := records[id]
        if matchesFilter(record, filter) {
            results = append(results, record)
        }
    }

    return results,nil
}

func intersect(a, b []int) []int {
    var result []int
    i, j := 0, 0
    for i < len(a) && j < len(b) {
        if a[i] == b[j] {
            result = append(result, a[i])
            i++
            j++
        } else if a[i] < b[j] {
            i++
        } else {
            j++
        }
    }
    return result
}

func matchesFilter(record models.LogEntry, filter Filter) bool {
    if filter.SeverityString != "" && derefString(&record.SeverityString) != filter.SeverityString {
        return false
    }
    if filter.AppName != "" && derefString(&record.AppName) != filter.AppName {
        return false
    }
    if filter.Namespace != "" && derefString(&record.Namespace) != filter.Namespace {
        return false
    }
	if filter.MsgId != "" && derefString(&record.MsgId) != filter.MsgId {
        return false
    }
    return true
}

func parseQuery(query string) map[string][]string {
	fieldQueries := make(map[string][]string)
	parts := strings.Fields(query)
	for _, part := range parts {
		if strings.Contains(part, ":") {
			fieldValue := strings.SplitN(part, ":", 2)
			if len(fieldValue) == 2 {
				field, value := fieldValue[0], fieldValue[1]
				words := tokenize(value)
				fieldQueries[field] = append(fieldQueries[field], words...)
			}
		} else {
			words := tokenize(part)
			fieldQueries[""] = append(fieldQueries[""], words...)
		}
	}
	return fieldQueries
}