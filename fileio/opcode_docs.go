package fileio

// Opcode documentation and IntelliSense-like signatures for Resident Evil 2 / Biohazard 2
// This provides function signatures, parameter hints, and documentation for script opcodes

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

// Opcode signature function type
type OpcodeSignature func([]byte) string

// Helper function to read binary data into a struct
func readInstruction[T any](lineBytes []byte) T {
	var instruction T
	byteArr := bytes.NewBuffer(lineBytes)
	binary.Read(byteArr, binary.LittleEndian, &instruction)
	return instruction
}

// Helper function to format array data
func formatArray(data []uint8) string {
	if len(data) == 0 {
		return "[]"
	}
	result := "["
	for i, v := range data {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", v)
	}
	return result + "]"
}

// Helper function to format 3D coordinates
func formatCoords3D(x, y, z int16) string {
	return fmt.Sprintf("[%d, %d, %d]", x, y, z)
}

// Individual opcode signature generators for each opcode
func formatGosubParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrGoSub](lineBytes)
	return fmt.Sprintf("Event=%d", instruction.Event)
}

func formatCheckBitParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrCheckBitTest](lineBytes)
	return fmt.Sprintf("BitArray=%d, BitNumber=%d, Value=%d",
		instruction.BitArray, instruction.BitNumber, instruction.Value)
}

func formatSetBitParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSetBit](lineBytes)
	return fmt.Sprintf("BitArray=%d, BitNumber=%d, Operation=%d",
		instruction.BitArray, instruction.BitNumber, instruction.Operation)
}

func formatCutChgParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrCutChg](lineBytes)
	return fmt.Sprintf("CameraId=%d", instruction.CameraId)
}

func formatAotSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrAotSet](lineBytes)
	return fmt.Sprintf("Aot=%d, Id=%d, Type=%d, Floor=%d, Super=%d, X=%d, Z=%d, Width=%d, Depth=%d, Data=%s",
		instruction.Aot, instruction.Id, instruction.Type, instruction.Floor, instruction.Super,
		instruction.X, instruction.Z, instruction.Width, instruction.Depth, formatArray(instruction.Data[:]))
}

func formatObjModelSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrObjModelSet](lineBytes)
	return fmt.Sprintf("ObjectIndex=%d, ObjectId=%d, Counter=%d, Wait=%d, Num=%d, Floor=%d, Flag0=%d, Type=%d, Flag1=%d, Attribute=%d, Position=%s, Direction=%s, Offset=%s, Dimensions=%s",
		instruction.ObjectIndex, instruction.ObjectId, instruction.Counter, instruction.Wait, instruction.Num,
		instruction.Floor, instruction.Flag0, instruction.Type, instruction.Flag1, instruction.Attribute,
		formatCoords3D(instruction.Position[0], instruction.Position[1], instruction.Position[2]),
		formatCoords3D(instruction.Direction[0], instruction.Direction[1], instruction.Direction[2]),
		formatCoords3D(instruction.Offset[0], instruction.Offset[1], instruction.Offset[2]),
		fmt.Sprintf("[%d, %d, %d]", instruction.Dimensions[0], instruction.Dimensions[1], instruction.Dimensions[2]))
}

func formatPosSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrPosSet](lineBytes)
	return fmt.Sprintf("Dummy=%d, X=%d, Y=%d, Z=%d",
		instruction.Dummy, instruction.X, instruction.Y, instruction.Z)
}

func formatSceEsprOnParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSceEsprOn](lineBytes)
	return fmt.Sprintf("Dummy=%d, Id=%d, Type=%d, Work=%d, Unknown1=%d, X=%d, Y=%d, Z=%d, DirY=%d",
		instruction.Dummy, instruction.Id, instruction.Type, instruction.Work, instruction.Unknown1,
		instruction.X, instruction.Y, instruction.Z, instruction.DirY)
}

func formatDoorAotSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrDoorAotSet](lineBytes)
	return fmt.Sprintf("Aot=%d, Id=%d, Type=%d, Floor=%d, Super=%d, X=%d, Z=%d, Width=%d, Depth=%d, NextX=%d, NextY=%d, NextZ=%d, NextDir=%d, Stage=%d, Room=%d, Camera=%d, NextFloor=%d, TextureType=%d, DoorType=%d, KnockType=%d, KeyId=%d, KeyType=%d, Free=%d",
		instruction.Aot, instruction.Id, instruction.Type, instruction.Floor, instruction.Super,
		instruction.X, instruction.Z, instruction.Width, instruction.Depth,
		instruction.NextX, instruction.NextY, instruction.NextZ, instruction.NextDir,
		instruction.Stage, instruction.Room, instruction.Camera, instruction.NextFloor,
		instruction.TextureType, instruction.DoorType, instruction.KnockType,
		instruction.KeyId, instruction.KeyType, instruction.Free)
}

func formatPlcNeckParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrPlcNeck](lineBytes)
	return fmt.Sprintf("Operation=%d, NeckX=%d, NeckY=%d, NeckZ=%d, Unknown=%s",
		instruction.Operation, instruction.NeckX, instruction.NeckY, instruction.NeckZ,
		formatArray([]uint8{uint8(instruction.Unknown[0]), uint8(instruction.Unknown[1])}))
}

func formatSceEmSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSceEmSet](lineBytes)
	return fmt.Sprintf("Dummy=%d, Aot=%d, Id=%d, Type=%d, Status=%d, Floor=%d, SoundFlag=%d, ModelType=%d, EmSetFlag=%d, X=%d, Y=%d, Z=%d, DirY=%d, Motion=%d, CtrFlag=%d",
		instruction.Dummy, instruction.Aot, instruction.Id, instruction.Type, instruction.Status,
		instruction.Floor, instruction.SoundFlag, instruction.ModelType, instruction.EmSetFlag,
		instruction.X, instruction.Y, instruction.Z, instruction.DirY, instruction.Motion, instruction.CtrFlag)
}

func formatAotResetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrAotReset](lineBytes)
	return fmt.Sprintf("Aot=%d, Id=%d, Type=%d, Data=%s",
		instruction.Aot, instruction.Id, instruction.Type, formatArray(instruction.Data[:]))
}

func formatItemAotSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrItemAotSet](lineBytes)
	return fmt.Sprintf("Aot=%d, Id=%d, Type=%d, Floor=%d, Super=%d, X=%d, Z=%d, Width=%d, Depth=%d, ItemId=%d, Amount=%d, ItemPickedIndex=%d, Md1ModelId=%d, Act=%d",
		instruction.Aot, instruction.Id, instruction.Type, instruction.Floor, instruction.Super,
		instruction.X, instruction.Z, instruction.Width, instruction.Depth,
		instruction.ItemId, instruction.Amount, instruction.ItemPickedIndex, instruction.Md1ModelId, instruction.Act)
}

func formatSceBgmControlParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSceBgmControl](lineBytes)
	return fmt.Sprintf("Id=%d, Operation=%d, Type=%d, LeftVolume=%d, RightVolume=%d",
		instruction.Id, instruction.Operation, instruction.Type, instruction.LeftVolume, instruction.RightVolume)
}

func formatAotSet4pParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrAotSet4p](lineBytes)
	return fmt.Sprintf("Aot=%d, Id=%d, Type=%d, Floor=%d, Super=%d, X1=%d, Z1=%d, X2=%d, Z2=%d, X3=%d, Z3=%d, X4=%d, Z4=%d, Data=%s",
		instruction.Aot, instruction.Id, instruction.Type, instruction.Floor, instruction.Super,
		instruction.X1, instruction.Z1, instruction.X2, instruction.Z2,
		instruction.X3, instruction.Z3, instruction.X4, instruction.Z4, formatArray(instruction.Data[:]))
}

func formatEventExecParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrEventExec](lineBytes)
	return fmt.Sprintf("ThreadNum=%d, ExOpcode=%d, Event=%d",
		instruction.ThreadNum, instruction.ExOpcode, instruction.Event)
}

func formatIfElseStartParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrIfElseStart](lineBytes)
	return fmt.Sprintf("Dummy=%d, BlockLength=%d",
		instruction.Dummy, instruction.BlockLength)
}

func formatElseStartParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrElseStart](lineBytes)
	return fmt.Sprintf("Dummy=%d, BlockLength=%d",
		instruction.Dummy, instruction.BlockLength)
}

func formatSleepParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSleep](lineBytes)
	return fmt.Sprintf("Dummy=%d, Count=%d",
		instruction.Dummy, instruction.Count)
}

func formatForStartParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrForStart](lineBytes)
	return fmt.Sprintf("Dummy=%d, BlockLength=%d, Count=%d",
		instruction.Dummy, instruction.BlockLength, instruction.Count)
}

func formatSwitchParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSwitch](lineBytes)
	return fmt.Sprintf("VarId=%d, BlockLength=%d",
		instruction.VarId, instruction.BlockLength)
}

func formatSwitchCaseParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSwitchCase](lineBytes)
	return fmt.Sprintf("Dummy=%d, BlockLength=%d, Value=%d",
		instruction.Dummy, instruction.BlockLength, instruction.Value)
}

func formatGotoParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrGoto](lineBytes)
	return fmt.Sprintf("IfElseCounter=%d, LoopLevel=%d, Unknown=%d, Offset=%d",
		instruction.IfElseCounter, instruction.LoopLevel, instruction.Unknown, instruction.Offset)
}

func formatCompareParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrCompare](lineBytes)
	return fmt.Sprintf("Dummy=%d, VarId=%d, Operation=%d, Value=%d",
		instruction.Dummy, instruction.VarId, instruction.Operation, instruction.Value)
}

func formatSaveParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSave](lineBytes)
	return fmt.Sprintf("VarId=%d, Value=%d",
		instruction.VarId, instruction.Value)
}

func formatCopyParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrCopy](lineBytes)
	return fmt.Sprintf("DestVarId=%d, SourceVarId=%d",
		instruction.DestVarId, instruction.SourceVarId)
}

func formatCalcParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrCalc](lineBytes)
	return fmt.Sprintf("Dummy=%d, Operation=%d, VarId=%d, Value=%d",
		instruction.Dummy, instruction.Operation, instruction.VarId, instruction.Value)
}

func formatCalc2Params(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrCalc2](lineBytes)
	return fmt.Sprintf("Operation=%d, VarId=%d, SourceVarId=%d",
		instruction.Operation, instruction.VarId, instruction.SourceVarId)
}

func formatWorkSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrWorkSet](lineBytes)
	return fmt.Sprintf("Component=%d, Index=%d",
		instruction.Component, instruction.Index)
}

func formatMemberSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrMemberSet](lineBytes)
	return fmt.Sprintf("MemberIndex=%d, Value=%d",
		instruction.MemberIndex, instruction.Value)
}

func formatScaIdSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrScaIdSet](lineBytes)
	return fmt.Sprintf("Id=%d, Flag=%d",
		instruction.Id, instruction.Flag)
}

func formatCutAutoParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrCutAuto](lineBytes)
	return fmt.Sprintf("FlagOn=%d", instruction.FlagOn)
}

func formatMemberCompareParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrMemberCompare](lineBytes)
	return fmt.Sprintf("Unknown0=%d, MemberIndex=%d, CompareOperation=%d, Value=%d",
		instruction.Unknown0, instruction.MemberIndex, instruction.CompareOperation, instruction.Value)
}

func formatPlcMotionParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrPlcMotion](lineBytes)
	return fmt.Sprintf("Action=%d, MoveNumber=%d, SceneFlag=%d",
		instruction.Action, instruction.MoveNumber, instruction.SceneFlag)
}

func formatPlcDestParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrPlcDest](lineBytes)
	return fmt.Sprintf("Dummy=%d, Action=%d, FlagNumber=%d, DestX=%d, DestZ=%d",
		instruction.Dummy, instruction.Action, instruction.FlagNumber, instruction.DestX, instruction.DestZ)
}

func formatPlcFlagParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrPlcFlag](lineBytes)
	return fmt.Sprintf("Operation=%d, Flag=%d",
		instruction.Operation, instruction.Flag)
}

func formatSceEsprKillParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSceEsprKill](lineBytes)
	return fmt.Sprintf("Id=%d, Type=%d, WorkComponent=%d, WorkIndex=%d",
		instruction.Id, instruction.Type, instruction.WorkComponent, instruction.WorkIndex)
}

func formatDoorModelSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrDoorModelSet](lineBytes)
	return fmt.Sprintf("Index=%d, Id=%d, Type=%d, Flag=%d, ModelNumber=%d, Unknown0=%d, Unknown1=%d, Position=%s, Direction=%s",
		instruction.Index, instruction.Id, instruction.Type, instruction.Flag, instruction.ModelNumber,
		instruction.Unknown0, instruction.Unknown1,
		formatCoords3D(instruction.Position[0], instruction.Position[1], instruction.Position[2]),
		formatCoords3D(instruction.Direction[0], instruction.Direction[1], instruction.Direction[2]))
}

func formatSceEsprControlParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSceEsprControl](lineBytes)
	return fmt.Sprintf("Id=%d, Type=%d, Action=%d, WorkComponent=%d, WorkIndex=%d",
		instruction.Id, instruction.Type, instruction.Action, instruction.WorkComponent, instruction.WorkIndex)
}

func formatSceEspr3DOnParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrSceEspr3DOn](lineBytes)
	return fmt.Sprintf("Dummy=%d, Unknown0=%d, Work=%d, Unknown1=%d, Vector1=%s, Vector2=%s, DirY=%d",
		instruction.Dummy, instruction.Unknown0, instruction.Work, instruction.Unknown1,
		formatCoords3D(instruction.Vector1[0], instruction.Vector1[1], instruction.Vector1[2]),
		formatCoords3D(instruction.Vector2[0], instruction.Vector2[1], instruction.Vector2[2]),
		instruction.DirY)
}

func formatPlcRotParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrPlcRot](lineBytes)
	return fmt.Sprintf("Index=%d, Value=%d",
		instruction.Index, instruction.Value)
}

func formatXaOnParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrXaOn](lineBytes)
	return fmt.Sprintf("Channel=%d, Id=%d",
		instruction.Channel, instruction.Id)
}

func formatMizuDivSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrMizuDivSet](lineBytes)
	return fmt.Sprintf("MizuDivMax=%d", instruction.MizuDivMax)
}

func formatKageSetParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrKageSet](lineBytes)
	return fmt.Sprintf("WorkSetComponent=%d, WorkSetIndex=%d, Color=%s, HalfX=%d, HalfZ=%d, OffsetX=%d, OffsetZ=%d",
		instruction.WorkSetComponent, instruction.WorkSetIndex,
		formatArray(instruction.Color[:]), instruction.HalfX, instruction.HalfZ, instruction.OffsetX, instruction.OffsetZ)
}

func formatDoorAotSet4pParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrDoorAotSet4p](lineBytes)
	return fmt.Sprintf("Aot=%d, Id=%d, Type=%d, Floor=%d, Super=%d, X1=%d, Z1=%d, X2=%d, Z2=%d, X3=%d, Z3=%d, X4=%d, Z4=%d, NextX=%d, NextY=%d, NextZ=%d, NextDir=%d, Stage=%d, Room=%d, Camera=%d, NextFloor=%d, TextureType=%d, DoorType=%d, KnockType=%d, KeyId=%d, KeyType=%d, Free=%d",
		instruction.Aot, instruction.Id, instruction.Type, instruction.Floor, instruction.Super,
		instruction.X1, instruction.Z1, instruction.X2, instruction.Z2,
		instruction.X3, instruction.Z3, instruction.X4, instruction.Z4,
		instruction.NextX, instruction.NextY, instruction.NextZ, instruction.NextDir,
		instruction.Stage, instruction.Room, instruction.Camera, instruction.NextFloor,
		instruction.TextureType, instruction.DoorType, instruction.KnockType,
		instruction.KeyId, instruction.KeyType, instruction.Free)
}

func formatItemAotSet4pParams(lineBytes []byte) string {
	instruction := readInstruction[ScriptInstrItemAotSet4p](lineBytes)
	return fmt.Sprintf("Aot=%d, Id=%d, Type=%d, Floor=%d, Super=%d, X1=%d, Z1=%d, X2=%d, Z2=%d, X3=%d, Z3=%d, X4=%d, Z4=%d, ItemId=%d, Amount=%d, ItemPickedIndex=%d, Md1ModelId=%d, Act=%d",
		instruction.Aot, instruction.Id, instruction.Type, instruction.Floor, instruction.Super,
		instruction.X1, instruction.Z1, instruction.X2, instruction.Z2,
		instruction.X3, instruction.Z3, instruction.X4, instruction.Z4,
		instruction.ItemId, instruction.Amount, instruction.ItemPickedIndex, instruction.Md1ModelId, instruction.Act)
}

// Placeholder signature generators for missing opcodes (based on Rust documentation)
func formatEvtEndParams(lineBytes []byte) string {
	return ""
}

func formatEvtNextParams(lineBytes []byte) string {
	return ""
}

func formatEvtChainParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d", lineBytes[1], lineBytes[2], lineBytes[3])
}

func formatEvtKillParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d", lineBytes[1])
}

func formatEndIfParams(lineBytes []byte) string {
	return ""
}

func formatSleepingParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d", lineBytes[1], lineBytes[2])
}

func formatWsleepParams(lineBytes []byte) string {
	return ""
}

func formatWsleepingParams(lineBytes []byte) string {
	return ""
}

func formatForEndParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d", lineBytes[1])
}

func formatWhileStartParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d", lineBytes[1], binary.LittleEndian.Uint16(lineBytes[2:4]))
}

func formatWhileEndParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d", lineBytes[1])
}

func formatDoStartParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d", lineBytes[1], binary.LittleEndian.Uint16(lineBytes[2:4]))
}

func formatDoEndParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d", lineBytes[1])
}

func formatEndSwitchParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d", lineBytes[1])
}

func formatDefaultCaseParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d", lineBytes[1], binary.LittleEndian.Uint16(lineBytes[2:4]), binary.LittleEndian.Uint16(lineBytes[4:6]))
}

func formatGosubReturnParams(lineBytes []byte) string {
	return ""
}

func formatBreakParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d", lineBytes[1])
}

func formatWorkCopyParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d", lineBytes[1], lineBytes[2], lineBytes[3])
}

func formatSceRndParams(lineBytes []byte) string {
	return ""
}

func formatCutOldParams(lineBytes []byte) string {
	return ""
}

func formatMessageOnParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d, param4=%d, param5=%d", lineBytes[1], lineBytes[2], lineBytes[3], lineBytes[4], lineBytes[5])
}

func formatSpeedSetParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d", lineBytes[1], lineBytes[2], lineBytes[3])
}

func formatAddSpeedParams(lineBytes []byte) string {
	return ""
}

func formatAddAspeedParams(lineBytes []byte) string {
	return ""
}

func formatDirSetParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d, param4=%d, param5=%d, param6=%d, param7=%d", lineBytes[1], lineBytes[2], lineBytes[3], lineBytes[4], lineBytes[5], lineBytes[6], lineBytes[7])
}

func formatMemberSet2Params(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d", lineBytes[1], lineBytes[2])
}

func formatSeOnParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d, param4=%d, param5=%d, param6=%d, param7=%d, param8=%d, param9=%d, param10=%d, param11=%d", lineBytes[1], lineBytes[2], lineBytes[3], lineBytes[4], lineBytes[5], lineBytes[6], lineBytes[7], lineBytes[8], lineBytes[9], lineBytes[10], lineBytes[11])
}

func formatDirCkParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d, param4=%d, param5=%d, param6=%d, param7=%d", lineBytes[1], lineBytes[2], lineBytes[3], lineBytes[4], lineBytes[5], lineBytes[6], lineBytes[7])
}

func formatMemberCopyParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d", lineBytes[1], lineBytes[2])
}

func formatPlcRetParams(lineBytes []byte) string {
	return ""
}

func formatAotOnParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d", lineBytes[1])
}

func formatCutReplaceParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d", lineBytes[1], lineBytes[2])
}

func formatSceBgmtblSetParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d, param4=%d, param5=%d, param6=%d, param7=%d", lineBytes[1], lineBytes[2], lineBytes[3], lineBytes[4], lineBytes[5], lineBytes[6], lineBytes[7])
}

func formatPlcCntParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d", lineBytes[1])
}

func formatXaVolParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d", lineBytes[1])
}

func formatCutBeSetParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d", lineBytes[1], lineBytes[2], lineBytes[3])
}

func formatSceItemLostParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d", lineBytes[1])
}

func formatSceEsprOn2Params(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d, param4=%d, param5=%d, param6=%d, param7=%d, param8=%d, param9=%d, param10=%d, param11=%d, param12=%d, param13=%d, param14=%d, param15=%d", lineBytes[1], lineBytes[2], lineBytes[3], lineBytes[4], lineBytes[5], lineBytes[6], lineBytes[7], lineBytes[8], lineBytes[9], lineBytes[10], lineBytes[11], lineBytes[12], lineBytes[13], lineBytes[14], lineBytes[15])
}

func formatPlcStopParams(lineBytes []byte) string {
	return ""
}

func formatLightPosSetParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d, param4=%d, param5=%d", lineBytes[1], lineBytes[2], lineBytes[3], lineBytes[4], lineBytes[5])
}

func formatLightKidoSetParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d", lineBytes[1], lineBytes[2], lineBytes[3])
}

func formatPartsSetParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d, param4=%d, param5=%d", lineBytes[1], lineBytes[2], lineBytes[3], lineBytes[4], lineBytes[5])
}

func formatScePartsBombParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d, param4=%d, param5=%d, param6=%d, param7=%d, param8=%d, param9=%d, param10=%d, param11=%d, param12=%d, param13=%d, param14=%d, param15=%d", lineBytes[1], lineBytes[2], lineBytes[3], lineBytes[4], lineBytes[5], lineBytes[6], lineBytes[7], lineBytes[8], lineBytes[9], lineBytes[10], lineBytes[11], lineBytes[12], lineBytes[13], lineBytes[14], lineBytes[15])
}

func formatScePartsDownParams(lineBytes []byte) string {
	return fmt.Sprintf("param1=%d, param2=%d, param3=%d, param4=%d, param5=%d, param6=%d, param7=%d, param8=%d, param9=%d, param10=%d, param11=%d, param12=%d, param13=%d, param14=%d, param15=%d", lineBytes[1], lineBytes[2], lineBytes[3], lineBytes[4], lineBytes[5], lineBytes[6], lineBytes[7], lineBytes[8], lineBytes[9], lineBytes[10], lineBytes[11], lineBytes[12], lineBytes[13], lineBytes[14], lineBytes[15])
}

func formatDefaultParams(lineBytes []byte) string {
	var params []string
	for i := 1; i < len(lineBytes); i++ {
		params = append(params, fmt.Sprintf("%d", lineBytes[i]))
	}
	return strings.Join(params, ", ")
}

