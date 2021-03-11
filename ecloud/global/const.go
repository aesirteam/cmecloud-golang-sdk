package global

type BootVolumeType int32
type DataVolumeType int32
type VmType int32
type OsType int32
type BillingType int32
type ImageType int32

const (
	//Linux操作系统
	OS_TYPE_LINUX OsType = 0
	//Windows操作系统
	OS_TYPE_WINDOWS OsType = 1

	//性能优化型系统盘
	BOOT_VOLUME_PERFORMANCEOPTIMIZATION BootVolumeType = 0
	//高性能型系统盘
	BOOT_VOLUME_HIGHPERFORMANCE BootVolumeType = 1

	//容量型云硬盘
	DATA_VOLUME_CAPEBS DataVolumeType = 0
	//性能优化型云硬盘
	DATA_VOLUME_SSDEBS DataVolumeType = 1
	//高性能型云硬盘
	DATA_VOLUME_SSD DataVolumeType = 2

	//通用型
	VM_TYPE_COMMON VmType = 0
	//通用入门型
	VM_TYPE_COMMONINTRODUCTORY VmType = 1
	//通用网络优化型
	VM_TYPE_COMMONNETIMPROVE VmType = 2
	//内存优化型
	VM_TYPE_MEMIMPROVE VmType = 3
	//计算型
	VM_TYPE_COMPUTE VmType = 4
	//计算网络优化型
	VM_TYPE_COMPUTENETIMPROVE VmType = 5
	//内存网络优化型
	VM_TYPE_MEMNETIMPROVE VmType = 6
	//本地存储型
	VM_TYPE_LOCALSTORAGE VmType = 7
	//超大内存型
	VM_TYPE_XLARGEMEMORY VmType = 8
	//超高主频型
	VM_TYPE_HIGHFREQUENCY VmType = 9
	//FPGA型
	VM_TYPE_FPGA VmType = 10
	//GPU型
	VM_TYPE_GPU VmType = 11
	//VGPU型
	VM_TYPE_VGPU VmType = 12
	//高IO型
	VM_TYPE_HIGHIO VmType = 13
	//独享型
	VM_TYPE_EXCLUSIVE VmType = 14

	//按时计费
	BILLING_TYPE_HOUR BillingType = 0
	//按月计费
	BILLING_TYPE_MONTH BillingType = 1
	//按年计费
	BILLING_TYPE_YEAR BillingType = 2

	IMAGE_BCLINUX_81_X64 ImageType = 1
	IMAGE_BCLINUX_77_X64 ImageType = 2
	IMAGE_BCLINUX_76_X64 ImageType = 3
	IMAGE_BCLINUX_75_X64 ImageType = 4
	IMAGE_BCLINUX_74_X64 ImageType = 5
	IMAGE_BCLINUX_71_X64 ImageType = 6
	IMAGE_BCLINUX_65_X64 ImageType = 7
)

func (b BootVolumeType) String() string {
	switch b {
	case BOOT_VOLUME_PERFORMANCEOPTIMIZATION:
		return "performanceOptimization"
	case BOOT_VOLUME_HIGHPERFORMANCE:
		return "highPerformance"
	default:
		return "performanceOptimization"
	}
}

func (d DataVolumeType) String() string {
	switch d {
	case DATA_VOLUME_CAPEBS:
		return "capebs"
	case DATA_VOLUME_SSDEBS:
		return "ssdebs"
	case DATA_VOLUME_SSD:
		return "ssd"
	default:
		return "capebs"
	}
}

func (v VmType) String() string {
	switch v {
	case VM_TYPE_COMMON:
		return "common"
	case VM_TYPE_COMMONINTRODUCTORY:
		return "commonIntroductory"
	case VM_TYPE_COMMONNETIMPROVE:
		return "commonNetImprove"
	case VM_TYPE_MEMIMPROVE:
		return "memImprove"
	case VM_TYPE_COMPUTE:
		return "compute"
	case VM_TYPE_COMPUTENETIMPROVE:
		return "computeNetImprove"
	case VM_TYPE_MEMNETIMPROVE:
		return "memNetImprove"
	case VM_TYPE_LOCALSTORAGE:
		return "localStorage"
	case VM_TYPE_XLARGEMEMORY:
		return "xlargeMemory"
	case VM_TYPE_HIGHFREQUENCY:
		return "highFrequency"
	case VM_TYPE_FPGA:
		return "fpga"
	case VM_TYPE_GPU:
		return "gpu"
	case VM_TYPE_VGPU:
		return "vgpu"
	case VM_TYPE_HIGHIO:
		return "highIO"
	case VM_TYPE_EXCLUSIVE:
		return "exclusive"
	default:
		return "common"
	}
}

func (o OsType) String() string {
	switch o {
	case OS_TYPE_LINUX:
		return "linux"
	case OS_TYPE_WINDOWS:
		return "windows"
	default:
		return "linux"
	}
}

func (b BillingType) String() string {
	switch b {
	case BILLING_TYPE_MONTH:
		return "MONTH"
	case BILLING_TYPE_HOUR:
		return "HOUR"
	default:
		return "HOUR"
	}
}

func (t ImageType) String() string {
	switch t {
	case IMAGE_BCLINUX_81_X64:
		return "BC-Linux 8.1 64位"
	case IMAGE_BCLINUX_77_X64:
		return "BC-Linux 7.7 64位"
	case IMAGE_BCLINUX_76_X64:
		return "BC-Linux 7.6 64位"
	case IMAGE_BCLINUX_75_X64:
		return "BC-Linux 7.5 64位"
	case IMAGE_BCLINUX_74_X64:
		return "BC-Linux 7.4 64位"
	case IMAGE_BCLINUX_71_X64:
		return "BC-Linux 7.1 64位"
	case IMAGE_BCLINUX_65_X64:
		return "BC-Linux 6.5 64位"
	default:
		return ""
	}
}

type VpcScale int32

const (
	//VPC规格
	VPC_SCALE_HIGH   VpcScale = 0
	VPC_SCALE_NORMAL VpcScale = 1
	VPC_SCALE_MIDDLE VpcScale = 2
)

func (v VpcScale) String() string {
	switch v {
	case VPC_SCALE_NORMAL:
		return "normal"
	case VPC_SCALE_HIGH:
		return "high"
	case VPC_SCALE_MIDDLE:
		return "middle"
	default:
		return "high"
	}
}
