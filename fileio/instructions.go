package fileio

// Script instruction struct definitions for Resident Evil 2 / Biohazard 2
// These structs represent the binary layout of different script commands

// ScriptInstrEventExec represents an EVT_EXEC instruction (0x04)
type ScriptInstrEventExec struct {
	Opcode    uint8 // 0x04
	ThreadNum uint8
	ExOpcode  uint8
	Event     uint8
}

// ScriptInstrIfElseStart represents an IF_START instruction (0x06)
type ScriptInstrIfElseStart struct {
	Opcode      uint8 // 0x06
	Dummy       uint8
	BlockLength uint16
}

// ScriptInstrElseStart represents an ELSE_START instruction (0x07)
type ScriptInstrElseStart struct {
	Opcode      uint8 // 0x07
	Dummy       uint8
	BlockLength uint16
}

// ScriptInstrSleep represents a SLEEP instruction (0x09)
type ScriptInstrSleep struct {
	Opcode uint8 // 0x09
	Dummy  uint8
	Count  uint16
}

// ScriptInstrForStart represents a FOR instruction (0x0d)
type ScriptInstrForStart struct {
	Opcode      uint8 // 0x0d
	Dummy       uint8
	BlockLength uint16
	Count       uint16
}

// ScriptInstrSwitch represents a SWITCH instruction (0x13)
type ScriptInstrSwitch struct {
	Opcode      uint8 // 0x13
	VarId       uint8
	BlockLength uint16
}

// ScriptInstrSwitchCase represents a CASE instruction (0x14)
type ScriptInstrSwitchCase struct {
	Opcode      uint8 // 0x14
	Dummy       uint8
	BlockLength uint16
	Value       uint16
}

// ScriptInstrGoto represents a GOTO instruction (0x17)
type ScriptInstrGoto struct {
	Opcode        uint8 // 0x17
	IfElseCounter int8
	LoopLevel     int8
	Unknown       uint8
	Offset        int16
}

// ScriptInstrGoSub represents a GOSUB instruction (0x18)
type ScriptInstrGoSub struct {
	Opcode uint8 // 0x18
	Event  uint8
}

// ScriptInstrCheckBitTest represents a CHECK instruction (0x21)
type ScriptInstrCheckBitTest struct {
	Opcode    uint8 // 0x21
	BitArray  uint8 // Index of array of bits to use
	BitNumber uint8 // Bit number to check
	Value     uint8 // Value to compare (0 or 1)
}

// ScriptInstrSetBit represents a SET_BIT instruction (0x22)
type ScriptInstrSetBit struct {
	Opcode    uint8 // 0x22
	BitArray  uint8 // Index of array of bits to use
	BitNumber uint8 // Bit number to check
	Operation uint8 // 0x0: clear, 0x1: set, 0x2-0x6: invalid, 0x7: flip bit
}

// ScriptInstrCompare represents a COMPARE instruction (0x23)
type ScriptInstrCompare struct {
	Opcode    uint8 // 0x23
	Dummy     uint8
	VarId     uint8
	Operation uint8
	Value     int16 // Value to compare against
}

// ScriptInstrSave represents a SAVE instruction (0x24)
type ScriptInstrSave struct {
	Opcode uint8 // 0x24
	VarId  uint8
	Value  int16
}

// ScriptInstrCopy represents a COPY instruction (0x25)
type ScriptInstrCopy struct {
	Opcode      uint8 // 0x25
	DestVarId   uint8
	SourceVarId uint8
}

// ScriptInstrCalc represents a CALC instruction (0x26)
type ScriptInstrCalc struct {
	Opcode    uint8 // 0x26
	Dummy     uint8
	Operation uint8
	VarId     uint8
	Value     uint8
}

// ScriptInstrCalc2 represents a CALC2 instruction (0x27)
type ScriptInstrCalc2 struct {
	Opcode      uint8 // 0x27
	Operation   uint8
	VarId       uint8
	SourceVarId uint8
}

