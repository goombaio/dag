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

func TestVertex(t *testing.T) {
	v := dag.NewVertex("1", nil)

	if v.ID == "" {
		t.Fatalf("Vertex ID expected to be not empty string.\n")
	}
	if v.Value != nil {
		t.Fatalf("Vertex Value expected to be nil.\n")
	}
}

func TestVertex_Parents(t *testing.T) {
	v := dag.NewVertex("1", nil)

	numParents := v.Parents.Size()
	if numParents != 0 {
		t.Fatalf("Vertex Parents expected to be 0 but got %d", v.Parents.Size())
	}
}

func TestVertex_Children(t *testing.T) {
	v := dag.NewVertex("1", nil)

	numParents := v.Children.Size()
	if numParents != 0 {
		t.Fatalf("Vertex Children expected to be 0 but got %d", v.Children.Size())
	}
}

func TestVertex_Degree(t *testing.T) {
	v := dag.NewVertex("1", nil)

	degree := v.Degree()
	if degree != 0 {
		t.Fatalf("Vertex Degree expected to be 0 but got %d", v.Degree())
	}
}

func TestVertex_String(t *testing.T) {
	v := dag.NewVertex("1", nil)
	vstr := v.String()

	expected := "ID: 1 - Parents: 0 - Children: 0 - Value: <nil>\n"
	if vstr != expected {
		t.Fatalf("Vertex stringer expected to be %q but got %q\n", expected, vstr)
	}
}

func TestVertex_String_WithStringValue(t *testing.T) {
	v := dag.NewVertex("1", "one")
	vstr := v.String()

	expected := "ID: 1 - Parents: 0 - Children: 0 - Value: one\n"
	if vstr != expected {
		t.Fatalf("Vertex stringer expected to be %q but got %q\n", expected, vstr)
	}
}

func TestVertex_WithStringValue(t *testing.T) {
	v := dag.NewVertex("1", "one")

	if v.ID == "" {
		t.Fatalf("Vertex ID expected to be not empty string.\n")
	}
	if v.Value != "one" {
		t.Fatalf("Vertex Value expected to be one.\n")
	}
}
