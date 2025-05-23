// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/928799934/gooxml"
)

type CT_PivotAreaReference struct {
	// Field Index
	FieldAttr *uint32
	// Item Index Count
	CountAttr *uint32
	// Selected
	SelectedAttr *bool
	// Positional Reference
	ByPositionAttr *bool
	// Relative Reference
	RelativeAttr *bool
	// Include Default Filter
	DefaultSubtotalAttr *bool
	// Include Sum Filter
	SumSubtotalAttr *bool
	// Include CountA Filter
	CountASubtotalAttr *bool
	// Include Average Filter
	AvgSubtotalAttr *bool
	// Include Maximum Filter
	MaxSubtotalAttr *bool
	// Include Minimum Filter
	MinSubtotalAttr *bool
	// Include Product Filter
	ProductSubtotalAttr *bool
	// Include Count Subtotal
	CountSubtotalAttr *bool
	// Include StdDev Filter
	StdDevSubtotalAttr *bool
	// Include StdDevP Filter
	StdDevPSubtotalAttr *bool
	// Include Var Filter
	VarSubtotalAttr *bool
	// Include VarP Filter
	VarPSubtotalAttr *bool
	// Field Item
	X      []*CT_Index
	ExtLst *CT_ExtensionList
}

func NewCT_PivotAreaReference() *CT_PivotAreaReference {
	ret := &CT_PivotAreaReference{}
	return ret
}

func (m *CT_PivotAreaReference) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "field"},
			Value: fmt.Sprintf("%v", *m.FieldAttr)})
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.SelectedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "selected"},
			Value: fmt.Sprintf("%d", b2i(*m.SelectedAttr))})
	}
	if m.ByPositionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "byPosition"},
			Value: fmt.Sprintf("%d", b2i(*m.ByPositionAttr))})
	}
	if m.RelativeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relative"},
			Value: fmt.Sprintf("%d", b2i(*m.RelativeAttr))})
	}
	if m.DefaultSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.DefaultSubtotalAttr))})
	}
	if m.SumSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sumSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.SumSubtotalAttr))})
	}
	if m.CountASubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "countASubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.CountASubtotalAttr))})
	}
	if m.AvgSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "avgSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.AvgSubtotalAttr))})
	}
	if m.MaxSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.MaxSubtotalAttr))})
	}
	if m.MinSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.MinSubtotalAttr))})
	}
	if m.ProductSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "productSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.ProductSubtotalAttr))})
	}
	if m.CountSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "countSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.CountSubtotalAttr))})
	}
	if m.StdDevSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stdDevSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.StdDevSubtotalAttr))})
	}
	if m.StdDevPSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stdDevPSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.StdDevPSubtotalAttr))})
	}
	if m.VarSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "varSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.VarSubtotalAttr))})
	}
	if m.VarPSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "varPSubtotal"},
			Value: fmt.Sprintf("%d", b2i(*m.VarPSubtotalAttr))})
	}
	e.EncodeToken(start)
	if m.X != nil {
		sex := xml.StartElement{Name: xml.Name{Local: "ma:x"}}
		for _, c := range m.X {
			e.EncodeElement(c, sex)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotAreaReference) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "maxSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MaxSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "minSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MinSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
		if attr.Name.Local == "productSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ProductSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "byPosition" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ByPositionAttr = &parsed
			continue
		}
		if attr.Name.Local == "countSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CountSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "defaultSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DefaultSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "countASubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CountASubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "avgSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AvgSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "field" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.FieldAttr = &pt
			continue
		}
		if attr.Name.Local == "selected" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SelectedAttr = &parsed
			continue
		}
		if attr.Name.Local == "relative" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RelativeAttr = &parsed
			continue
		}
		if attr.Name.Local == "sumSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SumSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "stdDevSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StdDevSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "stdDevPSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StdDevPSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "varSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VarSubtotalAttr = &parsed
			continue
		}
		if attr.Name.Local == "varPSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VarPSubtotalAttr = &parsed
			continue
		}
	}
lCT_PivotAreaReference:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "x"}:
				tmp := NewCT_Index()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.X = append(m.X, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_PivotAreaReference %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotAreaReference
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotAreaReference and its children
func (m *CT_PivotAreaReference) Validate() error {
	return m.ValidateWithPath("CT_PivotAreaReference")
}

// ValidateWithPath validates the CT_PivotAreaReference and its children, prefixing error messages with path
func (m *CT_PivotAreaReference) ValidateWithPath(path string) error {
	for i, v := range m.X {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/X[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
