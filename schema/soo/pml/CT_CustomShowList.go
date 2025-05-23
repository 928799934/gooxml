// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/928799934/gooxml"
)

type CT_CustomShowList struct {
	// Custom Show
	CustShow []*CT_CustomShow
}

func NewCT_CustomShowList() *CT_CustomShowList {
	ret := &CT_CustomShowList{}
	return ret
}

func (m *CT_CustomShowList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CustShow != nil {
		secustShow := xml.StartElement{Name: xml.Name{Local: "p:custShow"}}
		for _, c := range m.CustShow {
			e.EncodeElement(c, secustShow)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomShowList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustomShowList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "custShow"}:
				tmp := NewCT_CustomShow()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CustShow = append(m.CustShow, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_CustomShowList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomShowList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomShowList and its children
func (m *CT_CustomShowList) Validate() error {
	return m.ValidateWithPath("CT_CustomShowList")
}

// ValidateWithPath validates the CT_CustomShowList and its children, prefixing error messages with path
func (m *CT_CustomShowList) ValidateWithPath(path string) error {
	for i, v := range m.CustShow {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CustShow[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
