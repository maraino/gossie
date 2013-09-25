// Experimental fork of carloscm/gossie
package cassandra

import (
	"fmt"
	"github.com/pomack/thrift4go/lib/go/src/thrift"
	"math"
)

// This is a temporary safety measure to ensure that the `math'
// import does not trip up any generated output that may not
// happen to use the math import due to not having emited enums.
//
// Future clean-ups will deprecate the need for this.
func init() {
	var temporaryAndUnused int32 = math.MinInt32
	temporaryAndUnused++
}

type ICassandra interface {
	/**
	 * Parameters:
	 *  - AuthRequest
	 */
	Login(auth_request *AuthenticationRequest) (authnx *AuthenticationException, authzx *AuthorizationException, err error)
	/**
	 * Parameters:
	 *  - Keyspace
	 */
	SetKeyspace(keyspace string) (ire *InvalidRequestException, err error)
	/**
	 * Get the Column or SuperColumn at the given column_path. If no value is present, NotFoundException is thrown. (This is
	 * the only method that can throw an exception under non-failure conditions.)
	 * 
	 * Parameters:
	 *  - Key
	 *  - ColumnPath
	 *  - ConsistencyLevel
	 */
	Get(key []byte, column_path *ColumnPath, consistency_level ConsistencyLevel) (retval448 *ColumnOrSuperColumn, ire *InvalidRequestException, nfe *NotFoundException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * Get the group of columns contained by column_parent (either a ColumnFamily name or a ColumnFamily/SuperColumn name
	 * pair) specified by the given SlicePredicate. If no matching values are found, an empty list is returned.
	 * 
	 * Parameters:
	 *  - Key
	 *  - ColumnParent
	 *  - Predicate
	 *  - ConsistencyLevel
	 */
	GetSlice(key []byte, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval449 thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * returns the number of columns matching <code>predicate</code> for a particular <code>key</code>,
	 * <code>ColumnFamily</code> and optionally <code>SuperColumn</code>.
	 * 
	 * Parameters:
	 *  - Key
	 *  - ColumnParent
	 *  - Predicate
	 *  - ConsistencyLevel
	 */
	GetCount(key []byte, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval450 int32, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * Performs a get_slice for column_parent and predicate for the given keys in parallel.
	 * 
	 * Parameters:
	 *  - Keys
	 *  - ColumnParent
	 *  - Predicate
	 *  - ConsistencyLevel
	 */
	MultigetSlice(keys thrift.TList, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval451 thrift.TMap, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * Perform a get_count in parallel on the given list<binary> keys. The return value maps keys to the count found.
	 * 
	 * Parameters:
	 *  - Keys
	 *  - ColumnParent
	 *  - Predicate
	 *  - ConsistencyLevel
	 */
	MultigetCount(keys thrift.TList, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval452 thrift.TMap, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * returns a subset of columns for a contiguous range of keys.
	 * 
	 * Parameters:
	 *  - ColumnParent
	 *  - Predicate
	 *  - RangeA1
	 *  - ConsistencyLevel
	 */
	GetRangeSlices(column_parent *ColumnParent, predicate *SlicePredicate, range_a1 *KeyRange, consistency_level ConsistencyLevel) (retval453 thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * Returns the subset of columns specified in SlicePredicate for the rows matching the IndexClause
	 * 
	 * Parameters:
	 *  - ColumnParent
	 *  - IndexClause
	 *  - ColumnPredicate
	 *  - ConsistencyLevel
	 */
	GetIndexedSlices(column_parent *ColumnParent, index_clause *IndexClause, column_predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval454 thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * Insert a Column at the given column_parent.column_family and optional column_parent.super_column.
	 * 
	 * Parameters:
	 *  - Key
	 *  - ColumnParent
	 *  - Column
	 *  - ConsistencyLevel
	 */
	Insert(key []byte, column_parent *ColumnParent, column *Column, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * Increment or decrement a counter.
	 * 
	 * Parameters:
	 *  - Key
	 *  - ColumnParent
	 *  - Column
	 *  - ConsistencyLevel
	 */
	Add(key []byte, column_parent *ColumnParent, column *CounterColumn, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * Remove data from the row specified by key at the granularity specified by column_path, and the given timestamp. Note
	 * that all the values in column_path besides column_path.column_family are truly optional: you can remove the entire
	 * row by just specifying the ColumnFamily, or you can remove a SuperColumn or a single Column by specifying those levels too.
	 * 
	 * Parameters:
	 *  - Key
	 *  - ColumnPath
	 *  - Timestamp
	 *  - ConsistencyLevel
	 */
	Remove(key []byte, column_path *ColumnPath, timestamp int64, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * Remove a counter at the specified location.
	 * Note that counters have limited support for deletes: if you remove a counter, you must wait to issue any following update
	 * until the delete has reached all the nodes and all of them have been fully compacted.
	 * 
	 * Parameters:
	 *  - Key
	 *  - Path
	 *  - ConsistencyLevel
	 */
	RemoveCounter(key []byte, path *ColumnPath, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 *   Mutate many columns or super columns for many row keys. See also: Mutation.
	 * 
	 *   mutation_map maps key to column family to a list of Mutation objects to take place at that scope.
	 * *
	 * 
	 * Parameters:
	 *  - MutationMap
	 *  - ConsistencyLevel
	 */
	BatchMutate(mutation_map thrift.TMap, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error)
	/**
	 * Truncate will mark and entire column family as deleted.
	 * From the user's perspective a successful call to truncate will result complete data deletion from cfname.
	 * Internally, however, disk space will not be immediatily released, as with all deletes in cassandra, this one
	 * only marks the data as deleted.
	 * The operation succeeds only if all hosts in the cluster at available and will throw an UnavailableException if
	 * some hosts are down.
	 * 
	 * Parameters:
	 *  - Cfname
	 */
	Truncate(cfname string) (ire *InvalidRequestException, ue *UnavailableException, err error)
	/**
	 * for each schema version present in the cluster, returns a list of nodes at that version.
	 * hosts that do not respond will be under the key DatabaseDescriptor.INITIAL_VERSION.
	 * the cluster is all on the same version if the size of the map is 1.
	 */
	DescribeSchemaVersions() (retval461 thrift.TMap, ire *InvalidRequestException, err error)
	/**
	 * list the defined keyspaces in this cluster
	 */
	DescribeKeyspaces() (retval462 thrift.TList, ire *InvalidRequestException, err error)
	/**
	 * get the cluster name
	 */
	DescribeClusterName() (retval463 string, err error)
	/**
	 * get the thrift api version
	 */
	DescribeVersion() (retval464 string, err error)
	/**
	 * get the token ring: a map of ranges to host addresses,
	 * represented as a set of TokenRange instead of a map from range
	 * to list of endpoints, because you can't use Thrift structs as
	 * map keys:
	 * https://issues.apache.org/jira/browse/THRIFT-162
	 * 
	 * for the same reason, we can't return a set here, even though
	 * order is neither important nor predictable.
	 * 
	 * Parameters:
	 *  - Keyspace
	 */
	DescribeRing(keyspace string) (retval465 thrift.TList, ire *InvalidRequestException, err error)
	/**
	 * returns the partitioner used by this cluster
	 */
	DescribePartitioner() (retval466 string, err error)
	/**
	 * returns the snitch used by this cluster
	 */
	DescribeSnitch() (retval467 string, err error)
	/**
	 * describe specified keyspace
	 * 
	 * Parameters:
	 *  - Keyspace
	 */
	DescribeKeyspace(keyspace string) (retval468 *KsDef, nfe *NotFoundException, ire *InvalidRequestException, err error)
	/**
	 * experimental API for hadoop/parallel query support.
	 * may change violently and without warning.
	 * 
	 * returns list of token strings such that first subrange is (list[0], list[1]],
	 * next is (list[1], list[2]], etc.
	 * 
	 * Parameters:
	 *  - CfName
	 *  - StartToken
	 *  - EndToken
	 *  - KeysPerSplit
	 */
	DescribeSplits(cfName string, start_token string, end_token string, keys_per_split int32) (retval469 thrift.TList, ire *InvalidRequestException, err error)
	/**
	 * adds a column family. returns the new schema id.
	 * 
	 * Parameters:
	 *  - CfDef
	 */
	SystemAddColumnFamily(cf_def *CfDef) (retval470 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)
	/**
	 * drops a column family. returns the new schema id.
	 * 
	 * Parameters:
	 *  - ColumnFamily
	 */
	SystemDropColumnFamily(column_family string) (retval471 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)
	/**
	 * adds a keyspace and any column families that are part of it. returns the new schema id.
	 * 
	 * Parameters:
	 *  - KsDef
	 */
	SystemAddKeyspace(ks_def *KsDef) (retval472 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)
	/**
	 * drops a keyspace and any column families that are part of it. returns the new schema id.
	 * 
	 * Parameters:
	 *  - Keyspace
	 */
	SystemDropKeyspace(keyspace string) (retval473 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)
	/**
	 * updates properties of a keyspace. returns the new schema id.
	 * 
	 * Parameters:
	 *  - KsDef
	 */
	SystemUpdateKeyspace(ks_def *KsDef) (retval474 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)
	/**
	 * updates properties of a column family. returns the new schema id.
	 * 
	 * Parameters:
	 *  - CfDef
	 */
	SystemUpdateColumnFamily(cf_def *CfDef) (retval475 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error)
	/**
	 * Executes a CQL (Cassandra Query Language) statement and returns a
	 * CqlResult containing the results.
	 * 
	 * Parameters:
	 *  - Query
	 *  - Compression
	 */
	ExecuteCqlQuery(query []byte, compression Compression) (retval476 *CqlResult, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, sde *SchemaDisagreementException, err error)
}

type CassandraClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewCassandraClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *CassandraClient {
	return &CassandraClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewCassandraClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *CassandraClient {
	return &CassandraClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

/**
 * Parameters:
 *  - AuthRequest
 */
func (p *CassandraClient) Login(auth_request *AuthenticationRequest) (authnx *AuthenticationException, authzx *AuthorizationException, err error) {
	err = p.SendLogin(auth_request)
	if err != nil {
		return
	}
	return p.RecvLogin()
}

func (p *CassandraClient) SendLogin(auth_request *AuthenticationRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("login", thrift.CALL, p.SeqId)
	args478 := NewLoginArgs()
	args478.AuthRequest = auth_request
	err = args478.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvLogin() (authnx *AuthenticationException, authzx *AuthorizationException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error480 := thrift.NewTApplicationExceptionDefault()
		var error481 error
		error481, err = error480.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error481
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result479 := NewLoginResult()
	err = result479.Read(iprot)
	iprot.ReadMessageEnd()
	if result479.Authnx != nil {
		authnx = result479.Authnx
	}
	if result479.Authzx != nil {
		authzx = result479.Authzx
	}
	return
}

/**
 * Parameters:
 *  - Keyspace
 */
func (p *CassandraClient) SetKeyspace(keyspace string) (ire *InvalidRequestException, err error) {
	err = p.SendSetKeyspace(keyspace)
	if err != nil {
		return
	}
	return p.RecvSetKeyspace()
}

func (p *CassandraClient) SendSetKeyspace(keyspace string) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("set_keyspace", thrift.CALL, p.SeqId)
	args483 := NewSetKeyspaceArgs()
	args483.Keyspace = keyspace
	err = args483.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvSetKeyspace() (ire *InvalidRequestException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error485 := thrift.NewTApplicationExceptionDefault()
		var error486 error
		error486, err = error485.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error486
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result484 := NewSetKeyspaceResult()
	err = result484.Read(iprot)
	iprot.ReadMessageEnd()
	if result484.Ire != nil {
		ire = result484.Ire
	}
	return
}

/**
 * Get the Column or SuperColumn at the given column_path. If no value is present, NotFoundException is thrown. (This is
 * the only method that can throw an exception under non-failure conditions.)
 * 
 * Parameters:
 *  - Key
 *  - ColumnPath
 *  - ConsistencyLevel
 */
func (p *CassandraClient) Get(key []byte, column_path *ColumnPath, consistency_level ConsistencyLevel) (retval487 *ColumnOrSuperColumn, ire *InvalidRequestException, nfe *NotFoundException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendGet(key, column_path, consistency_level)
	if err != nil {
		return
	}
	return p.RecvGet()
}

func (p *CassandraClient) SendGet(key []byte, column_path *ColumnPath, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("get", thrift.CALL, p.SeqId)
	args488 := NewGetArgs()
	args488.Key = key
	args488.ColumnPath = column_path
	args488.ConsistencyLevel = consistency_level
	err = args488.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvGet() (value *ColumnOrSuperColumn, ire *InvalidRequestException, nfe *NotFoundException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error490 := thrift.NewTApplicationExceptionDefault()
		var error491 error
		error491, err = error490.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error491
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result489 := NewGetResult()
	err = result489.Read(iprot)
	iprot.ReadMessageEnd()
	value = result489.Success
	if result489.Ire != nil {
		ire = result489.Ire
	}
	if result489.Nfe != nil {
		nfe = result489.Nfe
	}
	if result489.Ue != nil {
		ue = result489.Ue
	}
	if result489.Te != nil {
		te = result489.Te
	}
	return
}

/**
 * Get the group of columns contained by column_parent (either a ColumnFamily name or a ColumnFamily/SuperColumn name
 * pair) specified by the given SlicePredicate. If no matching values are found, an empty list is returned.
 * 
 * Parameters:
 *  - Key
 *  - ColumnParent
 *  - Predicate
 *  - ConsistencyLevel
 */
func (p *CassandraClient) GetSlice(key []byte, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval492 thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendGetSlice(key, column_parent, predicate, consistency_level)
	if err != nil {
		return
	}
	return p.RecvGetSlice()
}

func (p *CassandraClient) SendGetSlice(key []byte, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("get_slice", thrift.CALL, p.SeqId)
	args493 := NewGetSliceArgs()
	args493.Key = key
	args493.ColumnParent = column_parent
	args493.Predicate = predicate
	args493.ConsistencyLevel = consistency_level
	err = args493.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvGetSlice() (value thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error495 := thrift.NewTApplicationExceptionDefault()
		var error496 error
		error496, err = error495.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error496
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result494 := NewGetSliceResult()
	err = result494.Read(iprot)
	iprot.ReadMessageEnd()
	value = result494.Success
	if result494.Ire != nil {
		ire = result494.Ire
	}
	if result494.Ue != nil {
		ue = result494.Ue
	}
	if result494.Te != nil {
		te = result494.Te
	}
	return
}

/**
 * returns the number of columns matching <code>predicate</code> for a particular <code>key</code>,
 * <code>ColumnFamily</code> and optionally <code>SuperColumn</code>.
 * 
 * Parameters:
 *  - Key
 *  - ColumnParent
 *  - Predicate
 *  - ConsistencyLevel
 */
func (p *CassandraClient) GetCount(key []byte, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval497 int32, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendGetCount(key, column_parent, predicate, consistency_level)
	if err != nil {
		return
	}
	return p.RecvGetCount()
}

func (p *CassandraClient) SendGetCount(key []byte, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("get_count", thrift.CALL, p.SeqId)
	args498 := NewGetCountArgs()
	args498.Key = key
	args498.ColumnParent = column_parent
	args498.Predicate = predicate
	args498.ConsistencyLevel = consistency_level
	err = args498.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvGetCount() (value int32, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error500 := thrift.NewTApplicationExceptionDefault()
		var error501 error
		error501, err = error500.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error501
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result499 := NewGetCountResult()
	err = result499.Read(iprot)
	iprot.ReadMessageEnd()
	value = result499.Success
	if result499.Ire != nil {
		ire = result499.Ire
	}
	if result499.Ue != nil {
		ue = result499.Ue
	}
	if result499.Te != nil {
		te = result499.Te
	}
	return
}

/**
 * Performs a get_slice for column_parent and predicate for the given keys in parallel.
 * 
 * Parameters:
 *  - Keys
 *  - ColumnParent
 *  - Predicate
 *  - ConsistencyLevel
 */
func (p *CassandraClient) MultigetSlice(keys thrift.TList, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval502 thrift.TMap, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendMultigetSlice(keys, column_parent, predicate, consistency_level)
	if err != nil {
		return
	}
	return p.RecvMultigetSlice()
}

func (p *CassandraClient) SendMultigetSlice(keys thrift.TList, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("multiget_slice", thrift.CALL, p.SeqId)
	args503 := NewMultigetSliceArgs()
	args503.Keys = keys
	args503.ColumnParent = column_parent
	args503.Predicate = predicate
	args503.ConsistencyLevel = consistency_level
	err = args503.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvMultigetSlice() (value thrift.TMap, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error505 := thrift.NewTApplicationExceptionDefault()
		var error506 error
		error506, err = error505.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error506
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result504 := NewMultigetSliceResult()
	err = result504.Read(iprot)
	iprot.ReadMessageEnd()
	value = result504.Success
	if result504.Ire != nil {
		ire = result504.Ire
	}
	if result504.Ue != nil {
		ue = result504.Ue
	}
	if result504.Te != nil {
		te = result504.Te
	}
	return
}

/**
 * Perform a get_count in parallel on the given list<binary> keys. The return value maps keys to the count found.
 * 
 * Parameters:
 *  - Keys
 *  - ColumnParent
 *  - Predicate
 *  - ConsistencyLevel
 */
func (p *CassandraClient) MultigetCount(keys thrift.TList, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval507 thrift.TMap, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendMultigetCount(keys, column_parent, predicate, consistency_level)
	if err != nil {
		return
	}
	return p.RecvMultigetCount()
}

func (p *CassandraClient) SendMultigetCount(keys thrift.TList, column_parent *ColumnParent, predicate *SlicePredicate, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("multiget_count", thrift.CALL, p.SeqId)
	args508 := NewMultigetCountArgs()
	args508.Keys = keys
	args508.ColumnParent = column_parent
	args508.Predicate = predicate
	args508.ConsistencyLevel = consistency_level
	err = args508.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvMultigetCount() (value thrift.TMap, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error510 := thrift.NewTApplicationExceptionDefault()
		var error511 error
		error511, err = error510.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error511
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result509 := NewMultigetCountResult()
	err = result509.Read(iprot)
	iprot.ReadMessageEnd()
	value = result509.Success
	if result509.Ire != nil {
		ire = result509.Ire
	}
	if result509.Ue != nil {
		ue = result509.Ue
	}
	if result509.Te != nil {
		te = result509.Te
	}
	return
}

/**
 * returns a subset of columns for a contiguous range of keys.
 * 
 * Parameters:
 *  - ColumnParent
 *  - Predicate
 *  - RangeA1
 *  - ConsistencyLevel
 */
func (p *CassandraClient) GetRangeSlices(column_parent *ColumnParent, predicate *SlicePredicate, range_a1 *KeyRange, consistency_level ConsistencyLevel) (retval512 thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendGetRangeSlices(column_parent, predicate, range_a1, consistency_level)
	if err != nil {
		return
	}
	return p.RecvGetRangeSlices()
}

func (p *CassandraClient) SendGetRangeSlices(column_parent *ColumnParent, predicate *SlicePredicate, range_a1 *KeyRange, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("get_range_slices", thrift.CALL, p.SeqId)
	args513 := NewGetRangeSlicesArgs()
	args513.ColumnParent = column_parent
	args513.Predicate = predicate
	args513.RangeA1 = range_a1
	args513.ConsistencyLevel = consistency_level
	err = args513.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvGetRangeSlices() (value thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error515 := thrift.NewTApplicationExceptionDefault()
		var error516 error
		error516, err = error515.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error516
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result514 := NewGetRangeSlicesResult()
	err = result514.Read(iprot)
	iprot.ReadMessageEnd()
	value = result514.Success
	if result514.Ire != nil {
		ire = result514.Ire
	}
	if result514.Ue != nil {
		ue = result514.Ue
	}
	if result514.Te != nil {
		te = result514.Te
	}
	return
}

/**
 * Returns the subset of columns specified in SlicePredicate for the rows matching the IndexClause
 * 
 * Parameters:
 *  - ColumnParent
 *  - IndexClause
 *  - ColumnPredicate
 *  - ConsistencyLevel
 */
func (p *CassandraClient) GetIndexedSlices(column_parent *ColumnParent, index_clause *IndexClause, column_predicate *SlicePredicate, consistency_level ConsistencyLevel) (retval517 thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendGetIndexedSlices(column_parent, index_clause, column_predicate, consistency_level)
	if err != nil {
		return
	}
	return p.RecvGetIndexedSlices()
}

func (p *CassandraClient) SendGetIndexedSlices(column_parent *ColumnParent, index_clause *IndexClause, column_predicate *SlicePredicate, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("get_indexed_slices", thrift.CALL, p.SeqId)
	args518 := NewGetIndexedSlicesArgs()
	args518.ColumnParent = column_parent
	args518.IndexClause = index_clause
	args518.ColumnPredicate = column_predicate
	args518.ConsistencyLevel = consistency_level
	err = args518.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvGetIndexedSlices() (value thrift.TList, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error520 := thrift.NewTApplicationExceptionDefault()
		var error521 error
		error521, err = error520.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error521
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result519 := NewGetIndexedSlicesResult()
	err = result519.Read(iprot)
	iprot.ReadMessageEnd()
	value = result519.Success
	if result519.Ire != nil {
		ire = result519.Ire
	}
	if result519.Ue != nil {
		ue = result519.Ue
	}
	if result519.Te != nil {
		te = result519.Te
	}
	return
}

/**
 * Insert a Column at the given column_parent.column_family and optional column_parent.super_column.
 * 
 * Parameters:
 *  - Key
 *  - ColumnParent
 *  - Column
 *  - ConsistencyLevel
 */
func (p *CassandraClient) Insert(key []byte, column_parent *ColumnParent, column *Column, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendInsert(key, column_parent, column, consistency_level)
	if err != nil {
		return
	}
	return p.RecvInsert()
}

func (p *CassandraClient) SendInsert(key []byte, column_parent *ColumnParent, column *Column, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("insert", thrift.CALL, p.SeqId)
	args523 := NewInsertArgs()
	args523.Key = key
	args523.ColumnParent = column_parent
	args523.Column = column
	args523.ConsistencyLevel = consistency_level
	err = args523.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvInsert() (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error525 := thrift.NewTApplicationExceptionDefault()
		var error526 error
		error526, err = error525.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error526
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result524 := NewInsertResult()
	err = result524.Read(iprot)
	iprot.ReadMessageEnd()
	if result524.Ire != nil {
		ire = result524.Ire
	}
	if result524.Ue != nil {
		ue = result524.Ue
	}
	if result524.Te != nil {
		te = result524.Te
	}
	return
}

/**
 * Increment or decrement a counter.
 * 
 * Parameters:
 *  - Key
 *  - ColumnParent
 *  - Column
 *  - ConsistencyLevel
 */
func (p *CassandraClient) Add(key []byte, column_parent *ColumnParent, column *CounterColumn, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendAdd(key, column_parent, column, consistency_level)
	if err != nil {
		return
	}
	return p.RecvAdd()
}

func (p *CassandraClient) SendAdd(key []byte, column_parent *ColumnParent, column *CounterColumn, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("add", thrift.CALL, p.SeqId)
	args528 := NewAddArgs()
	args528.Key = key
	args528.ColumnParent = column_parent
	args528.Column = column
	args528.ConsistencyLevel = consistency_level
	err = args528.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvAdd() (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error530 := thrift.NewTApplicationExceptionDefault()
		var error531 error
		error531, err = error530.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error531
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result529 := NewAddResult()
	err = result529.Read(iprot)
	iprot.ReadMessageEnd()
	if result529.Ire != nil {
		ire = result529.Ire
	}
	if result529.Ue != nil {
		ue = result529.Ue
	}
	if result529.Te != nil {
		te = result529.Te
	}
	return
}

/**
 * Remove data from the row specified by key at the granularity specified by column_path, and the given timestamp. Note
 * that all the values in column_path besides column_path.column_family are truly optional: you can remove the entire
 * row by just specifying the ColumnFamily, or you can remove a SuperColumn or a single Column by specifying those levels too.
 * 
 * Parameters:
 *  - Key
 *  - ColumnPath
 *  - Timestamp
 *  - ConsistencyLevel
 */
func (p *CassandraClient) Remove(key []byte, column_path *ColumnPath, timestamp int64, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendRemove(key, column_path, timestamp, consistency_level)
	if err != nil {
		return
	}
	return p.RecvRemove()
}

func (p *CassandraClient) SendRemove(key []byte, column_path *ColumnPath, timestamp int64, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("remove", thrift.CALL, p.SeqId)
	args533 := NewRemoveArgs()
	args533.Key = key
	args533.ColumnPath = column_path
	args533.Timestamp = timestamp
	args533.ConsistencyLevel = consistency_level
	err = args533.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvRemove() (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error535 := thrift.NewTApplicationExceptionDefault()
		var error536 error
		error536, err = error535.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error536
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result534 := NewRemoveResult()
	err = result534.Read(iprot)
	iprot.ReadMessageEnd()
	if result534.Ire != nil {
		ire = result534.Ire
	}
	if result534.Ue != nil {
		ue = result534.Ue
	}
	if result534.Te != nil {
		te = result534.Te
	}
	return
}

/**
 * Remove a counter at the specified location.
 * Note that counters have limited support for deletes: if you remove a counter, you must wait to issue any following update
 * until the delete has reached all the nodes and all of them have been fully compacted.
 * 
 * Parameters:
 *  - Key
 *  - Path
 *  - ConsistencyLevel
 */
func (p *CassandraClient) RemoveCounter(key []byte, path *ColumnPath, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendRemoveCounter(key, path, consistency_level)
	if err != nil {
		return
	}
	return p.RecvRemoveCounter()
}

func (p *CassandraClient) SendRemoveCounter(key []byte, path *ColumnPath, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("remove_counter", thrift.CALL, p.SeqId)
	args538 := NewRemoveCounterArgs()
	args538.Key = key
	args538.Path = path
	args538.ConsistencyLevel = consistency_level
	err = args538.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvRemoveCounter() (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error540 := thrift.NewTApplicationExceptionDefault()
		var error541 error
		error541, err = error540.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error541
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result539 := NewRemoveCounterResult()
	err = result539.Read(iprot)
	iprot.ReadMessageEnd()
	if result539.Ire != nil {
		ire = result539.Ire
	}
	if result539.Ue != nil {
		ue = result539.Ue
	}
	if result539.Te != nil {
		te = result539.Te
	}
	return
}

/**
 *   Mutate many columns or super columns for many row keys. See also: Mutation.
 * 
 *   mutation_map maps key to column family to a list of Mutation objects to take place at that scope.
 * *
 * 
 * Parameters:
 *  - MutationMap
 *  - ConsistencyLevel
 */
func (p *CassandraClient) BatchMutate(mutation_map thrift.TMap, consistency_level ConsistencyLevel) (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	err = p.SendBatchMutate(mutation_map, consistency_level)
	if err != nil {
		return
	}
	return p.RecvBatchMutate()
}

func (p *CassandraClient) SendBatchMutate(mutation_map thrift.TMap, consistency_level ConsistencyLevel) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("batch_mutate", thrift.CALL, p.SeqId)
	args543 := NewBatchMutateArgs()
	args543.MutationMap = mutation_map
	args543.ConsistencyLevel = consistency_level
	err = args543.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvBatchMutate() (ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error545 := thrift.NewTApplicationExceptionDefault()
		var error546 error
		error546, err = error545.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error546
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result544 := NewBatchMutateResult()
	err = result544.Read(iprot)
	iprot.ReadMessageEnd()
	if result544.Ire != nil {
		ire = result544.Ire
	}
	if result544.Ue != nil {
		ue = result544.Ue
	}
	if result544.Te != nil {
		te = result544.Te
	}
	return
}

/**
 * Truncate will mark and entire column family as deleted.
 * From the user's perspective a successful call to truncate will result complete data deletion from cfname.
 * Internally, however, disk space will not be immediatily released, as with all deletes in cassandra, this one
 * only marks the data as deleted.
 * The operation succeeds only if all hosts in the cluster at available and will throw an UnavailableException if
 * some hosts are down.
 * 
 * Parameters:
 *  - Cfname
 */
func (p *CassandraClient) Truncate(cfname string) (ire *InvalidRequestException, ue *UnavailableException, err error) {
	err = p.SendTruncate(cfname)
	if err != nil {
		return
	}
	return p.RecvTruncate()
}

func (p *CassandraClient) SendTruncate(cfname string) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("truncate", thrift.CALL, p.SeqId)
	args548 := NewTruncateArgs()
	args548.Cfname = cfname
	err = args548.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvTruncate() (ire *InvalidRequestException, ue *UnavailableException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error550 := thrift.NewTApplicationExceptionDefault()
		var error551 error
		error551, err = error550.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error551
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result549 := NewTruncateResult()
	err = result549.Read(iprot)
	iprot.ReadMessageEnd()
	if result549.Ire != nil {
		ire = result549.Ire
	}
	if result549.Ue != nil {
		ue = result549.Ue
	}
	return
}

/**
 * for each schema version present in the cluster, returns a list of nodes at that version.
 * hosts that do not respond will be under the key DatabaseDescriptor.INITIAL_VERSION.
 * the cluster is all on the same version if the size of the map is 1.
 */
func (p *CassandraClient) DescribeSchemaVersions() (retval552 thrift.TMap, ire *InvalidRequestException, err error) {
	err = p.SendDescribeSchemaVersions()
	if err != nil {
		return
	}
	return p.RecvDescribeSchemaVersions()
}

func (p *CassandraClient) SendDescribeSchemaVersions() (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("describe_schema_versions", thrift.CALL, p.SeqId)
	args553 := NewDescribeSchemaVersionsArgs()
	err = args553.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvDescribeSchemaVersions() (value thrift.TMap, ire *InvalidRequestException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error555 := thrift.NewTApplicationExceptionDefault()
		var error556 error
		error556, err = error555.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error556
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result554 := NewDescribeSchemaVersionsResult()
	err = result554.Read(iprot)
	iprot.ReadMessageEnd()
	value = result554.Success
	if result554.Ire != nil {
		ire = result554.Ire
	}
	return
}

/**
 * list the defined keyspaces in this cluster
 */
func (p *CassandraClient) DescribeKeyspaces() (retval557 thrift.TList, ire *InvalidRequestException, err error) {
	err = p.SendDescribeKeyspaces()
	if err != nil {
		return
	}
	return p.RecvDescribeKeyspaces()
}

func (p *CassandraClient) SendDescribeKeyspaces() (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("describe_keyspaces", thrift.CALL, p.SeqId)
	args558 := NewDescribeKeyspacesArgs()
	err = args558.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvDescribeKeyspaces() (value thrift.TList, ire *InvalidRequestException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error560 := thrift.NewTApplicationExceptionDefault()
		var error561 error
		error561, err = error560.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error561
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result559 := NewDescribeKeyspacesResult()
	err = result559.Read(iprot)
	iprot.ReadMessageEnd()
	value = result559.Success
	if result559.Ire != nil {
		ire = result559.Ire
	}
	return
}

/**
 * get the cluster name
 */
func (p *CassandraClient) DescribeClusterName() (retval562 string, err error) {
	err = p.SendDescribeClusterName()
	if err != nil {
		return
	}
	return p.RecvDescribeClusterName()
}

func (p *CassandraClient) SendDescribeClusterName() (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("describe_cluster_name", thrift.CALL, p.SeqId)
	args563 := NewDescribeClusterNameArgs()
	err = args563.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvDescribeClusterName() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error565 := thrift.NewTApplicationExceptionDefault()
		var error566 error
		error566, err = error565.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error566
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result564 := NewDescribeClusterNameResult()
	err = result564.Read(iprot)
	iprot.ReadMessageEnd()
	value = result564.Success
	return
}

/**
 * get the thrift api version
 */
func (p *CassandraClient) DescribeVersion() (retval567 string, err error) {
	err = p.SendDescribeVersion()
	if err != nil {
		return
	}
	return p.RecvDescribeVersion()
}

func (p *CassandraClient) SendDescribeVersion() (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("describe_version", thrift.CALL, p.SeqId)
	args568 := NewDescribeVersionArgs()
	err = args568.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvDescribeVersion() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error570 := thrift.NewTApplicationExceptionDefault()
		var error571 error
		error571, err = error570.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error571
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result569 := NewDescribeVersionResult()
	err = result569.Read(iprot)
	iprot.ReadMessageEnd()
	value = result569.Success
	return
}

/**
 * get the token ring: a map of ranges to host addresses,
 * represented as a set of TokenRange instead of a map from range
 * to list of endpoints, because you can't use Thrift structs as
 * map keys:
 * https://issues.apache.org/jira/browse/THRIFT-162
 * 
 * for the same reason, we can't return a set here, even though
 * order is neither important nor predictable.
 * 
 * Parameters:
 *  - Keyspace
 */
func (p *CassandraClient) DescribeRing(keyspace string) (retval572 thrift.TList, ire *InvalidRequestException, err error) {
	err = p.SendDescribeRing(keyspace)
	if err != nil {
		return
	}
	return p.RecvDescribeRing()
}

func (p *CassandraClient) SendDescribeRing(keyspace string) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("describe_ring", thrift.CALL, p.SeqId)
	args573 := NewDescribeRingArgs()
	args573.Keyspace = keyspace
	err = args573.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvDescribeRing() (value thrift.TList, ire *InvalidRequestException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error575 := thrift.NewTApplicationExceptionDefault()
		var error576 error
		error576, err = error575.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error576
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result574 := NewDescribeRingResult()
	err = result574.Read(iprot)
	iprot.ReadMessageEnd()
	value = result574.Success
	if result574.Ire != nil {
		ire = result574.Ire
	}
	return
}

/**
 * returns the partitioner used by this cluster
 */
func (p *CassandraClient) DescribePartitioner() (retval577 string, err error) {
	err = p.SendDescribePartitioner()
	if err != nil {
		return
	}
	return p.RecvDescribePartitioner()
}

func (p *CassandraClient) SendDescribePartitioner() (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("describe_partitioner", thrift.CALL, p.SeqId)
	args578 := NewDescribePartitionerArgs()
	err = args578.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvDescribePartitioner() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error580 := thrift.NewTApplicationExceptionDefault()
		var error581 error
		error581, err = error580.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error581
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result579 := NewDescribePartitionerResult()
	err = result579.Read(iprot)
	iprot.ReadMessageEnd()
	value = result579.Success
	return
}

/**
 * returns the snitch used by this cluster
 */
func (p *CassandraClient) DescribeSnitch() (retval582 string, err error) {
	err = p.SendDescribeSnitch()
	if err != nil {
		return
	}
	return p.RecvDescribeSnitch()
}

func (p *CassandraClient) SendDescribeSnitch() (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("describe_snitch", thrift.CALL, p.SeqId)
	args583 := NewDescribeSnitchArgs()
	err = args583.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvDescribeSnitch() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error585 := thrift.NewTApplicationExceptionDefault()
		var error586 error
		error586, err = error585.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error586
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result584 := NewDescribeSnitchResult()
	err = result584.Read(iprot)
	iprot.ReadMessageEnd()
	value = result584.Success
	return
}

/**
 * describe specified keyspace
 * 
 * Parameters:
 *  - Keyspace
 */
func (p *CassandraClient) DescribeKeyspace(keyspace string) (retval587 *KsDef, nfe *NotFoundException, ire *InvalidRequestException, err error) {
	err = p.SendDescribeKeyspace(keyspace)
	if err != nil {
		return
	}
	return p.RecvDescribeKeyspace()
}

func (p *CassandraClient) SendDescribeKeyspace(keyspace string) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("describe_keyspace", thrift.CALL, p.SeqId)
	args588 := NewDescribeKeyspaceArgs()
	args588.Keyspace = keyspace
	err = args588.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvDescribeKeyspace() (value *KsDef, nfe *NotFoundException, ire *InvalidRequestException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error590 := thrift.NewTApplicationExceptionDefault()
		var error591 error
		error591, err = error590.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error591
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result589 := NewDescribeKeyspaceResult()
	err = result589.Read(iprot)
	iprot.ReadMessageEnd()
	value = result589.Success
	if result589.Nfe != nil {
		nfe = result589.Nfe
	}
	if result589.Ire != nil {
		ire = result589.Ire
	}
	return
}

/**
 * experimental API for hadoop/parallel query support.
 * may change violently and without warning.
 * 
 * returns list of token strings such that first subrange is (list[0], list[1]],
 * next is (list[1], list[2]], etc.
 * 
 * Parameters:
 *  - CfName
 *  - StartToken
 *  - EndToken
 *  - KeysPerSplit
 */
func (p *CassandraClient) DescribeSplits(cfName string, start_token string, end_token string, keys_per_split int32) (retval592 thrift.TList, ire *InvalidRequestException, err error) {
	err = p.SendDescribeSplits(cfName, start_token, end_token, keys_per_split)
	if err != nil {
		return
	}
	return p.RecvDescribeSplits()
}

func (p *CassandraClient) SendDescribeSplits(cfName string, start_token string, end_token string, keys_per_split int32) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("describe_splits", thrift.CALL, p.SeqId)
	args593 := NewDescribeSplitsArgs()
	args593.CfName = cfName
	args593.StartToken = start_token
	args593.EndToken = end_token
	args593.KeysPerSplit = keys_per_split
	err = args593.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvDescribeSplits() (value thrift.TList, ire *InvalidRequestException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error595 := thrift.NewTApplicationExceptionDefault()
		var error596 error
		error596, err = error595.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error596
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result594 := NewDescribeSplitsResult()
	err = result594.Read(iprot)
	iprot.ReadMessageEnd()
	value = result594.Success
	if result594.Ire != nil {
		ire = result594.Ire
	}
	return
}

/**
 * adds a column family. returns the new schema id.
 * 
 * Parameters:
 *  - CfDef
 */
func (p *CassandraClient) SystemAddColumnFamily(cf_def *CfDef) (retval597 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	err = p.SendSystemAddColumnFamily(cf_def)
	if err != nil {
		return
	}
	return p.RecvSystemAddColumnFamily()
}

func (p *CassandraClient) SendSystemAddColumnFamily(cf_def *CfDef) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("system_add_column_family", thrift.CALL, p.SeqId)
	args598 := NewSystemAddColumnFamilyArgs()
	args598.CfDef = cf_def
	err = args598.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvSystemAddColumnFamily() (value string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error600 := thrift.NewTApplicationExceptionDefault()
		var error601 error
		error601, err = error600.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error601
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result599 := NewSystemAddColumnFamilyResult()
	err = result599.Read(iprot)
	iprot.ReadMessageEnd()
	value = result599.Success
	if result599.Ire != nil {
		ire = result599.Ire
	}
	if result599.Sde != nil {
		sde = result599.Sde
	}
	return
}

/**
 * drops a column family. returns the new schema id.
 * 
 * Parameters:
 *  - ColumnFamily
 */
func (p *CassandraClient) SystemDropColumnFamily(column_family string) (retval602 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	err = p.SendSystemDropColumnFamily(column_family)
	if err != nil {
		return
	}
	return p.RecvSystemDropColumnFamily()
}

func (p *CassandraClient) SendSystemDropColumnFamily(column_family string) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("system_drop_column_family", thrift.CALL, p.SeqId)
	args603 := NewSystemDropColumnFamilyArgs()
	args603.ColumnFamily = column_family
	err = args603.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvSystemDropColumnFamily() (value string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error605 := thrift.NewTApplicationExceptionDefault()
		var error606 error
		error606, err = error605.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error606
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result604 := NewSystemDropColumnFamilyResult()
	err = result604.Read(iprot)
	iprot.ReadMessageEnd()
	value = result604.Success
	if result604.Ire != nil {
		ire = result604.Ire
	}
	if result604.Sde != nil {
		sde = result604.Sde
	}
	return
}

/**
 * adds a keyspace and any column families that are part of it. returns the new schema id.
 * 
 * Parameters:
 *  - KsDef
 */
func (p *CassandraClient) SystemAddKeyspace(ks_def *KsDef) (retval607 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	err = p.SendSystemAddKeyspace(ks_def)
	if err != nil {
		return
	}
	return p.RecvSystemAddKeyspace()
}

func (p *CassandraClient) SendSystemAddKeyspace(ks_def *KsDef) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("system_add_keyspace", thrift.CALL, p.SeqId)
	args608 := NewSystemAddKeyspaceArgs()
	args608.KsDef = ks_def
	err = args608.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvSystemAddKeyspace() (value string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error610 := thrift.NewTApplicationExceptionDefault()
		var error611 error
		error611, err = error610.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error611
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result609 := NewSystemAddKeyspaceResult()
	err = result609.Read(iprot)
	iprot.ReadMessageEnd()
	value = result609.Success
	if result609.Ire != nil {
		ire = result609.Ire
	}
	if result609.Sde != nil {
		sde = result609.Sde
	}
	return
}

/**
 * drops a keyspace and any column families that are part of it. returns the new schema id.
 * 
 * Parameters:
 *  - Keyspace
 */
func (p *CassandraClient) SystemDropKeyspace(keyspace string) (retval612 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	err = p.SendSystemDropKeyspace(keyspace)
	if err != nil {
		return
	}
	return p.RecvSystemDropKeyspace()
}

func (p *CassandraClient) SendSystemDropKeyspace(keyspace string) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("system_drop_keyspace", thrift.CALL, p.SeqId)
	args613 := NewSystemDropKeyspaceArgs()
	args613.Keyspace = keyspace
	err = args613.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvSystemDropKeyspace() (value string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error615 := thrift.NewTApplicationExceptionDefault()
		var error616 error
		error616, err = error615.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error616
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result614 := NewSystemDropKeyspaceResult()
	err = result614.Read(iprot)
	iprot.ReadMessageEnd()
	value = result614.Success
	if result614.Ire != nil {
		ire = result614.Ire
	}
	if result614.Sde != nil {
		sde = result614.Sde
	}
	return
}

/**
 * updates properties of a keyspace. returns the new schema id.
 * 
 * Parameters:
 *  - KsDef
 */
func (p *CassandraClient) SystemUpdateKeyspace(ks_def *KsDef) (retval617 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	err = p.SendSystemUpdateKeyspace(ks_def)
	if err != nil {
		return
	}
	return p.RecvSystemUpdateKeyspace()
}

func (p *CassandraClient) SendSystemUpdateKeyspace(ks_def *KsDef) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("system_update_keyspace", thrift.CALL, p.SeqId)
	args618 := NewSystemUpdateKeyspaceArgs()
	args618.KsDef = ks_def
	err = args618.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvSystemUpdateKeyspace() (value string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error620 := thrift.NewTApplicationExceptionDefault()
		var error621 error
		error621, err = error620.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error621
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result619 := NewSystemUpdateKeyspaceResult()
	err = result619.Read(iprot)
	iprot.ReadMessageEnd()
	value = result619.Success
	if result619.Ire != nil {
		ire = result619.Ire
	}
	if result619.Sde != nil {
		sde = result619.Sde
	}
	return
}

/**
 * updates properties of a column family. returns the new schema id.
 * 
 * Parameters:
 *  - CfDef
 */
func (p *CassandraClient) SystemUpdateColumnFamily(cf_def *CfDef) (retval622 string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	err = p.SendSystemUpdateColumnFamily(cf_def)
	if err != nil {
		return
	}
	return p.RecvSystemUpdateColumnFamily()
}

func (p *CassandraClient) SendSystemUpdateColumnFamily(cf_def *CfDef) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("system_update_column_family", thrift.CALL, p.SeqId)
	args623 := NewSystemUpdateColumnFamilyArgs()
	args623.CfDef = cf_def
	err = args623.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvSystemUpdateColumnFamily() (value string, ire *InvalidRequestException, sde *SchemaDisagreementException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error625 := thrift.NewTApplicationExceptionDefault()
		var error626 error
		error626, err = error625.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error626
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result624 := NewSystemUpdateColumnFamilyResult()
	err = result624.Read(iprot)
	iprot.ReadMessageEnd()
	value = result624.Success
	if result624.Ire != nil {
		ire = result624.Ire
	}
	if result624.Sde != nil {
		sde = result624.Sde
	}
	return
}

/**
 * Executes a CQL (Cassandra Query Language) statement and returns a
 * CqlResult containing the results.
 * 
 * Parameters:
 *  - Query
 *  - Compression
 */
func (p *CassandraClient) ExecuteCqlQuery(query []byte, compression Compression) (retval627 *CqlResult, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, sde *SchemaDisagreementException, err error) {
	err = p.SendExecuteCqlQuery(query, compression)
	if err != nil {
		return
	}
	return p.RecvExecuteCqlQuery()
}

func (p *CassandraClient) SendExecuteCqlQuery(query []byte, compression Compression) (err error) {
	oprot := p.OutputProtocol
	if oprot != nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("execute_cql_query", thrift.CALL, p.SeqId)
	args628 := NewExecuteCqlQueryArgs()
	args628.Query = query
	args628.Compression = compression
	err = args628.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Transport().Flush()
	return
}

func (p *CassandraClient) RecvExecuteCqlQuery() (value *CqlResult, ire *InvalidRequestException, ue *UnavailableException, te *TimedOutException, sde *SchemaDisagreementException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error630 := thrift.NewTApplicationExceptionDefault()
		var error631 error
		error631, err = error630.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error631
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result629 := NewExecuteCqlQueryResult()
	err = result629.Read(iprot)
	iprot.ReadMessageEnd()
	value = result629.Success
	if result629.Ire != nil {
		ire = result629.Ire
	}
	if result629.Ue != nil {
		ue = result629.Ue
	}
	if result629.Te != nil {
		te = result629.Te
	}
	if result629.Sde != nil {
		sde = result629.Sde
	}
	return
}

type CassandraProcessor struct {
	handler      ICassandra
	processorMap map[string]thrift.TProcessorFunction
}

func (p *CassandraProcessor) Handler() ICassandra {
	return p.handler
}

func (p *CassandraProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *CassandraProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, exists bool) {
	processor, exists = p.processorMap[key]
	return processor, exists
}

func (p *CassandraProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewCassandraProcessor(handler ICassandra) *CassandraProcessor {

	self632 := &CassandraProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self632.processorMap["login"] = &cassandraProcessorLogin{handler: handler}
	self632.processorMap["set_keyspace"] = &cassandraProcessorSetKeyspace{handler: handler}
	self632.processorMap["get"] = &cassandraProcessorGet{handler: handler}
	self632.processorMap["get_slice"] = &cassandraProcessorGetSlice{handler: handler}
	self632.processorMap["get_count"] = &cassandraProcessorGetCount{handler: handler}
	self632.processorMap["multiget_slice"] = &cassandraProcessorMultigetSlice{handler: handler}
	self632.processorMap["multiget_count"] = &cassandraProcessorMultigetCount{handler: handler}
	self632.processorMap["get_range_slices"] = &cassandraProcessorGetRangeSlices{handler: handler}
	self632.processorMap["get_indexed_slices"] = &cassandraProcessorGetIndexedSlices{handler: handler}
	self632.processorMap["insert"] = &cassandraProcessorInsert{handler: handler}
	self632.processorMap["add"] = &cassandraProcessorAdd{handler: handler}
	self632.processorMap["remove"] = &cassandraProcessorRemove{handler: handler}
	self632.processorMap["remove_counter"] = &cassandraProcessorRemoveCounter{handler: handler}
	self632.processorMap["batch_mutate"] = &cassandraProcessorBatchMutate{handler: handler}
	self632.processorMap["truncate"] = &cassandraProcessorTruncate{handler: handler}
	self632.processorMap["describe_schema_versions"] = &cassandraProcessorDescribeSchemaVersions{handler: handler}
	self632.processorMap["describe_keyspaces"] = &cassandraProcessorDescribeKeyspaces{handler: handler}
	self632.processorMap["describe_cluster_name"] = &cassandraProcessorDescribeClusterName{handler: handler}
	self632.processorMap["describe_version"] = &cassandraProcessorDescribeVersion{handler: handler}
	self632.processorMap["describe_ring"] = &cassandraProcessorDescribeRing{handler: handler}
	self632.processorMap["describe_partitioner"] = &cassandraProcessorDescribePartitioner{handler: handler}
	self632.processorMap["describe_snitch"] = &cassandraProcessorDescribeSnitch{handler: handler}
	self632.processorMap["describe_keyspace"] = &cassandraProcessorDescribeKeyspace{handler: handler}
	self632.processorMap["describe_splits"] = &cassandraProcessorDescribeSplits{handler: handler}
	self632.processorMap["system_add_column_family"] = &cassandraProcessorSystemAddColumnFamily{handler: handler}
	self632.processorMap["system_drop_column_family"] = &cassandraProcessorSystemDropColumnFamily{handler: handler}
	self632.processorMap["system_add_keyspace"] = &cassandraProcessorSystemAddKeyspace{handler: handler}
	self632.processorMap["system_drop_keyspace"] = &cassandraProcessorSystemDropKeyspace{handler: handler}
	self632.processorMap["system_update_keyspace"] = &cassandraProcessorSystemUpdateKeyspace{handler: handler}
	self632.processorMap["system_update_column_family"] = &cassandraProcessorSystemUpdateColumnFamily{handler: handler}
	self632.processorMap["execute_cql_query"] = &cassandraProcessorExecuteCqlQuery{handler: handler}
	return self632
}

func (p *CassandraProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	process, nameFound := p.GetProcessorFunction(name)
	if !nameFound || process == nil {
		iprot.Skip(thrift.STRUCT)
		iprot.ReadMessageEnd()
		x633 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
		oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
		x633.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return false, x633
	}
	return process.Process(seqId, iprot, oprot)
}

type cassandraProcessorLogin struct {
	handler ICassandra
}

func (p *cassandraProcessorLogin) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewLoginArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("login", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewLoginResult()
	if result.Authnx, result.Authzx, err = p.handler.Login(args.AuthRequest); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing login: "+err.Error())
		oprot.WriteMessageBegin("login", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("login", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorSetKeyspace struct {
	handler ICassandra
}

func (p *cassandraProcessorSetKeyspace) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSetKeyspaceArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("set_keyspace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSetKeyspaceResult()
	if result.Ire, err = p.handler.SetKeyspace(args.Keyspace); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing set_keyspace: "+err.Error())
		oprot.WriteMessageBegin("set_keyspace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("set_keyspace", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorGet struct {
	handler ICassandra
}

func (p *cassandraProcessorGet) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewGetArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("get", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewGetResult()
	if result.Success, result.Ire, result.Nfe, result.Ue, result.Te, err = p.handler.Get(args.Key, args.ColumnPath, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get: "+err.Error())
		oprot.WriteMessageBegin("get", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("get", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorGetSlice struct {
	handler ICassandra
}

func (p *cassandraProcessorGetSlice) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewGetSliceArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("get_slice", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewGetSliceResult()
	if result.Success, result.Ire, result.Ue, result.Te, err = p.handler.GetSlice(args.Key, args.ColumnParent, args.Predicate, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get_slice: "+err.Error())
		oprot.WriteMessageBegin("get_slice", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("get_slice", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorGetCount struct {
	handler ICassandra
}

func (p *cassandraProcessorGetCount) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewGetCountArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("get_count", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewGetCountResult()
	if result.Success, result.Ire, result.Ue, result.Te, err = p.handler.GetCount(args.Key, args.ColumnParent, args.Predicate, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get_count: "+err.Error())
		oprot.WriteMessageBegin("get_count", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("get_count", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorMultigetSlice struct {
	handler ICassandra
}

func (p *cassandraProcessorMultigetSlice) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewMultigetSliceArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("multiget_slice", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewMultigetSliceResult()
	if result.Success, result.Ire, result.Ue, result.Te, err = p.handler.MultigetSlice(args.Keys, args.ColumnParent, args.Predicate, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing multiget_slice: "+err.Error())
		oprot.WriteMessageBegin("multiget_slice", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("multiget_slice", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorMultigetCount struct {
	handler ICassandra
}

func (p *cassandraProcessorMultigetCount) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewMultigetCountArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("multiget_count", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewMultigetCountResult()
	if result.Success, result.Ire, result.Ue, result.Te, err = p.handler.MultigetCount(args.Keys, args.ColumnParent, args.Predicate, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing multiget_count: "+err.Error())
		oprot.WriteMessageBegin("multiget_count", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("multiget_count", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorGetRangeSlices struct {
	handler ICassandra
}

func (p *cassandraProcessorGetRangeSlices) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewGetRangeSlicesArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("get_range_slices", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewGetRangeSlicesResult()
	if result.Success, result.Ire, result.Ue, result.Te, err = p.handler.GetRangeSlices(args.ColumnParent, args.Predicate, args.RangeA1, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get_range_slices: "+err.Error())
		oprot.WriteMessageBegin("get_range_slices", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("get_range_slices", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorGetIndexedSlices struct {
	handler ICassandra
}

func (p *cassandraProcessorGetIndexedSlices) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewGetIndexedSlicesArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("get_indexed_slices", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewGetIndexedSlicesResult()
	if result.Success, result.Ire, result.Ue, result.Te, err = p.handler.GetIndexedSlices(args.ColumnParent, args.IndexClause, args.ColumnPredicate, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get_indexed_slices: "+err.Error())
		oprot.WriteMessageBegin("get_indexed_slices", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("get_indexed_slices", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorInsert struct {
	handler ICassandra
}

func (p *cassandraProcessorInsert) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewInsertArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("insert", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewInsertResult()
	if result.Ire, result.Ue, result.Te, err = p.handler.Insert(args.Key, args.ColumnParent, args.Column, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing insert: "+err.Error())
		oprot.WriteMessageBegin("insert", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("insert", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorAdd struct {
	handler ICassandra
}

func (p *cassandraProcessorAdd) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewAddArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewAddResult()
	if result.Ire, result.Ue, result.Te, err = p.handler.Add(args.Key, args.ColumnParent, args.Column, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing add: "+err.Error())
		oprot.WriteMessageBegin("add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("add", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorRemove struct {
	handler ICassandra
}

func (p *cassandraProcessorRemove) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewRemoveArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("remove", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewRemoveResult()
	if result.Ire, result.Ue, result.Te, err = p.handler.Remove(args.Key, args.ColumnPath, args.Timestamp, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing remove: "+err.Error())
		oprot.WriteMessageBegin("remove", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("remove", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorRemoveCounter struct {
	handler ICassandra
}

func (p *cassandraProcessorRemoveCounter) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewRemoveCounterArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("remove_counter", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewRemoveCounterResult()
	if result.Ire, result.Ue, result.Te, err = p.handler.RemoveCounter(args.Key, args.Path, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing remove_counter: "+err.Error())
		oprot.WriteMessageBegin("remove_counter", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("remove_counter", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorBatchMutate struct {
	handler ICassandra
}

func (p *cassandraProcessorBatchMutate) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewBatchMutateArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("batch_mutate", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewBatchMutateResult()
	if result.Ire, result.Ue, result.Te, err = p.handler.BatchMutate(args.MutationMap, args.ConsistencyLevel); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing batch_mutate: "+err.Error())
		oprot.WriteMessageBegin("batch_mutate", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("batch_mutate", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorTruncate struct {
	handler ICassandra
}

func (p *cassandraProcessorTruncate) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewTruncateArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("truncate", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewTruncateResult()
	if result.Ire, result.Ue, err = p.handler.Truncate(args.Cfname); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing truncate: "+err.Error())
		oprot.WriteMessageBegin("truncate", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("truncate", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorDescribeSchemaVersions struct {
	handler ICassandra
}

func (p *cassandraProcessorDescribeSchemaVersions) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewDescribeSchemaVersionsArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("describe_schema_versions", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewDescribeSchemaVersionsResult()
	if result.Success, result.Ire, err = p.handler.DescribeSchemaVersions(); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing describe_schema_versions: "+err.Error())
		oprot.WriteMessageBegin("describe_schema_versions", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("describe_schema_versions", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorDescribeKeyspaces struct {
	handler ICassandra
}

func (p *cassandraProcessorDescribeKeyspaces) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewDescribeKeyspacesArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("describe_keyspaces", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewDescribeKeyspacesResult()
	if result.Success, result.Ire, err = p.handler.DescribeKeyspaces(); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing describe_keyspaces: "+err.Error())
		oprot.WriteMessageBegin("describe_keyspaces", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("describe_keyspaces", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorDescribeClusterName struct {
	handler ICassandra
}

func (p *cassandraProcessorDescribeClusterName) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewDescribeClusterNameArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("describe_cluster_name", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewDescribeClusterNameResult()
	if result.Success, err = p.handler.DescribeClusterName(); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing describe_cluster_name: "+err.Error())
		oprot.WriteMessageBegin("describe_cluster_name", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("describe_cluster_name", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorDescribeVersion struct {
	handler ICassandra
}

func (p *cassandraProcessorDescribeVersion) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewDescribeVersionArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("describe_version", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewDescribeVersionResult()
	if result.Success, err = p.handler.DescribeVersion(); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing describe_version: "+err.Error())
		oprot.WriteMessageBegin("describe_version", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("describe_version", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorDescribeRing struct {
	handler ICassandra
}

func (p *cassandraProcessorDescribeRing) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewDescribeRingArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("describe_ring", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewDescribeRingResult()
	if result.Success, result.Ire, err = p.handler.DescribeRing(args.Keyspace); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing describe_ring: "+err.Error())
		oprot.WriteMessageBegin("describe_ring", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("describe_ring", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorDescribePartitioner struct {
	handler ICassandra
}

func (p *cassandraProcessorDescribePartitioner) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewDescribePartitionerArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("describe_partitioner", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewDescribePartitionerResult()
	if result.Success, err = p.handler.DescribePartitioner(); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing describe_partitioner: "+err.Error())
		oprot.WriteMessageBegin("describe_partitioner", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("describe_partitioner", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorDescribeSnitch struct {
	handler ICassandra
}

func (p *cassandraProcessorDescribeSnitch) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewDescribeSnitchArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("describe_snitch", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewDescribeSnitchResult()
	if result.Success, err = p.handler.DescribeSnitch(); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing describe_snitch: "+err.Error())
		oprot.WriteMessageBegin("describe_snitch", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("describe_snitch", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorDescribeKeyspace struct {
	handler ICassandra
}

func (p *cassandraProcessorDescribeKeyspace) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewDescribeKeyspaceArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("describe_keyspace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewDescribeKeyspaceResult()
	if result.Success, result.Nfe, result.Ire, err = p.handler.DescribeKeyspace(args.Keyspace); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing describe_keyspace: "+err.Error())
		oprot.WriteMessageBegin("describe_keyspace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("describe_keyspace", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorDescribeSplits struct {
	handler ICassandra
}

func (p *cassandraProcessorDescribeSplits) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewDescribeSplitsArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("describe_splits", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewDescribeSplitsResult()
	if result.Success, result.Ire, err = p.handler.DescribeSplits(args.CfName, args.StartToken, args.EndToken, args.KeysPerSplit); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing describe_splits: "+err.Error())
		oprot.WriteMessageBegin("describe_splits", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("describe_splits", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorSystemAddColumnFamily struct {
	handler ICassandra
}

func (p *cassandraProcessorSystemAddColumnFamily) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSystemAddColumnFamilyArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("system_add_column_family", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSystemAddColumnFamilyResult()
	if result.Success, result.Ire, result.Sde, err = p.handler.SystemAddColumnFamily(args.CfDef); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing system_add_column_family: "+err.Error())
		oprot.WriteMessageBegin("system_add_column_family", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("system_add_column_family", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorSystemDropColumnFamily struct {
	handler ICassandra
}

func (p *cassandraProcessorSystemDropColumnFamily) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSystemDropColumnFamilyArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("system_drop_column_family", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSystemDropColumnFamilyResult()
	if result.Success, result.Ire, result.Sde, err = p.handler.SystemDropColumnFamily(args.ColumnFamily); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing system_drop_column_family: "+err.Error())
		oprot.WriteMessageBegin("system_drop_column_family", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("system_drop_column_family", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorSystemAddKeyspace struct {
	handler ICassandra
}

func (p *cassandraProcessorSystemAddKeyspace) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSystemAddKeyspaceArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("system_add_keyspace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSystemAddKeyspaceResult()
	if result.Success, result.Ire, result.Sde, err = p.handler.SystemAddKeyspace(args.KsDef); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing system_add_keyspace: "+err.Error())
		oprot.WriteMessageBegin("system_add_keyspace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("system_add_keyspace", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorSystemDropKeyspace struct {
	handler ICassandra
}

func (p *cassandraProcessorSystemDropKeyspace) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSystemDropKeyspaceArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("system_drop_keyspace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSystemDropKeyspaceResult()
	if result.Success, result.Ire, result.Sde, err = p.handler.SystemDropKeyspace(args.Keyspace); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing system_drop_keyspace: "+err.Error())
		oprot.WriteMessageBegin("system_drop_keyspace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("system_drop_keyspace", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorSystemUpdateKeyspace struct {
	handler ICassandra
}

func (p *cassandraProcessorSystemUpdateKeyspace) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSystemUpdateKeyspaceArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("system_update_keyspace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSystemUpdateKeyspaceResult()
	if result.Success, result.Ire, result.Sde, err = p.handler.SystemUpdateKeyspace(args.KsDef); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing system_update_keyspace: "+err.Error())
		oprot.WriteMessageBegin("system_update_keyspace", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("system_update_keyspace", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorSystemUpdateColumnFamily struct {
	handler ICassandra
}

func (p *cassandraProcessorSystemUpdateColumnFamily) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSystemUpdateColumnFamilyArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("system_update_column_family", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSystemUpdateColumnFamilyResult()
	if result.Success, result.Ire, result.Sde, err = p.handler.SystemUpdateColumnFamily(args.CfDef); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing system_update_column_family: "+err.Error())
		oprot.WriteMessageBegin("system_update_column_family", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("system_update_column_family", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type cassandraProcessorExecuteCqlQuery struct {
	handler ICassandra
}

func (p *cassandraProcessorExecuteCqlQuery) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewExecuteCqlQueryArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("execute_cql_query", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewExecuteCqlQueryResult()
	if result.Success, result.Ire, result.Ue, result.Te, result.Sde, err = p.handler.ExecuteCqlQuery(args.Query, args.Compression); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing execute_cql_query: "+err.Error())
		oprot.WriteMessageBegin("execute_cql_query", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Transport().Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("execute_cql_query", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Transport().Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

/**
 * Attributes:
 *  - AuthRequest
 */
type LoginArgs struct {
	thrift.TStruct
	AuthRequest *AuthenticationRequest "auth_request" // 1
}

func NewLoginArgs() *LoginArgs {
	output := &LoginArgs{
		TStruct: thrift.NewTStruct("login_args", []thrift.TField{
			thrift.NewTField("auth_request", thrift.STRUCT, 1),
		}),
	}
	{
	}
	return output
}

func (p *LoginArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "auth_request" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *LoginArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.AuthRequest = NewAuthenticationRequest()
	err636 := p.AuthRequest.Read(iprot)
	if err636 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.AuthRequestAuthenticationRequest", err636)
	}
	return err
}

func (p *LoginArgs) ReadFieldAuthRequest(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *LoginArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("login_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *LoginArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.AuthRequest != nil {
		err = oprot.WriteFieldBegin("auth_request", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "auth_request", p.ThriftName(), err)
		}
		err = p.AuthRequest.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("AuthenticationRequest", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "auth_request", p.ThriftName(), err)
		}
	}
	return err
}

func (p *LoginArgs) WriteFieldAuthRequest(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *LoginArgs) TStructName() string {
	return "LoginArgs"
}

func (p *LoginArgs) ThriftName() string {
	return "login_args"
}

func (p *LoginArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LoginArgs(%+v)", *p)
}

func (p *LoginArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*LoginArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *LoginArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.AuthRequest
	}
	return nil
}

func (p *LoginArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("auth_request", thrift.STRUCT, 1),
	})
}

/**
 * Attributes:
 *  - Authnx
 *  - Authzx
 */
type LoginResult struct {
	thrift.TStruct
	Authnx *AuthenticationException "authnx" // 1
	Authzx *AuthorizationException  "authzx" // 2
}

func NewLoginResult() *LoginResult {
	output := &LoginResult{
		TStruct: thrift.NewTStruct("login_result", []thrift.TField{
			thrift.NewTField("authnx", thrift.STRUCT, 1),
			thrift.NewTField("authzx", thrift.STRUCT, 2),
		}),
	}
	{
	}
	return output
}

func (p *LoginResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "authnx" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "authzx" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *LoginResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Authnx = NewAuthenticationException()
	err639 := p.Authnx.Read(iprot)
	if err639 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.AuthnxAuthenticationException", err639)
	}
	return err
}

func (p *LoginResult) ReadFieldAuthnx(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *LoginResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Authzx = NewAuthorizationException()
	err642 := p.Authzx.Read(iprot)
	if err642 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.AuthzxAuthorizationException", err642)
	}
	return err
}

func (p *LoginResult) ReadFieldAuthzx(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *LoginResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("login_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Authzx != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Authnx != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *LoginResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Authnx != nil {
		err = oprot.WriteFieldBegin("authnx", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "authnx", p.ThriftName(), err)
		}
		err = p.Authnx.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("AuthenticationException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "authnx", p.ThriftName(), err)
		}
	}
	return err
}

func (p *LoginResult) WriteFieldAuthnx(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *LoginResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Authzx != nil {
		err = oprot.WriteFieldBegin("authzx", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "authzx", p.ThriftName(), err)
		}
		err = p.Authzx.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("AuthorizationException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "authzx", p.ThriftName(), err)
		}
	}
	return err
}

func (p *LoginResult) WriteFieldAuthzx(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *LoginResult) TStructName() string {
	return "LoginResult"
}

func (p *LoginResult) ThriftName() string {
	return "login_result"
}

func (p *LoginResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LoginResult(%+v)", *p)
}

func (p *LoginResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*LoginResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *LoginResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Authnx
	case 2:
		return p.Authzx
	}
	return nil
}

func (p *LoginResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("authnx", thrift.STRUCT, 1),
		thrift.NewTField("authzx", thrift.STRUCT, 2),
	})
}

/**
 * Attributes:
 *  - Keyspace
 */
type SetKeyspaceArgs struct {
	thrift.TStruct
	Keyspace string "keyspace" // 1
}

func NewSetKeyspaceArgs() *SetKeyspaceArgs {
	output := &SetKeyspaceArgs{
		TStruct: thrift.NewTStruct("set_keyspace_args", []thrift.TField{
			thrift.NewTField("keyspace", thrift.STRING, 1),
		}),
	}
	{
	}
	return output
}

func (p *SetKeyspaceArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "keyspace" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SetKeyspaceArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v643, err644 := iprot.ReadString()
	if err644 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "keyspace", p.ThriftName(), err644)
	}
	p.Keyspace = v643
	return err
}

func (p *SetKeyspaceArgs) ReadFieldKeyspace(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SetKeyspaceArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("set_keyspace_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SetKeyspaceArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("keyspace", thrift.STRING, 1)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Keyspace))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	return err
}

func (p *SetKeyspaceArgs) WriteFieldKeyspace(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SetKeyspaceArgs) TStructName() string {
	return "SetKeyspaceArgs"
}

func (p *SetKeyspaceArgs) ThriftName() string {
	return "set_keyspace_args"
}

func (p *SetKeyspaceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SetKeyspaceArgs(%+v)", *p)
}

func (p *SetKeyspaceArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SetKeyspaceArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SetKeyspaceArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Keyspace
	}
	return nil
}

func (p *SetKeyspaceArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("keyspace", thrift.STRING, 1),
	})
}

/**
 * Attributes:
 *  - Ire
 */
type SetKeyspaceResult struct {
	thrift.TStruct
	Ire *InvalidRequestException "ire" // 1
}

func NewSetKeyspaceResult() *SetKeyspaceResult {
	output := &SetKeyspaceResult{
		TStruct: thrift.NewTStruct("set_keyspace_result", []thrift.TField{
			thrift.NewTField("ire", thrift.STRUCT, 1),
		}),
	}
	{
	}
	return output
}

func (p *SetKeyspaceResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SetKeyspaceResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err647 := p.Ire.Read(iprot)
	if err647 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err647)
	}
	return err
}

func (p *SetKeyspaceResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SetKeyspaceResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("set_keyspace_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SetKeyspaceResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SetKeyspaceResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SetKeyspaceResult) TStructName() string {
	return "SetKeyspaceResult"
}

func (p *SetKeyspaceResult) ThriftName() string {
	return "set_keyspace_result"
}

func (p *SetKeyspaceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SetKeyspaceResult(%+v)", *p)
}

func (p *SetKeyspaceResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SetKeyspaceResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SetKeyspaceResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Ire
	}
	return nil
}

func (p *SetKeyspaceResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("ire", thrift.STRUCT, 1),
	})
}

/**
 * Attributes:
 *  - Key
 *  - ColumnPath
 *  - ConsistencyLevel
 */
type GetArgs struct {
	thrift.TStruct
	Key              []byte           "key"               // 1
	ColumnPath       *ColumnPath      "column_path"       // 2
	ConsistencyLevel ConsistencyLevel "consistency_level" // 3
}

func NewGetArgs() *GetArgs {
	output := &GetArgs{
		TStruct: thrift.NewTStruct("get_args", []thrift.TField{
			thrift.NewTField("key", thrift.BINARY, 1),
			thrift.NewTField("column_path", thrift.STRUCT, 2),
			thrift.NewTField("consistency_level", thrift.I32, 3),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *GetArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *GetArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "key" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "column_path" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v648, err649 := iprot.ReadBinary()
	if err649 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "key", p.ThriftName(), err649)
	}
	p.Key = v648
	return err
}

func (p *GetArgs) ReadFieldKey(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *GetArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnPath = NewColumnPath()
	err652 := p.ColumnPath.Read(iprot)
	if err652 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnPathColumnPath", err652)
	}
	return err
}

func (p *GetArgs) ReadFieldColumnPath(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *GetArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v653, err654 := iprot.ReadI32()
	if err654 != nil {
		return thrift.NewTProtocolExceptionReadField(3, "consistency_level", p.ThriftName(), err654)
	}
	p.ConsistencyLevel = ConsistencyLevel(v653)
	return err
}

func (p *GetArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *GetArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("get_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Key != nil {
		err = oprot.WriteFieldBegin("key", thrift.BINARY, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteBinary(p.Key)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetArgs) WriteFieldKey(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *GetArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnPath != nil {
		err = oprot.WriteFieldBegin("column_path", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_path", p.ThriftName(), err)
		}
		err = p.ColumnPath.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnPath", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_path", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetArgs) WriteFieldColumnPath(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *GetArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *GetArgs) TStructName() string {
	return "GetArgs"
}

func (p *GetArgs) ThriftName() string {
	return "get_args"
}

func (p *GetArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetArgs(%+v)", *p)
}

func (p *GetArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*GetArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *GetArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Key
	case 2:
		return p.ColumnPath
	case 3:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *GetArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("key", thrift.BINARY, 1),
		thrift.NewTField("column_path", thrift.STRUCT, 2),
		thrift.NewTField("consistency_level", thrift.I32, 3),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Nfe
 *  - Ue
 *  - Te
 */
type GetResult struct {
	thrift.TStruct
	Success *ColumnOrSuperColumn     "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
	Nfe     *NotFoundException       "nfe"     // 2
	Ue      *UnavailableException    "ue"      // 3
	Te      *TimedOutException       "te"      // 4
}

func NewGetResult() *GetResult {
	output := &GetResult{
		TStruct: thrift.NewTStruct("get_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRUCT, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("nfe", thrift.STRUCT, 2),
			thrift.NewTField("ue", thrift.STRUCT, 3),
			thrift.NewTField("te", thrift.STRUCT, 4),
		}),
	}
	{
	}
	return output
}

func (p *GetResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "nfe" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Success = NewColumnOrSuperColumn()
	err657 := p.Success.Read(iprot)
	if err657 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.SuccessColumnOrSuperColumn", err657)
	}
	return err
}

func (p *GetResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *GetResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err660 := p.Ire.Read(iprot)
	if err660 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err660)
	}
	return err
}

func (p *GetResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *GetResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Nfe = NewNotFoundException()
	err663 := p.Nfe.Read(iprot)
	if err663 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.NfeNotFoundException", err663)
	}
	return err
}

func (p *GetResult) ReadFieldNfe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *GetResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err666 := p.Ue.Read(iprot)
	if err666 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err666)
	}
	return err
}

func (p *GetResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *GetResult) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err669 := p.Te.Read(iprot)
	if err669 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err669)
	}
	return err
}

func (p *GetResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *GetResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("get_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField4(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Nfe != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = p.Success.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnOrSuperColumn", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *GetResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *GetResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Nfe != nil {
		err = oprot.WriteFieldBegin("nfe", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "nfe", p.ThriftName(), err)
		}
		err = p.Nfe.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("NotFoundException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "nfe", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetResult) WriteFieldNfe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *GetResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *GetResult) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *GetResult) TStructName() string {
	return "GetResult"
}

func (p *GetResult) ThriftName() string {
	return "get_result"
}

func (p *GetResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetResult(%+v)", *p)
}

func (p *GetResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*GetResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *GetResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Nfe
	case 3:
		return p.Ue
	case 4:
		return p.Te
	}
	return nil
}

func (p *GetResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRUCT, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("nfe", thrift.STRUCT, 2),
		thrift.NewTField("ue", thrift.STRUCT, 3),
		thrift.NewTField("te", thrift.STRUCT, 4),
	})
}

/**
 * Attributes:
 *  - Key
 *  - ColumnParent
 *  - Predicate
 *  - ConsistencyLevel
 */
type GetSliceArgs struct {
	thrift.TStruct
	Key              []byte           "key"               // 1
	ColumnParent     *ColumnParent    "column_parent"     // 2
	Predicate        *SlicePredicate  "predicate"         // 3
	ConsistencyLevel ConsistencyLevel "consistency_level" // 4
}

func NewGetSliceArgs() *GetSliceArgs {
	output := &GetSliceArgs{
		TStruct: thrift.NewTStruct("get_slice_args", []thrift.TField{
			thrift.NewTField("key", thrift.BINARY, 1),
			thrift.NewTField("column_parent", thrift.STRUCT, 2),
			thrift.NewTField("predicate", thrift.STRUCT, 3),
			thrift.NewTField("consistency_level", thrift.I32, 4),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *GetSliceArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *GetSliceArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "key" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "column_parent" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "predicate" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetSliceArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v670, err671 := iprot.ReadBinary()
	if err671 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "key", p.ThriftName(), err671)
	}
	p.Key = v670
	return err
}

func (p *GetSliceArgs) ReadFieldKey(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *GetSliceArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnParent = NewColumnParent()
	err674 := p.ColumnParent.Read(iprot)
	if err674 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnParentColumnParent", err674)
	}
	return err
}

func (p *GetSliceArgs) ReadFieldColumnParent(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *GetSliceArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Predicate = NewSlicePredicate()
	err677 := p.Predicate.Read(iprot)
	if err677 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.PredicateSlicePredicate", err677)
	}
	return err
}

func (p *GetSliceArgs) ReadFieldPredicate(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *GetSliceArgs) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v678, err679 := iprot.ReadI32()
	if err679 != nil {
		return thrift.NewTProtocolExceptionReadField(4, "consistency_level", p.ThriftName(), err679)
	}
	p.ConsistencyLevel = ConsistencyLevel(v678)
	return err
}

func (p *GetSliceArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *GetSliceArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("get_slice_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField4(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetSliceArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Key != nil {
		err = oprot.WriteFieldBegin("key", thrift.BINARY, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteBinary(p.Key)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetSliceArgs) WriteFieldKey(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *GetSliceArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnParent != nil {
		err = oprot.WriteFieldBegin("column_parent", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
		err = p.ColumnParent.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnParent", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetSliceArgs) WriteFieldColumnParent(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *GetSliceArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Predicate != nil {
		err = oprot.WriteFieldBegin("predicate", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "predicate", p.ThriftName(), err)
		}
		err = p.Predicate.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SlicePredicate", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "predicate", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetSliceArgs) WriteFieldPredicate(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *GetSliceArgs) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetSliceArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *GetSliceArgs) TStructName() string {
	return "GetSliceArgs"
}

func (p *GetSliceArgs) ThriftName() string {
	return "get_slice_args"
}

func (p *GetSliceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetSliceArgs(%+v)", *p)
}

func (p *GetSliceArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*GetSliceArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *GetSliceArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Key
	case 2:
		return p.ColumnParent
	case 3:
		return p.Predicate
	case 4:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *GetSliceArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("key", thrift.BINARY, 1),
		thrift.NewTField("column_parent", thrift.STRUCT, 2),
		thrift.NewTField("predicate", thrift.STRUCT, 3),
		thrift.NewTField("consistency_level", thrift.I32, 4),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Ue
 *  - Te
 */
type GetSliceResult struct {
	thrift.TStruct
	Success thrift.TList             "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
	Ue      *UnavailableException    "ue"      // 2
	Te      *TimedOutException       "te"      // 3
}

func NewGetSliceResult() *GetSliceResult {
	output := &GetSliceResult{
		TStruct: thrift.NewTStruct("get_slice_result", []thrift.TField{
			thrift.NewTField("success", thrift.LIST, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *GetSliceResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.LIST {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetSliceResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_etype685, _size682, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Success", "", err)
	}
	p.Success = thrift.NewTList(_etype685, _size682)
	for _i686 := 0; _i686 < _size682; _i686++ {
		_elem687 := NewColumnOrSuperColumn()
		err690 := _elem687.Read(iprot)
		if err690 != nil {
			return thrift.NewTProtocolExceptionReadStruct("_elem687ColumnOrSuperColumn", err690)
		}
		p.Success.Push(_elem687)
	}
	err = iprot.ReadListEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
	}
	return err
}

func (p *GetSliceResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *GetSliceResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err693 := p.Ire.Read(iprot)
	if err693 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err693)
	}
	return err
}

func (p *GetSliceResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *GetSliceResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err696 := p.Ue.Read(iprot)
	if err696 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err696)
	}
	return err
}

func (p *GetSliceResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *GetSliceResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err699 := p.Te.Read(iprot)
	if err699 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err699)
	}
	return err
}

func (p *GetSliceResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *GetSliceResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("get_slice_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetSliceResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.LIST, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = oprot.WriteListBegin(thrift.STRUCT, p.Success.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		for Iter700 := range p.Success.Iter() {
			Iter701 := Iter700.(*ColumnOrSuperColumn)
			err = Iter701.Write(oprot)
			if err != nil {
				return thrift.NewTProtocolExceptionWriteStruct("ColumnOrSuperColumn", err)
			}
		}
		err = oprot.WriteListEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetSliceResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *GetSliceResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetSliceResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *GetSliceResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetSliceResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *GetSliceResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetSliceResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *GetSliceResult) TStructName() string {
	return "GetSliceResult"
}

func (p *GetSliceResult) ThriftName() string {
	return "get_slice_result"
}

func (p *GetSliceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetSliceResult(%+v)", *p)
}

func (p *GetSliceResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*GetSliceResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *GetSliceResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *GetSliceResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.LIST, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - Key
 *  - ColumnParent
 *  - Predicate
 *  - ConsistencyLevel
 */
type GetCountArgs struct {
	thrift.TStruct
	Key              []byte           "key"               // 1
	ColumnParent     *ColumnParent    "column_parent"     // 2
	Predicate        *SlicePredicate  "predicate"         // 3
	ConsistencyLevel ConsistencyLevel "consistency_level" // 4
}

func NewGetCountArgs() *GetCountArgs {
	output := &GetCountArgs{
		TStruct: thrift.NewTStruct("get_count_args", []thrift.TField{
			thrift.NewTField("key", thrift.BINARY, 1),
			thrift.NewTField("column_parent", thrift.STRUCT, 2),
			thrift.NewTField("predicate", thrift.STRUCT, 3),
			thrift.NewTField("consistency_level", thrift.I32, 4),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *GetCountArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *GetCountArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "key" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "column_parent" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "predicate" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetCountArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v702, err703 := iprot.ReadBinary()
	if err703 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "key", p.ThriftName(), err703)
	}
	p.Key = v702
	return err
}

func (p *GetCountArgs) ReadFieldKey(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *GetCountArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnParent = NewColumnParent()
	err706 := p.ColumnParent.Read(iprot)
	if err706 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnParentColumnParent", err706)
	}
	return err
}

func (p *GetCountArgs) ReadFieldColumnParent(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *GetCountArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Predicate = NewSlicePredicate()
	err709 := p.Predicate.Read(iprot)
	if err709 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.PredicateSlicePredicate", err709)
	}
	return err
}

func (p *GetCountArgs) ReadFieldPredicate(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *GetCountArgs) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v710, err711 := iprot.ReadI32()
	if err711 != nil {
		return thrift.NewTProtocolExceptionReadField(4, "consistency_level", p.ThriftName(), err711)
	}
	p.ConsistencyLevel = ConsistencyLevel(v710)
	return err
}

func (p *GetCountArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *GetCountArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("get_count_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField4(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetCountArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Key != nil {
		err = oprot.WriteFieldBegin("key", thrift.BINARY, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteBinary(p.Key)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetCountArgs) WriteFieldKey(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *GetCountArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnParent != nil {
		err = oprot.WriteFieldBegin("column_parent", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
		err = p.ColumnParent.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnParent", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetCountArgs) WriteFieldColumnParent(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *GetCountArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Predicate != nil {
		err = oprot.WriteFieldBegin("predicate", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "predicate", p.ThriftName(), err)
		}
		err = p.Predicate.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SlicePredicate", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "predicate", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetCountArgs) WriteFieldPredicate(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *GetCountArgs) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetCountArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *GetCountArgs) TStructName() string {
	return "GetCountArgs"
}

func (p *GetCountArgs) ThriftName() string {
	return "get_count_args"
}

func (p *GetCountArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetCountArgs(%+v)", *p)
}

func (p *GetCountArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*GetCountArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *GetCountArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Key
	case 2:
		return p.ColumnParent
	case 3:
		return p.Predicate
	case 4:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *GetCountArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("key", thrift.BINARY, 1),
		thrift.NewTField("column_parent", thrift.STRUCT, 2),
		thrift.NewTField("predicate", thrift.STRUCT, 3),
		thrift.NewTField("consistency_level", thrift.I32, 4),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Ue
 *  - Te
 */
type GetCountResult struct {
	thrift.TStruct
	Success int32                    "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
	Ue      *UnavailableException    "ue"      // 2
	Te      *TimedOutException       "te"      // 3
}

func NewGetCountResult() *GetCountResult {
	output := &GetCountResult{
		TStruct: thrift.NewTStruct("get_count_result", []thrift.TField{
			thrift.NewTField("success", thrift.I32, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *GetCountResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetCountResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v712, err713 := iprot.ReadI32()
	if err713 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err713)
	}
	p.Success = v712
	return err
}

func (p *GetCountResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *GetCountResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err716 := p.Ire.Read(iprot)
	if err716 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err716)
	}
	return err
}

func (p *GetCountResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *GetCountResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err719 := p.Ue.Read(iprot)
	if err719 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err719)
	}
	return err
}

func (p *GetCountResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *GetCountResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err722 := p.Te.Read(iprot)
	if err722 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err722)
	}
	return err
}

func (p *GetCountResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *GetCountResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("get_count_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetCountResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.I32, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteI32(int32(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *GetCountResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *GetCountResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetCountResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *GetCountResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetCountResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *GetCountResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetCountResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *GetCountResult) TStructName() string {
	return "GetCountResult"
}

func (p *GetCountResult) ThriftName() string {
	return "get_count_result"
}

func (p *GetCountResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetCountResult(%+v)", *p)
}

func (p *GetCountResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*GetCountResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *GetCountResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *GetCountResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.I32, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - Keys
 *  - ColumnParent
 *  - Predicate
 *  - ConsistencyLevel
 */
type MultigetSliceArgs struct {
	thrift.TStruct
	Keys             thrift.TList     "keys"              // 1
	ColumnParent     *ColumnParent    "column_parent"     // 2
	Predicate        *SlicePredicate  "predicate"         // 3
	ConsistencyLevel ConsistencyLevel "consistency_level" // 4
}

func NewMultigetSliceArgs() *MultigetSliceArgs {
	output := &MultigetSliceArgs{
		TStruct: thrift.NewTStruct("multiget_slice_args", []thrift.TField{
			thrift.NewTField("keys", thrift.LIST, 1),
			thrift.NewTField("column_parent", thrift.STRUCT, 2),
			thrift.NewTField("predicate", thrift.STRUCT, 3),
			thrift.NewTField("consistency_level", thrift.I32, 4),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *MultigetSliceArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *MultigetSliceArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "keys" {
			if fieldTypeId == thrift.LIST {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "column_parent" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "predicate" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *MultigetSliceArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_etype728, _size725, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Keys", "", err)
	}
	p.Keys = thrift.NewTList(_etype728, _size725)
	for _i729 := 0; _i729 < _size725; _i729++ {
		v731, err732 := iprot.ReadBinary()
		if err732 != nil {
			return thrift.NewTProtocolExceptionReadField(0, "_elem730", "", err732)
		}
		_elem730 := v731
		p.Keys.Push(_elem730)
	}
	err = iprot.ReadListEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
	}
	return err
}

func (p *MultigetSliceArgs) ReadFieldKeys(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *MultigetSliceArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnParent = NewColumnParent()
	err735 := p.ColumnParent.Read(iprot)
	if err735 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnParentColumnParent", err735)
	}
	return err
}

func (p *MultigetSliceArgs) ReadFieldColumnParent(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *MultigetSliceArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Predicate = NewSlicePredicate()
	err738 := p.Predicate.Read(iprot)
	if err738 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.PredicateSlicePredicate", err738)
	}
	return err
}

func (p *MultigetSliceArgs) ReadFieldPredicate(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *MultigetSliceArgs) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v739, err740 := iprot.ReadI32()
	if err740 != nil {
		return thrift.NewTProtocolExceptionReadField(4, "consistency_level", p.ThriftName(), err740)
	}
	p.ConsistencyLevel = ConsistencyLevel(v739)
	return err
}

func (p *MultigetSliceArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *MultigetSliceArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("multiget_slice_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField4(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *MultigetSliceArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Keys != nil {
		err = oprot.WriteFieldBegin("keys", thrift.LIST, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "keys", p.ThriftName(), err)
		}
		err = oprot.WriteListBegin(thrift.BINARY, p.Keys.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		for Iter741 := range p.Keys.Iter() {
			Iter742 := Iter741.([]byte)
			err = oprot.WriteBinary(Iter742)
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(0, "Iter742", "", err)
			}
		}
		err = oprot.WriteListEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "keys", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetSliceArgs) WriteFieldKeys(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *MultigetSliceArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnParent != nil {
		err = oprot.WriteFieldBegin("column_parent", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
		err = p.ColumnParent.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnParent", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetSliceArgs) WriteFieldColumnParent(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *MultigetSliceArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Predicate != nil {
		err = oprot.WriteFieldBegin("predicate", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "predicate", p.ThriftName(), err)
		}
		err = p.Predicate.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SlicePredicate", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "predicate", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetSliceArgs) WriteFieldPredicate(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *MultigetSliceArgs) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetSliceArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *MultigetSliceArgs) TStructName() string {
	return "MultigetSliceArgs"
}

func (p *MultigetSliceArgs) ThriftName() string {
	return "multiget_slice_args"
}

func (p *MultigetSliceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MultigetSliceArgs(%+v)", *p)
}

func (p *MultigetSliceArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*MultigetSliceArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *MultigetSliceArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Keys
	case 2:
		return p.ColumnParent
	case 3:
		return p.Predicate
	case 4:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *MultigetSliceArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("keys", thrift.LIST, 1),
		thrift.NewTField("column_parent", thrift.STRUCT, 2),
		thrift.NewTField("predicate", thrift.STRUCT, 3),
		thrift.NewTField("consistency_level", thrift.I32, 4),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Ue
 *  - Te
 */
type MultigetSliceResult struct {
	thrift.TStruct
	Success thrift.TMap              "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
	Ue      *UnavailableException    "ue"      // 2
	Te      *TimedOutException       "te"      // 3
}

func NewMultigetSliceResult() *MultigetSliceResult {
	output := &MultigetSliceResult{
		TStruct: thrift.NewTStruct("multiget_slice_result", []thrift.TField{
			thrift.NewTField("success", thrift.MAP, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *MultigetSliceResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.MAP {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *MultigetSliceResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_ktype746, _vtype747, _size745, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Success", "", err)
	}
	p.Success = thrift.NewTMap(_ktype746, _vtype747, _size745)
	for _i749 := 0; _i749 < _size745; _i749++ {
		v752, err753 := iprot.ReadBinary()
		if err753 != nil {
			return thrift.NewTProtocolExceptionReadField(0, "_key750", "", err753)
		}
		_key750 := v752
		_etype759, _size756, err := iprot.ReadListBegin()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(-1, "_val751", "", err)
		}
		_val751 := thrift.NewTList(_etype759, _size756)
		for _i760 := 0; _i760 < _size756; _i760++ {
			_elem761 := NewColumnOrSuperColumn()
			err764 := _elem761.Read(iprot)
			if err764 != nil {
				return thrift.NewTProtocolExceptionReadStruct("_elem761ColumnOrSuperColumn", err764)
			}
			_val751.Push(_elem761)
		}
		err = iprot.ReadListEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
		}
		p.Success.Set(_key750, _val751)
	}
	err = iprot.ReadMapEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "map", err)
	}
	return err
}

func (p *MultigetSliceResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *MultigetSliceResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err767 := p.Ire.Read(iprot)
	if err767 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err767)
	}
	return err
}

func (p *MultigetSliceResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *MultigetSliceResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err770 := p.Ue.Read(iprot)
	if err770 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err770)
	}
	return err
}

func (p *MultigetSliceResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *MultigetSliceResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err773 := p.Te.Read(iprot)
	if err773 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err773)
	}
	return err
}

func (p *MultigetSliceResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *MultigetSliceResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("multiget_slice_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *MultigetSliceResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.MAP, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = oprot.WriteMapBegin(thrift.BINARY, thrift.LIST, p.Success.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "map", err)
		}
		for Miter774 := range p.Success.Iter() {
			Kiter775, Viter776 := Miter774.Key().([]byte), Miter774.Value().(thrift.TList)
			err = oprot.WriteBinary(Kiter775)
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(0, "Kiter775", "", err)
			}
			err = oprot.WriteListBegin(thrift.STRUCT, Viter776.Len())
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
			}
			for Iter777 := range Viter776.Iter() {
				Iter778 := Iter777.(*ColumnOrSuperColumn)
				err = Iter778.Write(oprot)
				if err != nil {
					return thrift.NewTProtocolExceptionWriteStruct("ColumnOrSuperColumn", err)
				}
			}
			err = oprot.WriteListEnd()
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
			}
		}
		err = oprot.WriteMapEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "map", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetSliceResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *MultigetSliceResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetSliceResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *MultigetSliceResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetSliceResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *MultigetSliceResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetSliceResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *MultigetSliceResult) TStructName() string {
	return "MultigetSliceResult"
}

func (p *MultigetSliceResult) ThriftName() string {
	return "multiget_slice_result"
}

func (p *MultigetSliceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MultigetSliceResult(%+v)", *p)
}

func (p *MultigetSliceResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*MultigetSliceResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *MultigetSliceResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *MultigetSliceResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.MAP, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - Keys
 *  - ColumnParent
 *  - Predicate
 *  - ConsistencyLevel
 */
type MultigetCountArgs struct {
	thrift.TStruct
	Keys             thrift.TList     "keys"              // 1
	ColumnParent     *ColumnParent    "column_parent"     // 2
	Predicate        *SlicePredicate  "predicate"         // 3
	ConsistencyLevel ConsistencyLevel "consistency_level" // 4
}

func NewMultigetCountArgs() *MultigetCountArgs {
	output := &MultigetCountArgs{
		TStruct: thrift.NewTStruct("multiget_count_args", []thrift.TField{
			thrift.NewTField("keys", thrift.LIST, 1),
			thrift.NewTField("column_parent", thrift.STRUCT, 2),
			thrift.NewTField("predicate", thrift.STRUCT, 3),
			thrift.NewTField("consistency_level", thrift.I32, 4),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *MultigetCountArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *MultigetCountArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "keys" {
			if fieldTypeId == thrift.LIST {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "column_parent" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "predicate" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *MultigetCountArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_etype784, _size781, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Keys", "", err)
	}
	p.Keys = thrift.NewTList(_etype784, _size781)
	for _i785 := 0; _i785 < _size781; _i785++ {
		v787, err788 := iprot.ReadBinary()
		if err788 != nil {
			return thrift.NewTProtocolExceptionReadField(0, "_elem786", "", err788)
		}
		_elem786 := v787
		p.Keys.Push(_elem786)
	}
	err = iprot.ReadListEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
	}
	return err
}

func (p *MultigetCountArgs) ReadFieldKeys(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *MultigetCountArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnParent = NewColumnParent()
	err791 := p.ColumnParent.Read(iprot)
	if err791 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnParentColumnParent", err791)
	}
	return err
}

func (p *MultigetCountArgs) ReadFieldColumnParent(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *MultigetCountArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Predicate = NewSlicePredicate()
	err794 := p.Predicate.Read(iprot)
	if err794 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.PredicateSlicePredicate", err794)
	}
	return err
}

func (p *MultigetCountArgs) ReadFieldPredicate(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *MultigetCountArgs) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v795, err796 := iprot.ReadI32()
	if err796 != nil {
		return thrift.NewTProtocolExceptionReadField(4, "consistency_level", p.ThriftName(), err796)
	}
	p.ConsistencyLevel = ConsistencyLevel(v795)
	return err
}

func (p *MultigetCountArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *MultigetCountArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("multiget_count_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField4(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *MultigetCountArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Keys != nil {
		err = oprot.WriteFieldBegin("keys", thrift.LIST, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "keys", p.ThriftName(), err)
		}
		err = oprot.WriteListBegin(thrift.BINARY, p.Keys.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		for Iter797 := range p.Keys.Iter() {
			Iter798 := Iter797.([]byte)
			err = oprot.WriteBinary(Iter798)
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(0, "Iter798", "", err)
			}
		}
		err = oprot.WriteListEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "keys", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetCountArgs) WriteFieldKeys(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *MultigetCountArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnParent != nil {
		err = oprot.WriteFieldBegin("column_parent", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
		err = p.ColumnParent.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnParent", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetCountArgs) WriteFieldColumnParent(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *MultigetCountArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Predicate != nil {
		err = oprot.WriteFieldBegin("predicate", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "predicate", p.ThriftName(), err)
		}
		err = p.Predicate.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SlicePredicate", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "predicate", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetCountArgs) WriteFieldPredicate(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *MultigetCountArgs) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetCountArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *MultigetCountArgs) TStructName() string {
	return "MultigetCountArgs"
}

func (p *MultigetCountArgs) ThriftName() string {
	return "multiget_count_args"
}

func (p *MultigetCountArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MultigetCountArgs(%+v)", *p)
}

func (p *MultigetCountArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*MultigetCountArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *MultigetCountArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Keys
	case 2:
		return p.ColumnParent
	case 3:
		return p.Predicate
	case 4:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *MultigetCountArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("keys", thrift.LIST, 1),
		thrift.NewTField("column_parent", thrift.STRUCT, 2),
		thrift.NewTField("predicate", thrift.STRUCT, 3),
		thrift.NewTField("consistency_level", thrift.I32, 4),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Ue
 *  - Te
 */
type MultigetCountResult struct {
	thrift.TStruct
	Success thrift.TMap              "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
	Ue      *UnavailableException    "ue"      // 2
	Te      *TimedOutException       "te"      // 3
}

func NewMultigetCountResult() *MultigetCountResult {
	output := &MultigetCountResult{
		TStruct: thrift.NewTStruct("multiget_count_result", []thrift.TField{
			thrift.NewTField("success", thrift.MAP, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *MultigetCountResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.MAP {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *MultigetCountResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_ktype802, _vtype803, _size801, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Success", "", err)
	}
	p.Success = thrift.NewTMap(_ktype802, _vtype803, _size801)
	for _i805 := 0; _i805 < _size801; _i805++ {
		v808, err809 := iprot.ReadBinary()
		if err809 != nil {
			return thrift.NewTProtocolExceptionReadField(0, "_key806", "", err809)
		}
		_key806 := v808
		v810, err811 := iprot.ReadI32()
		if err811 != nil {
			return thrift.NewTProtocolExceptionReadField(0, "_val807", "", err811)
		}
		_val807 := v810
		p.Success.Set(_key806, _val807)
	}
	err = iprot.ReadMapEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "map", err)
	}
	return err
}

func (p *MultigetCountResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *MultigetCountResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err814 := p.Ire.Read(iprot)
	if err814 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err814)
	}
	return err
}

func (p *MultigetCountResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *MultigetCountResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err817 := p.Ue.Read(iprot)
	if err817 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err817)
	}
	return err
}

func (p *MultigetCountResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *MultigetCountResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err820 := p.Te.Read(iprot)
	if err820 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err820)
	}
	return err
}

func (p *MultigetCountResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *MultigetCountResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("multiget_count_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *MultigetCountResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.MAP, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = oprot.WriteMapBegin(thrift.BINARY, thrift.I32, p.Success.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "map", err)
		}
		for Miter821 := range p.Success.Iter() {
			Kiter822, Viter823 := Miter821.Key().([]byte), Miter821.Value().(int32)
			err = oprot.WriteBinary(Kiter822)
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(0, "Kiter822", "", err)
			}
			err = oprot.WriteI32(int32(Viter823))
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(0, "Viter823", "", err)
			}
		}
		err = oprot.WriteMapEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "map", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetCountResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *MultigetCountResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetCountResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *MultigetCountResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetCountResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *MultigetCountResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *MultigetCountResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *MultigetCountResult) TStructName() string {
	return "MultigetCountResult"
}

func (p *MultigetCountResult) ThriftName() string {
	return "multiget_count_result"
}

func (p *MultigetCountResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MultigetCountResult(%+v)", *p)
}

func (p *MultigetCountResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*MultigetCountResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *MultigetCountResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *MultigetCountResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.MAP, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - ColumnParent
 *  - Predicate
 *  - RangeA1
 *  - ConsistencyLevel
 */
type GetRangeSlicesArgs struct {
	thrift.TStruct
	ColumnParent     *ColumnParent    "column_parent"     // 1
	Predicate        *SlicePredicate  "predicate"         // 2
	RangeA1          *KeyRange        "range"             // 3
	ConsistencyLevel ConsistencyLevel "consistency_level" // 4
}

func NewGetRangeSlicesArgs() *GetRangeSlicesArgs {
	output := &GetRangeSlicesArgs{
		TStruct: thrift.NewTStruct("get_range_slices_args", []thrift.TField{
			thrift.NewTField("column_parent", thrift.STRUCT, 1),
			thrift.NewTField("predicate", thrift.STRUCT, 2),
			thrift.NewTField("range", thrift.STRUCT, 3),
			thrift.NewTField("consistency_level", thrift.I32, 4),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *GetRangeSlicesArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *GetRangeSlicesArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "column_parent" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "predicate" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "range" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetRangeSlicesArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnParent = NewColumnParent()
	err826 := p.ColumnParent.Read(iprot)
	if err826 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnParentColumnParent", err826)
	}
	return err
}

func (p *GetRangeSlicesArgs) ReadFieldColumnParent(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *GetRangeSlicesArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Predicate = NewSlicePredicate()
	err829 := p.Predicate.Read(iprot)
	if err829 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.PredicateSlicePredicate", err829)
	}
	return err
}

func (p *GetRangeSlicesArgs) ReadFieldPredicate(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *GetRangeSlicesArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.RangeA1 = NewKeyRange()
	err832 := p.RangeA1.Read(iprot)
	if err832 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.RangeA1KeyRange", err832)
	}
	return err
}

func (p *GetRangeSlicesArgs) ReadFieldRange(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *GetRangeSlicesArgs) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v833, err834 := iprot.ReadI32()
	if err834 != nil {
		return thrift.NewTProtocolExceptionReadField(4, "consistency_level", p.ThriftName(), err834)
	}
	p.ConsistencyLevel = ConsistencyLevel(v833)
	return err
}

func (p *GetRangeSlicesArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *GetRangeSlicesArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("get_range_slices_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField4(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetRangeSlicesArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnParent != nil {
		err = oprot.WriteFieldBegin("column_parent", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "column_parent", p.ThriftName(), err)
		}
		err = p.ColumnParent.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnParent", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "column_parent", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetRangeSlicesArgs) WriteFieldColumnParent(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *GetRangeSlicesArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Predicate != nil {
		err = oprot.WriteFieldBegin("predicate", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "predicate", p.ThriftName(), err)
		}
		err = p.Predicate.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SlicePredicate", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "predicate", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetRangeSlicesArgs) WriteFieldPredicate(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *GetRangeSlicesArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.RangeA1 != nil {
		err = oprot.WriteFieldBegin("range", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "range", p.ThriftName(), err)
		}
		err = p.RangeA1.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("KeyRange", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "range", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetRangeSlicesArgs) WriteFieldRange(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *GetRangeSlicesArgs) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetRangeSlicesArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *GetRangeSlicesArgs) TStructName() string {
	return "GetRangeSlicesArgs"
}

func (p *GetRangeSlicesArgs) ThriftName() string {
	return "get_range_slices_args"
}

func (p *GetRangeSlicesArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetRangeSlicesArgs(%+v)", *p)
}

func (p *GetRangeSlicesArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*GetRangeSlicesArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *GetRangeSlicesArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.ColumnParent
	case 2:
		return p.Predicate
	case 3:
		return p.RangeA1
	case 4:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *GetRangeSlicesArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("column_parent", thrift.STRUCT, 1),
		thrift.NewTField("predicate", thrift.STRUCT, 2),
		thrift.NewTField("range", thrift.STRUCT, 3),
		thrift.NewTField("consistency_level", thrift.I32, 4),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Ue
 *  - Te
 */
type GetRangeSlicesResult struct {
	thrift.TStruct
	Success thrift.TList             "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
	Ue      *UnavailableException    "ue"      // 2
	Te      *TimedOutException       "te"      // 3
}

func NewGetRangeSlicesResult() *GetRangeSlicesResult {
	output := &GetRangeSlicesResult{
		TStruct: thrift.NewTStruct("get_range_slices_result", []thrift.TField{
			thrift.NewTField("success", thrift.LIST, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *GetRangeSlicesResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.LIST {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetRangeSlicesResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_etype840, _size837, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Success", "", err)
	}
	p.Success = thrift.NewTList(_etype840, _size837)
	for _i841 := 0; _i841 < _size837; _i841++ {
		_elem842 := NewKeySlice()
		err845 := _elem842.Read(iprot)
		if err845 != nil {
			return thrift.NewTProtocolExceptionReadStruct("_elem842KeySlice", err845)
		}
		p.Success.Push(_elem842)
	}
	err = iprot.ReadListEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
	}
	return err
}

func (p *GetRangeSlicesResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *GetRangeSlicesResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err848 := p.Ire.Read(iprot)
	if err848 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err848)
	}
	return err
}

func (p *GetRangeSlicesResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *GetRangeSlicesResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err851 := p.Ue.Read(iprot)
	if err851 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err851)
	}
	return err
}

func (p *GetRangeSlicesResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *GetRangeSlicesResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err854 := p.Te.Read(iprot)
	if err854 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err854)
	}
	return err
}

func (p *GetRangeSlicesResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *GetRangeSlicesResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("get_range_slices_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetRangeSlicesResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.LIST, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = oprot.WriteListBegin(thrift.STRUCT, p.Success.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		for Iter855 := range p.Success.Iter() {
			Iter856 := Iter855.(*KeySlice)
			err = Iter856.Write(oprot)
			if err != nil {
				return thrift.NewTProtocolExceptionWriteStruct("KeySlice", err)
			}
		}
		err = oprot.WriteListEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetRangeSlicesResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *GetRangeSlicesResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetRangeSlicesResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *GetRangeSlicesResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetRangeSlicesResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *GetRangeSlicesResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetRangeSlicesResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *GetRangeSlicesResult) TStructName() string {
	return "GetRangeSlicesResult"
}

func (p *GetRangeSlicesResult) ThriftName() string {
	return "get_range_slices_result"
}

func (p *GetRangeSlicesResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetRangeSlicesResult(%+v)", *p)
}

func (p *GetRangeSlicesResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*GetRangeSlicesResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *GetRangeSlicesResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *GetRangeSlicesResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.LIST, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - ColumnParent
 *  - IndexClause
 *  - ColumnPredicate
 *  - ConsistencyLevel
 */
type GetIndexedSlicesArgs struct {
	thrift.TStruct
	ColumnParent     *ColumnParent    "column_parent"     // 1
	IndexClause      *IndexClause     "index_clause"      // 2
	ColumnPredicate  *SlicePredicate  "column_predicate"  // 3
	ConsistencyLevel ConsistencyLevel "consistency_level" // 4
}

func NewGetIndexedSlicesArgs() *GetIndexedSlicesArgs {
	output := &GetIndexedSlicesArgs{
		TStruct: thrift.NewTStruct("get_indexed_slices_args", []thrift.TField{
			thrift.NewTField("column_parent", thrift.STRUCT, 1),
			thrift.NewTField("index_clause", thrift.STRUCT, 2),
			thrift.NewTField("column_predicate", thrift.STRUCT, 3),
			thrift.NewTField("consistency_level", thrift.I32, 4),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *GetIndexedSlicesArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *GetIndexedSlicesArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "column_parent" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "index_clause" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "column_predicate" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetIndexedSlicesArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnParent = NewColumnParent()
	err859 := p.ColumnParent.Read(iprot)
	if err859 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnParentColumnParent", err859)
	}
	return err
}

func (p *GetIndexedSlicesArgs) ReadFieldColumnParent(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *GetIndexedSlicesArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.IndexClause = NewIndexClause()
	err862 := p.IndexClause.Read(iprot)
	if err862 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IndexClauseIndexClause", err862)
	}
	return err
}

func (p *GetIndexedSlicesArgs) ReadFieldIndexClause(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *GetIndexedSlicesArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnPredicate = NewSlicePredicate()
	err865 := p.ColumnPredicate.Read(iprot)
	if err865 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnPredicateSlicePredicate", err865)
	}
	return err
}

func (p *GetIndexedSlicesArgs) ReadFieldColumnPredicate(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *GetIndexedSlicesArgs) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v866, err867 := iprot.ReadI32()
	if err867 != nil {
		return thrift.NewTProtocolExceptionReadField(4, "consistency_level", p.ThriftName(), err867)
	}
	p.ConsistencyLevel = ConsistencyLevel(v866)
	return err
}

func (p *GetIndexedSlicesArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *GetIndexedSlicesArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("get_indexed_slices_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField4(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetIndexedSlicesArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnParent != nil {
		err = oprot.WriteFieldBegin("column_parent", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "column_parent", p.ThriftName(), err)
		}
		err = p.ColumnParent.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnParent", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "column_parent", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetIndexedSlicesArgs) WriteFieldColumnParent(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *GetIndexedSlicesArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IndexClause != nil {
		err = oprot.WriteFieldBegin("index_clause", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "index_clause", p.ThriftName(), err)
		}
		err = p.IndexClause.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("IndexClause", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "index_clause", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetIndexedSlicesArgs) WriteFieldIndexClause(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *GetIndexedSlicesArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnPredicate != nil {
		err = oprot.WriteFieldBegin("column_predicate", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "column_predicate", p.ThriftName(), err)
		}
		err = p.ColumnPredicate.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SlicePredicate", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "column_predicate", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetIndexedSlicesArgs) WriteFieldColumnPredicate(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *GetIndexedSlicesArgs) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetIndexedSlicesArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *GetIndexedSlicesArgs) TStructName() string {
	return "GetIndexedSlicesArgs"
}

func (p *GetIndexedSlicesArgs) ThriftName() string {
	return "get_indexed_slices_args"
}

func (p *GetIndexedSlicesArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetIndexedSlicesArgs(%+v)", *p)
}

func (p *GetIndexedSlicesArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*GetIndexedSlicesArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *GetIndexedSlicesArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.ColumnParent
	case 2:
		return p.IndexClause
	case 3:
		return p.ColumnPredicate
	case 4:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *GetIndexedSlicesArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("column_parent", thrift.STRUCT, 1),
		thrift.NewTField("index_clause", thrift.STRUCT, 2),
		thrift.NewTField("column_predicate", thrift.STRUCT, 3),
		thrift.NewTField("consistency_level", thrift.I32, 4),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Ue
 *  - Te
 */
type GetIndexedSlicesResult struct {
	thrift.TStruct
	Success thrift.TList             "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
	Ue      *UnavailableException    "ue"      // 2
	Te      *TimedOutException       "te"      // 3
}

func NewGetIndexedSlicesResult() *GetIndexedSlicesResult {
	output := &GetIndexedSlicesResult{
		TStruct: thrift.NewTStruct("get_indexed_slices_result", []thrift.TField{
			thrift.NewTField("success", thrift.LIST, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *GetIndexedSlicesResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.LIST {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetIndexedSlicesResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_etype873, _size870, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Success", "", err)
	}
	p.Success = thrift.NewTList(_etype873, _size870)
	for _i874 := 0; _i874 < _size870; _i874++ {
		_elem875 := NewKeySlice()
		err878 := _elem875.Read(iprot)
		if err878 != nil {
			return thrift.NewTProtocolExceptionReadStruct("_elem875KeySlice", err878)
		}
		p.Success.Push(_elem875)
	}
	err = iprot.ReadListEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
	}
	return err
}

func (p *GetIndexedSlicesResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *GetIndexedSlicesResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err881 := p.Ire.Read(iprot)
	if err881 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err881)
	}
	return err
}

func (p *GetIndexedSlicesResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *GetIndexedSlicesResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err884 := p.Ue.Read(iprot)
	if err884 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err884)
	}
	return err
}

func (p *GetIndexedSlicesResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *GetIndexedSlicesResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err887 := p.Te.Read(iprot)
	if err887 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err887)
	}
	return err
}

func (p *GetIndexedSlicesResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *GetIndexedSlicesResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("get_indexed_slices_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *GetIndexedSlicesResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.LIST, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = oprot.WriteListBegin(thrift.STRUCT, p.Success.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		for Iter888 := range p.Success.Iter() {
			Iter889 := Iter888.(*KeySlice)
			err = Iter889.Write(oprot)
			if err != nil {
				return thrift.NewTProtocolExceptionWriteStruct("KeySlice", err)
			}
		}
		err = oprot.WriteListEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetIndexedSlicesResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *GetIndexedSlicesResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetIndexedSlicesResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *GetIndexedSlicesResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetIndexedSlicesResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *GetIndexedSlicesResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *GetIndexedSlicesResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *GetIndexedSlicesResult) TStructName() string {
	return "GetIndexedSlicesResult"
}

func (p *GetIndexedSlicesResult) ThriftName() string {
	return "get_indexed_slices_result"
}

func (p *GetIndexedSlicesResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetIndexedSlicesResult(%+v)", *p)
}

func (p *GetIndexedSlicesResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*GetIndexedSlicesResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *GetIndexedSlicesResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *GetIndexedSlicesResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.LIST, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - Key
 *  - ColumnParent
 *  - Column
 *  - ConsistencyLevel
 */
type InsertArgs struct {
	thrift.TStruct
	Key              []byte           "key"               // 1
	ColumnParent     *ColumnParent    "column_parent"     // 2
	Column           *Column          "column"            // 3
	ConsistencyLevel ConsistencyLevel "consistency_level" // 4
}

func NewInsertArgs() *InsertArgs {
	output := &InsertArgs{
		TStruct: thrift.NewTStruct("insert_args", []thrift.TField{
			thrift.NewTField("key", thrift.BINARY, 1),
			thrift.NewTField("column_parent", thrift.STRUCT, 2),
			thrift.NewTField("column", thrift.STRUCT, 3),
			thrift.NewTField("consistency_level", thrift.I32, 4),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *InsertArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *InsertArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "key" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "column_parent" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "column" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *InsertArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v890, err891 := iprot.ReadBinary()
	if err891 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "key", p.ThriftName(), err891)
	}
	p.Key = v890
	return err
}

func (p *InsertArgs) ReadFieldKey(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *InsertArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnParent = NewColumnParent()
	err894 := p.ColumnParent.Read(iprot)
	if err894 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnParentColumnParent", err894)
	}
	return err
}

func (p *InsertArgs) ReadFieldColumnParent(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *InsertArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Column = NewColumn()
	err897 := p.Column.Read(iprot)
	if err897 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnColumn", err897)
	}
	return err
}

func (p *InsertArgs) ReadFieldColumn(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *InsertArgs) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v898, err899 := iprot.ReadI32()
	if err899 != nil {
		return thrift.NewTProtocolExceptionReadField(4, "consistency_level", p.ThriftName(), err899)
	}
	p.ConsistencyLevel = ConsistencyLevel(v898)
	return err
}

func (p *InsertArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *InsertArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("insert_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField4(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *InsertArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Key != nil {
		err = oprot.WriteFieldBegin("key", thrift.BINARY, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteBinary(p.Key)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
	}
	return err
}

func (p *InsertArgs) WriteFieldKey(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *InsertArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnParent != nil {
		err = oprot.WriteFieldBegin("column_parent", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
		err = p.ColumnParent.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnParent", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
	}
	return err
}

func (p *InsertArgs) WriteFieldColumnParent(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *InsertArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Column != nil {
		err = oprot.WriteFieldBegin("column", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "column", p.ThriftName(), err)
		}
		err = p.Column.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("Column", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "column", p.ThriftName(), err)
		}
	}
	return err
}

func (p *InsertArgs) WriteFieldColumn(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *InsertArgs) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *InsertArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *InsertArgs) TStructName() string {
	return "InsertArgs"
}

func (p *InsertArgs) ThriftName() string {
	return "insert_args"
}

func (p *InsertArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("InsertArgs(%+v)", *p)
}

func (p *InsertArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*InsertArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *InsertArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Key
	case 2:
		return p.ColumnParent
	case 3:
		return p.Column
	case 4:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *InsertArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("key", thrift.BINARY, 1),
		thrift.NewTField("column_parent", thrift.STRUCT, 2),
		thrift.NewTField("column", thrift.STRUCT, 3),
		thrift.NewTField("consistency_level", thrift.I32, 4),
	})
}

/**
 * Attributes:
 *  - Ire
 *  - Ue
 *  - Te
 */
type InsertResult struct {
	thrift.TStruct
	Ire *InvalidRequestException "ire" // 1
	Ue  *UnavailableException    "ue"  // 2
	Te  *TimedOutException       "te"  // 3
}

func NewInsertResult() *InsertResult {
	output := &InsertResult{
		TStruct: thrift.NewTStruct("insert_result", []thrift.TField{
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *InsertResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *InsertResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err902 := p.Ire.Read(iprot)
	if err902 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err902)
	}
	return err
}

func (p *InsertResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *InsertResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err905 := p.Ue.Read(iprot)
	if err905 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err905)
	}
	return err
}

func (p *InsertResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *InsertResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err908 := p.Te.Read(iprot)
	if err908 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err908)
	}
	return err
}

func (p *InsertResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *InsertResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("insert_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *InsertResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *InsertResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *InsertResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *InsertResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *InsertResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *InsertResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *InsertResult) TStructName() string {
	return "InsertResult"
}

func (p *InsertResult) ThriftName() string {
	return "insert_result"
}

func (p *InsertResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("InsertResult(%+v)", *p)
}

func (p *InsertResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*InsertResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *InsertResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *InsertResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - Key
 *  - ColumnParent
 *  - Column
 *  - ConsistencyLevel
 */
type AddArgs struct {
	thrift.TStruct
	Key              []byte           "key"               // 1
	ColumnParent     *ColumnParent    "column_parent"     // 2
	Column           *CounterColumn   "column"            // 3
	ConsistencyLevel ConsistencyLevel "consistency_level" // 4
}

func NewAddArgs() *AddArgs {
	output := &AddArgs{
		TStruct: thrift.NewTStruct("add_args", []thrift.TField{
			thrift.NewTField("key", thrift.BINARY, 1),
			thrift.NewTField("column_parent", thrift.STRUCT, 2),
			thrift.NewTField("column", thrift.STRUCT, 3),
			thrift.NewTField("consistency_level", thrift.I32, 4),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *AddArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *AddArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "key" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "column_parent" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "column" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *AddArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v909, err910 := iprot.ReadBinary()
	if err910 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "key", p.ThriftName(), err910)
	}
	p.Key = v909
	return err
}

func (p *AddArgs) ReadFieldKey(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *AddArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnParent = NewColumnParent()
	err913 := p.ColumnParent.Read(iprot)
	if err913 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnParentColumnParent", err913)
	}
	return err
}

func (p *AddArgs) ReadFieldColumnParent(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *AddArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Column = NewCounterColumn()
	err916 := p.Column.Read(iprot)
	if err916 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnCounterColumn", err916)
	}
	return err
}

func (p *AddArgs) ReadFieldColumn(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *AddArgs) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v917, err918 := iprot.ReadI32()
	if err918 != nil {
		return thrift.NewTProtocolExceptionReadField(4, "consistency_level", p.ThriftName(), err918)
	}
	p.ConsistencyLevel = ConsistencyLevel(v917)
	return err
}

func (p *AddArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *AddArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("add_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField4(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *AddArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Key != nil {
		err = oprot.WriteFieldBegin("key", thrift.BINARY, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteBinary(p.Key)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
	}
	return err
}

func (p *AddArgs) WriteFieldKey(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *AddArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnParent != nil {
		err = oprot.WriteFieldBegin("column_parent", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
		err = p.ColumnParent.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnParent", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_parent", p.ThriftName(), err)
		}
	}
	return err
}

func (p *AddArgs) WriteFieldColumnParent(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *AddArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Column != nil {
		err = oprot.WriteFieldBegin("column", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "column", p.ThriftName(), err)
		}
		err = p.Column.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("CounterColumn", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "column", p.ThriftName(), err)
		}
	}
	return err
}

func (p *AddArgs) WriteFieldColumn(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *AddArgs) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *AddArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *AddArgs) TStructName() string {
	return "AddArgs"
}

func (p *AddArgs) ThriftName() string {
	return "add_args"
}

func (p *AddArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddArgs(%+v)", *p)
}

func (p *AddArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*AddArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *AddArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Key
	case 2:
		return p.ColumnParent
	case 3:
		return p.Column
	case 4:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *AddArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("key", thrift.BINARY, 1),
		thrift.NewTField("column_parent", thrift.STRUCT, 2),
		thrift.NewTField("column", thrift.STRUCT, 3),
		thrift.NewTField("consistency_level", thrift.I32, 4),
	})
}

/**
 * Attributes:
 *  - Ire
 *  - Ue
 *  - Te
 */
type AddResult struct {
	thrift.TStruct
	Ire *InvalidRequestException "ire" // 1
	Ue  *UnavailableException    "ue"  // 2
	Te  *TimedOutException       "te"  // 3
}

func NewAddResult() *AddResult {
	output := &AddResult{
		TStruct: thrift.NewTStruct("add_result", []thrift.TField{
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *AddResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *AddResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err921 := p.Ire.Read(iprot)
	if err921 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err921)
	}
	return err
}

func (p *AddResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *AddResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err924 := p.Ue.Read(iprot)
	if err924 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err924)
	}
	return err
}

func (p *AddResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *AddResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err927 := p.Te.Read(iprot)
	if err927 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err927)
	}
	return err
}

func (p *AddResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *AddResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("add_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *AddResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *AddResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *AddResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *AddResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *AddResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *AddResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *AddResult) TStructName() string {
	return "AddResult"
}

func (p *AddResult) ThriftName() string {
	return "add_result"
}

func (p *AddResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddResult(%+v)", *p)
}

func (p *AddResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*AddResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *AddResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *AddResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - Key
 *  - ColumnPath
 *  - Timestamp
 *  - ConsistencyLevel
 */
type RemoveArgs struct {
	thrift.TStruct
	Key              []byte           "key"               // 1
	ColumnPath       *ColumnPath      "column_path"       // 2
	Timestamp        int64            "timestamp"         // 3
	ConsistencyLevel ConsistencyLevel "consistency_level" // 4
}

func NewRemoveArgs() *RemoveArgs {
	output := &RemoveArgs{
		TStruct: thrift.NewTStruct("remove_args", []thrift.TField{
			thrift.NewTField("key", thrift.BINARY, 1),
			thrift.NewTField("column_path", thrift.STRUCT, 2),
			thrift.NewTField("timestamp", thrift.I64, 3),
			thrift.NewTField("consistency_level", thrift.I32, 4),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *RemoveArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *RemoveArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "key" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "column_path" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "timestamp" {
			if fieldTypeId == thrift.I64 {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *RemoveArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v928, err929 := iprot.ReadBinary()
	if err929 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "key", p.ThriftName(), err929)
	}
	p.Key = v928
	return err
}

func (p *RemoveArgs) ReadFieldKey(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *RemoveArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.ColumnPath = NewColumnPath()
	err932 := p.ColumnPath.Read(iprot)
	if err932 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.ColumnPathColumnPath", err932)
	}
	return err
}

func (p *RemoveArgs) ReadFieldColumnPath(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *RemoveArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v933, err934 := iprot.ReadI64()
	if err934 != nil {
		return thrift.NewTProtocolExceptionReadField(3, "timestamp", p.ThriftName(), err934)
	}
	p.Timestamp = v933
	return err
}

func (p *RemoveArgs) ReadFieldTimestamp(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *RemoveArgs) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v935, err936 := iprot.ReadI32()
	if err936 != nil {
		return thrift.NewTProtocolExceptionReadField(4, "consistency_level", p.ThriftName(), err936)
	}
	p.ConsistencyLevel = ConsistencyLevel(v935)
	return err
}

func (p *RemoveArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *RemoveArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("remove_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField4(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *RemoveArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Key != nil {
		err = oprot.WriteFieldBegin("key", thrift.BINARY, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteBinary(p.Key)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveArgs) WriteFieldKey(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *RemoveArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.ColumnPath != nil {
		err = oprot.WriteFieldBegin("column_path", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_path", p.ThriftName(), err)
		}
		err = p.ColumnPath.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnPath", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "column_path", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveArgs) WriteFieldColumnPath(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *RemoveArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("timestamp", thrift.I64, 3)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(3, "timestamp", p.ThriftName(), err)
	}
	err = oprot.WriteI64(int64(p.Timestamp))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(3, "timestamp", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(3, "timestamp", p.ThriftName(), err)
	}
	return err
}

func (p *RemoveArgs) WriteFieldTimestamp(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *RemoveArgs) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *RemoveArgs) TStructName() string {
	return "RemoveArgs"
}

func (p *RemoveArgs) ThriftName() string {
	return "remove_args"
}

func (p *RemoveArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RemoveArgs(%+v)", *p)
}

func (p *RemoveArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*RemoveArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *RemoveArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Key
	case 2:
		return p.ColumnPath
	case 3:
		return p.Timestamp
	case 4:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *RemoveArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("key", thrift.BINARY, 1),
		thrift.NewTField("column_path", thrift.STRUCT, 2),
		thrift.NewTField("timestamp", thrift.I64, 3),
		thrift.NewTField("consistency_level", thrift.I32, 4),
	})
}

/**
 * Attributes:
 *  - Ire
 *  - Ue
 *  - Te
 */
type RemoveResult struct {
	thrift.TStruct
	Ire *InvalidRequestException "ire" // 1
	Ue  *UnavailableException    "ue"  // 2
	Te  *TimedOutException       "te"  // 3
}

func NewRemoveResult() *RemoveResult {
	output := &RemoveResult{
		TStruct: thrift.NewTStruct("remove_result", []thrift.TField{
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *RemoveResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *RemoveResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err939 := p.Ire.Read(iprot)
	if err939 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err939)
	}
	return err
}

func (p *RemoveResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *RemoveResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err942 := p.Ue.Read(iprot)
	if err942 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err942)
	}
	return err
}

func (p *RemoveResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *RemoveResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err945 := p.Te.Read(iprot)
	if err945 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err945)
	}
	return err
}

func (p *RemoveResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *RemoveResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("remove_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *RemoveResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *RemoveResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *RemoveResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *RemoveResult) TStructName() string {
	return "RemoveResult"
}

func (p *RemoveResult) ThriftName() string {
	return "remove_result"
}

func (p *RemoveResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RemoveResult(%+v)", *p)
}

func (p *RemoveResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*RemoveResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *RemoveResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *RemoveResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - Key
 *  - Path
 *  - ConsistencyLevel
 */
type RemoveCounterArgs struct {
	thrift.TStruct
	Key              []byte           "key"               // 1
	Path             *ColumnPath      "path"              // 2
	ConsistencyLevel ConsistencyLevel "consistency_level" // 3
}

func NewRemoveCounterArgs() *RemoveCounterArgs {
	output := &RemoveCounterArgs{
		TStruct: thrift.NewTStruct("remove_counter_args", []thrift.TField{
			thrift.NewTField("key", thrift.BINARY, 1),
			thrift.NewTField("path", thrift.STRUCT, 2),
			thrift.NewTField("consistency_level", thrift.I32, 3),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *RemoveCounterArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *RemoveCounterArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "key" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "path" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *RemoveCounterArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v946, err947 := iprot.ReadBinary()
	if err947 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "key", p.ThriftName(), err947)
	}
	p.Key = v946
	return err
}

func (p *RemoveCounterArgs) ReadFieldKey(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *RemoveCounterArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Path = NewColumnPath()
	err950 := p.Path.Read(iprot)
	if err950 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.PathColumnPath", err950)
	}
	return err
}

func (p *RemoveCounterArgs) ReadFieldPath(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *RemoveCounterArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v951, err952 := iprot.ReadI32()
	if err952 != nil {
		return thrift.NewTProtocolExceptionReadField(3, "consistency_level", p.ThriftName(), err952)
	}
	p.ConsistencyLevel = ConsistencyLevel(v951)
	return err
}

func (p *RemoveCounterArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *RemoveCounterArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("remove_counter_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *RemoveCounterArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Key != nil {
		err = oprot.WriteFieldBegin("key", thrift.BINARY, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteBinary(p.Key)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "key", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveCounterArgs) WriteFieldKey(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *RemoveCounterArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Path != nil {
		err = oprot.WriteFieldBegin("path", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "path", p.ThriftName(), err)
		}
		err = p.Path.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("ColumnPath", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "path", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveCounterArgs) WriteFieldPath(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *RemoveCounterArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveCounterArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *RemoveCounterArgs) TStructName() string {
	return "RemoveCounterArgs"
}

func (p *RemoveCounterArgs) ThriftName() string {
	return "remove_counter_args"
}

func (p *RemoveCounterArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RemoveCounterArgs(%+v)", *p)
}

func (p *RemoveCounterArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*RemoveCounterArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *RemoveCounterArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Key
	case 2:
		return p.Path
	case 3:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *RemoveCounterArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("key", thrift.BINARY, 1),
		thrift.NewTField("path", thrift.STRUCT, 2),
		thrift.NewTField("consistency_level", thrift.I32, 3),
	})
}

/**
 * Attributes:
 *  - Ire
 *  - Ue
 *  - Te
 */
type RemoveCounterResult struct {
	thrift.TStruct
	Ire *InvalidRequestException "ire" // 1
	Ue  *UnavailableException    "ue"  // 2
	Te  *TimedOutException       "te"  // 3
}

func NewRemoveCounterResult() *RemoveCounterResult {
	output := &RemoveCounterResult{
		TStruct: thrift.NewTStruct("remove_counter_result", []thrift.TField{
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *RemoveCounterResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *RemoveCounterResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err955 := p.Ire.Read(iprot)
	if err955 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err955)
	}
	return err
}

func (p *RemoveCounterResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *RemoveCounterResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err958 := p.Ue.Read(iprot)
	if err958 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err958)
	}
	return err
}

func (p *RemoveCounterResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *RemoveCounterResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err961 := p.Te.Read(iprot)
	if err961 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err961)
	}
	return err
}

func (p *RemoveCounterResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *RemoveCounterResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("remove_counter_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *RemoveCounterResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveCounterResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *RemoveCounterResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveCounterResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *RemoveCounterResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *RemoveCounterResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *RemoveCounterResult) TStructName() string {
	return "RemoveCounterResult"
}

func (p *RemoveCounterResult) ThriftName() string {
	return "remove_counter_result"
}

func (p *RemoveCounterResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RemoveCounterResult(%+v)", *p)
}

func (p *RemoveCounterResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*RemoveCounterResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *RemoveCounterResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *RemoveCounterResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - MutationMap
 *  - ConsistencyLevel
 */
type BatchMutateArgs struct {
	thrift.TStruct
	MutationMap      thrift.TMap      "mutation_map"      // 1
	ConsistencyLevel ConsistencyLevel "consistency_level" // 2
}

func NewBatchMutateArgs() *BatchMutateArgs {
	output := &BatchMutateArgs{
		TStruct: thrift.NewTStruct("batch_mutate_args", []thrift.TField{
			thrift.NewTField("mutation_map", thrift.MAP, 1),
			thrift.NewTField("consistency_level", thrift.I32, 2),
		}),
	}
	{
		output.ConsistencyLevel = 1
	}
	return output
}

func (p *BatchMutateArgs) IsSetConsistencyLevel() bool {
	return int64(p.ConsistencyLevel) != math.MinInt32-1
}

func (p *BatchMutateArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "mutation_map" {
			if fieldTypeId == thrift.MAP {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "consistency_level" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *BatchMutateArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_ktype965, _vtype966, _size964, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.MutationMap", "", err)
	}
	p.MutationMap = thrift.NewTMap(_ktype965, _vtype966, _size964)
	for _i968 := 0; _i968 < _size964; _i968++ {
		v971, err972 := iprot.ReadBinary()
		if err972 != nil {
			return thrift.NewTProtocolExceptionReadField(0, "_key969", "", err972)
		}
		_key969 := v971
		_ktype976, _vtype977, _size975, err := iprot.ReadMapBegin()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(-1, "_val970", "", err)
		}
		_val970 := thrift.NewTMap(_ktype976, _vtype977, _size975)
		for _i979 := 0; _i979 < _size975; _i979++ {
			v982, err983 := iprot.ReadString()
			if err983 != nil {
				return thrift.NewTProtocolExceptionReadField(0, "_key980", "", err983)
			}
			_key980 := v982
			_etype989, _size986, err := iprot.ReadListBegin()
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(-1, "_val981", "", err)
			}
			_val981 := thrift.NewTList(_etype989, _size986)
			for _i990 := 0; _i990 < _size986; _i990++ {
				_elem991 := NewMutation()
				err994 := _elem991.Read(iprot)
				if err994 != nil {
					return thrift.NewTProtocolExceptionReadStruct("_elem991Mutation", err994)
				}
				_val981.Push(_elem991)
			}
			err = iprot.ReadListEnd()
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
			}
			_val970.Set(_key980, _val981)
		}
		err = iprot.ReadMapEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(-1, "", "map", err)
		}
		p.MutationMap.Set(_key969, _val970)
	}
	err = iprot.ReadMapEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "map", err)
	}
	return err
}

func (p *BatchMutateArgs) ReadFieldMutationMap(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *BatchMutateArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v995, err996 := iprot.ReadI32()
	if err996 != nil {
		return thrift.NewTProtocolExceptionReadField(2, "consistency_level", p.ThriftName(), err996)
	}
	p.ConsistencyLevel = ConsistencyLevel(v995)
	return err
}

func (p *BatchMutateArgs) ReadFieldConsistencyLevel(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *BatchMutateArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("batch_mutate_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *BatchMutateArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.MutationMap != nil {
		err = oprot.WriteFieldBegin("mutation_map", thrift.MAP, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "mutation_map", p.ThriftName(), err)
		}
		err = oprot.WriteMapBegin(thrift.BINARY, thrift.MAP, p.MutationMap.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "map", err)
		}
		for Miter997 := range p.MutationMap.Iter() {
			Kiter998, Viter999 := Miter997.Key().([]byte), Miter997.Value().(thrift.TMap)
			err = oprot.WriteBinary(Kiter998)
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(0, "Kiter998", "", err)
			}
			err = oprot.WriteMapBegin(thrift.STRING, thrift.LIST, Viter999.Len())
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(-1, "", "map", err)
			}
			for Miter1000 := range Viter999.Iter() {
				Kiter1001, Viter1002 := Miter1000.Key().(string), Miter1000.Value().(thrift.TList)
				err = oprot.WriteString(string(Kiter1001))
				if err != nil {
					return thrift.NewTProtocolExceptionWriteField(0, "Kiter1001", "", err)
				}
				err = oprot.WriteListBegin(thrift.STRUCT, Viter1002.Len())
				if err != nil {
					return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
				}
				for Iter1003 := range Viter1002.Iter() {
					Iter1004 := Iter1003.(*Mutation)
					err = Iter1004.Write(oprot)
					if err != nil {
						return thrift.NewTProtocolExceptionWriteStruct("Mutation", err)
					}
				}
				err = oprot.WriteListEnd()
				if err != nil {
					return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
				}
			}
			err = oprot.WriteMapEnd()
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(-1, "", "map", err)
			}
		}
		err = oprot.WriteMapEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "map", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "mutation_map", p.ThriftName(), err)
		}
	}
	return err
}

func (p *BatchMutateArgs) WriteFieldMutationMap(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *BatchMutateArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetConsistencyLevel() {
		err = oprot.WriteFieldBegin("consistency_level", thrift.I32, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.ConsistencyLevel))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "consistency_level", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "consistency_level", p.ThriftName(), err)
		}
	}
	return err
}

func (p *BatchMutateArgs) WriteFieldConsistencyLevel(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *BatchMutateArgs) TStructName() string {
	return "BatchMutateArgs"
}

func (p *BatchMutateArgs) ThriftName() string {
	return "batch_mutate_args"
}

func (p *BatchMutateArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BatchMutateArgs(%+v)", *p)
}

func (p *BatchMutateArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*BatchMutateArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *BatchMutateArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.MutationMap
	case 2:
		return p.ConsistencyLevel
	}
	return nil
}

func (p *BatchMutateArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("mutation_map", thrift.MAP, 1),
		thrift.NewTField("consistency_level", thrift.I32, 2),
	})
}

/**
 * Attributes:
 *  - Ire
 *  - Ue
 *  - Te
 */
type BatchMutateResult struct {
	thrift.TStruct
	Ire *InvalidRequestException "ire" // 1
	Ue  *UnavailableException    "ue"  // 2
	Te  *TimedOutException       "te"  // 3
}

func NewBatchMutateResult() *BatchMutateResult {
	output := &BatchMutateResult{
		TStruct: thrift.NewTStruct("batch_mutate_result", []thrift.TField{
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
		}),
	}
	{
	}
	return output
}

func (p *BatchMutateResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *BatchMutateResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1007 := p.Ire.Read(iprot)
	if err1007 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1007)
	}
	return err
}

func (p *BatchMutateResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *BatchMutateResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err1010 := p.Ue.Read(iprot)
	if err1010 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err1010)
	}
	return err
}

func (p *BatchMutateResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *BatchMutateResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err1013 := p.Te.Read(iprot)
	if err1013 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err1013)
	}
	return err
}

func (p *BatchMutateResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *BatchMutateResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("batch_mutate_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *BatchMutateResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *BatchMutateResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *BatchMutateResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *BatchMutateResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *BatchMutateResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *BatchMutateResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *BatchMutateResult) TStructName() string {
	return "BatchMutateResult"
}

func (p *BatchMutateResult) ThriftName() string {
	return "batch_mutate_result"
}

func (p *BatchMutateResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BatchMutateResult(%+v)", *p)
}

func (p *BatchMutateResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*BatchMutateResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *BatchMutateResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	}
	return nil
}

func (p *BatchMutateResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
	})
}

/**
 * Attributes:
 *  - Cfname
 */
type TruncateArgs struct {
	thrift.TStruct
	Cfname string "cfname" // 1
}

func NewTruncateArgs() *TruncateArgs {
	output := &TruncateArgs{
		TStruct: thrift.NewTStruct("truncate_args", []thrift.TField{
			thrift.NewTField("cfname", thrift.STRING, 1),
		}),
	}
	{
	}
	return output
}

func (p *TruncateArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "cfname" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *TruncateArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1014, err1015 := iprot.ReadString()
	if err1015 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "cfname", p.ThriftName(), err1015)
	}
	p.Cfname = v1014
	return err
}

func (p *TruncateArgs) ReadFieldCfname(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *TruncateArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("truncate_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *TruncateArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("cfname", thrift.STRING, 1)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "cfname", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Cfname))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "cfname", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "cfname", p.ThriftName(), err)
	}
	return err
}

func (p *TruncateArgs) WriteFieldCfname(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *TruncateArgs) TStructName() string {
	return "TruncateArgs"
}

func (p *TruncateArgs) ThriftName() string {
	return "truncate_args"
}

func (p *TruncateArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TruncateArgs(%+v)", *p)
}

func (p *TruncateArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*TruncateArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *TruncateArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Cfname
	}
	return nil
}

func (p *TruncateArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("cfname", thrift.STRING, 1),
	})
}

/**
 * Attributes:
 *  - Ire
 *  - Ue
 */
type TruncateResult struct {
	thrift.TStruct
	Ire *InvalidRequestException "ire" // 1
	Ue  *UnavailableException    "ue"  // 2
}

func NewTruncateResult() *TruncateResult {
	output := &TruncateResult{
		TStruct: thrift.NewTStruct("truncate_result", []thrift.TField{
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
		}),
	}
	{
	}
	return output
}

func (p *TruncateResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *TruncateResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1018 := p.Ire.Read(iprot)
	if err1018 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1018)
	}
	return err
}

func (p *TruncateResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *TruncateResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err1021 := p.Ue.Read(iprot)
	if err1021 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err1021)
	}
	return err
}

func (p *TruncateResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *TruncateResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("truncate_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *TruncateResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *TruncateResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *TruncateResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *TruncateResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *TruncateResult) TStructName() string {
	return "TruncateResult"
}

func (p *TruncateResult) ThriftName() string {
	return "truncate_result"
}

func (p *TruncateResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TruncateResult(%+v)", *p)
}

func (p *TruncateResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*TruncateResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *TruncateResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	}
	return nil
}

func (p *TruncateResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
	})
}

type DescribeSchemaVersionsArgs struct {
	thrift.TStruct
}

func NewDescribeSchemaVersionsArgs() *DescribeSchemaVersionsArgs {
	output := &DescribeSchemaVersionsArgs{
		TStruct: thrift.NewTStruct("describe_schema_versions_args", []thrift.TField{}),
	}
	{
	}
	return output
}

func (p *DescribeSchemaVersionsArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		err = iprot.Skip(fieldTypeId)
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSchemaVersionsArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_schema_versions_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSchemaVersionsArgs) TStructName() string {
	return "DescribeSchemaVersionsArgs"
}

func (p *DescribeSchemaVersionsArgs) ThriftName() string {
	return "describe_schema_versions_args"
}

func (p *DescribeSchemaVersionsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeSchemaVersionsArgs(%+v)", *p)
}

func (p *DescribeSchemaVersionsArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeSchemaVersionsArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeSchemaVersionsArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	}
	return nil
}

func (p *DescribeSchemaVersionsArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 */
type DescribeSchemaVersionsResult struct {
	thrift.TStruct
	Success thrift.TMap              "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
}

func NewDescribeSchemaVersionsResult() *DescribeSchemaVersionsResult {
	output := &DescribeSchemaVersionsResult{
		TStruct: thrift.NewTStruct("describe_schema_versions_result", []thrift.TField{
			thrift.NewTField("success", thrift.MAP, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
		}),
	}
	{
	}
	return output
}

func (p *DescribeSchemaVersionsResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.MAP {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSchemaVersionsResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_ktype1025, _vtype1026, _size1024, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Success", "", err)
	}
	p.Success = thrift.NewTMap(_ktype1025, _vtype1026, _size1024)
	for _i1028 := 0; _i1028 < _size1024; _i1028++ {
		v1031, err1032 := iprot.ReadString()
		if err1032 != nil {
			return thrift.NewTProtocolExceptionReadField(0, "_key1029", "", err1032)
		}
		_key1029 := v1031
		_etype1038, _size1035, err := iprot.ReadListBegin()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(-1, "_val1030", "", err)
		}
		_val1030 := thrift.NewTList(_etype1038, _size1035)
		for _i1039 := 0; _i1039 < _size1035; _i1039++ {
			v1041, err1042 := iprot.ReadString()
			if err1042 != nil {
				return thrift.NewTProtocolExceptionReadField(0, "_elem1040", "", err1042)
			}
			_elem1040 := v1041
			_val1030.Push(_elem1040)
		}
		err = iprot.ReadListEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
		}
		p.Success.Set(_key1029, _val1030)
	}
	err = iprot.ReadMapEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "map", err)
	}
	return err
}

func (p *DescribeSchemaVersionsResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *DescribeSchemaVersionsResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1045 := p.Ire.Read(iprot)
	if err1045 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1045)
	}
	return err
}

func (p *DescribeSchemaVersionsResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *DescribeSchemaVersionsResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_schema_versions_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSchemaVersionsResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.MAP, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = oprot.WriteMapBegin(thrift.STRING, thrift.LIST, p.Success.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "map", err)
		}
		for Miter1046 := range p.Success.Iter() {
			Kiter1047, Viter1048 := Miter1046.Key().(string), Miter1046.Value().(thrift.TList)
			err = oprot.WriteString(string(Kiter1047))
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(0, "Kiter1047", "", err)
			}
			err = oprot.WriteListBegin(thrift.STRING, Viter1048.Len())
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
			}
			for Iter1049 := range Viter1048.Iter() {
				Iter1050 := Iter1049.(string)
				err = oprot.WriteString(string(Iter1050))
				if err != nil {
					return thrift.NewTProtocolExceptionWriteField(0, "Iter1050", "", err)
				}
			}
			err = oprot.WriteListEnd()
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
			}
		}
		err = oprot.WriteMapEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "map", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeSchemaVersionsResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *DescribeSchemaVersionsResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeSchemaVersionsResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *DescribeSchemaVersionsResult) TStructName() string {
	return "DescribeSchemaVersionsResult"
}

func (p *DescribeSchemaVersionsResult) ThriftName() string {
	return "describe_schema_versions_result"
}

func (p *DescribeSchemaVersionsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeSchemaVersionsResult(%+v)", *p)
}

func (p *DescribeSchemaVersionsResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeSchemaVersionsResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeSchemaVersionsResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	}
	return nil
}

func (p *DescribeSchemaVersionsResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.MAP, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
	})
}

type DescribeKeyspacesArgs struct {
	thrift.TStruct
}

func NewDescribeKeyspacesArgs() *DescribeKeyspacesArgs {
	output := &DescribeKeyspacesArgs{
		TStruct: thrift.NewTStruct("describe_keyspaces_args", []thrift.TField{}),
	}
	{
	}
	return output
}

func (p *DescribeKeyspacesArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		err = iprot.Skip(fieldTypeId)
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeKeyspacesArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_keyspaces_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeKeyspacesArgs) TStructName() string {
	return "DescribeKeyspacesArgs"
}

func (p *DescribeKeyspacesArgs) ThriftName() string {
	return "describe_keyspaces_args"
}

func (p *DescribeKeyspacesArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeKeyspacesArgs(%+v)", *p)
}

func (p *DescribeKeyspacesArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeKeyspacesArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeKeyspacesArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	}
	return nil
}

func (p *DescribeKeyspacesArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 */
type DescribeKeyspacesResult struct {
	thrift.TStruct
	Success thrift.TList             "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
}

func NewDescribeKeyspacesResult() *DescribeKeyspacesResult {
	output := &DescribeKeyspacesResult{
		TStruct: thrift.NewTStruct("describe_keyspaces_result", []thrift.TField{
			thrift.NewTField("success", thrift.LIST, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
		}),
	}
	{
	}
	return output
}

func (p *DescribeKeyspacesResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.LIST {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeKeyspacesResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_etype1056, _size1053, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Success", "", err)
	}
	p.Success = thrift.NewTList(_etype1056, _size1053)
	for _i1057 := 0; _i1057 < _size1053; _i1057++ {
		_elem1058 := NewKsDef()
		err1061 := _elem1058.Read(iprot)
		if err1061 != nil {
			return thrift.NewTProtocolExceptionReadStruct("_elem1058KsDef", err1061)
		}
		p.Success.Push(_elem1058)
	}
	err = iprot.ReadListEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
	}
	return err
}

func (p *DescribeKeyspacesResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *DescribeKeyspacesResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1064 := p.Ire.Read(iprot)
	if err1064 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1064)
	}
	return err
}

func (p *DescribeKeyspacesResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *DescribeKeyspacesResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_keyspaces_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeKeyspacesResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.LIST, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = oprot.WriteListBegin(thrift.STRUCT, p.Success.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		for Iter1065 := range p.Success.Iter() {
			Iter1066 := Iter1065.(*KsDef)
			err = Iter1066.Write(oprot)
			if err != nil {
				return thrift.NewTProtocolExceptionWriteStruct("KsDef", err)
			}
		}
		err = oprot.WriteListEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeKeyspacesResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *DescribeKeyspacesResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeKeyspacesResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *DescribeKeyspacesResult) TStructName() string {
	return "DescribeKeyspacesResult"
}

func (p *DescribeKeyspacesResult) ThriftName() string {
	return "describe_keyspaces_result"
}

func (p *DescribeKeyspacesResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeKeyspacesResult(%+v)", *p)
}

func (p *DescribeKeyspacesResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeKeyspacesResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeKeyspacesResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	}
	return nil
}

func (p *DescribeKeyspacesResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.LIST, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
	})
}

type DescribeClusterNameArgs struct {
	thrift.TStruct
}

func NewDescribeClusterNameArgs() *DescribeClusterNameArgs {
	output := &DescribeClusterNameArgs{
		TStruct: thrift.NewTStruct("describe_cluster_name_args", []thrift.TField{}),
	}
	{
	}
	return output
}

func (p *DescribeClusterNameArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		err = iprot.Skip(fieldTypeId)
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeClusterNameArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_cluster_name_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeClusterNameArgs) TStructName() string {
	return "DescribeClusterNameArgs"
}

func (p *DescribeClusterNameArgs) ThriftName() string {
	return "describe_cluster_name_args"
}

func (p *DescribeClusterNameArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeClusterNameArgs(%+v)", *p)
}

func (p *DescribeClusterNameArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeClusterNameArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeClusterNameArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	}
	return nil
}

func (p *DescribeClusterNameArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{})
}

/**
 * Attributes:
 *  - Success
 */
type DescribeClusterNameResult struct {
	thrift.TStruct
	Success string "success" // 0
}

func NewDescribeClusterNameResult() *DescribeClusterNameResult {
	output := &DescribeClusterNameResult{
		TStruct: thrift.NewTStruct("describe_cluster_name_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRING, 0),
		}),
	}
	{
	}
	return output
}

func (p *DescribeClusterNameResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeClusterNameResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1067, err1068 := iprot.ReadString()
	if err1068 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err1068)
	}
	p.Success = v1067
	return err
}

func (p *DescribeClusterNameResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *DescribeClusterNameResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_cluster_name_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeClusterNameResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.STRING, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *DescribeClusterNameResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *DescribeClusterNameResult) TStructName() string {
	return "DescribeClusterNameResult"
}

func (p *DescribeClusterNameResult) ThriftName() string {
	return "describe_cluster_name_result"
}

func (p *DescribeClusterNameResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeClusterNameResult(%+v)", *p)
}

func (p *DescribeClusterNameResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeClusterNameResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeClusterNameResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	}
	return nil
}

func (p *DescribeClusterNameResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRING, 0),
	})
}

