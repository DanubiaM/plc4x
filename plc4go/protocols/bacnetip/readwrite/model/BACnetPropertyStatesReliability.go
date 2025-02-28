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

// BACnetPropertyStatesReliability is the corresponding interface of BACnetPropertyStatesReliability
type BACnetPropertyStatesReliability interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetReliability returns Reliability (property field)
	GetReliability() BACnetReliabilityTagged
}

// BACnetPropertyStatesReliabilityExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesReliability.
// This is useful for switch cases.
type BACnetPropertyStatesReliabilityExactly interface {
	BACnetPropertyStatesReliability
	isBACnetPropertyStatesReliability() bool
}

// _BACnetPropertyStatesReliability is the data-structure of this message
type _BACnetPropertyStatesReliability struct {
	*_BACnetPropertyStates
	Reliability BACnetReliabilityTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesReliability) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesReliability) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesReliability) GetReliability() BACnetReliabilityTagged {
	return m.Reliability
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesReliability factory function for _BACnetPropertyStatesReliability
func NewBACnetPropertyStatesReliability(reliability BACnetReliabilityTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesReliability {
	_result := &_BACnetPropertyStatesReliability{
		Reliability:           reliability,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesReliability(structType interface{}) BACnetPropertyStatesReliability {
	if casted, ok := structType.(BACnetPropertyStatesReliability); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesReliability); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesReliability) GetTypeName() string {
	return "BACnetPropertyStatesReliability"
}

func (m *_BACnetPropertyStatesReliability) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (reliability)
	lengthInBits += m.Reliability.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesReliability) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesReliabilityParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesReliability, error) {
	return BACnetPropertyStatesReliabilityParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesReliabilityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesReliability, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesReliability"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesReliability")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reliability)
	if pullErr := readBuffer.PullContext("reliability"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reliability")
	}
	_reliability, _reliabilityErr := BACnetReliabilityTaggedParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _reliabilityErr != nil {
		return nil, errors.Wrap(_reliabilityErr, "Error parsing 'reliability' field of BACnetPropertyStatesReliability")
	}
	reliability := _reliability.(BACnetReliabilityTagged)
	if closeErr := readBuffer.CloseContext("reliability"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reliability")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesReliability"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesReliability")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesReliability{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		Reliability:           reliability,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesReliability) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesReliability) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesReliability"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesReliability")
		}

		// Simple Field (reliability)
		if pushErr := writeBuffer.PushContext("reliability"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reliability")
		}
		_reliabilityErr := writeBuffer.WriteSerializable(ctx, m.GetReliability())
		if popErr := writeBuffer.PopContext("reliability"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reliability")
		}
		if _reliabilityErr != nil {
			return errors.Wrap(_reliabilityErr, "Error serializing 'reliability' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesReliability"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesReliability")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesReliability) isBACnetPropertyStatesReliability() bool {
	return true
}

func (m *_BACnetPropertyStatesReliability) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
