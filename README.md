# go-quiz - Sebastian Miro

Simple quiz app, with an API and CLI tool made in Go.

Comes with two quizes by default, one about Go the programing language and
another about the city of Barcelona. It's easy to add more.
Each quiz has its own Leaderboard.

## Usage

The repository comes with a simple Makefile wich builds two binaries to a build
directory.

Once you cloned the reposiory:

``` make ```

Then you need to run the api server, it runs on the port 3000:

``` ./build/quiz-api ```

Now you can use the cli wich connects automaticaly to the server:

``` ./build/quiz-cli --help ```

