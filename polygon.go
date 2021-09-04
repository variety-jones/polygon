package polygon

import (
	"encoding/json"
)

// ProblemsList returns a list of problems, available to the user, according to search parameters.
//
// Parameters
//
// showDeleted : bool, optional - show/hide deleted problems (defaults to false)
//
// id : problem id, optional
//
// name : problem name, optional
//
// owner : problem owner login, optional
//
// Returns
//
// A list of Problem objects.
func (api *PolygonApi) ProblemsList(parameters map[string]string) (problems []ProblemObject, err error) {
	body, err := api.processRequest(parameters, problemsListEp)
	if err != nil {
		return problems, err
	}

	wrapper := wrapperProblemSlice{}
	err = json.Unmarshal(body, &wrapper)
	return wrapper.Result, err
}

// ProblemInfo return a ProbelmInfoObject representing metadata about the problem
func (api *PolygonApi) ProblemInfo(parameters map[string]string) (problemInfo ProblemInfoObject, err error) {
	body, err := api.processRequest(parameters, problemInfoEp)
	if err != nil {
		return problemInfo, err
	}

	wrapper := wrapperProblemInfo{}
	err = json.Unmarshal(body, &wrapper)
	return wrapper.Result, err
}

// ProblemUpdateInfo updates problem info.
// All parameters are optional.
//
// Parameters
//
// inputFile : problem’s input file
//
// outputFile : problem’s output file
//
// interactive : boolean - is problem interactive
//
// timeLimit : problem’s time limit in milliseconds
//
// memoryLimit : problem’s memory limit in MB
func (api *PolygonApi) ProblemUpdateInfo(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemUpdateInfoEp)
}

// ProblemStatements returns a map from language to a Statement object for that language.
func (api *PolygonApi) ProblemStatements(parameters map[string]string) (statementsMap map[string]StatementObject, err error) {
	body, err := api.processRequest(parameters, problemStatementsEp)
	if err != nil {
		return statementsMap, err
	}

	wrapper := wrapperStatementMap{}
	err = json.Unmarshal(body, &wrapper)
	return wrapper.Result, err
}

// ProblemSaveStatement updates or creates a problem’s statement.
// All parameters except for lang are optional.
//
// Parameters
//
// lang : (required) - statement’s language
//
// encoding : statement’s encoding
//
// name : problem’s name in statement’s language
//
// legend : problem’s language
//
// input : problem’s input format
//
// output : problem’s output format
//
// scoring : problem’s scoring
//
// notes : statement notes
//
// tutorial : problem’s tutorial
func (api *PolygonApi) ProblemSaveStatement(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSaveStatementEp)
}

// ProblemStatementResources returns a list of statement resources for the problem.
//
// Parameters
//
// None
func (api *PolygonApi) ProblemStatementResources(parameters map[string]string) (files []FileObject, err error) {
	body, err := api.processRequest(parameters, problemStatementResourcesEp)
	if err != nil {
		return files, err
	}

	wrapper := wrapperFileSlice{}
	err = json.Unmarshal(body, &wrapper)
	if err != nil {
		return files, err
	}
	return wrapper.Result, err

}

// ProblemSaveStatementResource adds or edit statement resource file
//
// Parameters
//
// checkExisting : boolean, optional - if true, only adding files is allowed
//
// name : file name
//
// file : file content
func (api *PolygonApi) ProblemSaveStatementResource(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSaveStatementResourceEp)
}

// ProblemChecker returns the name of currently set checker.
func (api *PolygonApi) ProblemChecker(parameters map[string]string) (checkerName string, err error) {
	return api.extractName(parameters, problemCheckerEp)
}

// ProblemValidator returns the name of currently set validator
func (api *PolygonApi) ProblemValidator(parameters map[string]string) (validatorName string, err error) {
	return api.extractName(parameters, problemValidatorEp)
}

// ProblemInteractor returns the name of currently set interactor
func (api *PolygonApi) ProblemInteractor(parameters map[string]string) (interactorName string, err error) {
	return api.extractName(parameters, problemInteractorEp)
}

