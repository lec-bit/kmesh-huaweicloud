// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package dualengine

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type KmeshCgroupSockWorkloadBpfSockTuple struct {
	Ipv4 struct {
		Saddr uint32
		Daddr uint32
		Sport uint16
		Dport uint16
	}
	_ [24]byte
}

type KmeshCgroupSockWorkloadBuf struct{ Data [40]int8 }

type KmeshCgroupSockWorkloadKmeshConfig struct {
	BpfLogLevel      uint32
	NodeIp           [4]uint32
	PodGateway       [4]uint32
	AuthzOffload     uint32
	EnableMonitoring uint32
}

type KmeshCgroupSockWorkloadManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshCgroupSockWorkloadOperationUsageData struct {
	StartTime     uint64
	EndTime       uint64
	PidTgid       uint64
	OperationType uint32
	_             [4]byte
}

type KmeshCgroupSockWorkloadOperationUsageKey struct {
	SocketCookie  uint64
	OperationType uint32
	_             [4]byte
}

type KmeshCgroupSockWorkloadSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshCgroupSockWorkload returns the embedded CollectionSpec for KmeshCgroupSockWorkload.
func LoadKmeshCgroupSockWorkload() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshCgroupSockWorkloadBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshCgroupSockWorkload: %w", err)
	}

	return spec, err
}

// LoadKmeshCgroupSockWorkloadObjects loads KmeshCgroupSockWorkload and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshCgroupSockWorkloadObjects
//	*KmeshCgroupSockWorkloadPrograms
//	*KmeshCgroupSockWorkloadMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshCgroupSockWorkloadObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshCgroupSockWorkload()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshCgroupSockWorkloadSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockWorkloadSpecs struct {
	KmeshCgroupSockWorkloadProgramSpecs
	KmeshCgroupSockWorkloadMapSpecs
}

// KmeshCgroupSockWorkloadSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockWorkloadProgramSpecs struct {
	CgroupConnect4Prog *ebpf.ProgramSpec `ebpf:"cgroup_connect4_prog"`
	CgroupConnect6Prog *ebpf.ProgramSpec `ebpf:"cgroup_connect6_prog"`
}

// KmeshCgroupSockWorkloadMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshCgroupSockWorkloadMapSpecs struct {
	KmAuthReq     *ebpf.MapSpec `ebpf:"km_auth_req"`
	KmAuthRes     *ebpf.MapSpec `ebpf:"km_auth_res"`
	KmBackend     *ebpf.MapSpec `ebpf:"km_backend"`
	KmCgrTailcall *ebpf.MapSpec `ebpf:"km_cgr_tailcall"`
	KmConfigmap   *ebpf.MapSpec `ebpf:"km_configmap"`
	KmEndpoint    *ebpf.MapSpec `ebpf:"km_endpoint"`
	KmFrontend    *ebpf.MapSpec `ebpf:"km_frontend"`
	KmLogEvent    *ebpf.MapSpec `ebpf:"km_log_event"`
	KmManage      *ebpf.MapSpec `ebpf:"km_manage"`
	KmOrigDst     *ebpf.MapSpec `ebpf:"km_orig_dst"`
	KmPerfInfo    *ebpf.MapSpec `ebpf:"km_perf_info"`
	KmPerfMap     *ebpf.MapSpec `ebpf:"km_perf_map"`
	KmService     *ebpf.MapSpec `ebpf:"km_service"`
	KmSockstorage *ebpf.MapSpec `ebpf:"km_sockstorage"`
	KmTcpProbe    *ebpf.MapSpec `ebpf:"km_tcp_probe"`
	KmTmpbuf      *ebpf.MapSpec `ebpf:"km_tmpbuf"`
	KmWlpolicy    *ebpf.MapSpec `ebpf:"km_wlpolicy"`
	KmXdpTailcall *ebpf.MapSpec `ebpf:"km_xdp_tailcall"`
	KmeshMap1600  *ebpf.MapSpec `ebpf:"kmesh_map1600"`
	KmeshMap192   *ebpf.MapSpec `ebpf:"kmesh_map192"`
	KmeshMap296   *ebpf.MapSpec `ebpf:"kmesh_map296"`
	KmeshMap64    *ebpf.MapSpec `ebpf:"kmesh_map64"`
}

// KmeshCgroupSockWorkloadObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockWorkloadObjects struct {
	KmeshCgroupSockWorkloadPrograms
	KmeshCgroupSockWorkloadMaps
}

func (o *KmeshCgroupSockWorkloadObjects) Close() error {
	return _KmeshCgroupSockWorkloadClose(
		&o.KmeshCgroupSockWorkloadPrograms,
		&o.KmeshCgroupSockWorkloadMaps,
	)
}

// KmeshCgroupSockWorkloadMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockWorkloadMaps struct {
	KmAuthReq     *ebpf.Map `ebpf:"km_auth_req"`
	KmAuthRes     *ebpf.Map `ebpf:"km_auth_res"`
	KmBackend     *ebpf.Map `ebpf:"km_backend"`
	KmCgrTailcall *ebpf.Map `ebpf:"km_cgr_tailcall"`
	KmConfigmap   *ebpf.Map `ebpf:"km_configmap"`
	KmEndpoint    *ebpf.Map `ebpf:"km_endpoint"`
	KmFrontend    *ebpf.Map `ebpf:"km_frontend"`
	KmLogEvent    *ebpf.Map `ebpf:"km_log_event"`
	KmManage      *ebpf.Map `ebpf:"km_manage"`
	KmOrigDst     *ebpf.Map `ebpf:"km_orig_dst"`
	KmPerfInfo    *ebpf.Map `ebpf:"km_perf_info"`
	KmPerfMap     *ebpf.Map `ebpf:"km_perf_map"`
	KmService     *ebpf.Map `ebpf:"km_service"`
	KmSockstorage *ebpf.Map `ebpf:"km_sockstorage"`
	KmTcpProbe    *ebpf.Map `ebpf:"km_tcp_probe"`
	KmTmpbuf      *ebpf.Map `ebpf:"km_tmpbuf"`
	KmWlpolicy    *ebpf.Map `ebpf:"km_wlpolicy"`
	KmXdpTailcall *ebpf.Map `ebpf:"km_xdp_tailcall"`
	KmeshMap1600  *ebpf.Map `ebpf:"kmesh_map1600"`
	KmeshMap192   *ebpf.Map `ebpf:"kmesh_map192"`
	KmeshMap296   *ebpf.Map `ebpf:"kmesh_map296"`
	KmeshMap64    *ebpf.Map `ebpf:"kmesh_map64"`
}

func (m *KmeshCgroupSockWorkloadMaps) Close() error {
	return _KmeshCgroupSockWorkloadClose(
		m.KmAuthReq,
		m.KmAuthRes,
		m.KmBackend,
		m.KmCgrTailcall,
		m.KmConfigmap,
		m.KmEndpoint,
		m.KmFrontend,
		m.KmLogEvent,
		m.KmManage,
		m.KmOrigDst,
		m.KmPerfInfo,
		m.KmPerfMap,
		m.KmService,
		m.KmSockstorage,
		m.KmTcpProbe,
		m.KmTmpbuf,
		m.KmWlpolicy,
		m.KmXdpTailcall,
		m.KmeshMap1600,
		m.KmeshMap192,
		m.KmeshMap296,
		m.KmeshMap64,
	)
}

// KmeshCgroupSockWorkloadPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshCgroupSockWorkloadObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshCgroupSockWorkloadPrograms struct {
	CgroupConnect4Prog *ebpf.Program `ebpf:"cgroup_connect4_prog"`
	CgroupConnect6Prog *ebpf.Program `ebpf:"cgroup_connect6_prog"`
}

func (p *KmeshCgroupSockWorkloadPrograms) Close() error {
	return _KmeshCgroupSockWorkloadClose(
		p.CgroupConnect4Prog,
		p.CgroupConnect6Prog,
	)
}

func _KmeshCgroupSockWorkloadClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshcgroupsockworkload_bpfel.o
var _KmeshCgroupSockWorkloadBytes []byte
