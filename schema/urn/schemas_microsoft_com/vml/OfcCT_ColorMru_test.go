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

func TestOfcCT_ColorMruConstructor(t *testing.T) {
	v := vml.NewOfcCT_ColorMru()
	if v == nil {
		t.Errorf("vml.NewOfcCT_ColorMru must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_ColorMru should validate: %s", err)
	}
}

func TestOfcCT_ColorMruMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_ColorMru()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_ColorMru()
	xml.Unmarshal(buf, v2)
}
