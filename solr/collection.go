// https:// lucene.apache.org/solr/guide/8_5/collection-management.html#collection-management
// 
package solr

import (
	"context"
	"net/http"
)

type WT string

const (
	JSON	WT	=	"json"
	XML		WT	=	"xml"
)

type CollectionAction string

const (
	CreateAction			CollectionAction			=	"CREATE"
	ReloadAction			CollectionAction			=	"RELOAD"
	ModifyCollectionAction	CollectionAction			=	"MODIFYCOLLECTION"
	ListAction				CollectionAction			=	"LIST"
	RenameAction			CollectionAction			=	"RENAME"
	DeleteAction			CollectionAction			=	"DELETE"
	CollectionPropAction	CollectionAction			=	"COLLECTIONPROP"
	MigrateAction			CollectionAction			=	"MIGRATE"
	ReindexCollectionAction	CollectionAction			=	"REINDEXCOLLECTION"
	ColStatusAction			CollectionAction			=	"COLSTATUS"
	BackupAction			CollectionAction			=	"BACKUP"
	RestoreAction			CollectionAction			=	"RESTORE"
	RebalanceLeadersAction	CollectionAction			=	"REBALANCELEADERS"
)

type ReindexCollectionCmd string

const (
	Start 					ReindexCollectionCmd 		= "start"
	Abort 					ReindexCollectionCmd 		= "abort"
	Status 					ReindexCollectionCmd 		= "status"
)

type RawSizeSamplingPercent string

const (
	FieldsBySize			RawSizeSamplingPercent		= "fieldsBySize"
	TypesBySize				RawSizeSamplingPercent		= "typesBySize"
	Summary					RawSizeSamplingPercent		= "summary"
	Details					RawSizeSamplingPercent		= "details"
	StoredFields			RawSizeSamplingPercent		= "storedFields"
	TermsTerms				RawSizeSamplingPercent		= "terms_terms"
	TermsPostings			RawSizeSamplingPercent		= "terms_postings"
	TermsPayloads			RawSizeSamplingPercent		= "terms_payloads"
	TermVectors				RawSizeSamplingPercent		= "termVectors"
	DocValues				RawSizeSamplingPercent		= "docValues_*"
	Points					RawSizeSamplingPercent		= "norms"
)

type collectionBase struct {
	Action 					CollectionAction		`url:"action,omitempty"`
	WT 						WT						`url:"wt,omitempty"`
}

