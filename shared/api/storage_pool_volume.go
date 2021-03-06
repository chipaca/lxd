package api

// StorageVolumesPost represents the fields of a new LXD storage pool volume
//
// API extension: storage
type StorageVolumesPost struct {
	StorageVolumePut `yaml:",inline"`

	Name string `json:"name" yaml:"name"`
	Type string `json:"type" yaml:"type"`

	// API extension: storage_api_local_volume_handling
	Source StorageVolumeSource `json:"source" yaml:"source"`
}

// StorageVolumePost represents the fields required to rename a LXD storage pool volume
//
// API extension: storage_api_volume_rename
type StorageVolumePost struct {
	Name string `json:"name" yaml:"name"`

	// API extension: storage_api_local_volume_handling
	Pool string `json:"pool,omitempty" yaml:"pool,omitempty"`
}

// StorageVolume represents the fields of a LXD storage volume.
//
// API extension: storage
type StorageVolume struct {
	StorageVolumePut `yaml:",inline"`
	Name             string   `json:"name" yaml:"name"`
	Type             string   `json:"type" yaml:"type"`
	UsedBy           []string `json:"used_by" yaml:"used_by"`
}

// StorageVolumePut represents the modifiable fields of a LXD storage volume.
//
// API extension: storage
type StorageVolumePut struct {
	Config map[string]string `json:"config" yaml:"config"`

	// API extension: entity_description
	Description string `json:"description" yaml:"description"`
}

// StorageVolumeSource represents the creation source for a new storage volume.
//
// API extension: storage_api_local_volume_handling
type StorageVolumeSource struct {
	Name string `json:"name" yaml:"name"`
	Type string `json:"type" yaml:"type"`
	Pool string `json:"pool" yaml:"pool"`
}

// Writable converts a full StorageVolume struct into a StorageVolumePut struct
// (filters read-only fields).
func (storageVolume *StorageVolume) Writable() StorageVolumePut {
	return storageVolume.StorageVolumePut
}
