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

package ncasn

import (
	"encoding/base32"
	"errors"
	"strings"
)

var I2P_ENCODING = base32.StdEncoding.WithPadding(base32.NoPadding)

func I2pRecordFromDomain(domain string) (*I2pB32Record, error) {
	b32, found := strings.CutSuffix(domain, ".b32.i2p")
	if !found {
		return nil, errors.New("Only .b32.i2p domains are supported.")
	}
	b32 = strings.ToUpper(b32)
	bytes, err := I2P_ENCODING.DecodeString(b32)
	if err != nil {
		return nil, err
	}

	return &I2pB32Record{Bytes: bytes}, nil
}

func (record *I2pB32Record) ToDomain() string {
	return strings.ToLower(I2P_ENCODING.EncodeToString(record.Bytes)) + ".b32.i2p"
}
