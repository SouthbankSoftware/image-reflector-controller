/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package policy

import (
	"reflect"
	"testing"
)

func TestPrefixMatchFilter(t *testing.T) {
	cases := []struct {
		label    string
		tags     []string
		include  []string
		exclude  []string
		expected []string
	}{
		{
			label:    "none",
			tags:     []string{"a"},
			expected: []string{"a"},
		},
		{
			label:    "include",
			tags:     []string{"ver1", "ver2", "ver3", "rel1"},
			include:  []string{"ver"},
			expected: []string{"ver1", "ver2", "ver3"},
		},
		{
			label:    "include",
			tags:     []string{"ver1", "ver2", "ver3", "rel1"},
			exclude:  []string{"rel"},
			expected: []string{"ver1", "ver2", "ver3"},
		},
		{
			label:    "include and exclude",
			tags:     []string{"ver1", "ver2", "rel1", "rel2", "patch1", "patch2"},
			exclude:  []string{"rel", "patch"},
			expected: []string{"ver1", "ver2"},
		},
	}
	for _, tt := range cases {
		t.Run(tt.label, func(t *testing.T) {
			r := PrefixMatchFilter(tt.tags, tt.include, tt.exclude)
			if !reflect.DeepEqual(r, tt.expected) {
				t.Errorf("incorrect value returned, got '%s', expected '%s'", r, tt.expected)
			}
		})
	}
}
