package ui

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"github.com/OpenBiohazard2/Bio2ScriptViewer/fileio"
)

func (a *App) openFileDialog() {
	dialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, a.mainWin)
			return
		}
		if err == nil && reader == nil {
			return
		}

		file, err := os.Open(reader.URI().String()[7:])
		if err != nil {
			dialog.ShowError(err, a.mainWin)
			return
		}

		err = a.open(file, true)
		if err != nil {
			dialog.ShowError(err, a.mainWin)
			return
		}
		defer reader.Close()
	}, a.mainWin)
	dialog.SetFilter(storage.NewExtensionFileFilter([]string{".rdt"}))
	dialog.Show()
}

func (a *App) open(file *os.File, folder bool) error {
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err
	}
	fileLength := fi.Size()

	streamReader := io.NewSectionReader(file, int64(0), fileLength)
	rdtOutput, err := fileio.LoadRDT(streamReader, fileLength)
	if err != nil {
		return err
	}

	scriptFiles := splitScriptDataIntoFiles(rdtOutput.RoomScriptData)
	// Add script from init
	scriptFiles["init.scd"] = convertInitialScriptIntoFile(rdtOutput.InitScriptData)

	filenames := make([]string, 0)
	for filename, _ := range scriptFiles {
		filenames = append(filenames, filename)
	}
	sort.Strings(filenames)

	layout := container.NewBorder(nil, a.loadStatusBar(), a.loadFileList(filenames, scriptFiles), nil, a.split)
	a.mainWin.SetContent(layout)

	return nil
}

func convertInitialScriptIntoFile(scriptFile *fileio.SCDOutput) [][]byte {
	programCounters := sortProgramCounters(scriptFile.ScriptData.Instructions)

	fileLines := make([][]byte, 0)
	for _, programCounter := range programCounters {
		fileLines = append(fileLines, scriptFile.ScriptData.Instructions[programCounter])
	}
	return fileLines
}

func splitScriptDataIntoFiles(scriptFile *fileio.SCDOutput) map[string][][]byte {
	programCounters := sortProgramCounters(scriptFile.ScriptData.Instructions)

	startCounterExists := make(map[int]bool)
	for _, start := range scriptFile.ScriptData.StartProgramCounter {
		startCounterExists[start] = true
	}

	scriptFiles := make(map[string][][]byte)
	fileLines := make([][]byte, 0)
	fileIndex := 0
	for _, programCounter := range programCounters {
		_, ok := startCounterExists[programCounter]
		if ok && programCounter > 0 {
			scriptFiles[fmt.Sprintf("sub%d.scd", fileIndex)] = fileLines
			fileIndex++
			fileLines = make([][]byte, 0)
		}

		fileLines = append(fileLines, scriptFile.ScriptData.Instructions[programCounter])
	}

	// Add last script function
	if len(fileLines) > 0 {
		scriptFiles[fmt.Sprintf("sub%d.scd", fileIndex)] = fileLines
	}

	return scriptFiles
}

func convertRawScriptInstructionsToString(instructions [][]byte) string {
	rawDataString := ""
	for _, lineBytes := range instructions {
		lineString := ""
		// print out hex values
		for i := 0; i < len(lineBytes); i++ {
			lineString += fmt.Sprintf("%02x ", lineBytes[i])
		}
		rawDataString += lineString + "\n"
	}

	return rawDataString
}

func convertScriptInstructionsToCode(instructions [][]byte) string {
	rawDataString := ""
	for _, lineBytes := range instructions {
		lineString := fmt.Sprintf("%s", getFunctionNameFromOpcode(lineBytes[0]))
		lineString += showParameters(lineBytes)
		rawDataString += lineString + "\n"
	}

	return rawDataString
}

func sortProgramCounters(instructions map[int][]byte) []int {
	// sort script commands in order
	programCounters := make([]int, 0)
	for counter, _ := range instructions {
		programCounters = append(programCounters, counter)
	}
	sort.Ints(programCounters)

	return programCounters
}

