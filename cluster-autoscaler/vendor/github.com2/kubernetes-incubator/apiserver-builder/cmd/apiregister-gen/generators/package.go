/*
Copyright 2017 The Kubernetes Authors.

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

package generators

import (
	"path"
	"path/filepath"
	"strings"

	"k8s.io2/apimachinery/pkg/util/sets"
	"k8s.io2/gengo/args"
	"k8s.io2/gengo/generator"
	"k8s.io2/gengo/namer"
	"k8s.io2/gengo/types"

	"github.com2/pkg/errors"
)

// CustomArgs is used tby the go2idl framework to pass args specific to this
// generator.
type CustomArgs struct{}

type Gen struct {
	p []generator.Package
}

func (g *Gen) Execute(arguments *args.GeneratorArgs) error {
	return arguments.Execute(
		g.NameSystems(),
		g.DefaultNameSystem(),
		g.Packages)
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func (g *Gen) DefaultNameSystem() string {
	return "public"
}

// NameSystems returns the name system used by the generators in this package.
func (g *Gen) NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public": namer.NewPublicNamer(1),
		"raw":    namer.NewRawNamer("", nil),
	}
}

func (g *Gen) ParsePackages(context *generator.Context, arguments *args.GeneratorArgs) (sets.String, sets.String, string, string) {
	versionedPkgs := sets.NewString()
	unversionedPkgs := sets.NewString()
	mainPkg := ""
	apisPkg := ""
	for _, o := range context.Order {
		if IsAPIResource(o) {
			versioned := o.Name.Package
			versionedPkgs.Insert(versioned)
			unversioned := filepath.Dir(versioned)
			unversionedPkgs.Insert(unversioned)

			if apis := filepath.Dir(unversioned); apis != apisPkg && len(apisPkg) > 0 {
				panic(errors.Errorf(
					"Found multiple apis directory paths: %v and %v", apisPkg, apis))
			} else {
				apisPkg = apis
				mainPkg = filepath.Dir(apisPkg)
			}
		}
	}
	return versionedPkgs, unversionedPkgs, apisPkg, mainPkg
}

func (g *Gen) Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	g.p = generator.Packages{}

	b := NewAPIsBuilder(context, arguments)
	for _, apigroup := range b.APIs.Groups {
		for _, apiversion := range apigroup.Versions {
			factory := &packageFactory{apiversion.Pkg.Path, arguments}
			// Add generators for versioned types
			gen := CreateVersionedGenerator(apiversion, apigroup, arguments.OutputFileBaseName)
			g.p = append(g.p, factory.createPackage(gen))
		}

		factory := &packageFactory{apigroup.Pkg.Path, arguments}
		gen := CreateUnversionedGenerator(apigroup, arguments.OutputFileBaseName)
		g.p = append(g.p, factory.createPackage(gen))

		factory = &packageFactory{path.Join(apigroup.Pkg.Path, "install"), arguments}
		gen = CreateInstallGenerator(apigroup, arguments.OutputFileBaseName)
		g.p = append(g.p, factory.createPackage(gen))
	}

	factory := &packageFactory{b.APIs.Pkg.Path, arguments}
	gen := CreateApisGenerator(b.APIs, arguments.OutputFileBaseName)
	g.p = append(g.p, factory.createPackage(gen))

	// Add generators for Controllers.
	repo := ""
	for _, c := range b.Controllers {
		repo = c.Repo
		factory = &packageFactory{c.Pkg.Path, arguments}
		cgen := CreateControllerGenerator(c, arguments.OutputFileBaseName)
		g.p = append(g.p, factory.createPackage(cgen))
	}

	if len(b.Controllers) > 0 {
		factory = &packageFactory{context.Universe[repo+"/pkg/controller"].Path, arguments}
		cgen := CreateAllControllerGenerator(b.Controllers, arguments.OutputFileBaseName)
		g.p = append(g.p, factory.createPackage(cgen))
	}

	if len(b.Controllers) > 0 {
		factory = &packageFactory{context.Universe[repo+"/pkg/controller/sharedinformers"].Path, arguments}
		cgen := CreateInformersGenerator(b.Controllers, arguments.OutputFileBaseName)
		g.p = append(g.p, factory.createPackage(cgen))
	}
	return g.p
}

type packageFactory struct {
	path      string
	arguments *args.GeneratorArgs
}

// Creates a package with a generator
func (f *packageFactory) createPackage(gen generator.Generator) generator.Package {
	path := f.path
	name := strings.Split(filepath.Base(f.path), ".")[0]
	return &generator.DefaultPackage{
		PackageName: name,
		PackagePath: path,
		HeaderText:  f.getHeader(),
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			return []generator.Generator{gen}
		},
		FilterFunc: func(c *generator.Context, t *types.Type) bool {
			return t.Name.Package == f.path
		},
	}
}

// Returns the header for generated files
func (f *packageFactory) getHeader() []byte {
	header := []byte(`/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

`)
	return header
}
