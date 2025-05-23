// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes

import (
	"encoding/xml"
	"fmt"
	"regexp"

	"github.com/928799934/gooxml"
)

const ST_CyPattern = `\s*[0-9]*\.[0-9]{4}\s*`

var ST_CyPatternRe = regexp.MustCompile(ST_CyPattern)

const ST_ErrorPattern = `\s*0x[0-9A-Za-z]{8}\s*`

var ST_ErrorPatternRe = regexp.MustCompile(ST_ErrorPattern)

type ST_VectorBaseType byte

const (
	ST_VectorBaseTypeUnset    ST_VectorBaseType = 0
	ST_VectorBaseTypeVariant  ST_VectorBaseType = 1
	ST_VectorBaseTypeI1       ST_VectorBaseType = 2
	ST_VectorBaseTypeI2       ST_VectorBaseType = 3
	ST_VectorBaseTypeI4       ST_VectorBaseType = 4
	ST_VectorBaseTypeI8       ST_VectorBaseType = 5
	ST_VectorBaseTypeUi1      ST_VectorBaseType = 6
	ST_VectorBaseTypeUi2      ST_VectorBaseType = 7
	ST_VectorBaseTypeUi4      ST_VectorBaseType = 8
	ST_VectorBaseTypeUi8      ST_VectorBaseType = 9
	ST_VectorBaseTypeR4       ST_VectorBaseType = 10
	ST_VectorBaseTypeR8       ST_VectorBaseType = 11
	ST_VectorBaseTypeLpstr    ST_VectorBaseType = 12
	ST_VectorBaseTypeLpwstr   ST_VectorBaseType = 13
	ST_VectorBaseTypeBstr     ST_VectorBaseType = 14
	ST_VectorBaseTypeDate     ST_VectorBaseType = 15
	ST_VectorBaseTypeFiletime ST_VectorBaseType = 16
	ST_VectorBaseTypeBool     ST_VectorBaseType = 17
	ST_VectorBaseTypeCy       ST_VectorBaseType = 18
	ST_VectorBaseTypeError    ST_VectorBaseType = 19
	ST_VectorBaseTypeClsid    ST_VectorBaseType = 20
)

func (e ST_VectorBaseType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VectorBaseTypeUnset:
		attr.Value = ""
	case ST_VectorBaseTypeVariant:
		attr.Value = "variant"
	case ST_VectorBaseTypeI1:
		attr.Value = "i1"
	case ST_VectorBaseTypeI2:
		attr.Value = "i2"
	case ST_VectorBaseTypeI4:
		attr.Value = "i4"
	case ST_VectorBaseTypeI8:
		attr.Value = "i8"
	case ST_VectorBaseTypeUi1:
		attr.Value = "ui1"
	case ST_VectorBaseTypeUi2:
		attr.Value = "ui2"
	case ST_VectorBaseTypeUi4:
		attr.Value = "ui4"
	case ST_VectorBaseTypeUi8:
		attr.Value = "ui8"
	case ST_VectorBaseTypeR4:
		attr.Value = "r4"
	case ST_VectorBaseTypeR8:
		attr.Value = "r8"
	case ST_VectorBaseTypeLpstr:
		attr.Value = "lpstr"
	case ST_VectorBaseTypeLpwstr:
		attr.Value = "lpwstr"
	case ST_VectorBaseTypeBstr:
		attr.Value = "bstr"
	case ST_VectorBaseTypeDate:
		attr.Value = "date"
	case ST_VectorBaseTypeFiletime:
		attr.Value = "filetime"
	case ST_VectorBaseTypeBool:
		attr.Value = "bool"
	case ST_VectorBaseTypeCy:
		attr.Value = "cy"
	case ST_VectorBaseTypeError:
		attr.Value = "error"
	case ST_VectorBaseTypeClsid:
		attr.Value = "clsid"
	}
	return attr, nil
}

func (e *ST_VectorBaseType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "variant":
		*e = 1
	case "i1":
		*e = 2
	case "i2":
		*e = 3
	case "i4":
		*e = 4
	case "i8":
		*e = 5
	case "ui1":
		*e = 6
	case "ui2":
		*e = 7
	case "ui4":
		*e = 8
	case "ui8":
		*e = 9
	case "r4":
		*e = 10
	case "r8":
		*e = 11
	case "lpstr":
		*e = 12
	case "lpwstr":
		*e = 13
	case "bstr":
		*e = 14
	case "date":
		*e = 15
	case "filetime":
		*e = 16
	case "bool":
		*e = 17
	case "cy":
		*e = 18
	case "error":
		*e = 19
	case "clsid":
		*e = 20
	}
	return nil
}

func (m ST_VectorBaseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_VectorBaseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "variant":
			*m = 1
		case "i1":
			*m = 2
		case "i2":
			*m = 3
		case "i4":
			*m = 4
		case "i8":
			*m = 5
		case "ui1":
			*m = 6
		case "ui2":
			*m = 7
		case "ui4":
			*m = 8
		case "ui8":
			*m = 9
		case "r4":
			*m = 10
		case "r8":
			*m = 11
		case "lpstr":
			*m = 12
		case "lpwstr":
			*m = 13
		case "bstr":
			*m = 14
		case "date":
			*m = 15
		case "filetime":
			*m = 16
		case "bool":
			*m = 17
		case "cy":
			*m = 18
		case "error":
			*m = 19
		case "clsid":
			*m = 20
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_VectorBaseType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "variant"
	case 2:
		return "i1"
	case 3:
		return "i2"
	case 4:
		return "i4"
	case 5:
		return "i8"
	case 6:
		return "ui1"
	case 7:
		return "ui2"
	case 8:
		return "ui4"
	case 9:
		return "ui8"
	case 10:
		return "r4"
	case 11:
		return "r8"
	case 12:
		return "lpstr"
	case 13:
		return "lpwstr"
	case 14:
		return "bstr"
	case 15:
		return "date"
	case 16:
		return "filetime"
	case 17:
		return "bool"
	case 18:
		return "cy"
	case 19:
		return "error"
	case 20:
		return "clsid"
	}
	return ""
}

