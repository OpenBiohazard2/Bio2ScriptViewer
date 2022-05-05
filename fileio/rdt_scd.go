package fileio

// .scd - Script data

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

const (
	OP_NO_OP            = 0
	OP_EVT_END          = 1
	OP_EVT_NEXT         = 2
	OP_EVT_CHAIN        = 3
	OP_EVT_EXEC         = 4
	OP_EVT_KILL         = 5
	OP_IF_START         = 6
	OP_ELSE_START       = 7
	OP_END_IF           = 8
	OP_SLEEP            = 9
	OP_SLEEPING         = 10
	OP_WSLEEP           = 11
	OP_WSLEEPING        = 12
	OP_FOR              = 13
	OP_FOR_END          = 14
	OP_WHILE_START      = 15
	OP_WHILE_END        = 16
	OP_DO_START         = 17
	OP_DO_END           = 18
	OP_SWITCH           = 19
	OP_CASE             = 20
	OP_DEFAULT          = 21
	OP_END_SWITCH       = 22
	OP_GOTO             = 23
	OP_GOSUB            = 24
	OP_GOSUB_RETURN     = 25
	OP_BREAK            = 26
	OP_WORK_COPY        = 29
	OP_NO_OP2           = 32
	OP_CHECK            = 33
	OP_SET_BIT          = 34
	OP_COMPARE          = 35
	OP_SAVE             = 36
	OP_COPY             = 37
	OP_CALC             = 38
	OP_CALC2            = 39
	OP_SCE_RND          = 40
	OP_CUT_CHG          = 41
	OP_CUT_OLD          = 42
	OP_MESSAGE_ON       = 43
	OP_AOT_SET          = 44
	OP_OBJ_MODEL_SET    = 45
	OP_WORK_SET         = 46
	OP_SPEED_SET        = 47
	OP_ADD_SPEED        = 48
	OP_ADD_ASPEED       = 49
	OP_POS_SET          = 50
	OP_DIR_SET          = 51
	OP_MEMBER_SET       = 52
	OP_MEMBER_SET2      = 53
	OP_SE_ON            = 54
	OP_SCA_ID_SET       = 55
	OP_DIR_CK           = 57
	OP_SCE_ESPR_ON      = 58
	OP_DOOR_AOT_SET     = 59
	OP_CUT_AUTO         = 60
	OP_MEMBER_COPY      = 61
	OP_MEMBER_CMP       = 62
	OP_PLC_MOTION       = 63
	OP_PLC_DEST         = 64
	OP_PLC_NECK         = 65
	OP_PLC_RET          = 66
	OP_PLC_FLAG         = 67
	OP_SCE_EM_SET       = 68
	OP_AOT_RESET        = 70
	OP_AOT_ON           = 71
	OP_SUPER_SET        = 72
	OP_CUT_REPLACE      = 75
	OP_SCE_ESPR_KILL    = 76
	OP_DOOR_MODEL_SET   = 77
	OP_ITEM_AOT_SET     = 78
	OP_SCE_TRG_CK       = 80
	OP_SCE_BGM_CONTROL  = 81
	OP_SCE_ESPR_CONTROL = 82
	OP_SCE_FADE_SET     = 83
	OP_SCE_ESPR3D_ON    = 84
	OP_SCE_BGMTBL_SET   = 87
	OP_PLC_ROT          = 88
	OP_XA_ON            = 89
	OP_WEAPON_CHG       = 90
	OP_PLC_CNT          = 91
	OP_SCE_SHAKE_ON     = 92
	OP_MIZU_DIV_SET     = 93
	OP_KEEP_ITEM_CK     = 94
	OP_XA_VOL           = 95
	OP_KAGE_SET         = 96
	OP_CUT_BE_SET       = 97
	OP_SCE_ITEM_LOST    = 98
	OP_PLC_GUN_EFF      = 99
	OP_SCE_ESPR_ON2     = 100
	OP_SCE_ESPR_KILL2   = 101
	OP_PLC_STOP         = 102
	OP_AOT_SET_4P       = 103
	OP_DOOR_AOT_SET_4P  = 104
	OP_ITEM_AOT_SET_4P  = 105
	OP_LIGHT_POS_SET    = 106
	OP_LIGHT_KIDO_SET   = 107
	OP_RBJ_RESET        = 108
	OP_SCE_SCR_MOVE     = 109
	OP_PARTS_SET        = 110
	OP_MOVIE_ON         = 111
	OP_SCE_PARTS_BOMB   = 122
	OP_SCE_PARTS_DOWN   = 123
)