type DescribeVersionArgs struct {
	thrift.TStruct
}

func NewDescribeVersionArgs() *DescribeVersionArgs {
	output := &DescribeVersionArgs{
		TStruct: thrift.NewTStruct("describe_version_args", []thrift.TField{}),
	}
	{
	}
	return output
}

func (p *DescribeVersionArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		err = iprot.Skip(fieldTypeId)
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeVersionArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_version_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeVersionArgs) TStructName() string {
	return "DescribeVersionArgs"
}

func (p *DescribeVersionArgs) ThriftName() string {
	return "describe_version_args"
}

func (p *DescribeVersionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeVersionArgs(%+v)", *p)
}

func (p *DescribeVersionArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeVersionArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeVersionArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	}
	return nil
}

func (p *DescribeVersionArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{})
}

/**
 * Attributes:
 *  - Success
 */
type DescribeVersionResult struct {
	thrift.TStruct
	Success string "success" // 0
}

func NewDescribeVersionResult() *DescribeVersionResult {
	output := &DescribeVersionResult{
		TStruct: thrift.NewTStruct("describe_version_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRING, 0),
		}),
	}
	{
	}
	return output
}

func (p *DescribeVersionResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeVersionResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1069, err1070 := iprot.ReadString()
	if err1070 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err1070)
	}
	p.Success = v1069
	return err
}

