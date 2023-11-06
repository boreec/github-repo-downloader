# github-repo-downloader

This tool allows you to download all the public repositories of GitHub users and
organizations. It doesn't require any access tokens to fetch repositories;
instead, it uses web scraping to retrieve data. 

Keep in mind that if there are any changes to the website's HTML structure, it
might affect the tool's functionality.

## usage

Build the program:

```sh
go build -v  
```

Specify the repositories to fetch for users or organization using the formats
`org:organization-name` or `user:user-name`. You can specify more than one org
or user:

```sh
./repo-downloader user:boreec org:golang user:octocat org:rust-lang
```
