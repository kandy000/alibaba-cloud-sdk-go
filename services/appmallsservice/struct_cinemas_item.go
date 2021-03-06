package appmallsservice

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

// CinemasItem is a nested struct in appmallsservice response
type CinemasItem struct {
	Address           string `json:"Address" xml:"Address"`
	CinemaName        string `json:"CinemaName" xml:"CinemaName"`
	CityId            int64  `json:"CityId" xml:"CityId"`
	Id                int64  `json:"Id" xml:"Id"`
	Latitude          string `json:"Latitude" xml:"Latitude"`
	Longitude         string `json:"Longitude" xml:"Longitude"`
	Phone             string `json:"Phone" xml:"Phone"`
	RegionName        string `json:"RegionName" xml:"RegionName"`
	ScheduleCloseTime int64  `json:"ScheduleCloseTime" xml:"ScheduleCloseTime"`
	StandardId        string `json:"StandardId" xml:"StandardId"`
}
