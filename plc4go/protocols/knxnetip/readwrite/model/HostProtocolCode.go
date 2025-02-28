/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// HostProtocolCode is an enum
type HostProtocolCode uint8

type IHostProtocolCode interface {
	fmt.Stringer
	utils.Serializable
}

const (
	HostProtocolCode_IPV4_UDP HostProtocolCode = 0x01
	HostProtocolCode_IPV4_TCP HostProtocolCode = 0x02
)

var HostProtocolCodeValues []HostProtocolCode

func init() {
	_ = errors.New
	HostProtocolCodeValues = []HostProtocolCode{
		HostProtocolCode_IPV4_UDP,
		HostProtocolCode_IPV4_TCP,
	}
}

func HostProtocolCodeByValue(value uint8) (enum HostProtocolCode, ok bool) {
	switch value {
	case 0x01:
		return HostProtocolCode_IPV4_UDP, true
	case 0x02:
		return HostProtocolCode_IPV4_TCP, true
	}
	return 0, false
}

func HostProtocolCodeByName(value string) (enum HostProtocolCode, ok bool) {
	switch value {
	case "IPV4_UDP":
		return HostProtocolCode_IPV4_UDP, true
	case "IPV4_TCP":
		return HostProtocolCode_IPV4_TCP, true
	}
	return 0, false
}

func HostProtocolCodeKnows(value uint8) bool {
	for _, typeValue := range HostProtocolCodeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastHostProtocolCode(structType interface{}) HostProtocolCode {
	castFunc := func(typ interface{}) HostProtocolCode {
		if sHostProtocolCode, ok := typ.(HostProtocolCode); ok {
			return sHostProtocolCode
		}
		return 0
	}
	return castFunc(structType)
}

func (m HostProtocolCode) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m HostProtocolCode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HostProtocolCodeParse(ctx context.Context, theBytes []byte) (HostProtocolCode, error) {
	return HostProtocolCodeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HostProtocolCodeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HostProtocolCode, error) {
	val, err := readBuffer.ReadUint8("HostProtocolCode", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading HostProtocolCode")
	}
	if enum, ok := HostProtocolCodeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return HostProtocolCode(val), nil
	} else {
		return enum, nil
	}
}

func (e HostProtocolCode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e HostProtocolCode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("HostProtocolCode", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e HostProtocolCode) PLC4XEnumName() string {
	switch e {
	case HostProtocolCode_IPV4_UDP:
		return "IPV4_UDP"
	case HostProtocolCode_IPV4_TCP:
		return "IPV4_TCP"
	}
	return ""
}

func (e HostProtocolCode) String() string {
	return e.PLC4XEnumName()
}
