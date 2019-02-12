package provisioner_test

import (
	"fmt"
	"io/ioutil"
	"net"
	"testing"

	"github.com/mahakamcloud/mahakam/pkg/network"

	"github.com/mahakamcloud/mahakam/pkg/utils"
	"github.com/mahakamcloud/mahakam/pkg/validation"

	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"

	"github.com/mahakamcloud/mahakam/pkg/node"
	. "github.com/mahakamcloud/mahakam/pkg/provisioner"
	store "github.com/mahakamcloud/mahakam/pkg/resource_store"

	gomock "github.com/golang/mock/gomock"
)

var (
	n = node.NodeCreateConfig{
		Host: net.ParseIP("10.10.10.10"),
	}

	cn = &network.ClusterNetwork{
		Name: "fake cluster network",
	}

	clusterName = "fake-cluster-name"
)

func nilLogger() *logrus.Logger {
	l := logrus.New()
	l.Out = ioutil.Discard

	return l
}

func TestPreCreateCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := nilLogger()
	s := store.NewMockResourceStore(ctrl)

	tests := []struct {
		name               string
		expectClusterExist bool
		expectError        error
	}{
		{
			name:               "test should return error if cluster exists",
			expectClusterExist: true,
			expectError:        fmt.Errorf("cluster %s already exists", clusterName),
		},
		{
			name:               "test should not return error if cluster doesn't exist",
			expectClusterExist: false,
			expectError:        nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s.EXPECT().KeyExists(gomock.Any()).Return(test.expectClusterExist)

			pc := NewPreCreateCheck(clusterName, l, s)
			err := pc.Run()

			assert.Equal(t, test.expectError, err)
		})
	}
}

func TestCreateNode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := nilLogger()
	p := NewMockProvisioner(ctrl)

	tests := []struct {
		name        string
		expectError error
	}{
		{
			name:        "test create node where provisioner successfully runs",
			expectError: nil,
		},
		{
			name:        "test create node where provisioner fails",
			expectError: fmt.Errorf("fake error from provisioner"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p.EXPECT().CreateNode(gomock.Any()).Return(test.expectError)

			cn := NewCreateNode(n, p, l)
			err := cn.Run()

			assert.Equal(t, test.expectError, err)
		})
	}
}

func TestCheckClusterNetworkNodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name        string
		expectError error
		expectReady bool
	}{
		{
			name:        "test check cluster network nodes where gateway is ready",
			expectError: nil,
			expectReady: true,
		},
		{
			name:        "test check cluster network nodes where gateway is not ready",
			expectError: fmt.Errorf("timeout waiting for cluster gateway to be up"),
			expectReady: false,
		},
	}

	l := nilLogger()
	pc := utils.NewMockPingChecker(ctrl)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pc.EXPECT().ICMPPingNWithDelay(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(test.expectReady)

			c := NewCheckClusterNetworkNodes(cn, l, pc)
			err := c.Run()

			if err != nil {
				assert.Error(t, err, test.expectError)
			}

			if err == nil {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCheckNode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name        string
		nodeIP      net.IP
		expectError error
		expectReady bool
	}{
		{
			name:        "test check node where node is ready",
			nodeIP:      net.ParseIP("1.2.3.4"),
			expectError: nil,
			expectReady: true,
		},
		{
			name:        "test check node where node is not ready",
			nodeIP:      net.ParseIP("1.2.3.4"),
			expectError: fmt.Errorf("timeout waiting for node to be up"),
			expectReady: false,
		},
	}

	l := nilLogger()
	pc := utils.NewMockPingChecker(ctrl)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pc.EXPECT().ICMPPingNWithDelay(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(test.expectReady)

			c := NewCheckNode(test.nodeIP, l, pc)
			err := c.Run()

			if err != nil {
				assert.Error(t, err, test.expectError)
			}

			if err == nil {
				assert.NoError(t, err)
			}
		})
	}
}

func TestClusterValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name        string
		expectError error
		expectReady bool
	}{
		{
			name:        "test cluster validation where cluster is ready",
			expectError: nil,
			expectReady: true,
		},
	}

	l := nilLogger()
	cv := validation.NewMockValidator(ctrl)
	s := store.NewMockResourceStore(ctrl)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cv.EXPECT().ValidateNWithDelay(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(test.expectReady)
			s.EXPECT().Get(gomock.Any()).
				Return(nil)
			s.EXPECT().Update(gomock.Any()).
				Return(int64(0), nil)

			v := NewClusterValidation("fake-owner", "fake-cluster", cv, s, l)
			err := v.Run()

			if err != nil {
				assert.Error(t, err, test.expectError)
			}
			if err == nil {
				assert.NoError(t, err)
			}
		})
	}
}
