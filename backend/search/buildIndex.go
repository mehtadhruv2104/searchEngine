package search

import (
	"encoding/json"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/mehtadhruv2104/searchEngine/backend/models"
)


type Index map[string][]int

func BuildIndex(records []models.LogEntry) (Index,error) {
    index := make(Index)

    for i, record := range records {
        fields := []string{
            derefString(&record.Tag),
            derefString(&record.Sender),
            derefString(&record.AppName),
            derefString(&record.Namespace),
            derefString(&record.SeverityString),
            derefString(&record.Hostname),
            derefString(&record.Event),
            derefString(&record.MsgId),
            derefIntToString(&record.Severity),
            derefIntToString(&record.PartitionId),
            derefIntToString(&record.Facility),
            derefIntToString(&record.Priority),
            derefString(&record.Message),
        }

        if raw := derefString(&record.MessageRaw); raw != "" {
            var rawData map[string]interface{}
            if err := json.Unmarshal([]byte(raw), &rawData); err == nil {
                if log, ok := rawData["log"].(string); ok {
                    words := tokenize(log)
                    for _, word := range words {
                        index[word] = append(index[word], i)
                    }
                }
                if k8s, ok := rawData["kubernetes"].(map[string]interface{}); ok {
                    for _, key := range []string{"container_name", "pod_name", "namespace_name"} {
                        if val, ok := k8s[key].(string); ok {
                            words := tokenize(val)
                            for _, word := range words {
                                index[key+"."+word] = append(index[key+"."+word], i)
                            }
                        }
                    }
                }
            }
        }

        if sd := derefString(&record.StructuredData); sd != "" {
            var sdData map[string][]string
            if err := json.Unmarshal([]byte(sd), &sdData); err == nil {
                for key, values := range sdData {
                    words := tokenize(key)
                    for _, word := range words {
                        index["structured."+word] = append(index["structured."+word], i)
                    }
                    for _, value := range values {
                        words := tokenize(value)
                        for _, word := range words {
                            index["structured."+key+"."+word] = append(index["structured."+key+"."+word], i)
                        }
                    }
                }
            }
        }

        for _, field := range fields {
            if field != "" {
                words := tokenize(field)
                for _, word := range words {
                    index[word] = append(index[word], i)
                }
            }
        }

    }
    return index,nil
}

func UniqueSorted(ids []int) []int {
    if len(ids) == 0 {
        return ids
    }

    idMap := make(map[int]bool)
    for _, id := range ids {
        idMap[id] = true
    }

    uniqueIDs := make([]int, 0, len(idMap))
    for id := range idMap {
        uniqueIDs = append(uniqueIDs, id)
    }
    sort.Ints(uniqueIDs)
    return uniqueIDs
}


func tokenize(text string) []string {

    text = strings.ToLower(text)
    text = strings.ReplaceAll(text, ",", " ")
    text = strings.ReplaceAll(text, ".", " ")
    text = strings.ReplaceAll(text, ";", " ")
    text = strings.ReplaceAll(text, "(", " ")
    text = strings.ReplaceAll(text, ")", " ")
    text = strings.ReplaceAll(text, "[", " ")
    text = strings.ReplaceAll(text, "]", " ")
    text = strings.ReplaceAll(text, "=", " ")
    text = strings.ReplaceAll(text, "/", " ")
    text = strings.ReplaceAll(text, "\\", " ")

    timestampRegex := regexp.MustCompile(`[0-9]{1,4}/[0-9]{1,2}/[0-9]{1,2} [0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2}`)
    text = timestampRegex.ReplaceAllString(text, " ")
    numberRegex := regexp.MustCompile(`\b\d+\b`)
    text = numberRegex.ReplaceAllString(text, " ")

    words := strings.Fields(text)

    stopWords := map[string]bool{
        "stdout": true,
        "f":      true,
        "in":     true,
        "with":   true,
        "at":     true,
        "for":    true,
        "org":    true,
        "have":   true,
        "wrote":    true,
        "because": true,
    }
    var filtered []string
    for _, word := range words {
        if !stopWords[word] && len(word) > 2 {
            filtered = append(filtered, word)
        }
    }
    return filtered
}

func derefString(s *string) string {
    if s == nil {
        return ""
    }
    return *s
}


func derefIntToString(n *int32) string{
    if n == nil{
        return ""
    }
    return strconv.FormatInt(int64(*n),10)
}