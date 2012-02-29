package gossie

import (
    "os"
    "thrift"
    "cassandra"
    "time"
)

type Column struct {
    Name []byte
    Value []byte
    Ttl int32
    Timestamp int64
}

type Row struct {
    Key []byte
    Columns []*Column
}

type Slice struct {
    Start []byte
    End []byte
    Count int
    Reversed bool
}

type Range struct {
    Start []byte
    End []byte
    Count int
}

// Read from Cassandra one or more rows.
type Query interface {
    ConsistencyLevel(int) Query
    Cf(string) Query
    Slice(*Slice) Query
    Columns([][]byte) Query
    // Index
    Get([]byte) (*Row, os.Error)
    MultiGet([][]byte) ([]*Row, os.Error)
    //Count([]byte) ([]Row, os.Error)
    //MultiCount([]byte) ([]Row, os.Error)
    RangeGet(*Range) ([]*Row, os.Error)
}

type query struct {
    pool *connectionPool
    consistencyLevel int
    cf string
    slice Slice
    setSlice bool
    columns [][]byte
    setColumns bool
}

func (q *query) ConsistencyLevel(l int) Query {
    q.consistencyLevel = l
    return q
}

func (q *query) Cf(cf string) Query {
    q.cf = cf
    return q
}

func (q *query) Slice(s *Slice) Query {
    q.slice = *s
    q.setSlice = true
    return q
}

func (q *query) Columns(c [][]byte) Query {
    copy(q.columns, c)
    q.setColumns = true
    return q
}

func sliceToCassandra(slice *Slice) *cassandra.SliceRange {
    sr := cassandra.NewSliceRange()
    sr.Start = slice.Start
    sr.Finish = slice.End
    sr.Count = int32(slice.Count)
    sr.Reversed = slice.Reversed
    // workaround some uninitialized slice == nil quirks that trickle down into the generated thrift4go code
    if sr.Start == nil {
        sr.Start = make([]byte, 0)
    }
    if sr.Finish == nil {
        sr.Finish = make([]byte, 0)
    }
    return sr
}

func fullSlice() *cassandra.SliceRange {
    sr := cassandra.NewSliceRange()
    // workaround some uninitialized slice == nil quirks that trickle down into the generated thrift4go code
    sr.Start = make([]byte, 0)
    sr.Finish = make([]byte, 0)
    return sr
}

func (q *query) buildPredicate() *cassandra.SlicePredicate {
    sp := cassandra.NewSlicePredicate()
    if q.setColumns {
        sp.ColumnNames = thrift.NewTList(thrift.BINARY, 1)
        for _, col := range q.columns {
            sp.ColumnNames.Push(col)
        }
    } else if q.setSlice {
        sp.SliceRange = sliceToCassandra(&q.slice)
    } else {
        sp.SliceRange = fullSlice()
    }
    return sp
}

func (q *query) Get(key []byte) (*Row, os.Error) {
    if q.cf == "" {
        return nil, os.NewError("No column family specified")
    }

    cp := cassandra.NewColumnParent()
    cp.ColumnFamily = q.cf

    sp := q.buildPredicate()

    var ret thrift.TList
    err := q.pool.run(func(c *connection) (*cassandra.InvalidRequestException, *cassandra.UnavailableException, *cassandra.TimedOutException, os.Error) {
        var ire *cassandra.InvalidRequestException
        var ue *cassandra.UnavailableException
        var te *cassandra.TimedOutException
        var err os.Error
        ret, ire, ue, te, err = c.client.GetSlice(key, cp, sp, cassandra.ConsistencyLevel(q.consistencyLevel))
        return ire, ue, te, err
    })

    if err != nil {
        return nil, err
    }

    return rowFromTListColumns(key, ret), nil
}

