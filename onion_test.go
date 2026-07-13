/*
Copyright (C) Namecoin

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package ncasn_test

import (
	"slices"
	"testing"

	"github.com/sofia-nep/ncasn"
)

const SAMPLE_ONION = "rw6nbpjrmcpdxszn3air4bt7t75rpz4cp3c2kbdu72ptua57tzvin4id.onion"

var SAMPLE_ONION_RECORD = ncasn.OnionV3Record{
	Bytes: []byte{
		0x8d, 0xbc, 0xd0, 0xbd,
		0x31, 0x60, 0x9e, 0x3b,
		0xcb, 0x2d, 0xd8, 0x11,
		0x1e, 0x06, 0x7f, 0x9f,
		0xfb, 0x17, 0xe7, 0x82,
		0x7e, 0xc5, 0xa5, 0x04,
		0x74, 0xfe, 0x9f, 0x3a,
		0x03, 0xbf, 0x9e, 0x6a,
		0x86, 0xf1,
	},
}

func TestOnionRecordFromDomain(t *testing.T) {
	record, err := ncasn.OnionRecordFromDomain(SAMPLE_ONION)
	if err != nil {
		t.Logf("err != nil: %s", err.Error())
		t.FailNow()
	}

	if !slices.Equal(record.Bytes, SAMPLE_ONION_RECORD.Bytes) {
		t.Error("record != sample")
	}
}

func TestOnionDomainFromRecord(t *testing.T) {
	domain := SAMPLE_ONION_RECORD.ToDomain()
	if domain != SAMPLE_ONION {
		t.Errorf("domain != sample: %s != %s", domain, SAMPLE_ONION)
	}
}
