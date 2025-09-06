package fileio

// .scd - Script data parsing logic

import (
	"encoding/binary"
	"fmt"
	"io"
)

func LoadRDT_SCDStream(fileReader io.ReaderAt, fileLength int64) (*SCDOutput, error) {
	streamReader := io.NewSectionReader(fileReader, int64(0), fileLength)
	firstOffset := uint16(0)
	if err := binary.Read(streamReader, binary.LittleEndian, &firstOffset); err != nil {
		return nil, err
	}

	functionOffsets := make([]uint16, 0)
	functionOffsets = append(functionOffsets, firstOffset)
	for i := 2; i < int(firstOffset); i += 2 {
		nextOffset := uint16(0)
		if err := binary.Read(streamReader, binary.LittleEndian, &nextOffset); err != nil {
			return nil, err
		}
		functionOffsets = append(functionOffsets, nextOffset)
	}

	programCounter := 0
	scriptData := ScriptFunction{}
	scriptData.Instructions = make(map[int][]byte)
	scriptData.StartProgramCounter = make([]int, 0)
	for functionNum := 0; functionNum < len(functionOffsets); functionNum++ {
		scriptData.StartProgramCounter = append(scriptData.StartProgramCounter, programCounter)

		var functionLength int64
		if functionNum != len(functionOffsets)-1 {
			functionLength = int64(functionOffsets[functionNum+1]) - int64(functionOffsets[functionNum])
		} else {
			functionLength = fileLength - int64(functionOffsets[functionNum])
		}

		streamReader = io.NewSectionReader(fileReader, int64(functionOffsets[functionNum]), functionLength)
		for lineNum := 0; lineNum < int(functionLength); lineNum++ {
			opcode := byte(0)
			if err := binary.Read(streamReader, binary.LittleEndian, &opcode); err != nil {
				return nil, err
			}

			byteSize, exists := InstructionSize[opcode]
			if !exists {
				fmt.Println("Unknown opcode:", opcode)
			}

			scriptLine, err := generateScriptLine(streamReader, byteSize, opcode)
			if err != nil {
				return nil, err
			}
			scriptData.Instructions[programCounter] = scriptLine

			// Sleep contains sleep and sleeping commands
			if opcode == OP_SLEEP {
				scriptData.Instructions[programCounter+1] = scriptData.Instructions[programCounter][1:]
			}

			programCounter += byteSize

			// return
			if opcode == OP_EVT_END {
				break
			}
		}
	}

	output := &SCDOutput{
		ScriptData: scriptData,
	}
	return output, nil
}

func generateScriptLine(streamReader *io.SectionReader, totalByteSize int, opcode byte) ([]byte, error) {
	scriptLine := make([]byte, 0)
	scriptLine = append(scriptLine, opcode)

	if totalByteSize == 1 {
		return scriptLine, nil
	}

	parameters, err := readRemainingBytes(streamReader, totalByteSize-1)
	if err != nil {
		return nil, fmt.Errorf("error reading script for opcode %d: %w", opcode, err)
	}
	scriptLine = append(scriptLine, parameters...)
	return scriptLine, nil
}

func readRemainingBytes(streamReader *io.SectionReader, byteSize int) ([]byte, error) {
	parameters := make([]byte, byteSize)
	if err := binary.Read(streamReader, binary.LittleEndian, &parameters); err != nil {
		return nil, err
	}
	return parameters, nil
}
