// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/cockroachdb/errors"
)

type target string

func (t target) query() *template.Template {
	return queryTemplates.Lookup(string(t))
}

func (t target) execQuery(qd *queryData) (results []string, _ error) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		return nil, err
	}
	defer func() { _ = os.Remove(f.Name()) }()
	if err := t.query().Execute(f, qd); err != nil {
		return nil, err
	}
	if err := f.Close(); err != nil {
		return nil, err
	}
	cmd := exec.Command("bazel", "query", "--query_file", f.Name())
	var stdout bytes.Buffer
	var stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, errors.Wrapf(err,
			"failed to run %s: (stderr)\n%s", cmd, &stderr)
	}
	for sc := bufio.NewScanner(&stdout); sc.Scan(); {
		results = append(results, sc.Text())
	}
	sort.Strings(results)
	return results, nil
}

func (t target) filename() string {
	return string(t) + ".bzl"
}

func (t target) variable() string {
	return strings.ToUpper(string(t)) + "_SRCS"
}

func (t target) write(outDir string, out []string) error {
	var buf bytes.Buffer
	if err := outputVariableTemplate.Execute(&buf, variableData{
		Variable: t.variable(),
		Targets:  out,
	}); err != nil {
		return errors.Wrapf(err, "failed to execute template for %s", t)
	}
	f, err := os.Create(filepath.Join(outDir, t.filename()))
	if err != nil {
		return errors.Wrapf(err, "failed to open file for %s", t)
	}
	if _, err := io.Copy(f, &buf); err != nil {
		return errors.Wrapf(err, "failed to write file for %s", t)
	}
	if err := f.Close(); err != nil {
		return errors.Wrapf(err, "failed to write file for %s", t)
	}
	return nil
}

var outputVariableTemplate = template.Must(template.New("file").Parse(
	`# Generated by genbzl

{{ .Variable }} = [{{ range .Targets }}
    "{{ . }}",{{end}}
]
`))

type variableData struct {
	Variable string
	Targets  []string
}
