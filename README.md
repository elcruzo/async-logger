# Asynchronous Logger in Go

![Go Version](https://img.shields.io/badge/Go-v1.17-blue)
![License](https://img.shields.io/badge/License-MIT-green)

A simple asynchronous logger implementation in Go for non-blocking logging to a file.

## Table of Contents

- [Asynchronous Logger in Go](#asynchronous-logger-in-go)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Usage](#usage)
  - [Examples](#examples)
  - [Contributing](#contributing)
  - [License](#license)

## Introduction

The Asynchronous Logger in Go is designed to provide a simple and efficient way to log messages asynchronously to a file, preventing blocking of your application during the logging process. It utilizes goroutines and channels to achieve non-blocking log writes.

## Features

- Asynchronous logging to a file.
- Non-blocking log writes.
- Simple and easy-to-use API.
- Automatic timestamping of log messages.
- Customizable log file location and format.

## Getting Started

### Installation

To use this Asynchronous Logger in your Go project, you can use Go modules. Import it as follows:

```go
import "github.com/elcruzo/async-logger"
```

### Usage

```go
// Create a new logger instance.
logger := asynclogger.NewLogger()

// Log messages asynchronously.
logger.Log("This is a log message.")
logger.Log("Another log message.")

// Stop the logger when it's no longer needed.
logger.Stop()
```

## Examples

For more usage examples and advanced features, check the [examples](examples/) directory in this repository.

## Contributing

Contributions are welcome! If you find a bug or have an enhancement in mind, please open an issue or submit a pull request. For more details, see [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.