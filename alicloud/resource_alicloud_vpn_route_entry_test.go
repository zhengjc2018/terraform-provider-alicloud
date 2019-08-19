package alicloud

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func init() {
	resource.AddTestSweepers("alicloud_vpn_route_entry", &resource.Sweeper{
		Name: "alicloud_vpn_router_entry",
		F:    testSweepVPNRouterEntry,
		Dependencies: []string{
			"alicloud_vpn_gateway",
		},
	})
}

func testSweepVPNRouterEntry(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("error getting Alicloud client: %s", err)
	}
	client := rawClient.(*connectivity.AliyunClient)

	// prefixes := []string{
	// 	"tf-testAcc",
	// 	"tf_testAcc",
	// 	"tf_test_",
	// 	"tf-test-",
	// 	"testAcc",
	// }

	var rns []vpc.VpnRouteEntry
	req := vpc.CreateDescribeVpnRouteEntriesRequest()
	req.RegionId = client.RegionId
	req.PageSize = requests.NewInteger(PageSizeLarge)
	req.PageNumber = requests.NewInteger(1)

	for {
		raw, err := client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
			return vpcClient.DescribeVpnRouteEntries(req)
		})
		if err != nil {
			log.Printf("[ERROR] Error retrieving VPN Route Entry: %s", err)
		}
		resp, _ := raw.(*vpc.DescribeVpnRouteEntriesResponse)
		if resp == nil || len(resp.VpnRouteEntries.VpnRouteEntry) < 1 {
			break
		}
		rns = append(rns, resp.VpnRouteEntries.VpnRouteEntry...)

		if len(resp.VpnRouteEntries.VpnRouteEntry) < PageSizeLarge {
			break
		}

		if page, err := getNextpageNumber(req.PageNumber); err != nil {
			return err
		} else {
			req.PageNumber = page
		}
	}

	for _, v := range rns {
		req := vpc.CreateDeleteVpnRouteEntryRequest()
		req.VpnGatewayId = v.VpnInstanceId
		req.RouteDest = v.RouteDest
		req.NextHop = v.NextHop
		req.Weight = requests.NewInteger(v.Weight)

		_, err := client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
			return vpcClient.DeleteVpnRouteEntry(req)
		})
		if err != nil {
			log.Printf("[ERROR] Failed to delete VPN Route Entry (%s): %s", v.VpnInstanceId, err)
		}

		time.Sleep(10 * time.Second)
	}
	return nil
}

func testAccCheckVpnRouteEntryDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*connectivity.AliyunClient)
	vpnRouteEntryService := VpnRouteEntryService{client}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "alicloud_vpn_route_entry" {
			continue
		}

		fmt.Printf("rs: %v \t %[1]T\n\n\n", rs)
		_, err := vpnRouteEntryService.DescribeVpnRouteEntry(rs.Primary.ID, "test")

		if err != nil {
			if NotFoundError(err) {
				continue
			}
			return WrapError(err)
		}

		// if instance != "" {
		// 	return WrapError(Error("VPN %s still exist", instance.VpnGatewayId))
		// }
	}

	return nil
}

func TestAccAlicloudVpnRouteEntryBasic(t *testing.T) {
	var v vpc.DescribeVpnRouteEntriesResponse

	resourceId := "alicloud_vpn_route_entry.default"
	ra := resourceAttrInit(resourceId, testAccVpnGatewayCheckMap)
	serviceFunc := func() interface{} {
		return &VpnRouteEntryService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000, 9999)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithAccountSiteType(t, IntlSite)
		},

		// module name
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckVpnRouteEntryDestroy,
		Steps: []resource.TestStep{
			// {
			// 	Config: testAccVpnRouteEntryConfigBasic(rand),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheck(map[string]string{
			// 			"route_dest": fmt.Sprintf("tf-testAccVpnRouteEntryConfig%d", rand),
			// 		}),
			// 	),
			// },
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVpnRouteEntryConfigRouteDest(rand),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"route_dest": fmt.Sprintf("12.0.0.%d", rand),
					}),
				),
			},
			{
				Config: testAccVpnRouteEntryConfigWeight(rand),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"weight": fmt.Sprintf("%d", 100),
					}),
				),
			},

			{
				Config: testAccVpnRouteEntryConfigAll(rand),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"publish_vpc": fmt.Sprintf("%t", false),
						"description": fmt.Sprintf("tf-testAccVpnRouteEntryConfig%d", rand),
						"weight":      fmt.Sprintf("%d", 0),
					}),
				),
			},
		},
	})

}

