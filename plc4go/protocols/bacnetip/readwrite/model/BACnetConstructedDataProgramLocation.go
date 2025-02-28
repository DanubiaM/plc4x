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

// BACnetConstructedDataProgramLocation is the corresponding interface of BACnetConstructedDataProgramLocation
type BACnetConstructedDataProgramLocation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProgramLocation returns ProgramLocation (property field)
	GetProgramLocation() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataProgramLocationExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProgramLocation.
// This is useful for switch cases.
type BACnetConstructedDataProgramLocationExactly interface {
	BACnetConstructedDataProgramLocation
	isBACnetConstructedDataProgramLocation() bool
}

// _BACnetConstructedDataProgramLocation is the data-structure of this message
type _BACnetConstructedDataProgramLocation struct {
	*_BACnetConstructedData
	ProgramLocation BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProgramLocation) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProgramLocation) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROGRAM_LOCATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProgramLocation) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProgramLocation) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProgramLocation) GetProgramLocation() BACnetApplicationTagCharacterString {
	return m.ProgramLocation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProgramLocation) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetProgramLocation())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProgramLocation factory function for _BACnetConstructedDataProgramLocation
func NewBACnetConstructedDataProgramLocation(programLocation BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProgramLocation {
	_result := &_BACnetConstructedDataProgramLocation{
		ProgramLocation:        programLocation,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProgramLocation(structType interface{}) BACnetConstructedDataProgramLocation {
	if casted, ok := structType.(BACnetConstructedDataProgramLocation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProgramLocation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProgramLocation) GetTypeName() string {
	return "BACnetConstructedDataProgramLocation"
}

func (m *_BACnetConstructedDataProgramLocation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (programLocation)
	lengthInBits += m.ProgramLocation.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProgramLocation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataProgramLocationParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProgramLocation, error) {
	return BACnetConstructedDataProgramLocationParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataProgramLocationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProgramLocation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProgramLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProgramLocation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (programLocation)
	if pullErr := readBuffer.PullContext("programLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for programLocation")
	}
	_programLocation, _programLocationErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _programLocationErr != nil {
		return nil, errors.Wrap(_programLocationErr, "Error parsing 'programLocation' field of BACnetConstructedDataProgramLocation")
	}
	programLocation := _programLocation.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("programLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for programLocation")
	}

	// Virtual field
	_actualValue := programLocation
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProgramLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProgramLocation")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProgramLocation{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ProgramLocation: programLocation,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProgramLocation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProgramLocation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProgramLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProgramLocation")
		}

		// Simple Field (programLocation)
		if pushErr := writeBuffer.PushContext("programLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for programLocation")
		}
		_programLocationErr := writeBuffer.WriteSerializable(ctx, m.GetProgramLocation())
		if popErr := writeBuffer.PopContext("programLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for programLocation")
		}
		if _programLocationErr != nil {
			return errors.Wrap(_programLocationErr, "Error serializing 'programLocation' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProgramLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProgramLocation")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProgramLocation) isBACnetConstructedDataProgramLocation() bool {
	return true
}

func (m *_BACnetConstructedDataProgramLocation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
