package fileio

// Opcode constants for Resident Evil 2 / Biohazard 2 script engine
// These represent the different script commands available in the game

const (
	// Control flow opcodes
	OP_NO_OP        = 0
	OP_EVT_END      = 1
	OP_EVT_NEXT     = 2
	OP_EVT_CHAIN    = 3
	OP_EVT_EXEC     = 4
	OP_EVT_KILL     = 5
	OP_IF_START     = 6
	OP_ELSE_START   = 7
	OP_END_IF       = 8
	OP_SLEEP        = 9
	OP_SLEEPING     = 10
	OP_WSLEEP       = 11
	OP_WSLEEPING    = 12
	OP_FOR          = 13
	OP_FOR_END      = 14
	OP_WHILE_START  = 15
	OP_WHILE_END    = 16
	OP_DO_START     = 17
	OP_DO_END       = 18
	OP_SWITCH       = 19
	OP_CASE         = 20
	OP_DEFAULT      = 21
	OP_END_SWITCH   = 22
	OP_GOTO         = 23
	OP_GOSUB        = 24
	OP_GOSUB_RETURN = 25
	OP_BREAK        = 26

	// Data manipulation opcodes
	OP_WORK_COPY = 29
	OP_NO_OP2    = 32
	OP_CHECK     = 33
	OP_SET_BIT   = 34
	OP_COMPARE   = 35
	OP_SAVE      = 36
	OP_COPY      = 37
	OP_CALC      = 38
	OP_CALC2     = 39

	// Scene and camera opcodes
	OP_SCE_RND     = 40
	OP_CUT_CHG     = 41
	OP_CUT_OLD     = 42
	OP_MESSAGE_ON  = 43
	OP_CUT_AUTO    = 60
	OP_CUT_REPLACE = 75

	// Area of Trigger (AOT) opcodes
	OP_AOT_SET    = 44
	OP_AOT_RESET  = 70
	OP_AOT_ON     = 71
	OP_AOT_SET_4P = 103

	// Object and model opcodes
	OP_OBJ_MODEL_SET  = 45
	OP_DOOR_MODEL_SET = 77

	// Work and member opcodes
	OP_WORK_SET    = 46
	OP_MEMBER_SET  = 52
	OP_MEMBER_SET2 = 53
	OP_MEMBER_COPY = 61
	OP_MEMBER_CMP  = 62

	// Movement and positioning opcodes
	OP_SPEED_SET  = 47
	OP_ADD_SPEED  = 48
	OP_ADD_ASPEED = 49
	OP_POS_SET    = 50
	OP_DIR_SET    = 51
	OP_DIR_CK     = 57

	// Audio opcodes
	OP_SE_ON      = 54
	OP_SCA_ID_SET = 55
	OP_XA_ON      = 89
	OP_XA_VOL     = 95

	// Effects and sprites opcodes
	OP_SCE_ESPR_ON      = 58
	OP_SCE_ESPR_KILL    = 76
	OP_SCE_ESPR_CONTROL = 82
	OP_SCE_ESPR3D_ON    = 84
	OP_SCE_ESPR_ON2     = 100
	OP_SCE_ESPR_KILL2   = 101

	// Door opcodes
	OP_DOOR_AOT_SET    = 59
	OP_DOOR_AOT_SET_4P = 104

	// Player control opcodes
	OP_PLC_MOTION  = 63
	OP_PLC_DEST    = 64
	OP_PLC_NECK    = 65
	OP_PLC_RET     = 66
	OP_PLC_FLAG    = 67
	OP_PLC_ROT     = 88
	OP_PLC_CNT     = 91
	OP_PLC_GUN_EFF = 99
	OP_PLC_STOP    = 102

	// Entity management opcodes
	OP_SCE_EM_SET = 68
	OP_SUPER_SET  = 72

	// Item opcodes
	OP_ITEM_AOT_SET    = 78
	OP_ITEM_AOT_SET_4P = 105
	OP_KEEP_ITEM_CK    = 94
	OP_SCE_ITEM_LOST   = 98

	// Trigger and collision opcodes
	OP_SCE_TRG_CK = 80

	// Audio control opcodes
	OP_SCE_BGM_CONTROL = 81
	OP_SCE_BGMTBL_SET  = 87

	// Visual effects opcodes
	OP_SCE_FADE_SET = 83
	OP_SCE_SHAKE_ON = 92
	OP_KAGE_SET     = 96
	OP_CUT_BE_SET   = 97

	// Weapon opcodes
	OP_WEAPON_CHG = 90

	// Environment opcodes
	OP_MIZU_DIV_SET = 93

	// Light opcodes
	OP_LIGHT_POS_SET  = 106
	OP_LIGHT_KIDO_SET = 107

	// System opcodes
	OP_RBJ_RESET    = 108
	OP_SCE_SCR_MOVE = 109
	OP_PARTS_SET    = 110
	OP_MOVIE_ON     = 111

	// Parts and destruction opcodes
	OP_SCE_PARTS_BOMB = 122
	OP_SCE_PARTS_DOWN = 123
)

// InstructionSize maps opcodes to their byte sizes
var InstructionSize = map[byte]int{
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

// FunctionName maps opcodes to their human-readable names
var FunctionName = map[byte]string{
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
