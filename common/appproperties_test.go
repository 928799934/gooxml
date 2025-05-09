// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/928799934/gooxml"

	"github.com/928799934/gooxml/common"
)

func TestNewAppDefaultProperties(t *testing.T) {
	ap := common.NewAppProperties()
	if ap.X() == nil {
		t.Errorf("expected non-nil internal element")
	}
	if got := ap.Application(); got != "github.com/928799934/gooxml" {
		t.Errorf("unexpected application: %s", got)
	}

	ap.X().Application = nil
	if got := ap.Application(); got != "" {
		t.Errorf("unexpected application: %s", got)
	}

	fv, _ := strconv.ParseFloat(strings.Replace(gooxml.ReleaseVersion, "v", "", -1), 64)
	if got := ap.ApplicationVersion(); got != fmt.Sprintf("%07.4f", fv) {
		t.Errorf("unexpected application version: %s", got)
	}
	ap.X().AppVersion = nil
	if got := ap.ApplicationVersion(); got != "" {
		t.Errorf("unexpected application version: %s", got)
	}
}

func TestAppPropertiesSetApplication(t *testing.T) {
	ap := common.NewAppProperties()
	if ap.X() == nil {
		t.Errorf("expected non-nil internal element")
	}
	ap.SetApplication("foo")
	if got := ap.Application(); got != "foo" {
		t.Errorf("unexpected application: %s", got)
	}
}

func TestAppPropertiesSetApplicationVersion(t *testing.T) {
	ap := common.NewAppProperties()
	if ap.X() == nil {
		t.Errorf("expected non-nil internal element")
	}
	ap.SetApplicationVersion("foo")
	if got := ap.ApplicationVersion(); got != "foo" {
		t.Errorf("unexpected application: %s", got)
	}
}

func TestAppPropertiesSetCompany(t *testing.T) {
	ap := common.NewAppProperties()
	if ap.X() == nil {
		t.Errorf("expected non-nil internal element")
	}
	ap.SetCompany("foo")
	if got := ap.Company(); got != "foo" {
		t.Errorf("unexpected company: %s", got)
	}
}
