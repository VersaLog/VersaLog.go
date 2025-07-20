# What is VersaLog.go?

What is VersaLog.go?
VersaLog is a powerful and flexible logging library for Golang.
It supports everything from simple usage to advanced, highly customizable configurations to meet a wide range of needs.

## Installation

```
go get github.com/kayu0514/VersaLog.go
```

### Mode

| Mode         |Description                                   |
| ------------ | -------------------------------------------- |
| `detailed`   | Logs including execution time and log levels |
| `file`       | Logs with filename and line number           |
| `simple`     | Simple and easy-to-read logs                 |

### Options

| Options      |Description                                   |
| ------------ | --------------------------------------------------------------          |
| `show_file`  | True : Display filename and line number (for simple and detailed modes) |
| `show_tag`   | True : Show self.tag if no explicit tag is provided                     |
| `tag`        | Default tag to use when show_tag is enabled                             |
| `all`        | Shortcut to enable both show_file and show_tag                          |

## Sample

**Simple** : [Tap](/tests/simple_test.go)  
**Detailed** : [Tap](/tests/detailed_test.go)
**File** : [Tap](/tests/file_test.go)
