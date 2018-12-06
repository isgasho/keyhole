// Copyright 2018 Kuei-chun Chen. All rights reserved.

package ftdc

import (
	"io/ioutil"
	"testing"

	"github.com/globalsign/mgo/bson"
)

var filename = "../test_data/diagnostic.data/metrics.2017-10-12T20-08-53Z-00000"

func TestDecode(t *testing.T) {
	var err error
	var buffer []byte

	if buffer, err = ioutil.ReadFile(filename); err != nil {
		t.Fatal(err)
	}
	m := NewMetrics()
	m.ReadAllMetrics(buffer)
	if len(m.Blocks) == 0 {
		t.Fatal()
	}

	if _, err = m.decode([]byte(m.Blocks[0])); err != nil {
		t.Fatal(err)
	}
}

func TestTraverseDocElem(t *testing.T) {
	var err error
	var buffer []byte

	if buffer, err = ioutil.ReadFile(filename); err != nil {
		t.Fatal(err)
	}
	m := NewMetrics()
	m.ReadAllMetrics(buffer)
	if len(m.Blocks) == 0 {
		t.Fatal()
	}

	var dp = MetricsData{DataPointsMap: map[string][]int64{}}
	var docElem = bson.D{}
	var attribsList = []string{}
	bson.Unmarshal([]byte(m.Blocks[0]), &docElem) // first document
	traverseDocElem(&attribsList, &dp.DataPointsMap, docElem, "")
	if len(attribsList) == 0 || len(dp.DataPointsMap) != len(attribsList) {
		t.Fatal()
	}
}