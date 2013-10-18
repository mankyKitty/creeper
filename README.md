creeper
=======

Super light weight IF file_changed THEN do_this cli app written in Go. Creeper is single target at the moment, there is no support for creeping on an entire directory or list of files, but I'm working on it.

## Basic Usage
```
Usage of ./creeper:
  -cmd="": This is the cmd to execute
  -file="": This is the file to watch!
  -q=false: Intent message won't be displayed (optional)
  -shuttup=false: Creeper will make a best error to not make any std/out/err noise at all. (optional)
  -wait=1000: Time in Milliseconds between creeping on the file (optional - defaults to 1 second)
```

To use simply point it at a file and give it something to do when the file changes:

```
creeper -file foo.go -cmd go build bar.go
```

This will attempt a build of foo.go whenever <code>foo.go</code> is updated, checking once a second.

### Alternate usage

Maybe you type super fast, or you want something to happen if an error.log is updated somewhere. You can mess with the <code>-wait</code> option to tell creeper how often you want to creep on the file in question.

```
creeper -file bad-errors.log -cmd notify_someone.sh -wait 3600000

...or

creeper -file thesis.tex -wait 200 -cmd pdflatex thesis.tex
```

Any trailing arguments will be packaged as parameters for the given command when it runs.