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

// BACnetPropertyStatesPolarity is the corresponding interface of BACnetPropertyStatesPolarity
type BACnetPropertyStatesPolarity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetPolarity returns Polarity (property field)
	GetPolarity() BACnetPolarityTagged
}

// BACnetPropertyStatesPolarityExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesPolarity.
// This is useful for switch cases.
type BACnetPropertyStatesPolarityExactly interface {
	BACnetPropertyStatesPolarity
	isBACnetPropertyStatesPolarity() bool
}

// _BACnetPropertyStatesPolarity is the data-structure of this message
type _BACnetPropertyStatesPolarity struct {
	*_BACnetPropertyStates
	Polarity BACnetPolarityTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesPolarity) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesPolarity) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesPolarity) GetPolarity() BACnetPolarityTagged {
	return m.Polarity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesPolarity factory function for _BACnetPropertyStatesPolarity
func NewBACnetPropertyStatesPolarity(polarity BACnetPolarityTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesPolarity {
	_result := &_BACnetPropertyStatesPolarity{
		Polarity:              polarity,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesPolarity(structType interface{}) BACnetPropertyStatesPolarity {
	if casted, ok := structType.(BACnetPropertyStatesPolarity); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesPolarity); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesPolarity) GetTypeName() string {
	return "BACnetPropertyStatesPolarity"
}

func (m *_BACnetPropertyStatesPolarity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (polarity)
	lengthInBits += m.Polarity.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesPolarity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesPolarityParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesPolarity, error) {
	return BACnetPropertyStatesPolarityParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesPolarityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesPolarity, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesPolarity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesPolarity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (polarity)
	if pullErr := readBuffer.PullContext("polarity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for polarity")
	}
	_polarity, _polarityErr := BACnetPolarityTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _polarityErr != nil {
		return nil, errors.Wrap(_polarityErr, "Error parsing 'polarity' field of BACnetPropertyStatesPolarity")
	}
	polarity := _polarity.(BACnetPolarityTagged)
	if closeErr := readBuffer.CloseContext("polarity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for polarity")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesPolarity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesPolarity")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesPolarity{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		Polarity:              polarity,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesPolarity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesPolarity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesPolarity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesPolarity")
		}

		// Simple Field (polarity)
		if pushErr := writeBuffer.PushContext("polarity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for polarity")
		}
		_polarityErr := writeBuffer.WriteSerializable(ctx, m.GetPolarity())
		if popErr := writeBuffer.PopContext("polarity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for polarity")
		}
		if _polarityErr != nil {
			return errors.Wrap(_polarityErr, "Error serializing 'polarity' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesPolarity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesPolarity")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesPolarity) isBACnetPropertyStatesPolarity() bool {
	return true
}

func (m *_BACnetPropertyStatesPolarity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
