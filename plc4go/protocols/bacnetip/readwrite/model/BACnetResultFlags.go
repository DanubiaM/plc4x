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

// BACnetResultFlags is an enum
type BACnetResultFlags uint8

type IBACnetResultFlags interface {
	fmt.Stringer
	utils.Serializable
}

const (
	BACnetResultFlags_FIRST_ITEM BACnetResultFlags = 0
	BACnetResultFlags_LAST_ITEM  BACnetResultFlags = 1
	BACnetResultFlags_MORE_ITEMS BACnetResultFlags = 2
)

var BACnetResultFlagsValues []BACnetResultFlags

func init() {
	_ = errors.New
	BACnetResultFlagsValues = []BACnetResultFlags{
		BACnetResultFlags_FIRST_ITEM,
		BACnetResultFlags_LAST_ITEM,
		BACnetResultFlags_MORE_ITEMS,
	}
}

func BACnetResultFlagsByValue(value uint8) (enum BACnetResultFlags, ok bool) {
	switch value {
	case 0:
		return BACnetResultFlags_FIRST_ITEM, true
	case 1:
		return BACnetResultFlags_LAST_ITEM, true
	case 2:
		return BACnetResultFlags_MORE_ITEMS, true
	}
	return 0, false
}

func BACnetResultFlagsByName(value string) (enum BACnetResultFlags, ok bool) {
	switch value {
	case "FIRST_ITEM":
		return BACnetResultFlags_FIRST_ITEM, true
	case "LAST_ITEM":
		return BACnetResultFlags_LAST_ITEM, true
	case "MORE_ITEMS":
		return BACnetResultFlags_MORE_ITEMS, true
	}
	return 0, false
}

func BACnetResultFlagsKnows(value uint8) bool {
	for _, typeValue := range BACnetResultFlagsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetResultFlags(structType interface{}) BACnetResultFlags {
	castFunc := func(typ interface{}) BACnetResultFlags {
		if sBACnetResultFlags, ok := typ.(BACnetResultFlags); ok {
			return sBACnetResultFlags
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetResultFlags) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetResultFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetResultFlagsParse(ctx context.Context, theBytes []byte) (BACnetResultFlags, error) {
	return BACnetResultFlagsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetResultFlagsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetResultFlags, error) {
	val, err := readBuffer.ReadUint8("BACnetResultFlags", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetResultFlags")
	}
	if enum, ok := BACnetResultFlagsByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetResultFlags(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetResultFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetResultFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetResultFlags", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetResultFlags) PLC4XEnumName() string {
	switch e {
	case BACnetResultFlags_FIRST_ITEM:
		return "FIRST_ITEM"
	case BACnetResultFlags_LAST_ITEM:
		return "LAST_ITEM"
	case BACnetResultFlags_MORE_ITEMS:
		return "MORE_ITEMS"
	}
	return ""
}

func (e BACnetResultFlags) String() string {
	return e.PLC4XEnumName()
}
