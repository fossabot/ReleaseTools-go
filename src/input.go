package main

import (
	"bufio"
	"fmt"
	"github.com/tcnksm/go-gitconfig"
	"os"
	"strconv"
	"strings"
)

/**
 * Default questions
 */
const ASK_VERSION = "What version?"
const ASK_TITLE = "Whats the summary of this merge request?"
const ASK_MRID = "What merge request id?"
const ASK_REPONAME = "Whats the name of the new repository?"

/**
 * Struct for slice AskedQuestions
 */
type questions struct {
	question string
	awnser   string
}

/**
 * Slice with all asked questions
 */
var AskedQuestions []questions

var active bool = false

/**
 * Ask an question and save it to a slice so if asked again,
 * we can just return the already asked answer
 */
func askQuestion(question string) string {
	if len(AskedQuestions) != 0 {
		for index, _ := range AskedQuestions {
			if AskedQuestions[index].question == question {
				active = false
				return AskedQuestions[index].awnser
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)

	// For some reason when the iteration is high enough
	// the platform will print two times? Weird, work-a-round
	// in use by var active
	if active == false {
		fmt.Print(question + ": ")
		active = true
	}

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)

	if len(text) < 1 {
		return askQuestion(question)
	}

	AskedQuestions = append(AskedQuestions, questions{question, text})
	active = false
	return text
}

/**
 * Set answer before asking!
 */
func setAwnser(question string, answer string) {
	AskedQuestions = append(AskedQuestions, questions{question, answer})
}

/**
 * Ask user for confirmation
 */
func askConfirmation() bool {
	var s string

	fmt.Printf("Do you want to continue? (y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}

	return false
}

func askMergeType() string {
	mergeRequestTypes := [8]string{
		"Other",
		"New feature",
		"Bug fix",
		"Feature change",
		"New deprecation",
		"Feature removal",
		"Security fix",
		"Style fix",
	}

	for index, element := range mergeRequestTypes {
		fmt.Printf("[%d] %s\n", index, element)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is this merge request about? [0-7]: ")
	index, _ := reader.ReadString('\n')

	index = strings.Replace(index, "\r", "", -1)
	index = strings.Replace(index, "\n", "", -1)

	int64Index, _ := strconv.ParseInt(index, 10, 64)

	if len(mergeRequestTypes) >= int(int64Index) {
		return mergeRequestTypes[int64Index]
	} else {
		panic("Input is out of bounds!")
	}
}

func askUsername() string {

	reader := bufio.NewReader(os.Stdin)

	name, _ := gitconfig.Username()

	fmt.Print("What is your name [", name, "]: ")
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)

	if len(text) < 3 {
		text = name
	}

	return text
}
