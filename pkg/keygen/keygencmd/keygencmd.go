package keygencmd

import (
	"log"

	"github.com/mrlutik/autoflowhub/pkg/keygen/docker"
	"github.com/mrlutik/autoflowhub/pkg/keygen/usecase"
	"github.com/spf13/cobra"
)

const (
	use              = "keygen"
	shortDescription = "command to generate accounts for Kira Network testnet"
	longDescription  = "This command accepts three parameters: home of sekaid, keyring-backend value count of keys which will be generated directory for saved address and mnemonic for keys There is no default values!"
)

func New() *cobra.Command {
	var keygenCmd = &cobra.Command{
		Use:   use,
		Short: shortDescription,
		Long:  longDescription,
		Run: func(cmd *cobra.Command, _ []string) {
			home, err := cmd.Flags().GetString("home")
			if err != nil {
				log.Fatalf("Error reading home flag: %v", err)
			}

			keyringBackend, err := cmd.Flags().GetString("keyring-backend")
			if err != nil {
				log.Fatalf("Error reading keyring-backend flag: %v", err)
			}

			dirOfKeys, err := cmd.Flags().GetString("keys-dir")
			if err != nil {
				log.Fatalf("Error reading keys-dir flag: %v", err)
			}

			sekaiContainer, err := cmd.Flags().GetString("sekai")
			if err != nil {
				log.Fatalf("Error reading sekai flag: %v", err)
			}

			count, err := cmd.Flags().GetInt("count")
			if err != nil {
				log.Fatalf("Error reading count flag: %v", err)
			}

			if home == "" || sekaiContainer == "" || keyringBackend == "" || dirOfKeys == "" || count <= 0 {
				cmd.Help()
				log.Fatal("Please provide all required parameters: home, sekai, backend, keys-dir and a positive count")
			}

			log.Println("Sekai Container: ", sekaiContainer)
			log.Println("Home: ", home)
			log.Println("Backend: ", keyringBackend)
			log.Println("Keys directory path: ", dirOfKeys)
			log.Println("Count: ", count)

			generating(sekaiContainer, home, keyringBackend, dirOfKeys, count)
		},
	}

	keygenCmd.PersistentFlags().String("home", "", "Path to sekaid home")
	keygenCmd.PersistentFlags().StringP("keys-dir", "d", "", "Keys directory path(relative or absolute path)")
	keygenCmd.PersistentFlags().StringP("keyring-backend", "k", "", "Keyring backend")
	keygenCmd.PersistentFlags().StringP("sekai", "s", "", "Sekaid container name")
	keygenCmd.PersistentFlags().IntP("count", "c", 0, "Count of keys which will be added")

	return keygenCmd
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
	// create a map to store the presence of the string
	map1 := make(map[string]bool)
	for _, val := range slice1 {
		map1[val] = true
	}

	for _, val := range slice2 {
		if !map1[val] {
			return false
		}
	}
	return true
}
