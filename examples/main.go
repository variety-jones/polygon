package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/variety-jones/polygon"
)

// A sample credentials.txt
/*
{
	"apiKey" :  "KEY_HERE",
	"secret" : "SECRET_HERE",
	"problemId" : "PROBLEM_ID_HERE"
}
*/

// Reads a credential.txt from your current directory and finds a JSON with 3 keys
func CreateApiObjectFromLocal() (api polygon.PolygonApi) {
	api = polygon.PolygonApi{}

	// Uncomment this part if you want to write the key manually
	/*
		api.ApiKey = "KEY_HERE"
		api.Secret = "SECRET_HERE"
		api.ProblemId = "ID_HERE"
		return api
	*/

	data, err := os.ReadFile("credentials.txt")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &api)
	if err != nil {
		panic(err)
	}

	return api
}

// Demonstration of how to view a problem's info
func ProblemInfoExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	problemInfo, err := api.ProblemInfo(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	prettyInfo, err := problemInfo.Prettify()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(prettyInfo)
}

// Demonstration of how to update a problem's info
func ProblemUpdateInfoExample(api polygon.PolygonApi, interactive string) {
	parameters := make(map[string]string)
	parameters["timeLimit"] = "3000"
	parameters["memoryLimit"] = "512"
	parameters["interactive"] = interactive
	err := api.ProblemUpdateInfo(parameters)
	if err != nil {
		fmt.Println(err)
	}
}

// Demonstration of how to view a problem's statements
func ProblemStatementsExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	statementsMap, err := api.ProblemStatements(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}

	for key, value := range statementsMap {
		fmt.Printf("Here is the statement for the language: %v\n", key)
		prettyStatement, err := value.Prettify()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(prettyStatement)
	}
}

// Demonstration of how to update the problem statement
func ProblemSaveStatementExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["lang"] = "english"
	parameters["name"] = "A Plus B"
	parameters["legend"] = "Hello, this is the latest version of the english statement."
	parameters["notes"] = "I just added some notes here."
	err := api.ProblemSaveStatement(parameters)
	if err != nil {
		fmt.Println(err)
	}
}

// Demonstration of how to print all statement resources
func ProblemStatementResourcesExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	files, err := api.ProblemStatementResources(parameters)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		prettyFile, err := file.Prettify()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(prettyFile)
		}
	}
}

// Demonstration of how to update a problem's resource
func ProblemSaveStatementResourceExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["checkExisting"] = "false"
	parameters["name"] = "hello-world.txt"
	parameters["file"] = "Hi there, this resource just got updated"
	err := api.ProblemSaveStatementResource(parameters)
	if err != nil {
		fmt.Println(err)
	}
}

// Demonstration of how to view the currently set checker
func ProblemCheckerExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	checker, err := api.ProblemChecker(parameters)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(checker)
}

// Demonstration of how to set a checker
func ProblemSetCheckerExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["checker"] = "std::lcmp.cpp"
	err := api.ProblemSetChecker(parameters)
	if err != nil {
		fmt.Println(err)
	}
}

// Demonstration of how to view the currently set validator
func ProblemValidatorExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	validator, err := api.ProblemValidator(parameters)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(validator)
}

// Demonstration of how to set a validator
func ProblemSetValidatorExample(api polygon.PolygonApi, validator string) {
	parameters := make(map[string]string)
	parameters["validator"] = validator
	err := api.ProblemSetValidator(parameters)
	if err != nil {
		fmt.Println(err)
	}
}

// Demonstration of how to view the currently set interactor
func ProblemInteractorExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	interactor, err := api.ProblemInteractor(parameters)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(interactor)
}

// Demonstration of how to set a validator
func ProblemSetInteractorExample(api polygon.PolygonApi, interactor string) {
	parameters := make(map[string]string)
	parameters["interactor"] = interactor
	err := api.ProblemSetInteractor(parameters)
	if err != nil {
		fmt.Println(err)
	}
}

// Demonstration of how to print the problem resource, source and aux files metadata
func ProblemFilesExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	rsa, err := api.ProblemFiles(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Uncomment the commented outputs if you want verbose output
	/*
		fmt.Println("Here are the resource files of the problem")
		for _, resourceFile := range rsa.ResourceFiles {
			prettyFile, err := resourceFile.Prettify()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(prettyFile)
		}
	*/

	fmt.Println("Here are the source files of the problem")
	for _, sourceFile := range rsa.SourceFiles {
		prettyFile, err := sourceFile.Prettify()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(prettyFile)
	}

	/*
		fmt.Println("Here are the aux files of the problem")
		for _, auxFile := range rsa.AuxFiles {
			prettyFile, err := auxFile.Prettify()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(prettyFile)
		}
	*/
}

