// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package dag_test

import (
	"testing"

	"github.com/goombaio/dag"
)

func TestDAG(t *testing.T) {
	d := dag.NewDAG()

	if d.Vertices.Size() != 0 {
		t.Fatalf("DAG number of vertices expected to be 0 but got %d", d.Vertices.Size())
	}
}

func TestDAG_AddVertex(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	if dag1.Vertices.Size() != 1 {
		t.Fatalf("DAG number of vertices expected to be 1 but got %d", dag1.Vertices.Size())
	}
}

func TestDAG_DeleteVertex(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG")
	}

	if dag1.Vertices.Size() != 1 {
		t.Fatalf("DAG number of vertices expected to be 1 but got %d", dag1.Vertices.Size())
	}

	err = dag1.DeleteVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't delete vertex from DAG: %s", err)
	}

	if dag1.Vertices.Size() != 0 {
		t.Fatalf("DAG number of vertices expected to be 0 but got %d", dag1.Vertices.Size())
	}

	err = dag1.DeleteVertex(vertex1)
	if err == nil {
		t.Fatalf("Vertex don't exist, AddEdge should fail but it doesn't: %s", err)
	}
}

func TestDAG_AddEdge(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", "two")

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}
}

func TestDAG_AddEdge_FailsVertextDontExist(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex3, vertex2)
	if err == nil {
		t.Fatalf("Vertex don't exist, AddEdge should fail but it doesn't")
	}

	err = dag1.AddEdge(vertex2, vertex3)
	if err == nil {
		t.Fatalf("Vertex don't exist, AddEdge should fail but it doesn't")
	}
}

func TestDAG_AddEdge_FailsAlreadyExists(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err == nil {
		t.Fatalf("Edge already exists, AddEdge should fail but it doesn't")
	}
}

func TestDAG_DeleteEdge(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG")
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG")
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG")
	}

	size := dag1.Size()
	if size != 1 {
		t.Fatalf("Dag expected to have 1 edge but got %d", size)
	}

	err = dag1.DeleteEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't delete edge from DAG")
	}

	size = dag1.Size()
	if size != 0 {
		t.Fatalf("Dag expected to have 0 edges but got %d", size)
	}
}

func TestDAG_Order(t *testing.T) {
	dag1 := dag.NewDAG()

	expected_order := 0
	order := dag1.Order()
	if order != expected_order {
		t.Fatalf("Expected order to be %d but got %d", expected_order, order)
	}

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex3)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	expected_order = 3
	order = dag1.Order()
	if order != expected_order {
		t.Fatalf("Expected order to be %d but got %d", expected_order, order)
	}
}

func TestDAG_Size(t *testing.T) {
	dag1 := dag.NewDAG()

	expected_size := 0
	size := dag1.Size()
	if size != expected_size {
		t.Fatalf("Expected size to be %d but got %d", expected_size, size)
	}

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)
	vertex4 := dag.NewVertex("4", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex3)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex4)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	expected_size = 0
	size = dag1.Size()
	if size != expected_size {
		t.Fatalf("Expected size to be %d but got %d", expected_size, size)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex2, vertex3)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex2, vertex4)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	expected_size = 3
	size = dag1.Size()
	if size != expected_size {
		t.Fatalf("Expected size to be %d but got %d", expected_size, size)
	}
}

func TestDAG_SinkVertices(t *testing.T) {
	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		t.Fatalf("Can't add vertex to DAG: %s", err)
	}

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		t.Fatalf("Can't add edge to DAG: %s", err)
	}

	sinkVertices := dag1.SinkVertices()
	if len(sinkVertices) != 1 {
		t.Fatalf("Expected to have 1 Sink vertex but got %d", len(sinkVertices))
	}
}

func TestDAG_Validate(t *testing.T) {
	dag1 := dag.NewDAG()

	valid := dag1.Validate()
	if !valid {
		t.Fatalf("Expected to be valid, but it does not.")
	}
}