func TestAccAlicloudVpnRouteEntryMulti(t *testing.T) {
	var v vpc.DescribeVpcAttributeResponse
	rand := acctest.RandInt()
	resourceId := "alicloud_vpc.default.9"
	ra := resourceAttrInit(resourceId, testAccCheckVpcCheckMap)
	serviceFunc := func() interface{} {
		return &VpcService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		// module name
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckVpnRouteEntryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnRouteEntryConfigAll(rand),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"publish_vpc": fmt.Sprintf("%t", true),
						"description": fmt.Sprintf("tf-testAccVpnRouteEntryConfig%d", rand),
						"weight":      fmt.Sprintf("%d", 100),
					}),
				),
			},
			{
				Config: testAccVpnRouteEntryConfigDescription(rand),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": fmt.Sprintf("tf-testAccVpnRouteEntryConfig%d", rand),
					}),
				),
			},
			{
				Config: testAccVpnRouteEntryConfigPublishVpc(rand),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"publish_vpc": fmt.Sprintf("%t", true),
					}),
				),
			},
			{
				Config: testAccVpnRouteEntryConfigWeight(rand),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"weight": fmt.Sprintf("%d", 100),
					}),
				),
			},
		},
	})
}

func testAccVpnRouteEntryConfigWeight(rand int) string {
	return fmt.Sprintf(`resource "alicloud_vpc" "default" {
  name       = "tf-testAccVpnRouteEntryConfig%d"
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vpn_gateway" "default" {
  name       = "tf-testAccVpnRouteEntryConfig%[1]d"
  vpc_id =  "${alicloud_vpc.default.id}"
  bandwidth = 5
  instance_charge_type = "PostPaid"
}

resource "alicloud_vpn_connection" "default" {
  name = "tf-testAccVpnRouteEntryConfig%[1]d"
  customer_gateway_id = "${alicloud_vpn_customer_gateway.default.id}"
  vpn_gateway_id = "${alicloud_vpn_gateway.default.id}"
  local_subnet = ["192.168.2.0/24"]
  remote_subnet = ["192.168.3.0/24"]
}

resource "alicloud_vpn_customer_gateway" "default" {
  name = "tf-testAccVpnRouteEntryConfig%[1]d"
  ip_address = "192.168.1.1"
}

resource "alicloud_vpn_route_entry" "default" {
  description       = ""
  vpn_gateway_id = "${alicloud_vpn_gateway.default.id}"
  route_dest = "12.0.0.2"
  next_hop = "${alicloud_vpn_connection.default.id}"
  weight =100
  publish_vpc = false
}`, rand)
}

