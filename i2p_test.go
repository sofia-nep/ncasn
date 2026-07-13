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

const SAMPLE_I2P = "ukeu3k5oycgaauneqgtnvselmt4yemvoilkln7jpvamvfx7dnkdq.b32.i2p"

var SAMPLE_I2P_RECORD = ncasn.I2PB32{
	Bytes: []byte{
		0xa2, 0x89, 0x4d, 0xab,
		0xae, 0xc0, 0x8c, 0x00,
		0x51, 0xa4, 0x81, 0xa6,
		0xda, 0xc8, 0x8b, 0x64,
		0xf9, 0x82, 0x32, 0xae,
		0x42, 0xd4, 0xb6, 0xfd,
		0x2f, 0xa8, 0x19, 0x52,
		0xdf, 0xe3, 0x6a, 0x87,
	},
}

func TestI2pRecordFromDomain(t *testing.T) {
	record, err := ncasn.I2pRecordFromDomain(SAMPLE_I2P)
	if err != nil {
		t.Logf("err != nil: %s", err.Error())
		t.FailNow()
	}

	if !slices.Equal(record.Bytes, SAMPLE_I2P_RECORD.Bytes) {
		t.Error("record != sample")
	}
}

func TestI2pDomainFromRecord(t *testing.T) {
	domain := SAMPLE_I2P_RECORD.ToDomain()
	if domain != SAMPLE_I2P {
		t.Errorf("domain != sample: %s != %s", domain, SAMPLE_I2P)
	}
}
