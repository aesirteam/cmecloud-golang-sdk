package ecloud

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/ELB"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/FloatingIP"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/storage/CloudBlockStorage"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/Image"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/KeyPair"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/SecurityGroup"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/Server"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/ServerBackup"
	"github.com/caarlos0/env/v6"
)

type core struct {
	global.CoreInterface
}

type vm struct {
	Image.ImageInterface
	KeyPair.KeypairInterface
	SecurityGroup.SecurityGroupInterface
	Server.ServerInterface
	ServerBackup.ServerBackupInterface
}

type net struct {
	ELB.ELBInterface
	FloatingIP.EIPInterface
	VirtualPrivateCloud.VPCInterface
}

type storage struct {
	CloudBlockStorage.CBSInterface
}

type ClientSet struct {
	coreAPIv2 *global.APIv2

	//net
	elbAPIv2 *ELB.APIv2
	eipAPIv2 *FloatingIP.APIv2
	vpcAPIv2 *VirtualPrivateCloud.APIv2

	//vm
	imageAPIv2         *Image.APIv2
	keypairAPIv2       *KeyPair.APIv2
	securityGroupAPIv2 *SecurityGroup.APIv2
	serverAPIv2        *Server.APIv2
	serverBackupAPIv2  *ServerBackup.APIv2

	//storage
	cbsAPIv2 *CloudBlockStorage.APIv2
}

func (c *ClientSet) Core() *core {
	return &core{c.coreAPIv2}
}

func (c *ClientSet) Net() *net {
	return &net{c.elbAPIv2, c.eipAPIv2, c.vpcAPIv2}
}

func (c *ClientSet) VM() *vm {
	return &vm{c.imageAPIv2, c.keypairAPIv2, c.securityGroupAPIv2, c.serverAPIv2, c.serverBackupAPIv2}
}

func (c *ClientSet) Storage() *storage {
	return &storage{c.cbsAPIv2}
}

func NewForConfig(conf *global.Config) (*ClientSet, error) {
	var err error
	if err = env.Parse(conf); err != nil {
		return nil, err
	}

	var cs ClientSet

	if cs.coreAPIv2, err = global.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.elbAPIv2, err = ELB.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.eipAPIv2, err = FloatingIP.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.vpcAPIv2, err = VirtualPrivateCloud.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.imageAPIv2, err = Image.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.keypairAPIv2, err = KeyPair.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.securityGroupAPIv2, err = SecurityGroup.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.serverAPIv2, err = Server.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.serverBackupAPIv2, err = ServerBackup.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.cbsAPIv2, err = CloudBlockStorage.NewForConfig(conf); err != nil {
		return nil, err
	}

	return &cs, nil
}

func NewForConfigDie(conf *global.Config) *ClientSet {
	client, err := NewForConfig(conf)
	if err != nil {
		panic(err)
	}
	return client
}
