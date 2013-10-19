creeper
=======

Super light weight IF file_changed THEN do_this cli app written in Go. A lot of environments come with their own mechanisms for this, such as [Compass](http://thesassway.com/beginner/getting-started-with-sass-and-compass) or [CoffeeScript](http://coffeescript.org/#usage). But sometimes you want to have the some behaviour on any type of file, with any type of command being triggered. Maybe you like to stage all the things, all the time. Perhaps you want to run an external linting process against a file but the built-in system is a little lacking. 

This is where Creeper comes in. Creeper is a ludicrously simple application that will keep an eye on the file you specify, checking at specific intervals to see if the file has changed (It checks the modification time on the file). If the file has been modified then the command that was provided is run and the output is printed to the terminal. Simple!

Creeper is single target at the moment, there is no support for creeping on an entire directory or list of files, but I'm working on it. Additionally the command that is given is execute in a blocking manner, which may cause issues or a 'backing up' of requests if the command is particularly long running. Plans are afoot that will offer some alternatives to this. Fiddling with the wait period can aleviate this shoudl the need arise.

ON TO THE USAGE!

## Basic Usage
```
Usage of ./creeper:
  -cmd="": This is the cmd to execute
  -file="": This is the file to watch!
  -q=false: Intent message won't be displayed
  -shuttup=false: Creeper will make a best error to not make any std/out/err noise at all.
  -wait=1s: An interval duration, defaults to '1s' for one second.
```

To use simply point it at a file and give it something to do when the file changes:

```
creeper -file foo.go -cmd go build bar.go
```

This will attempt a build of bar.go whenever <code>foo.go</code> is updated, checking once a second.

### Alternate usage

Maybe you type super fast, or you want something to happen if an error.log is updated somewhere. You can mess with the <code>-wait</code> option to tell creeper how often you want to creep on the file in question.

```
creeper -file bad-errors.log -cmd notify_someone.sh -wait 1h

...or

creeper -file thesis.tex -wait 200ms -cmd pdflatex thesis.tex
```

Any trailing arguments will be packaged as parameters for the given command when it runs, thus...

```
creeper -file gargleblaster.code -cmd builder arg1 arg2 arg3 arg4
```

This will run the <code>builder</code> with all of the trailing arguments: <code>arg1 arg2 arg3 arg4</code>. 