type CollectionCreate struct {
	collectionBase
	// The name of the collection to be created. This parameter is required.
	Name 					string 		`url:"name,omitempty"`

	// The router name that will be used. The router defines how
	// documents will be distributed among the shards. Possible
	// values are implicit or compositeId, which is the default.
	// 
	// The implicit router does not automatically route documents
	// to different shards. Whichever shard you indicate on the
	// indexing request (or within each document) will be used as
	// the destination for those documents.
	// 
	// The compositeId router hashes the value in the uniqueKey field
	// and looks up that hash in the collection’s clusterstate to
	// determine which shard will receive the document, with the
	// additional ability to manually direct the routing.
	// 
	// When using the implicit router, the shards parameter is required.
	// When using the compositeId router, the numShards parameter is required.
	// 
	// For more information, see also the section Document Routing.
	RouterName 				string 		`url:"router.name,omitempty"`

	// The number of shards to be created as part of the collection.
	// This is a required parameter when the router.name is compositeId.
	NumShards 				int 		`url:"numShards,omitempty"`

	// A comma separated list of shard names, e.g., shard-x,shard-y,shard-z.
	// This is a required parameter when the router.name is implicit.
	Shards 					int 		`url:"shards,omitempty"`

	// The number of replicas to be created for each shard. The default is 1.
	// 
	// This will create a NRT type of replica. If you want another type of replica,
	// see the tlogReplicas and pullReplica parameters below. See the section Types
	// of Replicas for more information about replica types.
	ReplicationFactor 		int 		`url:"replicationFactor,omitempty"`

	// The number of NRT (Near-Real-Time) replicas to create for this collection.
	// This type of replica maintains a transaction log and updates its index locally.
	// If you want all of your replicas to be of this type, you can simply use
	// replicationFactor instead.
	NrtReplicas 			int 		`url:"nrtReplicas,omitempty"`

	// The number of TLOG replicas to create for this collection.
	// This type of replica maintains a transaction log but only updates
	// its index via replication from a leader. See the section Types of
	// Replicas for more information about replica types.
	TLogReplicas 			int 		`url:"tlogReplicas,omitempty"`

	// The number of PULL replicas to create for this collection.
	// This type of replica does not maintain a transaction log and only
	// updates its index via replication from a leader. This type is not
	// eligible to become a leader and should not be the only type of replicas
	// in the collection. See the section Types of Replicas for more information
	// about replica types.
	PullReplicas 			int 		`url:"pullReplicas,omitempty"`

	// When creating collections, the shards and/or replicas are spread across all
	// available (i.e., live) nodes, and two replicas of the same shard will never
	// be on the same node.
	// 
	// If a node is not live when the CREATE action is called, it will not get any
	// parts of the new collection, which could lead to too many replicas being created
	// on a single live node. Defining maxShardsPerNode sets a limit on the number of
	// replicas the CREATE action will spread to each node.
	// 
	// If the entire collection can not be fit into the live nodes, no collection
	// will be created at all. The default maxShardsPerNode value is 1. A value of -1 means
	// unlimited. If a policy is also specified then the stricter of maxShardsPerNode and policy rules apply.
	MaxShardsPerNode 		string 		`url:"maxShardsPerNode,omitempty"`

	// Allows defining the nodes to spread the new collection across. The format is a
	// comma-separated list of node_names, such as localhost:8983_solr,localhost:8984_solr,localhost:8985_solr.
	// 
	// If not provided, the CREATE operation will create shard-replicas spread across all live Solr nodes.
	// 
	// Alternatively, use the special value of EMPTY to initially create no shard-replica
	// within the new collection and then later use the ADDREPLICA operation to add shard-replicas
	// when and where required.
	CreateNodeSet 			string 		`url:"createNodeSet,omitempty"`

	// Controls whether or not the shard-replicas created for this collection will be assigned to
	// the nodes specified by the createNodeSet in a sequential manner, or if the list of nodes
	// should be shuffled prior to creating individual replicas.
	// 
	// A false value makes the results of a collection creation predictable and gives
	// more exact control over the location of the individual shard-replicas, but true can
	// be a better choice for ensuring replicas are distributed evenly across nodes. The default is true.
	// 
	// This parameter is ignored if createNodeSet is not also specified.
	CreateNodeSetShuffle 	string 		`url:"createNodeSet.shuffle,omitempty"`

	// Defines the name of the configuration (which must already be stored in ZooKeeper) to use for
	// this collection. If not provided, Solr will use the configuration of _default configset
	// to create a new (and mutable) configset named <collectionName>.AUTOCREATED and will use it for
	// the new collection. When such a collection (that uses a copy of the _default configset)
	// is deleted, the autocreated configset is not deleted by default.
	CollectionConfigName 	string 		`url:"collection.configName,omitempty"`

	// If this parameter is specified, the router will look at the value of the field in an input
	// document to compute the hash and identify a shard instead of looking at the uniqueKey field.
	// If the field specified is null in the document, the document will be rejected.
	// 
	// Please note that RealTime Get or retrieval by document ID would also require the
	// parameter _route_ (or shard.keys) to avoid a distributed search.
	RouterField 			string 		`url:"router.field,omitempty"`

	// Set core property name to value. See the section Defining core.properties for
	// details on supported properties and values.
	PropertyName 			string 		`url:"property.name,omitempty"`

	// When set to true, enables automatic addition of replicas when the number of active replicas
	// falls below the value set for replicationFactor. This may occur if a replica goes down,
	// for example. The default is false, which means new replicas will not be added.
	// 
	// While this parameter is provided as part of Solr’s set of features to provide
	// autoscaling of clusters, it is available even when you have not implemented any
	// other part of autoscaling (such as a policy). See the section SolrCloud Autoscaling
	// Automatically Adding Replicas for more details about this option and how it can be used.
	AutoAddReplicas 		bool 		`url:"autoAddReplicas,omitempty"`

	// Replica placement rules. See the section Rule-based Replica Placement for details.
	Rule 					string 		`url:"rule,omitempty"`

	// Details of the snitch provider. See the section Rule-based Replica Placement for details.
	Snitch 					string 		`url:"snitch,omitempty"`

	// Name of the collection-level policy. See Defining Collection-Specific Policies for details.
	Policy 					string 		`url:"policy,omitempty"`

	// If true, the request will complete only when all affected replicas become active. The
	// default is false, which means that the API will return the status of the single action,
	// which may be before the new replica is online and active.
	WaitForFinalState 		string 		`url:"waitForFinalState,omitempty"`

	// The name of the collection with which all replicas of this collection must be co-located.
	// The collection must already exist and must have a single shard named shard1. See Colocating
	// collections for more details.
	WithCollection 			string 		`url:"withCollection,omitempty"`

	// Starting with version 8.1 when a collection is created additionally an alias can be created
	// that points to this collection. This parameter allows specifying the name of this alias,
	// effectively combining this operation with CREATEALIAS
	Alias 					string 		`url:"alias,omitempty"`

	// Request ID to track this action which will be processed asynchronously.
	Async 					bool 		`url:"async,omitempty"`
}

