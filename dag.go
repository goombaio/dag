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

	"github.com/google/uuid"
)

// DAG type implements a Directed acyclic graph data structure.
// https://en.wikipedia.org/wiki/Directed_acyclic_graph.
type DAG struct {
	mu       sync.Mutex
	Vertices map[uuid.UUID]*Vertex
}

// NewDAG creates a new directed acyclic graph instance.
func NewDAG() *DAG {
	d := &DAG{
		Vertices: make(map[uuid.UUID]*Vertex, 0),
	}

	return d
}

// AddVertex adds a vertex to the graph.
func (d *DAG) AddVertex(v *Vertex) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.Vertices[v.ID] = v

	return nil
}

// DeleteVertex deletes a verrtex and all the edges referencing it from the
// graph.
func (d *DAG) DeleteVertex(vertex *Vertex) error {
	existsVertex := false

	d.mu.Lock()
	defer d.mu.Unlock()

	// Check if vertexs exists.
	for _, v := range d.Vertices {
		if v == vertex {
			existsVertex = true
		}
	}
	if existsVertex == false {
		return fmt.Errorf("Vertex with ID %v not found", vertex.ID)
	}

	delete(d.Vertices, vertex.ID)

	return nil
}

// AddEdge adds a directed edge between two existing vertices to the graph.
func (d *DAG) AddEdge(tailVertex *Vertex, headVertex *Vertex) error {
	tailExists := false
	toExists := false

	d.mu.Lock()
	defer d.mu.Unlock()

	// Check if vertexs exists.
	for _, vertex := range d.Vertices {
		if vertex == tailVertex {
			tailExists = true
		}
		if vertex == headVertex {
			toExists = true
		}
	}
	if tailExists == false {
		return fmt.Errorf("Vertex with ID %v not found", tailVertex.ID)
	}
	if toExists == false {
		return fmt.Errorf("Vertex with ID %v not found", headVertex.ID)
	}

	// Check if edge already exists.
	for _, childVertex := range tailVertex.Children {
		if childVertex == headVertex {
			return fmt.Errorf("Edge (%v,%v) already exists", tailVertex.ID, headVertex.ID)
		}
	}

	// Add edge.
	tailVertex.Children = append(tailVertex.Children, headVertex)

	return nil
}

// DeleteEdge deletes a directed edge between two existing vertices from the
// graph.
func (d *DAG) DeleteEdge(tailVertex *Vertex, headVertex *Vertex) error {
	for i, childVertex := range tailVertex.Children {
		if childVertex == headVertex {
			tailVertex.Children[i] = tailVertex.Children[len(tailVertex.Children)-1]
			tailVertex.Children = tailVertex.Children[:len(tailVertex.Children)-1]
		}
	}

	return nil
}

// Order return the number of vertices in the graph.
func (d *DAG) Order() int {
	numVertices := len(d.Vertices)

	return numVertices
}

// Size return the number of edges in the graph.
func (d *DAG) Size() int {
	numEdges := 0
	for _, vertex := range d.Vertices {
		numEdges = numEdges + len(vertex.Children)
	}

	return numEdges
}

// SinkVertices return vertices with no children defined by the graph edges.
func (d *DAG) SinkVertices() []*Vertex {
	var sinkVertices []*Vertex

	for _, vertex := range d.Vertices {
		if len(vertex.Children) == 0 {
			sinkVertices = append(sinkVertices, vertex)
		}
	}

	return sinkVertices
}
