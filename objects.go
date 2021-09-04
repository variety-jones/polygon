package polygon

import (
	"encoding/json"
)

// PolygonApi stores metadata for API calls
type PolygonApi struct {
	ApiKey    string
	Secret    string
	ProblemId string
}

// A handy struct to unmarshal response which returns a string as result
type wrapperString struct {
	Status  string
	Comment string
	Result  string
}

// A handy struct to unmarshal response which returns a string slice as result
type wrapperStringSlice struct {
	Status  string
	Comment string
	Result  []string
}

// ProblemObject represents a polygon problem
//
// Parameter Description
//
// Id            : problem id
//
// Owner         : problem owner handle
//
// Name          : problem name
//
// Deleted       : is a problem deleted
//
// Favourite     : is a problem in user's favourites
//
// AccessType    : user’s access type for the problem
//
// Revision      : current problem revision
//
// LatestPackage : latest revision with package available (may be absent)
//
// Modified      : is the problem modified
type ProblemObject struct {
	Id            int
	Owner         string
	Name          string
	Deleted       bool
	Favourite     bool
	AccessType    string
	Revision      int
	LatestPackage int
	Modified      bool
}

type wrapperProblemSlice struct {
	Status  string
	Comment string
	Result  []ProblemObject
}

// Prettify pretty prints a ProblemObject
func (problem *ProblemObject) Prettify() (prettyProblem string, err error) {
	problemJSON, err := json.MarshalIndent(problem, "", " ")
	if err != nil {
		return prettyProblem, err
	}
	return string(problemJSON), err
}

// ProblemInfoObject represents problem’s general information
//
// Parameter Description
//
// InputFile   : problem’s input file
//
// OutputFile  : problem’s output file
//
// Interactive : is problem interactive
//
// TimeLimit   : problem’s time limit in milliseconds
//
// MemoryLimit : problem’s memory limit in MB
type ProblemInfoObject struct {
	InputFile   string
	OutputFile  string
	Interactive bool
	TimeLimit   int
	MemoryLimit int
}

type wrapperProblemInfo struct {
	Status  string
	Comment string
	Result  ProblemInfoObject
}

// Prettify pretty prints a ProblemInfoObject
func (problemInfo *ProblemInfoObject) Prettify() (prettyInfo string, err error) {
	infoJSON, err := json.MarshalIndent(problemInfo, "", " ")
	if err != nil {
		return prettyInfo, err
	}
	return string(infoJSON), err
}

// StatementObject represents a problem’s statement
//
// Parameter Description
//
// Encoding : statement’s encoding
//
// Name     : problem’s name in statement’s language
//
// Legend   : problem’s language
//
// Input    : problem’s input format
//
// Output   : problem’s output format
//
// Scoring  : problem’s scoring
//
// Notes    : statement notes
//
// Tutorial : problem’s tutorial
type StatementObject struct {
	Encoding string
	Name     string
	Legend   string
	Input    string
	Output   string
	Scoring  string
	Notes    string
	Tutorial string
}

type wrapperStatementMap struct {
	Status  string
	Comment string
	Result  map[string]StatementObject
}

// Prettify pretty prints a StatementObject
func (statement *StatementObject) Prettify() (prettyStatement string, err error) {
	statementJSON, err := json.MarshalIndent(statement, "", " ")
	if err != nil {
		return prettyStatement, err
	}
	return string(statementJSON), err
}

// FileObject represents a resource, source or aux file.
//
// Parameter Description
//
// Name                       : name
//
// ModificationTimeSeconds    : file’s modification time in unix format
//
// Length                     : file length
//
// SourceType                 : (present only for source files) - source file type
//
// ResourceAdvancedProperties : (may be absent) Problem resource files can have extra
// property called ResourceAdvancedProperties of type ResourceAdvancedProperties
//
// TODO
//
// Address name conflicts
type FileObject struct {
	Name                       string
	ModificationTimeSeconds    int64
	Length                     int64
	SourceType                 string
	ResourceAdvancedProperties ResourceAdvancedPropertiesObject
}

type wrapperFileSlice struct {
	Status  string
	Comment string
	Result  []FileObject
}

// Prettify pretty prints a fileObject
func (file *FileObject) Prettify() (prettyFile string, err error) {
	fileJSON, err := json.MarshalIndent(file, "", " ")
	if err != nil {
		return prettyFile, err
	}
	return string(fileJSON), err
}

// RsaObject represents a slice of resource, source and aux fileObjects
type RsaObject struct {
	ResourceFiles []FileObject
	SourceFiles   []FileObject
	AuxFiles      []FileObject
}

type wrapperRSA struct {
	Status  string
	Comment string
	Result  RsaObject
}

