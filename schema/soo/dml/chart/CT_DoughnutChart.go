// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/928799934/gooxml"
)

type CT_DoughnutChart struct {
	VaryColors    *CT_Boolean
	Ser           []*CT_PieSer
	DLbls         *CT_DLbls
	FirstSliceAng *CT_FirstSliceAng
	HoleSize      *CT_HoleSize
	ExtLst        *CT_ExtensionList
}

func NewCT_DoughnutChart() *CT_DoughnutChart {
	ret := &CT_DoughnutChart{}
	return ret
}

func (m *CT_DoughnutChart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.VaryColors != nil {
		sevaryColors := xml.StartElement{Name: xml.Name{Local: "c:varyColors"}}
		e.EncodeElement(m.VaryColors, sevaryColors)
	}
	if m.Ser != nil {
		seser := xml.StartElement{Name: xml.Name{Local: "c:ser"}}
		for _, c := range m.Ser {
			e.EncodeElement(c, seser)
		}
	}
	if m.DLbls != nil {
		sedLbls := xml.StartElement{Name: xml.Name{Local: "c:dLbls"}}
		e.EncodeElement(m.DLbls, sedLbls)
	}
	if m.FirstSliceAng != nil {
		sefirstSliceAng := xml.StartElement{Name: xml.Name{Local: "c:firstSliceAng"}}
		e.EncodeElement(m.FirstSliceAng, sefirstSliceAng)
	}
	if m.HoleSize != nil {
		seholeSize := xml.StartElement{Name: xml.Name{Local: "c:holeSize"}}
		e.EncodeElement(m.HoleSize, seholeSize)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DoughnutChart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DoughnutChart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "varyColors"}:
				m.VaryColors = NewCT_Boolean()
				if err := d.DecodeElement(m.VaryColors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ser"}:
				tmp := NewCT_PieSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLbls"}:
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "firstSliceAng"}:
				m.FirstSliceAng = NewCT_FirstSliceAng()
				if err := d.DecodeElement(m.FirstSliceAng, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "holeSize"}:
				m.HoleSize = NewCT_HoleSize()
				if err := d.DecodeElement(m.HoleSize, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_DoughnutChart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DoughnutChart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DoughnutChart and its children
func (m *CT_DoughnutChart) Validate() error {
	return m.ValidateWithPath("CT_DoughnutChart")
}

// ValidateWithPath validates the CT_DoughnutChart and its children, prefixing error messages with path
func (m *CT_DoughnutChart) ValidateWithPath(path string) error {
	if m.VaryColors != nil {
		if err := m.VaryColors.ValidateWithPath(path + "/VaryColors"); err != nil {
			return err
		}
	}
	for i, v := range m.Ser {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ser[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DLbls != nil {
		if err := m.DLbls.ValidateWithPath(path + "/DLbls"); err != nil {
			return err
		}
	}
	if m.FirstSliceAng != nil {
		if err := m.FirstSliceAng.ValidateWithPath(path + "/FirstSliceAng"); err != nil {
			return err
		}
	}
	if m.HoleSize != nil {
		if err := m.HoleSize.ValidateWithPath(path + "/HoleSize"); err != nil {
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
