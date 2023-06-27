// Code generated by "stringer -output types_string.go -type=MapType,ProgramType,PinType"; DO NOT EDIT.

package ebpf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnspecifiedMap-0]
	_ = x[Hash-1]
	_ = x[Array-2]
	_ = x[ProgramArray-3]
	_ = x[PerfEventArray-4]
	_ = x[PerCPUHash-5]
	_ = x[PerCPUArray-6]
	_ = x[StackTrace-7]
	_ = x[CGroupArray-8]
	_ = x[LRUHash-9]
	_ = x[LRUCPUHash-10]
	_ = x[LPMTrie-11]
	_ = x[ArrayOfMaps-12]
	_ = x[HashOfMaps-13]
	_ = x[DevMap-14]
	_ = x[SockMap-15]
	_ = x[CPUMap-16]
	_ = x[XSKMap-17]
	_ = x[SockHash-18]
	_ = x[CGroupStorage-19]
	_ = x[ReusePortSockArray-20]
	_ = x[PerCPUCGroupStorage-21]
	_ = x[Queue-22]
	_ = x[Stack-23]
	_ = x[SkStorage-24]
	_ = x[DevMapHash-25]
	_ = x[StructOpsMap-26]
	_ = x[RingBuf-27]
	_ = x[InodeStorage-28]
	_ = x[TaskStorage-29]
}

const _MapType_name = "UnspecifiedMapHashArrayProgramArrayPerfEventArrayPerCPUHashPerCPUArrayStackTraceCGroupArrayLRUHashLRUCPUHashLPMTrieArrayOfMapsHashOfMapsDevMapSockMapCPUMapXSKMapSockHashCGroupStorageReusePortSockArrayPerCPUCGroupStorageQueueStackSkStorageDevMapHashStructOpsMapRingBufInodeStorageTaskStorage"

var _MapType_index = [...]uint16{0, 14, 18, 23, 35, 49, 59, 70, 80, 91, 98, 108, 115, 126, 136, 142, 149, 155, 161, 169, 182, 200, 219, 224, 229, 238, 248, 260, 267, 279, 290}

func (i MapType) String() string {
	if i >= MapType(len(_MapType_index)-1) {
		return "MapType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MapType_name[_MapType_index[i]:_MapType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnspecifiedProgram-0]
	_ = x[SocketFilter-1]
	_ = x[Kprobe-2]
	_ = x[SchedCLS-3]
	_ = x[SchedACT-4]
	_ = x[TracePoint-5]
	_ = x[XDP-6]
	_ = x[PerfEvent-7]
	_ = x[CGroupSKB-8]
	_ = x[CGroupSock-9]
	_ = x[LWTIn-10]
	_ = x[LWTOut-11]
	_ = x[LWTXmit-12]
	_ = x[SockOps-13]
	_ = x[SkSKB-14]
	_ = x[CGroupDevice-15]
	_ = x[SkMsg-16]
	_ = x[RawTracepoint-17]
	_ = x[CGroupSockAddr-18]
	_ = x[LWTSeg6Local-19]
	_ = x[LircMode2-20]
	_ = x[SkReuseport-21]
	_ = x[FlowDissector-22]
	_ = x[CGroupSysctl-23]
	_ = x[RawTracepointWritable-24]
	_ = x[CGroupSockopt-25]
	_ = x[Tracing-26]
	_ = x[StructOps-27]
	_ = x[Extension-28]
	_ = x[LSM-29]
	_ = x[SkLookup-30]
	_ = x[Boesfs-31]
	_ = x[Syscall-32]
}

const _ProgramType_name = "UnspecifiedProgramSocketFilterKprobeSchedCLSSchedACTTracePointXDPPerfEventCGroupSKBCGroupSockLWTInLWTOutLWTXmitSockOpsSkSKBCGroupDeviceSkMsgRawTracepointCGroupSockAddrLWTSeg6LocalLircMode2SkReuseportFlowDissectorCGroupSysctlRawTracepointWritableCGroupSockoptTracingStructOpsExtensionLSMSkLookupBoesfsSyscall"

var _ProgramType_index = [...]uint16{0, 18, 30, 36, 44, 52, 62, 65, 74, 83, 93, 98, 104, 111, 118, 123, 135, 140, 153, 167, 179, 188, 199, 212, 224, 245, 258, 265, 274, 283, 286, 294, 300, 307}

func (i ProgramType) String() string {
	if i >= ProgramType(len(_ProgramType_index)-1) {
		return "ProgramType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ProgramType_name[_ProgramType_index[i]:_ProgramType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PinNone-0]
	_ = x[PinByName-1]
}

const _PinType_name = "PinNonePinByName"

var _PinType_index = [...]uint8{0, 7, 16}

func (i PinType) String() string {
	if i < 0 || i >= PinType(len(_PinType_index)-1) {
		return "PinType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PinType_name[_PinType_index[i]:_PinType_index[i+1]]
}