type CollectionReload struct {
	collectionBase

	// The name of the collection to reload. This parameter is required.
	Name 					string 		`url:"name,omitempty"`

	// Request ID to track this action which will be processed asynchronously.
	Async 					bool 		`url:"async,omitempty"`
}

type CollectionModifyCollection struct {
	collectionBase
	// The name of the collection to be modified. This parameter is required.
	Collection 				string 							`url:"collection,omitempty"`

	// Key-value pairs of attribute names and attribute values.
	// At least one attribute parameter is required.
	// 
	// The attributes that can be modified are:
	// 
	// Read-Only Mode
	// Setting the readOnly attribute to true puts the collection in read-only mode, in which any
	// index update requests are rejected. Other collection-level actions (e.g., adding / removing / moving replicas)
	// are still available in this mode.
	// 
	// The transition from the (default) read-write to read-only mode consists of the following steps:
	// 
	// the readOnly flag is changed in collection state,
	// any new update requests are rejected with 403 FORBIDDEN error code (ongoing long-running requests
	// are aborted, too),
	// a forced commit is performed to flush and commit any in-flight updates.
	MaxShardsPerNode 		int 							`url:"maxShardsPerNode,omitempty"`
	ReplicationFactor 		int 							`url:"replicationFactor,omitempty"`
	AutoAddReplicas 		int 							`url:"autoAddReplicas,omitempty"`
	ConfigName 				string 							`url:"collection.configName,omitempty"`
	Rule 					string 							`url:"rule,omitempty"`
	Snitch 					string 							`url:"snitch,omitempty"`
	Policy 					string 							`url:"policy,omitempty"`
	WithCollection 			string 							`url:"withCollection,omitempty"`
	ReadOnly 				bool 							`url:"readOnly,omitempty"`
}

type CollectionList struct {
	collectionBase
}

type CollectionRename struct {
	collectionBase

	// Name of the existing SolrCloud collection or an alias that refers to exactly one collection
	// and is not a Routed Alias.
	Name 					string 		`url:"name,omitempty"`

	// Target name of the collection. This will be the new alias that refers to the underlying
	// SolrCloud collection. The original name (or alias) of the collection will be replaced also
	// in the existing aliases so that they also refer to the new name. Target name must not be an
	// existing alias.
	Target 					string 		`url:"target,omitempty"`
}

type CollectionDelete struct {
	collectionBase

	// The name of the collection to delete. This parameter is required.
	Name 					string 		`url:"name,omitempty"`

	// Request ID to track this action which will be processed asynchronously.
	Async 					bool 		`url:"async,omitempty"`
}

type CollectionProp struct {
	collectionBase

	// The name of the collection for which the property would be set.
	Name 					string 		`url:"name,omitempty"`

	// The name of the property.
	PropertyName 			string 		`url:"propertyName,omitempty"`

	// The value of the property. When not provided, the property is deleted.
	PropertyValue 			string 		`url:"propertyValue,omitempty"`
}

type CollectionMigrate struct {
	collectionBase

	// The name of the source collection from which documents will be split. This parameter is required.
	Collection 			string 		`url:"collection,omitempty"`

	// The name of the target collection to which documents will be migrated. This parameter is required.
	TargetCollection 		string 		`url:"target.collection,omitempty"`

	// The routing key prefix. For example, if the uniqueKey of a document is "a!123",
	// then you would use split.key=a!. This parameter is required.
	SplitKey 				string 		`url:"split.key,omitempty"`

	// The timeout, in seconds, until which write requests made to the source collection for the
	// given split.key will be forwarded to the target shard. The default is 60 seconds.
	ForwardTimeout 			int 		`url:"forward.timeout,omitempty"`

	// Set core property name to value. See the section Defining core.properties for details
	// on supported properties and values.
	PropertyName 			string 		`url:"property.name,omitempty"`

	// Request ID to track this action which will be processed asynchronously.
	Async 					bool 		`url:"async,omitempty"`
}