func (p *DescribeVersionResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *DescribeVersionResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_version_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeVersionResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.STRING, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *DescribeVersionResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *DescribeVersionResult) TStructName() string {
	return "DescribeVersionResult"
}

func (p *DescribeVersionResult) ThriftName() string {
	return "describe_version_result"
}

func (p *DescribeVersionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeVersionResult(%+v)", *p)
}

func (p *DescribeVersionResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeVersionResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeVersionResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	}
	return nil
}

func (p *DescribeVersionResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRING, 0),
	})
}

/**
 * Attributes:
 *  - Keyspace
 */
type DescribeRingArgs struct {
	thrift.TStruct
	Keyspace string "keyspace" // 1
}

func NewDescribeRingArgs() *DescribeRingArgs {
	output := &DescribeRingArgs{
		TStruct: thrift.NewTStruct("describe_ring_args", []thrift.TField{
			thrift.NewTField("keyspace", thrift.STRING, 1),
		}),
	}
	{
	}
	return output
}

func (p *DescribeRingArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "keyspace" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeRingArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1071, err1072 := iprot.ReadString()
	if err1072 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "keyspace", p.ThriftName(), err1072)
	}
	p.Keyspace = v1071
	return err
}

func (p *DescribeRingArgs) ReadFieldKeyspace(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *DescribeRingArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_ring_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeRingArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("keyspace", thrift.STRING, 1)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Keyspace))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	return err
}