func (q *query) MultiGet(keys [][]byte) ([]*Row, os.Error) {
    if q.cf == "" {
        return nil, os.NewError("No column family specified")
    }

    if len(keys) <= 0 {
        return nil, nil
    }

    cp := cassandra.NewColumnParent()
    cp.ColumnFamily = q.cf

    sp := q.buildPredicate()

    tkeys := thrift.NewTList(thrift.BINARY, 1)
    for _, k := range keys {
        tkeys.Push(k)
    }

    var ret thrift.TMap
    err := q.pool.run(func(c *connection) (*cassandra.InvalidRequestException, *cassandra.UnavailableException, *cassandra.TimedOutException, os.Error) {
        var ire *cassandra.InvalidRequestException
        var ue *cassandra.UnavailableException
        var te *cassandra.TimedOutException
        var err os.Error
        ret, ire, ue, te, err = c.client.MultigetSlice(tkeys, cp, sp, cassandra.ConsistencyLevel(q.consistencyLevel))
        return ire, ue, te, err
    })

    if err != nil {
        return nil, err
    }

    return rowsFromTMap(ret), nil
}

func (q *query) RangeGet(rang *Range) ([]*Row, os.Error) {
    if q.cf == "" {
        return nil, os.NewError("No column family specified")
    }

    if rang == nil || rang.Count <= 0 {
        return nil, nil
    }

    kr := cassandra.NewKeyRange()
    kr.StartKey = rang.Start
    kr.EndKey = rang.End
    kr.Count = int32(rang.Count)
    // workaround some uninitialized slice == nil quirks that trickle down into the generated thrift4go code
    if kr.StartKey == nil {
        kr.StartKey = make([]byte, 0)
    }
    if kr.EndKey == nil {
        kr.EndKey = make([]byte, 0)
    }

    cp := cassandra.NewColumnParent()
    cp.ColumnFamily = q.cf

    sp := q.buildPredicate()

    var ret thrift.TList
    err := q.pool.run(func(c *connection) (*cassandra.InvalidRequestException, *cassandra.UnavailableException, *cassandra.TimedOutException, os.Error) {
        var ire *cassandra.InvalidRequestException
        var ue *cassandra.UnavailableException
        var te *cassandra.TimedOutException
        var err os.Error
        ret, ire, ue, te, err = c.client.GetRangeSlices(cp, sp, kr, cassandra.ConsistencyLevel(q.consistencyLevel))
        return ire, ue, te, err
    })

    if err != nil {
        return nil, err
    }

    return rowsFromTListKeySlice(ret), nil
}

func rowFromTListColumns(key []byte, tl thrift.TList) *Row {
    if tl == nil {
        return nil
    }
    if tl.Len() <= 0 {
        return nil
    }
    r := &Row{Key:key}
    for colI := range tl.Iter() {
        var col *cassandra.ColumnOrSuperColumn = colI.(*cassandra.ColumnOrSuperColumn)
        if col.Column != nil {
            c := &Column{
                Name:col.Column.Name,
                Value:col.Column.Value,
                Timestamp:col.Column.Timestamp,
                Ttl:col.Column.Ttl,
            }
            r.Columns = append(r.Columns, c)
        } else if col.CounterColumn != nil {
            v, _ := Marshal(col.CounterColumn.Value, LongType)
            c := &Column{
                Name:col.CounterColumn.Name,
                Value:v,
            }
            r.Columns = append(r.Columns, c)
        }
    }
    return r
}

func rowsFromTMap(tm thrift.TMap) []*Row {
    if tm != nil && tm.Len() > 0 {
        var r []*Row
        for rowI := range tm.Iter() {
            // workaround some issues with the way the key->row array gets built by thrift4go
            rawKey := rowI.Key()
            var key []byte
            switch k := rawKey.(type) {
                case []uint8: key = []byte(k)
                case string: key = []byte(k)
            }
            columns := (rowI.Value()).(thrift.TList)
            row := rowFromTListColumns(key, columns)
            if row != nil {
                r = append(r, row)
            }
        }
        return r
    }
    return nil
}

func rowsFromTListKeySlice(tl thrift.TList) []*Row {
    if tl != nil && tl.Len() > 0 {
        var r []*Row
        for keySliceI := range tl.Iter() {
            keySlice := keySliceI.(*cassandra.KeySlice)
            key := keySlice.Key
            row := rowFromTListColumns(key, keySlice.Columns)
            if row != nil {
                r = append(r, row)
            }
        }
        return r
    }
    return nil
}

