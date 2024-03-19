/* The Gas Station problem is a classic algorithmic problem that can be solved using a greedy approach. Here's how the problem is typically stated:

You have a circular route with gas stations at different points. Each gas station has a certain amount of gas. You also have a car that can start at any gas station and must travel around the route, stopping at each gas station to refuel. The car consumes 1 unit of gas to travel from one station to the next. You need to find the starting gas station that will allow you to complete the entire route.

Here's the algorithm to solve this problem:

Initialize two variables: totalTank to keep track of the total gas available and currentTank to keep track of the gas in the current trip.

Iterate over each gas station:

a. Add the gas from the current station to totalTank.

b. Subtract the distance to the next station from the gas required for the trip (cost).

c. If currentTank + gas from the current station < cost, update startStation to the next station and reset currentTank to 0.

If totalTank is greater than or equal to 0 after visiting all stations, return startStation. Otherwise, return -1. */

package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	fmt.Printf("starting Gas Station:%v\n", canCompeleteCircuit(gas, cost))

}

func canCompeleteCircuit(gas []int, cost []int) int {
	totalTank, currentTank := 0, 0
	startStation := 0

	for i := 0; i < len(gas); i++ {
		totalTank += gas[i] - cost[i]
		currentTank += gas[i] - cost[i]

		if currentTank < 0 {
			startStation = i + 1
			currentTank = 0
		}
	}

	if totalTank >= 0 {
		return startStation
	}
	return -1

}
