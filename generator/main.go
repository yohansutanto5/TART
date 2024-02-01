package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type ModelConfig struct {
	Model  string `yaml:"model"`
	Fields []struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
	} `yaml:"fields"`
}

func main() {
	// Read YAML file
	yamlFile, err := ioutil.ReadFile("animal.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Unmarshal YAML data into struct
	var config ModelConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	// Generate Go file content
	goModelContent := generateGoFile(config)
	// Write to a new Go file
	err = ioutil.WriteFile("../model/generated.go", []byte(goModelContent), 0644)
	if err != nil {
		log.Fatalf("Error writing Go file: %v", err)
	}
	goDbContent := generateGoDbFile(config)
	// Write to a new Go file
	err = ioutil.WriteFile("../db/generated.go", []byte(goDbContent), 0644)
	if err != nil {
		log.Fatalf("Error writing Go file: %v", err)
	}
	fmt.Println("Generated Go file successfully.")
}

func generateGoFile(config ModelConfig) string {
	goFileContent := fmt.Sprintf(`package model

type %s struct {
	ID   int    `+"`gorm:\"primaryKey;autoIncrement\"`"+`
`, config.Model)

	for _, field := range config.Fields {
		goFileContent += fmt.Sprintf("\t%s %s\n", field.Name, getTypeMapping(field.Type))
	}

	goFileContent += "}\n\n"

	// DTO input
	goFileContent += fmt.Sprintf(`type Add%sIn struct {
`, config.Model)

	for _, field := range config.Fields {
		goFileContent += fmt.Sprintf("\t%s %s `json:\"%s\" binding:\"required\"`\n", field.Name, getTypeMapping(field.Type), field.Name)
	}

	goFileContent += "}\n\n"

	// PopulateFromDTOInput method
	goFileContent += fmt.Sprintf(`func (m *%s) PopulateFromDTOInput(input Add%sIn) {
`, config.Model, config.Model)

	for _, field := range config.Fields {
		goFileContent += fmt.Sprintf("\tm.%s = input.%s\n", field.Name, field.Name)
	}

	goFileContent += "}\n"

	return goFileContent
}

func getTypeMapping(yamlType string) string {
	switch yamlType {
	case "string":
		return "string"
	case "bool":
		return "bool"
	// Add more type mappings as needed
	default:
		return "interface{}"
	}
}

func generateGoDbFile(config ModelConfig) string {
	packageName := strings.ToLower(config.Model)
	goFileContent := fmt.Sprintf(`package db

import (
	"app/model"
	"app/pkg/error"
)

func (d *DataStore) GetList%s() (%ss []model.%s, err *error.Error) {
	e := d.Db.Find(&%ss).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) Insert%s(%s *model.%s) (err *error.Error) {
	e := d.Db.Create(%s).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) Delete%sByID(id int) (err *error.Error) {
	e := d.Db.Where("id = ?", id).Delete(&model.%s{}).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) Update%s(%s *model.%s) (err *error.Error) {
	e := d.Db.Save(&%s).Error
	err.ParseMysqlError(e)
	return
}
`, config.Model, packageName, config.Model, packageName, config.Model, packageName, config.Model, packageName, config.Model, packageName, config.Model, packageName, config.Model, packageName)

	return goFileContent
}
