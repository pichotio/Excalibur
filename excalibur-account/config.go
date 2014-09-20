////////////////////////////////////////////////////////////////////////////////////
// Copyright (C) 2014 Putta Khunchalee.
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"encoding/xml"
	"os"
)

type IPEndPoint struct {
	XMLName xml.Name `xml:"address"`
	Name    string   `xml:"name,attr"`
	Port    uint16   `xml:"port,attr"`
}

type Server struct {
	XMLName   xml.Name     `xml:"server"`
	Addresses []IPEndPoint `xml:"address"`
}

type Config struct {
	XMLName xml.Name `xml:"excalibur-account"`
	Server  Server   `xml:"server"`
}

func LoadConfig(name string) (file *Config, err error) {
	f, e := os.Open(name)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	c := new(Config)
	x := xml.NewDecoder(f)
	e = x.Decode(c)
	if e != nil {
		return nil, e
	}

	return c, nil
}
