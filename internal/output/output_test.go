// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package output

import (
	"testing"

	"github.com/issue9/assert"

	"github.com/caixw/apidoc/doc"
	opt "github.com/caixw/apidoc/options"
)

func getTestDoc() *doc.Doc {
	return &doc.Doc{
		Tags: []*doc.Tag{&doc.Tag{Name: "t1"}, &doc.Tag{Name: "t2"}},
		Apis: []*doc.API{
			&doc.API{Tags: []string{"t1", "tag1"}},
			&doc.API{Tags: []string{"t2", "tag2"}},
		},
	}
}

func TestFilterDoc(t *testing.T) {
	a := assert.New(t)

	d := getTestDoc()
	o := &options{}
	filterDoc(d, o)
	a.Equal(2, len(d.Tags))

	d = getTestDoc()
	o = &options{
		Output: opt.Output{
			Tags: []string{"t1"},
		},
	}
	filterDoc(d, o)
	a.Equal(1, len(d.Tags)).
		Equal(1, len(d.Apis))

	d = getTestDoc()
	o = &options{
		Output: opt.Output{
			Tags: []string{"t1", "t2"},
		},
	}
	filterDoc(d, o)
	a.Equal(2, len(d.Tags)).
		Equal(2, len(d.Apis))

	d = getTestDoc()
	o = &options{
		Output: opt.Output{
			Tags: []string{"tag1"},
		},
	}
	filterDoc(d, o)
	a.Equal(0, len(d.Tags)).
		Equal(1, len(d.Apis))

	d = getTestDoc()
	o = &options{
		Output: opt.Output{
			Tags: []string{"not-exists"},
		},
	}
	filterDoc(d, o)
	a.Equal(0, len(d.Tags)).
		Equal(0, len(d.Apis))
}
