package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/12yanogden/clipboard"
	"github.com/12yanogden/ints"
	"github.com/spf13/cobra"
)

// Base command
var rootCmd = &cobra.Command{
	Use:   "goint",
	Short: "Return compatible types",
	Long: `Return a list of compatible go types, given the input.
	
	For example: goint 256`,

	Run: getType,
}

func init() {
	// Set flags, if applicable
}

func getType(cmd *cobra.Command, args []string) {
	max := -1

	if len(args) == 0 {
		fmt.Println("goint: must pass a maximum value")
		os.Exit(1)
	}

	if ints.IsNum(args[0]) {
		max, _ = strconv.Atoi(args[0])
	} else {
		fmt.Printf("goint: maximum value must be a number. Found: %v\n", max)
		os.Exit(1)
	}

	unsignedType := calcUnsignedType(max)
	signedType := calcSignedType(max)

	fmt.Printf("Signed: %s\n", signedType)
	fmt.Printf("Unsigned: %s\n", unsignedType)

	fmt.Println()

	clipboard.Push(unsignedType)
	fmt.Printf("Copied '%s' to your clipboard.\n\n", unsignedType)
}

func calcUnsignedType(max int) string {
	switch {
	case max < 0:
		break
	case max <= 255:
		return "uint8"
	case max <= 65535:
		return "uint16"
	case max <= 4294967295:
		return "uint32"
	// uint64 can reach 18446744073709551615, but the 'max' signed int can't
	case max <= 9223372036854775807:
		return "uint64"
	}

	return "Out of range"
}

func calcSignedType(max int) string {
	switch {
	case max >= -128 && max <= 127:
		return "int8"
	case max >= -32768 && max <= 32767:
		return "int16"
	case max >= -2147483648 && max <= 2147483647:
		return "int32"
	case max >= -9223372036854775808 && max <= 9223372036854775807:
		return "int64"
	}

	return "Out of range"
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
