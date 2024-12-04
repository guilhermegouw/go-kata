/*
The Challenge

Students would like to go on a school trip at the end of the year, but they need
money to pay for it.
To raise money, they have organized a brunch.
To attend the brunch, a student in
- their ﬁrst year pays $12
- their second year pays $10
- their third year pays $7
- their fourth year pays $5.

Of all of the money raised at the brunch, 50 percent of it can be used to pay for
the school trip (the other 50 percent is used to pay for the brunch itself).
We are told the cost of the school trip, the proportion of students in each year,
and the total number of students.
Determine whether the students need to raise more money for the school trip.

# Input

  - The ﬁrst line contains the cost in dollars of the school trip;
  - The second line contains four numbers indicating the proportion of brunching students
    who are in ﬁrst, second, third, and fourth year, respectively.
    Each number is between 0 and 1, and their sum is 1 (for 100 percent).
  - The third line contains integer n, the number of students attending the brunch.

# Output

For each test case: if the students need to raise more money for the school trip,
output YES; otherwise, output NO.
*/
package main

import (
	"fmt"
)

func NeedsMoreMoney(cost int, proportion [4]float64, students int) (string, error) {
	if cost < 0 || students < 0 {
		return "", fmt.Errorf("cost: %d or students: %d cannot be less than 0", cost, students)
	}
	fees := [4]int{12, 10, 7, 5}
	total := 0.0
	for i := range proportion {
		total += float64(students) * proportion[i] * float64(fees[i])
	}

	result := "YES"
	if total/2 >= float64(cost) {
		result = "NO"
	}
	return result, nil
}
