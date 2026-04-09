package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// command-line flags
	tokenString := flag.String("token", "", "JWT token to decode")
	tFlag := flag.String("t", "", "Shorthand for -token")
	secretString := flag.String("secret", "", "Optional secret key for signature verification")
	sFlag := flag.String("s", "", "Shorthand for -secret")
	helpFlag := flag.Bool("help", false, "Show usage")
	hFlag := flag.Bool("h", false, "Shorthand for -help")
	flag.Parse()

	if *helpFlag || *hFlag {
		fmt.Println("Usage: jwt-decoder -token <JWT_TOKEN> [-secret <SECRET_KEY>] [-h|-help]")
		fmt.Println("  -t, -token  JWT token to decode (required)")
		fmt.Println("  -s, -secret  Secret key for signature verification (optional, HMAC only)")
		fmt.Println("  -h, -help  Show this help message")
		os.Exit(0)
	}

	// allow -t as shorthand for -token
	if *tokenString == "" && *tFlag != "" {
		*tokenString = *tFlag
	}
	if *tokenString == "" {
		fmt.Println("Error: Please provide a JWT token using the -token or -t flag.")
		fmt.Println("Usage: jwt-decoder -token <JWT_TOKEN> [-secret <SECRET_KEY>] [-h|-help]")
		os.Exit(1)
	}

	// allow -s as shorthand for -secret
	if *secretString == "" && *sFlag != "" {
		*secretString = *sFlag
	}

	// Split token
	parts := strings.Split(*tokenString, ".")
	if len(parts) != 3 {
		fmt.Println("Error: Invalid token format.")
		os.Exit(1)
	}

	// Decode header
	header, err := decodeJWT(parts[0])
	if err != nil {
		fmt.Printf("Error decoding header: %v\n", err)
		os.Exit(1)
	}
	printJSON("Header", header)

	//Decode payload
	payload, err := decodeJWT(parts[1])
	if err != nil {
		fmt.Printf("Error decoding payload: %v\n", err)
		os.Exit(1)
	}
	printJSON("Payload", payload)

	// Verify the signature if a secret is provided
	if *secretString != "" {
		token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (any, error) {
			// right now only HMAC is supported
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method, %v", token.Header["alg"])
			}
			return []byte(*secretString), nil
		})

		if err != nil {
			fmt.Printf("Signature verification failed: %v\n", err)
			os.Exit(1)
		}

		if token.Valid {
			fmt.Println("Signature checks out")
		} else {
			fmt.Println("Invalid signature")
		}
	} else {
		fmt.Println("Signature validation skipped (no secret provided)")
	}
}

// base64 decode to get JSON
func decodeJWT(segment string) (map[string]any, error) {
	decoded, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil {
		return nil, err
	}
	var result map[string]any
	if err := json.Unmarshal(decoded, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// print decoded JWT stuff
func printJSON(label string, data map[string]any) {
	fmt.Printf("%s:\n", label)
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting %s: %v\n", label, err)
		return
	}
	fmt.Println(string(output))
}
