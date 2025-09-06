package fileio

// .rdt - Room data

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type RDTHeader struct {
	NumSprites uint8
	NumCameras uint8
	NumModels  uint8
	NumItems   uint8
	NumDoors   uint8
	NumRooms   uint8
	NumReverb  uint8 // related to sound
	SpriteMax  uint8 // max number of .pri sprites used by one of the room's cameras
}

type RDTOffsets struct {
	OffsetRoomSound              uint32 // offset to room .snd sound table data
	OffsetRoomVABHeader          uint32 // .vh file
	OffsetRoomVABData            uint32 // .vb file
	OffsetEnemyVABHeader         uint32 // .vh file
	OffsetEnemyVABData           uint32 // .vb file
	OffsetOTA                    uint32
	OffsetCollisionData          uint32 // .sca file
	OffsetCameraPosition         uint32 // .rid file
	OffsetCameraSwitches         uint32 // .rvd file
	OffsetLights                 uint32 // .lit file
	OffsetItems                  uint32
	OffsetFloorSound             uint32 // .flr file
	OffsetBlocks                 uint32 // .blk file
	OffsetLang1                  uint32 // .msg file
	OffsetLang2                  uint32 // .msg file
	OffsetScrollTexture          uint32 // .tim file
	OffsetInitScript             uint32 // .scd file
	OffsetExecuteScript          uint32 // .scd file
	OffsetSpriteAnimations       uint32 // .esp file
	OffsetSpriteAnimationsOffset uint32 // .esp file
	OffsetSpriteImage            uint32 // .tim file
	OffsetModelImage             uint32 // .tim file
	OffsetRBJ                    uint32 // .rbj file
}

type RDTOutput struct {
	InitScriptData *SCDOutput
	RoomScriptData *SCDOutput
}

func LoadRDTFile(filename string) (*RDTOutput, error) {
	rdtFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open RDT file %s: %w", filename, err)
	}
	defer rdtFile.Close()

	fi, err := rdtFile.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info for %s: %w", filename, err)
	}

	fileLength := fi.Size()
	return LoadRDT(rdtFile, fileLength)
}

func LoadRDT(r io.ReaderAt, fileLength int64) (*RDTOutput, error) {
	reader := io.NewSectionReader(r, int64(0), fileLength)

	rdtHeader := RDTHeader{}
	if err := binary.Read(reader, binary.LittleEndian, &rdtHeader); err != nil {
		return nil, err
	}

	offsets := RDTOffsets{}
	if err := binary.Read(reader, binary.LittleEndian, &offsets); err != nil {
		return nil, err
	}

	// Script data
	// Run once when the level loads
	offset := int64(offsets.OffsetInitScript)
	initSCDReader := io.NewSectionReader(r, offset, fileLength-offset)
	initSCDOutput, err := LoadRDT_SCDStream(initSCDReader, fileLength)
	if err != nil {
		return nil, err
	}

	// Run during the game
	offset = int64(offsets.OffsetExecuteScript)
	roomSCDReader := io.NewSectionReader(r, offset, fileLength-offset)
	roomSCDOutput, err := LoadRDT_SCDStream(roomSCDReader, fileLength)
	if err != nil {
		return nil, err
	}

	output := &RDTOutput{
		InitScriptData: initSCDOutput,
		RoomScriptData: roomSCDOutput,
	}
	return output, nil
}
