// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"

	"github.com/928799934/gooxml"
)

type CT_BorderBoxPr struct {
	HideTop    *CT_OnOff
	HideBot    *CT_OnOff
	HideLeft   *CT_OnOff
	HideRight  *CT_OnOff
	StrikeH    *CT_OnOff
	StrikeV    *CT_OnOff
	StrikeBLTR *CT_OnOff
	StrikeTLBR *CT_OnOff
	CtrlPr     *CT_CtrlPr
}

func NewCT_BorderBoxPr() *CT_BorderBoxPr {
	ret := &CT_BorderBoxPr{}
	return ret
}

func (m *CT_BorderBoxPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.HideTop != nil {
		sehideTop := xml.StartElement{Name: xml.Name{Local: "m:hideTop"}}
		e.EncodeElement(m.HideTop, sehideTop)
	}
	if m.HideBot != nil {
		sehideBot := xml.StartElement{Name: xml.Name{Local: "m:hideBot"}}
		e.EncodeElement(m.HideBot, sehideBot)
	}
	if m.HideLeft != nil {
		sehideLeft := xml.StartElement{Name: xml.Name{Local: "m:hideLeft"}}
		e.EncodeElement(m.HideLeft, sehideLeft)
	}
	if m.HideRight != nil {
		sehideRight := xml.StartElement{Name: xml.Name{Local: "m:hideRight"}}
		e.EncodeElement(m.HideRight, sehideRight)
	}
	if m.StrikeH != nil {
		sestrikeH := xml.StartElement{Name: xml.Name{Local: "m:strikeH"}}
		e.EncodeElement(m.StrikeH, sestrikeH)
	}
	if m.StrikeV != nil {
		sestrikeV := xml.StartElement{Name: xml.Name{Local: "m:strikeV"}}
		e.EncodeElement(m.StrikeV, sestrikeV)
	}
	if m.StrikeBLTR != nil {
		sestrikeBLTR := xml.StartElement{Name: xml.Name{Local: "m:strikeBLTR"}}
		e.EncodeElement(m.StrikeBLTR, sestrikeBLTR)
	}
	if m.StrikeTLBR != nil {
		sestrikeTLBR := xml.StartElement{Name: xml.Name{Local: "m:strikeTLBR"}}
		e.EncodeElement(m.StrikeTLBR, sestrikeTLBR)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BorderBoxPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BorderBoxPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "hideTop"}:
				m.HideTop = NewCT_OnOff()
				if err := d.DecodeElement(m.HideTop, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "hideBot"}:
				m.HideBot = NewCT_OnOff()
				if err := d.DecodeElement(m.HideBot, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "hideLeft"}:
				m.HideLeft = NewCT_OnOff()
				if err := d.DecodeElement(m.HideLeft, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "hideRight"}:
				m.HideRight = NewCT_OnOff()
				if err := d.DecodeElement(m.HideRight, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "strikeH"}:
				m.StrikeH = NewCT_OnOff()
				if err := d.DecodeElement(m.StrikeH, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "strikeV"}:
				m.StrikeV = NewCT_OnOff()
				if err := d.DecodeElement(m.StrikeV, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "strikeBLTR"}:
				m.StrikeBLTR = NewCT_OnOff()
				if err := d.DecodeElement(m.StrikeBLTR, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "strikeTLBR"}:
				m.StrikeTLBR = NewCT_OnOff()
				if err := d.DecodeElement(m.StrikeTLBR, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_BorderBoxPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BorderBoxPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BorderBoxPr and its children
func (m *CT_BorderBoxPr) Validate() error {
	return m.ValidateWithPath("CT_BorderBoxPr")
}

// ValidateWithPath validates the CT_BorderBoxPr and its children, prefixing error messages with path
func (m *CT_BorderBoxPr) ValidateWithPath(path string) error {
	if m.HideTop != nil {
		if err := m.HideTop.ValidateWithPath(path + "/HideTop"); err != nil {
			return err
		}
	}
	if m.HideBot != nil {
		if err := m.HideBot.ValidateWithPath(path + "/HideBot"); err != nil {
			return err
		}
	}
	if m.HideLeft != nil {
		if err := m.HideLeft.ValidateWithPath(path + "/HideLeft"); err != nil {
			return err
		}
	}
	if m.HideRight != nil {
		if err := m.HideRight.ValidateWithPath(path + "/HideRight"); err != nil {
			return err
		}
	}
	if m.StrikeH != nil {
		if err := m.StrikeH.ValidateWithPath(path + "/StrikeH"); err != nil {
			return err
		}
	}
	if m.StrikeV != nil {
		if err := m.StrikeV.ValidateWithPath(path + "/StrikeV"); err != nil {
			return err
		}
	}
	if m.StrikeBLTR != nil {
		if err := m.StrikeBLTR.ValidateWithPath(path + "/StrikeBLTR"); err != nil {
			return err
		}
	}
	if m.StrikeTLBR != nil {
		if err := m.StrikeTLBR.ValidateWithPath(path + "/StrikeTLBR"); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