func (p *DescribeRingArgs) WriteFieldKeyspace(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *DescribeRingArgs) TStructName() string {
	return "DescribeRingArgs"
}

func (p *DescribeRingArgs) ThriftName() string {
	return "describe_ring_args"
}

func (p *DescribeRingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeRingArgs(%+v)", *p)
}

func (p *DescribeRingArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeRingArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeRingArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Keyspace
	}
	return nil
}

func (p *DescribeRingArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("keyspace", thrift.STRING, 1),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 */
type DescribeRingResult struct {
	thrift.TStruct
	Success thrift.TList             "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
}

func NewDescribeRingResult() *DescribeRingResult {
	output := &DescribeRingResult{
		TStruct: thrift.NewTStruct("describe_ring_result", []thrift.TField{
			thrift.NewTField("success", thrift.LIST, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
		}),
	}
	{
	}
	return output
}

func (p *DescribeRingResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.LIST {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeRingResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_etype1078, _size1075, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Success", "", err)
	}
	p.Success = thrift.NewTList(_etype1078, _size1075)
	for _i1079 := 0; _i1079 < _size1075; _i1079++ {
		_elem1080 := NewTokenRange()
		err1083 := _elem1080.Read(iprot)
		if err1083 != nil {
			return thrift.NewTProtocolExceptionReadStruct("_elem1080TokenRange", err1083)
		}
		p.Success.Push(_elem1080)
	}
	err = iprot.ReadListEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
	}
	return err
}

func (p *DescribeRingResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *DescribeRingResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1086 := p.Ire.Read(iprot)
	if err1086 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1086)
	}
	return err
}

