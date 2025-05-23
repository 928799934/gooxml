// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/928799934/gooxml/schema/soo/pml"
)

func TestCT_NotesMasterIdListEntryConstructor(t *testing.T) {
	v := pml.NewCT_NotesMasterIdListEntry()
	if v == nil {
		t.Errorf("pml.NewCT_NotesMasterIdListEntry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NotesMasterIdListEntry should validate: %s", err)
	}
}

func TestCT_NotesMasterIdListEntryMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NotesMasterIdListEntry()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NotesMasterIdListEntry()
	xml.Unmarshal(buf, v2)
}
