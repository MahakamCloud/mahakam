package provisioner

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/mahakamcloud/mahakam/pkg/config"
	"github.com/mahakamcloud/mahakam/pkg/network"
	"github.com/mahakamcloud/mahakam/pkg/node"
	"github.com/mahakamcloud/mahakam/pkg/utils"
	"github.com/sirupsen/logrus"
)

type CreateNode struct {
	Config node.NodeCreateConfig
	p      Provisioner
	log    logrus.FieldLogger
}

func NewCreateNode(config node.NodeCreateConfig, p Provisioner, log logrus.FieldLogger) *CreateNode {
	createNodeLog := log.WithField("task", fmt.Sprintf("create node in %s", config.Host))

	return &CreateNode{
		Config: config,
		p:      p,
		log:    createNodeLog,
	}
}

func (n *CreateNode) Run() error {
	err := n.p.CreateNode(n.Config)
	if err != nil {
		n.log.Errorf("error creating node '%v': %s", n.Config, err)
		return err
	}
	return nil
}

type CheckClusterNetworkNodes struct {
	clusterNetwork *network.ClusterNetwork
	log            logrus.FieldLogger
}

func NewCheckClusterNetworkNodes(clusterNetwork *network.ClusterNetwork, log logrus.FieldLogger) *CheckClusterNetworkNodes {
	checkClusterNetworkNodesLog := log.WithField("task", fmt.Sprintf("check cluster network nodes in %v", clusterNetwork))

	return &CheckClusterNetworkNodes{
		clusterNetwork: clusterNetwork,
		log:            checkClusterNetworkNodesLog,
	}
}

func (c *CheckClusterNetworkNodes) Run() error {
	// Blocking check waiting for cluster gateway to be up
	gwReady := utils.ICMPPingNWithDelay(c.clusterNetwork.Gateway.String(), config.NodePingTimeout, c.log,
		config.NodePingRetry, config.NodePingDelay)

	// Cluster gateway still not ready after max retry
	if !gwReady {
		return fmt.Errorf("timeout waiting for cluster gateway to be up '%v'", c.clusterNetwork)
	}

	return nil
}

type CheckNode struct {
	ip  net.IP
	log logrus.FieldLogger
}

func NewCheckNode(ip net.IP, log logrus.FieldLogger) *CheckNode {
	checkNodeLog := log.WithField("task", fmt.Sprintf("check node with address %v", ip.String()))

	return &CheckNode{
		ip:  ip,
		log: checkNodeLog,
	}
}

func (c *CheckNode) Run() error {
	// Blocking check waiting for node to be up
	nodeReady := utils.ICMPPingNWithDelay(c.ip.String(), config.NodePingTimeout, c.log,
		config.NodePingRetry, config.NodePingDelay)

	// Cluster gateway still not ready after max retry
	if !nodeReady {
		return fmt.Errorf("timeout waiting for node to be up '%v'", c.ip)
	}

	return nil
}

type CreateAdminKubeconfig struct {
	clustername      string
	apiServerAddress string
	apiServerPort    string
	utils.SCPConfig
	log logrus.FieldLogger
}

func NewCreateAdminKubeconfig(clustername, apiServerAddress, apiServerPort string,
	config utils.SCPConfig) *CreateAdminKubeconfig {

	createAdminKubeconfigLog := logrus.WithField("task", fmt.Sprintf("copying kubeconfig from %s to local system", config.RemoteIPAddress))

	return &CreateAdminKubeconfig{
		clustername:      clustername,
		apiServerAddress: apiServerAddress,
		apiServerPort:    apiServerPort,
		SCPConfig:        config,
		log:              createAdminKubeconfigLog,
	}
}

func (k *CreateAdminKubeconfig) Run() error {
	// Blocking check waiting control plane to be up
	apiServer := fmt.Sprintf("%s:%s", k.apiServerAddress, k.apiServerPort)
	ready := utils.PortPingNWithDelay(apiServer, config.NodePingTimeout, k.log,
		config.NodePingRetry, config.NodePingDelay)

	// Control plane node still not up after max retry
	if !ready {
		return fmt.Errorf("timeout waiting for control plane to be up '%v'", k)
	}
	// TODO(giri): wait until kubeadm finishes bootstraping,
	// hardcoded wait time 120 sec in cloud init script
	time.Sleep(3 * config.NodePingDelay)

	err := os.MkdirAll(config.MahakamMultiKubeconfigPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating mahakam directory to hold kubeconfig files")
	}

	// Copy over kubeconfig generated by kubeadm to mahakam server
	s := utils.NewSCPClient()
	_, err = s.CopyRemoteFile(k.SCPConfig)
	if err != nil {
		return fmt.Errorf("error creating admin kubeconfig file for cluster %s '%v': %s", k.clustername, k.SCPConfig, err)
	}

	k.log.Infof("admin kubeconfig has been copied over successfully '%v'", k)
	return nil
}