// ResourceAdvancedPropertiesObject represents special properties of resource files.
// Basically, they stand for compile- or run-time resources for specific file types and asset types.
// The most important application is IOI-style graders.
//
// 	Example: {"forTypes":"cpp.*","main":false,"stages":["COMPILE"],"assets":["SOLUTION"]}
//
// Parameter Description
//
// ForTypes : colon or semicolon separated list of file types this resource if applied,
// wildcards are supported (example: “cpp.*” or “pascal.*;java.11”)
//
// Main : currently reserved to be false,
//
// Stages : array of possible string values COMPILE or RUN, meaning the phase when the resource is applicable,
//
// Assets : array of possible string values VALIDATOR, INTERACTOR, CHECKER, SOLUTION,
// meaning the asset types the resource is applicable.
type ResourceAdvancedPropertiesObject struct {
	ForTypes string
	Main     bool
	Stages   []string
	Assets   []string
}

// SolutionObject represents a problem solution
//
// Parameter Description
//
// Name                    : solution name
//
// ModificationTimeSeconds : solution’s modification time in unix format
//
// Length                  : solution length
//
// SourceType              : source file type
//
// Tag                     : solution tag
type SolutionObject struct {
	Name                    string
	ModificationTimeSeconds int64
	Length                  int64
	SourceType              string
	Tag                     string
}

type wrapperSolutionSlice struct {
	Status  string
	Comment string
	Result  []SolutionObject
}

// Prettify pretty prints a SolutionObject
func (solution *SolutionObject) Prettify() (prettySolution string, err error) {
	solutionJSON, err := json.MarshalIndent(solution, "", " ")
	if err != nil {
		return prettySolution, err
	}
	return string(solutionJSON), err
}

// TestObject represents a test for the problem.
//
// Parameter Description
//
// Index                          : test index
//
// Manual                         : boolean - whether test is manual or generated
//
// Input                          : test input (absent for generated tests)
//
// Description                    : test description (may be absent)
//
// UseInStatements                : boolean - whether test is included in statements
//
// ScriptLine                     : script line for generating test (absent for manual tests)
//
// Groups                         : test group (may be absent)
//
// Points                         : test points (may be absent)
//
// InputForStatement              : input for statements (may be absent)
//
// OutputForStatement             : output for statements (may be absent)
//
// VerifyInputOutputForStatements : boolean - whether to verify input and output for statements (may be absent)
type TestObject struct {
	Index                          int
	Manual                         bool
	Input                          string
	Description                    string
	UseInStatements                bool
	ScriptLine                     string
	Groups                         string
	Points                         string
	InputForStatement              string
	OutputForStatement             string
	VerifyInputOutputForStatements bool
}

type wrapperTestSlice struct {
	Status  string
	Comment string
	Result  []TestObject
}

// Prettify pretty prints a TestObject
func (test *TestObject) Prettify() (prettyTest string, err error) {
	testJSON, err := json.MarshalIndent(test, "", " ")
	if err != nil {
		return prettyTest, err
	}
	return string(testJSON), err
}

// TestGroupObject represents a test group in testset
//
// Parameter Description
//
// Name           : test group name
//
// PointsPolicy   : test group points policy,
// 	COMPLETE_GROUP for the complete group points policy,
// 	EACH_TEST for the each test points policy
//
// FeedbackPolicy : test group feedback policy,
// 	COMPLETE for the complete feedback policy,
// 	ICPC for the first error feedback policy, POINTS for the only points feedback policy,
// 	NONE for no feedback policy.
//
// Dependencies   : list of group names from which this group depends on (may be empty)
type TestGroupObject struct {
	Name           string
	PointsPolicy   string
	FeedbackPolicy string
	Dependencies   string
}

type wrapperTestGroupSlice struct {
	Status  string
	Comment string
	Result  []TestGroupObject
}

// Prettify pretty prints a TestGroupObject
func (test *TestGroupObject) Prettify() (prettyTestGroup string, err error) {
	testGroupJSON, err := json.MarshalIndent(test, "", " ")
	if err != nil {
		return prettyTestGroup, err
	}
	return string(testGroupJSON), err
}

// PackageObject represents a package
//
// Parameter Description
//
// Id                  : package’s id
//
// Revision            : revision of the problem for which the package was created
//
// CreationTimeSeconds : package’s creation time in unix format
//
// State               : PENDING/RUNNING/READY/FAILED package’s state
//
// Comment             : comment for the package
type PackageObject struct {
	Id                  int64
	Revision            int
	CreationTimeSeconds int64
	State               string
	Comment             string
}

type wrapperPackageSlice struct {
	Status  string
	Comment string
	Result  []PackageObject
}

// Prettify pretty prints a PackageObject
func (packageObj *PackageObject) Prettify() (prettyPackage string, err error) {
	packageObjJSON, err := json.MarshalIndent(packageObj, "", " ")
	if err != nil {
		return prettyPackage, err
	}
	return string(packageObjJSON), err
}