func (p *DescribeRingResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *DescribeRingResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_ring_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeRingResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.LIST, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = oprot.WriteListBegin(thrift.STRUCT, p.Success.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		for Iter1087 := range p.Success.Iter() {
			Iter1088 := Iter1087.(*TokenRange)
			err = Iter1088.Write(oprot)
			if err != nil {
				return thrift.NewTProtocolExceptionWriteStruct("TokenRange", err)
			}
		}
		err = oprot.WriteListEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeRingResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *DescribeRingResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeRingResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *DescribeRingResult) TStructName() string {
	return "DescribeRingResult"
}

func (p *DescribeRingResult) ThriftName() string {
	return "describe_ring_result"
}

func (p *DescribeRingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeRingResult(%+v)", *p)
}

func (p *DescribeRingResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeRingResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeRingResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	}
	return nil
}

func (p *DescribeRingResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.LIST, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
	})
}

type DescribePartitionerArgs struct {
	thrift.TStruct
}

func NewDescribePartitionerArgs() *DescribePartitionerArgs {
	output := &DescribePartitionerArgs{
		TStruct: thrift.NewTStruct("describe_partitioner_args", []thrift.TField{}),
	}
	{
	}
	return output
}

func (p *DescribePartitionerArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		err = iprot.Skip(fieldTypeId)
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribePartitionerArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_partitioner_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribePartitionerArgs) TStructName() string {
	return "DescribePartitionerArgs"
}

