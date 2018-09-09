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

package dag

import (
	"fmt"
	"sync"

	"github.com/goombaio/orderedmap"
)

// DAG type implements a Directed acyclic graph data structure.
type DAG struct {
	mu       sync.Mutex
	Vertices orderedmap.OrderedMap
}

// NewDAG creates a new directed acyclic graph instance.
func NewDAG() *DAG {
	d := &DAG{
		Vertices: *orderedmap.NewOrderedMap(),
	}

	return d
}

// AddVertex adds a vertex to the graph.
func (d *DAG) AddVertex(v *Vertex) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.Vertices.Put(v.ID, v)

	return nil
}

// DeleteVertex deletes a verrtex and all the edges referencing it from the
// graph.
func (d *DAG) DeleteVertex(vertex *Vertex) error {
	existsVertex := false

	d.mu.Lock()
	defer d.mu.Unlock()

	// Check if vertexs exists.
	for _, v := range d.Vertices.Values() {
		if v == vertex {
			existsVertex = true
		}
	}
	if existsVertex == false {
		return fmt.Errorf("Vertex with ID %v not found", vertex.ID)
	}

	d.Vertices.Remove(vertex.ID)

	return nil
}

// AddEdge adds a directed edge between two existing vertices to the graph.
func (d *DAG) AddEdge(tailVertex *Vertex, headVertex *Vertex) error {
	tailExists := false
	headExists := false

	d.mu.Lock()
	defer d.mu.Unlock()

	// Check if vertexs exists.
	for _, vertex := range d.Vertices.Values() {
		if vertex == tailVertex {
			tailExists = true
		}
		if vertex == headVertex {
			headExists = true
		}
	}
	if tailExists == false {
		return fmt.Errorf("Vertex with ID %v not found", tailVertex.ID)
	}
	if headExists == false {
		return fmt.Errorf("Vertex with ID %v not found", headVertex.ID)
	}

	// Check if edge already exists.
	for _, childVertex := range tailVertex.Children.Values() {
		if childVertex == headVertex {
			return fmt.Errorf("Edge (%v,%v) already exists", tailVertex.ID, headVertex.ID)
		}
	}

	// Add edge.
	tailVertex.Children.Add(headVertex)
	headVertex.Parents.Add(tailVertex)

	return nil
}

// DeleteEdge deletes a directed edge between two existing vertices from the
// graph.
func (d *DAG) DeleteEdge(tailVertex *Vertex, headVertex *Vertex) error {
	for _, childVertex := range tailVertex.Children.Values() {
		if childVertex == headVertex {
			tailVertex.Children.Remove(childVertex)
		}
	}

	return nil
}

// Order return the number of vertices in the graph.
func (d *DAG) Order() int {
	numVertices := d.Vertices.Size()

	return numVertices
}

// Size return the number of edges in the graph.
func (d *DAG) Size() int {
	numEdges := 0
	for _, vertex := range d.Vertices.Values() {
		numEdges = numEdges + vertex.(*Vertex).Children.Size()
	}

	return numEdges
}

// SinkVertices return vertices with no children defined by the graph edges.
func (d *DAG) SinkVertices() []*Vertex {
	var sinkVertices []*Vertex

	for _, vertex := range d.Vertices.Values() {
		if vertex.(*Vertex).Children.Size() == 0 {
			sinkVertices = append(sinkVertices, vertex.(*Vertex))
		}
	}

	return sinkVertices
}

// Validate return a boolean value whether DAG is valid or not.
// A DAG is valid if all edges in the graph point to existing vertices, and
// that there are no dependency cycles.
func (d *DAG) Validate() bool {
	// If there are no vertices
	if d.Order() == 0 {
		// TODO:
		// Sure? A DAG without vertexs and edges is valid?
		return true
	}

	return false
}

// String implements stringer interface and prints an string representation
// of this instance.
func (d *DAG) String() string {
	var result string
	result = fmt.Sprintf("DAG Vertices: %d - Edges: %d\n", d.Order(), d.Size())
	result = result + fmt.Sprintf("Vertexs:\n")
	for _, vertex := range d.Vertices.Values() {
		vertex = vertex.(*Vertex)

		result = result + fmt.Sprintf("%s", vertex)
	}

	return result
}
