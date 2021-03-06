package resource

import (
	"strconv"
	"strings"

	"github.com/mahakamcloud/mahakam/pkg/config"
	"github.com/mahakamcloud/mahakam/pkg/utils"
)

const (
	// ClusterSizeExtraSmall is represents cluster with extra small size nodes
	ClusterSizeExtraSmall = "EXTRASMALL"
	// ClusterSizeSmall is represents cluster with small size nodes
	ClusterSizeSmall = "SMALL"
	// ClusterSizeMedium is represents cluster with medium size nodes
	ClusterSizeMedium = "MEDIUM"
	// ClusterSizeLarge is represents cluster with large size nodes
	ClusterSizeLarge = "LARGE"
	// ClusterSizeExtraLarge is represents cluster with extra large size nodes
	ClusterSizeExtraLarge = "EXTRALARGE"
	// ClusterSizeDefault is represents cluster with default size nodes
	ClusterSizeDefault = "DEFAULT"
)

// clusterNodeSizes represents available node size configurations.
var (
	clusterNodeSizes = map[string]map[string]string{
		ClusterSizeExtraSmall: map[string]string{"cpu": "2", "ram": "4GB"},
		ClusterSizeSmall:      map[string]string{"cpu": "4", "ram": "16GB"},
		ClusterSizeMedium:     map[string]string{"cpu": "8", "ram": "32GB"},
		ClusterSizeLarge:      map[string]string{"cpu": "16", "ram": "64GB"},
		ClusterSizeExtraLarge: map[string]string{"cpu": "32", "ram": "128GB"},
		ClusterSizeDefault:    map[string]string{"cpu": "2", "ram": "4GB"},
	}

	availableClusterSizes = []string{ClusterSizeDefault, ClusterSizeExtraSmall, ClusterSizeSmall, ClusterSizeMedium, ClusterSizeLarge, ClusterSizeExtraLarge}
)

// GetClusterNodeCPUs returns number of CPUs for a cluster node
func GetClusterNodeCPUs(size string) string {
	size = strings.ToUpper(size)
	cpuInString := clusterNodeSizes[size]["cpu"]
	return cpuInString
}

// GetClusterNodeMemoryInMB returns memory for a cluster node in bytes from default GB representationss
func GetClusterNodeMemoryInMB(size string) (string, error) {
	size = strings.ToUpper(size)

	memoryInGB := clusterNodeSizes[size]["ram"]

	memoryInMB, err := utils.ToMegabytes(memoryInGB)
	if err != nil {
		// TODO: Figure out a better return value as default
		return "", err
	}
	memory := strconv.FormatUint(memoryInMB, 10)
	return memory, nil
}

func ClusterSizeValidate(size string) bool {
	for _, availableSize := range availableClusterSizes {
		if strings.ToUpper(size) == availableSize {
			return true
		}
	}
	return false
}

// Cluster represents stored resource with cluster kind
type Cluster struct {
	BaseResource
	NodeSize       string
	NumNodes       int
	NetworkName    string
	KubeconfigPath string
}

// NewCluster creates new resource cluster
func NewCluster(name string) *Cluster {
	return &Cluster{
		BaseResource: BaseResource{
			Name:  name,
			Kind:  string(KindCluster),
			Owner: config.ResourceOwnerGojek,
		},
	}
}

// ClusterList represents a group of Clusters
type ClusterList struct {
	Items []*Cluster
}

// Resource returns a empty Cluster
func (l *ClusterList) Resource() Resource {
	return &Cluster{}
}

// WithItems returns list of Cluster
func (l *ClusterList) WithItems(items []Resource) {
	for _, i := range items {
		cluster := i.(*Cluster)
		l.Items = append(l.Items, cluster)
	}
}
