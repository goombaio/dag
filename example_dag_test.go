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
	"fmt"

	"github.com/goombaio/dag"
)

func ExampleDAG_vertices() {

	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)
	vertex4 := dag.NewVertex("4", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}
	err = dag1.AddVertex(vertex3)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}
	err = dag1.AddVertex(vertex4)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}

	fmt.Println(dag1.String())
	// Output:
	// DAG Vertices: 4 - Edges: 0
	// Vertexs:
	// ID: 1 - Parents: 0 - Children: 0 - Value: <nil>
	// ID: 2 - Parents: 0 - Children: 0 - Value: <nil>
	// ID: 3 - Parents: 0 - Children: 0 - Value: <nil>
	// ID: 4 - Parents: 0 - Children: 0 - Value: <nil>
}

func ExampleDAG_edges() {

	dag1 := dag.NewDAG()

	vertex1 := dag.NewVertex("1", nil)
	vertex2 := dag.NewVertex("2", nil)
	vertex3 := dag.NewVertex("3", nil)
	vertex4 := dag.NewVertex("4", nil)

	err := dag1.AddVertex(vertex1)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}
	err = dag1.AddVertex(vertex2)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}
	err = dag1.AddVertex(vertex3)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}
	err = dag1.AddVertex(vertex4)
	if err != nil {
		fmt.Printf("Can't add vertex to DAG: %s", err)
		panic(err)
	}

	// Edges

	err = dag1.AddEdge(vertex1, vertex2)
	if err != nil {
		fmt.Printf("Can't add edge to DAG: %s", err)
		panic(err)
	}

	err = dag1.AddEdge(vertex2, vertex3)
	if err != nil {
		fmt.Printf("Can't add edge to DAG: %s", err)
		panic(err)
	}

	err = dag1.AddEdge(vertex3, vertex4)
	if err != nil {
		fmt.Printf("Can't add edge to DAG: %s", err)
		panic(err)
	}

	fmt.Println(dag1.String())
	// Output:
	// DAG Vertices: 4 - Edges: 3
	// Vertexs:
	// ID: 1 - Parents: 0 - Children: 1 - Value: <nil>
	// ID: 2 - Parents: 1 - Children: 1 - Value: <nil>
	// ID: 3 - Parents: 1 - Children: 1 - Value: <nil>
	// ID: 4 - Parents: 1 - Children: 0 - Value: <nil>
}
