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

// BACnetConstructedDataRequestedUpdateInterval is the corresponding interface of BACnetConstructedDataRequestedUpdateInterval
type BACnetConstructedDataRequestedUpdateInterval interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRequestedUpdateInterval returns RequestedUpdateInterval (property field)
	GetRequestedUpdateInterval() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataRequestedUpdateIntervalExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRequestedUpdateInterval.
// This is useful for switch cases.
type BACnetConstructedDataRequestedUpdateIntervalExactly interface {
	BACnetConstructedDataRequestedUpdateInterval
	isBACnetConstructedDataRequestedUpdateInterval() bool
}

// _BACnetConstructedDataRequestedUpdateInterval is the data-structure of this message
type _BACnetConstructedDataRequestedUpdateInterval struct {
	*_BACnetConstructedData
	RequestedUpdateInterval BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REQUESTED_UPDATE_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRequestedUpdateInterval) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetRequestedUpdateInterval() BACnetApplicationTagUnsignedInteger {
	return m.RequestedUpdateInterval
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetRequestedUpdateInterval())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRequestedUpdateInterval factory function for _BACnetConstructedDataRequestedUpdateInterval
func NewBACnetConstructedDataRequestedUpdateInterval(requestedUpdateInterval BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRequestedUpdateInterval {
	_result := &_BACnetConstructedDataRequestedUpdateInterval{
		RequestedUpdateInterval: requestedUpdateInterval,
		_BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRequestedUpdateInterval(structType interface{}) BACnetConstructedDataRequestedUpdateInterval {
	if casted, ok := structType.(BACnetConstructedDataRequestedUpdateInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRequestedUpdateInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetTypeName() string {
	return "BACnetConstructedDataRequestedUpdateInterval"
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestedUpdateInterval)
	lengthInBits += m.RequestedUpdateInterval.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataRequestedUpdateIntervalParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRequestedUpdateInterval, error) {
	return BACnetConstructedDataRequestedUpdateIntervalParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataRequestedUpdateIntervalParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRequestedUpdateInterval, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRequestedUpdateInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRequestedUpdateInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestedUpdateInterval)
	if pullErr := readBuffer.PullContext("requestedUpdateInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestedUpdateInterval")
	}
	_requestedUpdateInterval, _requestedUpdateIntervalErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _requestedUpdateIntervalErr != nil {
		return nil, errors.Wrap(_requestedUpdateIntervalErr, "Error parsing 'requestedUpdateInterval' field of BACnetConstructedDataRequestedUpdateInterval")
	}
	requestedUpdateInterval := _requestedUpdateInterval.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("requestedUpdateInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestedUpdateInterval")
	}

	// Virtual field
	_actualValue := requestedUpdateInterval
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRequestedUpdateInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRequestedUpdateInterval")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataRequestedUpdateInterval{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		RequestedUpdateInterval: requestedUpdateInterval,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRequestedUpdateInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRequestedUpdateInterval")
		}

		// Simple Field (requestedUpdateInterval)
		if pushErr := writeBuffer.PushContext("requestedUpdateInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestedUpdateInterval")
		}
		_requestedUpdateIntervalErr := writeBuffer.WriteSerializable(ctx, m.GetRequestedUpdateInterval())
		if popErr := writeBuffer.PopContext("requestedUpdateInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestedUpdateInterval")
		}
		if _requestedUpdateIntervalErr != nil {
			return errors.Wrap(_requestedUpdateIntervalErr, "Error serializing 'requestedUpdateInterval' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRequestedUpdateInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRequestedUpdateInterval")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) isBACnetConstructedDataRequestedUpdateInterval() bool {
	return true
}

func (m *_BACnetConstructedDataRequestedUpdateInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
