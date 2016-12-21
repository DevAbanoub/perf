// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 21 Dec 2016 18:21:29 GMT.
// By https://git.io/cgogen. DO NOT EDIT.

package gen

/*
#include <linux/perf_event.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// SampleBranchPlmAll as defined in linux/perf_event.h:198
	SampleBranchPlmAll = 0
	// AttrSizeVer0 as defined in linux/perf_event.h:263
	AttrSizeVer0 = 64
	// AttrSizeVer1 as defined in linux/perf_event.h:264
	AttrSizeVer1 = 72
	// AttrSizeVer2 as defined in linux/perf_event.h:265
	AttrSizeVer2 = 80
	// AttrSizeVer3 as defined in linux/perf_event.h:266
	AttrSizeVer3 = 96
	// AttrSizeVer4 as defined in linux/perf_event.h:268
	AttrSizeVer4 = 104
	// AttrSizeVer5 as defined in linux/perf_event.h:269
	AttrSizeVer5 = 112
	// EventIocEnable as defined in linux/perf_event.h:389
	EventIocEnable = 9216
	// EventIocDisable as defined in linux/perf_event.h:390
	EventIocDisable = 9217
	// EventIocRefresh as defined in linux/perf_event.h:391
	EventIocRefresh = 9218
	// EventIocReset as defined in linux/perf_event.h:392
	EventIocReset = 9219
	// EventIocSetOutput as defined in linux/perf_event.h:394
	EventIocSetOutput = 9221
	// RecordMiscCPUModeMask as defined in linux/perf_event.h:564
	RecordMiscCPUModeMask = (7 << 0)
	// RecordMiscCPUModeUnknown as defined in linux/perf_event.h:565
	RecordMiscCPUModeUnknown = (0 << 0)
	// RecordMiscKernel as defined in linux/perf_event.h:566
	RecordMiscKernel = (1 << 0)
	// RecordMiscUser as defined in linux/perf_event.h:567
	RecordMiscUser = (2 << 0)
	// RecordMiscHypervisor as defined in linux/perf_event.h:568
	RecordMiscHypervisor = (3 << 0)
	// RecordMiscGuestKernel as defined in linux/perf_event.h:569
	RecordMiscGuestKernel = (4 << 0)
	// RecordMiscGuestUser as defined in linux/perf_event.h:570
	RecordMiscGuestUser = (5 << 0)
	// RecordMiscProcMapParseTimeout as defined in linux/perf_event.h:575
	RecordMiscProcMapParseTimeout = (1 << 12)
	// RecordMiscMmapData as defined in linux/perf_event.h:581
	RecordMiscMmapData = (1 << 13)
	// RecordMiscCommExec as defined in linux/perf_event.h:582
	RecordMiscCommExec = (1 << 13)
	// RecordMiscSwitchOut as defined in linux/perf_event.h:583
	RecordMiscSwitchOut = (1 << 13)
	// RecordMiscExactIp as defined in linux/perf_event.h:589
	RecordMiscExactIp = (1 << 14)
	// RecordMiscExtReserved as defined in linux/perf_event.h:593
	RecordMiscExtReserved = (1 << 15)
	// MaxStackDepth as defined in linux/perf_event.h:856
	MaxStackDepth = 127
	// AuxFlagTruncated as defined in linux/perf_event.h:873
	AuxFlagTruncated = 0x01
	// AuxFlagOverwrite as defined in linux/perf_event.h:874
	AuxFlagOverwrite = 0x02
	// FlagFdNoGroup as defined in linux/perf_event.h:876
	FlagFdNoGroup = (uint64(1) << 0)
	// FlagFdOutput as defined in linux/perf_event.h:877
	FlagFdOutput = (uint64(1) << 1)
	// FlagPidCgroup as defined in linux/perf_event.h:878
	FlagPidCgroup = (uint64(1) << 2)
	// FlagFdCloexec as defined in linux/perf_event.h:879
	FlagFdCloexec = (uint64(1) << 3)
	// MemOpNa as defined in linux/perf_event.h:894
	MemOpNa = 0x01
	// MemOpLoad as defined in linux/perf_event.h:895
	MemOpLoad = 0x02
	// MemOpStore as defined in linux/perf_event.h:896
	MemOpStore = 0x04
	// MemOpPfetch as defined in linux/perf_event.h:897
	MemOpPfetch = 0x08
	// MemOpExec as defined in linux/perf_event.h:898
	MemOpExec = 0x10
	// MemOpShift as defined in linux/perf_event.h:899
	MemOpShift = 0
	// MemLvlNa as defined in linux/perf_event.h:902
	MemLvlNa = 0x01
	// MemLvlHit as defined in linux/perf_event.h:903
	MemLvlHit = 0x02
	// MemLvlMiss as defined in linux/perf_event.h:904
	MemLvlMiss = 0x04
	// MemLvlL1 as defined in linux/perf_event.h:905
	MemLvlL1 = 0x08
	// MemLvlLfb as defined in linux/perf_event.h:906
	MemLvlLfb = 0x10
	// MemLvlL2 as defined in linux/perf_event.h:907
	MemLvlL2 = 0x20
	// MemLvlL3 as defined in linux/perf_event.h:908
	MemLvlL3 = 0x40
	// MemLvlLocRam as defined in linux/perf_event.h:909
	MemLvlLocRam = 0x80
	// MemLvlRemRam1 as defined in linux/perf_event.h:910
	MemLvlRemRam1 = 0x100
	// MemLvlRemRam2 as defined in linux/perf_event.h:911
	MemLvlRemRam2 = 0x200
	// MemLvlRemCce1 as defined in linux/perf_event.h:912
	MemLvlRemCce1 = 0x400
	// MemLvlRemCce2 as defined in linux/perf_event.h:913
	MemLvlRemCce2 = 0x800
	// MemLvlIo as defined in linux/perf_event.h:914
	MemLvlIo = 0x1000
	// MemLvlUnc as defined in linux/perf_event.h:915
	MemLvlUnc = 0x2000
	// MemLvlShift as defined in linux/perf_event.h:916
	MemLvlShift = 5
	// MemSnoopNa as defined in linux/perf_event.h:919
	MemSnoopNa = 0x01
	// MemSnoopNone as defined in linux/perf_event.h:920
	MemSnoopNone = 0x02
	// MemSnoopHit as defined in linux/perf_event.h:921
	MemSnoopHit = 0x04
	// MemSnoopMiss as defined in linux/perf_event.h:922
	MemSnoopMiss = 0x08
	// MemSnoopHitm as defined in linux/perf_event.h:923
	MemSnoopHitm = 0x10
	// MemSnoopShift as defined in linux/perf_event.h:924
	MemSnoopShift = 19
	// MemLockNa as defined in linux/perf_event.h:927
	MemLockNa = 0x01
	// MemLockLocked as defined in linux/perf_event.h:928
	MemLockLocked = 0x02
	// MemLockShift as defined in linux/perf_event.h:929
	MemLockShift = 24
	// MemTlbNa as defined in linux/perf_event.h:932
	MemTlbNa = 0x01
	// MemTlbHit as defined in linux/perf_event.h:933
	MemTlbHit = 0x02
	// MemTlbMiss as defined in linux/perf_event.h:934
	MemTlbMiss = 0x04
	// MemTlbL1 as defined in linux/perf_event.h:935
	MemTlbL1 = 0x08
	// MemTlbL2 as defined in linux/perf_event.h:936
	MemTlbL2 = 0x10
	// MemTlbWk as defined in linux/perf_event.h:937
	MemTlbWk = 0x20
	// MemTlbOs as defined in linux/perf_event.h:938
	MemTlbOs = 0x40
	// MemTlbShift as defined in linux/perf_event.h:939
	MemTlbShift = 26
)

// SampleFormat as declared in linux/perf_event.h:122
type SampleFormat int32

// SampleFormat enumeration from linux/perf_event.h:122
const (
	SampleIp          = C.PERF_SAMPLE_IP
	SampleTid         = C.PERF_SAMPLE_TID
	SampleTime        = C.PERF_SAMPLE_TIME
	SampleAddr        = C.PERF_SAMPLE_ADDR
	SampleRead        = C.PERF_SAMPLE_READ
	SampleCallchain   = C.PERF_SAMPLE_CALLCHAIN
	SampleID          = C.PERF_SAMPLE_ID
	SampleCpu         = C.PERF_SAMPLE_CPU
	SamplePeriod      = C.PERF_SAMPLE_PERIOD
	SampleStreamID    = C.PERF_SAMPLE_STREAM_ID
	SampleRaw         = C.PERF_SAMPLE_RAW
	SampleBranchStack = C.PERF_SAMPLE_BRANCH_STACK
	SampleRegsUser    = C.PERF_SAMPLE_REGS_USER
	SampleStackUser   = C.PERF_SAMPLE_STACK_USER
	SampleWeight      = C.PERF_SAMPLE_WEIGHT
	SampleDataSrc     = C.PERF_SAMPLE_DATA_SRC
	SampleIDentifier  = C.PERF_SAMPLE_IDENTIFIER
	SampleTransaction = C.PERF_SAMPLE_TRANSACTION
	SampleRegsIntr    = C.PERF_SAMPLE_REGS_INTR
	SampleMax         = C.PERF_SAMPLE_MAX
)

// ReadFormat as declared in linux/perf_event.h:254
type ReadFormat int32

// ReadFormat enumeration from linux/perf_event.h:254
const (
	FormatTotalTimeEnabled = C.PERF_FORMAT_TOTAL_TIME_ENABLED
	FormatTotalTimeRunning = C.PERF_FORMAT_TOTAL_TIME_RUNNING
	FormatID               = C.PERF_FORMAT_ID
	FormatGroup            = C.PERF_FORMAT_GROUP
	FormatMax              = C.PERF_FORMAT_MAX
)

// IocFlags as declared in linux/perf_event.h:399
type IocFlags int32

// IocFlags enumeration from linux/perf_event.h:399
const (
	IocFlagGroup = C.PERF_IOC_FLAG_GROUP
)

// Type as declared in linux/perf_event.h:601
type Type int32

// Type enumeration from linux/perf_event.h:601
const (
	RecordMmap          = C.PERF_RECORD_MMAP
	RecordLost          = C.PERF_RECORD_LOST
	RecordComm          = C.PERF_RECORD_COMM
	RecordExit          = C.PERF_RECORD_EXIT
	RecordThrottle      = C.PERF_RECORD_THROTTLE
	RecordUnthrottle    = C.PERF_RECORD_UNTHROTTLE
	RecordFork          = C.PERF_RECORD_FORK
	RecordRead          = C.PERF_RECORD_READ
	RecordSample        = C.PERF_RECORD_SAMPLE
	RecordMmap2         = C.PERF_RECORD_MMAP2
	RecordAux           = C.PERF_RECORD_AUX
	RecordItraceStart   = C.PERF_RECORD_ITRACE_START
	RecordLostSamples   = C.PERF_RECORD_LOST_SAMPLES
	RecordSwitch        = C.PERF_RECORD_SWITCH
	RecordSwitchCPUWide = C.PERF_RECORD_SWITCH_CPU_WIDE
	RecordMax           = C.PERF_RECORD_MAX
)

// PerfTypeID as declared in linux/perf_event.h:28
type PerfTypeID int32

// PerfTypeID enumeration from linux/perf_event.h:28
const (
	TypeHardware   = C.PERF_TYPE_HARDWARE
	TypeSoftware   = C.PERF_TYPE_SOFTWARE
	TypeTracepoint = C.PERF_TYPE_TRACEPOINT
	TypeHWCache    = C.PERF_TYPE_HW_CACHE
	TypeRaw        = C.PERF_TYPE_RAW
	TypeBreakpoint = C.PERF_TYPE_BREAKPOINT
	TypeMax        = C.PERF_TYPE_MAX
)

// PerfHwID as declared in linux/perf_event.h:44
type PerfHwID int32

// PerfHwID enumeration from linux/perf_event.h:44
const (
	CountHWCPUCycles             = C.PERF_COUNT_HW_CPU_CYCLES
	CountHWInstructions          = C.PERF_COUNT_HW_INSTRUCTIONS
	CountHWCacheReferences       = C.PERF_COUNT_HW_CACHE_REFERENCES
	CountHWCacheMisses           = C.PERF_COUNT_HW_CACHE_MISSES
	CountHWBranchInstructions    = C.PERF_COUNT_HW_BRANCH_INSTRUCTIONS
	CountHWBranchMisses          = C.PERF_COUNT_HW_BRANCH_MISSES
	CountHWBusCycles             = C.PERF_COUNT_HW_BUS_CYCLES
	CountHWStalledCyclesFrontend = C.PERF_COUNT_HW_STALLED_CYCLES_FRONTEND
	CountHWStalledCyclesBackend  = C.PERF_COUNT_HW_STALLED_CYCLES_BACKEND
	CountHWRefCPUCycles          = C.PERF_COUNT_HW_REF_CPU_CYCLES
	CountHWMax                   = C.PERF_COUNT_HW_MAX
)

// PerfHwCacheID as declared in linux/perf_event.h:69
type PerfHwCacheID int32

// PerfHwCacheID enumeration from linux/perf_event.h:69
const (
	CountHWCacheL1d  = C.PERF_COUNT_HW_CACHE_L1D
	CountHWCacheL1i  = C.PERF_COUNT_HW_CACHE_L1I
	CountHWCacheLl   = C.PERF_COUNT_HW_CACHE_LL
	CountHWCacheDtlb = C.PERF_COUNT_HW_CACHE_DTLB
	CountHWCacheItlb = C.PERF_COUNT_HW_CACHE_ITLB
	CountHWCacheBpu  = C.PERF_COUNT_HW_CACHE_BPU
	CountHWCacheNode = C.PERF_COUNT_HW_CACHE_NODE
	CountHWCacheMax  = C.PERF_COUNT_HW_CACHE_MAX
)

// PerfHwCacheOpID as declared in linux/perf_event.h:81
type PerfHwCacheOpID int32

// PerfHwCacheOpID enumeration from linux/perf_event.h:81
const (
	CountHWCacheOpRead     = C.PERF_COUNT_HW_CACHE_OP_READ
	CountHWCacheOpWrite    = C.PERF_COUNT_HW_CACHE_OP_WRITE
	CountHWCacheOpPrefetch = C.PERF_COUNT_HW_CACHE_OP_PREFETCH
	CountHWCacheOpMax      = C.PERF_COUNT_HW_CACHE_OP_MAX
)

// PerfHwCacheOpResultID as declared in linux/perf_event.h:89
type PerfHwCacheOpResultID int32

// PerfHwCacheOpResultID enumeration from linux/perf_event.h:89
const (
	CountHWCacheResultAccess = C.PERF_COUNT_HW_CACHE_RESULT_ACCESS
	CountHWCacheResultMiss   = C.PERF_COUNT_HW_CACHE_RESULT_MISS
	CountHWCacheResultMax    = C.PERF_COUNT_HW_CACHE_RESULT_MAX
)

// PerfSwIds as declared in linux/perf_event.h:102
type PerfSwIds int32

// PerfSwIds enumeration from linux/perf_event.h:102
const (
	CountSWCPUClock        = C.PERF_COUNT_SW_CPU_CLOCK
	CountSWTaskClock       = C.PERF_COUNT_SW_TASK_CLOCK
	CountSWPageFaults      = C.PERF_COUNT_SW_PAGE_FAULTS
	CountSWContextSwitches = C.PERF_COUNT_SW_CONTEXT_SWITCHES
	CountSWCPUMigrations   = C.PERF_COUNT_SW_CPU_MIGRATIONS
	CountSWPageFaultsMin   = C.PERF_COUNT_SW_PAGE_FAULTS_MIN
	CountSWPageFaultsMaj   = C.PERF_COUNT_SW_PAGE_FAULTS_MAJ
	CountSWAlignmentFaults = C.PERF_COUNT_SW_ALIGNMENT_FAULTS
	CountSWEmulationFaults = C.PERF_COUNT_SW_EMULATION_FAULTS
	CountSWDummy           = C.PERF_COUNT_SW_DUMMY
	CountSWBpfOutput       = C.PERF_COUNT_SW_BPF_OUTPUT
	CountSWMax             = C.PERF_COUNT_SW_MAX
)

// PerfBranchSampleTypeShift as declared in linux/perf_event.h:156
type PerfBranchSampleTypeShift int32

// PerfBranchSampleTypeShift enumeration from linux/perf_event.h:156
const (
	SampleBranchUserShift      = C.PERF_SAMPLE_BRANCH_USER_SHIFT
	SampleBranchKernelShift    = C.PERF_SAMPLE_BRANCH_KERNEL_SHIFT
	SampleBranchHVShift        = C.PERF_SAMPLE_BRANCH_HV_SHIFT
	SampleBranchAnyShift       = C.PERF_SAMPLE_BRANCH_ANY_SHIFT
	SampleBranchAnyCallShift   = C.PERF_SAMPLE_BRANCH_ANY_CALL_SHIFT
	SampleBranchAnyReturnShift = C.PERF_SAMPLE_BRANCH_ANY_RETURN_SHIFT
	SampleBranchIndCallShift   = C.PERF_SAMPLE_BRANCH_IND_CALL_SHIFT
	SampleBranchAbortTxShift   = C.PERF_SAMPLE_BRANCH_ABORT_TX_SHIFT
	SampleBranchInTxShift      = C.PERF_SAMPLE_BRANCH_IN_TX_SHIFT
	SampleBranchNoTxShift      = C.PERF_SAMPLE_BRANCH_NO_TX_SHIFT
	SampleBranchCondShift      = C.PERF_SAMPLE_BRANCH_COND_SHIFT
	SampleBranchCallStackShift = C.PERF_SAMPLE_BRANCH_CALL_STACK_SHIFT
	SampleBranchIndJumpShift   = C.PERF_SAMPLE_BRANCH_IND_JUMP_SHIFT
	SampleBranchCallShift      = C.PERF_SAMPLE_BRANCH_CALL_SHIFT
	SampleBranchMaxShift       = C.PERF_SAMPLE_BRANCH_MAX_SHIFT
)

// PerfBranchSampleType as declared in linux/perf_event.h:177
type PerfBranchSampleType int32

// PerfBranchSampleType enumeration from linux/perf_event.h:177
const (
	SampleBranchUser      = C.PERF_SAMPLE_BRANCH_USER
	SampleBranchKernel    = C.PERF_SAMPLE_BRANCH_KERNEL
	SampleBranchHV        = C.PERF_SAMPLE_BRANCH_HV
	SampleBranchAny       = C.PERF_SAMPLE_BRANCH_ANY
	SampleBranchAnyCall   = C.PERF_SAMPLE_BRANCH_ANY_CALL
	SampleBranchAnyReturn = C.PERF_SAMPLE_BRANCH_ANY_RETURN
	SampleBranchIndCall   = C.PERF_SAMPLE_BRANCH_IND_CALL
	SampleBranchAbortTx   = C.PERF_SAMPLE_BRANCH_ABORT_TX
	SampleBranchInTx      = C.PERF_SAMPLE_BRANCH_IN_TX
	SampleBranchNoTx      = C.PERF_SAMPLE_BRANCH_NO_TX
	SampleBranchCond      = C.PERF_SAMPLE_BRANCH_COND
	SampleBranchCallStack = C.PERF_SAMPLE_BRANCH_CALL_STACK
	SampleBranchIndJump   = C.PERF_SAMPLE_BRANCH_IND_JUMP
	SampleBranchCall      = C.PERF_SAMPLE_BRANCH_CALL
	SampleBranchMax       = C.PERF_SAMPLE_BRANCH_MAX
)

// PerfSampleRegsAbi as declared in linux/perf_event.h:206
type PerfSampleRegsAbi int32

// PerfSampleRegsAbi enumeration from linux/perf_event.h:206
const (
	SampleRegsABINone = C.PERF_SAMPLE_REGS_ABI_NONE
	SampleRegsABI32   = C.PERF_SAMPLE_REGS_ABI_32
	SampleRegsABI64   = C.PERF_SAMPLE_REGS_ABI_64
)

const (
	// TxnElision as declared in linux/perf_event.h:217
	TxnElision = C.PERF_TXN_ELISION
	// TxnTransaction as declared in linux/perf_event.h:218
	TxnTransaction = C.PERF_TXN_TRANSACTION
	// TxnSync as declared in linux/perf_event.h:219
	TxnSync = C.PERF_TXN_SYNC
	// TxnAsync as declared in linux/perf_event.h:220
	TxnAsync = C.PERF_TXN_ASYNC
	// TxnRetry as declared in linux/perf_event.h:221
	TxnRetry = C.PERF_TXN_RETRY
	// TxnConflict as declared in linux/perf_event.h:222
	TxnConflict = C.PERF_TXN_CONFLICT
	// TxnCapacityWrite as declared in linux/perf_event.h:223
	TxnCapacityWrite = C.PERF_TXN_CAPACITY_WRITE
	// TxnCapacityRead as declared in linux/perf_event.h:224
	TxnCapacityRead = C.PERF_TXN_CAPACITY_READ
	// TxnMax as declared in linux/perf_event.h:226
	TxnMax = C.PERF_TXN_MAX
	// TxnAbortMask as declared in linux/perf_event.h:230
	TxnAbortMask = C.PERF_TXN_ABORT_MASK
	// TxnAbortShift as declared in linux/perf_event.h:231
	TxnAbortShift = C.PERF_TXN_ABORT_SHIFT
)

// PerfCallchainContext as declared in linux/perf_event.h:858
type PerfCallchainContext int32

// PerfCallchainContext enumeration from linux/perf_event.h:858
const (
	ContextHV          = C.PERF_CONTEXT_HV
	ContextKernel      = C.PERF_CONTEXT_KERNEL
	ContextUser        = C.PERF_CONTEXT_USER
	ContextGuest       = C.PERF_CONTEXT_GUEST
	ContextGuestKernel = C.PERF_CONTEXT_GUEST_KERNEL
	ContextGuestUser   = C.PERF_CONTEXT_GUEST_USER
	ContextMax         = C.PERF_CONTEXT_MAX
)
