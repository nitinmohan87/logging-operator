// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filter

import (
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/secret"
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/types"
)

// +kubebuilder:object:generate=true
// +docName:"The grep filter plugin "greps" events by the values of specified fields."
// More info at https://docs.fluentd.org/filter/grep

// #### Example grep filter configurations
// ```
// spec:
//  filters:
//    - regexp:
//        - key: elso
//          pattern: /^5\d\d$/
//        - key: masodik
//          pattern: /\.css$/
//    - and:
//        - regexp:
//          - key: elso
//            pattern: /^5\d\d$/
//          exclude:
//          - key: masodik
//            pattern: /\.css$/
// ```
//
// #### Fluentd Config Result
// ```
//	<filter **>
//		@type grep
//		@id test_grep
//		<regexp>
//		  key elso
//		  pattern /^5\d\d$/
//		</regexp>
//		<regexp>
//		  key masodik
//		  pattern /\.css$/
//		</regexp>
//		<and>
//		  <regexp>
//			key elso
//			pattern /^5\d\d$/
//		  </regexp>
//		  <exclude>
//			key masodik
//			pattern /\.css$/
//		  </exclude>
//		</and>
//	  </filter>
// ```
type _docGrep interface{}

type GrepConfig struct {
	// +docLink:"Regexp Section,#Regex-Section"
	Regexp []RegexpSection `json:"regexp,omitempty"`
	// +docLink:"Exclude Section,#Exclude-Section"
	Exclude []ExcludeSection `json:"exclude,omitempty"`
	// +docLink:"Or Section,#Or-Section"
	Or []OrSection `json:"or,omitempty"`
	// +docLink:"And Section,#And-Section"
	And []AndSection `json:"and,omitempty"`
}

// +kubebuilder:object:generate=true
// +docName:"Regexp Section"
type RegexpSection struct {
	// Specify field name in the record to parse.
	Key string `json:"key"`
	// Pattern expression to evaluate
	Pattern string `json:"pattern"`
}

// +kubebuilder:object:generate=true
// +docName:"Exclude Section"
type ExcludeSection struct {
	// Specify field name in the record to parse.
	Key string `json:"key"`
	// Pattern expression to evaluate
	Pattern string `json:"pattern"`
}

// +kubebuilder:object:generate=true
// +docName:"Or directive"
type OrSection struct {
	// +docLink:"Regexp Section,#Regex-Section"
	Regexp []RegexpSection `json:"regexp,omitempty"`
	// +docLink:"Exclude Section,#Exclude-Section"
	Exclude []ExcludeSection `json:"exclude,omitempty"`
}

// +kubebuilder:object:generate=true
// +docName:"And directive"
type AndSection struct {
	// +docLink:"Regexp Section,#Regex-Section"
	Regexp []RegexpSection `json:"regexp,omitempty"`
	// +docLink:"Exclude Section,#Exclude-Section"
	Exclude []ExcludeSection `json:"exclude,omitempty"`
}

func (r *RegexpSection) ToDirective(secretLoader secret.SecretLoader, id string) (types.Directive, error) {
	meta := types.PluginMeta{
		Directive: "regexp",
	}
	return types.NewFlatDirective(meta, r, secretLoader)
}

func (r *ExcludeSection) ToDirective(secretLoader secret.SecretLoader, id string) (types.Directive, error) {
	meta := types.PluginMeta{
		Directive: "exclude",
	}
	return types.NewFlatDirective(meta, r, secretLoader)
}

func (r *OrSection) ToDirective(secretLoader secret.SecretLoader, id string) (types.Directive, error) {

	or := &types.GenericDirective{
		PluginMeta: types.PluginMeta{
			Directive: "or",
		},
	}

	if len(r.Regexp) > 0 {
		for _, regexp := range r.Regexp {
			if meta, err := regexp.ToDirective(secretLoader, ""); err != nil {
				return nil, err
			} else {
				or.SubDirectives = append(or.SubDirectives, meta)
			}
		}
	}

	if len(r.Exclude) > 0 {
		for _, exclude := range r.Exclude {
			if meta, err := exclude.ToDirective(secretLoader, ""); err != nil {
				return nil, err
			} else {
				or.SubDirectives = append(or.SubDirectives, meta)
			}
		}
	}
	return or, nil
}

func (r *AndSection) ToDirective(secretLoader secret.SecretLoader, id string) (types.Directive, error) {
	and := &types.GenericDirective{
		PluginMeta: types.PluginMeta{
			Directive: "and",
		},
	}

	if len(r.Regexp) > 0 {
		for _, regexp := range r.Regexp {
			if meta, err := regexp.ToDirective(secretLoader, ""); err != nil {
				return nil, err
			} else {
				and.SubDirectives = append(and.SubDirectives, meta)
			}
		}
	}

	if len(r.Exclude) > 0 {
		for _, exclude := range r.Exclude {
			if meta, err := exclude.ToDirective(secretLoader, ""); err != nil {
				return nil, err
			} else {
				and.SubDirectives = append(and.SubDirectives, meta)
			}
		}
	}
	return and, nil
}

func (g *GrepConfig) ToDirective(secretLoader secret.SecretLoader, id string) (types.Directive, error) {
	pluginType := "grep"
	grep := &types.GenericDirective{
		PluginMeta: types.PluginMeta{
			Type:      pluginType,
			Directive: "filter",
			Tag:       "**",
			Id:        id + "_" + pluginType,
		},
	}
	if len(g.Regexp) > 0 {
		for _, regexp := range g.Regexp {
			if meta, err := regexp.ToDirective(secretLoader, ""); err != nil {
				return nil, err
			} else {
				grep.SubDirectives = append(grep.SubDirectives, meta)
			}
		}
	}
	if len(g.Exclude) > 0 {
		for _, exclude := range g.Exclude {
			if meta, err := exclude.ToDirective(secretLoader, ""); err != nil {
				return nil, err
			} else {
				grep.SubDirectives = append(grep.SubDirectives, meta)
			}
		}
	}
	if len(g.Or) > 0 {
		for _, or := range g.Or {
			if meta, err := or.ToDirective(secretLoader, ""); err != nil {
				return nil, err
			} else {
				grep.SubDirectives = append(grep.SubDirectives, meta)
			}
		}
	}
	if len(g.And) > 0 {
		for _, and := range g.And {
			if meta, err := and.ToDirective(secretLoader, ""); err != nil {
				return nil, err
			} else {
				grep.SubDirectives = append(grep.SubDirectives, meta)
			}
		}
	}

	return grep, nil
}
