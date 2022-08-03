# proc

Process control, logging and child processes

## Description

Golang is a great language for concurrency, but sometimes you want parallelism, and it's often simpler to design apps
as little units that plug together via RPC APIs. Services don't always have easy start/stop/reconfig/restart
controls, and sometimes they use their own custom configuration scheme that is complicated to integrate with yours.

In addition, you may want to design your applications as neat little pieces, but how to attach them together and
orchestrate them starting up and coming down cleanly, and not have to deal with several different ttys to get the logs.

Proc creates a simple framework for creating observable, controllable execution units out of your code, and more
easily integrating the code of others.

This project is the merge of several libraries for logging, spawning and controlling child processes, and creating
an RPC to child processes that controls the run of the child process. Due to the confusing duplication of signals as
a means of control and the lack of uniformity in signals (ie windows) the goal of `proc` is to create one way to do,
following the principles of design used for Go itself.

## Badges

<insert badges here

## Installation

### For developers:

To make nice tagged versions with the version number in the commit as well as matching to the tag, there is a tool
called [bumper](cmd/bumper) that bumps the version to the next patch version (vX.X.PATCH), embeds this new version
number into [version.go](./version.go) with the matching corresponding git commit hash. This will make importing
this library at an exact commit much more human.

In addition, it makes it easy to make a nice multi line commit message as many repositories request in their 
CONTRIBUTION file by replacing ` -- ` with two carriage returns.

To install:

    go install ./bumper/.

To use:

    bumper make a commit comment here -- text after double \
        hyphen will be separated by a carriage return -- \
        anywhere in the text

To automatically bump the minor version use `minor` as the first word of the comment. For the major version `major`.

## Usage

reasonably include in the README.

## Support

## Roadmap

- [x] Implement basic logger v0.0.x
- [ ] Create concurrent safe configuration CLI/env/config system v0.1.x
  - [x] Created types with mutex/atomic locks to prevent concurrent access
  - [x] Created key value map type containing collections of concurrent safe values with concurrent safe access to 
    members
  - [x] Create JSON marshal/unmarshal for configuration collections
  - [x] Created tests for generating and concurrently accessing/mutating data
  - [x] Save and load configuration from file (using json, with all values stored in file, 
    no unnecessary comments)
  - [ ] Generate CLI help texts from configs specifications
  - [ ] Read values from environment variables overlay on config file values
  - [ ] Created command line parsing system overlay values above previous
- [ ] Child process control system v0.2.x
  - [ ] Launch, pause, continue and stop child process. Use only one method: the IPC API, no complication with signals.
  - [ ] Read and pass through logging from child process
  - [ ] Correctly handle process signals from OS/tty to trigger orderly shutdown of child processes and leave none 
    orphaned

## Contributing

## Authors and acknowledgment

David Vennik david@cybriq.systems

## License

Unlicenced: see [here](./LICENSE)

## Project status

In the process of revision and merging together several related libraries that need to be unified.