// ProblemFiles returns the list of resource, source and aux files.
// Method returns a JSON object with three fields: resource, source, aux.
// EAch of them is a list of FileObject
func (api *PolygonApi) ProblemFiles(parameters map[string]string) (rsa RsaObject, err error) {
	body, err := api.processRequest(parameters, problemFilesEp)
	if err != nil {
		return rsa, err
	}

	wrapper := wrapperRSA{}
	err = json.Unmarshal(body, &wrapper)
	if err != nil {
		return rsa, err
	}
	return wrapper.Result, err
}

// ProblemSolutions returns the list of Solution objects.
func (api *PolygonApi) ProblemSolutions(parameters map[string]string) (solutions []SolutionObject, err error) {
	body, err := api.processRequest(parameters, problemSolutionsEp)
	if err != nil {
		return solutions, err
	}

	wrapper := wrapperSolutionSlice{}
	err = json.Unmarshal(body, &wrapper)
	if err != nil {
		return solutions, err
	}
	return wrapper.Result, err
}

// ProblemViewFile returns resource, source or aux file.
// It returns plain view of the file with the corresponding mime-type set.
//
// Parameters
//
// type : resource/aux/source - requested file’s type
//
// name : file name
func (api *PolygonApi) ProblemViewFile(parameters map[string]string) (fileView string, err error) {
	return api.extractView(parameters, problemViewFileEp)
}

// ProblemViewSolution returns a view of the solution file
//
// Parameters
//
// name : solution’s name
func (api *PolygonApi) ProblemViewSolution(parameters map[string]string) (solutionView string, err error) {
	return api.extractView(parameters, problemViewSolutionEp)
}

// ProblemScript returns script for generating tests. it returns plain view of the script.
//
// Parameters
//
// testset : testset for which the script is requested
func (api *PolygonApi) ProblemScript(parameters map[string]string) (scriptView string, err error) {
	return api.extractView(parameters, problemScriptEp)
}

// ProblemTests Rreturns tests for the given testset
// It returns a list of Test objects.
//
// Parameters
//
// testset : testset for which tests are requested
func (api *PolygonApi) ProblemTests(parameters map[string]string) (tests []TestObject, err error) {
	body, err := api.processRequest(parameters, problemTestsEp)
	if err != nil {
		return tests, err
	}

	wrapper := wrapperTestSlice{}
	err = json.Unmarshal(body, &wrapper)
	return wrapper.Result, err
}

// ProblemTestInput returns generated test input.
// It returns plain view of the resource.
//
// Parameters
//
// testset : testset of the test
//
// testIndex : index of the test
func (api *PolygonApi) ProblemTestInput(parameters map[string]string) (testInputView string, err error) {
	return api.extractView(parameters, problemTestInputEp)
}

// ProblemTestAnswer returns generated test answer.
// It returns a plain view of the answer
//
// Parameters
//
// testset : testset of the test
//
// testIndex : index of the test
func (api *PolygonApi) ProblemTestAnswer(parameters map[string]string) (testAnswerView string, err error) {
	return api.extractView(parameters, problemTestAnswerEp)
}

// ProblemSetValidator updates validatdor
//
// Parameters
//
// validator : name of the validator (one of the source files)
func (api *PolygonApi) ProblemSetValidator(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSetValidatorEp)
}

// ProblemSetChecker updates checker
//
// Parameters
//
// checker : name of the checker (one of the source files)
func (api *PolygonApi) ProblemSetChecker(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSetCheckerEp)
}

// ProblemSetInteractor updates interactor
//
// Parameters
//
// interactor : name of the interactor (one of the source files)
func (api *PolygonApi) ProblemSetInteractor(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSetInteractorEp)
}

// ProblemSaveFile is used to add or edit resource, source or aux file.
// In case of editing, all parameters except for type and name are optional.
//
// Parameters
//
// checkExisting - boolean, optional - if true, only adding files is allowed
//
// type - file type (resource/source or aux)
//
// name - file name
//
// file - file content
//
// sourceType - optional - source type (in case of a source file)
//
// In case of type=resource there are some possible additional parameters
//
// forTypes - optional - semicolon separated list of applicable file types (see ResourceAdvancedProperties)
//
// stages - optional - semicolon separated list of values COMPILE or RUN,
// meaning the phase when the resource is applicable (currently only stages=COMPILE is supported)
//
// assets - optional - semicolon separated list of values VALIDATOR, INTERACTOR, CHECKER, SOLUTION,
// meaning the asset types the resource is applicable (currently only assets=SOLUTION is supported)
//
// The parameters forTypes, stages, assets can be present only together
// (it means, all of them are absent or all of them are used at the same time). They can be used only for type=resource.
// To delete ResourceAdvancedProperties of a resource file pass empty “forTypes=”.
func (api *PolygonApi) ProblemSaveFile(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSaveFileEp)
}

