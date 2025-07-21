# Todor

Query your todos in your code.


## Syntax 

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

## Supporpted Lanauges 
* Golang
* (WIP) Javascript & JSX
* (WIP) Typescript & TSX

# Building 

## Prerequisites
* Golang 1.23 (To be dockersized and removed as a requirement)
* docker
* docker compose
* make installed

Then you should be able to run `make build` to build, which should output to `dist/` folder. 
