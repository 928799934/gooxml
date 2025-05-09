// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/928799934/gooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_ShapetypeConstructor(t *testing.T) {
	v := vml.NewCT_Shapetype()
	if v == nil {
		t.Errorf("vml.NewCT_Shapetype must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Shapetype should validate: %s", err)
	}
}

func TestCT_ShapetypeMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Shapetype()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Shapetype()
	xml.Unmarshal(buf, v2)
}
