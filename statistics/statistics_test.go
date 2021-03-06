// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package statistics

import (
	"testing"

	. "github.com/pingcap/check"
	"github.com/pingcap/tidb/ast"
	"github.com/pingcap/tidb/model"
	"github.com/pingcap/tidb/mysql"
	"github.com/pingcap/tidb/sessionctx/variable"
	"github.com/pingcap/tidb/util/mock"
	"github.com/pingcap/tidb/util/types"
)

func TestT(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&testStatisticsSuite{})

type testStatisticsSuite struct {
	count   int64
	samples []types.Datum
	rc      ast.RecordSet
	pk      ast.RecordSet
}

type dataTable struct {
	count   int64
	samples []types.Datum
}

type recordSet struct {
	data   []types.Datum
	count  int64
	cursor int64
}

func (r *recordSet) Fields() ([]*ast.ResultField, error) {
	return nil, nil
}

func (r *recordSet) Next() (*ast.Row, error) {
	if r.cursor == r.count {
		return nil, nil
	}
	r.cursor++
	return &ast.Row{Data: []types.Datum{r.data[r.cursor-1]}}, nil
}

func (r *recordSet) Close() error {
	r.cursor = 0
	return nil
}

func (s *testStatisticsSuite) SetUpSuite(c *C) {
	s.count = 100000
	samples := make([]types.Datum, 10000)
	start := 1000 // 1000 values is null
	for i := start; i < len(samples); i++ {
		samples[i].SetInt64(int64(i))
	}
	for i := start; i < len(samples); i += 3 {
		samples[i].SetInt64(samples[i].GetInt64() + 1)
	}
	for i := start; i < len(samples); i += 5 {
		samples[i].SetInt64(samples[i].GetInt64() + 2)
	}
	sc := new(variable.StatementContext)
	err := types.SortDatums(sc, samples)
	c.Check(err, IsNil)
	s.samples = samples

	rc := &recordSet{
		data:   make([]types.Datum, s.count),
		count:  s.count,
		cursor: 0,
	}
	for i := int64(start); i < rc.count; i++ {
		rc.data[i].SetInt64(int64(i))
	}
	for i := int64(start); i < rc.count; i += 3 {
		rc.data[i].SetInt64(rc.data[i].GetInt64() + 1)
	}
	for i := int64(start); i < rc.count; i += 5 {
		rc.data[i].SetInt64(rc.data[i].GetInt64() + 2)
	}
	err = types.SortDatums(sc, rc.data)
	c.Check(err, IsNil)
	s.rc = rc

	pk := &recordSet{
		data:   make([]types.Datum, s.count),
		count:  s.count,
		cursor: 0,
	}
	for i := int64(0); i < rc.count; i++ {
		pk.data[i].SetInt64(int64(i))
	}
	s.pk = pk
}

func (s *testStatisticsSuite) TestTable(c *C) {
	tblInfo := &model.TableInfo{
		ID: 1,
	}
	columns := []*model.ColumnInfo{
		{
			ID:        2,
			Name:      model.NewCIStr("a"),
			FieldType: *types.NewFieldType(mysql.TypeLonglong),
		},
		{
			ID:        3,
			Name:      model.NewCIStr("b"),
			FieldType: *types.NewFieldType(mysql.TypeLonglong),
		},
		{
			ID:        4,
			Name:      model.NewCIStr("c"),
			FieldType: *types.NewFieldType(mysql.TypeLonglong),
		},
	}
	indices := []*model.IndexInfo{
		{
			ID: 1,
			Columns: []*model.IndexColumn{
				{
					Name:   model.NewCIStr("b"),
					Length: types.UnspecifiedLength,
					Offset: 1,
				},
			},
		},
	}
	tblInfo.Columns = columns
	tblInfo.Indices = indices
	timestamp := int64(10)
	bucketCount := int64(256)
	_, ndv, _ := buildFMSketch(s.rc.(*recordSet).data, 1000)
	builder := &Builder{
		Ctx:           mock.NewContext(),
		TblInfo:       tblInfo,
		StartTS:       timestamp,
		Count:         s.count,
		NumBuckets:    bucketCount,
		ColumnSamples: [][]types.Datum{s.samples},
		ColIDs:        []int64{2},
		ColNDVs:       []int64{ndv},
		IdxRecords:    []ast.RecordSet{s.rc},
		IdxIDs:        []int64{1},
		PkRecords:     ast.RecordSet(s.pk),
		PkID:          4,
	}
	sc := builder.Ctx.GetSessionVars().StmtCtx
	t, err := builder.NewTable()
	c.Check(err, IsNil)

	col := t.Columns[2]
	c.Check(len(col.Buckets), Equals, 232)
	count, err := col.EqualRowCount(sc, types.NewIntDatum(1000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(1))
	count, err = col.LessRowCount(sc, types.NewIntDatum(2000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(19964))
	count, err = col.GreaterRowCount(sc, types.NewIntDatum(2000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(80035))
	count, err = col.LessRowCount(sc, types.NewIntDatum(200000000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(100000))
	count, err = col.GreaterRowCount(sc, types.NewIntDatum(200000000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(0))
	count, err = col.EqualRowCount(sc, types.NewIntDatum(200000000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(0))
	count, err = col.BetweenRowCount(sc, types.NewIntDatum(3000), types.NewIntDatum(3500))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(5079))

	col = t.Columns[3]
	count, err = col.EqualRowCount(sc, types.NewIntDatum(10000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(1))
	count, err = col.LessRowCount(sc, types.NewIntDatum(20000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(19983))
	count, err = col.BetweenRowCount(sc, types.NewIntDatum(30000), types.NewIntDatum(35000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(4618))

	col = t.Columns[4]
	count, err = col.EqualRowCount(sc, types.NewIntDatum(10000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(1))
	count, err = col.LessRowCount(sc, types.NewIntDatum(20000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(20223))
	count, err = col.BetweenRowCount(sc, types.NewIntDatum(30000), types.NewIntDatum(35000))
	c.Check(err, IsNil)
	c.Check(count, Equals, int64(5120))

	str := t.String()
	c.Check(len(str), Greater, 0)

}

func (s *testStatisticsSuite) TestPseudoTable(c *C) {
	ti := &model.TableInfo{}
	colInfo := &model.ColumnInfo{
		ID:        1,
		FieldType: *types.NewFieldType(mysql.TypeLonglong),
	}
	ti.Columns = append(ti.Columns, colInfo)
	tbl := PseudoTable(ti)
	c.Assert(tbl.Count, Greater, int64(0))
	col := tbl.Columns[1]
	c.Assert(col.ID, Greater, int64(0))
	c.Assert(col.NDV, Greater, int64(0))
	sc := new(variable.StatementContext)
	count, err := tbl.ColumnLessRowCount(sc, types.NewIntDatum(100), colInfo)
	c.Assert(err, IsNil)
	c.Assert(count, Equals, int64(3333333))
	count, err = tbl.ColumnEqualRowCount(sc, types.NewIntDatum(1000), colInfo)
	c.Assert(err, IsNil)
	c.Assert(count, Equals, int64(10000))
	count, err = tbl.ColumnBetweenRowCount(sc, types.NewIntDatum(1000), types.NewIntDatum(5000), colInfo)
	c.Assert(err, IsNil)
	c.Assert(count, Equals, int64(250000))
}
