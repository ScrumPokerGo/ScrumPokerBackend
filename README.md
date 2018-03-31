# ScrumPoker

BackendForScrumPokerRealtime

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/ScrumPoker
```

### Testing

``make test``

### Endpoints

POST /api/auth/

```
{
    "username":"sample_user",
    "password":"sample_pass"
}
```


GET /api/account  
`List account information`

PUT /api/account/gitlab/sync   
`Syncronize Projects,Milestones and Issues from gitlab`

GET /api/account/project   
`List all projects from the account`   

GET /api/account/project/{ID}   
`Get project info from account`   

GET /api/account/project/{ID}/milestone   
`Show milestones of a project`   

GET /api/account/project/{ID}/milestone/{ID}   
`Show milestone info of a project`   

GET /api/account/project/{ID}/milestone/{ID}/issue   
`Show list of issues from milestones`   

GET /api/account/project/{ID}/milestone/{ID}/issue/{ID} 
`Show info from the issue selected`  

PUT /api/account/project/{ID}/milestone/{ID}/issue/{ID}/estimate
`Update issue with the estimated time`  