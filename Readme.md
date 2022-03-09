# Elevator Navigation App

## Description

The project designed with **Model-View-Controller** (MVC) pattern. I've used [Fyne](https://developer.fyne.io/) (GUI) library as a view component.

## Running & Using Instructions

`make run`

After deciding which floor you want to go to with a very simple interface, just click on the relevant floor.

Meanwhile, you can watch the console output and make a another request to call the elevator to a different floor while the program is running.


## Test & Coverage

`make test`

## Help

`make help`

## Notes

Due to my limited time, I want to share the missing points.

- I've tried to do a multi-stage Docker build, however there were some errors due to Fyne components.
- Some parts of *ElevatorController* is still procedural. *OnPress and handle* methods can be divided to sub functions.
- From the above statement, after creating sub functions, **test** coverage can also be increased.