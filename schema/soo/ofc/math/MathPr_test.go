// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/928799934/gooxml/schema/soo/ofc/math"
)

func TestMathPrConstructor(t *testing.T) {
	v := math.NewMathPr()
	if v == nil {
		t.Errorf("math.NewMathPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.MathPr should validate: %s", err)
	}
}

func TestMathPrMarshalUnmarshal(t *testing.T) {
	v := math.NewMathPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewMathPr()
	xml.Unmarshal(buf, v2)
}
