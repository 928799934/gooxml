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
	"strconv"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/soo/dml"
)

type CT_SlideMaster struct {
	// Preserve Slide Master
	PreserveAttr *bool
	// Common slide data for slide masters
	CSld *CT_CommonSlideData
	// Color Scheme Map
	ClrMap *dml.CT_ColorMapping
	// List of Slide Layouts
	SldLayoutIdLst *CT_SlideLayoutIdList
	// Slide Transition for a Slide Master
	Transition *CT_SlideTransition
	// Slide Timing Information for Slide Masters
	Timing *CT_SlideTiming
	// Header/Footer information for a slide master
	Hf *CT_HeaderFooter
	// Slide Master Text Styles
	TxStyles *CT_SlideMasterTextStyles
	ExtLst   *CT_ExtensionListModify
}

func NewCT_SlideMaster() *CT_SlideMaster {
	ret := &CT_SlideMaster{}
	ret.CSld = NewCT_CommonSlideData()
	ret.ClrMap = dml.NewCT_ColorMapping()
	return ret
}

func (m *CT_SlideMaster) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PreserveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preserve"},
			Value: fmt.Sprintf("%d", b2i(*m.PreserveAttr))})
	}
	e.EncodeToken(start)
	secSld := xml.StartElement{Name: xml.Name{Local: "p:cSld"}}
	e.EncodeElement(m.CSld, secSld)
	seclrMap := xml.StartElement{Name: xml.Name{Local: "p:clrMap"}}
	e.EncodeElement(m.ClrMap, seclrMap)
	if m.SldLayoutIdLst != nil {
		sesldLayoutIdLst := xml.StartElement{Name: xml.Name{Local: "p:sldLayoutIdLst"}}
		e.EncodeElement(m.SldLayoutIdLst, sesldLayoutIdLst)
	}
	if m.Transition != nil {
		setransition := xml.StartElement{Name: xml.Name{Local: "p:transition"}}
		e.EncodeElement(m.Transition, setransition)
	}
	if m.Timing != nil {
		setiming := xml.StartElement{Name: xml.Name{Local: "p:timing"}}
		e.EncodeElement(m.Timing, setiming)
	}
	if m.Hf != nil {
		sehf := xml.StartElement{Name: xml.Name{Local: "p:hf"}}
		e.EncodeElement(m.Hf, sehf)
	}
	if m.TxStyles != nil {
		setxStyles := xml.StartElement{Name: xml.Name{Local: "p:txStyles"}}
		e.EncodeElement(m.TxStyles, setxStyles)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideMaster) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CSld = NewCT_CommonSlideData()
	m.ClrMap = dml.NewCT_ColorMapping()
	for _, attr := range start.Attr {
		if attr.Name.Local == "preserve" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveAttr = &parsed
			continue
		}
	}
lCT_SlideMaster:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cSld"}:
				if err := d.DecodeElement(m.CSld, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "clrMap"}:
				if err := d.DecodeElement(m.ClrMap, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldLayoutIdLst"}:
				m.SldLayoutIdLst = NewCT_SlideLayoutIdList()
				if err := d.DecodeElement(m.SldLayoutIdLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "transition"}:
				m.Transition = NewCT_SlideTransition()
				if err := d.DecodeElement(m.Transition, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "timing"}:
				m.Timing = NewCT_SlideTiming()
				if err := d.DecodeElement(m.Timing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "hf"}:
				m.Hf = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.Hf, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "txStyles"}:
				m.TxStyles = NewCT_SlideMasterTextStyles()
				if err := d.DecodeElement(m.TxStyles, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_SlideMaster %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideMaster
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideMaster and its children
func (m *CT_SlideMaster) Validate() error {
	return m.ValidateWithPath("CT_SlideMaster")
}

// ValidateWithPath validates the CT_SlideMaster and its children, prefixing error messages with path
func (m *CT_SlideMaster) ValidateWithPath(path string) error {
	if err := m.CSld.ValidateWithPath(path + "/CSld"); err != nil {
		return err
	}
	if err := m.ClrMap.ValidateWithPath(path + "/ClrMap"); err != nil {
		return err
	}
	if m.SldLayoutIdLst != nil {
		if err := m.SldLayoutIdLst.ValidateWithPath(path + "/SldLayoutIdLst"); err != nil {
			return err
		}
	}
	if m.Transition != nil {
		if err := m.Transition.ValidateWithPath(path + "/Transition"); err != nil {
			return err
		}
	}
	if m.Timing != nil {
		if err := m.Timing.ValidateWithPath(path + "/Timing"); err != nil {
			return err
		}
	}
	if m.Hf != nil {
		if err := m.Hf.ValidateWithPath(path + "/Hf"); err != nil {
			return err
		}
	}
	if m.TxStyles != nil {
		if err := m.TxStyles.ValidateWithPath(path + "/TxStyles"); err != nil {
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
