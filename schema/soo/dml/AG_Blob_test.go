// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/928799934/gooxml/schema/soo/dml"
)

func TestAG_BlobConstructor(t *testing.T) {
	v := dml.NewAG_Blob()
	if v == nil {
		t.Errorf("dml.NewAG_Blob must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.AG_Blob should validate: %s", err)
	}
}

func TestAG_BlobMarshalUnmarshal(t *testing.T) {
	v := dml.NewAG_Blob()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewAG_Blob()
	xml.Unmarshal(buf, v2)
}
