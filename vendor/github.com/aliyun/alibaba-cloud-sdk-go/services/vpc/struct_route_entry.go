package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// RouteEntry is a nested struct in vpc response
type RouteEntry struct {
	PrivateIpAddress         string                        `json:"PrivateIpAddress" xml:"PrivateIpAddress"`
	NextHopOppsiteInstanceId string                        `json:"NextHopOppsiteInstanceId" xml:"NextHopOppsiteInstanceId"`
	NextHopType              string                        `json:"NextHopType" xml:"NextHopType"`
	IpVersion                string                        `json:"IpVersion" xml:"IpVersion"`
	RouteTableId             string                        `json:"RouteTableId" xml:"RouteTableId"`
	RouteEntryName           string                        `json:"RouteEntryName" xml:"RouteEntryName"`
	InstanceId               string                        `json:"InstanceId" xml:"InstanceId"`
	RouteEntryId             string                        `json:"RouteEntryId" xml:"RouteEntryId"`
	Status                   string                        `json:"Status" xml:"Status"`
	DestinationCidrBlock     string                        `json:"DestinationCidrBlock" xml:"DestinationCidrBlock"`
	NextHopRegionId          string                        `json:"NextHopRegionId" xml:"NextHopRegionId"`
	Description              string                        `json:"Description" xml:"Description"`
	NextHopOppsiteType       string                        `json:"NextHopOppsiteType" xml:"NextHopOppsiteType"`
	NextHopOppsiteRegionId   string                        `json:"NextHopOppsiteRegionId" xml:"NextHopOppsiteRegionId"`
	Type                     string                        `json:"Type" xml:"Type"`
	NextHops                 NextHopsInDescribeRouteTables `json:"NextHops" xml:"NextHops"`
}
