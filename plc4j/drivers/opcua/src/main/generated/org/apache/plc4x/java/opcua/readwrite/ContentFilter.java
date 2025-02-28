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
package org.apache.plc4x.java.opcua.readwrite;

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

public class ContentFilter extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "588";
  }

  // Properties.
  protected final int noOfElements;
  protected final List<ExtensionObjectDefinition> elements;

  public ContentFilter(int noOfElements, List<ExtensionObjectDefinition> elements) {
    super();
    this.noOfElements = noOfElements;
    this.elements = elements;
  }

  public int getNoOfElements() {
    return noOfElements;
  }

  public List<ExtensionObjectDefinition> getElements() {
    return elements;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ContentFilter");

    // Simple Field (noOfElements)
    writeSimpleField("noOfElements", noOfElements, writeSignedInt(writeBuffer, 32));

    // Array Field (elements)
    writeComplexTypeArrayField("elements", elements, writeBuffer);

    writeBuffer.popContext("ContentFilter");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ContentFilter _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (noOfElements)
    lengthInBits += 32;

    // Array field
    if (elements != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : elements) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= elements.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("ContentFilter");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfElements = readSimpleField("noOfElements", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> elements =
        readCountArrayField(
            "elements",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("585")),
                readBuffer),
            noOfElements);

    readBuffer.closeContext("ContentFilter");
    // Create the instance
    return new ContentFilterBuilderImpl(noOfElements, elements);
  }

  public static class ContentFilterBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final int noOfElements;
    private final List<ExtensionObjectDefinition> elements;

    public ContentFilterBuilderImpl(int noOfElements, List<ExtensionObjectDefinition> elements) {
      this.noOfElements = noOfElements;
      this.elements = elements;
    }

    public ContentFilter build() {
      ContentFilter contentFilter = new ContentFilter(noOfElements, elements);
      return contentFilter;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ContentFilter)) {
      return false;
    }
    ContentFilter that = (ContentFilter) o;
    return (getNoOfElements() == that.getNoOfElements())
        && (getElements() == that.getElements())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNoOfElements(), getElements());
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