var (
	InstructionSize = map[byte]int{
		OP_NO_OP:            1,
		OP_EVT_END:          1,
		OP_EVT_NEXT:         1,
		OP_EVT_CHAIN:        4,
		OP_EVT_EXEC:         4,
		OP_EVT_KILL:         2,
		OP_IF_START:         4,
		OP_ELSE_START:       4,
		OP_END_IF:           1,
		OP_SLEEP:            4,
		OP_SLEEPING:         3,
		OP_WSLEEP:           1,
		OP_WSLEEPING:        1,
		OP_FOR:              6,
		OP_FOR_END:          2,
		OP_WHILE_START:      4,
		OP_WHILE_END:        2,
		OP_DO_START:         4,
		OP_DO_END:           2,
		OP_SWITCH:           4,
		OP_CASE:             6,
		OP_DEFAULT:          2,
		OP_END_SWITCH:       2,
		OP_GOTO:             6,
		OP_GOSUB:            2,
		OP_GOSUB_RETURN:     2,
		OP_BREAK:            2,
		OP_WORK_COPY:        4,
		OP_NO_OP2:           1,
		OP_CHECK:            4,
		OP_SET_BIT:          4,
		OP_COMPARE:          6,
		OP_SAVE:             4,
		OP_COPY:             3,
		OP_CALC:             6,
		OP_CALC2:            4,
		OP_SCE_RND:          1,
		OP_CUT_CHG:          2,
		OP_CUT_OLD:          1,
		OP_MESSAGE_ON:       6,
		OP_AOT_SET:          20,
		OP_OBJ_MODEL_SET:    38,
		OP_WORK_SET:         3,
		OP_SPEED_SET:        4,
		OP_ADD_SPEED:        1,
		OP_ADD_ASPEED:       1,
		OP_POS_SET:          8,
		OP_DIR_SET:          8,
		OP_MEMBER_SET:       4,
		OP_MEMBER_SET2:      3,
		OP_SE_ON:            12,
		OP_SCA_ID_SET:       4,
		OP_DIR_CK:           8,
		OP_SCE_ESPR_ON:      16,
		OP_DOOR_AOT_SET:     32,
		OP_CUT_AUTO:         2,
		OP_MEMBER_COPY:      3,
		OP_MEMBER_CMP:       6,
		OP_PLC_MOTION:       4,
		OP_PLC_DEST:         8,
		OP_PLC_NECK:         10,
		OP_PLC_RET:          1,
		OP_PLC_FLAG:         4,
		OP_SCE_EM_SET:       22,
		OP_AOT_RESET:        10,
		OP_AOT_ON:           2,
		OP_SUPER_SET:        16,
		OP_CUT_REPLACE:      3,
		OP_SCE_ESPR_KILL:    5,
		OP_DOOR_MODEL_SET:   22,
		OP_ITEM_AOT_SET:     22,
		OP_SCE_TRG_CK:       4,
		OP_SCE_BGM_CONTROL:  6,
		OP_SCE_ESPR_CONTROL: 6,
		OP_SCE_FADE_SET:     6,
		OP_SCE_ESPR3D_ON:    22,
		OP_SCE_BGMTBL_SET:   8,
		OP_PLC_ROT:          4,
		OP_XA_ON:            4,
		OP_WEAPON_CHG:       2,
		OP_PLC_CNT:          2,
		OP_SCE_SHAKE_ON:     3,
		OP_MIZU_DIV_SET:     2,
		OP_KEEP_ITEM_CK:     2,
		OP_XA_VOL:           2,
		OP_KAGE_SET:         14,
		OP_CUT_BE_SET:       4,
		OP_SCE_ITEM_LOST:    2,
		OP_PLC_GUN_EFF:      1,
		OP_SCE_ESPR_ON2:     16,
		OP_SCE_ESPR_KILL2:   2,
		OP_PLC_STOP:         1,
		OP_AOT_SET_4P:       28,
		OP_DOOR_AOT_SET_4P:  40,
		OP_ITEM_AOT_SET_4P:  30,
		OP_LIGHT_POS_SET:    6,
		OP_LIGHT_KIDO_SET:   4,
		OP_RBJ_RESET:        1,
		OP_SCE_SCR_MOVE:     4,
		OP_PARTS_SET:        6,
		OP_MOVIE_ON:         2,
		OP_SCE_PARTS_BOMB:   16,
		OP_SCE_PARTS_DOWN:   16,
	}
	FunctionName = map[byte]string{
		OP_NO_OP:            "NoOp",
		OP_EVT_END:          "EvtEnd",
		OP_EVT_NEXT:         "EvtNext",
		OP_EVT_CHAIN:        "EvtChain",
		OP_EVT_EXEC:         "EvtExec",
		OP_EVT_KILL:         "EvtKill",
		OP_IF_START:         "IfStart",
		OP_ELSE_START:       "ElseStart",
		OP_END_IF:           "EndIf",
		OP_SLEEP:            "Sleep",
		OP_SLEEPING:         "Sleeping",
		OP_WSLEEP:           "Wsleep",
		OP_WSLEEPING:        "Wsleeping",
		OP_FOR:              "ForStart",
		OP_FOR_END:          "ForEnd",
		OP_WHILE_START:      "WhileStart",
		OP_WHILE_END:        "WhileEnd",
		OP_DO_START:         "DoStart",
		OP_DO_END:           "DoEnd",
		OP_SWITCH:           "Switch",
		OP_CASE:             "Case",
		OP_DEFAULT:          "Default",
		OP_END_SWITCH:       "EndSwitch",
		OP_GOTO:             "Goto",
		OP_GOSUB:            "Gosub",
		OP_GOSUB_RETURN:     "GosubReturn",
		OP_BREAK:            "Break",
		OP_WORK_COPY:        "WorkCopy",
		OP_NO_OP2:           "NoOp2",
		OP_CHECK:            "CheckBit",
		OP_SET_BIT:          "SetBit",
		OP_COMPARE:          "Compare",
		OP_SAVE:             "Save",
		OP_COPY:             "Copy",
		OP_CALC:             "Calc",
		OP_CALC2:            "Calc2",
		OP_SCE_RND:          "SceRnd",
		OP_CUT_CHG:          "CutChg",
		OP_CUT_OLD:          "CutOld",
		OP_MESSAGE_ON:       "MessageOn",
		OP_AOT_SET:          "AotSet",
		OP_OBJ_MODEL_SET:    "ObjModelSet",
		OP_WORK_SET:         "WorkSet",
		OP_SPEED_SET:        "SpeedSet",
		OP_ADD_SPEED:        "AddSpeed",
		OP_ADD_ASPEED:       "AddAspeed",
		OP_POS_SET:          "PosSet",
		OP_DIR_SET:          "DirSet",
		OP_MEMBER_SET:       "MemberSet",
		OP_MEMBER_SET2:      "MemberSet2",
		OP_SE_ON:            "SeOn",
		OP_SCA_ID_SET:       "ScaIdSet",
		OP_DIR_CK:           "DirCk",
		OP_SCE_ESPR_ON:      "SceEsprOn",
		OP_DOOR_AOT_SET:     "DoorAotSet",
		OP_CUT_AUTO:         "CutAuto",
		OP_MEMBER_COPY:      "MemberCopy",
		OP_MEMBER_CMP:       "MemberCmp",
		OP_PLC_MOTION:       "PlcMotion",
		OP_PLC_DEST:         "PlcDest",
		OP_PLC_NECK:         "PlcNeck",
		OP_PLC_RET:          "PlcRet",
		OP_PLC_FLAG:         "PlcFlag",
		OP_SCE_EM_SET:       "SceEmSet",
		OP_AOT_RESET:        "AotReset",
		OP_AOT_ON:           "AotOn",
		OP_SUPER_SET:        "SuperSet",
		OP_CUT_REPLACE:      "CutReplace",
		OP_SCE_ESPR_KILL:    "SceEsprKill",
		OP_DOOR_MODEL_SET:   "DoorModelSet",
		OP_ITEM_AOT_SET:     "ItemAotSet",
		OP_SCE_TRG_CK:       "SceTrgCk",
		OP_SCE_BGM_CONTROL:  "SceBgmControl",
		OP_SCE_ESPR_CONTROL: "SceEsprControl",
		OP_SCE_FADE_SET:     "SceFadeSet",
		OP_SCE_ESPR3D_ON:    "SceEspr3dOn",
		OP_SCE_BGMTBL_SET:   "SceBgmTblSet",
		OP_PLC_ROT:          "PlcRot",
		OP_XA_ON:            "XaOn",
		OP_WEAPON_CHG:       "WeaponChg",
		OP_PLC_CNT:          "PlcCnt",
		OP_SCE_SHAKE_ON:     "SceShakeOn",
		OP_MIZU_DIV_SET:     "MizuDivSet",
		OP_KEEP_ITEM_CK:     "KeepItemCk",
		OP_XA_VOL:           "XaVol",
		OP_KAGE_SET:         "KageSet",
		OP_CUT_BE_SET:       "CutBeSet",
		OP_SCE_ITEM_LOST:    "SceItemLost",
		OP_PLC_GUN_EFF:      "PlcGunEff",
		OP_SCE_ESPR_ON2:     "SceEsprOn2",
		OP_SCE_ESPR_KILL2:   "SceEsprKill2",
		OP_PLC_STOP:         "PlcStop",
		OP_AOT_SET_4P:       "AotSet4P",
		OP_DOOR_AOT_SET_4P:  "DoorAotSet4P",
		OP_ITEM_AOT_SET_4P:  "ItemAotSet4P",
		OP_LIGHT_POS_SET:    "LightPosSet",
		OP_LIGHT_KIDO_SET:   "LightKidoSet",
		OP_RBJ_RESET:        "RbjReset",
		OP_SCE_SCR_MOVE:     "SceScrMove",
		OP_PARTS_SET:        "PartsSet",
		OP_MOVIE_ON:         "MovieOn",
		OP_SCE_PARTS_BOMB:   "ScePartsBomb",
		OP_SCE_PARTS_DOWN:   "ScePartsDown",
	}
)