// Map of opcodes to their signature generators
var OpcodeSignatures = map[byte]OpcodeSignature{
	// Control flow opcodes
	OP_EVT_EXEC:   formatEventExecParams,
	OP_IF_START:   formatIfElseStartParams,
	OP_ELSE_START: formatElseStartParams,
	OP_SLEEP:      formatSleepParams,
	OP_FOR:        formatForStartParams,
	OP_SWITCH:     formatSwitchParams,
	OP_CASE:       formatSwitchCaseParams,
	OP_GOTO:       formatGotoParams,
	OP_GOSUB:      formatGosubParams,

	// Data manipulation opcodes
	OP_CHECK:   formatCheckBitParams,
	OP_SET_BIT: formatSetBitParams,
	OP_COMPARE: formatCompareParams,
	OP_SAVE:    formatSaveParams,
	OP_COPY:    formatCopyParams,
	OP_CALC:    formatCalcParams,
	OP_CALC2:   formatCalc2Params,

	// Scene and camera opcodes
	OP_CUT_CHG:  formatCutChgParams,
	OP_CUT_AUTO: formatCutAutoParams,

	// Area of Trigger (AOT) opcodes
	OP_AOT_SET:    formatAotSetParams,
	OP_AOT_RESET:  formatAotResetParams,
	OP_AOT_SET_4P: formatAotSet4pParams,

	// Object and model opcodes
	OP_OBJ_MODEL_SET:  formatObjModelSetParams,
	OP_DOOR_MODEL_SET: formatDoorModelSetParams,

	// Work and member opcodes
	OP_WORK_SET:   formatWorkSetParams,
	OP_MEMBER_SET: formatMemberSetParams,
	OP_MEMBER_CMP: formatMemberCompareParams,

	// Position and movement opcodes
	OP_POS_SET: formatPosSetParams,

	// Scene ID opcodes
	OP_SCA_ID_SET: formatScaIdSetParams,

	// Effect and sprite opcodes
	OP_SCE_ESPR_ON:      formatSceEsprOnParams,
	OP_SCE_ESPR_KILL:    formatSceEsprKillParams,
	OP_SCE_ESPR_CONTROL: formatSceEsprControlParams,
	OP_SCE_ESPR3D_ON:    formatSceEspr3DOnParams,

	// Door opcodes
	OP_DOOR_AOT_SET:    formatDoorAotSetParams,
	OP_DOOR_AOT_SET_4P: formatDoorAotSet4pParams,

	// Player control opcodes
	OP_PLC_MOTION: formatPlcMotionParams,
	OP_PLC_DEST:   formatPlcDestParams,
	OP_PLC_NECK:   formatPlcNeckParams,
	OP_PLC_FLAG:   formatPlcFlagParams,
	OP_PLC_ROT:    formatPlcRotParams,

	// Entity management opcodes
	OP_SCE_EM_SET: formatSceEmSetParams,

	// Item opcodes
	OP_ITEM_AOT_SET:    formatItemAotSetParams,
	OP_ITEM_AOT_SET_4P: formatItemAotSet4pParams,

	// Audio opcodes
	OP_SCE_BGM_CONTROL: formatSceBgmControlParams,
	OP_XA_ON:           formatXaOnParams,

	// Visual effects opcodes
	OP_KAGE_SET:     formatKageSetParams,
	OP_MIZU_DIV_SET: formatMizuDivSetParams,

	// Additional placeholder formatters for complete coverage
	OP_EVT_END:        formatEvtEndParams,
	OP_EVT_NEXT:       formatEvtNextParams,
	OP_EVT_CHAIN:      formatEvtChainParams,
	OP_EVT_KILL:       formatEvtKillParams,
	OP_END_IF:         formatEndIfParams,
	OP_SLEEPING:       formatSleepingParams,
	OP_WSLEEP:         formatWsleepParams,
	OP_WSLEEPING:      formatWsleepingParams,
	OP_FOR_END:        formatForEndParams,
	OP_WHILE_START:    formatWhileStartParams,
	OP_WHILE_END:      formatWhileEndParams,
	OP_DO_START:       formatDoStartParams,
	OP_DO_END:         formatDoEndParams,
	OP_END_SWITCH:     formatEndSwitchParams,
	OP_DEFAULT:        formatDefaultCaseParams,
	OP_GOSUB_RETURN:   formatGosubReturnParams,
	OP_BREAK:          formatBreakParams,
	OP_WORK_COPY:      formatWorkCopyParams,
	OP_SCE_RND:        formatSceRndParams,
	OP_CUT_OLD:        formatCutOldParams,
	OP_MESSAGE_ON:     formatMessageOnParams,
	OP_SPEED_SET:      formatSpeedSetParams,
	OP_ADD_SPEED:      formatAddSpeedParams,
	OP_ADD_ASPEED:     formatAddAspeedParams,
	OP_DIR_SET:        formatDirSetParams,
	OP_MEMBER_SET2:    formatMemberSet2Params,
	OP_SE_ON:          formatSeOnParams,
	OP_DIR_CK:         formatDirCkParams,
	OP_MEMBER_COPY:    formatMemberCopyParams,
	OP_PLC_RET:        formatPlcRetParams,
	OP_AOT_ON:         formatAotOnParams,
	OP_CUT_REPLACE:    formatCutReplaceParams,
	OP_SCE_BGMTBL_SET: formatSceBgmtblSetParams,
	OP_PLC_CNT:        formatPlcCntParams,
	OP_XA_VOL:         formatXaVolParams,
	OP_CUT_BE_SET:     formatCutBeSetParams,
	OP_KEEP_ITEM_CK:   formatSceItemLostParams,
	OP_SCE_ITEM_LOST:  formatSceItemLostParams,
	OP_SCE_ESPR_ON2:   formatSceEsprOn2Params,
	OP_SCE_ESPR_KILL2: formatSceEsprOn2Params,
	OP_PLC_STOP:       formatPlcStopParams,
	OP_LIGHT_POS_SET:  formatLightPosSetParams,
	OP_LIGHT_KIDO_SET: formatLightKidoSetParams,
	OP_RBJ_RESET:      formatPartsSetParams,
	OP_SCE_SCR_MOVE:   formatPartsSetParams,
	OP_PARTS_SET:      formatPartsSetParams,
	OP_MOVIE_ON:       formatPartsSetParams,
	OP_SCE_PARTS_BOMB: formatScePartsBombParams,
	OP_SCE_PARTS_DOWN: formatScePartsDownParams,
}

// GetOpcodeSignature converts binary instruction data to IntelliSense-like function signature
func GetOpcodeSignature(lineBytes []byte) string {
	opcode := lineBytes[0]

	signature, exists := OpcodeSignatures[opcode]
	if !exists {
		return "(" + formatDefaultParams(lineBytes) + ");"
	}

	return "(" + signature(lineBytes) + ");"
}
