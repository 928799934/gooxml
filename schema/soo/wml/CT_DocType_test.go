// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/928799934/gooxml/schema/soo/wml"
)

func TestCT_DocTypeConstructor(t *testing.T) {
	v := wml.NewCT_DocType()
	if v == nil {
		t.Errorf("wml.NewCT_DocType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocType should validate: %s", err)
	}
}

func TestCT_DocTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocType()
	xml.Unmarshal(buf, v2)
}
