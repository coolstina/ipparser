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

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

// NewDistrict ...
func NewDistrict(fn string, ops ...Option) (*District, error) {

	db := &District{}

	if err := db.load(fn, ops...); err != nil {
		return nil, err
	}

	return db, nil
}

// District ...
type District struct {
	file io.ReadSeeker

	index []byte
	data  []byte
}

func (db *District) load(fn string, ops ...Option) error {
	options := &option{}
	for _, o := range ops {
		o.apply(options)
	}

	var err error
	if options.embedsfs != nil {
		embeds := options.embedsfs.Embeds()
		data, err := embeds.ReadFile(fn)
		if err != nil {
			return err
		}

		db.file = bytes.NewReader(data)
	} else {

		db.file, err = os.Open(fn)
		if err != nil {
			return err
		}
	}

	b4 := make([]byte, 4)
	_, err = db.file.Read(b4)
	if err != nil {
		return err
	}

	off := int(binary.BigEndian.Uint32(b4))
	_, err = db.file.Seek(262148, 0)
	if err != nil {
		return err
	}
	db.index = make([]byte, off-262148-262144)
	_, err = db.file.Read(db.index)
	if err != nil {
		return err
	}
	db.data, err = ioutil.ReadAll(db.file)
	if err != nil {
		return err
	}
	//	fmt.Println(len(db.data))
	return nil
}

// Find ...
func (db *District) Find(s string) ([]string, error) {

	ipv := net.ParseIP(s)
	if ipv == nil {
		return nil, fmt.Errorf("%s", "ip format error.")
	}

	low := 0
	high := int(len(db.index)/13) - 1
	mid := 0

	val := binary.BigEndian.Uint32(ipv.To4())

	for low <= high {
		mid = int((low + high) / 2)
		pos := mid * 13

		start := binary.BigEndian.Uint32(db.index[pos : pos+4])
		end := binary.BigEndian.Uint32(db.index[pos+4 : pos+8])

		if val < start {
			high = mid - 1
		} else if val > end {
			low = mid + 1
		} else {

			off := int(binary.LittleEndian.Uint32(db.index[pos+8 : pos+12]))

			return strings.Split(string(db.data[off:off+int(db.index[pos+12])]), "\t"), nil
		}
	}
	return nil, fmt.Errorf("%s", "not found")
}