func getFunctionNameFromOpcode(opcode byte) string {
	return fileio.FunctionName[opcode]
}

func showParameters(lineBytes []byte) string {
	opcode := lineBytes[0]
	parameterString := "("
	switch opcode {
	case fileio.OP_GOSUB: // 0x18
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrGoSub{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("Event=%d", instruction.Event)
	case fileio.OP_CHECK: // 0x21
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrCheckBitTest{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("BitArray=%d, ", instruction.BitArray)
		parameterString += fmt.Sprintf("BitNumber=%d, ", instruction.BitNumber)
		parameterString += fmt.Sprintf("Value=%d", instruction.Value)
	case fileio.OP_SET_BIT: // 0x22
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrSetBit{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("BitArray=%d, ", instruction.BitArray)
		parameterString += fmt.Sprintf("BitNumber=%d, ", instruction.BitNumber)
		parameterString += fmt.Sprintf("Operation=%d", instruction.Operation)
	case fileio.OP_CUT_CHG: // 0x29
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrCutChg{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("CameraId=%d", instruction.CameraId)
	case fileio.OP_AOT_SET: // 0x2c
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrAotSet{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("Aot=%d, ", instruction.Aot)
		parameterString += fmt.Sprintf("Id=%d, ", instruction.Id)
		parameterString += fmt.Sprintf("Type=%d, ", instruction.Type)
		parameterString += fmt.Sprintf("Floor=%d, ", instruction.Floor)
		parameterString += fmt.Sprintf("Super=%d, ", instruction.Super)
		parameterString += fmt.Sprintf("X=%d, Z=%d, ", instruction.X, instruction.Z)
		parameterString += fmt.Sprintf("Width=%d, Depth=%d, ", instruction.Width, instruction.Depth)
		parameterString += fmt.Sprintf("Data=[%d,%d,%d,%d,%d,%d]", instruction.Data[0], instruction.Data[1], instruction.Data[2],
			instruction.Data[3], instruction.Data[4], instruction.Data[5])
	case fileio.OP_OBJ_MODEL_SET: // 0x2d
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrObjModelSet{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("ObjectIndex=%d, ", instruction.ObjectIndex)
		parameterString += fmt.Sprintf("ObjectId=%d, ", instruction.ObjectId)
		parameterString += fmt.Sprintf("Counter=%d, ", instruction.Counter)
		parameterString += fmt.Sprintf("Wait=%d, ", instruction.Wait)
		parameterString += fmt.Sprintf("Num=%d, ", instruction.Num)
		parameterString += fmt.Sprintf("Floor=%d, ", instruction.Floor)
		parameterString += fmt.Sprintf("Flag0=%d, ", instruction.Flag0)
		parameterString += fmt.Sprintf("Type=%d, ", instruction.Type)
		parameterString += fmt.Sprintf("Flag1=%d, ", instruction.Flag1)
		parameterString += fmt.Sprintf("Attribute=%d, ", instruction.Attribute)
		parameterString += fmt.Sprintf("Position=[%d, %d, %d], ", instruction.Position[0], instruction.Position[1], instruction.Position[2])
		parameterString += fmt.Sprintf("Direction=[%d, %d, %d], ", instruction.Direction[0], instruction.Direction[1], instruction.Direction[2])
		parameterString += fmt.Sprintf("Offset=[%d, %d, %d], ", instruction.Offset[0], instruction.Offset[1], instruction.Offset[2])
		parameterString += fmt.Sprintf("Dimensions=[%d, %d, %d]", instruction.Dimensions[0], instruction.Dimensions[1], instruction.Dimensions[2])
	case fileio.OP_POS_SET: // 0x32
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrPosSet{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("Dummy=%d, ", instruction.Dummy)
		parameterString += fmt.Sprintf("X=%d, Y=%d, Z=%d", instruction.X, instruction.Y, instruction.Z)
	case fileio.OP_SCE_ESPR_ON: // 0x3a
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrSceEsprOn{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("Dummy=%d, ", instruction.Dummy)
		parameterString += fmt.Sprintf("Id=%d, ", instruction.Id)
		parameterString += fmt.Sprintf("Type=%d, ", instruction.Type)
		parameterString += fmt.Sprintf("Work=%d, ", instruction.Work)
		parameterString += fmt.Sprintf("Unknown1=%d, ", instruction.Unknown1)
		parameterString += fmt.Sprintf("X=%d, Y=%d, Z=%d, ", instruction.X, instruction.Y, instruction.Z)
		parameterString += fmt.Sprintf("DirY=%d", instruction.DirY)
	case fileio.OP_DOOR_AOT_SET: // 0x3b
		byteArr := bytes.NewBuffer(lineBytes)
		door := fileio.ScriptInstrDoorAotSet{}
		binary.Read(byteArr, binary.LittleEndian, &door)
		parameterString += fmt.Sprintf("Aot=%d, ", door.Aot)
		parameterString += fmt.Sprintf("Id=%d, ", door.Id)
		parameterString += fmt.Sprintf("Type=%d, ", door.Type)
		parameterString += fmt.Sprintf("Floor=%d, ", door.Floor)
		parameterString += fmt.Sprintf("Super=%d, ", door.Super)
		parameterString += fmt.Sprintf("X=%d, Z=%d, ", door.X, door.Z)
		parameterString += fmt.Sprintf("Width=%d, Depth=%d, ", door.Width, door.Depth)
		parameterString += fmt.Sprintf("NextX=%d, NextY=%d, ", door.NextX, door.NextY)
		parameterString += fmt.Sprintf("NextZ=%d, NextDir=%d, ", door.NextZ, door.NextDir)
		parameterString += fmt.Sprintf("Stage=%d, Room=%d, Camera=%d, ", door.Stage, door.Room, door.Camera)
		parameterString += fmt.Sprintf("NextFloor=%d, ", door.NextFloor)
		parameterString += fmt.Sprintf("TextureType=%d, ", door.TextureType)
		parameterString += fmt.Sprintf("DoorType=%d, ", door.DoorType)
		parameterString += fmt.Sprintf("KnockType=%d, ", door.KnockType)
		parameterString += fmt.Sprintf("KeyId=%d, ", door.KeyId)
		parameterString += fmt.Sprintf("KeyType=%d, ", door.KeyType)
		parameterString += fmt.Sprintf("Free=%d", door.Free)
	case fileio.OP_PLC_NECK: // 0x41
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrPlcNeck{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("Operation=%d, ", instruction.Operation)
		parameterString += fmt.Sprintf("NeckX=%d, ", instruction.NeckX)
		parameterString += fmt.Sprintf("NeckY=%d, ", instruction.NeckY)
		parameterString += fmt.Sprintf("NeckZ=%d, ", instruction.NeckZ)
		parameterString += fmt.Sprintf("Unknown=[%d, %d]", instruction.Unknown[0], instruction.Unknown[1])
	case fileio.OP_SCE_EM_SET: // 0x44
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrSceEmSet{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("Dummy=%d, ", instruction.Dummy)
		parameterString += fmt.Sprintf("Aot=%d, ", instruction.Aot)
		parameterString += fmt.Sprintf("Id=%d, ", instruction.Id)
		parameterString += fmt.Sprintf("Type=%d, ", instruction.Type)
		parameterString += fmt.Sprintf("Status=%d, ", instruction.Status)
		parameterString += fmt.Sprintf("Floor=%d, ", instruction.Floor)
		parameterString += fmt.Sprintf("SoundFlag=%d, ", instruction.SoundFlag)
		parameterString += fmt.Sprintf("ModelType=%d, ", instruction.ModelType)
		parameterString += fmt.Sprintf("EmSetFlag=%d, ", instruction.EmSetFlag)
		parameterString += fmt.Sprintf("X=%d, Y=%d, Z=%d, ", instruction.X, instruction.Y, instruction.Z)
		parameterString += fmt.Sprintf("DirY=%d, ", instruction.DirY)
		parameterString += fmt.Sprintf("Motion=%d, ", instruction.Motion)
		parameterString += fmt.Sprintf("CtrFlag=%d", instruction.CtrFlag)
	case fileio.OP_AOT_RESET: // 0x46
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrAotReset{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("Aot=%d, ", instruction.Aot)
		parameterString += fmt.Sprintf("Id=%d, ", instruction.Id)
		parameterString += fmt.Sprintf("Type=%d, ", instruction.Type)
		parameterString += fmt.Sprintf("Data=[%d,%d,%d,%d,%d,%d]", instruction.Data[0], instruction.Data[1], instruction.Data[2],
			instruction.Data[3], instruction.Data[4], instruction.Data[5])
	case fileio.OP_ITEM_AOT_SET: // 0x4e
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrItemAotSet{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("Aot=%d, ", instruction.Aot)
		parameterString += fmt.Sprintf("Id=%d, ", instruction.Id)
		parameterString += fmt.Sprintf("Type=%d, ", instruction.Type)
		parameterString += fmt.Sprintf("Floor=%d, ", instruction.Floor)
		parameterString += fmt.Sprintf("Super=%d, ", instruction.Super)
		parameterString += fmt.Sprintf("X=%d, Z=%d, ", instruction.X, instruction.Z)
		parameterString += fmt.Sprintf("Width=%d, Depth=%d, ", instruction.Width, instruction.Depth)
		parameterString += fmt.Sprintf("ItemId=%d, ", instruction.ItemId)
		parameterString += fmt.Sprintf("Amount=%d, ", instruction.Amount)
		parameterString += fmt.Sprintf("ItemPickedIndex=%d, ", instruction.ItemPickedIndex)
		parameterString += fmt.Sprintf("Md1ModelId=%d, ", instruction.Md1ModelId)
		parameterString += fmt.Sprintf("Act=%d", instruction.Act)
	case fileio.OP_SCE_BGM_CONTROL: // 0x51
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrSceBgmControl{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("Id=%d, ", instruction.Id)
		parameterString += fmt.Sprintf("Operation=%d, ", instruction.Operation)
		parameterString += fmt.Sprintf("Type=%d, ", instruction.Type)
		parameterString += fmt.Sprintf("LeftVolume=%d, ", instruction.LeftVolume)
		parameterString += fmt.Sprintf("RightVolume=%d", instruction.RightVolume)
	case fileio.OP_AOT_SET_4P: // 0x67
		byteArr := bytes.NewBuffer(lineBytes)
		instruction := fileio.ScriptInstrAotSet4p{}
		binary.Read(byteArr, binary.LittleEndian, &instruction)
		parameterString += fmt.Sprintf("Aot=%d, ", instruction.Aot)
		parameterString += fmt.Sprintf("Id=%d, ", instruction.Id)
		parameterString += fmt.Sprintf("Type=%d, ", instruction.Type)
		parameterString += fmt.Sprintf("Floor=%d, ", instruction.Floor)
		parameterString += fmt.Sprintf("Super=%d, ", instruction.Super)
		parameterString += fmt.Sprintf("X1=%d, Z1=%d, ", instruction.X1, instruction.Z1)
		parameterString += fmt.Sprintf("X2=%d, Z2=%d, ", instruction.X2, instruction.Z2)
		parameterString += fmt.Sprintf("X3=%d, Z3=%d, ", instruction.X3, instruction.Z3)
		parameterString += fmt.Sprintf("X4=%d, Z4=%d, ", instruction.X4, instruction.Z4)
		parameterString += fmt.Sprintf("Data=[%d,%d,%d,%d,%d,%d]", instruction.Data[0], instruction.Data[1], instruction.Data[2],
			instruction.Data[3], instruction.Data[4], instruction.Data[5])
	default:
		// Log each byte as its own parameter
		for i := 1; i < len(lineBytes); i++ {
			parameterString += fmt.Sprintf("%d, ", lineBytes[i])
		}
	}
	parameterString += ");"
	return parameterString
}