type CollectionReindex struct {
	collectionBase

	// The name of the source collection from which documents will be split. This parameter is required.
	Name 					string 						`url:"name,omitempty"`

	// Optional command. Default command is start. Currently supported commands are:

	// start - default, starts processing if not already running,
	// abort - aborts an already running reindexing (or clears a left-over status after a crash), and deletes partial results,
	// status - returns detailed status of a running reindexing command.
	CMD 					ReindexCollectionCmd 		`url:"cmd,omitempty"`

	// Target collection name, optional. If not specified a unique name will be generated and after
	// all documents have been copied an alias will be created that points from the source collection
	// name to the unique sequentially-named collection, effectively "hiding" the original source
	// collection from regular update and search operations.
	Target 					string 						`url:"target,omitempty"`

	// Optional query to select documents for reindexing. Default value is *:*.
	Q 						string 						`url:"q,omitempty"`

	// Optional list of fields to reindex. Default value is *.
	FL 						string 						`url:"fl,omitempty"`

	// Documents are transferred in batches. Depending on the average size of the document large batch
	// sizes may cause memory issues. Default value is 100.
	Rows 					string 						`url:"rows,omitempty"`

	// Optional name of the configset for the target collection. Default is the same as the source collection.
	// There’s a number of optional parameters that determine the target collection layout. If they are
	// not specified in the request then their values are copied from the source collection.
	// The following parameters are currently supported (described in details in the CREATE collection section):
	// numShards, replicationFactor, nrtReplicas, tlogReplicas, pullReplicas, maxShardsPerNode,
	// autoAddReplicas, shards, policy, createNodeSet, createNodeSet.shuffle, router.*.
	ConfigName 				string 						`url:"configName,omitempty"`

	// Optional boolean. If true then after the processing is successfully finished the source
	// collection will be deleted.
	RemoveSource 			string 						`url:"removeSource,omitempty"`

	// Request ID to track this action which will be processed asynchronously.
	Async 					bool 						`url:"async,omitempty"`
}

type CollectionColStatus struct {
	collectionBase

	// Collection name (optional). If missing then it means all collections.
	Collection 				string 		`url:"collection,omitempty"`

	// Optional boolean. If true then additional information will be provided about SolrCore of shard leaders.
	CoreInfo 					string 		`url:"coreInfo,omitempty"`

	// Optional boolean. If true then segment information will be provided.
	Segments 					string 		`url:"segments,omitempty"`

	// Optional boolean. If true then detailed Lucene field information will be provided and their
	// corresponding Solr schema types.
	FieldInfo 					string 		`url:"fieldInfo,omitempty"`

	// Optional boolean. If true then additional information about the index files size and their
	// RAM usage will be provided
	SizeInfo 					string 		`url:"sizeInfo,omitempty"`

	// Optional boolean. If true then run the raw index data analysis tool (other boolean
	// options below imply this option if any of them are true). Command response will include
	// sections that show estimated breakdown of data size per field and per data type.
	RawSize 					string 		`url:"rawSize,omitempty"`

	// Optional boolean. If true then include also a more detailed breakdown of data size per field and per type.
	RawSizeSummary 				string 		`url:"rawSizeSummary,omitempty"`

	// Optional boolean. If true then provide exhaustive details that include statistical distribution
	// of items per field and per type as well as top 20 largest items per field.
	RawSizeDetails 				string 		`url:"rawSizeDetails,omitempty"`

	// Optional float. When the index is larger than a certain threshold (100k documents per shard)
	// only a part of data is actually retrieved and analyzed in order to reduce the IO load, and then
	// the final results are extrapolated. Values must be greater than 0 and less or equal to 100.0.
	// Default value is 5.0. Very small values (between 0.0 and 1.0) may introduce significant estimation
	// errors. Also, values that would result in less than 10 documents being sampled are rejected with an
	// exception.
	// 
	// Response for this command always contains two sections:
	// 
	// fieldsBySize is a map where field names are keys and values are estimated sizes of raw (uncompressed) data
	// that belongs to the field. The map is sorted by size so that it’s easy to see what field occupies most space.
	// typesBySize is a map where data types are the keys and values are estimates sizes of raw (uncompressed)
	// data of particular type. This map is also sorted by size.
	// Optional sections include:
	// 
	// summary section containing a breakdown of data sizes for each field by data type.
	// details section containing detailed statistical summary of size distribution within each field, per data type.
	// This section also shows topN values by size from each field.
	// Data types shown in the response can be roughly divided into the following groups:
	// 
	// storedFields - represents the raw uncompressed data in stored fields. For example, for UTF-8 strings this
	// represents the aggregated sum of the number of bytes in the strings' UTF-8 representation, for long numbers
	// this is 8 bytes per value, etc.
	// 
	// terms_terms - represents the aggregated size of the term dictionary. The size of this data is affected by
	// the the number and length of unique terms, which in turn depends on the field size and the analysis chain.
	// 
	// terms_postings - represents the aggregated size of all term position and offset information, if present.
	// This information may be absent if position-based searching, such as phrase queries, is not needed.
	// 
	// terms_payloads - represents the aggregated size of all per-term payload data, if present.
	// norms - represents the aggregated size of field norm information. This information may be omitted if a field
	// has an omitNorms flag in the schema, which is common for fields that don’t need weighting or scoring by
	// field length.
	// 
	// termVectors - represents the aggregated size of term vectors.
	// 
	// docValues_* - represents aggregated size of doc values, by type (e.g., docValues_numeric, docValues_binary, etc).
	// 
	// points - represents aggregated size of point values.
	RawSizeSamplingPercent 		RawSizeSamplingPercent 		`url:"rawSizeSamplingPercent,omitempty"`
}

