package instructions

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/instructions/comparisons"
	"create_jvm/main02/instructions/constants"
	"create_jvm/main02/instructions/control"
	"create_jvm/main02/instructions/conversions"
	"create_jvm/main02/instructions/extended"
	"create_jvm/main02/instructions/loads"
	"create_jvm/main02/instructions/math"
	"create_jvm/main02/instructions/references"
	"create_jvm/main02/instructions/stack"
	"create_jvm/main02/instructions/stores"
	"fmt"
)

var (
	Nop         = &base.Nop{}
	ACONST_NULL = &constants.ACONST_NULL{}
	DCONST_0    = &constants.DCONST_0{}
	DCONST_1    = &constants.DCONST_1{}

	FCONST_0 = &constants.FCONST_0{}
	FCONST_1 = &constants.FCONST_1{}
	FCONST_2 = &constants.FCONST_2{}

	ICONST_M1 = &constants.ICONST_M1{}
	ICONST_0  = &constants.ICONST_0{}
	ICONST_1  = &constants.ICONST_1{}
	ICONST_2  = &constants.ICONST_2{}
	ICONST_3  = &constants.ICONST_3{}
	ICONST_4  = &constants.ICONST_4{}
	ICONST_5  = &constants.ICONST_5{}

	LCONST_0 = &constants.LCONST_0{}
	LCONST_1 = &constants.LCONST_1{}

	BIPUSH = &constants.BIPUSH{}
	SIPUSH = &constants.SIPUSH{}

	GOTO          = &control.GOTO{}
	LOOKUP_SWITCH = &control.LOOKUP_SWITCH{}
	TABLE_SWITCH  = &control.TABLE_SWITCH{}

	D2F = &conversions.D2F{}
	D2L = &conversions.D2L{}
	D2I = &conversions.D2I{}

	GOTO_W    = &extended.GOTO_W{}
	IFNULL    = &extended.IFNULL{}
	IFNONNULL = &extended.IFNONNULL{}
	WIDE      = &extended.WIDE{}

	ILOAD   = &loads.ILOAD{}
	ILOAD_0 = &loads.ILOAD_0{}
	ILOAD_1 = &loads.ILOAD_1{}
	ILOAD_2 = &loads.ILOAD_2{}
	ILOAD_3 = &loads.ILOAD_3{}

	ALOAD   = &loads.A_LOAD{}
	ALOAD_0 = &loads.A_LOAD_0{}
	ALOAD_1 = &loads.A_LOAD_1{}
	ALOAD_2 = &loads.A_LOAD_2{}
	ALOAD_3 = &loads.A_LOAD_3{}

	IAND  = &math.IAND{}
	LAND  = &math.LAND{}
	IINNC = &math.IINNC{}

	DREM = &math.DREM{}
	FREM = &math.FREM{}
	LREM = &math.LREM{}
	IREM = &math.IREM{}

	ISHL  = &math.ISHL{}
	ISHR  = &math.ISHR{}
	IUSHR = &math.IUSHR{}
	LSHL  = &math.LSHL{}
	LSHR  = &math.LSHR{}
	LUSHR = &math.LUSHR{}

	DUP = &stack.DUP{}

	POP  = &stack.POP{}
	POP2 = &stack.POP2{}
	SWAP = &stack.SWAP{}

	LSTORE   = &stores.LSTORE{}
	LSTORE_0 = &stores.LSTORE_0{}
	LSTORE_1 = &stores.LSTORE_1{}
	LSTORE_2 = &stores.LSTORE_2{}
	LSTORE_3 = &stores.LSTORE_3{}

	FCMPG = &comparisons.FCMPG{}
	FCMPL = &comparisons.FCMPL{}

	IF_ACMPEQ = &comparisons.IF_ACMPEQ{}
	IF_ACMPNE = &comparisons.IF_ACMPNE{}

	IF_ICMPEQ = &comparisons.IF_ICMPEQ{}
	IF_ICMPNE = &comparisons.IF_ICMPNE{}
	IF_ICMPLT = &comparisons.IF_ICMPLT{}
	IF_ICMPLE = &comparisons.IF_ICMPLE{}
	IF_ICMPGT = &comparisons.IF_ICMPGT{}
	IF_ICMPGE = &comparisons.IF_ICMPGE{}

	IFEQ = &comparisons.IFEQ{}
	IFNE = &comparisons.IFNE{}
	IFLT = &comparisons.IFLT{}
	IFLE = &comparisons.IFLE{}
	IFGT = &comparisons.IFGT{}
	IFGE = &comparisons.IFGE{}

	LCMP = &comparisons.LCMP{}

	ISTORE   = &stores.ISTORE{}
	ISTORE_0 = &stores.ISTORE_0{}
	ISTORE_1 = &stores.ISTORE_1{}
	ISTORE_2 = &stores.ISTORE_2{}
	ISTORE_3 = &stores.ISTORE_3{}

	IADD = &math.IADD{}

	CHECK_CAST     = &references.CHECK_CAST{}
	GET_FIELD      = &references.GET_FIELD{}
	GET_STATIC     = &references.GET_STATIC{}
	INSTANCE_OF    = &references.INSTANCEOF{}
	INVOKE_VIRTUAL = &references.INVOKE_VIRTUAL{}
	INVOKE_SPECIAL = &references.INVOKE_SPECIAL{}

	NEW        = &references.NEW{}
	PUT_STATIC = &references.PUT_STATIC{}
	PUT_FIELD  = &references.PUT_FIELD{}

	LDC   = &constants.LDC{}
	LDC_W = &constants.LDC_W{}

	ASTORE_0 = &stores.A_STORE_0{}
	ASTORE_1 = &stores.A_STORE_1{}
	ASTORE_2 = &stores.A_STORE_2{}
	ASTORE_3 = &stores.A_STORE_3{}
)

