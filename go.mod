module github.com/CsCN-CBJ/ebpf

go 1.19

require (
	github.com/cilium/ebpf v0.10.0
	github.com/frankban/quicktest v1.14.5
	github.com/google/go-cmp v0.5.9
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2
	golang.org/x/sys v0.6.0
)

replace github.com/cilium/ebpf v0.10.0 => github.com/CsCN-CBJ/ebpf v0.10.0

require (
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
)