type CollectionBackup struct {
	collectionBase

	// The collection where the indexes will be restored into. This parameter is required.
	Collection 			string 		`url:"collection,omitempty"`

	// The name of the existing backup that you want to restore. This parameter is required.
	Name 					string 		`url:"name,omitempty"`

	// The location on a shared drive for the RESTORE command to read from. Alternately it can be set
	// as a cluster property.
	Location 				string 		`url:"location,omitempty"`

	// Request ID to track this action which will be processed asynchronously.
	Async 					bool 		`url:"async,omitempty"`
}

type CollectionRestore struct {
	collectionBase

	// The collection where the indexes will be restored into. This parameter is required.
	Collection 			string 		`url:"collection,omitempty"`

	// The name of the existing backup that you want to restore. This parameter is required.
	Name 					string 		`url:"name,omitempty"`

	// The location on a shared drive for the RESTORE command to read from. Alternately it can be set
	// as a cluster property.
	Location 				string 		`url:"location,omitempty"`

	// Request ID to track this action which will be processed asynchronously.
	Async 					bool 		`url:"async,omitempty"`

	// The name of a repository to be used for the backup. If no repository is specified then the
	// local filesystem repository will be used automatically.
	Repository 				string 		`url:"repository,omitempty"`

	// Defines the name of the configurations to use for this collection. These must already be stored
	// in ZooKeeper. If not provided, Solr will default to the collection name as the configuration name.
	CollectionConfigName 	string 		`url:"collection.configName,omitempty"`

	// The number of replicas to be created for each shard.
	ReplicationFactor 		int 		`url:"replicationFactor,omitempty"`

	// The number of NRT (Near-Real-Time) replicas to create for this collection. This type of replica
	// maintains a transaction log and updates its index locally. This parameter behaves the same way
	// as setting replicationFactor parameter.
	NrtReplicas 			int 		`url:"nrtReplicas,omitempty"`

	// The number of TLOG replicas to create for this collection. This type of replica maintains a
	// transaction log but only updates its index via replication from a leader. See the section Types of
	// Replicas for more information about replica types.
	TLogReplicas 			int 		`url:"tlogReplicas,omitempty"`

	// The number of TLOG replicas to create for this collection. This type of replica maintains a
	// transaction log but only updates its index via replication from a leader. See the section Types
	// of Replicas for more information about replica types.
	PullReplicas 			int 		`url:"pullReplicas,omitempty"`

	// When creating collections, the shards and/or replicas are spread across all available (i.e., live)
	// nodes, and two replicas of the same shard will never be on the same node.
	// 
	// If a node is not live when the CREATE operation is called, it will not get any parts of the new
	// collection, which could lead to too many replicas being created on a single live node. Defining
	// maxShardsPerNode sets a limit on the number of replicas CREATE will spread to each node. If the
	// entire collection can not be fit into the live nodes, no collection will be created at all.
	MaxShardsPerNode 		string 		`url:"maxShardsPerNode,omitempty"`

	// When set to true, enables auto addition of replicas on shared file systems. See the section
	// Automatically Add Replicas in SolrCloud for more details on settings and overrides.
	AutoAddReplicas 		string 		`url:"autoAddReplicas,omitempty"`

	// Set core property name to value. See the section Defining core.properties for details on
	// supported properties and values.
	PropertyName 			string 		`url:"property.name,omitempty"`
}

