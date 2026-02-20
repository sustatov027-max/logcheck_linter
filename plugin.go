package main

import (
    "github.com/sustatov027-max/logcheck_linter/analyzer"
    "golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
    return []*analysis.Analyzer{analyzer.Analyzer}, nil
}