func testAccVpnRouteEntryConfigDescription(rand int) string {
	return fmt.Sprintf(`
resource "alicloud_vpc" "default" {
  name       = "tf-testAccVpnRouteEntryConfig%d"
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vpn_gateway" "default" {
  name       = "tf-testAccVpnRouteEntryConfig%[1]d"
  vpc_id =  "${alicloud_vpc.default.id}"
  bandwidth = 5
  instance_charge_type = "PostPaid"
}

resource "alicloud_vpn_connection" "default" {
  name = "tf-testAccVpnRouteEntryConfig%[1]d"
  customer_gateway_id = "${alicloud_vpn_customer_gateway.default.id}"
  vpn_gateway_id = "${alicloud_vpn_gateway.default.id}"
  local_subnet = ["192.168.2.0/24"]
  remote_subnet = ["192.168.3.0/24"]
}

resource "alicloud_vpn_customer_gateway" "default" {
  name = "tf-testAccVpnRouteEntryConfig%[1]d"
  ip_address = "192.168.1.1"
}

resource "alicloud_vpn_route_entry" "default" {
  description       = "tf-testAccVpnRouteEntryConfig%[1]d"
  vpn_gateway_id = "${alicloud_vpn_gateway.default.id}"
  route_dest = "12.0.0.2"
  next_hop = "${alicloud_vpn_connection.default.id}"
  weight =100
  publish_vpc = false
}`, rand)
}

func testAccVpnRouteEntryConfigPublishVpc(rand int) string {
	return fmt.Sprintf(`
resource "alicloud_vpc" "default" {
  name       = "tf-testAccVpnRouteEntryConfig%d"
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vpn_gateway" "default" {
  name       = "tf-testAccVpnRouteEntryConfig%[1]d"
  vpc_id =  "${alicloud_vpc.default.id}"
  bandwidth = 5
  instance_charge_type = "PostPaid"
}

resource "alicloud_vpn_connection" "default" {
  name = "tf-testAccVpnRouteEntryConfig%[1]d"
  customer_gateway_id = "${alicloud_vpn_customer_gateway.default.id}"
  vpn_gateway_id = "${alicloud_vpn_gateway.default.id}"
  local_subnet = ["192.168.2.0/24"]
  remote_subnet = ["192.168.3.0/24"]
}

resource "alicloud_vpn_customer_gateway" "default" {
  name = "tf-testAccVpnRouteEntryConfig%[1]d"
  ip_address = "192.168.1.1"
}

resource "alicloud_vpn_route_entry" "default" {
  description       = ""
  vpn_gateway_id = "${alicloud_vpn_gateway.default.id}"
  route_dest = "12.0.0.2"
  next_hop = "${alicloud_vpn_connection.default.id}"
  weight =0
  publish_vpc = true
}`, rand)
}

func testAccVpnRouteEntryConfigAll(rand int) string {
	return fmt.Sprintf(`
resource "alicloud_vpc" "default" {
  name       = "tf-testAccVpnRouteEntryConfig%d"
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vpn_gateway" "default" {
  name       = "tf-testAccVpnRouteEntryConfig%[1]d"
  vpc_id =  "${alicloud_vpc.default.id}"
  bandwidth = 5
  instance_charge_type = "PostPaid"
}

resource "alicloud_vpn_connection" "default" {
  name = "tf-testAccVpnRouteEntryConfig%[1]d"
  customer_gateway_id = "${alicloud_vpn_customer_gateway.default.id}"
  vpn_gateway_id = "${alicloud_vpn_gateway.default.id}"
  local_subnet = ["192.168.2.0/24"]
  remote_subnet = ["192.168.3.0/24"]
}

resource "alicloud_vpn_customer_gateway" "default" {
  name = "tf-testAccVpnRouteEntryConfig%[1]d"
  ip_address = "192.168.1.1"
}

resource "alicloud_vpn_route_entry" "default" {
  description       = "tf-testAccVpnRouteEntryConfig%[1]d"
  vpn_gateway_id = "${alicloud_vpn_gateway.default.id}"
  route_dest = "12.0.0.2"
  next_hop = "${alicloud_vpn_connection.default.id}"
  weight =0
  publish_vpc = false
}`, rand)
}

var testAccVpnRoutrEntryCheckMap = map[string]string{
	"vpn_gateway_id": CHECKSET,
	"next_hop":       CHECKSET,
	"route_dest":     "12.0.0.2",
	"weight":         "0",
	"publish_vpc":    "false",
	"description":    "",
}
