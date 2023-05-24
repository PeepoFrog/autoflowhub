package main

import (
	"log"

	"github.com/mrlutik/autoflowhub/pkg/keygen/docker"
	"github.com/mrlutik/autoflowhub/pkg/keygen/usecase"
	"github.com/spf13/cobra"
)

const (
	use              = "keysgen"
	shortDescription = "CLI application for generate accounts for Kira Network"
)

const longDescription = `This application accepts three parameters: 
home of sekaid, 
keyring-backend value
count of keys which will be generated
directory for saved address and mnemonic for keys
There is no default values!`

var rootCmd = &cobra.Command{
	Use:   use,
	Short: shortDescription,
	Long:  longDescription,
	Run: func(cmd *cobra.Command, _ []string) {
		home, _ := cmd.Flags().GetString("home")
		keyringBackend, _ := cmd.Flags().GetString("keyring-backend")
		dirOfKeys, _ := cmd.Flags().GetString("keys-dir")
		sekaiContainer, _ := cmd.Flags().GetString("sekai")
		count, _ := cmd.Flags().GetInt("count")

		if home == "" || sekaiContainer == "" || keyringBackend == "" || dirOfKeys == "" || count <= 0 {
			cmd.Help()
			log.Fatal("Please provide all required parameters: home, backend and positive count")
		}

		log.Println("Sekai Container:", sekaiContainer)
		log.Println("Home:", home)
		log.Println("Backend:", keyringBackend)
		log.Println("Directory of keys:", dirOfKeys)
		log.Println("Count:", count)

		generating(sekaiContainer, home, keyringBackend, dirOfKeys, count)
	},
}

func main() {
	// Usage: keygen --home "/root/.sekaid-testnetwork-1" -k "test" -c 4 -d ./data -s sekai
	rootCmd.PersistentFlags().String("home", "", "Path to sekaid home")
	rootCmd.PersistentFlags().StringP("keys-dir", "d", "", "Keys directory (relative or absolute path)")
	rootCmd.PersistentFlags().StringP("keyring-backend", "k", "", "Keyring backend")
	rootCmd.PersistentFlags().StringP("sekai", "s", "", "Sekaid container name")
	rootCmd.PersistentFlags().IntP("count", "c", 0, "Count of keys which will be added")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func generating(containerName, homePath, keyringBackend, dirOfKeys string, count int) {
	dockerClient := docker.NewDockerCommandRunner()
	keysUsecase := usecase.NewKeysClient(dockerClient, containerName, homePath, keyringBackend, dirOfKeys)

	var err error
	addresses, err := keysUsecase.GenerateKeys(count)
	if err != nil {
		log.Fatal(err)
	}

	allAddresses, err := keysUsecase.ListOfKeys()
	if err != nil {
		log.Fatal(err)
	}

	// NEXT BUSINESS LOGIC HERE
	// allAddresses includes all users in Kira network
	// addresses includes only generated keys

	log.Println("Checking generated addresses in the list of all...")
	if containsAllStrings(allAddresses, addresses) {
		log.Fatal("Error: not all generated addresses are saved")
	}

	log.Println("All is O'key!")
}

// Temporary helpers which are used for checking program
func sliceContains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func containsAllStrings(slice1 []string, slice2 []string) bool {
	for _, b := range slice2 {
		if !sliceContains(slice1, b) {
			return false
		}
	}
	return true
}