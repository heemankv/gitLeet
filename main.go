package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/heemankv/LeetCode-Solutions/helpers"
	"github.com/joho/godotenv"
)

// @dev this a function that returns a slice of all the questions solved by the user
// @returns list of objects containing frontendQuestionId, titleSlug, Name
func getListOfAllQuestionsSolved( skip, limit int, status ,cookie string ) (questions []interface {}, total int){
	// By default, Golang defines all the named variables with the zero value and function will able to use them. In case function doesn’t modify the values then automatically zero value will return.
	// So no need for the below 2 lines
	// var questions []interface {}
	// var total int

	var query = fmt.Sprintf("{\"query\":\"\\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\\n  problemsetQuestionList: questionList(\\n    categorySlug: $categorySlug\\n    limit: $limit\\n    skip: $skip\\n    filters: $filters\\n  ) {\\n    total: totalNum\\n    questions: data {\\n      frontendQuestionId: questionFrontendId\\n      difficulty\\n\\n      title\\n      titleSlug\\n    }\\n  }\\n}\\n    \",\"variables\":{\"categorySlug\":\"\",\"skip\":%d,\"limit\":%d,\"filters\":{\"status\":\"%s\"}}}",skip,limit,status)
	
	m := helpers.QueryWrapper(query, cookie)
	
	data := m["data"].(map[string]interface{})
	problemsetQuestionList :=  data["problemsetQuestionList"].(map[string]interface{})
	
	questions = problemsetQuestionList["questions"].([]interface{})
	total = int(problemsetQuestionList["total"].(float64))
	return
}

// @dev this is a function that returns a description the given problem slug 
// @returns a formatted string explaining the question
// @params takes in title-slug 
func getQuestionDescription(titleSlug , cookie string) (description string ){
	// make request to fetch question Description
	var query = fmt.Sprintf("{\"query\":\"\\n    query questionContent($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    content\\n  }\\n}\\n    \",\"variables\":{\"titleSlug\":\"%s\"}}", titleSlug)
	
	m := helpers.QueryWrapper(query, cookie)
	data := m["data"].(map[string]interface{})
	questionDesc :=  data["question"].(map[string]interface{})	
	// format the question description 
	description = questionDesc["content"].(string)
	return
}

// @dev this is a function that returns user's solution to the given questionId 
// @returns a string that is the solution and a timestamp as to when the soln was submitted 
// @params takes in questionId
func getQuestionSolution(lang int, frontendId, cookie string)(timestamp int, code string, validation bool){
	// make request to fetch solution
	// format the solution 
	// return the formatted value, timestamp
	var query = fmt.Sprintf("{\"query\":\"\\n  query syncedCode($questionId: Int!, $lang: Int!) {\\n  syncedCode(questionId: $questionId, lang: $lang) {\\n    timestamp\\n    code\\n  }\\n}\\n    \",\"variables\":{\"lang\": %d,\"questionId\": %s}}", lang, frontendId)
	
	m := helpers.QueryWrapper(query, cookie)
	data := m["data"].(map[string]interface{})
	problemsetQuestionList, ok :=  data["syncedCode"].(map[string]interface{})
	if(!ok){
		code =""
		timestamp = 1
		validation = false
		return 
	}
	timestamp = int(problemsetQuestionList["timestamp"].(float64))
	code = problemsetQuestionList["code"].(string)
	validation = true
	return
}



func makeCodeFile(object map[string]interface{}, folderName string){
	fileName := fmt.Sprintf("%s/%s.cpp",folderName,folderName)
	file, err := os.Create(fileName)
     
	if err != nil {
			log.Fatalf("failed creating file: %s", err)
	}
	// file is made now time for code put in

	file, errr := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if errr != nil {
		log.Fatalf("failed creating file: %s", errr)
	}
 
	datawriter := bufio.NewWriter(file)
 
	datawriter.WriteString((object["solution"]).(string) + "\n")
 
	datawriter.Flush()
	file.Close()
}

