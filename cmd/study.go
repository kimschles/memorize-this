/*
Copyright © 2021 Kim Schlesinger

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// studyCmd represents the study command
var studyCmd = &cobra.Command{
	Use:   "study <file-name>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		readFile(args)
	},
}

func init() {
	rootCmd.AddCommand(studyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// studyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// studyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readFile(args []string) {
	csvFile := args[0]

	file, err := os.Open(csvFile)
	handleErr(err)

	reader := csv.NewReader(file)

	var questions []string
	var answers []string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		handleErr(err)
		// appended to a slice?
		questions = append(questions, record[0])
		answers = append(answers, record[1])

		// err = file.Close()
	}
	askQuestion(questions[0])
	showAnswer(answers[0])
	recordDifficultyRating()
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func askQuestion(question string) {
	fmt.Println(question)
}

func showAnswer(answer string) {
	fmt.Println(answer)
}

func recordDifficultyRating() {
	fmt.Println("Easy (E) or Hard (H)?")
	inputReader := bufio.NewReader(os.Stdin)
	userText, err := inputReader.ReadString('\n')
	handleErr(err)
	fmt.Println(userText)
}
