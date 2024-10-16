package internal

import (
	"go/ast"
	"go/token"

	"github.com/gnoswap-labs/tlin/internal/lints"
	tt "github.com/gnoswap-labs/tlin/internal/types"
)

/*
* Implement each lint rule as a separate struct
 */

// LintRule defines the interface for all lint rules.
type LintRule interface {
	// Check runs the lint rule on the given file and returns a slice of Issues.
	Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error)

	// Name returns the name of the lint rule.
	Name() string
}

type GolangciLintRule struct{}

func (r *GolangciLintRule) Check(filename string, _ *ast.File, _ *token.FileSet) ([]tt.Issue, error) {
	return lints.RunGolangciLint(filename)
}

func (r *GolangciLintRule) Name() string {
	return "golangci-lint"
}

type SimplifySliceExprRule struct{}

func (r *SimplifySliceExprRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectUnnecessarySliceLength(filename, node, fset)
}

func (r *SimplifySliceExprRule) Name() string {
	return "simplify-slice-range"
}

type UnnecessaryConversionRule struct{}

func (r *UnnecessaryConversionRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectUnnecessaryConversions(filename, node, fset)
}

func (r *UnnecessaryConversionRule) Name() string {
	return "unnecessary-type-conversion"
}

type LoopAllocationRule struct{}

func (r *LoopAllocationRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectLoopAllocation(filename, node, fset)
}

func (r *LoopAllocationRule) Name() string {
	return "loop-allocation"
}

type DetectCycleRule struct{}

func (r *DetectCycleRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectCycle(filename, node, fset)
}

func (r *DetectCycleRule) Name() string {
	return "cycle-detection"
}

type EmitFormatRule struct{}

func (r *EmitFormatRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectEmitFormat(filename, node, fset)
}

func (r *EmitFormatRule) Name() string {
	return "emit-format"
}

type SliceBoundCheckRule struct{}

func (r *SliceBoundCheckRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectSliceBoundCheck(filename, node, fset)
}

func (r *SliceBoundCheckRule) Name() string {
	return "slice-bounds-check"
}

type UselessBreakRule struct{}

func (r *UselessBreakRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectUselessBreak(filename, node, fset)
}

func (r *UselessBreakRule) Name() string {
	return "useless-break"
}

type EarlyReturnOpportunityRule struct{}

func (r *EarlyReturnOpportunityRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectEarlyReturnOpportunities(filename, node, fset)
}

func (r *EarlyReturnOpportunityRule) Name() string {
	return "early-return-opportunity"
}

type DeferRule struct{}

func (r *DeferRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectDeferIssues(filename, node, fset)
}

func (r *DeferRule) Name() string {
	return "defer-issues"
}

type MissingModPackageRule struct{}

func (r *MissingModPackageRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectMissingModPackage(filename, node, fset)
}

func (r *MissingModPackageRule) Name() string {
	return "gno-mod-tidy"
}

// -----------------------------------------------------------------------------
// Regex related rules

type RepeatedRegexCompilationRule struct{}

func (r *RepeatedRegexCompilationRule) Check(filename string, node *ast.File, fset *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectRepeatedRegexCompilation(filename)
}

func (r *RepeatedRegexCompilationRule) Name() string {
	return "repeated-regex-compilation"
}

// -----------------------------------------------------------------------------

type CyclomaticComplexityRule struct {
	Threshold int
}

func (r *CyclomaticComplexityRule) Check(filename string, node *ast.File) ([]tt.Issue, error) {
	return lints.DetectHighCyclomaticComplexity(filename, r.Threshold)
}

func (r *CyclomaticComplexityRule) Name() string {
	return "high-cyclomatic-complexity"
}

// -----------------------------------------------------------------------------

// GnoSpecificRule checks for gno-specific package imports. (p, r and std)
type GnoSpecificRule struct{}

func (r *GnoSpecificRule) Check(filename string, _ *ast.File, _ *token.FileSet) ([]tt.Issue, error) {
	return lints.DetectGnoPackageImports(filename)
}

func (r *GnoSpecificRule) Name() string {
	return "unused-package"
}