// ProblemSaveSolution adds or edits solution
// In case of editing, all parameters except for name are optional.
//
// Parameters
//
// checkExisting : boolean, optional - if true, only adding solutions is allowed
//
// name : solution name
//
// file : solution content
//
// sourceType : optional - source type
//
// tag : solution’s tag (MA - Main, OK, RJ, TL, TO - Time Limit or Accepted, WA, PE, ML or RE)
func (api *PolygonApi) ProblemSaveSolution(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSaveSolutionEp)
}

// ProblemEditSolutionExtraTags adds or remove testset or test group extra tag for solution.
//
// Parameters
//
// remove : boolean, if true - remove extra tag, if false - add extra tag
//
// name : solution name
//
// testset : optional - testset name for which edit extra tag
//
// testGroup : optional - test group name for which edit extra tag.
// Exactly one from testset and testGroup should be specified
//
// tag : optional - when you add extra tag - solution’s extra tag
// (OK, RJ, TL, TO - Time Limit or Accepted, WA, PE, ML or RE)
func (api *PolygonApi) ProblemEditSolutionExtraTags(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemEditSolutionExtraTagsEp)
}

// ProblemSaveScript edits script.
//
// Parameters
//
// testset : testset of the script
// source : script source
func (api *PolygonApi) ProblemSaveScript(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSaveScriptEp)
}

// ProblemSaveTest adds or edit test.
// In case of editing, all parameters except for testset and testIndex are optional.
//
// Parameters
//
// checkExisting : boolean, optional - if true, only adding new test is allowed
//
// testset : testset of the test
//
// testIndex : index of the test
//
// testInput : test input
//
// testGroup : optional - test group (groups should be enabled for the testset)
//
// testPoints : optional - test points (points should be enabled for the problem)
//
// testDescription : optional - test description
//
// testUseInStatements : bool, optional - whether to use test in statements
//
// testInputForStatements : optional - test input for viewing in the statements
//
// testOutputForStatements : optional - test output for viewing in the statements
//
// verifyInputOutputForStatements : bool, optional - whether to verify input and output for statements
func (api *PolygonApi) ProblemSaveTest(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSaveTestEp)
}

// ProblemSetTestGroup sets test group for one or more tests.
// It expects that for specified testset test groups are enabled.
//
// Parameters
//
// testset : testset of the test(s)
//
// testGroup : test group name to set
//
// testIndex : index of the test, you can specify multiple parameters with the same name testIndex,
// to set test group to many tests of the same testset at the same time.
//
// testIndices : list of test indices, separated by a comma.
// It’s alternative for testIndex, you should use only one from these two ways
func (api *PolygonApi) ProblemSetTestGroup(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSetTestGroupEp)
}

// ProblemEnableGroups enable or disable test groups for the specified testset.
//
// Parameters
//
// testset : testset to enable or disable groups
//
// enable : bool - if it is true test groups become enabled, else test groups become disabled
func (api *PolygonApi) ProblemEnableGroups(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemEnableGroupsEp)
}

// ProblemEnablePoints enable or disable test points for the problem.
//
// Parameters
//
// enable : bool - if it is true test points become enabled, else test points become disabled
func (api *PolygonApi) ProblemEnablePoints(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemEnablePointsEp)
}

// ProblemViewTestGroup returns test groups for the specified testset.
//
// Parameters
//
// testset : testset name
//
// group : optional - test group name
//
// Returns
//
// A list of TestGroup objects.
func (api *PolygonApi) ProblemViewTestGroup(parameters map[string]string) (testGroups []TestGroupObject, err error) {
	body, err := api.processRequest(parameters, problemViewTestGroupEp)
	if err != nil {
		return testGroups, err
	}

	wrapper := wrapperTestGroupSlice{}
	err = json.Unmarshal(body, &wrapper)
	return wrapper.Result, err
}

