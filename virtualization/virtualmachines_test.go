package virtualization

import (
	"github.com/sapcc/go-netbox-go/dcim"
	"github.com/sapcc/go-netbox-go/ipam"
	"github.com/sapcc/go-netbox-go/models"
	"github.com/sapcc/go-netbox-go/tenancy"
	"github.com/seborama/govcr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_ListVirtualMachines(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("ListVirtualMachines", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListVirtualMachinesRequest{}
	opts.Id = 1060
	res, err := client.ListVirtualMachines(opts)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(res)
	t.Log(res.Results[0].Name)
	t.Log(res.Results[0].Cluster.Name)
	t.Log(res.Results[0].Status)
	assert.NotEqual(t, 0, res.Count)
}

func TestClient_CreateVirtualMachine(t *testing.T) {
	client, err := New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	ipamClient, err := ipam.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	tenantClient, err := tenancy.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	dcimClient, err := dcim.New(os.Getenv("NETBOX_URL"), os.Getenv("NETBOX_TOKEN"), true)
	if err != nil {
		t.Fatal(err)
	}
	vcrConf := &govcr.VCRConfig{}
	vcrConf.Client = client.HttpClient
	vcr := govcr.NewVCR("CreateVirtualMachine", vcrConf)
	client.HttpClient = vcr.Client
	opts := models.ListClusterRequest{}
	res, err := client.ListClusters(opts)
	if err != nil {
		t.Fatal(err)
	}
	rOpts := models.ListRolesRequest{}
	roles, err := ipamClient.ListRoles(rOpts)
	if err != nil {
		t.Fatal(err)
	}
	tOpts := models.ListTenantsRequest{}
	tenants, err := tenantClient.ListTenants(tOpts)
	if err != nil {
		t.Fatal(err)
	}
	pOpts := models.ListPlatformsRequest{}
	platforms, err := dcimClient.ListPlatforms(pOpts)
	if err != nil {
		t.Fatal(err)
	}
	vm := models.WriteableVirtualMachine{
		Name: "test-d062260.cc.qa-de-1.cloud.sap",
		Cluster: res.Results[0].Id,
		Status: "active",
		Role: roles.Results[0].Id,
		Tenant: tenants.Results[0].Id,
		Platform: platforms.Results[0].Id,
	}
	err = client.CreateVirtualMachine(vm)
	if err != nil {
		t.Fatal(err)
	}
}
