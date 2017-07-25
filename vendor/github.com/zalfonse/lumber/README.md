# Lumber
ANSI Logging module for Go. Each log type has it's own color!!

### Initialize a new logger like so: 

`logger = lumber.NewLogger(lumber.TRACE)`


### 6 Levels of logging:
- `lumber.TRACE`
- `lumber.DEBUG`
- `lumber.SUCCESS`
- `lumber.INFO`
- `lumber.WARNING`
- `lumber.ERROR`


### Log by calling:

`logger.Info("Here's the number: ", 1)`

`logger.Warning("Uh oh, this number doesn't look right: ", 3)`

etc..


### Color codes:

| Log Level | Color Code | Supposed to be |
|-----------|------------|----------------|
| TRACE     | 32m        | Green          |
| DEBUG     | 35m        | Purple         |
| SUCCESS   | 32m        | Green          |
| INFO      | 36m        | Cyan           |
| WARNING   | 31m        | Red            |
| ERROR     | 31m        | Red            |

### Looks like:
![](https://s.alfnz.com/pUOkI.png)
