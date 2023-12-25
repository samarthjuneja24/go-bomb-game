# go-bomb-game


<h3>
What does this do?
</h3>
Passing a bomb between two players with a timeout set in context.The player having the bomb at the end of the timer expiry loses.

We're passing the sender and receiver in the context and pass it through the channels. The receiver at the end of context timeout is the loser.
<h3>
How to run the code?
</h3>
<code>
go run main.go
</code>

<h3>
Why did I write this code?
</h3>
To practice the concept of channels in Go

<h3>Idea inspiration
</h3>
https://www.reddit.com/r/golang/comments/187joet/bomb_game_goroutines_and_channels/