func NewInstructions(opcode byte) base.Instruction {

	switch opcode {
	case 0x00:
		return Nop
	case 0x01:
		return ACONST_NULL

	case 0x0e:
		return DCONST_0

	case 0x0f:
		return DCONST_1

	case 0x0b:
		return FCONST_0

	case 0x0c:
		return FCONST_1

	case 0x0d:
		return FCONST_2

	case 0x02:
		return ICONST_M1

	case 0x03:
		return ICONST_0

	case 0x04:
		return ICONST_1

	case 0x05:
		return ICONST_2

	case 0x06:
		return ICONST_3

	case 0x07:
		return ICONST_4

	case 0x08:
		return ICONST_5

	case 0x09:
		return LCONST_0

	case 0x0a:
		return LCONST_1

	case 0x10:
		return BIPUSH

	case 0x11:
		return SIPUSH

	case 0x12:
		return LDC

	case 0xa7:
		return GOTO

	case 0xab:
		return LOOKUP_SWITCH
	case 0xaa:
		return TABLE_SWITCH

	case 0x90:
		return D2F

	case 0x8f:
		return D2L
	case 0x8e:
		return D2I
	case 0xa8:
		return GOTO_W
	case 0xc6:
		return IFNULL
	case 0xc7:
		return IFNONNULL
	case 0x15:
		return ILOAD

	case 0x1a:
		return ILOAD_0
	case 0x1b:
		return ILOAD_1
	case 0x1c:
		return ILOAD_2
	case 0x1d:
		return ILOAD_3

	case 0x2a:
		return ALOAD_0

	case 0x2b:
		return ALOAD_1

	case 0x2c:
		return ALOAD_2

	case 0x2d:
		return ALOAD_3

	case 0x7e:
		return IAND

	case 0x7f:
		return LAND
	case 0x84:
		return IINNC
	case 0x73:
		return DREM
	case 0x71:
		return LREM
	case 0x70:
		return IREM
	case 0x78:
		return ISHL

	case 0x7a:
		return ISHR
	case 0x7c:
		return IUSHR
	case 0x79:
		return LSHL
	case 0x59:
		return DUP
	case 0x57:
		return POP
	case 0x58:
		return POP2
	case 0x5f:
		return SWAP

	case 0x36:
		return ISTORE
	case 0x3b:
		return ISTORE_0
	case 0x3c:
		return ISTORE_1
	case 0x3d:
		return ISTORE_2
	case 0x3e:
		return ISTORE_3

	case 0x37:
		return LSTORE
	case 0x3f:
		return LSTORE_0
	case 0x40:
		return LSTORE_1
	case 0x41:
		return LSTORE_2
	case 0x42:
		return LSTORE_3

	case 0x4b:
		return ASTORE_0
	case 0x4c:
		return ASTORE_1
	case 0x4d:
		return ASTORE_2
	case 0x4e:
		return ASTORE_3

	case 0x96:
		return FCMPG
	case 0x95:
		return FCMPL

	case 0xa5:
		return IF_ACMPEQ
	case 0xa6:
		return IF_ACMPNE
	case 0xa0:
		return IF_ICMPNE
	case 0xa1:
		return IF_ICMPLT
	case 0xa4:
		return IF_ICMPLE
	case 0xa3:
		return IF_ICMPGT
	case 0xa2:
		return IF_ICMPGE

	case 0x99:
		return IFEQ
	case 0x9a:
		return IFNE
	case 0x9b:
		return IFLT
	case 0x9e:
		return IFLE
	case 0x9d:
		return IFGT
	case 0x9c:
		return IFGE
	case 0x94:
		return LCMP

	case 0x60:
		return IADD

	case 0xc0:
		return CHECK_CAST
	case 0xb4:
		return GET_FIELD
	case 0xb2:
		return GET_STATIC
	case 0xc1:
		return INSTANCE_OF
	case 0xb6:
		return INVOKE_VIRTUAL

	case 0xb7:
		return INVOKE_SPECIAL

	case 0xbb:
		return NEW
	case 0xb3:
		return PUT_STATIC
	case 0xb5:
		return PUT_FIELD

	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))

	}

	return nil
}
