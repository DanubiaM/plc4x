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

// ErrorReportingSystemCategoryTypeForInputUnits is an enum
type ErrorReportingSystemCategoryTypeForInputUnits uint8

type IErrorReportingSystemCategoryTypeForInputUnits interface {
	fmt.Stringer
	utils.Serializable
}

const (
	ErrorReportingSystemCategoryTypeForInputUnits_KEY_UNITS                    ErrorReportingSystemCategoryTypeForInputUnits = 0x0
	ErrorReportingSystemCategoryTypeForInputUnits_TELECOMMAND_AND_REMOTE_ENTRY ErrorReportingSystemCategoryTypeForInputUnits = 0x1
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_2                   ErrorReportingSystemCategoryTypeForInputUnits = 0x2
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_3                   ErrorReportingSystemCategoryTypeForInputUnits = 0x3
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_4                   ErrorReportingSystemCategoryTypeForInputUnits = 0x4
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_5                   ErrorReportingSystemCategoryTypeForInputUnits = 0x5
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_6                   ErrorReportingSystemCategoryTypeForInputUnits = 0x6
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_7                   ErrorReportingSystemCategoryTypeForInputUnits = 0x7
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_8                   ErrorReportingSystemCategoryTypeForInputUnits = 0x8
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_9                   ErrorReportingSystemCategoryTypeForInputUnits = 0x9
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_10                  ErrorReportingSystemCategoryTypeForInputUnits = 0xA
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_11                  ErrorReportingSystemCategoryTypeForInputUnits = 0xB
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_12                  ErrorReportingSystemCategoryTypeForInputUnits = 0xC
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_13                  ErrorReportingSystemCategoryTypeForInputUnits = 0xD
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_14                  ErrorReportingSystemCategoryTypeForInputUnits = 0xE
	ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_15                  ErrorReportingSystemCategoryTypeForInputUnits = 0xF
)

var ErrorReportingSystemCategoryTypeForInputUnitsValues []ErrorReportingSystemCategoryTypeForInputUnits

func init() {
	_ = errors.New
	ErrorReportingSystemCategoryTypeForInputUnitsValues = []ErrorReportingSystemCategoryTypeForInputUnits{
		ErrorReportingSystemCategoryTypeForInputUnits_KEY_UNITS,
		ErrorReportingSystemCategoryTypeForInputUnits_TELECOMMAND_AND_REMOTE_ENTRY,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_2,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_3,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_4,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_5,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_6,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_7,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_8,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_9,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_10,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_11,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_12,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_13,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_14,
		ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_15,
	}
}

func ErrorReportingSystemCategoryTypeForInputUnitsByValue(value uint8) (enum ErrorReportingSystemCategoryTypeForInputUnits, ok bool) {
	switch value {
	case 0x0:
		return ErrorReportingSystemCategoryTypeForInputUnits_KEY_UNITS, true
	case 0x1:
		return ErrorReportingSystemCategoryTypeForInputUnits_TELECOMMAND_AND_REMOTE_ENTRY, true
	case 0x2:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_2, true
	case 0x3:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_3, true
	case 0x4:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_4, true
	case 0x5:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_5, true
	case 0x6:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_6, true
	case 0x7:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_7, true
	case 0x8:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_8, true
	case 0x9:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_9, true
	case 0xA:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_10, true
	case 0xB:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_11, true
	case 0xC:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_12, true
	case 0xD:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_13, true
	case 0xE:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_14, true
	case 0xF:
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_15, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryTypeForInputUnitsByName(value string) (enum ErrorReportingSystemCategoryTypeForInputUnits, ok bool) {
	switch value {
	case "KEY_UNITS":
		return ErrorReportingSystemCategoryTypeForInputUnits_KEY_UNITS, true
	case "TELECOMMAND_AND_REMOTE_ENTRY":
		return ErrorReportingSystemCategoryTypeForInputUnits_TELECOMMAND_AND_REMOTE_ENTRY, true
	case "RESERVED_2":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_2, true
	case "RESERVED_3":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_3, true
	case "RESERVED_4":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_4, true
	case "RESERVED_5":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_5, true
	case "RESERVED_6":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_6, true
	case "RESERVED_7":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_7, true
	case "RESERVED_8":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_8, true
	case "RESERVED_9":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_9, true
	case "RESERVED_10":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_10, true
	case "RESERVED_11":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_11, true
	case "RESERVED_12":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_12, true
	case "RESERVED_13":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_13, true
	case "RESERVED_14":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_14, true
	case "RESERVED_15":
		return ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_15, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryTypeForInputUnitsKnows(value uint8) bool {
	for _, typeValue := range ErrorReportingSystemCategoryTypeForInputUnitsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastErrorReportingSystemCategoryTypeForInputUnits(structType interface{}) ErrorReportingSystemCategoryTypeForInputUnits {
	castFunc := func(typ interface{}) ErrorReportingSystemCategoryTypeForInputUnits {
		if sErrorReportingSystemCategoryTypeForInputUnits, ok := typ.(ErrorReportingSystemCategoryTypeForInputUnits); ok {
			return sErrorReportingSystemCategoryTypeForInputUnits
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorReportingSystemCategoryTypeForInputUnits) GetLengthInBits(ctx context.Context) uint16 {
	return 4
}

func (m ErrorReportingSystemCategoryTypeForInputUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryTypeForInputUnitsParse(ctx context.Context, theBytes []byte) (ErrorReportingSystemCategoryTypeForInputUnits, error) {
	return ErrorReportingSystemCategoryTypeForInputUnitsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingSystemCategoryTypeForInputUnitsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSystemCategoryTypeForInputUnits, error) {
	val, err := readBuffer.ReadUint8("ErrorReportingSystemCategoryTypeForInputUnits", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorReportingSystemCategoryTypeForInputUnits")
	}
	if enum, ok := ErrorReportingSystemCategoryTypeForInputUnitsByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ErrorReportingSystemCategoryTypeForInputUnits(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorReportingSystemCategoryTypeForInputUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ErrorReportingSystemCategoryTypeForInputUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ErrorReportingSystemCategoryTypeForInputUnits", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorReportingSystemCategoryTypeForInputUnits) PLC4XEnumName() string {
	switch e {
	case ErrorReportingSystemCategoryTypeForInputUnits_KEY_UNITS:
		return "KEY_UNITS"
	case ErrorReportingSystemCategoryTypeForInputUnits_TELECOMMAND_AND_REMOTE_ENTRY:
		return "TELECOMMAND_AND_REMOTE_ENTRY"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_2:
		return "RESERVED_2"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_3:
		return "RESERVED_3"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_4:
		return "RESERVED_4"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_5:
		return "RESERVED_5"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_6:
		return "RESERVED_6"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_7:
		return "RESERVED_7"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_8:
		return "RESERVED_8"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_9:
		return "RESERVED_9"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_10:
		return "RESERVED_10"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_11:
		return "RESERVED_11"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_12:
		return "RESERVED_12"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_13:
		return "RESERVED_13"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_14:
		return "RESERVED_14"
	case ErrorReportingSystemCategoryTypeForInputUnits_RESERVED_15:
		return "RESERVED_15"
	}
	return ""
}

func (e ErrorReportingSystemCategoryTypeForInputUnits) String() string {
	return e.PLC4XEnumName()
}
