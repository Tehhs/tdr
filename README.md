# Todor

Query your TODOs in your code.

Install using `go install github.com/Tehhs/tdr@latest`



# Features

* Query your TODOs in your code
* Provides a small "TODO" syntax/language for use in your code.
* Output in JSON for connections to other tools and workflows (WIP) 


# Syntax 

Todor supports the typical todo syntax: 
```
//todo: this needs to be done
```

Optionally, todor allows you to add on top of this.

### Syntax: Tags 

You can add tags to your todos like so:

```
//todo(important, security): this needs to be done asap
```

And then query your tags by doing something like `tdr -t security` or `tdr -t important`.

# Supporpted Lanauges 
* Golang
* Javascript (including JSX)
* (WIP) Typescript & TSX
* (WIP) C#

# Building 

### Prerequisites
* Golang 1.23 (To be dockersized and removed as a requirement)
* docker
* docker compose
* make installed

Then you should be able to run `make build` to build, which should output to `dist/` folder. 
