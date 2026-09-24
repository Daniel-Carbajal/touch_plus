<h1>touch_plus</h1>
A lightweight command-line file utility built in Go using only the standard library. It accepts a file path, creates and initializes the file if it does not exist, and provides an interactive loop for reading its contents, overwriting the file with user input, and quitting.

Built as my first Go project for a distributed systems course, with a focus on file I/O, command-line input, error handling, and resource cleanup.

## CSC 376 Programming Assignment 1
**Jarvis College of Computing and Digital Media - DePaul University**

**Student**: Daniel Carbajal (dcarbaj4@depaul.edu)  
**Solution programming language**: Go 

<h2>Compile and run instructions:</h2>
**Perform these steps within project directory**<br>
<h3>1. Run 'go build' in project directory</h3>

```bash
go build
```

<p>You can now run the executable using `./touch_plus [filePath]`. However to run the executable without needing to specify or be withing the project directory continue with the following steps.</p>
<h3>2. Install the executable:</h3>

```bash
go install .
```
<h3>3. Add the Go install directory to your system's shell path</h3>
This enables you to run the programs executable without specifying where the executable is.

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

<h2>Using touch_plus</h2>
Run the executable with:

```bash
touch_plus <filePath>
```

Where filePath leads to a pre-existing file or a new one you want to create.
<br>
Within the program you can run the following commands to perform various actions:
<ul>
    <li>[read] - display the contents of the file</li>
    <li>[write] [args] - overwrite the current contents of the file with the entered arguments</li>
    <li>[quit] - close the program</li>
</ul>
    
