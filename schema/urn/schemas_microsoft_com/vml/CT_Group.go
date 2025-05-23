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
	"fmt"
	"strconv"

	"github.com/928799934/gooxml"
	"github.com/928799934/gooxml/schema/soo/ofc/sharedTypes"
	"github.com/928799934/gooxml/schema/urn/schemas_microsoft_com/office/excel"
	"github.com/928799934/gooxml/schema/urn/schemas_microsoft_com/office/powerpoint"
	"github.com/928799934/gooxml/schema/urn/schemas_microsoft_com/office/word"
)

type CT_Group struct {
	EditasAttr            ST_EditAs
	TablepropertiesAttr   *string
	TablelimitsAttr       *string
	Group                 []*Group
	Shape                 []*Shape
	Shapetype             []*Shapetype
	Arc                   []*Arc
	Curve                 []*Curve
	Image                 []*Image
	Line                  []*Line
	Oval                  []*Oval
	Polyline              []*Polyline
	Rect                  []*Rect
	Roundrect             []*Roundrect
	Diagram               []*OfcDiagram
	EG_ShapeElements      []*EG_ShapeElements
	HrefAttr              *string
	TargetAttr            *string
	ClassAttr             *string
	TitleAttr             *string
	AltAttr               *string
	CoordsizeAttr         *string
	CoordoriginAttr       *string
	WrapcoordsAttr        *string
	PrintAttr             sharedTypes.ST_TrueFalse
	IdAttr                *string
	StyleAttr             *string
	SpidAttr              *string
	OnedAttr              sharedTypes.ST_TrueFalse
	RegroupidAttr         *int64
	DoubleclicknotifyAttr sharedTypes.ST_TrueFalse
	ButtonAttr            sharedTypes.ST_TrueFalse
	UserhiddenAttr        sharedTypes.ST_TrueFalse
	BulletAttr            sharedTypes.ST_TrueFalse
	HrAttr                sharedTypes.ST_TrueFalse
	HrstdAttr             sharedTypes.ST_TrueFalse
	HrnoshadeAttr         sharedTypes.ST_TrueFalse
	HrpctAttr             *float32
	HralignAttr           OfcST_HrAlign
	AllowincellAttr       sharedTypes.ST_TrueFalse
	AllowoverlapAttr      sharedTypes.ST_TrueFalse
	UserdrawnAttr         sharedTypes.ST_TrueFalse
	BordertopcolorAttr    *string
	BorderleftcolorAttr   *string
	BorderbottomcolorAttr *string
	BorderrightcolorAttr  *string
	DgmlayoutAttr         OfcST_DiagramLayout
	DgmnodekindAttr       *int64
	DgmlayoutmruAttr      OfcST_DiagramLayout
	InsetmodeAttr         OfcST_InsetMode
	FilledAttr            sharedTypes.ST_TrueFalse
	FillcolorAttr         *string
}

func NewCT_Group() *CT_Group {
	ret := &CT_Group{}
	return ret
}