// ScriptInstrCutChg represents a CUT_CHG instruction (0x29)
type ScriptInstrCutChg struct {
	Opcode   uint8 // 0x29
	CameraId uint8
}

// ScriptInstrAotSet represents an AOT_SET instruction (0x2c)
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

// ScriptInstrObjModelSet represents an OBJ_MODEL_SET instruction (0x2d)
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

// ScriptInstrWorkSet represents a WORK_SET instruction (0x2e)
type ScriptInstrWorkSet struct {
	Opcode    uint8 // 0x2e
	Component uint8
	Index     uint8
}

// ScriptInstrPosSet represents a POS_SET instruction (0x32)
type ScriptInstrPosSet struct {
	Opcode uint8 // 0x32
	Dummy  uint8
	X      int16
	Y      int16
	Z      int16
}

// ScriptInstrMemberSet represents a MEMBER_SET instruction (0x34)
type ScriptInstrMemberSet struct {
	Opcode      uint8 // 0x34
	MemberIndex uint8
	Value       uint16
}

// ScriptInstrScaIdSet represents a SCA_ID_SET instruction (0x37)
type ScriptInstrScaIdSet struct {
	Opcode uint8 // 0x37
	Id     uint8
	Flag   uint16
}

// ScriptInstrSceEsprOn represents a SCE_ESPR_ON instruction (0x3a)
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

// ScriptInstrDoorAotSet represents a DOOR_AOT_SET instruction (0x3b)
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

// ScriptInstrCutAuto represents a CUT_AUTO instruction (0x3c)
type ScriptInstrCutAuto struct {
	Opcode uint8 // 0x3c
	FlagOn uint8
}

// ScriptInstrMemberCompare represents a MEMBER_CMP instruction (0x3e)
type ScriptInstrMemberCompare struct {
	Opcode           uint8 // 0x3e
	Unknown0         uint8
	MemberIndex      uint8
	CompareOperation uint8 // 0 - 6
	Value            int16
}

// ScriptInstrPlcMotion represents a PLC_MOTION instruction (0x3f)
type ScriptInstrPlcMotion struct {
	Opcode     uint8 // 0x3f
	Action     uint8
	MoveNumber uint8
	SceneFlag  uint8
}

// ScriptInstrPlcDest represents a PLC_DEST instruction (0x40)
type ScriptInstrPlcDest struct {
	Opcode     uint8 // 0x40
	Dummy      uint8
	Action     uint8
	FlagNumber uint8
	DestX      int16
	DestZ      int16
}

// ScriptInstrPlcNeck represents a PLC_NECK instruction (0x41)
type ScriptInstrPlcNeck struct {
	Opcode    uint8 // 0x41
	Operation uint8
	NeckX     int16
	NeckY     int16
	NeckZ     int16
	Unknown   [2]int8
}

// ScriptInstrPlcFlag represents a PLC_FLAG instruction (0x43)
type ScriptInstrPlcFlag struct {
	Opcode    uint8 // 0x43
	Operation uint8 // 0: OR, 1: Set, 2: XOR
	Flag      uint16
}

// ScriptInstrSceEmSet represents a SCE_EM_SET instruction (0x44)
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

// ScriptInstrAotReset represents an AOT_RESET instruction (0x46)
type ScriptInstrAotReset struct {
	Opcode uint8 // 0x46
	Aot    uint8
	Id     uint8
	Type   uint8
	Data   [6]uint8
}

// ScriptInstrSceEsprKill represents a SCE_ESPR_KILL instruction (0x4c)
type ScriptInstrSceEsprKill struct {
	Opcode        uint8 // 0x4c
	Id            uint8
	Type          uint8
	WorkComponent uint8
	WorkIndex     uint8
}

// ScriptInstrDoorModelSet represents a DOOR_MODEL_SET instruction (0x4d)
type ScriptInstrDoorModelSet struct {
	Opcode      uint8 // 0x4d
	Index       uint8
	Id          uint8
	Type        uint8
	Flag        uint8
	ModelNumber uint8
	Unknown0    uint16
	Unknown1    uint16
	Position    [3]int16
	Direction   [3]int16
}