func (m ST_VectorBaseType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_VectorBaseType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ArrayBaseType byte

const (
	ST_ArrayBaseTypeUnset   ST_ArrayBaseType = 0
	ST_ArrayBaseTypeVariant ST_ArrayBaseType = 1
	ST_ArrayBaseTypeI1      ST_ArrayBaseType = 2
	ST_ArrayBaseTypeI2      ST_ArrayBaseType = 3
	ST_ArrayBaseTypeI4      ST_ArrayBaseType = 4
	ST_ArrayBaseTypeInt     ST_ArrayBaseType = 5
	ST_ArrayBaseTypeUi1     ST_ArrayBaseType = 6
	ST_ArrayBaseTypeUi2     ST_ArrayBaseType = 7
	ST_ArrayBaseTypeUi4     ST_ArrayBaseType = 8
	ST_ArrayBaseTypeUint    ST_ArrayBaseType = 9
	ST_ArrayBaseTypeR4      ST_ArrayBaseType = 10
	ST_ArrayBaseTypeR8      ST_ArrayBaseType = 11
	ST_ArrayBaseTypeDecimal ST_ArrayBaseType = 12
	ST_ArrayBaseTypeBstr    ST_ArrayBaseType = 13
	ST_ArrayBaseTypeDate    ST_ArrayBaseType = 14
	ST_ArrayBaseTypeBool    ST_ArrayBaseType = 15
	ST_ArrayBaseTypeCy      ST_ArrayBaseType = 16
	ST_ArrayBaseTypeError   ST_ArrayBaseType = 17
)

func (e ST_ArrayBaseType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ArrayBaseTypeUnset:
		attr.Value = ""
	case ST_ArrayBaseTypeVariant:
		attr.Value = "variant"
	case ST_ArrayBaseTypeI1:
		attr.Value = "i1"
	case ST_ArrayBaseTypeI2:
		attr.Value = "i2"
	case ST_ArrayBaseTypeI4:
		attr.Value = "i4"
	case ST_ArrayBaseTypeInt:
		attr.Value = "int"
	case ST_ArrayBaseTypeUi1:
		attr.Value = "ui1"
	case ST_ArrayBaseTypeUi2:
		attr.Value = "ui2"
	case ST_ArrayBaseTypeUi4:
		attr.Value = "ui4"
	case ST_ArrayBaseTypeUint:
		attr.Value = "uint"
	case ST_ArrayBaseTypeR4:
		attr.Value = "r4"
	case ST_ArrayBaseTypeR8:
		attr.Value = "r8"
	case ST_ArrayBaseTypeDecimal:
		attr.Value = "decimal"
	case ST_ArrayBaseTypeBstr:
		attr.Value = "bstr"
	case ST_ArrayBaseTypeDate:
		attr.Value = "date"
	case ST_ArrayBaseTypeBool:
		attr.Value = "bool"
	case ST_ArrayBaseTypeCy:
		attr.Value = "cy"
	case ST_ArrayBaseTypeError:
		attr.Value = "error"
	}
	return attr, nil
}

func (e *ST_ArrayBaseType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "variant":
		*e = 1
	case "i1":
		*e = 2
	case "i2":
		*e = 3
	case "i4":
		*e = 4
	case "int":
		*e = 5
	case "ui1":
		*e = 6
	case "ui2":
		*e = 7
	case "ui4":
		*e = 8
	case "uint":
		*e = 9
	case "r4":
		*e = 10
	case "r8":
		*e = 11
	case "decimal":
		*e = 12
	case "bstr":
		*e = 13
	case "date":
		*e = 14
	case "bool":
		*e = 15
	case "cy":
		*e = 16
	case "error":
		*e = 17
	}
	return nil
}

func (m ST_ArrayBaseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ArrayBaseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "variant":
			*m = 1
		case "i1":
			*m = 2
		case "i2":
			*m = 3
		case "i4":
			*m = 4
		case "int":
			*m = 5
		case "ui1":
			*m = 6
		case "ui2":
			*m = 7
		case "ui4":
			*m = 8
		case "uint":
			*m = 9
		case "r4":
			*m = 10
		case "r8":
			*m = 11
		case "decimal":
			*m = 12
		case "bstr":
			*m = 13
		case "date":
			*m = 14
		case "bool":
			*m = 15
		case "cy":
			*m = 16
		case "error":
			*m = 17
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ArrayBaseType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "variant"
	case 2:
		return "i1"
	case 3:
		return "i2"
	case 4:
		return "i4"
	case 5:
		return "int"
	case 6:
		return "ui1"
	case 7:
		return "ui2"
	case 8:
		return "ui4"
	case 9:
		return "uint"
	case 10:
		return "r4"
	case 11:
		return "r8"
	case 12:
		return "decimal"
	case 13:
		return "bstr"
	case 14:
		return "date"
	case 15:
		return "bool"
	case 16:
		return "cy"
	case 17:
		return "error"
	}
	return ""
}

func (m ST_ArrayBaseType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ArrayBaseType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "CT_Empty", NewCT_Empty)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "CT_Null", NewCT_Null)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "CT_Vector", NewCT_Vector)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "CT_Array", NewCT_Array)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "CT_Variant", NewCT_Variant)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "CT_Vstream", NewCT_Vstream)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "variant", NewVariant)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "vector", NewVector)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "array", NewArray)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "empty", NewEmpty)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "null", NewNull)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", "vstream", NewVstream)
}
