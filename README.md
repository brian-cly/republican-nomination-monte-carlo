election2012.go is a Monte Carlo simulator of the 2012 presidential election written in Go. 
* It reads polling data from the Huffington Post API.
* It ignores Rasmussen polls (http://fivethirtyeight.blogs.nytimes.com/2010/11/04/rasmussen-polls-were-biased-and-inaccurate-quinnipiac-surveyusa-performed-strongly/)
* It transforms the polling data into probabilities.
* It assumes that where there's no polling data, the state will vote as it did in 2008.
* It runs 25,000 simulations to determine the most likely electoral college total for Obama.

## Usage ##

To run:
	$ go run election2012.go state.go api.go cumDist.go parse.go college.go

or

	$ go build election2012.go state.go api.go cumDist.go parse.go college.go
	$ ./election2012

For additional details about the data that was gathered and the simulation, see the the logfile.

	$ less logfile


## Notes ##

Why does the Monte Carlo simulation make the election appear more certain for Obama than the national polls that show a close race?

First, you need to understand that the national polling reported in the media is irrelevant. 47% Obama, 46% Romney? Irrelevant. Why? Because popular polls don't elect the president; the electoral college does. To predict the winner, you need to use the polls to simulate the various combinations of electoral college wins/losses for each candidate. This simulation uses 25,000 to simulated elections. Most of the time, Obama wins.

Second, you need to understand that if the polls are 52% Obama in a state with a margin of error of 4%, then Obama is 69% likely to win that state. You might think that 52% Obama in the polling means that he's 52% percent likely to win the state. But that's not how the statistics work. If you have a coin that comes out heads 52% of the time with a 4% margin of error, then that means that the probability of exceeding 50% is about 70%. To see this, go to a cumulative distribution function calculator such as http://www.danielsoper.com/statcalc3/calc.aspx?id=53 and plug in 52, 5, 50 and then calculate 1 - the answer given. In Obama's case, exceeding 50% is a win for that state. It turns out to have a probability of ~70%. 

Finally, realize that there is more certainty available in the data than is described in the polls. That's because 1) the electoral college elects presidents and 2) there are only so many ways that the electoral college can add up to a win for either candidate. For example, CA, HI, and NY will almost certainly vote for Obama; no need to poll in those states. That's why we talk about swing states; they can change the outcome. But there are only a limited number of ways the swing states can combine to a victory for one candidate or the other. As it turns out, the current polling makes these combinations favor Obama. It's nearly impossible for swing states to combine in ways that add up to a Romney win.