func (m *CT_Group) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.EditasAttr != ST_EditAsUnset {
		attr, err := m.EditasAttr.MarshalXMLAttr(xml.Name{Local: "editas"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TablepropertiesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:tableproperties"},
			Value: fmt.Sprintf("%v", *m.TablepropertiesAttr)})
	}
	if m.TablelimitsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:tablelimits"},
			Value: fmt.Sprintf("%v", *m.TablelimitsAttr)})
	}
	if m.HrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "href"},
			Value: fmt.Sprintf("%v", *m.HrefAttr)})
	}
	if m.TargetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "target"},
			Value: fmt.Sprintf("%v", *m.TargetAttr)})
	}
	if m.ClassAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "class"},
			Value: fmt.Sprintf("%v", *m.ClassAttr)})
	}
	if m.TitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "title"},
			Value: fmt.Sprintf("%v", *m.TitleAttr)})
	}
	if m.AltAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "alt"},
			Value: fmt.Sprintf("%v", *m.AltAttr)})
	}
	if m.CoordsizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coordsize"},
			Value: fmt.Sprintf("%v", *m.CoordsizeAttr)})
	}
	if m.CoordoriginAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coordorigin"},
			Value: fmt.Sprintf("%v", *m.CoordoriginAttr)})
	}
	if m.WrapcoordsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "wrapcoords"},
			Value: fmt.Sprintf("%v", *m.WrapcoordsAttr)})
	}
	if m.PrintAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.PrintAttr.MarshalXMLAttr(xml.Name{Local: "print"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.StyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"},
			Value: fmt.Sprintf("%v", *m.StyleAttr)})
	}
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.OnedAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OnedAttr.MarshalXMLAttr(xml.Name{Local: "oned"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RegroupidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:regroupid"},
			Value: fmt.Sprintf("%v", *m.RegroupidAttr)})
	}
	if m.DoubleclicknotifyAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.DoubleclicknotifyAttr.MarshalXMLAttr(xml.Name{Local: "doubleclicknotify"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ButtonAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ButtonAttr.MarshalXMLAttr(xml.Name{Local: "button"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.UserhiddenAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.UserhiddenAttr.MarshalXMLAttr(xml.Name{Local: "userhidden"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BulletAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.BulletAttr.MarshalXMLAttr(xml.Name{Local: "bullet"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.HrAttr.MarshalXMLAttr(xml.Name{Local: "hr"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrstdAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.HrstdAttr.MarshalXMLAttr(xml.Name{Local: "hrstd"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrnoshadeAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.HrnoshadeAttr.MarshalXMLAttr(xml.Name{Local: "hrnoshade"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrpctAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:hrpct"},
			Value: fmt.Sprintf("%v", *m.HrpctAttr)})
	}
	if m.HralignAttr != OfcST_HrAlignUnset {
		attr, err := m.HralignAttr.MarshalXMLAttr(xml.Name{Local: "hralign"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AllowincellAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AllowincellAttr.MarshalXMLAttr(xml.Name{Local: "allowincell"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AllowoverlapAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AllowoverlapAttr.MarshalXMLAttr(xml.Name{Local: "allowoverlap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.UserdrawnAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.UserdrawnAttr.MarshalXMLAttr(xml.Name{Local: "userdrawn"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BordertopcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:bordertopcolor"},
			Value: fmt.Sprintf("%v", *m.BordertopcolorAttr)})
	}
	if m.BorderleftcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:borderleftcolor"},
			Value: fmt.Sprintf("%v", *m.BorderleftcolorAttr)})
	}
	if m.BorderbottomcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:borderbottomcolor"},
			Value: fmt.Sprintf("%v", *m.BorderbottomcolorAttr)})
	}
	if m.BorderrightcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:borderrightcolor"},
			Value: fmt.Sprintf("%v", *m.BorderrightcolorAttr)})
	}
	if m.DgmlayoutAttr != OfcST_DiagramLayoutUnset {
		attr, err := m.DgmlayoutAttr.MarshalXMLAttr(xml.Name{Local: "dgmlayout"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DgmnodekindAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:dgmnodekind"},
			Value: fmt.Sprintf("%v", *m.DgmnodekindAttr)})
	}
	if m.DgmlayoutmruAttr != OfcST_DiagramLayoutUnset {
		attr, err := m.DgmlayoutmruAttr.MarshalXMLAttr(xml.Name{Local: "dgmlayoutmru"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.InsetmodeAttr != OfcST_InsetModeUnset {
		attr, err := m.InsetmodeAttr.MarshalXMLAttr(xml.Name{Local: "insetmode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FilledAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.FilledAttr.MarshalXMLAttr(xml.Name{Local: "filled"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FillcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillcolor"},
			Value: fmt.Sprintf("%v", *m.FillcolorAttr)})
	}
	e.EncodeToken(start)
	if m.Group != nil {
		segroup := xml.StartElement{Name: xml.Name{Local: "v:group"}}
		for _, c := range m.Group {
			e.EncodeElement(c, segroup)
		}
	}
	if m.Shape != nil {
		seshape := xml.StartElement{Name: xml.Name{Local: "v:shape"}}
		for _, c := range m.Shape {
			e.EncodeElement(c, seshape)
		}
	}
	if m.Shapetype != nil {
		seshapetype := xml.StartElement{Name: xml.Name{Local: "v:shapetype"}}
		for _, c := range m.Shapetype {
			e.EncodeElement(c, seshapetype)
		}
	}
	if m.Arc != nil {
		searc := xml.StartElement{Name: xml.Name{Local: "v:arc"}}
		for _, c := range m.Arc {
			e.EncodeElement(c, searc)
		}
	}
	if m.Curve != nil {
		securve := xml.StartElement{Name: xml.Name{Local: "v:curve"}}
		for _, c := range m.Curve {
			e.EncodeElement(c, securve)
		}
	}
	if m.Image != nil {
		seimage := xml.StartElement{Name: xml.Name{Local: "v:image"}}
		for _, c := range m.Image {
			e.EncodeElement(c, seimage)
		}
	}
	if m.Line != nil {
		seline := xml.StartElement{Name: xml.Name{Local: "v:line"}}
		for _, c := range m.Line {
			e.EncodeElement(c, seline)
		}
	}
	if m.Oval != nil {
		seoval := xml.StartElement{Name: xml.Name{Local: "v:oval"}}
		for _, c := range m.Oval {
			e.EncodeElement(c, seoval)
		}
	}
	if m.Polyline != nil {
		sepolyline := xml.StartElement{Name: xml.Name{Local: "v:polyline"}}
		for _, c := range m.Polyline {
			e.EncodeElement(c, sepolyline)
		}
	}
	if m.Rect != nil {
		serect := xml.StartElement{Name: xml.Name{Local: "v:rect"}}
		for _, c := range m.Rect {
			e.EncodeElement(c, serect)
		}
	}
	if m.Roundrect != nil {
		seroundrect := xml.StartElement{Name: xml.Name{Local: "v:roundrect"}}
		for _, c := range m.Roundrect {
			e.EncodeElement(c, seroundrect)
		}
	}
	if m.Diagram != nil {
		sediagram := xml.StartElement{Name: xml.Name{Local: "o:diagram"}}
		for _, c := range m.Diagram {
			e.EncodeElement(c, sediagram)
		}
	}
	if m.EG_ShapeElements != nil {
		for _, c := range m.EG_ShapeElements {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Group) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
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
lCT_Group:
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
				gooxml.Log("skipping unsupported element on CT_Group %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Group
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Group and its children
func (m *CT_Group) Validate() error {
	return m.ValidateWithPath("CT_Group")
}

// ValidateWithPath validates the CT_Group and its children, prefixing error messages with path
func (m *CT_Group) ValidateWithPath(path string) error {
	if err := m.EditasAttr.ValidateWithPath(path + "/EditasAttr"); err != nil {
		return err
	}
	for i, v := range m.Group {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Group[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Shape {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Shape[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Shapetype {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Shapetype[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Arc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Arc[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Curve {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Curve[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Image {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Image[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Line {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Line[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Oval {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Oval[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Polyline {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Polyline[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rect {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rect[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Roundrect {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Roundrect[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Diagram {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Diagram[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.EG_ShapeElements {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ShapeElements[%d]", path, i)); err != nil {
			return err
		}
	}
	if err := m.PrintAttr.ValidateWithPath(path + "/PrintAttr"); err != nil {
		return err
	}
	if err := m.OnedAttr.ValidateWithPath(path + "/OnedAttr"); err != nil {
		return err
	}
	if err := m.DoubleclicknotifyAttr.ValidateWithPath(path + "/DoubleclicknotifyAttr"); err != nil {
		return err
	}
	if err := m.ButtonAttr.ValidateWithPath(path + "/ButtonAttr"); err != nil {
		return err
	}
	if err := m.UserhiddenAttr.ValidateWithPath(path + "/UserhiddenAttr"); err != nil {
		return err
	}
	if err := m.BulletAttr.ValidateWithPath(path + "/BulletAttr"); err != nil {
		return err
	}
	if err := m.HrAttr.ValidateWithPath(path + "/HrAttr"); err != nil {
		return err
	}
	if err := m.HrstdAttr.ValidateWithPath(path + "/HrstdAttr"); err != nil {
		return err
	}
	if err := m.HrnoshadeAttr.ValidateWithPath(path + "/HrnoshadeAttr"); err != nil {
		return err
	}
	if err := m.HralignAttr.ValidateWithPath(path + "/HralignAttr"); err != nil {
		return err
	}
	if err := m.AllowincellAttr.ValidateWithPath(path + "/AllowincellAttr"); err != nil {
		return err
	}
	if err := m.AllowoverlapAttr.ValidateWithPath(path + "/AllowoverlapAttr"); err != nil {
		return err
	}
	if err := m.UserdrawnAttr.ValidateWithPath(path + "/UserdrawnAttr"); err != nil {
		return err
	}
	if err := m.DgmlayoutAttr.ValidateWithPath(path + "/DgmlayoutAttr"); err != nil {
		return err
	}
	if err := m.DgmlayoutmruAttr.ValidateWithPath(path + "/DgmlayoutmruAttr"); err != nil {
		return err
	}
	if err := m.InsetmodeAttr.ValidateWithPath(path + "/InsetmodeAttr"); err != nil {
		return err
	}
	if err := m.FilledAttr.ValidateWithPath(path + "/FilledAttr"); err != nil {
		return err
	}
	return nil
}
