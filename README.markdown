# Random Text Generator written in Go

Install Go (http://golang.org) and compile the gotext.go file. If you want a pre-compiled executable that's working on an Intel-based OS X 10.5, grab gotext and run it.

The rest works almost exactly like in the original Random Text Generator in http://rtext.heroku.com.

Have fun but don't shoot me if the code seems a bit ugly. This is my first attempt at writing a Go application.

Here are some of the things you can do.

## Generate 4 random words

This generates 4 random words. 

    http://localhost:5678/words

Best usage scenario for this.

![Random password generator](http://imgs.xkcd.com/comics/password_strength.png)

## Generate 10 random words

This will generate 10 random words.

    http://localhost:5678/words/10
    

## Generate a random line from a classic book

This generates a single line of text from "Princess of Mars" by Edgar Rice Burroughs starting at a random location

    http://localhost:5678/book/barsoom

## Generate 10 lines from a classic book

This generates 10 lines of text from "The Adventure of Sherlock Holmes" by Arthur Conan Doyle starting at a random location

    http://localhost:5678/book/holmes/10
