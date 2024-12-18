package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Morse code map for alphabet and numbers
var morseCode = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
	'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
	'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
	'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
	'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
	'Z': "--..", '0': "-----", '1': ".----", '2': "..---", '3': "...--",
	'4': "....-", '5': ".....", '6': "-....", '7': "--...", '8': "---..",
	'9': "----.",
	' ': " ", // Spaces remain as separators between words
}

// Function to convert text to Morse code
func textToMorse(input string) string {
	var morseResult []string
	for _, char := range strings.ToUpper(input) {
		if morse, exists := morseCode[char]; exists {
			morseResult = append(morseResult, morse)
		} else {
			morseResult = append(morseResult, "") //If the character is unknown, leave it blank.
		}
	}
	return strings.Join(morseResult, " ")
}

func main() {
	fmt.Println("===================================")
	fmt.Println("MorsecX Code Converter App V1")
	fmt.Println("Input Message and it will be converted to Morse Code")
	fmt.Println("===================================")

	// Reading input from terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the message you want to convert: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Convert input to Morse code
	morseOutput := textToMorse(input)

	// Displays Morse code output
	fmt.Println("\n=== Morse Code Conversion Results ===")
	fmt.Println(morseOutput)

	// Reading computer IP Address
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Failed to get IP Address:", err)
		return
	}

	titleIp := textToMorse("505 IP Address on this sender")
	fmt.Println("\n" + titleIp)

	// Combining IP in Morse code
	var ipMorseList []string
	for _, addr := range addrs {
		// Check if the address is of type *net.IPNet* (IP Network)
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() {
			// Only displays IPv4
			if ipNet.IP.To4() != nil {
				ipMorse := textToMorse(ipNet.IP.String())
				ipMorseList = append(ipMorseList, ipMorse)
				fmt.Println(ipMorse)
			}
		}
	}

	// Combine all Morse IPs into one string
	finalIpMorse := strings.Join(ipMorseList, "\n")

	// Save output to file
	fmt.Println("\nSave the results to file: output.txt")
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	// Write output to file
	_, err = file.WriteString(morseOutput + "\n\n" + titleIp + "\n" + finalIpMorse)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}

	fmt.Println("\nDone. The results have been saved to output.txt.")
}
