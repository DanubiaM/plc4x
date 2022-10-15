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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ErrorReportingSystemCategoryClass is an enum
type ErrorReportingSystemCategoryClass uint8

type IErrorReportingSystemCategoryClass interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	ErrorReportingSystemCategoryClass_RESERVED_0                  ErrorReportingSystemCategoryClass = 0x0
	ErrorReportingSystemCategoryClass_RESERVED_1                  ErrorReportingSystemCategoryClass = 0x1
	ErrorReportingSystemCategoryClass_RESERVED_2                  ErrorReportingSystemCategoryClass = 0x2
	ErrorReportingSystemCategoryClass_RESERVED_3                  ErrorReportingSystemCategoryClass = 0x3
	ErrorReportingSystemCategoryClass_RESERVED_4                  ErrorReportingSystemCategoryClass = 0x4
	ErrorReportingSystemCategoryClass_INPUT_UNITS                 ErrorReportingSystemCategoryClass = 0x5
	ErrorReportingSystemCategoryClass_RESERVED_6                  ErrorReportingSystemCategoryClass = 0x6
	ErrorReportingSystemCategoryClass_RESERVED_7                  ErrorReportingSystemCategoryClass = 0x7
	ErrorReportingSystemCategoryClass_RESERVED_8                  ErrorReportingSystemCategoryClass = 0x8
	ErrorReportingSystemCategoryClass_SUPPORT_UNITS               ErrorReportingSystemCategoryClass = 0x9
	ErrorReportingSystemCategoryClass_RESERVED_10                 ErrorReportingSystemCategoryClass = 0xA
	ErrorReportingSystemCategoryClass_BUILDING_MANAGEMENT_SYSTEMS ErrorReportingSystemCategoryClass = 0xB
	ErrorReportingSystemCategoryClass_RESERVED_12                 ErrorReportingSystemCategoryClass = 0xC
	ErrorReportingSystemCategoryClass_OUTPUT_UNITS                ErrorReportingSystemCategoryClass = 0xD
	ErrorReportingSystemCategoryClass_RESERVED_14                 ErrorReportingSystemCategoryClass = 0xE
	ErrorReportingSystemCategoryClass_CLIMATE_CONTROLLERS         ErrorReportingSystemCategoryClass = 0xF
)

var ErrorReportingSystemCategoryClassValues []ErrorReportingSystemCategoryClass

func init() {
	_ = errors.New
	ErrorReportingSystemCategoryClassValues = []ErrorReportingSystemCategoryClass{
		ErrorReportingSystemCategoryClass_RESERVED_0,
		ErrorReportingSystemCategoryClass_RESERVED_1,
		ErrorReportingSystemCategoryClass_RESERVED_2,
		ErrorReportingSystemCategoryClass_RESERVED_3,
		ErrorReportingSystemCategoryClass_RESERVED_4,
		ErrorReportingSystemCategoryClass_INPUT_UNITS,
		ErrorReportingSystemCategoryClass_RESERVED_6,
		ErrorReportingSystemCategoryClass_RESERVED_7,
		ErrorReportingSystemCategoryClass_RESERVED_8,
		ErrorReportingSystemCategoryClass_SUPPORT_UNITS,
		ErrorReportingSystemCategoryClass_RESERVED_10,
		ErrorReportingSystemCategoryClass_BUILDING_MANAGEMENT_SYSTEMS,
		ErrorReportingSystemCategoryClass_RESERVED_12,
		ErrorReportingSystemCategoryClass_OUTPUT_UNITS,
		ErrorReportingSystemCategoryClass_RESERVED_14,
		ErrorReportingSystemCategoryClass_CLIMATE_CONTROLLERS,
	}
}

