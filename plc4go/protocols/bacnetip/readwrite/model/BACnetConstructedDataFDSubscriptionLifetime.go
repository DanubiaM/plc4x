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

// BACnetConstructedDataFDSubscriptionLifetime is the corresponding interface of BACnetConstructedDataFDSubscriptionLifetime
type BACnetConstructedDataFDSubscriptionLifetime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFdSubscriptionLifetime returns FdSubscriptionLifetime (property field)
	GetFdSubscriptionLifetime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataFDSubscriptionLifetimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataFDSubscriptionLifetime.
// This is useful for switch cases.
type BACnetConstructedDataFDSubscriptionLifetimeExactly interface {
	BACnetConstructedDataFDSubscriptionLifetime
	isBACnetConstructedDataFDSubscriptionLifetime() bool
}

// _BACnetConstructedDataFDSubscriptionLifetime is the data-structure of this message
type _BACnetConstructedDataFDSubscriptionLifetime struct {
	*_BACnetConstructedData
	FdSubscriptionLifetime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FD_SUBSCRIPTION_LIFETIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFDSubscriptionLifetime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetFdSubscriptionLifetime() BACnetApplicationTagUnsignedInteger {
	return m.FdSubscriptionLifetime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetFdSubscriptionLifetime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFDSubscriptionLifetime factory function for _BACnetConstructedDataFDSubscriptionLifetime
func NewBACnetConstructedDataFDSubscriptionLifetime(fdSubscriptionLifetime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFDSubscriptionLifetime {
	_result := &_BACnetConstructedDataFDSubscriptionLifetime{
		FdSubscriptionLifetime: fdSubscriptionLifetime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFDSubscriptionLifetime(structType interface{}) BACnetConstructedDataFDSubscriptionLifetime {
	if casted, ok := structType.(BACnetConstructedDataFDSubscriptionLifetime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFDSubscriptionLifetime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetTypeName() string {
	return "BACnetConstructedDataFDSubscriptionLifetime"
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (fdSubscriptionLifetime)
	lengthInBits += m.FdSubscriptionLifetime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataFDSubscriptionLifetimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFDSubscriptionLifetime, error) {
	return BACnetConstructedDataFDSubscriptionLifetimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataFDSubscriptionLifetimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFDSubscriptionLifetime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFDSubscriptionLifetime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFDSubscriptionLifetime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fdSubscriptionLifetime)
	if pullErr := readBuffer.PullContext("fdSubscriptionLifetime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fdSubscriptionLifetime")
	}
	_fdSubscriptionLifetime, _fdSubscriptionLifetimeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _fdSubscriptionLifetimeErr != nil {
		return nil, errors.Wrap(_fdSubscriptionLifetimeErr, "Error parsing 'fdSubscriptionLifetime' field of BACnetConstructedDataFDSubscriptionLifetime")
	}
	fdSubscriptionLifetime := _fdSubscriptionLifetime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("fdSubscriptionLifetime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fdSubscriptionLifetime")
	}

	// Virtual field
	_actualValue := fdSubscriptionLifetime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFDSubscriptionLifetime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFDSubscriptionLifetime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataFDSubscriptionLifetime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FdSubscriptionLifetime: fdSubscriptionLifetime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFDSubscriptionLifetime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFDSubscriptionLifetime")
		}

		// Simple Field (fdSubscriptionLifetime)
		if pushErr := writeBuffer.PushContext("fdSubscriptionLifetime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fdSubscriptionLifetime")
		}
		_fdSubscriptionLifetimeErr := writeBuffer.WriteSerializable(ctx, m.GetFdSubscriptionLifetime())
		if popErr := writeBuffer.PopContext("fdSubscriptionLifetime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fdSubscriptionLifetime")
		}
		if _fdSubscriptionLifetimeErr != nil {
			return errors.Wrap(_fdSubscriptionLifetimeErr, "Error serializing 'fdSubscriptionLifetime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFDSubscriptionLifetime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFDSubscriptionLifetime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) isBACnetConstructedDataFDSubscriptionLifetime() bool {
	return true
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