// ProblemSaveTestGroups Saves test group.
// Use if only to save a test group.
// If you want to create new test group, just add new test with such test group.
//
// Parameters
//
// testset : testset name
//
// group : test group name
//
// pointsPolicy : optional - COMPLETE_GROUP or EACH_TEST (leaves old value if no specified)
//
// feedbackPolicy : optional - NONE, POINTS, ICPC or COMPLETE (leaves old value if no specified)
//
// dependencies : optional - string of group names from which group should depends on separated by a comma
func (api *PolygonApi) ProblemSaveTestGroups(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSaveTestGroupEp)
}

// ProblemViewTags returns tags for the problem.
//
// Parameters
//
// None.
//
// Returns
//
// A list of strings – tags for the problem.
func (api *PolygonApi) ProblemViewTags(parameters map[string]string) (tags []string, err error) {
	body, err := api.processRequest(parameters, problemViewTagsEp)
	if err != nil {
		return tags, err
	}

	wrapper := wrapperStringSlice{}
	err = json.Unmarshal(body, &wrapper)
	return wrapper.Result, err
}

// ProblemSaveTags Saves tags for the problem. Existed tags will be replaced by new tags.
//
// Parameters
//
// tags – string of tags, separated by a comma. If you specified several same tags will be add only one of them.
func (api *PolygonApi) ProblemSaveTags(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSaveTagsEp)
}

// ProblemViewGeneralDescription returns problem general description.
//
// Parameters
//
// None.
//
// Returns
//
// A string – the problem general description.
func (api *PolygonApi) ProblemViewGeneralDescription(parameters map[string]string) (description string, err error) {
	body, err := api.processRequest(parameters, problemViewGeneralDescriptionEp)
	if err != nil {
		return description, err
	}

	wrapper := wrapperString{}
	err = json.Unmarshal(body, &wrapper)
	return wrapper.Result, err
}

// ProblemSaveGeneralDescription saves problem general description.
//
// Parameters
//
// description : string – the problem general description to save. The description may be empty.
func (api *PolygonApi) ProblemSaveGeneralDescription(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSaveGeneralDescriptionEp)
}

// ProblemViewGeneralTutorial returns problem general tutorial.
//
// Parameters
//
// None
//
// Returns
//
// A string – the problem general  tutorial.
func (api *PolygonApi) ProblemViewGeneralTutorial(parameters map[string]string) (tutorial string, err error) {
	body, err := api.processRequest(parameters, problemViewGeneralTutorialEp)
	if err != nil {
		return tutorial, err
	}

	wrapper := wrapperString{}
	err = json.Unmarshal(body, &wrapper)
	return wrapper.Result, err
}

// ProblemSaveGeneralTutorial saves problem general tutorial.
//
// Parameters:
//
// tutorial : string – the problem general tutorial to save. The tutorial may be empty.
func (api *PolygonApi) ProblemSaveGeneralTutorial(parameters map[string]string) (err error) {
	return api.checkForErrors(parameters, problemSaveGeneralTutorialEp)
}

// ProblemPackages returns a list of Package objects - list all packages available for the problem.
//
// Parameters
//
// None
func (api *PolygonApi) ProblemPackages(parameters map[string]string) (packages []PackageObject, err error) {
	body, err := api.processRequest(parameters, problemPackagesEp)
	if err != nil {
		return packages, err
	}

	wrapper := wrapperPackageSlice{}
	err = json.Unmarshal(body, &wrapper)
	return wrapper.Result, err
}

// ContestProblems returns a list of Problem objects - problems of the contest.
//
// Parameters
//
// contestId : The ID of the contest
func (api *PolygonApi) ContestProblems(parameters map[string]string) (problems []ProblemObject, err error) {
	body, err := api.processRequest(parameters, contestProblemsEp)
	if err != nil {
		return problems, err
	}

	wrapper := wrapperProblemSlice{}
	err = json.Unmarshal(body, &wrapper)
	return wrapper.Result, err
}

// Skipped : ProblemPackage
