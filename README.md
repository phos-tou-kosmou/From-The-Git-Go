# from-the-Git-Go

cli-tool that will allow you to easily make new repos in github

first:
    `$go build main.go`

then to run:
    `$ ./main -u $GITHUB_USERNAME -p $NEW_GITHUB_REPO_NAME`

###Current Status as of July 12, 2019
    It seems that if you would like to validate a tool like from-the-Git-Go,
    then you must authenticate with a bearer token; however, from the Linux
    CLI one would be able to authenticate with credentials after executing the curl.

    I was hoping this was the case, but Golang requires a method to issue
    an http protocol and for good reason.  Therefore, I could most likely dump the 
    data into a bash script and easily fire the command.  To be honest, this is just
    a super fancy alias and most developer probably set this as newrepo or something
    of the like; however, I had a lot of fun making this.

    Refactoring thoughts:
        - Golang
            Although I am still in the process of completing this, if I were to rewrite
            it with the flags library to try and recreate and actual CLI argumentation array
        - Best Practice
            Bash...