type Mutation interface {
    Insert(cf string, row *Row) Mutation
    DeltaCounters(cf string, row *Row) Mutation
    Delete(cf string, key []byte) Mutation
    DeleteColumns(cf string, key []byte, columns [][]byte) Mutation
    //DeleteSlice(cf string, key []byte, slice *Slice) Mutation
    Run() os.Error
}

type mutation struct {
    pool *connectionPool
    consistencyLevel int
    mutations thrift.TMap
}

func makeMutation(cp *connectionPool, cl int) *mutation {
    return &mutation {
        pool: cp,
        consistencyLevel: cl,
        mutations: thrift.NewTMap(thrift.BINARY, thrift.MAP, 1),
    }
}

func now() int64 {
    return time.Nanoseconds()/1000
}

func (m *mutation) addMutation(cf string, key []byte) *cassandra.Mutation {
    tm := cassandra.NewMutation()
    var cfMuts thrift.TMap
    im, exists := m.mutations.Get(key)
    if !exists {
        cfMuts = thrift.NewTMap(thrift.STRING, thrift.LIST, 1)
        m.mutations.Set(key, cfMuts)
    } else {
        cfMuts = im.(thrift.TMap)
    }
    var mutList thrift.TList
    im, exists = cfMuts.Get(cf)
    if !exists {
        mutList = thrift.NewTList(thrift.STRUCT, 1)
        cfMuts.Set(cf, mutList)
    } else {
        mutList = im.(thrift.TList)
    }
    mutList.Push(tm)
    return tm
}

func (m *mutation) Insert(cf string, row *Row) Mutation {
    t := now()
    for _, col := range row.Columns {
        tm := m.addMutation(cf, row.Key)
        c := cassandra.NewColumn()
        c.Name = col.Name
        c.Value = col.Value
        c.Ttl = col.Ttl
        if col.Timestamp > 0 {
            c.Timestamp = col.Timestamp
        } else {
            c.Timestamp = t
        }
        cs := cassandra.NewColumnOrSuperColumn()
        cs.Column = c
        tm.ColumnOrSupercolumn = cs
    }
    return m
}

func (m *mutation) DeltaCounters(cf string, row *Row) Mutation {
    for _, col := range row.Columns {
        tm := m.addMutation(cf, row.Key)
        c := cassandra.NewCounterColumn()
        c.Name = col.Name
        Unmarshal(col.Value, LongType, &c.Value)
        cs := cassandra.NewColumnOrSuperColumn()
        cs.CounterColumn = c
        tm.ColumnOrSupercolumn = cs
    }
    return m
}

func (m *mutation) Delete(cf string, key []byte) Mutation {
    tm := m.addMutation(cf, key)
    d := cassandra.NewDeletion()
    d.Timestamp = now()
    tm.Deletion = d
    return m
}

func (m *mutation) DeleteColumns(cf string, key []byte, columns [][]byte) Mutation {
    tm := m.addMutation(cf, key)
    d := cassandra.NewDeletion()
    d.Timestamp = now()
    sp := cassandra.NewSlicePredicate()
    sp.ColumnNames = thrift.NewTList(thrift.BINARY, 1)
    for _, name := range columns {
        sp.ColumnNames.Push(name)
    }
    d.Predicate = sp
    tm.Deletion = d
    return m
}

/* InvalidRequestException({TStruct:InvalidRequestException Why:Deletion does not yet support SliceRange predicates.})
func (m *mutation) DeleteSlice(cf string, key []byte, slice *Slice) Mutation {
    tm := m.addMutation(cf, key)
    d := cassandra.NewDeletion()
    d.Timestamp = now()
    sp := cassandra.NewSlicePredicate()
    sp.SliceRange = sliceToCassandra(slice)
    d.Predicate = sp
    tm.Deletion = d
    return m
}
*/

func (m *mutation) Run() os.Error {
    return m.pool.run(func(c *connection) (*cassandra.InvalidRequestException, *cassandra.UnavailableException, *cassandra.TimedOutException, os.Error) {
        var ire *cassandra.InvalidRequestException
        var ue *cassandra.UnavailableException
        var te *cassandra.TimedOutException
        var err os.Error
        ire, ue, te, err = c.client.BatchMutate(m.mutations, cassandra.ConsistencyLevel(m.consistencyLevel))
        return ire, ue, te, err
    })
}