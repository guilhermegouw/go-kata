/*
The Challenge

Martha goes to a casino and brings n quarters.
The casino has three slot machines, and she plays them in order until she has
no quarters left.
That is, she plays the ﬁrst slot machine, then the second, then the third,
then back to the ﬁrst, then the second, and so on.
Each play costs one quarter.

The slot machines operate according to the following rules:
- The ﬁrst slot machine pays 30 quarters every 35th time it is played.
- The second slot machine pays 60 quarters every 100th time it is played.
- The third slot machine pays 9 quarters every 10th time it is played.
No other plays pay anything.

Determine the number of times Martha plays before she has no quarters left.

# Input

The input consists of four lines.
- The ﬁrst line contains an integer n, the number of quarters that Martha brings
to the casino. n is between 1 and 1,000.
- The second line contains an integer indicating the number of times that the ﬁrst
slot machine has been played since it last paid.
These plays occurred prior to Martha arriving, and Martha’s plays continue from there.
For example:

	Suppose that the ﬁrst slot machine has been played 34 times since it last paid.
	Then, Martha will win 30 quarters the ﬁrst time she plays it.

- The third line contains an integer indicating the number of times that the second slot
machine has been played since it last paid.
- The fourth line contains an integer indicating the number of times that the third slot
machine has been played since it last paid.

# Output

The number of times Martha plays before she has no quarters left
*/
package main

func CountPlays(quarters, firstMachine, secondMachine, thirdMachine int) int {
	plays := 0
	machine := 1
	for quarters > 0 {
		plays++
		quarters--

		switch machine {
		case 1:
			firstMachine++
			if firstMachine == 35 {
				quarters += 30
				firstMachine = 0
			}
			machine = 2
		case 2:
			secondMachine++
			if secondMachine == 100 {
				quarters += 60
				secondMachine = 0
			}
			machine = 3
		case 3:
			thirdMachine++
			if thirdMachine == 10 {
				quarters += 9
				thirdMachine = 0
			}
			machine = 1
		}
	}
	return plays
}
