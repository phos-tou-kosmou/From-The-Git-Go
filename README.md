![alt text][logo]

[logo]: https://cdn-images-1.medium.com/max/2000/0*V7hnAm-RUHa0-7HS.

# From-The-Git-Go

cli tool that will allow you to easily make new repos in github

first:
    `$go build main.go`

then to run:
    `$ ./main -u $GITHUB_USERNAME -p $NEW_GITHUB_REPO_NAME`

### Current Status as of July 12, 2019
    
    It seems that if you would like to validate a tool, such as, From-The-Git-Go,
    then you must authenticate with a bearer token utilizing oAuth 2.0; however, from the Linux
    CLI one would be able to authenticate with credentials.
    
    To be honest, From-The-Git-Go is just a super fancy alias and most developers probably 
    set this bash command as `newrepo` or just add a new remote to a non-existant repo, 
    which might be confusing for some; consequently, I had a lot of fun making this 
    and learned a lot about flag, os.Args, and http requests using Go.

    Refactoring thoughts:
    
        - Golang
            Although I am still in the process of completing this, if I were to rewrite
            it with Golang, I would use the flags library to try and recreate an actual 
            CLI argumentation array
            
        - Best Practice
            Bash...just..Bash