// Demonstration of how to view an individual file
func ProblemViewFileExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["name"] = "demo.cpp"
	parameters["type"] = "source"
	view, err := api.ProblemViewFile(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(view)
}

// Demonstration of how to add a source/resource/aux file
func ProblemSaveFileExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)

	// Let's get a random file from my Github gists.
	// You can also upload your local ".cpp" file.
	// Just read its contents and store it as a string
	contents, err := extractRandomCodeFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	parameters["checkExisting"] = "false"
	parameters["name"] = "demo.cpp"
	parameters["type"] = "source"
	parameters["file"] = contents

	err = api.ProblemSaveFile(parameters)
	if err != nil {
		fmt.Println(err)
	}

}

// Demonstration of how to add/edit a problem's solution
func ProblemSaveSolutionExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["name"] = "sol.cpp"
	codeFile, err := extractRandomCodeFile()
	if err != nil {
		fmt.Println(err)
	}

	parameters["file"] = codeFile
	err = api.ProblemSaveSolution(parameters)
	if err != nil {
		fmt.Println(err)
	}
}

// Demonstration of how to get metadata of all solutions
func ProblemSolutionsExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	solutions, err := api.ProblemSolutions(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, solution := range solutions {
		prettySolution, err := solution.Prettify()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(prettySolution)
	}
}

// Demonstration of how to view an individual solution
func ProblemViewSolutionExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["name"] = "sol.cpp"
	solution, err := api.ProblemViewSolution(parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(solution)
	}
}

// Demonstration of how to view a testset's script
func ProblemScriptExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	// "tests" is the default name
	parameters["testset"] = "tests"
	script, err := api.ProblemScript(parameters)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(script)
	}
}

// Demonstration of how to update a problem's script
func ProblemSaveScriptExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	// "tests" is the default name
	parameters["testset"] = "tests"

	// You can also add your local script
	parameters["source"] = `demo seed_1 > 1
	demo seed_2 > 2
	demo seed_3 > 3
	demo seed_4 > 4`
	err := api.ProblemSaveScript(parameters)
	if err != nil {
		fmt.Println(err)
	}
}

// Demonstration of how to print all test resources
func ProblemTestsExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["testset"] = "tests"
	tests, err := api.ProblemTests(parameters)
	if err != nil {
		fmt.Println(err)
	}

	for _, test := range tests {
		prettyTest, err := test.Prettify()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(prettyTest)
		}
	}
}

// Demonstration of how to view the test input at a particular index
func ProblemTestInputExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["testset"] = "tests"
	parameters["testIndex"] = "3"
	input, err := api.ProblemTestInput(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)
}

// Demonstration of how to view the test answer at a particular index
func ProblemTestAnswerExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["testset"] = "tests"
	parameters["testIndex"] = "3"
	answer, err := api.ProblemTestAnswer(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(answer)
}

// Demonstration of how to view the general description of the problem
func ProblemViewGeneralDescriptionExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	view, err := api.ProblemViewGeneralDescription(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(view)
}

// Demonstration of how to save the general description of the problem
func ProblemSaveGeneralDescriptionExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["description"] = "The problem description just got updated"

	err := api.ProblemSaveGeneralDescription(parameters)
	if err != nil {
		fmt.Println(err)
	}
}

// Demonstration of how to view the general tutorial of the problem
func ProblemViewGeneralTutorialExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	view, err := api.ProblemViewGeneralTutorial(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(view)
}

// Demonstration of how to save the general tutorial of the problem
func ProblemSaveGeneralTutorialExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)
	parameters["tutorial"] = "The tutorial just got updated"

	err := api.ProblemSaveGeneralTutorial(parameters)
	if err != nil {
		fmt.Println(err)
	}
}

// Demonstration of how to view package metadata
func ProblemPackagesExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)

	packages, err := api.ProblemPackages(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, pacakgeObj := range packages {
		prettyPackage, err := pacakgeObj.Prettify()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(prettyPackage)
	}
}