func (p *DescribePartitionerArgs) ThriftName() string {
	return "describe_partitioner_args"
}

func (p *DescribePartitionerArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribePartitionerArgs(%+v)", *p)
}

func (p *DescribePartitionerArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribePartitionerArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribePartitionerArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	}
	return nil
}

func (p *DescribePartitionerArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{})
}

/**
 * Attributes:
 *  - Success
 */
type DescribePartitionerResult struct {
	thrift.TStruct
	Success string "success" // 0
}

func NewDescribePartitionerResult() *DescribePartitionerResult {
	output := &DescribePartitionerResult{
		TStruct: thrift.NewTStruct("describe_partitioner_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRING, 0),
		}),
	}
	{
	}
	return output
}

func (p *DescribePartitionerResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribePartitionerResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1089, err1090 := iprot.ReadString()
	if err1090 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err1090)
	}
	p.Success = v1089
	return err
}

func (p *DescribePartitionerResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *DescribePartitionerResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_partitioner_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribePartitionerResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.STRING, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *DescribePartitionerResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *DescribePartitionerResult) TStructName() string {
	return "DescribePartitionerResult"
}

func (p *DescribePartitionerResult) ThriftName() string {
	return "describe_partitioner_result"
}

func (p *DescribePartitionerResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribePartitionerResult(%+v)", *p)
}