type ScriptInstrGoSub struct {
	Opcode uint8 // 0x18
	Event  uint8
}

type ScriptInstrCheckBitTest struct {
	Opcode    uint8 // 0x21
	BitArray  uint8 // Index of array of bits to use
	BitNumber uint8 // Bit number to check
	Value     uint8 // Value to compare (0 or 1)
}

type ScriptInstrSetBit struct {
	Opcode    uint8 // 0x22
	BitArray  uint8 // Index of array of bits to use
	BitNumber uint8 // Bit number to check
	Operation uint8 // 0x0: clear, 0x1: set, 0x2-0x6: invalid, 0x7: flip bit
}

type ScriptInstrCutChg struct {
	Opcode   uint8 // 0x29
	CameraId uint8
}

type ScriptInstrAotSet struct {
	Opcode       uint8 // 0x2c
	Aot          uint8
	Id           uint8
	Type         uint8
	Floor        uint8
	Super        uint8
	X, Z         int16
	Width, Depth int16
	Data         [6]uint8
}

type ScriptInstrObjModelSet struct {
	Opcode      uint8 // 0x2d
	ObjectIndex uint8
	ObjectId    uint8
	Counter     uint8
	Wait        uint8
	Num         uint8
	Floor       uint8
	Flag0       uint8
	Type        uint16
	Flag1       uint16
	Attribute   int16
	Position    [3]int16
	Direction   [3]int16
	Offset      [3]int16
	Dimensions  [3]uint16
}

