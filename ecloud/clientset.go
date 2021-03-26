package ecloud

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/ELB"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/FloatingIP"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/IPSecVpn"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/storage/CloudBlockStorage"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/Image"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/KeyPair"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/SecurityGroup"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/Server"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/ServerBackup"
	"github.com/caarlos0/env/v6"
)

type ClientSet struct {
	//net
	elbAPIv2 *ELB.APIv2
	fipAPIv2 *FloatingIP.APIv2
	vpcAPIv2 *VirtualPrivateCloud.APIv2
	vpnAPIv2 *IPSecVpn.APIv2
	//vm
	imageAPIv1         *Image.APIv1
	keypairAPIv2       *KeyPair.APIv2
	securityGroupAPIv1 *SecurityGroup.APIv1
	serverAPIv2        *Server.APIv2
	serverBackupAPIv1  *ServerBackup.APIv1
	//storage
	cbsAPIv2 *CloudBlockStorage.APIv2
}

func NewForConfig(conf *global.Config) (*ClientSet, error) {
	if err := env.Parse(conf); err != nil {
		return nil, err
	}

	client, err := global.NewEcloudClient(conf)
	if err != nil {
		return nil, err
	}

	return &ClientSet{
		elbAPIv2:           ELB.New(client),
		fipAPIv2:           FloatingIP.New(client),
		vpcAPIv2:           VirtualPrivateCloud.New(client),
		vpnAPIv2:           IPSecVpn.New(client),
		imageAPIv1:         Image.New(client),
		keypairAPIv2:       KeyPair.New(client),
		securityGroupAPIv1: SecurityGroup.New(client),
		serverAPIv2:        Server.New(client),
		serverBackupAPIv1:  ServerBackup.New(client),
		cbsAPIv2:           CloudBlockStorage.New(client),
	}, nil
}

func NewForConfigDie(conf *global.Config) *ClientSet {
	client, err := NewForConfig(conf)
	if err != nil {
		panic(err)
	}
	return client
}

type vmInterface struct {
	Image.ImageInterface
	KeyPair.KeypairInterface
	SecurityGroup.SecurityGroupInterface
	Server.ServerInterface
	ServerBackup.ServerBackupInterface
}

type netInterface struct {
	ELB.ELBInterface
	FloatingIP.FIPInterface
	VirtualPrivateCloud.VPCInterface
	IPSecVpn.VPNInterface
}

type storageInterface struct {
	CloudBlockStorage.CBSInterface
}

func (cs *ClientSet) Net() *netInterface {
	return &netInterface{cs.elbAPIv2, cs.fipAPIv2, cs.vpcAPIv2, cs.vpnAPIv2}
}

func (cs *ClientSet) VM() *vmInterface {
	return &vmInterface{cs.imageAPIv1, cs.keypairAPIv2, cs.securityGroupAPIv1, cs.serverAPIv2, cs.serverBackupAPIv1}
}

func (cs *ClientSet) Storage() *storageInterface {
	return &storageInterface{cs.cbsAPIv2}
}
