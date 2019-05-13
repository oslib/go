# The Go++ Programming Language

Go++ is a derivative of the popular open source programming language that makes it easy to build simple,
reliable, and efficient software, named go, and sometimes referred to as golang. The basic idea is to enhance the core language with support for classes (a kind of hybrid between structs and interfaces), and to improve the usability and readability by adding new keywords and simplifying some of the heavier syntax. 

![Go++ image](doc/images/Go++.jpg)
*Go++ image by [Glenn Fiedler] [https://gafferongames.com/post/shape_of_the_go_stone].*

Unless otherwise noted, the Go++ source files are distributed under the BSD-style license found in the LICENSE file.

### Download and Install

To give Go++ a spin, you will need to download the source files from github and build a new version of the compilers and related tools. First, clone a copy of the github repository: 
```
> git clone https://github.com/oslib/gopp.git go++
```
This step is exactly the same on the three supported platforms, Linux, Mac, and Windows (basically, the only ones that have been tested). The "go++" directory, specified by the last parameter, can be any local name or, if omitted, it will default to the git supported name "gopp" (git does not allow the "++" special characters when naming repositories). 

The next step is a short form of the golang bootstrap process. It starts with a binary installation of the standard golang environment, and must match the exact release version that was used as the initial branch fot G++ (currently 1.12.5). These files may be retrieved from the official binary distributions available at https://golang.org/dl/. 

If a more current release has been made it will show at the top of the page, in which case you'll need an archived version which can be accessed via the link at the bottom of the main page. After downloading a binary release, visit https://golang.org/doc/install for installation instructions. Once installed:
```
make a copy of the "bin" and "pkg" subdirectories from the 1.12.5 binary distribution 
pasting them into your Go++ source directory, "go++" in this example. 
```
Of course, this step is a little different on each platform, but is most easily performed with the local file explorer. The GOROOT environment variable should be set to the go++ directory and the GOPATH should be unset or set to some other path (it's considered an error to have these pointing to the same location).  
```
> export GOROOT=~/go++  
> export GOPATH= 
```
The "export" command should work for Linux and Mac. On windows the "set" command would be used to accomplish the same result. On each platform these variable may be setup as permanent for each new session with GOROOT pointing to the go++ directory, and GOPATH set to whatever directory is being used for local development. The Go++ compiler should be fully backwards compatible and be able to compile any standard golang packages (with the exception of any name conflicts with new keywords added for the Go++ language). 

The last step is to build the Go++ compiler and related tools: 
```
> cd ~/go++/src   
> ./BuildGo++.bash  
```
The build command should be executed from the "src" subdirectory withing the "go++" main directory. On windows the ".bat" command file should be used: 
```
C:> cd \go++\src   
C:\go++\src> BuildGo++.bat  
```
Hopefully, all of that worked without incident and you're ready to give the test program a try: 
```
> cd ~/go++/src/hello   
> go build
> ./hello 
H
He
Hel
Hell
Hello
Hello World!

Welcome to Go++ - The friendlier golang...
>
```
The program is in two parts, "src/hello" for the main program, and "src/hellosub" to demonstrate the inclusion of a seperate package. Of course, the Go++ language is strictly experimental and still a work in progress, it should only be used at your own risk or perhaps, peril (there are a fews details still being worked out).  

