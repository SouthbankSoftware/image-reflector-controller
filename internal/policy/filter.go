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

import "strings"

// PrefixMatchFilter applies filtering based on the provided lists of included and
// excluded prefixes
func PrefixMatchFilter(list []string, include []string, exclude []string) []string {
	var filtered []string
	for _, item := range list {
		// Keep by default if no include prefixes specified
		var keep bool = len(include) == 0
		for _, in := range include {
			if strings.HasPrefix(item, in) {
				keep = true
			}
		}

		for _, out := range exclude {
			if strings.HasPrefix(item, out) {
				keep = false
			}
		}

		if keep {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
