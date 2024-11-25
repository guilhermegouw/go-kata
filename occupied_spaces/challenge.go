/*
The Challenge

You supervise a parking lot with n parking spaces.
Yesterday, you recorded whether each parking space was occupied by a car or was empty.
Today, you again recorded whether each parking space was occupied by a car or was empty.
Indicate the number of parking spaces that were occupied on both days.

# Input

The input consists of three lines.
- The first line contains integer n, the number of parking spaces.
- The second line contains a string of n characters for yesterday’s
information,  one character for each parking space. A C indicates an occupied
parking space (C for car), and a .  indicates an empty parking space.
For example, CC. means that the ﬁrst two parking spaces were occupied and the
third was empty.
- The third line contains a string of n characters for today’s information, in the
same format as the second line.

# Output

Output the number of parking spaces that were occupied on both days.
*/
package challenge

import "fmt"

func OccupiedOnBothDays(numSpaces int, yesterday, today string) (int, error) {
	if numSpaces != len(yesterday) || numSpaces != len(today) {
		return 0, fmt.Errorf("invalid input: numSpaces %d does not match string lengths %d, %d", numSpaces, len(yesterday), len(today))
	}
	count := 0
	for i := 0; i < numSpaces; i++ {
		if yesterday[i] == 'C' && today[i] == 'C' {
			count++
		}
	}
	return count, nil
}
