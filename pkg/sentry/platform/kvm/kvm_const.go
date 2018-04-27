// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kvm

// KVM ioctls.
//
// Only the ioctls we need in Go appear here; some additional ioctls are used
// within the assembly stubs (KVM_INTERRUPT, etc.).
const (
	_KVM_CREATE_VM              = 0xae01
	_KVM_GET_VCPU_MMAP_SIZE     = 0xae04
	_KVM_CREATE_VCPU            = 0xae41
	_KVM_SET_TSS_ADDR           = 0xae47
	_KVM_RUN                    = 0xae80
	_KVM_INTERRUPT              = 0x4004ae86
	_KVM_SET_MSRS               = 0x4008ae89
	_KVM_SET_USER_MEMORY_REGION = 0x4020ae46
	_KVM_SET_REGS               = 0x4090ae82
	_KVM_SET_SREGS              = 0x4138ae84
	_KVM_GET_SUPPORTED_CPUID    = 0xc008ae05
	_KVM_SET_CPUID2             = 0x4008ae90
	_KVM_SET_SIGNAL_MASK        = 0x4004ae8b
)

// KVM exit reasons.
const (
	_KVM_EXIT_EXCEPTION       = 0x1
	_KVM_EXIT_IO              = 0x2
	_KVM_EXIT_HYPERCALL       = 0x3
	_KVM_EXIT_DEBUG           = 0x4
	_KVM_EXIT_HLT             = 0x5
	_KVM_EXIT_MMIO            = 0x6
	_KVM_EXIT_IRQ_WINDOW_OPEN = 0x7
	_KVM_EXIT_SHUTDOWN        = 0x8
	_KVM_EXIT_FAIL_ENTRY      = 0x9
	_KVM_EXIT_INTERNAL_ERROR  = 0x11
)

// KVM limits.
const (
	_KVM_NR_VCPUS         = 0x100
	_KVM_NR_INTERRUPTS    = 0x100
	_KVM_NR_CPUID_ENTRIES = 0x100
)
