package solr

import "net/http"


type Response struct {
	HttpResponse 		*http.Response
	ResponseHeader 		ResponseHeader      	`json:"responseHeader,omitempty"`
	Collections 		[]*string      			`json:"collections,omitempty"`
	Response 			Result          		`json:"response,omitempty"`
	Error  				Error          	 		`json:"error,omitempty"`
	Exception  			Exception          		`json:"exception,omitempty"`
	Terms 	 			map[string]interface{} 	`json:"terms,omitempty"`
	Schema 	 			Schema 					`json:"schema,omitempty"`
	Fields 	 			[]Field 				`json:"fields,omitempty"`
	DynamicFields 	 	[]Field 				`json:"dynamicFields,omitempty"`
	UniqueKey 	 		[]Field 				`json:"uniqueKey,omitempty"`
	ReindexStatus 		ReindexStatus      		`json:"reindexStatus,omitempty"`
	GettingStarted 		GettingStarted      	`json:"gettingstarted,omitempty"`
	Summary 			map[string]interface{}  `json:"Summary,omitempty"`
	AlreadyLeaders 		map[string]interface{}  `json:"alreadyLeaders,omitempty"`
	InactivePreferreds 	map[string]interface{}  `json:"inactivePreferreds,omitempty"`
	Successes 			map[string]interface{}  `json:"successes,omitempty"`
}

type ResponseHeader struct {
	Status 				int 				`json:"status,omitempty"`
	QTime  				int 				`json:"QTime,omitempty"`
	Params 				Params 				`json:"params,omitempty"`
	ZKConnected 		bool 				`json:"zkConnected,omitempty"`
}

type Doc map[string]interface{}

type Result struct {
	NumFound 			int   				`json:"numFound,omitempty"`
	Start    			int   				`json:"start,omitempty"`
	Docs     			[]Doc 				`json:"docs,omitempty"`
}

type Error struct {
	Msg  				string 				`json:"msg,omitempty"`
	Code 				int 				`json:"code,omitempty"`
	Metadata 			[]*string 			`json:"metadata,omitempty"`
}

type Params struct {
	Indent 				bool 				`json:"indent,omitempty"`
	Q      				string				`json:"q,omitempty"`
	WT     				string				`json:"wt,omitempty"`
	Json 				string				`json:"json,omitempty"`
}

type Exception struct {
	Msg 				string 				`json:"msg"`
	RspCode 			int64 				`json:"rspCode"`
}

type Schema struct {
	Name 				string 				`json:"name,omitempty"`
	Version 			float32 			`json:"version,omitempty"`
	UniqueKey 			string 				`json:"uniqueKey,omitempty"`
	FieldTypes 			[]interface{} 		`json:"fieldTypes,omitempty"`
}

type Field struct {
	Name 					string 			`json:"name,omitempty"`
	Type 					string 			`json:"type,omitempty"`
	MultiValued 			bool 			`json:"multiValued,omitempty"`
	Indexed 				bool 			`json:"indexed,omitempty"`
	Stored 					bool 			`json:"stored,omitempty"`
	Required 				bool 			`json:"required,omitempty"`
	UseDocValuesAsStored	bool 			`json:"useDocValuesAsStored,omitempty"`
}

type ReindexStatus struct {
	Phase 						string 			`json:"phase,omitempty"`
	InputDocs 					int64 			`json:"inputDocs,omitempty"`
	ProcessedDocs 				int64 			`json:"processedDocs,omitempty"`
	ActualSourceCollection 		bool 			`json:"actualSourceCollection,omitempty"`
	State 						bool 			`json:"state,omitempty"`
	ActualTargetCollection 		bool 			`json:"actualTargetCollection,omitempty"`
	CheckpointCollection		bool 			`json:"checkpointCollection,omitempty"`
}

type GettingStarted struct {
	Phase 						string 					`json:"stateFormat,omitempty"`
	InputDocs 					int64 					`json:"znodeVersion,omitempty"`
	ProcessedDocs 				map[string]interface{} 	`json:"properties,omitempty"`
	ActiveShards 				bool 					`json:"activeShards,omitempty"`
	InactiveShards 				bool 					`json:"inactiveShards,omitempty"`
	SchemaNonCompliant 			[]*string 				`json:"schemaNonCompliant,omitempty"`
	Shards						map[string]interface{} 	`json:"shards,omitempty"`
}



