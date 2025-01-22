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

type KmeshXDPAuthBpfSockTuple struct {
	Ipv4 struct {
		Saddr uint32
		Daddr uint32
		Sport uint16
		Dport uint16
	}
	_ [24]byte
}

type KmeshXDPAuthBuf struct{ Data [40]int8 }

type KmeshXDPAuthKmeshConfig struct {
	BpfLogLevel      uint32
	NodeIp           [4]uint32
	PodGateway       [4]uint32
	AuthzOffload     uint32
	EnableMonitoring uint32
}

type KmeshXDPAuthManagerKey struct {
	NetnsCookie uint64
	_           [8]byte
}

type KmeshXDPAuthSockStorageData struct {
	ConnectNs      uint64
	Direction      uint8
	ConnectSuccess uint8
	_              [6]byte
}

// LoadKmeshXDPAuth returns the embedded CollectionSpec for KmeshXDPAuth.
func LoadKmeshXDPAuth() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_KmeshXDPAuthBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load KmeshXDPAuth: %w", err)
	}

	return spec, err
}

// LoadKmeshXDPAuthObjects loads KmeshXDPAuth and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*KmeshXDPAuthObjects
//	*KmeshXDPAuthPrograms
//	*KmeshXDPAuthMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadKmeshXDPAuthObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadKmeshXDPAuth()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// KmeshXDPAuthSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshXDPAuthSpecs struct {
	KmeshXDPAuthProgramSpecs
	KmeshXDPAuthMapSpecs
}

// KmeshXDPAuthSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshXDPAuthProgramSpecs struct {
	PoliciesCheck          *ebpf.ProgramSpec `ebpf:"policies_check"`
	PolicyCheck            *ebpf.ProgramSpec `ebpf:"policy_check"`
	XdpAuthz               *ebpf.ProgramSpec `ebpf:"xdp_authz"`
	XdpShutdownInUserspace *ebpf.ProgramSpec `ebpf:"xdp_shutdown_in_userspace"`
}

// KmeshXDPAuthMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type KmeshXDPAuthMapSpecs struct {
	KmAuthReq     *ebpf.MapSpec `ebpf:"km_auth_req"`
	KmAuthRes     *ebpf.MapSpec `ebpf:"km_auth_res"`
	KmAuthzPolicy *ebpf.MapSpec `ebpf:"km_authz_policy"`
	KmBackend     *ebpf.MapSpec `ebpf:"km_backend"`
	KmCgrTailcall *ebpf.MapSpec `ebpf:"km_cgr_tailcall"`
	KmConfigmap   *ebpf.MapSpec `ebpf:"km_configmap"`
	KmEndpoint    *ebpf.MapSpec `ebpf:"km_endpoint"`
	KmFrontend    *ebpf.MapSpec `ebpf:"km_frontend"`
	KmLogEvent    *ebpf.MapSpec `ebpf:"km_log_event"`
	KmManage      *ebpf.MapSpec `ebpf:"km_manage"`
	KmService     *ebpf.MapSpec `ebpf:"km_service"`
	KmSockstorage *ebpf.MapSpec `ebpf:"km_sockstorage"`
	KmTcargs      *ebpf.MapSpec `ebpf:"km_tcargs"`
	KmTmpbuf      *ebpf.MapSpec `ebpf:"km_tmpbuf"`
	KmWlpolicy    *ebpf.MapSpec `ebpf:"km_wlpolicy"`
	KmXdpTailcall *ebpf.MapSpec `ebpf:"km_xdp_tailcall"`
	KmeshMap1600  *ebpf.MapSpec `ebpf:"kmesh_map1600"`
	KmeshMap192   *ebpf.MapSpec `ebpf:"kmesh_map192"`
	KmeshMap296   *ebpf.MapSpec `ebpf:"kmesh_map296"`
	KmeshMap64    *ebpf.MapSpec `ebpf:"kmesh_map64"`
}

// KmeshXDPAuthObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshXDPAuthObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshXDPAuthObjects struct {
	KmeshXDPAuthPrograms
	KmeshXDPAuthMaps
}

func (o *KmeshXDPAuthObjects) Close() error {
	return _KmeshXDPAuthClose(
		&o.KmeshXDPAuthPrograms,
		&o.KmeshXDPAuthMaps,
	)
}

// KmeshXDPAuthMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshXDPAuthObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshXDPAuthMaps struct {
	KmAuthReq     *ebpf.Map `ebpf:"km_auth_req"`
	KmAuthRes     *ebpf.Map `ebpf:"km_auth_res"`
	KmAuthzPolicy *ebpf.Map `ebpf:"km_authz_policy"`
	KmBackend     *ebpf.Map `ebpf:"km_backend"`
	KmCgrTailcall *ebpf.Map `ebpf:"km_cgr_tailcall"`
	KmConfigmap   *ebpf.Map `ebpf:"km_configmap"`
	KmEndpoint    *ebpf.Map `ebpf:"km_endpoint"`
	KmFrontend    *ebpf.Map `ebpf:"km_frontend"`
	KmLogEvent    *ebpf.Map `ebpf:"km_log_event"`
	KmManage      *ebpf.Map `ebpf:"km_manage"`
	KmService     *ebpf.Map `ebpf:"km_service"`
	KmSockstorage *ebpf.Map `ebpf:"km_sockstorage"`
	KmTcargs      *ebpf.Map `ebpf:"km_tcargs"`
	KmTmpbuf      *ebpf.Map `ebpf:"km_tmpbuf"`
	KmWlpolicy    *ebpf.Map `ebpf:"km_wlpolicy"`
	KmXdpTailcall *ebpf.Map `ebpf:"km_xdp_tailcall"`
	KmeshMap1600  *ebpf.Map `ebpf:"kmesh_map1600"`
	KmeshMap192   *ebpf.Map `ebpf:"kmesh_map192"`
	KmeshMap296   *ebpf.Map `ebpf:"kmesh_map296"`
	KmeshMap64    *ebpf.Map `ebpf:"kmesh_map64"`
}

func (m *KmeshXDPAuthMaps) Close() error {
	return _KmeshXDPAuthClose(
		m.KmAuthReq,
		m.KmAuthRes,
		m.KmAuthzPolicy,
		m.KmBackend,
		m.KmCgrTailcall,
		m.KmConfigmap,
		m.KmEndpoint,
		m.KmFrontend,
		m.KmLogEvent,
		m.KmManage,
		m.KmService,
		m.KmSockstorage,
		m.KmTcargs,
		m.KmTmpbuf,
		m.KmWlpolicy,
		m.KmXdpTailcall,
		m.KmeshMap1600,
		m.KmeshMap192,
		m.KmeshMap296,
		m.KmeshMap64,
	)
}

// KmeshXDPAuthPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadKmeshXDPAuthObjects or ebpf.CollectionSpec.LoadAndAssign.
type KmeshXDPAuthPrograms struct {
	PoliciesCheck          *ebpf.Program `ebpf:"policies_check"`
	PolicyCheck            *ebpf.Program `ebpf:"policy_check"`
	XdpAuthz               *ebpf.Program `ebpf:"xdp_authz"`
	XdpShutdownInUserspace *ebpf.Program `ebpf:"xdp_shutdown_in_userspace"`
}

func (p *KmeshXDPAuthPrograms) Close() error {
	return _KmeshXDPAuthClose(
		p.PoliciesCheck,
		p.PolicyCheck,
		p.XdpAuthz,
		p.XdpShutdownInUserspace,
	)
}

func _KmeshXDPAuthClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed kmeshxdpauth_bpfel.o
var _KmeshXDPAuthBytes []byte