type ScriptInstrPosSet struct {
	Opcode uint8 // 0x32
	Dummy  uint8
	X      int16
	Y      int16
	Z      int16
}

type ScriptInstrSceEsprOn struct {
	Opcode   uint8 // 0x3a
	Dummy    uint8
	Id       uint8
	Type     uint8
	Work     uint16
	Unknown1 int16
	X, Y, Z  int16
	DirY     uint16
}

type ScriptInstrDoorAotSet struct {
	Opcode                       uint8 // 0x3b
	Aot                          uint8 // Index of item in array of room objects list
	Id                           uint8
	Type                         uint8
	Floor                        uint8
	Super                        uint8
	X, Z                         int16 // Location of door
	Width, Depth                 int16 // Size of door
	NextX, NextY, NextZ, NextDir int16 // Position and direction of player after door entered
	Stage, Room, Camera          uint8 // Stage, room, camera after door entered
	NextFloor                    uint8
	TextureType                  uint8
	DoorType                     uint8
	KnockType                    uint8
	KeyId                        uint8
	KeyType                      uint8
	Free                         uint8
}

type ScriptInstrPlcNeck struct {
	Opcode    uint8 // 0x41
	Operation uint8
	NeckX     int16
	NeckY     int16
	NeckZ     int16
	Unknown   [2]int8
}

