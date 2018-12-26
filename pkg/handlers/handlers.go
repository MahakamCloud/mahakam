package handlers

import (
	"github.com/mahakamcloud/mahakam/pkg/config"
	"github.com/mahakamcloud/mahakam/pkg/network"
	"github.com/mahakamcloud/mahakam/pkg/provisioner"
	store "github.com/mahakamcloud/mahakam/pkg/resource_store"
	log "github.com/sirupsen/logrus"
)

// Handlers holds common modules that each handler needs
type Handlers struct {
	Store       store.ResourceStore
	Network     *network.NetworkManager
	Provisioner provisioner.Provisioner
}

// New creates new handlers
func New(storeConfig config.StorageBackendConfig, networkConfig config.NetworkConfig,
	provisioner provisioner.Provisioner) *Handlers {
	rs, err := store.New(storeConfig)
	if err != nil {
		log.Fatalf("Error initializing resource store in handlers: %s", err)
		return nil
	}

	n, err := network.New(rs, networkConfig)
	if err != nil {
		log.Fatalf("Error initializing network in handlers: %s", err)
		return nil
	}

	return &Handlers{
		Store:       rs,
		Network:     n,
		Provisioner: provisioner,
	}
}
