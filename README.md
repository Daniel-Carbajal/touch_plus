<h1>touch_plus</h1>
A lightweight command-line file utility built in Go using only the standard library. It accepts a file path, creates and initializes the file if it does not exist, and provides an interactive loop for reading its contents, appending text, and quitting.

Built as my first Go project for a distributed systems course, with a focus on file I/O, command-line input, error handling, and resource cleanup.

## CSC 376 Programming Assignment 1
**Jarvis College of Computing and Digital Media - DePaul University**

**Student**: Daniel Carbajal (dcarbaj4@depaul.edu)  
**Solution programming language**: Go 

<h2>Compile and run instructions:</h2>
Perform these steps within project directory
<ol>
    <li>Run 'go build' in project directory</li>
        '''bash
        go build
        '''
        You can now run the executable using "./touch_plus [filePath]". However to run the executable without needing to specify or be withing the project directory continue with the following steps.
    <li>Install the executable</li>
        '''bash
        go install .
        '''
    <li>Add the Go install directory to your system's shell path </li>
        This enables you to run the programs executable without specifying where the executable is.
        '''bash
        export PATH="$PATH:$(go env GOPATH)/bin"
        ''' 
</ol>

<h2>Using touch_plus</h2>
Run the executable with:
'''bash
touch_plus <filePath>
'''

Where filePath leads to a pre-existing file or a new one you want to create.
<br>
Within the program you can run the following commands to perform various actions:
<ul>
    <li>[read] - display the contents of the file</li>
    <li>[write] [args] - append the entered arguments to the file</li>
    <li>[quit] - close the program</li>
</ul>
    
