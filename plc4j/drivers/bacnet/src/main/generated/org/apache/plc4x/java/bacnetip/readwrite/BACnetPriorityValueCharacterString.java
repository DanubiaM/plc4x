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

public class BACnetPriorityValueCharacterString extends BACnetPriorityValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagCharacterString characterStringValue;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;

  public BACnetPriorityValueCharacterString(
      BACnetTagHeader peekedTagHeader,
      BACnetApplicationTagCharacterString characterStringValue,
      BACnetObjectType objectTypeArgument) {
    super(peekedTagHeader, objectTypeArgument);
    this.characterStringValue = characterStringValue;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetApplicationTagCharacterString getCharacterStringValue() {
    return characterStringValue;
  }

  @Override
  protected void serializeBACnetPriorityValueChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetPriorityValueCharacterString");

    // Simple Field (characterStringValue)
    writeSimpleField(
        "characterStringValue", characterStringValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPriorityValueCharacterString");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPriorityValueCharacterString _value = this;

    // Simple field (characterStringValue)
    lengthInBits += characterStringValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPriorityValueCharacterStringBuilder staticParseBuilder(
      ReadBuffer readBuffer, BACnetObjectType objectTypeArgument) throws ParseException {
    readBuffer.pullContext("BACnetPriorityValueCharacterString");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetApplicationTagCharacterString characterStringValue =
        readSimpleField(
            "characterStringValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagCharacterString)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetPriorityValueCharacterString");
    // Create the instance
    return new BACnetPriorityValueCharacterStringBuilder(characterStringValue, objectTypeArgument);
  }

  public static class BACnetPriorityValueCharacterStringBuilder
      implements BACnetPriorityValue.BACnetPriorityValueBuilder {
    private final BACnetApplicationTagCharacterString characterStringValue;
    private final BACnetObjectType objectTypeArgument;

    public BACnetPriorityValueCharacterStringBuilder(
        BACnetApplicationTagCharacterString characterStringValue,
        BACnetObjectType objectTypeArgument) {

      this.characterStringValue = characterStringValue;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetPriorityValueCharacterString build(
        BACnetTagHeader peekedTagHeader, BACnetObjectType objectTypeArgument) {
      BACnetPriorityValueCharacterString bACnetPriorityValueCharacterString =
          new BACnetPriorityValueCharacterString(
              peekedTagHeader, characterStringValue, objectTypeArgument);
      return bACnetPriorityValueCharacterString;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPriorityValueCharacterString)) {
      return false;
    }
    BACnetPriorityValueCharacterString that = (BACnetPriorityValueCharacterString) o;
    return (getCharacterStringValue() == that.getCharacterStringValue())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCharacterStringValue());
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