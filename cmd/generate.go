/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const fileTemplate = `package {{ .Module }}

// TODO: Implement {{ .Filename }} logic
`

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [module-name]",
	Short: "Generate a new module with default files",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]
		basePath := filepath.Join("pkg", moduleName)

		// Create base directory
		if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}

		files := []string{"endpoints.go", "service.go", "request.go", "response.go"}
		for _, file := range files {
			filePath := filepath.Join(basePath, file)

			f, err := os.Create(filePath)
			if err != nil {
				log.Fatalf("Failed to create file: %v", err)
			}
			defer f.Close()

			tmpl, err := template.New(file).Parse(fileTemplate)
			if err != nil {
				log.Fatalf("Failed to parse template: %v", err)
			}

			err = tmpl.Execute(f, map[string]string{
				"Module":   moduleName,
				"Filename": file,
			})
			if err != nil {
				log.Fatalf("Failed to write to file: %v", err)
			}

			fmt.Printf("Created: %s\n", filePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
