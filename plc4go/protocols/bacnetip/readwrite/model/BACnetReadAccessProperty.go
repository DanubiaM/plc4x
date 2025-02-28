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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetReadAccessProperty is the corresponding interface of BACnetReadAccessProperty
type BACnetReadAccessProperty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetReadResult returns ReadResult (property field)
	GetReadResult() BACnetReadAccessPropertyReadResult
}

// BACnetReadAccessPropertyExactly can be used when we want exactly this type and not a type which fulfills BACnetReadAccessProperty.
// This is useful for switch cases.
type BACnetReadAccessPropertyExactly interface {
	BACnetReadAccessProperty
	isBACnetReadAccessProperty() bool
}

// _BACnetReadAccessProperty is the data-structure of this message
type _BACnetReadAccessProperty struct {
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	ReadResult         BACnetReadAccessPropertyReadResult

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetReadAccessProperty) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetReadAccessProperty) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetReadAccessProperty) GetReadResult() BACnetReadAccessPropertyReadResult {
	return m.ReadResult
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetReadAccessProperty factory function for _BACnetReadAccessProperty
func NewBACnetReadAccessProperty(propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, readResult BACnetReadAccessPropertyReadResult, objectTypeArgument BACnetObjectType) *_BACnetReadAccessProperty {
	return &_BACnetReadAccessProperty{PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, ReadResult: readResult, ObjectTypeArgument: objectTypeArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetReadAccessProperty(structType interface{}) BACnetReadAccessProperty {
	if casted, ok := structType.(BACnetReadAccessProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetReadAccessProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetReadAccessProperty) GetTypeName() string {
	return "BACnetReadAccessProperty"
}

func (m *_BACnetReadAccessProperty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (readResult)
	if m.ReadResult != nil {
		lengthInBits += m.ReadResult.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetReadAccessProperty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetReadAccessPropertyParse(theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetReadAccessProperty, error) {
	return BACnetReadAccessPropertyParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetReadAccessPropertyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetReadAccessProperty, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetReadAccessProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetReadAccessProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetReadAccessProperty")
	}
	propertyIdentifier := _propertyIdentifier.(BACnetPropertyIdentifierTagged)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for arrayIndex")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field of BACnetReadAccessProperty")
		default:
			arrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	// Optional Field (readResult) (Can be skipped, if a given expression evaluates to false)
	var readResult BACnetReadAccessPropertyReadResult = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("readResult"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for readResult")
		}
		_val, _err := BACnetReadAccessPropertyReadResultParseWithBuffer(ctx, readBuffer, objectTypeArgument, propertyIdentifier.GetValue(), (CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() interface{} { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() interface{} { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'readResult' field of BACnetReadAccessProperty")
		default:
			readResult = _val.(BACnetReadAccessPropertyReadResult)
			if closeErr := readBuffer.CloseContext("readResult"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for readResult")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetReadAccessProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetReadAccessProperty")
	}

	// Create the instance
	return &_BACnetReadAccessProperty{
		ObjectTypeArgument: objectTypeArgument,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		ReadResult:         readResult,
	}, nil
}

func (m *_BACnetReadAccessProperty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetReadAccessProperty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetReadAccessProperty"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetReadAccessProperty")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
	}
	_propertyIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyIdentifier())
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	if m.GetArrayIndex() != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for arrayIndex")
		}
		arrayIndex = m.GetArrayIndex()
		_arrayIndexErr := writeBuffer.WriteSerializable(ctx, arrayIndex)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for arrayIndex")
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Optional Field (readResult) (Can be skipped, if the value is null)
	var readResult BACnetReadAccessPropertyReadResult = nil
	if m.GetReadResult() != nil {
		if pushErr := writeBuffer.PushContext("readResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for readResult")
		}
		readResult = m.GetReadResult()
		_readResultErr := writeBuffer.WriteSerializable(ctx, readResult)
		if popErr := writeBuffer.PopContext("readResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for readResult")
		}
		if _readResultErr != nil {
			return errors.Wrap(_readResultErr, "Error serializing 'readResult' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetReadAccessProperty"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetReadAccessProperty")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetReadAccessProperty) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetReadAccessProperty) isBACnetReadAccessProperty() bool {
	return true
}

func (m *_BACnetReadAccessProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