// ScriptInstrItemAotSet represents an ITEM_AOT_SET instruction (0x4e)
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

// ScriptInstrSceBgmControl represents a SCE_BGM_CONTROL instruction (0x51)
type ScriptInstrSceBgmControl struct {
	Opcode      uint8 // 0x51
	Id          uint8 // 0: Main, 1: sub0, 2: sub1
	Operation   uint8 // 0: nop, 1: start, 2: stop, 3: restart, 4: pause, 5: fadeout
	Type        uint8 // 0: MAIN_VOL, 1: PROG0_VOL, 2: PROG1_VOL, 3: PROG2_VOL
	LeftVolume  uint8
	RightVolume uint8
}

// ScriptInstrSceEsprControl represents a SCE_ESPR_CONTROL instruction (0x52)
type ScriptInstrSceEsprControl struct {
	Opcode        uint8 // 0x52
	Id            uint8
	Type          uint8
	Action        uint8
	WorkComponent uint8
	WorkIndex     uint8
}

// ScriptInstrSceEspr3DOn represents a SCE_ESPR3D_ON instruction (0x54)
type ScriptInstrSceEspr3DOn struct {
	Opcode   uint8 // 0x54
	Dummy    uint8
	Unknown0 uint16
	Work     uint16
	Unknown1 uint16
	Vector1  [3]int16
	Vector2  [3]int16
	DirY     uint16
}

// ScriptInstrPlcRot represents a PLC_ROT instruction (0x58)
type ScriptInstrPlcRot struct {
	Opcode uint8 // 0x58
	Index  uint8 // 0 or 1
	Value  int16
}

// ScriptInstrXaOn represents an XA_ON instruction (0x59)
type ScriptInstrXaOn struct {
	Opcode  uint8 // 0x59
	Channel uint8 // channel on which to play sound
	Id      int16 // ID of sound to play
}

// ScriptInstrMizuDivSet represents a MIZU_DIV_SET instruction (0x5d)
type ScriptInstrMizuDivSet struct {
	Opcode     uint8 // 0x5d
	MizuDivMax uint8
}

// ScriptInstrKageSet represents a KAGE_SET instruction (0x60)
type ScriptInstrKageSet struct {
	Opcode           uint8 // 0x60
	WorkSetComponent uint8
	WorkSetIndex     uint8
	Color            [3]uint8
	HalfX, HalfZ     int16
	OffsetX, OffsetZ int16
}

// ScriptInstrAotSet4p represents an AOT_SET_4P instruction (0x67)
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

// ScriptInstrDoorAotSet4p represents a DOOR_AOT_SET_4P instruction (0x68)
type ScriptInstrDoorAotSet4p struct {
	Opcode                       uint8 // 0x68
	Aot                          uint8 // Index of item in array of room objects list
	Id                           uint8
	Type                         uint8
	Floor                        uint8
	Super                        uint8
	X1, Z1                       int16
	X2, Z2                       int16
	X3, Z3                       int16
	X4, Z4                       int16
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

// ScriptInstrItemAotSet4p represents an ITEM_AOT_SET_4P instruction (0x69)
type ScriptInstrItemAotSet4p struct {
	Opcode          uint8 // 0x69
	Aot             uint8
	Id              uint8
	Type            uint8
	Floor           uint8
	Super           uint8
	X1, Z1          int16
	X2, Z2          int16
	X3, Z3          int16
	X4, Z4          int16
	ItemId          uint16
	Amount          uint16
	ItemPickedIndex uint16 // flag to check if item is picked up
	Md1ModelId      uint8
	Act             uint8
}

// SCDOutput represents the parsed output from a script data file
type SCDOutput struct {
	ScriptData ScriptFunction
}

// ScriptFunction represents a parsed script function with its instructions
type ScriptFunction struct {
	Instructions        map[int][]byte // key is program counter, value is command
	StartProgramCounter []int          // set per function
}