func ErrorReportingSystemCategoryClassByValue(value uint8) (enum ErrorReportingSystemCategoryClass, ok bool) {
	switch value {
	case 0x0:
		return ErrorReportingSystemCategoryClass_RESERVED_0, true
	case 0x1:
		return ErrorReportingSystemCategoryClass_RESERVED_1, true
	case 0x2:
		return ErrorReportingSystemCategoryClass_RESERVED_2, true
	case 0x3:
		return ErrorReportingSystemCategoryClass_RESERVED_3, true
	case 0x4:
		return ErrorReportingSystemCategoryClass_RESERVED_4, true
	case 0x5:
		return ErrorReportingSystemCategoryClass_INPUT_UNITS, true
	case 0x6:
		return ErrorReportingSystemCategoryClass_RESERVED_6, true
	case 0x7:
		return ErrorReportingSystemCategoryClass_RESERVED_7, true
	case 0x8:
		return ErrorReportingSystemCategoryClass_RESERVED_8, true
	case 0x9:
		return ErrorReportingSystemCategoryClass_SUPPORT_UNITS, true
	case 0xA:
		return ErrorReportingSystemCategoryClass_RESERVED_10, true
	case 0xB:
		return ErrorReportingSystemCategoryClass_BUILDING_MANAGEMENT_SYSTEMS, true
	case 0xC:
		return ErrorReportingSystemCategoryClass_RESERVED_12, true
	case 0xD:
		return ErrorReportingSystemCategoryClass_OUTPUT_UNITS, true
	case 0xE:
		return ErrorReportingSystemCategoryClass_RESERVED_14, true
	case 0xF:
		return ErrorReportingSystemCategoryClass_CLIMATE_CONTROLLERS, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryClassByName(value string) (enum ErrorReportingSystemCategoryClass, ok bool) {
	switch value {
	case "RESERVED_0":
		return ErrorReportingSystemCategoryClass_RESERVED_0, true
	case "RESERVED_1":
		return ErrorReportingSystemCategoryClass_RESERVED_1, true
	case "RESERVED_2":
		return ErrorReportingSystemCategoryClass_RESERVED_2, true
	case "RESERVED_3":
		return ErrorReportingSystemCategoryClass_RESERVED_3, true
	case "RESERVED_4":
		return ErrorReportingSystemCategoryClass_RESERVED_4, true
	case "INPUT_UNITS":
		return ErrorReportingSystemCategoryClass_INPUT_UNITS, true
	case "RESERVED_6":
		return ErrorReportingSystemCategoryClass_RESERVED_6, true
	case "RESERVED_7":
		return ErrorReportingSystemCategoryClass_RESERVED_7, true
	case "RESERVED_8":
		return ErrorReportingSystemCategoryClass_RESERVED_8, true
	case "SUPPORT_UNITS":
		return ErrorReportingSystemCategoryClass_SUPPORT_UNITS, true
	case "RESERVED_10":
		return ErrorReportingSystemCategoryClass_RESERVED_10, true
	case "BUILDING_MANAGEMENT_SYSTEMS":
		return ErrorReportingSystemCategoryClass_BUILDING_MANAGEMENT_SYSTEMS, true
	case "RESERVED_12":
		return ErrorReportingSystemCategoryClass_RESERVED_12, true
	case "OUTPUT_UNITS":
		return ErrorReportingSystemCategoryClass_OUTPUT_UNITS, true
	case "RESERVED_14":
		return ErrorReportingSystemCategoryClass_RESERVED_14, true
	case "CLIMATE_CONTROLLERS":
		return ErrorReportingSystemCategoryClass_CLIMATE_CONTROLLERS, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryClassKnows(value uint8) bool {
	for _, typeValue := range ErrorReportingSystemCategoryClassValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastErrorReportingSystemCategoryClass(structType interface{}) ErrorReportingSystemCategoryClass {
	castFunc := func(typ interface{}) ErrorReportingSystemCategoryClass {
		if sErrorReportingSystemCategoryClass, ok := typ.(ErrorReportingSystemCategoryClass); ok {
			return sErrorReportingSystemCategoryClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorReportingSystemCategoryClass) GetLengthInBits() uint16 {
	return 4
}

func (m ErrorReportingSystemCategoryClass) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ErrorReportingSystemCategoryClassParse(readBuffer utils.ReadBuffer) (ErrorReportingSystemCategoryClass, error) {
	val, err := readBuffer.ReadUint8("ErrorReportingSystemCategoryClass", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorReportingSystemCategoryClass")
	}
	if enum, ok := ErrorReportingSystemCategoryClassByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ErrorReportingSystemCategoryClass(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorReportingSystemCategoryClass) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ErrorReportingSystemCategoryClass", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorReportingSystemCategoryClass) PLC4XEnumName() string {
	switch e {
	case ErrorReportingSystemCategoryClass_RESERVED_0:
		return "RESERVED_0"
	case ErrorReportingSystemCategoryClass_RESERVED_1:
		return "RESERVED_1"
	case ErrorReportingSystemCategoryClass_RESERVED_2:
		return "RESERVED_2"
	case ErrorReportingSystemCategoryClass_RESERVED_3:
		return "RESERVED_3"
	case ErrorReportingSystemCategoryClass_RESERVED_4:
		return "RESERVED_4"
	case ErrorReportingSystemCategoryClass_INPUT_UNITS:
		return "INPUT_UNITS"
	case ErrorReportingSystemCategoryClass_RESERVED_6:
		return "RESERVED_6"
	case ErrorReportingSystemCategoryClass_RESERVED_7:
		return "RESERVED_7"
	case ErrorReportingSystemCategoryClass_RESERVED_8:
		return "RESERVED_8"
	case ErrorReportingSystemCategoryClass_SUPPORT_UNITS:
		return "SUPPORT_UNITS"
	case ErrorReportingSystemCategoryClass_RESERVED_10:
		return "RESERVED_10"
	case ErrorReportingSystemCategoryClass_BUILDING_MANAGEMENT_SYSTEMS:
		return "BUILDING_MANAGEMENT_SYSTEMS"
	case ErrorReportingSystemCategoryClass_RESERVED_12:
		return "RESERVED_12"
	case ErrorReportingSystemCategoryClass_OUTPUT_UNITS:
		return "OUTPUT_UNITS"
	case ErrorReportingSystemCategoryClass_RESERVED_14:
		return "RESERVED_14"
	case ErrorReportingSystemCategoryClass_CLIMATE_CONTROLLERS:
		return "CLIMATE_CONTROLLERS"
	}
	return ""
}

func (e ErrorReportingSystemCategoryClass) String() string {
	return e.PLC4XEnumName()
}