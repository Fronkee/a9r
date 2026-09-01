package models

import "time"

type Ec2instance struct {
	ID        string
	State     string
	Type      string
	Name      string
	AZ        string
	PrivateIP string
	PublicIP  string
}

type S3Bucket struct {
	Name         string
	CreationDate time.Time
}

type EC2Detail struct {
	// ==================================================
	// BASIC
	// ==================================================
	InstanceID       string
	Name             string
	InstanceType     string
	State            string
	StateCode        int32
	StateReason      string
	AMI              string
	Architecture     string
	Platform         string
	PlatformDetails  string
	UsageOperation   string
	LaunchTime       time.Time
	// ==================================================
	// COMPUTE
	// ==================================================
	VCPU             int
	MemoryMB         int
	CPUCoreCount     int
	ThreadsPerCore   int
	CPUOptions       string
	Hypervisor       string
	ENAEnabled       bool
	NitroEnclaves    bool
	EBSOptimized      bool
	// ==================================================
	// NETWORK
	// ==================================================
	PrivateIP        string
	PublicIP         string
	PrivateDNS       string
	PublicDNS        string
	IPv6Addresses    []string
	VPCID            string
	SubnetID         string
	AvailabilityZone string
	AvailabilityZoneID string
	// ==================================================
	// SECURITY
	// ==================================================
	KeyName          string
	IAMRole          string
	IAMInstanceProfile string
	SecurityGroups   []SecurityGroup
	NetworkInterfaces []NetworkInterface
	// ==================================================
	// STORAGE
	// ==================================================
	RootDeviceName   string
	RootDeviceType   string
	BlockDevices     []BlockDevice
	// ==================================================
	// PLACEMENT
	// ==================================================
	Tenancy          string
	PlacementGroup   string
	HostID           string
	HostResourceGroupARN string
	PartitionNumber  int
	// ==================================================
	// MONITORING
	// ==================================================
	MonitoringState  string
	// ==================================================
	// LIFECYCLE
	// ==================================================
	InstanceLifecycle string
	SpotInstanceRequestID string
	TerminationProtection bool
	StopProtection        bool
	// ==================================================
	// METADATA
	// ==================================================
	MetadataHTTPToken   string
	MetadataHTTPEndpoint string
	MetadataHopLimit    int
	// ==================================================
	// TAGS
	// ==================================================
	Tags []Tag
}

type SecurityGroup struct {
	ID   string
	Name string
}

type NetworkInterface struct {
	ID              string
	Description     string
	PrivateIP       string
	PublicIP        string
	PrivateDNS      string
	SubnetID        string
	VPCID           string
	Status          string
	MacAddress      string
	IPv6Addresses   []string
	SecurityGroups  []SecurityGroup
}

type BlockDevice struct {
	DeviceName string
	VolumeID   string
	Status     string
	SizeGB     int
	VolumeType string
	IOPS       int
	Throughput int
	Encrypted  bool
	DeleteOnTermination bool
}

type Tag struct {
	Key   string
	Value string
}



type S3BucketDetail struct {
	Name            string
	Region          string
	CreationDate    time.Time
	ObjectCount     int64
	SizeBytes       int64
	Versioning      string
	Encryption      string
	ObjectLock      string
	PublicAccess    string
	ObjectOwnership string
	ACL             string
	Policy          string
	LifecycleRules  int
	Replication     string
	AccessLogging   string
	Tags            map[string]string
}