func (p *DescribePartitionerResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribePartitionerResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribePartitionerResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	}
	return nil
}

func (p *DescribePartitionerResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRING, 0),
	})
}

type DescribeSnitchArgs struct {
	thrift.TStruct
}

func NewDescribeSnitchArgs() *DescribeSnitchArgs {
	output := &DescribeSnitchArgs{
		TStruct: thrift.NewTStruct("describe_snitch_args", []thrift.TField{}),
	}
	{
	}
	return output
}

func (p *DescribeSnitchArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		err = iprot.Skip(fieldTypeId)
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSnitchArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_snitch_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSnitchArgs) TStructName() string {
	return "DescribeSnitchArgs"
}

func (p *DescribeSnitchArgs) ThriftName() string {
	return "describe_snitch_args"
}

func (p *DescribeSnitchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeSnitchArgs(%+v)", *p)
}

func (p *DescribeSnitchArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeSnitchArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeSnitchArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	}
	return nil
}

func (p *DescribeSnitchArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{})
}

/**
 * Attributes:
 *  - Success
 */
type DescribeSnitchResult struct {
	thrift.TStruct
	Success string "success" // 0
}

func NewDescribeSnitchResult() *DescribeSnitchResult {
	output := &DescribeSnitchResult{
		TStruct: thrift.NewTStruct("describe_snitch_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRING, 0),
		}),
	}
	{
	}
	return output
}

func (p *DescribeSnitchResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSnitchResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1091, err1092 := iprot.ReadString()
	if err1092 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err1092)
	}
	p.Success = v1091
	return err
}

func (p *DescribeSnitchResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *DescribeSnitchResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_snitch_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSnitchResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.STRING, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSnitchResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *DescribeSnitchResult) TStructName() string {
	return "DescribeSnitchResult"
}

func (p *DescribeSnitchResult) ThriftName() string {
	return "describe_snitch_result"
}

func (p *DescribeSnitchResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeSnitchResult(%+v)", *p)
}

func (p *DescribeSnitchResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeSnitchResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeSnitchResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	}
	return nil
}

func (p *DescribeSnitchResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRING, 0),
	})
}

/**
 * Attributes:
 *  - Keyspace
 */
type DescribeKeyspaceArgs struct {
	thrift.TStruct
	Keyspace string "keyspace" // 1
}

func NewDescribeKeyspaceArgs() *DescribeKeyspaceArgs {
	output := &DescribeKeyspaceArgs{
		TStruct: thrift.NewTStruct("describe_keyspace_args", []thrift.TField{
			thrift.NewTField("keyspace", thrift.STRING, 1),
		}),
	}
	{
	}
	return output
}

func (p *DescribeKeyspaceArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "keyspace" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeKeyspaceArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1093, err1094 := iprot.ReadString()
	if err1094 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "keyspace", p.ThriftName(), err1094)
	}
	p.Keyspace = v1093
	return err
}

func (p *DescribeKeyspaceArgs) ReadFieldKeyspace(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *DescribeKeyspaceArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_keyspace_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeKeyspaceArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("keyspace", thrift.STRING, 1)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Keyspace))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	return err
}

func (p *DescribeKeyspaceArgs) WriteFieldKeyspace(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *DescribeKeyspaceArgs) TStructName() string {
	return "DescribeKeyspaceArgs"
}

func (p *DescribeKeyspaceArgs) ThriftName() string {
	return "describe_keyspace_args"
}

func (p *DescribeKeyspaceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeKeyspaceArgs(%+v)", *p)
}

func (p *DescribeKeyspaceArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeKeyspaceArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeKeyspaceArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Keyspace
	}
	return nil
}

func (p *DescribeKeyspaceArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("keyspace", thrift.STRING, 1),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Nfe
 *  - Ire
 */
type DescribeKeyspaceResult struct {
	thrift.TStruct
	Success *KsDef                   "success" // 0
	Nfe     *NotFoundException       "nfe"     // 1
	Ire     *InvalidRequestException "ire"     // 2
}

func NewDescribeKeyspaceResult() *DescribeKeyspaceResult {
	output := &DescribeKeyspaceResult{
		TStruct: thrift.NewTStruct("describe_keyspace_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRUCT, 0),
			thrift.NewTField("nfe", thrift.STRUCT, 1),
			thrift.NewTField("ire", thrift.STRUCT, 2),
		}),
	}
	{
	}
	return output
}

func (p *DescribeKeyspaceResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "nfe" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeKeyspaceResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Success = NewKsDef()
	err1097 := p.Success.Read(iprot)
	if err1097 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.SuccessKsDef", err1097)
	}
	return err
}

func (p *DescribeKeyspaceResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *DescribeKeyspaceResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Nfe = NewNotFoundException()
	err1100 := p.Nfe.Read(iprot)
	if err1100 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.NfeNotFoundException", err1100)
	}
	return err
}

func (p *DescribeKeyspaceResult) ReadFieldNfe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *DescribeKeyspaceResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1103 := p.Ire.Read(iprot)
	if err1103 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1103)
	}
	return err
}

func (p *DescribeKeyspaceResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *DescribeKeyspaceResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_keyspace_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Ire != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Nfe != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeKeyspaceResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = p.Success.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("KsDef", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeKeyspaceResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *DescribeKeyspaceResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Nfe != nil {
		err = oprot.WriteFieldBegin("nfe", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "nfe", p.ThriftName(), err)
		}
		err = p.Nfe.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("NotFoundException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "nfe", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeKeyspaceResult) WriteFieldNfe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *DescribeKeyspaceResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeKeyspaceResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *DescribeKeyspaceResult) TStructName() string {
	return "DescribeKeyspaceResult"
}

func (p *DescribeKeyspaceResult) ThriftName() string {
	return "describe_keyspace_result"
}

func (p *DescribeKeyspaceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeKeyspaceResult(%+v)", *p)
}

func (p *DescribeKeyspaceResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeKeyspaceResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeKeyspaceResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Nfe
	case 2:
		return p.Ire
	}
	return nil
}

func (p *DescribeKeyspaceResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRUCT, 0),
		thrift.NewTField("nfe", thrift.STRUCT, 1),
		thrift.NewTField("ire", thrift.STRUCT, 2),
	})
}

/**
 * Attributes:
 *  - CfName
 *  - StartToken
 *  - EndToken
 *  - KeysPerSplit
 */
type DescribeSplitsArgs struct {
	thrift.TStruct
	CfName       string "cfName"         // 1
	StartToken   string "start_token"    // 2
	EndToken     string "end_token"      // 3
	KeysPerSplit int32  "keys_per_split" // 4
}

func NewDescribeSplitsArgs() *DescribeSplitsArgs {
	output := &DescribeSplitsArgs{
		TStruct: thrift.NewTStruct("describe_splits_args", []thrift.TField{
			thrift.NewTField("cfName", thrift.STRING, 1),
			thrift.NewTField("start_token", thrift.STRING, 2),
			thrift.NewTField("end_token", thrift.STRING, 3),
			thrift.NewTField("keys_per_split", thrift.I32, 4),
		}),
	}
	{
	}
	return output
}

func (p *DescribeSplitsArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "cfName" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "start_token" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "end_token" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "keys_per_split" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSplitsArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1104, err1105 := iprot.ReadString()
	if err1105 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "cfName", p.ThriftName(), err1105)
	}
	p.CfName = v1104
	return err
}

func (p *DescribeSplitsArgs) ReadFieldCfName(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *DescribeSplitsArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1106, err1107 := iprot.ReadString()
	if err1107 != nil {
		return thrift.NewTProtocolExceptionReadField(2, "start_token", p.ThriftName(), err1107)
	}
	p.StartToken = v1106
	return err
}

func (p *DescribeSplitsArgs) ReadFieldStartToken(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *DescribeSplitsArgs) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1108, err1109 := iprot.ReadString()
	if err1109 != nil {
		return thrift.NewTProtocolExceptionReadField(3, "end_token", p.ThriftName(), err1109)
	}
	p.EndToken = v1108
	return err
}

func (p *DescribeSplitsArgs) ReadFieldEndToken(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *DescribeSplitsArgs) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1110, err1111 := iprot.ReadI32()
	if err1111 != nil {
		return thrift.NewTProtocolExceptionReadField(4, "keys_per_split", p.ThriftName(), err1111)
	}
	p.KeysPerSplit = v1110
	return err
}

func (p *DescribeSplitsArgs) ReadFieldKeysPerSplit(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *DescribeSplitsArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_splits_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField3(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField4(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSplitsArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("cfName", thrift.STRING, 1)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "cfName", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.CfName))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "cfName", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "cfName", p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSplitsArgs) WriteFieldCfName(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *DescribeSplitsArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("start_token", thrift.STRING, 2)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(2, "start_token", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.StartToken))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(2, "start_token", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(2, "start_token", p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSplitsArgs) WriteFieldStartToken(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *DescribeSplitsArgs) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("end_token", thrift.STRING, 3)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(3, "end_token", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.EndToken))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(3, "end_token", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(3, "end_token", p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSplitsArgs) WriteFieldEndToken(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *DescribeSplitsArgs) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("keys_per_split", thrift.I32, 4)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(4, "keys_per_split", p.ThriftName(), err)
	}
	err = oprot.WriteI32(int32(p.KeysPerSplit))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(4, "keys_per_split", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(4, "keys_per_split", p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSplitsArgs) WriteFieldKeysPerSplit(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *DescribeSplitsArgs) TStructName() string {
	return "DescribeSplitsArgs"
}

func (p *DescribeSplitsArgs) ThriftName() string {
	return "describe_splits_args"
}

func (p *DescribeSplitsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeSplitsArgs(%+v)", *p)
}

func (p *DescribeSplitsArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeSplitsArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeSplitsArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.CfName
	case 2:
		return p.StartToken
	case 3:
		return p.EndToken
	case 4:
		return p.KeysPerSplit
	}
	return nil
}

func (p *DescribeSplitsArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("cfName", thrift.STRING, 1),
		thrift.NewTField("start_token", thrift.STRING, 2),
		thrift.NewTField("end_token", thrift.STRING, 3),
		thrift.NewTField("keys_per_split", thrift.I32, 4),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 */
type DescribeSplitsResult struct {
	thrift.TStruct
	Success thrift.TList             "success" // 0
	Ire     *InvalidRequestException "ire"     // 1
}

func NewDescribeSplitsResult() *DescribeSplitsResult {
	output := &DescribeSplitsResult{
		TStruct: thrift.NewTStruct("describe_splits_result", []thrift.TField{
			thrift.NewTField("success", thrift.LIST, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
		}),
	}
	{
	}
	return output
}

func (p *DescribeSplitsResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.LIST {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSplitsResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_etype1117, _size1114, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "p.Success", "", err)
	}
	p.Success = thrift.NewTList(_etype1117, _size1114)
	for _i1118 := 0; _i1118 < _size1114; _i1118++ {
		v1120, err1121 := iprot.ReadString()
		if err1121 != nil {
			return thrift.NewTProtocolExceptionReadField(0, "_elem1119", "", err1121)
		}
		_elem1119 := v1120
		p.Success.Push(_elem1119)
	}
	err = iprot.ReadListEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadField(-1, "", "list", err)
	}
	return err
}

func (p *DescribeSplitsResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *DescribeSplitsResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1124 := p.Ire.Read(iprot)
	if err1124 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1124)
	}
	return err
}

func (p *DescribeSplitsResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *DescribeSplitsResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("describe_splits_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *DescribeSplitsResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.LIST, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = oprot.WriteListBegin(thrift.STRING, p.Success.Len())
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		for Iter1125 := range p.Success.Iter() {
			Iter1126 := Iter1125.(string)
			err = oprot.WriteString(string(Iter1126))
			if err != nil {
				return thrift.NewTProtocolExceptionWriteField(0, "Iter1126", "", err)
			}
		}
		err = oprot.WriteListEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(-1, "", "list", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeSplitsResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *DescribeSplitsResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *DescribeSplitsResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *DescribeSplitsResult) TStructName() string {
	return "DescribeSplitsResult"
}

func (p *DescribeSplitsResult) ThriftName() string {
	return "describe_splits_result"
}

func (p *DescribeSplitsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DescribeSplitsResult(%+v)", *p)
}

func (p *DescribeSplitsResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*DescribeSplitsResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *DescribeSplitsResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	}
	return nil
}

func (p *DescribeSplitsResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.LIST, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
	})
}

/**
 * Attributes:
 *  - CfDef
 */
type SystemAddColumnFamilyArgs struct {
	thrift.TStruct
	CfDef *CfDef "cf_def" // 1
}

func NewSystemAddColumnFamilyArgs() *SystemAddColumnFamilyArgs {
	output := &SystemAddColumnFamilyArgs{
		TStruct: thrift.NewTStruct("system_add_column_family_args", []thrift.TField{
			thrift.NewTField("cf_def", thrift.STRUCT, 1),
		}),
	}
	{
	}
	return output
}

func (p *SystemAddColumnFamilyArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "cf_def" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemAddColumnFamilyArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.CfDef = NewCfDef()
	err1129 := p.CfDef.Read(iprot)
	if err1129 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.CfDefCfDef", err1129)
	}
	return err
}

func (p *SystemAddColumnFamilyArgs) ReadFieldCfDef(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemAddColumnFamilyArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_add_column_family_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemAddColumnFamilyArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.CfDef != nil {
		err = oprot.WriteFieldBegin("cf_def", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "cf_def", p.ThriftName(), err)
		}
		err = p.CfDef.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("CfDef", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "cf_def", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemAddColumnFamilyArgs) WriteFieldCfDef(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemAddColumnFamilyArgs) TStructName() string {
	return "SystemAddColumnFamilyArgs"
}

func (p *SystemAddColumnFamilyArgs) ThriftName() string {
	return "system_add_column_family_args"
}

func (p *SystemAddColumnFamilyArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemAddColumnFamilyArgs(%+v)", *p)
}

func (p *SystemAddColumnFamilyArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemAddColumnFamilyArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemAddColumnFamilyArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.CfDef
	}
	return nil
}

func (p *SystemAddColumnFamilyArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("cf_def", thrift.STRUCT, 1),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Sde
 */
type SystemAddColumnFamilyResult struct {
	thrift.TStruct
	Success string                       "success" // 0
	Ire     *InvalidRequestException     "ire"     // 1
	Sde     *SchemaDisagreementException "sde"     // 2
}

func NewSystemAddColumnFamilyResult() *SystemAddColumnFamilyResult {
	output := &SystemAddColumnFamilyResult{
		TStruct: thrift.NewTStruct("system_add_column_family_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRING, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("sde", thrift.STRUCT, 2),
		}),
	}
	{
	}
	return output
}

func (p *SystemAddColumnFamilyResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "sde" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemAddColumnFamilyResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1130, err1131 := iprot.ReadString()
	if err1131 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err1131)
	}
	p.Success = v1130
	return err
}

func (p *SystemAddColumnFamilyResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *SystemAddColumnFamilyResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1134 := p.Ire.Read(iprot)
	if err1134 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1134)
	}
	return err
}

func (p *SystemAddColumnFamilyResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemAddColumnFamilyResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Sde = NewSchemaDisagreementException()
	err1137 := p.Sde.Read(iprot)
	if err1137 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.SdeSchemaDisagreementException", err1137)
	}
	return err
}

func (p *SystemAddColumnFamilyResult) ReadFieldSde(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *SystemAddColumnFamilyResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_add_column_family_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Sde != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemAddColumnFamilyResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.STRING, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *SystemAddColumnFamilyResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *SystemAddColumnFamilyResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemAddColumnFamilyResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemAddColumnFamilyResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Sde != nil {
		err = oprot.WriteFieldBegin("sde", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
		err = p.Sde.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SchemaDisagreementException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemAddColumnFamilyResult) WriteFieldSde(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *SystemAddColumnFamilyResult) TStructName() string {
	return "SystemAddColumnFamilyResult"
}

func (p *SystemAddColumnFamilyResult) ThriftName() string {
	return "system_add_column_family_result"
}

func (p *SystemAddColumnFamilyResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemAddColumnFamilyResult(%+v)", *p)
}

func (p *SystemAddColumnFamilyResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemAddColumnFamilyResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemAddColumnFamilyResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Sde
	}
	return nil
}

func (p *SystemAddColumnFamilyResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRING, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("sde", thrift.STRUCT, 2),
	})
}

/**
 * Attributes:
 *  - ColumnFamily
 */
type SystemDropColumnFamilyArgs struct {
	thrift.TStruct
	ColumnFamily string "column_family" // 1
}

func NewSystemDropColumnFamilyArgs() *SystemDropColumnFamilyArgs {
	output := &SystemDropColumnFamilyArgs{
		TStruct: thrift.NewTStruct("system_drop_column_family_args", []thrift.TField{
			thrift.NewTField("column_family", thrift.STRING, 1),
		}),
	}
	{
	}
	return output
}

func (p *SystemDropColumnFamilyArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "column_family" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropColumnFamilyArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1138, err1139 := iprot.ReadString()
	if err1139 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "column_family", p.ThriftName(), err1139)
	}
	p.ColumnFamily = v1138
	return err
}

func (p *SystemDropColumnFamilyArgs) ReadFieldColumnFamily(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemDropColumnFamilyArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_drop_column_family_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropColumnFamilyArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("column_family", thrift.STRING, 1)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "column_family", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.ColumnFamily))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "column_family", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "column_family", p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropColumnFamilyArgs) WriteFieldColumnFamily(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemDropColumnFamilyArgs) TStructName() string {
	return "SystemDropColumnFamilyArgs"
}

func (p *SystemDropColumnFamilyArgs) ThriftName() string {
	return "system_drop_column_family_args"
}

func (p *SystemDropColumnFamilyArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemDropColumnFamilyArgs(%+v)", *p)
}

func (p *SystemDropColumnFamilyArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemDropColumnFamilyArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemDropColumnFamilyArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.ColumnFamily
	}
	return nil
}

func (p *SystemDropColumnFamilyArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("column_family", thrift.STRING, 1),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Sde
 */
type SystemDropColumnFamilyResult struct {
	thrift.TStruct
	Success string                       "success" // 0
	Ire     *InvalidRequestException     "ire"     // 1
	Sde     *SchemaDisagreementException "sde"     // 2
}

func NewSystemDropColumnFamilyResult() *SystemDropColumnFamilyResult {
	output := &SystemDropColumnFamilyResult{
		TStruct: thrift.NewTStruct("system_drop_column_family_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRING, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("sde", thrift.STRUCT, 2),
		}),
	}
	{
	}
	return output
}

func (p *SystemDropColumnFamilyResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "sde" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropColumnFamilyResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1140, err1141 := iprot.ReadString()
	if err1141 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err1141)
	}
	p.Success = v1140
	return err
}

func (p *SystemDropColumnFamilyResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *SystemDropColumnFamilyResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1144 := p.Ire.Read(iprot)
	if err1144 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1144)
	}
	return err
}

func (p *SystemDropColumnFamilyResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemDropColumnFamilyResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Sde = NewSchemaDisagreementException()
	err1147 := p.Sde.Read(iprot)
	if err1147 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.SdeSchemaDisagreementException", err1147)
	}
	return err
}

func (p *SystemDropColumnFamilyResult) ReadFieldSde(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *SystemDropColumnFamilyResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_drop_column_family_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Sde != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropColumnFamilyResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.STRING, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropColumnFamilyResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *SystemDropColumnFamilyResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemDropColumnFamilyResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemDropColumnFamilyResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Sde != nil {
		err = oprot.WriteFieldBegin("sde", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
		err = p.Sde.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SchemaDisagreementException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemDropColumnFamilyResult) WriteFieldSde(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *SystemDropColumnFamilyResult) TStructName() string {
	return "SystemDropColumnFamilyResult"
}

func (p *SystemDropColumnFamilyResult) ThriftName() string {
	return "system_drop_column_family_result"
}

func (p *SystemDropColumnFamilyResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemDropColumnFamilyResult(%+v)", *p)
}

func (p *SystemDropColumnFamilyResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemDropColumnFamilyResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemDropColumnFamilyResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Sde
	}
	return nil
}

func (p *SystemDropColumnFamilyResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRING, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("sde", thrift.STRUCT, 2),
	})
}

/**
 * Attributes:
 *  - KsDef
 */
type SystemAddKeyspaceArgs struct {
	thrift.TStruct
	KsDef *KsDef "ks_def" // 1
}

func NewSystemAddKeyspaceArgs() *SystemAddKeyspaceArgs {
	output := &SystemAddKeyspaceArgs{
		TStruct: thrift.NewTStruct("system_add_keyspace_args", []thrift.TField{
			thrift.NewTField("ks_def", thrift.STRUCT, 1),
		}),
	}
	{
	}
	return output
}

func (p *SystemAddKeyspaceArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "ks_def" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemAddKeyspaceArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.KsDef = NewKsDef()
	err1150 := p.KsDef.Read(iprot)
	if err1150 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.KsDefKsDef", err1150)
	}
	return err
}

func (p *SystemAddKeyspaceArgs) ReadFieldKsDef(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemAddKeyspaceArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_add_keyspace_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemAddKeyspaceArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.KsDef != nil {
		err = oprot.WriteFieldBegin("ks_def", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ks_def", p.ThriftName(), err)
		}
		err = p.KsDef.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("KsDef", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ks_def", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemAddKeyspaceArgs) WriteFieldKsDef(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemAddKeyspaceArgs) TStructName() string {
	return "SystemAddKeyspaceArgs"
}

func (p *SystemAddKeyspaceArgs) ThriftName() string {
	return "system_add_keyspace_args"
}

func (p *SystemAddKeyspaceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemAddKeyspaceArgs(%+v)", *p)
}

func (p *SystemAddKeyspaceArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemAddKeyspaceArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemAddKeyspaceArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.KsDef
	}
	return nil
}

func (p *SystemAddKeyspaceArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("ks_def", thrift.STRUCT, 1),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Sde
 */
type SystemAddKeyspaceResult struct {
	thrift.TStruct
	Success string                       "success" // 0
	Ire     *InvalidRequestException     "ire"     // 1
	Sde     *SchemaDisagreementException "sde"     // 2
}

func NewSystemAddKeyspaceResult() *SystemAddKeyspaceResult {
	output := &SystemAddKeyspaceResult{
		TStruct: thrift.NewTStruct("system_add_keyspace_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRING, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("sde", thrift.STRUCT, 2),
		}),
	}
	{
	}
	return output
}

func (p *SystemAddKeyspaceResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "sde" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemAddKeyspaceResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1151, err1152 := iprot.ReadString()
	if err1152 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err1152)
	}
	p.Success = v1151
	return err
}

func (p *SystemAddKeyspaceResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *SystemAddKeyspaceResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1155 := p.Ire.Read(iprot)
	if err1155 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1155)
	}
	return err
}