type CollectionRebalanceLeaders struct {
	collectionBase

	// The collection where the indexes will be restored into. This parameter is required.
	Collection 			string 		`url:"collection,omitempty"`

	// The maximum number of reassignments to have queue up at once. Values <=0 are use the default
	// value Integer.MAX_VALUE.
	// 
	// When this number is reached, the process waits for one or more leaders to be successfully
	// assigned before adding more to the queue.
	MaxAtOnce 				string 		`url:"maxAtOnce,omitempty"`

	// Defaults to 60. This is the timeout value when waiting for leaders to be reassigned. If maxAtOnce
	// is less than the number of reassignments that will take place, this is the maximum interval that
	// any single wait for at least one reassignment.
	// 
	// For example, if 10 reassignments are to take place and maxAtOnce is 1 and maxWaitSeconds is 60,
	// the upper bound on the time that the command may wait is 10 minutes.
	MaxWaitSeconds 			string 		`url:"maxWaitSeconds,omitempty"`
}

type CollectionAPI struct {
	client *Client
}


// CREATE: Create a Collection
func (c *CollectionAPI) Create(ctx context.Context, collection CollectionCreate) (*Response, error) {
	collection.WT = JSON
	collection.Action = CreateAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}
	
	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// Reload: Reload a Collection
func (c *CollectionAPI) Reload(ctx context.Context, collection CollectionReload) (*Response, error) {
	collection.WT = JSON
	collection.Action = ReloadAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// Modify: Modify Attributes of a Collection
func (c *CollectionAPI) Modify(ctx context.Context, collection CollectionModifyCollection) (*Response, error) {
	collection.WT = JSON
	collection.Action = ModifyCollectionAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// List: List Collections
func (c *CollectionAPI) List(ctx context.Context) (*Response, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collectionBase{
		Action: ListAction,
		WT:     JSON,
	})
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// Rename: Rename a Collection
func (c *CollectionAPI) Rename(ctx context.Context, collection CollectionRename) (*Response, error) {
	collection.WT = JSON
	collection.Action = RenameAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// Delete: Delete a Collection
func (c *CollectionAPI) Delete(ctx context.Context, collection CollectionDelete) (*Response, error) {
	collection.WT = JSON
	collection.Action = DeleteAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// CollectionProp: Collection Properties
// Add, edit or delete a collection property.
func (c *CollectionAPI) CollectionProp(ctx context.Context, collection CollectionProp) (*Response, error) {
	collection.WT = JSON
	collection.Action = CollectionPropAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// Migrate: Migrate Documents to Another Collection
func (c *CollectionAPI) Migrate(ctx context.Context, collection CollectionMigrate) (*Response, error) {
	collection.WT = JSON
	collection.Action = MigrateAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// ReindexCollection: Re-Index a Collection
func (c *CollectionAPI) ReindexCollection(ctx context.Context, collection CollectionReindex) (*Response, error) {
	collection.WT = JSON
	collection.Action = ReindexCollectionAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// ColStatus: Detailed Status of a Collection’s Indexes
// The COLSTATUS command provides a detailed description of the collection status, including low-level
// index information about segments and field data.
func (c *CollectionAPI) ColStatus(ctx context.Context, collection CollectionColStatus) (*Response, error) {
	collection.WT = JSON
	collection.Action = ColStatusAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// Backup: Backup Collection
// Backs up Solr collections and associated configurations to a shared filesystem - for example a Network File System.
func (c *CollectionAPI) Backup(ctx context.Context, collection CollectionBackup) (*Response, error) {
	collection.WT = JSON
	collection.Action = BackupAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// Restore: Restore Collection
// Restores Solr indexes and associated configurations.
func (c *CollectionAPI) Restore(ctx context.Context, collection CollectionRestore) (*Response, error) {
	collection.WT = JSON
	collection.Action = RestoreAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

// RebalanceLeaders: Rebalance Leaders
// Reassigns leaders in a collection according to the preferredLeader property across active nodes.
func (c *CollectionAPI) RebalanceLeaders(ctx context.Context, collection CollectionRebalanceLeaders) (*Response, error) {
	collection.WT = JSON
	collection.Action = RebalanceLeadersAction

	req, err := c.client.NewRequest(ctx, http.MethodGet, "/solr/admin/collections", nil, collection)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}