type ScriptInstrSceEmSet struct {
	Opcode    uint8 // 0x44
	Dummy     uint8
	Aot       uint8
	Id        uint8
	Type      uint8
	Status    uint8
	Floor     uint8
	SoundFlag uint8
	ModelType uint8
	EmSetFlag int8
	X, Y, Z   int16
	DirY      uint16
	Motion    uint16
	CtrFlag   uint16
}

type ScriptInstrAotReset struct {
	Opcode uint8 // 0x46
	Aot    uint8
	Id     uint8
	Type   uint8
	Data   [6]uint8
}

type ScriptInstrItemAotSet struct {
	Opcode          uint8 // 0x4e
	Aot             uint8
	Id              uint8
	Type            uint8
	Floor           uint8
	Super           uint8
	X, Z            int16
	Width, Depth    int16
	ItemId          uint16
	Amount          uint16
	ItemPickedIndex uint16 // flag to check if item is picked up
	Md1ModelId      uint8
	Act             uint8
}

type ScriptInstrSceBgmControl struct {
	Opcode      uint8 // 0x51
	Id          uint8 // 0: Main, 1: sub0, 2: sub1
	Operation   uint8 // 0: nop, 1: start, 2: stop, 3: restart, 4: pause, 5: fadeout
	Type        uint8 // 0: MAIN_VOL, 1: PROG0_VOL, 2: PROG1_VOL, 3: PROG2_VOL
	LeftVolume  uint8
	RightVolume uint8
}

type ScriptInstrAotSet4p struct {
	Opcode uint8 // 0x67
	Aot    uint8
	Id     uint8
	Type   uint8
	Floor  uint8
	Super  uint8
	X1, Z1 int16
	X2, Z2 int16
	X3, Z3 int16
	X4, Z4 int16
	Data   [6]uint8
}

type SCDOutput struct {
	ScriptData ScriptFunction
}

type ScriptFunction struct {
	Instructions        map[int][]byte // key is program counter, value is command
	StartProgramCounter []int          // set per function
}

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

			scriptData.Instructions[programCounter] = generateScriptLine(streamReader, byteSize, opcode)

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

func generateScriptLine(streamReader *io.SectionReader, totalByteSize int, opcode byte) []byte {
	scriptLine := make([]byte, 0)
	scriptLine = append(scriptLine, opcode)

	if totalByteSize == 1 {
		return scriptLine
	}

	parameters, err := readRemainingBytes(streamReader, totalByteSize-1)
	if err != nil {
		log.Fatal("Error reading script for opcode %v\n", opcode)
	}
	scriptLine = append(scriptLine, parameters...)
	return scriptLine
}

func readRemainingBytes(streamReader *io.SectionReader, byteSize int) ([]byte, error) {
	parameters := make([]byte, byteSize)
	if err := binary.Read(streamReader, binary.LittleEndian, &parameters); err != nil {
		return nil, err
	}
	return parameters, nil
}
