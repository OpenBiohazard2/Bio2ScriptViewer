package main

import (
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
		lineString := fmt.Sprintf("%s(", getFunctioNameFromOpcode(lineBytes[0]))
		for i := 1; i < len(lineBytes); i++ {
			lineString += fmt.Sprintf("%d, ", lineBytes[i])
		}
		lineString += ");"

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

func getFunctioNameFromOpcode(opcode byte) string {
	return fileio.FunctionName[opcode]
}
