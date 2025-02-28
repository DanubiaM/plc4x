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

// Constant values.
const AdsConstants_ADSTCPDEFAULTPORT uint16 = uint16(48898)

// AdsConstants is the corresponding interface of AdsConstants
type AdsConstants interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// AdsConstantsExactly can be used when we want exactly this type and not a type which fulfills AdsConstants.
// This is useful for switch cases.
type AdsConstantsExactly interface {
	AdsConstants
	isAdsConstants() bool
}

// _AdsConstants is the data-structure of this message
type _AdsConstants struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsConstants) GetAdsTcpDefaultPort() uint16 {
	return AdsConstants_ADSTCPDEFAULTPORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsConstants factory function for _AdsConstants
func NewAdsConstants() *_AdsConstants {
	return &_AdsConstants{}
}

// Deprecated: use the interface for direct cast
func CastAdsConstants(structType interface{}) AdsConstants {
	if casted, ok := structType.(AdsConstants); ok {
		return casted
	}
	if casted, ok := structType.(*AdsConstants); ok {
		return *casted
	}
	return nil
}

func (m *_AdsConstants) GetTypeName() string {
	return "AdsConstants"
}

func (m *_AdsConstants) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (adsTcpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_AdsConstants) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsConstantsParse(theBytes []byte) (AdsConstants, error) {
	return AdsConstantsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AdsConstantsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsConstants, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (adsTcpDefaultPort)
	adsTcpDefaultPort, _adsTcpDefaultPortErr := readBuffer.ReadUint16("adsTcpDefaultPort", 16)
	if _adsTcpDefaultPortErr != nil {
		return nil, errors.Wrap(_adsTcpDefaultPortErr, "Error parsing 'adsTcpDefaultPort' field of AdsConstants")
	}
	if adsTcpDefaultPort != AdsConstants_ADSTCPDEFAULTPORT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AdsConstants_ADSTCPDEFAULTPORT) + " but got " + fmt.Sprintf("%d", adsTcpDefaultPort))
	}

	if closeErr := readBuffer.CloseContext("AdsConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsConstants")
	}

	// Create the instance
	return &_AdsConstants{}, nil
}

func (m *_AdsConstants) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsConstants) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AdsConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsConstants")
	}

	// Const Field (adsTcpDefaultPort)
	_adsTcpDefaultPortErr := writeBuffer.WriteUint16("adsTcpDefaultPort", 16, 48898)
	if _adsTcpDefaultPortErr != nil {
		return errors.Wrap(_adsTcpDefaultPortErr, "Error serializing 'adsTcpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("AdsConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsConstants")
	}
	return nil
}

func (m *_AdsConstants) isAdsConstants() bool {
	return true
}

func (m *_AdsConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
