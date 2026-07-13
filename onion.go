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
	"slices"
	"strings"
)

func OnionRecordFromDomain(domain string) (*OnionV3Record, error) {
	b32, found := strings.CutSuffix(domain, ".onion")
	if !found {
		return nil, errors.New("Not an onion address.")
	}

	bytes, err := base32.StdEncoding.DecodeString(strings.ToUpper(b32))
	if err != nil {
		return nil, err
	}

	// Omit version byte
	return &OnionV3Record{Bytes: bytes[:34]}, nil
}

func (record *OnionV3Record) ToDomain() string {
	// append() could modify the original, and we don't want that.
	bytes := slices.Concat(record.Bytes, []byte{0x03})
	return strings.ToLower(base32.StdEncoding.EncodeToString(bytes)) + ".onion"
}
