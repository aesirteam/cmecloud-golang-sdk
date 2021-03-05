package VirtualPrivateCloud

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
