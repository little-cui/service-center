/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package util

import (
	"fmt"
	"reflect"
	"testing"
)

type testStru struct {
	f1 int
	f2 string
	f3 testField
	f4 *testField
}

type testField struct {
}

func TestLoadStruct(t *testing.T) {
	obj1 := testStru{}
	v := ReflectObject(obj1)
	if v.Type.String() != "util.testStru" {
		t.Fatalf("TestLoadStruct failed, %s != 'testStru'", v.Type.String())
	}
	if len(v.Fields) != 4 {
		t.Fatalf("TestLoadStruct failed, wrong count of fields")
	}
	for _, f := range v.Fields {
		fmt.Println(f.Name, f.Type.String())
	}

	obj2 := testStru{}
	v = ReflectObject(obj2)
	v = ReflectObject(&obj2)
	v = ReflectObject(&obj2)
}

func BenchmarkLoadStruct(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ReflectObject(testStru{})
		}
	})
	b.ReportAllocs()
	// 20000000	        86.9 ns/op	      32 B/op	       1 allocs/op
}

var (
	sliceSize  = uint64(reflect.TypeOf(reflect.SliceHeader{}).Size())
	stringSize = uint64(reflect.TypeOf(reflect.StringHeader{}).Size())
)

type S struct {
	a  int
	s  string
	p  *S
	m  map[int32]uint32
	u  []uint64
	ua [8]uint64
	ch chan int
	i  interface{}
}

func BenchmarkSizeof(b *testing.B) {
	s := &S{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Sizeof(S{
				p: s,
				i: s,
			})
		}
	})
	b.ReportAllocs()
	// 2000000	       650 ns/op	     160 B/op	       1 allocs/op
}
