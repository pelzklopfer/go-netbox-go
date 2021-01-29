package models

import (
	"encoding/json"
	"github.com/sapcc/go-netbox-go/common"
)

type VMInterface struct {
	Enabled          bool          `json:"enabled"`
	Url        		 string        `json:"url"`
	Id               int           `json:"id"`
	Name             string        `json:"name"`
	MTU              int           `json:"mtu"`
	MacAddress       string        `json:"mac_address"`
	Description      string        `json:"description"`
	Mode             string		   `json:"mode"`
	Tags             []NestedTag   `json:"tags"`
	VirtualMachine 	 NestedVirtualMachine
	UntaggedVlan     NestedVLAN
	TaggedVlans 	 []NestedVLAN
}

type WritableVMInterface struct {
	Id 				int				`json:"id,omitempty"`
	Url 			string			`json:"url,omitempty"`
	VirtualMachine 	int				`json:"virtual_machine,omitempty"`
	Name 			string			`json:"name"`
	Enabled 		bool			`json:"enabled,omitempty"`
	MTU 			int				`json:"mtu,omitempty"`
	MacAddress 		*string			`json:"mac_address,omitempty"`
	Description 	string			`json:"description,omitempty"`
	Mode 			string			`json:"mode,omitempty"`
	UntaggedVlan 	*int			`json:"untagged_vlan,omitempty"`
	TaggedVlans 	[]NestedVLAN	`json:"tagged_vlans,omitempty"`
	Tags        	[]NestedTag		`json:"tags,omitempty"`
}

type ListVMInterfacesRequest struct {
	common.ListParams
	VmId int
}

type ListVMInterfacesResponse struct {
	common.ReturnValues
	Results []VMInterface `json:"results"`
}

func (vmIf *VMInterface) UnmarshalJSON(b []byte) error {
	var tmp map[string]json.RawMessage
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	var tags []NestedTag
	if err := json.Unmarshal(tmp["tags"], &tags); err != nil {
		return err
	}
	vmIf.Tags = tags
	var vm NestedVirtualMachine
	if err := json.Unmarshal(tmp["virtual_machine"], &vm); err != nil {
			return err
	}
	vmIf.VirtualMachine = vm
	var tagVlan []NestedVLAN
	if err := json.Unmarshal(tmp["tagged_vlans"],&tagVlan); err != nil {
		return err
	}
	vmIf.TaggedVlans = tagVlan
	var untagVlan NestedVLAN
	if err := json.Unmarshal(tmp["untagged_vlan"],&untagVlan); err != nil {
		return err
	}
	vmIf.UntaggedVlan = untagVlan
	var id int
	if err := json.Unmarshal(tmp["id"], &id); err != nil {
		return err
	}
	vmIf.Id = id
	var url string
	if err := json.Unmarshal(tmp["url"], &url); err != nil {
		return err
	}
	vmIf.Url = url
	var Name string
	if err := json.Unmarshal(tmp["name"], &Name); err != nil {
		return err
	}
	vmIf.Name = Name
	var descr string
	if err := json.Unmarshal(tmp["description"],&descr); err != nil {
		return err
	}
	vmIf.Description = descr
	var mac string
	if err := json.Unmarshal(tmp["mac_address"],&mac); err != nil {
		return err
	}
	vmIf.MacAddress = mac
	var mtu int
	if err := json.Unmarshal(tmp["mtu"],&mtu); err != nil {
		return err
	}
	vmIf.MTU = mtu
	var en bool
	if err := json.Unmarshal(tmp["enabled"],&en); err != nil {
		return err
	}
	vmIf.Enabled = en
	return nil
}
