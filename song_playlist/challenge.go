/*
The Challenge

We have ﬁve favorite songs named A, B, C, D, and E.
We’ve created a playlist of these songs and are using an app to manage the playlist.
The songs start oﬀ in the order A, B, C, D, E. The app has four buttons:
- Button 1: Moves the ﬁrst song of the playlist to the end of the playlist.
For example:
If the playlist is currently A, B, C, D, E, then it changes to B, C, D, E, A.
- Button 2: Moves the last song of the playlist to the beginning of the playlist.
For example:
If the playlist is currently A, B, C, D, E, then it changes to E, A, B, C, D.
- Button 3: Swaps the ﬁrst two songs of the playlist.
For example:
If the playlist is currently A, B, C, D, E, then it changes to be B, A, C, D, E.
- Button 4: Plays the playlist!

We’re provided a user’s button presses.
When the user presses button 4, output the order of songs in the playlist.

# Input

The input consists of pairs of lines:
- First line of a pair gives the number of a button (1, 2, 3, or 4)
- Second gives the number of times that the user pressed this button.
- Third line is the number of a button
- Fourth line is the number of times it is pressed, and so on.
- The input ends with these two lines:
4
1
indicating that the user pressed button 4 once.

# Output

Output the order of songs in the playlist after all button presses.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetPlaylist(presses string) (string, error) {
	playlist := []string{"A", "B", "C", "D", "E"}
	moves := strings.Split(presses, " ")

	for _, move := range moves {
		if len(move) != 2 {
			return "", fmt.Errorf("invalid input: a move cannot have lenght equal to %d", len(move))
		}

		button := string(move[0])
		if button != "1" && button != "2" && button != "3" && button != "4" {
			return "", fmt.Errorf("invalid button: %s", button)
		}

		timesPressed, err := strconv.Atoi(string(move[1]))
		if err != nil {
			return "", fmt.Errorf("error converting second item of move %s to integer: %v", move, err)
		}

		if button == "4" {
			return strings.Join(playlist, " "), nil
		}

		for range timesPressed {
			switch button {
			case "1":
				playlist = append(playlist[1:], playlist[0])
			case "2":
				playlist = append([]string{playlist[len(playlist)-1]}, playlist[:len(playlist)-1]...)
			case "3":
				playlist[0], playlist[1] = playlist[1], playlist[0]
			}
		}
	}
	return strings.Join(playlist, " "), nil
}
