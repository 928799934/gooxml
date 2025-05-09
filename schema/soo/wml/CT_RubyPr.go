// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"

	"github.com/928799934/gooxml"
)

type CT_RubyPr struct {
	// Phonetic Guide Text Alignment
	RubyAlign *CT_RubyAlign
	// Phonetic Guide Text Font Size
	Hps *CT_HpsMeasure
	// Distance Between Phonetic Guide Text and Phonetic Guide Base Text
	HpsRaise *CT_HpsMeasure
	// Phonetic Guide Base Text Font Size
	HpsBaseText *CT_HpsMeasure
	// Language ID for Phonetic Guide
	Lid *CT_Lang
	// Invalidated Field Cache
	Dirty *CT_OnOff
}

func NewCT_RubyPr() *CT_RubyPr {
	ret := &CT_RubyPr{}
	ret.RubyAlign = NewCT_RubyAlign()
	ret.Hps = NewCT_HpsMeasure()
	ret.HpsRaise = NewCT_HpsMeasure()
	ret.HpsBaseText = NewCT_HpsMeasure()
	ret.Lid = NewCT_Lang()
	return ret
}

func (m *CT_RubyPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	serubyAlign := xml.StartElement{Name: xml.Name{Local: "w:rubyAlign"}}
	e.EncodeElement(m.RubyAlign, serubyAlign)
	sehps := xml.StartElement{Name: xml.Name{Local: "w:hps"}}
	e.EncodeElement(m.Hps, sehps)
	sehpsRaise := xml.StartElement{Name: xml.Name{Local: "w:hpsRaise"}}
	e.EncodeElement(m.HpsRaise, sehpsRaise)
	sehpsBaseText := xml.StartElement{Name: xml.Name{Local: "w:hpsBaseText"}}
	e.EncodeElement(m.HpsBaseText, sehpsBaseText)
	selid := xml.StartElement{Name: xml.Name{Local: "w:lid"}}
	e.EncodeElement(m.Lid, selid)
	if m.Dirty != nil {
		sedirty := xml.StartElement{Name: xml.Name{Local: "w:dirty"}}
		e.EncodeElement(m.Dirty, sedirty)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RubyPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RubyAlign = NewCT_RubyAlign()
	m.Hps = NewCT_HpsMeasure()
	m.HpsRaise = NewCT_HpsMeasure()
	m.HpsBaseText = NewCT_HpsMeasure()
	m.Lid = NewCT_Lang()
lCT_RubyPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rubyAlign"}:
				if err := d.DecodeElement(m.RubyAlign, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hps"}:
				if err := d.DecodeElement(m.Hps, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hpsRaise"}:
				if err := d.DecodeElement(m.HpsRaise, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hpsBaseText"}:
				if err := d.DecodeElement(m.HpsBaseText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "lid"}:
				if err := d.DecodeElement(m.Lid, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "dirty"}:
				m.Dirty = NewCT_OnOff()
				if err := d.DecodeElement(m.Dirty, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_RubyPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RubyPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RubyPr and its children
func (m *CT_RubyPr) Validate() error {
	return m.ValidateWithPath("CT_RubyPr")
}

// ValidateWithPath validates the CT_RubyPr and its children, prefixing error messages with path
func (m *CT_RubyPr) ValidateWithPath(path string) error {
	if err := m.RubyAlign.ValidateWithPath(path + "/RubyAlign"); err != nil {
		return err
	}
	if err := m.Hps.ValidateWithPath(path + "/Hps"); err != nil {
		return err
	}
	if err := m.HpsRaise.ValidateWithPath(path + "/HpsRaise"); err != nil {
		return err
	}
	if err := m.HpsBaseText.ValidateWithPath(path + "/HpsBaseText"); err != nil {
		return err
	}
	if err := m.Lid.ValidateWithPath(path + "/Lid"); err != nil {
		return err
	}
	if m.Dirty != nil {
		if err := m.Dirty.ValidateWithPath(path + "/Dirty"); err != nil {
			return err
		}
	}
	return nil
}
