// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipparser

import "github.com/coolstina/embedsfs"

type Option interface {
	apply(*option)
}

type OptionFunc func(option *option)

func (o OptionFunc) apply(option *option) {
	o(option)
}

type option struct {
	embedsfs *embedsfs.EmbedsFS
}

func WithEmbedsFS(fs *embedsfs.EmbedsFS) OptionFunc {
	return func(option *option) {
		option.embedsfs = fs
	}
}
