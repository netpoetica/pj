# pj
## Convenience utility to easily navigate to projects (subfolders) in your primary project directory.

### What does it do?
After installing this minimal command line tool, you can simply call

```
pj my-project-name
```

instead of having to navigate all the way to it via terminal:

```
cd /Users/username/Desktop/WebProjects/my-project-name
```

### Setup
The pj executable will need to be installed into somewhere in your path, like your /usr/bin or /usr/local/bin directory, and it depends on two modifications to your .bashrc or .zshrc

#### Note: this was tested on Mac OSX Mountain Lion (UNIX) with zsh

First, move the pj executable to a location in your environment $PATH var:

```
sudo cp pj /usr/bin
```

Next, add the PJ_DIRS environment variable to your rc file:

#### Note: do not leave a trailing slash at the end of the directory.
```
export PJ_DIRS=/Users/netpoetica/Desktop/Projects
```

and then add a simple bash function right below it:

```
function pj() {
  cd "`/usr/bin/pj --project $1`"
}
```

finally, source your shell rc file:

```
source ~/.zshrc
```

or whatever rc file for whatever shell you are using.

##### Umm... why are you using bash?
Because in order to use the CD command, I am running a go program which outputs the precise directory you want to navigate to to std out. There is no way for go to finish its process AND THEN do something in UNIX terminal :-) (I learned this the hard way, even though it should have been a no-brainer).

#### Note: here, you could change /usr/bin/ to /usr/local/bin, etc, if you want to install pj there instead.

### Future Updates
Future updates will add even more convenience:
- Create aliases for your projects that pj will remember
- Allow multiple project directories
- Investigate autocompletion of project names
- Allow you to specify where pj is installed (instead of /usr/bin)

This was a fun little project written in the Go programming language, which is a really cool language from what I've seen so far. Looking forward to making more!

### Working with Go
Go is a ton of fun! Here are some resources to help get you started:
- Installing: http://golang.org/doc/install
- Clearly laid out extra install steps if anything goes wrong: http://stackoverflow.com/questions/10130341/go-go-get-go-install-local-packages-and-version-control/10142340#10142340
- All of the Go folks who would love to help you (including me): https://groups.google.com/forum/#!forum/golang-nuts
- A book, and free! at https://gobyexample.com/
- THE "book", also free: http://golang.org/doc/effective_go.html