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

To bump patch version and make a commit, install the program in [bumper/](./bumper), run with the commit string 
after the command. To install:

    go install ./bumper/.

To use:

    bumper make a commit comment here -- text after double \
        hyphen will be separated by a carriage return -- \
        anywhere in the text

## Usage

reasonably include in the README.

## Support

## Roadmap

- [x] Implement basic logger
- [ ] Create concurrent safe configuration CLI/env/config system

## Contributing

## Authors and acknowledgment

David Vennik david@cybriq.systems

## License

Unlicenced: see [here](./LICENSE)

## Project status

In the process of revision and merging together several related libraries that need to be unified.