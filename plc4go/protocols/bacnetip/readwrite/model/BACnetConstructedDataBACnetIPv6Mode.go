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

// BACnetConstructedDataBACnetIPv6Mode is the corresponding interface of BACnetConstructedDataBACnetIPv6Mode
type BACnetConstructedDataBACnetIPv6Mode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBacnetIpv6Mode returns BacnetIpv6Mode (property field)
	GetBacnetIpv6Mode() BACnetIPModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetIPModeTagged
}

// BACnetConstructedDataBACnetIPv6ModeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBACnetIPv6Mode.
// This is useful for switch cases.
type BACnetConstructedDataBACnetIPv6ModeExactly interface {
	BACnetConstructedDataBACnetIPv6Mode
	isBACnetConstructedDataBACnetIPv6Mode() bool
}

// _BACnetConstructedDataBACnetIPv6Mode is the data-structure of this message
type _BACnetConstructedDataBACnetIPv6Mode struct {
	*_BACnetConstructedData
	BacnetIpv6Mode BACnetIPModeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IPV6_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPv6Mode) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetBacnetIpv6Mode() BACnetIPModeTagged {
	return m.BacnetIpv6Mode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetActualValue() BACnetIPModeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetIPModeTagged(m.GetBacnetIpv6Mode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBACnetIPv6Mode factory function for _BACnetConstructedDataBACnetIPv6Mode
func NewBACnetConstructedDataBACnetIPv6Mode(bacnetIpv6Mode BACnetIPModeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPv6Mode {
	_result := &_BACnetConstructedDataBACnetIPv6Mode{
		BacnetIpv6Mode:         bacnetIpv6Mode,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPv6Mode(structType interface{}) BACnetConstructedDataBACnetIPv6Mode {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPv6Mode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPv6Mode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPv6Mode"
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (bacnetIpv6Mode)
	lengthInBits += m.BacnetIpv6Mode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBACnetIPv6ModeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPv6Mode, error) {
	return BACnetConstructedDataBACnetIPv6ModeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBACnetIPv6ModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPv6Mode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPv6Mode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPv6Mode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bacnetIpv6Mode)
	if pullErr := readBuffer.PullContext("bacnetIpv6Mode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bacnetIpv6Mode")
	}
	_bacnetIpv6Mode, _bacnetIpv6ModeErr := BACnetIPModeTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _bacnetIpv6ModeErr != nil {
		return nil, errors.Wrap(_bacnetIpv6ModeErr, "Error parsing 'bacnetIpv6Mode' field of BACnetConstructedDataBACnetIPv6Mode")
	}
	bacnetIpv6Mode := _bacnetIpv6Mode.(BACnetIPModeTagged)
	if closeErr := readBuffer.CloseContext("bacnetIpv6Mode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bacnetIpv6Mode")
	}

	// Virtual field
	_actualValue := bacnetIpv6Mode
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPv6Mode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPv6Mode")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBACnetIPv6Mode{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BacnetIpv6Mode: bacnetIpv6Mode,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPv6Mode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPv6Mode")
		}

		// Simple Field (bacnetIpv6Mode)
		if pushErr := writeBuffer.PushContext("bacnetIpv6Mode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bacnetIpv6Mode")
		}
		_bacnetIpv6ModeErr := writeBuffer.WriteSerializable(ctx, m.GetBacnetIpv6Mode())
		if popErr := writeBuffer.PopContext("bacnetIpv6Mode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bacnetIpv6Mode")
		}
		if _bacnetIpv6ModeErr != nil {
			return errors.Wrap(_bacnetIpv6ModeErr, "Error serializing 'bacnetIpv6Mode' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPv6Mode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPv6Mode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) isBACnetConstructedDataBACnetIPv6Mode() bool {
	return true
}

func (m *_BACnetConstructedDataBACnetIPv6Mode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
