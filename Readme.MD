# Morse Code Converter & IP Address Encoder (MorsecX)

A simple Golang application that converts text into Morse code and retrieves the computer's IP address, encoding it into Morse code as well. This application runs offline and is designed to output the results to both the terminal and a file (`output.txt`).

## Features

- **Text to Morse Conversion**: Convert user input text into Morse code.
- **IP Address to Morse Conversion**: Automatically detect the computer's IPv4 addresses and encode them into Morse code.
- **Offline Mode**: Runs completely offline without any external dependencies.
- **Output to File**: Saves the Morse code output, including the IP addresses, to `output.txt`.

## How It Works

1. The user enters a text input via the terminal.
2. The application converts the input into Morse code and displays it in the terminal.
3. The program retrieves the computer's IP addresses, encodes them in Morse code, and displays them in the terminal.
4. All outputs (text Morse code and IP Morse code) are saved in a file called `output.txt`.

### Example Output

**Input**:  
`Hello World`

## How to Run

1. Clone the repository:

```bash
   git clone <repository_url>
   cd <repository_name>
```

2. Run the program:

```bash
go run main.go
```

3. Build the executable (optional):

```bash
go build -o morse_converter.exe main.go
```

4. Execute the binary:

On Windows: `morse_converter.exe`

On Linux/Mac: `./morse_converter`

## Requirements

Golang: Make sure Go is installed. Download it from golang.org.
