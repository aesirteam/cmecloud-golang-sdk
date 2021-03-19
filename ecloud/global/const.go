package global

type BootVolumeType int32
type DataVolumeType int32
type VmType int32
type OsType int32
type BillingType int32
type ImageFormat int32
type ImageName int32

const (
	// Linux操作系统
	OS_TYPE_LINUX OsType = 0
	// Windows操作系统
	OS_TYPE_WINDOWS OsType = 1

	// 性能优化型系统盘
	BOOT_VOLUME_PERFORMANCEOPTIMIZATION BootVolumeType = 0
	// 高性能型系统盘
	BOOT_VOLUME_HIGHPERFORMANCE BootVolumeType = 1

	// 容量型云硬盘
	DATA_VOLUME_CAPEBS DataVolumeType = 0
	// 性能优化型云硬盘
	DATA_VOLUME_SSDEBS DataVolumeType = 1
	// 高性能型云硬盘
	DATA_VOLUME_SSD DataVolumeType = 2

	// 通用型
	VM_TYPE_COMMON VmType = 0
	// 通用入门型
	VM_TYPE_COMMONINTRODUCTORY VmType = 1
	// 通用网络优化型
	VM_TYPE_COMMONNETIMPROVE VmType = 2
	// 内存优化型
	VM_TYPE_MEMIMPROVE VmType = 3
	// 计算型
	VM_TYPE_COMPUTE VmType = 4
	// 计算网络优化型
	VM_TYPE_COMPUTENETIMPROVE VmType = 5
	// 内存网络优化型
	VM_TYPE_MEMNETIMPROVE VmType = 6
	// 本地存储型
	VM_TYPE_LOCALSTORAGE VmType = 7
	// 超大内存型
	VM_TYPE_XLARGEMEMORY VmType = 8
	// 超高主频型
	VM_TYPE_HIGHFREQUENCY VmType = 9
	// FPGA型
	VM_TYPE_FPGA VmType = 10
	// GPU型
	VM_TYPE_GPU VmType = 11
	// VGPU型
	VM_TYPE_VGPU VmType = 12
	// 高IO型
	VM_TYPE_HIGHIO VmType = 13
	// 独享型
	VM_TYPE_EXCLUSIVE VmType = 14

	// 按时计费
	BILLING_TYPE_HOUR BillingType = 0
	// 按月计费
	BILLING_TYPE_MONTH BillingType = 1
	// 按年计费
	BILLING_TYPE_YEAR BillingType = 2

	// 镜像格式
	IMAGE_FORMAT_QCOW2 ImageFormat = 0
	IMAGE_FORMAT_VHD   ImageFormat = 1
	IMAGE_FORMAT_RAW   ImageFormat = 2

	// 云主机镜像
	IMAGE_BCLINUX_81_X64 ImageName = 1
	IMAGE_BCLINUX_77_X64 ImageName = 2
	IMAGE_BCLINUX_76_X64 ImageName = 3
	IMAGE_BCLINUX_75_X64 ImageName = 4
	IMAGE_BCLINUX_74_X64 ImageName = 5
	IMAGE_BCLINUX_71_X64 ImageName = 6
	IMAGE_BCLINUX_65_X64 ImageName = 7
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

func (t ImageFormat) String() string {
	switch t {
	case IMAGE_FORMAT_QCOW2:
		return "qcow2"
	case IMAGE_FORMAT_VHD:
		return "vhd"
	case IMAGE_FORMAT_RAW:
		return "raw"
	default:
		return "qcow2"
	}
}

func (t ImageName) String() string {
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

type SecurityGroupDirection int32
type SecurityGroupEtherType int32
type SecurityGroupProtocol int32
type SecurityGroupRemoteType int32

const (
	// 安全组方向(流入)
	SECURITYGROUP_DIRECTION_INGRESS SecurityGroupDirection = 0
	// 安全组方向(流出)
	SECURITYGROUP_DIRECTION_EGRESS SecurityGroupDirection = 1

	// 安全组IP类型
	SECURITYGROUP_ETHERTYPE_IPV4 SecurityGroupEtherType = 0
	SECURITYGROUP_ETHERTYPE_IPV6 SecurityGroupEtherType = 1

	// 安全组协议(TCP)
	SECURITYGROUP_PROTOCOL_TCP SecurityGroupProtocol = 0
	// 安全组协议(UDP)
	SECURITYGROUP_PROTOCOL_UDP SecurityGroupProtocol = 1
	// 安全组协议(ICMP)
	SECURITYGROUP_PROTOCOL_ICMP SecurityGroupProtocol = 2
	// 安全组协议(IGMP)
	SECURITYGROUP_PROTOCOL_IGMP SecurityGroupProtocol = 3
	// 安全组协议(ANY)
	SECURITYGROUP_PROTOCOL_ANY SecurityGroupProtocol = 4

	// 安全组授权类型
	SECURITYGROUP_REMOTETYPE_CIDR          SecurityGroupRemoteType = 0
	SECURITYGROUP_REMOTETYPE_SECURITYGROUP SecurityGroupRemoteType = 1
)

func (s SecurityGroupDirection) String() string {
	switch s {
	case SECURITYGROUP_DIRECTION_INGRESS:
		return "ingress"
	case SECURITYGROUP_DIRECTION_EGRESS:
		return "egress"
	default:
		return "ingress"
	}
}

func (s SecurityGroupEtherType) String() string {
	switch s {
	case SECURITYGROUP_ETHERTYPE_IPV4:
		return "IPv4"
	case SECURITYGROUP_ETHERTYPE_IPV6:
		return "IPv6"
	default:
		return "IPv6"
	}
}

func (s SecurityGroupProtocol) String() string {
	switch s {
	case SECURITYGROUP_PROTOCOL_TCP:
		return "TCP"
	case SECURITYGROUP_PROTOCOL_UDP:
		return "UDP"
	case SECURITYGROUP_PROTOCOL_ICMP:
		return "ICMP"
	case SECURITYGROUP_PROTOCOL_IGMP:
		return "IGMP"
	case SECURITYGROUP_PROTOCOL_ANY:
		return "ANY"
	default:
		return "TCP"
	}
}

func (s SecurityGroupRemoteType) String() string {
	switch s {
	case SECURITYGROUP_REMOTETYPE_CIDR:
		return "cidr"
	case SECURITYGROUP_REMOTETYPE_SECURITYGROUP:
		return "security_group"
	default:
		return "cidr"
	}
}

type VpcScale int32

const (
	// VPC规格
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

type FloatingIpChargeMode int32
type FloatingIpQuotaName int32
type FloatingIpBindStatus int32
type FloatingIpBindType int32

const (
	// 弹性公网IP按流量计费
	FLOATINGIP_CHARGEMODE_TRAFFIC FloatingIpChargeMode = 0
	// 弹性公网IP按带宽计费方式
	FLOATINGIP_CHARGEMODE_BANDWIDTH FloatingIpChargeMode = 1

	// 弹性公网IP配额名称
	FLOATINGIP_QUOTANAME_IP        FloatingIpQuotaName = 0
	FLOATINGIP_QUOTANAME_BANDWIDTH FloatingIpQuotaName = 1

	// 弹性公网IP绑定状态
	// 未绑定
	FLOATINGIP_BINDSTATUS_UNBOUND FloatingIpBindStatus = 0
	// 绑定
	FLOATINGIP_BINDSTATUS_BOUND FloatingIpBindStatus = 1
	// 冻结
	FLOATINGIP_BINDSTATUS_FROZEN FloatingIpBindStatus = 2

	// 弹性公网IP绑定类型
	// 虚拟机
	FLOATINGIP_BINDTYPE_VM FloatingIpBindType = 0
	// 裸金属
	FLOATINGIP_BINDTYPE_IRONIC FloatingIpBindType = 1
	// 负载均衡
	FLOATINGIP_BINDTYPE_ELB FloatingIpBindType = 2
	// 共享带宽
	FLOATINGIP_BINDTYPE_SBW FloatingIpBindType = 3
	// 云数据库
	FLOATINGIP_BINDTYPE_RDS FloatingIpBindType = 4
	// HAVIP
	FLOATINGIP_BINDTYPE_HAVIP FloatingIpBindType = 5
	// 云堡垒机
	FLOATINGIP_BINDTYPE_BASTION FloatingIpBindType = 6
	// SNAT
	FLOATINGIP_BINDTYPE_SNAT FloatingIpBindType = 7
	// DNAT
	FLOATINGIP_BINDTYPE_DNAT FloatingIpBindType = 8
	// 虚拟网卡
	FLOATINGIP_BINDTYPE_PORT FloatingIpBindType = 9
)

func (f FloatingIpChargeMode) String() string {
	switch f {
	case FLOATINGIP_CHARGEMODE_TRAFFIC:
		return "trafficCharge"
	case FLOATINGIP_CHARGEMODE_BANDWIDTH:
		return "bandwidthCharge"
	default:
		return "trafficCharge"
	}
}

func (f FloatingIpQuotaName) String() string {
	switch f {
	case FLOATINGIP_QUOTANAME_IP:
		return "FLOATING_IP"
	case FLOATINGIP_QUOTANAME_BANDWIDTH:
		return "CAPACITY_OF_BANDWIDTH"
	default:
		return "FLOATING_IP"
	}
}

func (f FloatingIpBindStatus) String() string {
	switch f {
	case FLOATINGIP_BINDSTATUS_UNBOUND:
		return "UNBOUND"
	case FLOATINGIP_BINDSTATUS_BOUND:
		return "BINDING"
	case FLOATINGIP_BINDSTATUS_FROZEN:
		return "FROZEN"
	default:
		return "UNBOUND"
	}
}

func (f FloatingIpBindType) String() string {
	switch f {
	case FLOATINGIP_BINDTYPE_VM:
		return "vm"
	case FLOATINGIP_BINDTYPE_IRONIC:
		return "ironic"
	case FLOATINGIP_BINDTYPE_ELB:
		return "elb"
	case FLOATINGIP_BINDTYPE_SBW:
		return "sbw"
	case FLOATINGIP_BINDTYPE_RDS:
		return "rds"
	case FLOATINGIP_BINDTYPE_HAVIP:
		return "havip"
	case FLOATINGIP_BINDTYPE_BASTION:
		return "bastion"
	case FLOATINGIP_BINDTYPE_SNAT:
		return "snat"
	case FLOATINGIP_BINDTYPE_DNAT:
		return "dnat"
	case FLOATINGIP_BINDTYPE_PORT:
		return "port"
	default:
		return "vm"
	}
}

type VpnAuthAlgorithm int32
type VpnEncryptionAlgorithm int32
type VpnPfs int32
type VpnEncapsulationMode int32
type VpnIkeVersion int32
type VpnIkePhase1NegotiationMode int32

const (
	//  认证算法
	VPN_AUTHALGORITHM_SHA1   VpnAuthAlgorithm = 0
	VPN_AUTHALGORITHM_SHA256 VpnAuthAlgorithm = 1
	VPN_AUTHALGORITHM_SHA512 VpnAuthAlgorithm = 2

	// 加密算法
	VPN_ENCRYPTIONALGORITHM_AES128 VpnEncryptionAlgorithm = 0
	VPN_ENCRYPTIONALGORITHM_DES3   VpnEncryptionAlgorithm = 1

	// 完美向前加密
	VPN_PFS_GROUP2 VpnPfs = 0
	VPN_PFS_GROUP5 VpnPfs = 1

	// 传输封装模式
	VPN_ENCAPSULATIONMODE_TUNNEL    VpnEncapsulationMode = 0
	VPN_ENCAPSULATIONMODE_TRANSPORT VpnEncapsulationMode = 1

	// IKE策略版本
	VPN_IKE_VERSION_V1 VpnIkeVersion = 0
	VPN_IKE_VERSION_V2 VpnIkeVersion = 1

	// 协商模式
	VPN_IKE_PHASE1NEGOTIATIONMODE_MAIN       VpnIkePhase1NegotiationMode = 0
	VPN_IKE_PHASE1NEGOTIATIONMODE_AGGRESSIVE VpnIkePhase1NegotiationMode = 1
)

func (v VpnAuthAlgorithm) String() string {
	switch v {
	case VPN_AUTHALGORITHM_SHA1:
		return "sha1"
	case VPN_AUTHALGORITHM_SHA256:
		return "sha256"
	case VPN_AUTHALGORITHM_SHA512:
		return "sha512"
	default:
		return "sha1"
	}
}

func (v VpnEncryptionAlgorithm) String() string {
	switch v {
	case VPN_ENCRYPTIONALGORITHM_AES128:
		return "AES_128"
	case VPN_ENCRYPTIONALGORITHM_DES3:
		return "DES3"
	default:
		return "AES_128"
	}
}

func (v VpnPfs) String() string {
	switch v {
	case VPN_PFS_GROUP2:
		return "group2"
	case VPN_PFS_GROUP5:
		return "group5"
	default:
		return "group2"
	}
}

func (v VpnEncapsulationMode) String() string {
	switch v {
	case VPN_ENCAPSULATIONMODE_TUNNEL:
		return "tunnel"
	case VPN_ENCAPSULATIONMODE_TRANSPORT:
		return "transport"
	default:
		return "tunnel"
	}
}

func (v VpnIkeVersion) String() string {
	switch v {
	case VPN_IKE_VERSION_V1:
		return "v1"
	case VPN_IKE_VERSION_V2:
		return "v2"
	default:
		return "v1"
	}
}

func (v VpnIkePhase1NegotiationMode) String() string {
	switch v {
	case VPN_IKE_PHASE1NEGOTIATIONMODE_MAIN:
		return "main"
	case VPN_IKE_PHASE1NEGOTIATIONMODE_AGGRESSIVE:
		return "aggressive"
	default:
		return "main"
	}
}
