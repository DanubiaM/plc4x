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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public abstract class BACnetPriorityValue implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final BACnetTagHeader peekedTagHeader;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;

  public BACnetPriorityValue(BACnetTagHeader peekedTagHeader, BACnetObjectType objectTypeArgument) {
    super();
    this.peekedTagHeader = peekedTagHeader;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetTagHeader getPeekedTagHeader() {
    return peekedTagHeader;
  }

  public short getPeekedTagNumber() {
    return (short) (getPeekedTagHeader().getActualTagNumber());
  }

  public boolean getPeekedIsContextTag() {
    return (boolean) ((getPeekedTagHeader().getTagClass()) == (TagClass.CONTEXT_SPECIFIC_TAGS));
  }

  protected abstract void serializeBACnetPriorityValueChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetPriorityValue");

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    short peekedTagNumber = getPeekedTagNumber();
    writeBuffer.writeVirtual("peekedTagNumber", peekedTagNumber);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean peekedIsContextTag = getPeekedIsContextTag();
    writeBuffer.writeVirtual("peekedIsContextTag", peekedIsContextTag);

    // Switch field (Serialize the sub-type)
    serializeBACnetPriorityValueChild(writeBuffer);

    writeBuffer.popContext("BACnetPriorityValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetPriorityValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static BACnetPriorityValue staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    BACnetObjectType objectTypeArgument;
    if (args[0] instanceof BACnetObjectType) {
      objectTypeArgument = (BACnetObjectType) args[0];
    } else if (args[0] instanceof String) {
      objectTypeArgument = BACnetObjectType.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type BACnetObjectType or a string which is parseable but"
              + " was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, objectTypeArgument);
  }

  public static BACnetPriorityValue staticParse(
      ReadBuffer readBuffer, BACnetObjectType objectTypeArgument) throws ParseException {
    readBuffer.pullContext("BACnetPriorityValue");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetTagHeader peekedTagHeader =
        readPeekField(
            "peekedTagHeader",
            new DataReaderComplexDefault<>(
                () -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    short peekedTagNumber =
        readVirtualField("peekedTagNumber", short.class, peekedTagHeader.getActualTagNumber());
    boolean peekedIsContextTag =
        readVirtualField(
            "peekedIsContextTag",
            boolean.class,
            (peekedTagHeader.getTagClass()) == (TagClass.CONTEXT_SPECIFIC_TAGS));
    // Validation
    if (!(((!(peekedIsContextTag)))
        || ((((peekedIsContextTag) && ((peekedTagHeader.getLengthValueType()) != (0x6)))
            && ((peekedTagHeader.getLengthValueType()) != (0x7)))))) {
      throw new ParseValidationException("unexpected opening or closing tag");
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    BACnetPriorityValueBuilder builder = null;
    if (EvaluationHelper.equals(peekedTagNumber, (short) 0x0)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueNull.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x4)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueReal.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x9)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueEnumerated.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x2)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueUnsigned.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x1)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueBoolean.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x3)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueInteger.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x5)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueDouble.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0xB)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueTime.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x7)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueCharacterString.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x6)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueOctetString.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0x8)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueBitString.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0xA)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueDate.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0xC)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) false)) {
      builder =
          BACnetPriorityValueObjectidentifier.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 0)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) true)) {
      builder =
          BACnetPriorityValueConstructedValue.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    } else if (EvaluationHelper.equals(peekedTagNumber, (short) 1)
        && EvaluationHelper.equals(peekedIsContextTag, (boolean) true)) {
      builder =
          BACnetPriorityValueDateTime.staticParseBACnetPriorityValueBuilder(
              readBuffer, objectTypeArgument);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "peekedTagNumber="
              + peekedTagNumber
              + " "
              + "peekedIsContextTag="
              + peekedIsContextTag
              + "]");
    }

    readBuffer.closeContext("BACnetPriorityValue");
    // Create the instance
    BACnetPriorityValue _bACnetPriorityValue = builder.build(peekedTagHeader, objectTypeArgument);
    return _bACnetPriorityValue;
  }

  public interface BACnetPriorityValueBuilder {
    BACnetPriorityValue build(BACnetTagHeader peekedTagHeader, BACnetObjectType objectTypeArgument);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPriorityValue)) {
      return false;
    }
    BACnetPriorityValue that = (BACnetPriorityValue) o;
    return (getPeekedTagHeader() == that.getPeekedTagHeader()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getPeekedTagHeader());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