func (p *SystemAddKeyspaceResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemAddKeyspaceResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Sde = NewSchemaDisagreementException()
	err1158 := p.Sde.Read(iprot)
	if err1158 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.SdeSchemaDisagreementException", err1158)
	}
	return err
}

func (p *SystemAddKeyspaceResult) ReadFieldSde(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *SystemAddKeyspaceResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_add_keyspace_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Sde != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemAddKeyspaceResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.STRING, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *SystemAddKeyspaceResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *SystemAddKeyspaceResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemAddKeyspaceResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemAddKeyspaceResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Sde != nil {
		err = oprot.WriteFieldBegin("sde", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
		err = p.Sde.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SchemaDisagreementException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemAddKeyspaceResult) WriteFieldSde(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *SystemAddKeyspaceResult) TStructName() string {
	return "SystemAddKeyspaceResult"
}

func (p *SystemAddKeyspaceResult) ThriftName() string {
	return "system_add_keyspace_result"
}

func (p *SystemAddKeyspaceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemAddKeyspaceResult(%+v)", *p)
}

func (p *SystemAddKeyspaceResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemAddKeyspaceResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemAddKeyspaceResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Sde
	}
	return nil
}

func (p *SystemAddKeyspaceResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRING, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("sde", thrift.STRUCT, 2),
	})
}

/**
 * Attributes:
 *  - Keyspace
 */
type SystemDropKeyspaceArgs struct {
	thrift.TStruct
	Keyspace string "keyspace" // 1
}

func NewSystemDropKeyspaceArgs() *SystemDropKeyspaceArgs {
	output := &SystemDropKeyspaceArgs{
		TStruct: thrift.NewTStruct("system_drop_keyspace_args", []thrift.TField{
			thrift.NewTField("keyspace", thrift.STRING, 1),
		}),
	}
	{
	}
	return output
}

func (p *SystemDropKeyspaceArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "keyspace" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropKeyspaceArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1159, err1160 := iprot.ReadString()
	if err1160 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "keyspace", p.ThriftName(), err1160)
	}
	p.Keyspace = v1159
	return err
}

func (p *SystemDropKeyspaceArgs) ReadFieldKeyspace(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemDropKeyspaceArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_drop_keyspace_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropKeyspaceArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("keyspace", thrift.STRING, 1)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Keyspace))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(1, "keyspace", p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropKeyspaceArgs) WriteFieldKeyspace(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemDropKeyspaceArgs) TStructName() string {
	return "SystemDropKeyspaceArgs"
}

func (p *SystemDropKeyspaceArgs) ThriftName() string {
	return "system_drop_keyspace_args"
}

func (p *SystemDropKeyspaceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemDropKeyspaceArgs(%+v)", *p)
}

func (p *SystemDropKeyspaceArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemDropKeyspaceArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemDropKeyspaceArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Keyspace
	}
	return nil
}

func (p *SystemDropKeyspaceArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("keyspace", thrift.STRING, 1),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Sde
 */
type SystemDropKeyspaceResult struct {
	thrift.TStruct
	Success string                       "success" // 0
	Ire     *InvalidRequestException     "ire"     // 1
	Sde     *SchemaDisagreementException "sde"     // 2
}

func NewSystemDropKeyspaceResult() *SystemDropKeyspaceResult {
	output := &SystemDropKeyspaceResult{
		TStruct: thrift.NewTStruct("system_drop_keyspace_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRING, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("sde", thrift.STRUCT, 2),
		}),
	}
	{
	}
	return output
}

func (p *SystemDropKeyspaceResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "sde" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropKeyspaceResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1161, err1162 := iprot.ReadString()
	if err1162 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err1162)
	}
	p.Success = v1161
	return err
}

func (p *SystemDropKeyspaceResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *SystemDropKeyspaceResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1165 := p.Ire.Read(iprot)
	if err1165 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1165)
	}
	return err
}

func (p *SystemDropKeyspaceResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemDropKeyspaceResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Sde = NewSchemaDisagreementException()
	err1168 := p.Sde.Read(iprot)
	if err1168 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.SdeSchemaDisagreementException", err1168)
	}
	return err
}

func (p *SystemDropKeyspaceResult) ReadFieldSde(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *SystemDropKeyspaceResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_drop_keyspace_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Sde != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropKeyspaceResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.STRING, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *SystemDropKeyspaceResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *SystemDropKeyspaceResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemDropKeyspaceResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemDropKeyspaceResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Sde != nil {
		err = oprot.WriteFieldBegin("sde", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
		err = p.Sde.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SchemaDisagreementException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemDropKeyspaceResult) WriteFieldSde(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *SystemDropKeyspaceResult) TStructName() string {
	return "SystemDropKeyspaceResult"
}

func (p *SystemDropKeyspaceResult) ThriftName() string {
	return "system_drop_keyspace_result"
}

func (p *SystemDropKeyspaceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemDropKeyspaceResult(%+v)", *p)
}

func (p *SystemDropKeyspaceResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemDropKeyspaceResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemDropKeyspaceResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Sde
	}
	return nil
}

func (p *SystemDropKeyspaceResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRING, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("sde", thrift.STRUCT, 2),
	})
}

/**
 * Attributes:
 *  - KsDef
 */
type SystemUpdateKeyspaceArgs struct {
	thrift.TStruct
	KsDef *KsDef "ks_def" // 1
}

func NewSystemUpdateKeyspaceArgs() *SystemUpdateKeyspaceArgs {
	output := &SystemUpdateKeyspaceArgs{
		TStruct: thrift.NewTStruct("system_update_keyspace_args", []thrift.TField{
			thrift.NewTField("ks_def", thrift.STRUCT, 1),
		}),
	}
	{
	}
	return output
}

func (p *SystemUpdateKeyspaceArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "ks_def" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemUpdateKeyspaceArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.KsDef = NewKsDef()
	err1171 := p.KsDef.Read(iprot)
	if err1171 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.KsDefKsDef", err1171)
	}
	return err
}

func (p *SystemUpdateKeyspaceArgs) ReadFieldKsDef(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemUpdateKeyspaceArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_update_keyspace_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemUpdateKeyspaceArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.KsDef != nil {
		err = oprot.WriteFieldBegin("ks_def", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ks_def", p.ThriftName(), err)
		}
		err = p.KsDef.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("KsDef", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ks_def", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemUpdateKeyspaceArgs) WriteFieldKsDef(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemUpdateKeyspaceArgs) TStructName() string {
	return "SystemUpdateKeyspaceArgs"
}

func (p *SystemUpdateKeyspaceArgs) ThriftName() string {
	return "system_update_keyspace_args"
}

func (p *SystemUpdateKeyspaceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemUpdateKeyspaceArgs(%+v)", *p)
}

func (p *SystemUpdateKeyspaceArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemUpdateKeyspaceArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemUpdateKeyspaceArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.KsDef
	}
	return nil
}

func (p *SystemUpdateKeyspaceArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("ks_def", thrift.STRUCT, 1),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Sde
 */
type SystemUpdateKeyspaceResult struct {
	thrift.TStruct
	Success string                       "success" // 0
	Ire     *InvalidRequestException     "ire"     // 1
	Sde     *SchemaDisagreementException "sde"     // 2
}

func NewSystemUpdateKeyspaceResult() *SystemUpdateKeyspaceResult {
	output := &SystemUpdateKeyspaceResult{
		TStruct: thrift.NewTStruct("system_update_keyspace_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRING, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("sde", thrift.STRUCT, 2),
		}),
	}
	{
	}
	return output
}

func (p *SystemUpdateKeyspaceResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "sde" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemUpdateKeyspaceResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1172, err1173 := iprot.ReadString()
	if err1173 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err1173)
	}
	p.Success = v1172
	return err
}

func (p *SystemUpdateKeyspaceResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *SystemUpdateKeyspaceResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1176 := p.Ire.Read(iprot)
	if err1176 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1176)
	}
	return err
}

func (p *SystemUpdateKeyspaceResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemUpdateKeyspaceResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Sde = NewSchemaDisagreementException()
	err1179 := p.Sde.Read(iprot)
	if err1179 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.SdeSchemaDisagreementException", err1179)
	}
	return err
}

func (p *SystemUpdateKeyspaceResult) ReadFieldSde(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *SystemUpdateKeyspaceResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_update_keyspace_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Sde != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemUpdateKeyspaceResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.STRING, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *SystemUpdateKeyspaceResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *SystemUpdateKeyspaceResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemUpdateKeyspaceResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemUpdateKeyspaceResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Sde != nil {
		err = oprot.WriteFieldBegin("sde", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
		err = p.Sde.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SchemaDisagreementException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemUpdateKeyspaceResult) WriteFieldSde(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *SystemUpdateKeyspaceResult) TStructName() string {
	return "SystemUpdateKeyspaceResult"
}

func (p *SystemUpdateKeyspaceResult) ThriftName() string {
	return "system_update_keyspace_result"
}

func (p *SystemUpdateKeyspaceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemUpdateKeyspaceResult(%+v)", *p)
}

func (p *SystemUpdateKeyspaceResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemUpdateKeyspaceResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemUpdateKeyspaceResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Sde
	}
	return nil
}

func (p *SystemUpdateKeyspaceResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRING, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("sde", thrift.STRUCT, 2),
	})
}

/**
 * Attributes:
 *  - CfDef
 */
type SystemUpdateColumnFamilyArgs struct {
	thrift.TStruct
	CfDef *CfDef "cf_def" // 1
}

func NewSystemUpdateColumnFamilyArgs() *SystemUpdateColumnFamilyArgs {
	output := &SystemUpdateColumnFamilyArgs{
		TStruct: thrift.NewTStruct("system_update_column_family_args", []thrift.TField{
			thrift.NewTField("cf_def", thrift.STRUCT, 1),
		}),
	}
	{
	}
	return output
}

func (p *SystemUpdateColumnFamilyArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "cf_def" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemUpdateColumnFamilyArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.CfDef = NewCfDef()
	err1182 := p.CfDef.Read(iprot)
	if err1182 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.CfDefCfDef", err1182)
	}
	return err
}

func (p *SystemUpdateColumnFamilyArgs) ReadFieldCfDef(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemUpdateColumnFamilyArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_update_column_family_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemUpdateColumnFamilyArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.CfDef != nil {
		err = oprot.WriteFieldBegin("cf_def", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "cf_def", p.ThriftName(), err)
		}
		err = p.CfDef.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("CfDef", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "cf_def", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemUpdateColumnFamilyArgs) WriteFieldCfDef(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemUpdateColumnFamilyArgs) TStructName() string {
	return "SystemUpdateColumnFamilyArgs"
}

func (p *SystemUpdateColumnFamilyArgs) ThriftName() string {
	return "system_update_column_family_args"
}

func (p *SystemUpdateColumnFamilyArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemUpdateColumnFamilyArgs(%+v)", *p)
}

func (p *SystemUpdateColumnFamilyArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemUpdateColumnFamilyArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemUpdateColumnFamilyArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.CfDef
	}
	return nil
}

func (p *SystemUpdateColumnFamilyArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("cf_def", thrift.STRUCT, 1),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Sde
 */
type SystemUpdateColumnFamilyResult struct {
	thrift.TStruct
	Success string                       "success" // 0
	Ire     *InvalidRequestException     "ire"     // 1
	Sde     *SchemaDisagreementException "sde"     // 2
}

func NewSystemUpdateColumnFamilyResult() *SystemUpdateColumnFamilyResult {
	output := &SystemUpdateColumnFamilyResult{
		TStruct: thrift.NewTStruct("system_update_column_family_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRING, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("sde", thrift.STRUCT, 2),
		}),
	}
	{
	}
	return output
}

func (p *SystemUpdateColumnFamilyResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "sde" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemUpdateColumnFamilyResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1183, err1184 := iprot.ReadString()
	if err1184 != nil {
		return thrift.NewTProtocolExceptionReadField(0, "success", p.ThriftName(), err1184)
	}
	p.Success = v1183
	return err
}

func (p *SystemUpdateColumnFamilyResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *SystemUpdateColumnFamilyResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1187 := p.Ire.Read(iprot)
	if err1187 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1187)
	}
	return err
}

func (p *SystemUpdateColumnFamilyResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *SystemUpdateColumnFamilyResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Sde = NewSchemaDisagreementException()
	err1190 := p.Sde.Read(iprot)
	if err1190 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.SdeSchemaDisagreementException", err1190)
	}
	return err
}

func (p *SystemUpdateColumnFamilyResult) ReadFieldSde(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *SystemUpdateColumnFamilyResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("system_update_column_family_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Sde != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *SystemUpdateColumnFamilyResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteFieldBegin("success", thrift.STRING, 0)
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteString(string(p.Success))
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
	}
	return err
}

func (p *SystemUpdateColumnFamilyResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *SystemUpdateColumnFamilyResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemUpdateColumnFamilyResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *SystemUpdateColumnFamilyResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Sde != nil {
		err = oprot.WriteFieldBegin("sde", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
		err = p.Sde.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SchemaDisagreementException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "sde", p.ThriftName(), err)
		}
	}
	return err
}

func (p *SystemUpdateColumnFamilyResult) WriteFieldSde(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *SystemUpdateColumnFamilyResult) TStructName() string {
	return "SystemUpdateColumnFamilyResult"
}

func (p *SystemUpdateColumnFamilyResult) ThriftName() string {
	return "system_update_column_family_result"
}

func (p *SystemUpdateColumnFamilyResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SystemUpdateColumnFamilyResult(%+v)", *p)
}

func (p *SystemUpdateColumnFamilyResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*SystemUpdateColumnFamilyResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *SystemUpdateColumnFamilyResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Sde
	}
	return nil
}

func (p *SystemUpdateColumnFamilyResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRING, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("sde", thrift.STRUCT, 2),
	})
}

/**
 * Attributes:
 *  - Query
 *  - Compression
 */
type ExecuteCqlQueryArgs struct {
	thrift.TStruct
	Query       []byte      "query"       // 1
	Compression Compression "compression" // 2
}

func NewExecuteCqlQueryArgs() *ExecuteCqlQueryArgs {
	output := &ExecuteCqlQueryArgs{
		TStruct: thrift.NewTStruct("execute_cql_query_args", []thrift.TField{
			thrift.NewTField("query", thrift.BINARY, 1),
			thrift.NewTField("compression", thrift.I32, 2),
		}),
	}
	{
		output.Compression = math.MinInt32 - 1
	}
	return output
}

func (p *ExecuteCqlQueryArgs) IsSetCompression() bool {
	return int64(p.Compression) != math.MinInt32-1
}

func (p *ExecuteCqlQueryArgs) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 1 || fieldName == "query" {
			if fieldTypeId == thrift.STRING {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "compression" {
			if fieldTypeId == thrift.I32 {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *ExecuteCqlQueryArgs) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1191, err1192 := iprot.ReadBinary()
	if err1192 != nil {
		return thrift.NewTProtocolExceptionReadField(1, "query", p.ThriftName(), err1192)
	}
	p.Query = v1191
	return err
}

func (p *ExecuteCqlQueryArgs) ReadFieldQuery(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *ExecuteCqlQueryArgs) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	v1193, err1194 := iprot.ReadI32()
	if err1194 != nil {
		return thrift.NewTProtocolExceptionReadField(2, "compression", p.ThriftName(), err1194)
	}
	p.Compression = Compression(v1193)
	return err
}

func (p *ExecuteCqlQueryArgs) ReadFieldCompression(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *ExecuteCqlQueryArgs) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("execute_cql_query_args")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	err = p.WriteField1(oprot)
	if err != nil {
		return err
	}
	err = p.WriteField2(oprot)
	if err != nil {
		return err
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *ExecuteCqlQueryArgs) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Query != nil {
		err = oprot.WriteFieldBegin("query", thrift.BINARY, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "query", p.ThriftName(), err)
		}
		err = oprot.WriteBinary(p.Query)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "query", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "query", p.ThriftName(), err)
		}
	}
	return err
}

func (p *ExecuteCqlQueryArgs) WriteFieldQuery(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *ExecuteCqlQueryArgs) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.IsSetCompression() {
		err = oprot.WriteFieldBegin("compression", thrift.I32, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "compression", p.ThriftName(), err)
		}
		err = oprot.WriteI32(int32(p.Compression))
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "compression", p.ThriftName(), err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "compression", p.ThriftName(), err)
		}
	}
	return err
}

func (p *ExecuteCqlQueryArgs) WriteFieldCompression(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *ExecuteCqlQueryArgs) TStructName() string {
	return "ExecuteCqlQueryArgs"
}

func (p *ExecuteCqlQueryArgs) ThriftName() string {
	return "execute_cql_query_args"
}

func (p *ExecuteCqlQueryArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ExecuteCqlQueryArgs(%+v)", *p)
}

func (p *ExecuteCqlQueryArgs) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*ExecuteCqlQueryArgs)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *ExecuteCqlQueryArgs) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 1:
		return p.Query
	case 2:
		return p.Compression
	}
	return nil
}

func (p *ExecuteCqlQueryArgs) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("query", thrift.BINARY, 1),
		thrift.NewTField("compression", thrift.I32, 2),
	})
}

/**
 * Attributes:
 *  - Success
 *  - Ire
 *  - Ue
 *  - Te
 *  - Sde
 */
type ExecuteCqlQueryResult struct {
	thrift.TStruct
	Success *CqlResult                   "success" // 0
	Ire     *InvalidRequestException     "ire"     // 1
	Ue      *UnavailableException        "ue"      // 2
	Te      *TimedOutException           "te"      // 3
	Sde     *SchemaDisagreementException "sde"     // 4
}

func NewExecuteCqlQueryResult() *ExecuteCqlQueryResult {
	output := &ExecuteCqlQueryResult{
		TStruct: thrift.NewTStruct("execute_cql_query_result", []thrift.TField{
			thrift.NewTField("success", thrift.STRUCT, 0),
			thrift.NewTField("ire", thrift.STRUCT, 1),
			thrift.NewTField("ue", thrift.STRUCT, 2),
			thrift.NewTField("te", thrift.STRUCT, 3),
			thrift.NewTField("sde", thrift.STRUCT, 4),
		}),
	}
	{
	}
	return output
}

func (p *ExecuteCqlQueryResult) Read(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	_, err = iprot.ReadStructBegin()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	for {
		fieldName, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if fieldId < 0 {
			fieldId = int16(p.FieldIdFromFieldName(fieldName))
		} else if fieldName == "" {
			fieldName = p.FieldNameFromFieldId(int(fieldId))
		}
		if fieldTypeId == thrift.GENERIC {
			fieldTypeId = p.FieldFromFieldId(int(fieldId)).TypeId()
		}
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if fieldId == 0 || fieldName == "success" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField0(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 1 || fieldName == "ire" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField1(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 2 || fieldName == "ue" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField2(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 3 || fieldName == "te" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField3(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else if fieldId == 4 || fieldName == "sde" {
			if fieldTypeId == thrift.STRUCT {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else if fieldTypeId == thrift.VOID {
				err = iprot.Skip(fieldTypeId)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			} else {
				err = p.ReadField4(iprot)
				if err != nil {
					return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
				}
			}
		} else {
			err = iprot.Skip(fieldTypeId)
			if err != nil {
				return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
			}
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionReadField(int(fieldId), fieldName, p.ThriftName(), err)
		}
	}
	err = iprot.ReadStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionReadStruct(p.ThriftName(), err)
	}
	return err
}

func (p *ExecuteCqlQueryResult) ReadField0(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Success = NewCqlResult()
	err1197 := p.Success.Read(iprot)
	if err1197 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.SuccessCqlResult", err1197)
	}
	return err
}

func (p *ExecuteCqlQueryResult) ReadFieldSuccess(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField0(iprot)
}

func (p *ExecuteCqlQueryResult) ReadField1(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ire = NewInvalidRequestException()
	err1200 := p.Ire.Read(iprot)
	if err1200 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.IreInvalidRequestException", err1200)
	}
	return err
}

func (p *ExecuteCqlQueryResult) ReadFieldIre(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField1(iprot)
}

func (p *ExecuteCqlQueryResult) ReadField2(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Ue = NewUnavailableException()
	err1203 := p.Ue.Read(iprot)
	if err1203 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.UeUnavailableException", err1203)
	}
	return err
}

func (p *ExecuteCqlQueryResult) ReadFieldUe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField2(iprot)
}

func (p *ExecuteCqlQueryResult) ReadField3(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Te = NewTimedOutException()
	err1206 := p.Te.Read(iprot)
	if err1206 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.TeTimedOutException", err1206)
	}
	return err
}

func (p *ExecuteCqlQueryResult) ReadFieldTe(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField3(iprot)
}

func (p *ExecuteCqlQueryResult) ReadField4(iprot thrift.TProtocol) (err thrift.TProtocolException) {
	p.Sde = NewSchemaDisagreementException()
	err1209 := p.Sde.Read(iprot)
	if err1209 != nil {
		return thrift.NewTProtocolExceptionReadStruct("p.SdeSchemaDisagreementException", err1209)
	}
	return err
}

func (p *ExecuteCqlQueryResult) ReadFieldSde(iprot thrift.TProtocol) thrift.TProtocolException {
	return p.ReadField4(iprot)
}

func (p *ExecuteCqlQueryResult) Write(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	err = oprot.WriteStructBegin("execute_cql_query_result")
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	switch {
	case p.Sde != nil:
		if err = p.WriteField4(oprot); err != nil {
			return err
		}
	case p.Te != nil:
		if err = p.WriteField3(oprot); err != nil {
			return err
		}
	case p.Ue != nil:
		if err = p.WriteField2(oprot); err != nil {
			return err
		}
	case p.Ire != nil:
		if err = p.WriteField1(oprot); err != nil {
			return err
		}
	default:
		if err = p.WriteField0(oprot); err != nil {
			return err
		}
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteField(-1, "STOP", p.ThriftName(), err)
	}
	err = oprot.WriteStructEnd()
	if err != nil {
		return thrift.NewTProtocolExceptionWriteStruct(p.ThriftName(), err)
	}
	return err
}

func (p *ExecuteCqlQueryResult) WriteField0(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Success != nil {
		err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
		err = p.Success.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("CqlResult", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(0, "success", p.ThriftName(), err)
		}
	}
	return err
}

func (p *ExecuteCqlQueryResult) WriteFieldSuccess(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField0(oprot)
}

func (p *ExecuteCqlQueryResult) WriteField1(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ire != nil {
		err = oprot.WriteFieldBegin("ire", thrift.STRUCT, 1)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
		err = p.Ire.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("InvalidRequestException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(1, "ire", p.ThriftName(), err)
		}
	}
	return err
}

func (p *ExecuteCqlQueryResult) WriteFieldIre(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField1(oprot)
}

func (p *ExecuteCqlQueryResult) WriteField2(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Ue != nil {
		err = oprot.WriteFieldBegin("ue", thrift.STRUCT, 2)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
		err = p.Ue.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("UnavailableException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(2, "ue", p.ThriftName(), err)
		}
	}
	return err
}

func (p *ExecuteCqlQueryResult) WriteFieldUe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField2(oprot)
}

func (p *ExecuteCqlQueryResult) WriteField3(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Te != nil {
		err = oprot.WriteFieldBegin("te", thrift.STRUCT, 3)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
		err = p.Te.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("TimedOutException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(3, "te", p.ThriftName(), err)
		}
	}
	return err
}

func (p *ExecuteCqlQueryResult) WriteFieldTe(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField3(oprot)
}

func (p *ExecuteCqlQueryResult) WriteField4(oprot thrift.TProtocol) (err thrift.TProtocolException) {
	if p.Sde != nil {
		err = oprot.WriteFieldBegin("sde", thrift.STRUCT, 4)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "sde", p.ThriftName(), err)
		}
		err = p.Sde.Write(oprot)
		if err != nil {
			return thrift.NewTProtocolExceptionWriteStruct("SchemaDisagreementException", err)
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return thrift.NewTProtocolExceptionWriteField(4, "sde", p.ThriftName(), err)
		}
	}
	return err
}

func (p *ExecuteCqlQueryResult) WriteFieldSde(oprot thrift.TProtocol) thrift.TProtocolException {
	return p.WriteField4(oprot)
}

func (p *ExecuteCqlQueryResult) TStructName() string {
	return "ExecuteCqlQueryResult"
}

func (p *ExecuteCqlQueryResult) ThriftName() string {
	return "execute_cql_query_result"
}

func (p *ExecuteCqlQueryResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ExecuteCqlQueryResult(%+v)", *p)
}

func (p *ExecuteCqlQueryResult) CompareTo(other interface{}) (int, bool) {
	if other == nil {
		return 1, true
	}
	data, ok := other.(*ExecuteCqlQueryResult)
	if !ok {
		return 0, false
	}
	return thrift.TType(thrift.STRUCT).Compare(p, data)
}

func (p *ExecuteCqlQueryResult) AttributeByFieldId(id int) interface{} {
	switch id {
	default:
		return nil
	case 0:
		return p.Success
	case 1:
		return p.Ire
	case 2:
		return p.Ue
	case 3:
		return p.Te
	case 4:
		return p.Sde
	}
	return nil
}

func (p *ExecuteCqlQueryResult) TStructFields() thrift.TFieldContainer {
	return thrift.NewTFieldContainer([]thrift.TField{
		thrift.NewTField("success", thrift.STRUCT, 0),
		thrift.NewTField("ire", thrift.STRUCT, 1),
		thrift.NewTField("ue", thrift.STRUCT, 2),
		thrift.NewTField("te", thrift.STRUCT, 3),
		thrift.NewTField("sde", thrift.STRUCT, 4),
	})
}
