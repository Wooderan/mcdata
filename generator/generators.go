package main

import (
	"fmt"
	"github.com/Wooderan/mcdata"
	"github.com/atombender/go-jsonschema/pkg/generator"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
// ErrVersionNotSupported = errors.New("version provided not supported")
)

func GenerateTypes(dest string) error {
	// Get embedded files
	schemasJsonPath := filepath.Join(mcdata.SubmodulePath, "schemas")
	schemasFiles, err := os.ReadDir(schemasJsonPath)
	if err != nil {
		return err
	}

	// Prepare output folder
	curDir, _ := os.Getwd()
	outpath := filepath.Join(curDir, dest, "types")
	if err := os.MkdirAll(outpath, 0777); err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	return walkAndGenerateTypes(schemasJsonPath, outpath, schemasFiles)
}

func walkAndGenerateTypes(srcRoot string, outRoot string, schemaFiles []fs.DirEntry) error {
	for _, file := range schemaFiles {
		srcPath := filepath.Join(srcRoot, file.Name())

		if file.IsDir() {
			dstPath := filepath.Join(outRoot, file.Name())

			// Create the new directory in the destination root
			if err := os.MkdirAll(dstPath, os.ModePerm); err != nil {
				return err
			}

			// Read the directory
			newFiles, err := os.ReadDir(srcPath)
			if err != nil {
				return err
			}

			// Recurse into the directory
			if err := walkAndGenerateTypes(srcPath, dstPath, newFiles); err != nil {
				return err
			}
		} else {
			basicConfig := generator.Config{
				SchemaMappings:     []generator.SchemaMapping{},
				ExtraImports:       false,
				DefaultPackageName: "types",
				DefaultOutputName:  "-",
				ResolveExtensions:  []string{".json", ".yaml"},
				YAMLExtensions:     []string{".yaml", ".yml"},
				Warner: func(message string) {
					log.Printf("[from warner] %s", message)
				},
				Tags:                []string{"json", "yaml", "mapstructure"},
				StructNameFromTitle: true,
			}

			g, err := generator.New(basicConfig)
			if err != nil {
				panic(err)
			}

			if err := g.DoFile(srcPath); err != nil {
				fmt.Printf("Failed to generate type from schema: %s\n", srcPath)
				continue
			}

			if len(g.Sources()) == 0 {
				panic("Expected sources to contain something")
			}

			for outputName, source := range g.Sources() {
				if outputName == "-" {
					outputName = strings.TrimSuffix(filepath.Base(file.Name()), filepath.Ext(file.Name()))
					if strings.HasSuffix(outputName, "_schema") {
						outputName = strings.TrimSuffix(outputName, "_schema")
					}
					outputName += ".go"
				}
				dstPath := filepath.Join(outRoot, outputName)

				err = os.WriteFile(dstPath, source, 0644)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
