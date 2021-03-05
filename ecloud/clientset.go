package ecloud

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/ELB"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/FloatingIP"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/Image"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/KeyPair"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/SecurityGroup"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/Server"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/ServerBackup"
	"github.com/caarlos0/env/v6"
)

type ClientSet struct {
	//net
	elb *ELB.APIv2
	fip *FloatingIP.APIv2
	vpc *VirtualPrivateCloud.APIv2

	//vm
	image         *Image.APIv2
	keypair       *KeyPair.APIv2
	securityGroup *SecurityGroup.APIv2
	server        *Server.APIv2
	serverBackup  *ServerBackup.APIv2
}

func NewForConfig(conf *core.Config) (*ClientSet, error) {
	var err error
	if err = env.Parse(conf); err != nil {
		return nil, err
	}

	var cs ClientSet

	if cs.elb, err = ELB.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.fip, err = FloatingIP.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.vpc, err = VirtualPrivateCloud.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.image, err = Image.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.keypair, err = KeyPair.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.securityGroup, err = SecurityGroup.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.server, err = Server.NewForConfig(conf); err != nil {
		return nil, err
	}

	if cs.serverBackup, err = ServerBackup.NewForConfig(conf); err != nil {
		return nil, err
	}

	return &cs, nil
}

func NewForConfigDie(conf *core.Config) *ClientSet {
	client, err := NewForConfig(conf)
	if err != nil {
		panic(err)
	}
	return client
}

func (c *ClientSet) ELB() ELB.Interface {
	return c.elb
}

func (c *ClientSet) FloatingIP() FloatingIP.Interface {
	return c.fip
}

func (c *ClientSet) VPC() VirtualPrivateCloud.Interface {
	return c.vpc
}

func (c *ClientSet) Image() Image.Interface {
	return c.image
}

func (c *ClientSet) Keypair() KeyPair.Interface {
	return c.keypair
}

func (c *ClientSet) SecurityGroup() SecurityGroup.Interface {
	return c.securityGroup
}

func (c *ClientSet) Server() Server.Interface {
	return c.server
}

func (c *ClientSet) ServerBackup() ServerBackup.Interface {
	return c.serverBackup
}
