// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"strconv"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/urn/schemas_microsoft_com/office/excel"
	"github.com/928799934/gooxml/schema/urn/schemas_microsoft_com/office/powerpoint"
	"github.com/928799934/gooxml/schema/urn/schemas_microsoft_com/office/word"
)

type Group struct {
	CT_Group
}

func NewGroup() *Group {
	ret := &Group{}
	ret.CT_Group = *NewCT_Group()
	return ret
}

func (m *Group) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Group.MarshalXML(e, start)
}

func (m *Group) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Group = *NewCT_Group()
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "allowincell" {
			m.AllowincellAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bullet" {
			m.BulletAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hr" {
			m.HrAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "oned" {
			m.OnedAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hralign" {
			m.HralignAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "tablelimits" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TablelimitsAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "regroupid" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.RegroupidAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bordertopcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BordertopcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "spid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpidAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hrpct" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.HrpctAttr = &pt
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "borderrightcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderrightcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "button" {
			m.ButtonAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hrstd" {
			m.HrstdAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "dgmlayout" {
			m.DgmlayoutAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "userdrawn" {
			m.UserdrawnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "doubleclicknotify" {
			m.DoubleclicknotifyAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "userhidden" {
			m.UserhiddenAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "borderleftcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderleftcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "insetmode" {
			m.InsetmodeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "allowoverlap" {
			m.AllowoverlapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "tableproperties" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TablepropertiesAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hrnoshade" {
			m.HrnoshadeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "dgmlayoutmru" {
			m.DgmlayoutmruAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "dgmnodekind" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmnodekindAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "borderbottomcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderbottomcolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "print" {
			m.PrintAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
			continue
		}
		if attr.Name.Local == "coordorigin" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CoordoriginAttr = &parsed
			continue
		}
		if attr.Name.Local == "wrapcoords" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WrapcoordsAttr = &parsed
			continue
		}
		if attr.Name.Local == "coordsize" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CoordsizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "fillcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FillcolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "class" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ClassAttr = &parsed
			continue
		}
		if attr.Name.Local == "target" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TargetAttr = &parsed
			continue
		}
		if attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HrefAttr = &parsed
			continue
		}
		if attr.Name.Local == "editas" {
			m.EditasAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "filled" {
			m.FilledAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "alt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AltAttr = &parsed
			continue
		}
	}
lGroup:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "group"}:
				tmp := NewGroup()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Group = append(m.Group, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "shape"}:
				tmp := NewShape()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Shape = append(m.Shape, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "shapetype"}:
				tmp := NewShapetype()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Shapetype = append(m.Shapetype, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "arc"}:
				tmp := NewArc()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Arc = append(m.Arc, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "curve"}:
				tmp := NewCurve()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Curve = append(m.Curve, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "image"}:
				tmp := NewImage()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Image = append(m.Image, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "line"}:
				tmp := NewLine()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Line = append(m.Line, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "oval"}:
				tmp := NewOval()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Oval = append(m.Oval, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "polyline"}:
				tmp := NewPolyline()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Polyline = append(m.Polyline, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "rect"}:
				tmp := NewRect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rect = append(m.Rect, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "roundrect"}:
				tmp := NewRoundrect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Roundrect = append(m.Roundrect, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "diagram"}:
				tmp := NewOfcDiagram()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Diagram = append(m.Diagram, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "path"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Path = NewPath()
				if err := d.DecodeElement(tmpshapeelements.Path, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "formulas"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Formulas = NewFormulas()
				if err := d.DecodeElement(tmpshapeelements.Formulas, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "handles"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Handles = NewHandles()
				if err := d.DecodeElement(tmpshapeelements.Handles, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "fill"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Fill = NewFill()
				if err := d.DecodeElement(tmpshapeelements.Fill, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "stroke"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Stroke = NewStroke()
				if err := d.DecodeElement(tmpshapeelements.Stroke, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "shadow"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Shadow = NewShadow()
				if err := d.DecodeElement(tmpshapeelements.Shadow, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textbox"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Textbox = NewTextbox()
				if err := d.DecodeElement(tmpshapeelements.Textbox, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textpath"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Textpath = NewTextpath()
				if err := d.DecodeElement(tmpshapeelements.Textpath, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "imagedata"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Imagedata = NewImagedata()
				if err := d.DecodeElement(tmpshapeelements.Imagedata, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "skew"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Skew = NewOfcSkew()
				if err := d.DecodeElement(tmpshapeelements.Skew, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "extrusion"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Extrusion = NewOfcExtrusion()
				if err := d.DecodeElement(tmpshapeelements.Extrusion, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "callout"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Callout = NewOfcCallout()
				if err := d.DecodeElement(tmpshapeelements.Callout, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "lock"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Lock = NewOfcLock()
				if err := d.DecodeElement(tmpshapeelements.Lock, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "clippath"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Clippath = NewOfcClippath()
				if err := d.DecodeElement(tmpshapeelements.Clippath, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "signatureline"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Signatureline = NewOfcSignatureline()
				if err := d.DecodeElement(tmpshapeelements.Signatureline, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "wrap"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Wrap = word.NewWrap()
				if err := d.DecodeElement(tmpshapeelements.Wrap, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "anchorlock"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Anchorlock = word.NewAnchorlock()
				if err := d.DecodeElement(tmpshapeelements.Anchorlock, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "bordertop"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Bordertop = word.NewBordertop()
				if err := d.DecodeElement(tmpshapeelements.Bordertop, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderbottom"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Borderbottom = word.NewBorderbottom()
				if err := d.DecodeElement(tmpshapeelements.Borderbottom, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderleft"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Borderleft = word.NewBorderleft()
				if err := d.DecodeElement(tmpshapeelements.Borderleft, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderright"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Borderright = word.NewBorderright()
				if err := d.DecodeElement(tmpshapeelements.Borderright, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ClientData"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.ClientData = excel.NewClientData()
				if err := d.DecodeElement(tmpshapeelements.ClientData, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:powerpoint", Local: "textdata"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Textdata = powerpoint.NewTextdata()
				if err := d.DecodeElement(tmpshapeelements.Textdata, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			default:
				gooxml.Log("skipping unsupported element on Group %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lGroup
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Group and its children
func (m *Group) Validate() error {
	return m.ValidateWithPath("Group")
}

// ValidateWithPath validates the Group and its children, prefixing error messages with path
func (m *Group) ValidateWithPath(path string) error {
	if err := m.CT_Group.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
