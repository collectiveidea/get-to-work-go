# get-to-work
Start and stop project specific, annotated Harvest timers with information from Pivotal Tracker

## Workflow

### Initialize Your Project Directory
Navigate to your project directory and initialize `get-to-work`:
```shell
$ get-to-work init
```

Follow the directions on the screen.

### Start a Timer
Start a new timer with the following command:
```shell
$ get-to-work start [Pivotal Tracker URL]
```

Start your last timer with the following:
```shell
$ get-to-work start
```

### Stop a Timer
```shell
$ get-to-work stop
```

## Contributing
Install development dependencies:
```shell
$ brew install direnv
$ go get github.com/constabulary/gb/...
```

Setup your environment:
```shell
$ cp .apprc{.example,}
# edit your .apprc example with your Harvest and Pivotal Tracker credentials
$ direnv allow
```

Run the tests:
```shell
$ gb test
```
