// O(n), O(d), d is height, n is number of nodes

// use getInfo, counter and lowest common
// if meet one, counter++ , lowest node with counter = 2 means found
// if lowestCommon found, return prev orgInfo until root
package main

import (
	"fmt"
)

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

type OrgInfo struct {
	counter             int
	lowestCommonManager *OrgChart
}

func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	// Write your code here.
	return getInfo(org, reportOne, reportTwo).lowestCommonManager
}

// OrgInfo will change for every call, creating new OrgInfo every recursion
// to find lowest, traverse first, from the bottom start checking
func getInfo(manager, one, two *OrgChart) OrgInfo {
	counter := 0

	// start from bottom, traverse first
	for _, node := range manager.DirectReports {
		orgInfo := getInfo(node, one, two)
		fmt.Print(orgInfo)
		fmt.Print("\n")
		// if alr found lowestCommon
		if orgInfo.lowestCommonManager != nil {
			return orgInfo
		}
		// keep track of prev counter
		counter += orgInfo.counter
	}

	// check if found
	if manager == one || manager == two {
		counter += 1
	}
	var lowestCommon *OrgChart
	if counter == 2 {
		lowestCommon = manager
	}
	// return new OrgInfo
	return OrgInfo{
		counter,
		lowestCommon,
	}
}

// orgCharts := getOrgCharts()
// 	orgCharts['A'].addDirectReports(orgCharts['B'], orgCharts['C'])
// 	orgCharts['B'].addDirectReports(orgCharts['D'], orgCharts['E'])
// 	orgCharts['C'].addDirectReports(orgCharts['F'], orgCharts['G'])
// 	orgCharts['D'].addDirectReports(orgCharts['H'], orgCharts['I'])
// 	lcm := GetLowestCommonManager(orgCharts['A'], orgCharts['E'], orgCharts['I'])
// 	require.Equal(t, lcm, orgCharts['B'], lcm)
