// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml_test

import "testing"
import "github.com/928799934/gooxml"

func TestOpenSourceLicense(t *testing.T) {
	gooxml.GetLicense()
}
