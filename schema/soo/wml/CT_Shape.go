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
	"fmt"
	"github.com/928799934/gooxml"
	"strconv"
)

type CT_Shape struct {
	//↓↓↓  属性  ↓↓↓

	// 复制仿写  - 文本属性
	IdAttr           *string
	SpidAttr         *string
	TypedAttr        *string
	StyleAttr        *string
	FillcolorAttr    *string
	FilledAttr       *string
	CoordsizeAttr    *string
	AllowoverlapAttr *string

	// 复制仿写  - 数字属性
	SptAttr *uint32
	AdjAttr *uint32

	//  ↓↓↓ 下级结构体  ↓↓↓
	// Spacing Above Paragraph

	Path      *CT_ShapePath
	Fill      *CT_ShapeFill
	Stroke    *CT_ShapeStroke
	Imagedata *CT_ShapeImagedata
	Lock      *CT_ShapeLock
	TextPath  *CT_ShapTextPath
	Wrap      *CT_ShapeWrap
}

func NewCT_Shape() *CT_Shape {
	ret := &CT_Shape{}
	return ret
}

func (m *CT_Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// 属性编码 -  文本
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.TypedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"},
			Value: fmt.Sprintf("%v", *m.TypedAttr)})
	}
	if m.StyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"},
			Value: fmt.Sprintf("%v", *m.StyleAttr)})
	}
	if m.FillcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillcolor"},
			Value: fmt.Sprintf("%v", *m.FillcolorAttr)})
	}
	if m.FilledAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "filled"},
			Value: fmt.Sprintf("%v", *m.FilledAttr)})
	}
	if m.CoordsizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coordsize"},
			Value: fmt.Sprintf("%v", *m.CoordsizeAttr)})
	}
	if m.AllowoverlapAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:allowoverlap"},
			Value: fmt.Sprintf("%v", *m.AllowoverlapAttr)})
	}

	// 属性编码 -  数字
	if m.SptAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:spt"},
			Value: fmt.Sprintf("%v", *m.SptAttr)})
	}
	if m.AdjAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "adj"},
			Value: fmt.Sprintf("%v", *m.AdjAttr)})
	}

	e.EncodeToken(start)

	// 节点编码
	if m.Path != nil {
		path := xml.StartElement{Name: xml.Name{Local: "v:path"}}
		e.EncodeElement(m.Path, path)
	}
	if m.Fill != nil {
		fill := xml.StartElement{Name: xml.Name{Local: "v:fill"}}
		e.EncodeElement(m.Fill, fill)
	}
	if m.Stroke != nil {
		stroke := xml.StartElement{Name: xml.Name{Local: "v:stroke"}}
		e.EncodeElement(m.Stroke, stroke)
	}
	if m.Imagedata != nil {
		imagedata := xml.StartElement{Name: xml.Name{Local: "v:imagedata"}}
		e.EncodeElement(m.Imagedata, imagedata)
	}
	if m.Lock != nil {
		lock := xml.StartElement{Name: xml.Name{Local: "o:lock"}}
		e.EncodeElement(m.Lock, lock)
	}
	if m.TextPath != nil {
		textpath := xml.StartElement{Name: xml.Name{Local: "v:textpath"}}
		e.EncodeElement(m.TextPath, textpath)
	}
	if m.Wrap != nil {
		wrap := xml.StartElement{Name: xml.Name{Local: "w10:wrap"}}
		e.EncodeElement(m.Wrap, wrap)
	}

	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Shape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	// initialize to default
	for _, attr := range start.Attr {
		//fmt.Printf("反序列化xml的时候输出属性：%+v\n", attr.Name.Local)
		// 数字属性
		if attr.Name.Local == "spt" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SptAttr = &pt
			continue
		}
		if attr.Name.Local == "adj" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.AdjAttr = &pt
			continue
		}
		// 文本属性
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "spid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpidAttr = &parsed
			continue
		}
		if attr.Name.Local == "type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypedAttr = &parsed
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
		if attr.Name.Local == "fillcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FillcolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "filled" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FilledAttr = &parsed
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
		if attr.Name.Local == "allowoverlap" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AllowoverlapAttr = &parsed
			continue
		}

	}

lCT_Path:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "path"}:
				m.Path = NewCT_ShapePath()
				if err := d.DecodeElement(m.Path, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "fill"}:
				m.Fill = NewCT_ShapeFill()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "stroke"}:
				m.Stroke = NewCT_ShapeStroke()
				if err := d.DecodeElement(m.Stroke, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "imagedata"}:
				m.Imagedata = NewCT_ShapeImagedata()
				if err := d.DecodeElement(m.Imagedata, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "lock"}:
				m.Lock = NewCT_ShapeLock()
				if err := d.DecodeElement(m.Lock, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textpath"}:
				m.TextPath = NewCT_ShapTextPath()
				if err := d.DecodeElement(m.TextPath, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "wrap"}:
				m.Wrap = NewCT_ShapeWrap()
				if err := d.DecodeElement(m.Wrap, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Path %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Path
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Shape and its children
func (m *CT_Shape) Validate() error {
	return m.ValidateWithPath("CT_Shape")
}

// ValidateWithPath validates the CT_Shape and its children, prefixing error messages with path
func (m *CT_Shape) ValidateWithPath(path string) error {
	//if m.BeforeAttr != nil {
	//	if err := m.BeforeAttr.ValidateWithPath(path + "/BeforeAttr"); err != nil {
	//		return err
	//	}
	//}
	//if m.BeforeAutospacingAttr != nil {
	//	if err := m.BeforeAutospacingAttr.ValidateWithPath(path + "/BeforeAutospacingAttr"); err != nil {
	//		return err
	//	}
	//}
	//if m.AfterAttr != nil {
	//	if err := m.AfterAttr.ValidateWithPath(path + "/AfterAttr"); err != nil {
	//		return err
	//	}
	//}
	//if m.AfterAutospacingAttr != nil {
	//	if err := m.AfterAutospacingAttr.ValidateWithPath(path + "/AfterAutospacingAttr"); err != nil {
	//		return err
	//	}
	//}
	//if m.LineAttr != nil {
	//	if err := m.LineAttr.ValidateWithPath(path + "/LineAttr"); err != nil {
	//		return err
	//	}
	//}
	//if err := m.LineRuleAttr.ValidateWithPath(path + "/LineRuleAttr"); err != nil {
	//	return err
	//}
	return nil
}