// Demonstration of how to view all problems of a certain user
func ProblemsListExample(api polygon.PolygonApi) {
	parameters := make(map[string]string)

	problems, err := api.ProblemsList(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, problem := range problems {
		prettyProblem, err := problem.Prettify()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(prettyProblem)
	}
}

// A utility function to download a random CPP file from the internet
func extractRandomCodeFile() (code string, err error) {
	location := "https://gist.githubusercontent.com/variety-jones/478ffaaaf23c007f2e3c8954b61677d4/raw/0314dd1a6926d0c21e84b726eacfbe4a5d37b654/demo.cpp"
	resp, err := http.Get(location)
	if err != nil {
		return code, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return code, err
	}

	return string(body), err

}

func main() {
	//api := polygon.PolygonApi{ApiKey: apiKey, Secret: secret, ProblemId: problemId}

	api := CreateApiObjectFromLocal()

	// Examples on how to update problem info
	fmt.Println("Here is the old info: ")
	ProblemInfoExample(api)
	ProblemUpdateInfoExample(api, "true")
	fmt.Println("Here is the new info: ")
	ProblemInfoExample(api)

	// Examples on how to update problem statements
	fmt.Println("Here is the old statement: ")
	ProblemStatementsExample(api)
	ProblemSaveStatementExample(api)
	fmt.Println("Here is the new statement: ")
	ProblemStatementsExample(api)

	// Examples on how to update problem statements
	fmt.Println("Here is the old statement resources: ")
	ProblemStatementResourcesExample(api)
	ProblemSaveStatementResourceExample(api)
	fmt.Println("Here is the new statement resources: ")
	ProblemStatementResourcesExample(api)

	// Examples on how to update checker
	fmt.Println("Here is the old checker: ")
	ProblemCheckerExample(api)
	ProblemSetCheckerExample(api)
	fmt.Println("Here is the new checker: ")
	ProblemCheckerExample(api)

	// Now, we'll update validator and interactor, but we don't have any code files as of now.
	// Let's add a demo.cpp to simulate a validator and interactor
	ProblemSaveFileExample(api)
	// We'll also inspect the list of files once
	ProblemFilesExample(api)
	// Let's inspect the contents of the file that we've added
	fmt.Println("Here is the file that we just uploaded")
	ProblemViewFileExample(api)

	// Now that we've added a code file, let's set the validator and interactor to that.
	// Examples on how to update validator
	fmt.Println("Here is the old validator: ")
	ProblemValidatorExample(api)
	ProblemSetValidatorExample(api, "demo.cpp")
	fmt.Println("Here is the new validator: ")
	ProblemValidatorExample(api)

	// Examples on how to update interactor
	fmt.Println("Here is the old interactor: ")
	ProblemInteractorExample(api)
	ProblemSetInteractorExample(api, "demo.cpp")
	fmt.Println("Here is the new interactor: ")
	ProblemInteractorExample(api)

	// Let's make problem non interactive, and remove validators, as they weren't correct.
	ProblemSetValidatorExample(api, "")
	ProblemSetInteractorExample(api, "")
	ProblemUpdateInfoExample(api, "false")

	// Examples on how to update solution
	// First, we add a new solution
	ProblemSaveSolutionExample(api)
	// Now, let's view the solution metadata
	fmt.Println("Here are the solution files metadata")
	ProblemSolutionsExample(api)
	// Finally, print an actual solution file
	fmt.Println("These are the contents of sol.cpp")
	ProblemViewSolutionExample(api)

	// Examples on how to update a script
	ProblemSaveScriptExample(api)
	fmt.Println("Here is the updated script: ")
	ProblemScriptExample(api)

	// Now that we've updated the script, let's inspect test metadata
	fmt.Println("Here is the test metadata: ")
	ProblemTestsExample(api)
	// Let's view the input and output of test and index 3.
	fmt.Println("Input of the test at desired index is ")
	ProblemTestInputExample(api)
	fmt.Println("Output of the test at desired index is")
	ProblemTestAnswerExample(api)

	// Let's update general description
	ProblemSaveGeneralDescriptionExample(api)
	fmt.Println("Here is the new general description: ")
	ProblemViewGeneralDescriptionExample(api)

	// Let's update general tutorial
	ProblemSaveGeneralTutorialExample(api)
	fmt.Println("Here is the new general tutorial: ")
	ProblemViewGeneralTutorialExample(api)

	// Let's inspect the package metadata
	// At this stage, there's no pacage created yet, bbut you can commit the changes and create a package
	// Then, run this example (or whole program again), it would be reflected.
	ProblemPackagesExample(api)

	// Finally, let's view a non problem specific method
	// We will print meta data of all problems for the current user.
	ProblemsListExample(api)

	// That's all! Have left out a couple of examples related to tests, but it's similar to the above ones.
	// Make sure to read the documentation to get more insights.
}