func makeReadmeFile(object map[string]interface{}, folderName string){
	// make a readme file in the given folder
	// add appropriate data to the file
	// return and handle error
	fileNameArr := []string{folderName ,"Readme.md"}
	fileName := strings.Join(fileNameArr, "/")
	file, err := os.Create(fileName)
     
	if err != nil {
			log.Fatalf("failed creating file: %s", err)
	}
	 
	// closing the running file after the main
	// method has completed execution and
	// the writing to the file is complete
	
	// What all to write in the Readme
	// Difficulty : title
	// Submission Time in Date format
	// question link
	// question content
	// writing data to the file using
	// WriteString() method and the
	// length of the string is stored
	// in len variable

	firstLine := []string{(object["difficulty"]).(string), (object["title"]).(string)}
	
	secondSegment :=time.Unix(int64((object["submissionTime"]).(int)), 0).Local().String()
	secondLine := []string{"Submission Time", secondSegment}

	thirdLine := fmt.Sprintf("https://leetcode.com/problems/%s/", object["titleSlug"])

	converter := md.NewConverter("", true, nil)
	html := (object["description"]).(string)

	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}

	sampledata := []string{ 
		strings.Join(firstLine, " : "),
		strings.Join(secondLine, " : "),
		thirdLine,
		markdown,
	}
 
	file, errr := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if errr != nil {
		log.Fatalf("failed creating file: %s", errr)
	}
 
	datawriter := bufio.NewWriter(file)
 
	for _, data := range sampledata {
		_, _ = datawriter.WriteString(data + "\n")
	}
 
	datawriter.Flush()
	file.Close()

}

func makeFolderGivenName(title string)(output string, eror error){
	title = strings.ReplaceAll(title, " ", "_") // folderName is always made with underScore
  err := os.Mkdir(title, 0755) //create a directory and give it required permissions
	return title , err
}

// @params this function takes in the title, title-slug and timestamp
// @devs this function commits the folder given by title with commit name title-slug solved, at the time given by timestamp
func commitFolderToGithub(object map[string]interface{}, excType string ){
	// GOOD: Lets your script choose its own interpreter
	excTime := fmt.Sprint(((object["submissionTime"]).(int)))
	excMessage := fmt.Sprintf("Solved: %s : %s ",object["difficulty"], object["title"])
	cmd := exec.Command("./helpers/bash/github.sh", excType, excTime, excMessage)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(stdout)
		return
	}
}


func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	cookie := os.Getenv("COOKIE")

	// Default Values
	skip := 0
	limit := 9999
	status := "AC"
	lang := 0

	fmt.Println("Getting List of all questions...")
	questions, total := getListOfAllQuestionsSolved(skip, limit, status, cookie)
	fmt.Println("Found" , total, " Questions")
	var finalValue []map[string]interface{}
	for i := 110; i < total ; i++ {
		input := questions[i].(map[string]interface{})
		submissionTime, solution, validation := getQuestionSolution(lang, input["frontendQuestionId"].(string), cookie)
		if(!validation){
			fmt.Println("Empty Data fetched for: ", i ," : ", input["title"].(string), " Skipping")
			continue
		}
		var question = make(map[string]interface{}) 
		question["description"] = getQuestionDescription((input["titleSlug"]).(string), cookie)
		question["difficulty"] = input["difficulty"].(string)
		question["frontendQuestionId"] = input["frontendQuestionId"].(string)
		question["title"] = input["title"].(string)
		question["titleSlug"] = (input["titleSlug"]).(string)
		question["solution"] = solution
		question["submissionTime"] = submissionTime
		fmt.Println("Data fetched for: ", i ," : ", (question["title"]).(string))
		finalValue = append(finalValue,question)
	}
	// Sort Interface by submissionTime
	
	sort.Slice(finalValue, func(i, j int) bool {
		input1 := finalValue[i]
		input2 := finalValue[j]
		return (input1["submissionTime"]).(int) < (input2["submissionTime"]).(int)
		})

	for i := 0; i < len(finalValue) ; i++ {
		slug := finalValue[i]
		output , err := makeFolderGivenName((slug["title"]).(string))
		if err != nil {
      fmt.Println(err)
			break;
   }
	 makeReadmeFile(slug, output)
	 makeCodeFile(slug,output)
	 commitFolderToGithub(slug, "commit")
	 fmt.Println("Folder Ready and Commited for: ", (slug["title"]).(string))

	}

	var userChoice string
	fmt.Println("Ready to push? : ")

	fmt.Scanln(&userChoice)
	if(userChoice == "push"){
		commitFolderToGithub(finalValue[0], "push")
		fmt.Println("Done")
	}
	
}