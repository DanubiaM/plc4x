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

// BACnetShedState is an enum
type BACnetShedState uint8

type IBACnetShedState interface {
	fmt.Stringer
	utils.Serializable
}

const (
	BACnetShedState_SHED_INACTIVE        BACnetShedState = 0
	BACnetShedState_SHED_REQUEST_PENDING BACnetShedState = 1
	BACnetShedState_SHED_COMPLIANT       BACnetShedState = 2
	BACnetShedState_SHED_NON_COMPLIANT   BACnetShedState = 3
)

var BACnetShedStateValues []BACnetShedState

func init() {
	_ = errors.New
	BACnetShedStateValues = []BACnetShedState{
		BACnetShedState_SHED_INACTIVE,
		BACnetShedState_SHED_REQUEST_PENDING,
		BACnetShedState_SHED_COMPLIANT,
		BACnetShedState_SHED_NON_COMPLIANT,
	}
}

func BACnetShedStateByValue(value uint8) (enum BACnetShedState, ok bool) {
	switch value {
	case 0:
		return BACnetShedState_SHED_INACTIVE, true
	case 1:
		return BACnetShedState_SHED_REQUEST_PENDING, true
	case 2:
		return BACnetShedState_SHED_COMPLIANT, true
	case 3:
		return BACnetShedState_SHED_NON_COMPLIANT, true
	}
	return 0, false
}

func BACnetShedStateByName(value string) (enum BACnetShedState, ok bool) {
	switch value {
	case "SHED_INACTIVE":
		return BACnetShedState_SHED_INACTIVE, true
	case "SHED_REQUEST_PENDING":
		return BACnetShedState_SHED_REQUEST_PENDING, true
	case "SHED_COMPLIANT":
		return BACnetShedState_SHED_COMPLIANT, true
	case "SHED_NON_COMPLIANT":
		return BACnetShedState_SHED_NON_COMPLIANT, true
	}
	return 0, false
}

func BACnetShedStateKnows(value uint8) bool {
	for _, typeValue := range BACnetShedStateValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetShedState(structType interface{}) BACnetShedState {
	castFunc := func(typ interface{}) BACnetShedState {
		if sBACnetShedState, ok := typ.(BACnetShedState); ok {
			return sBACnetShedState
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetShedState) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetShedState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetShedStateParse(ctx context.Context, theBytes []byte) (BACnetShedState, error) {
	return BACnetShedStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetShedStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetShedState, error) {
	val, err := readBuffer.ReadUint8("BACnetShedState", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetShedState")
	}
	if enum, ok := BACnetShedStateByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetShedState(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetShedState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetShedState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetShedState", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetShedState) PLC4XEnumName() string {
	switch e {
	case BACnetShedState_SHED_INACTIVE:
		return "SHED_INACTIVE"
	case BACnetShedState_SHED_REQUEST_PENDING:
		return "SHED_REQUEST_PENDING"
	case BACnetShedState_SHED_COMPLIANT:
		return "SHED_COMPLIANT"
	case BACnetShedState_SHED_NON_COMPLIANT:
		return "SHED_NON_COMPLIANT"
	}
	return ""
}

func (e BACnetShedState) String() string {
	return e.PLC4XEnumName()
}
