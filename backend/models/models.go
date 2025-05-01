package models

type LogEntry struct {
    MsgId         string            `parquet:"name=MsgId, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    PartitionId   int32             `parquet:"name=PartitionId, type=INT32"`
    Timestamp     string            `parquet:"name=Timestamp, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    Hostname      string            `parquet:"name=Hostname, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    Priority      int32             `parquet:"name=Priority, type=INT32"`
    Facility      int32             `parquet:"name=Facility, type=INT32"`
    FacilityString string           `parquet:"name=FacilityString, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    Severity      int32             `parquet:"name=Severity, type=INT32"`
    SeverityString string           `parquet:"name=SeverityString, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    AppName       string            `parquet:"name=AppName, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    ProcId        string            `parquet:"name=ProcId, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    Message       string            `parquet:"name=Message, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    MessageRaw    string            `parquet:"name=MessageRaw, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    StructuredData string           `parquet:"name=StructuredData, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"` // flatten map
    Tag           string            `parquet:"name=Tag, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    Sender        string            `parquet:"name=Sender, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    Groupings     string            `parquet:"name=Groupings, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    Event         string            `parquet:"name=Event, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    EventId       string            `parquet:"name=EventId, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    NanoTimeStamp string            `parquet:"name=NanoTimeStamp, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
    Namespace     string            `parquet:"name=namespace, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type SearchRequest struct {
	Query     string  `json:"query"`
	Severity  *string `json:"severity"`
	AppName   *string `json:"app_name"`
	Namespace *string `json:"namespace"`
	MsgID     *string `json:"msg_id"`
}

type SearchResponse struct {
	Records   []LogEntry `json:"records"`
	Count     int               `json:"count"`
	TimeTaken int64             `json:"time_ms"`
}


