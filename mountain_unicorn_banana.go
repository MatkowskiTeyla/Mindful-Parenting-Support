package main

import (
	"fmt"
)

// ParentCoachingAndSupportServices - the interface for Parent Coaching & Support Services
type ParentCoachingAndSupportServices interface {
	ProvideSupport()
	ProvideCoaching()
	ConnectResources()
	ProvideGuidance()
}

// ParentCoach - a type representing a Parent Coach
type ParentCoach struct {
	name   string
	skills []string
}

// ProvideSupport - provides support to the client
func (pc *ParentCoach) ProvideSupport() {
	fmt.Println("Providing support to client...")
}

// ProvideCoaching - provides coaching to the client
func (pc *ParentCoach) ProvideCoaching() {
	fmt.Println("Providing coaching to client...")
}

// ConnectResources - connects the client to the necessary resources
func (pc *ParentCoach) ConnectResources() {
	fmt.Println("Connecting client to necessary resources...")
}

// ProvideGuidance - provides guidance to the client
func (pc *ParentCoach) ProvideGuidance() {
	fmt.Println("Providing guidence to client...")
}

// ParentSupportService - a type representing a Parent Support Service
type ParentSupportService struct {
	name   string
	skills []string
}

// ProvideSupport - provides support to the client
func (pss *ParentSupportService) ProvideSupport() {
	fmt.Println("Providing support to client...")
}

// ConnectResources - connects the client to the necessary resources
func (pss *ParentSupportService) ConnectResources() {
	fmt.Println("Connecting client to necessary resources...")
}

// ProvideGuidance - provides guidance to the client
func (pss *ParentSupportService) ProvideGuidance() {
	fmt.Println("Providing guidence to client...")
}

// NewParentCoach - creates and returns a new ParentCoach instance
func NewParentCoach(name string, skills []string) *ParentCoach {
	return &ParentCoach{
		name:   name,
		skills: skills,
	}
}

// NewParentSupportService - creates and returns a new ParentSupportService instance
func NewParentSupportService(name string, skills []string) *ParentSupportService {
	return &ParentSupportService{
		name:   name,
		skills: skills,
	}
}

// CreateParentCoachService - creates and returns a new ParentCoachingAndSupportServices instance
func CreateParentCoachService(name string, skills []string) ParentCoachingAndSupportServices {
	if name == "Parent Coach" {
		return NewParentCoach(name, skills)
	} else if name == "Parent Support Service" {
		return NewParentSupportService(name, skills)
	}

	return nil
}

func main() {
	// Create a new Parent Coach
	parentCoach := CreateParentCoachService("Parent Coach", []string{"Listening", "Advising", "Talking"})

	// Call the methods
	parentCoach.ProvideSupport()
	parentCoach.ProvideCoaching()
	parentCoach.ConnectResources()
	parentCoach.ProvideGuidance()

	// Create a new Parent Support Service
	parentSupportService := CreateParentCoachService("Parent Support Service", []string{"Listening", "Advising", "Talking", "Referring"})

	// Call the methods
	parentSupportService.ProvideSupport()
	parentSupportService.ConnectResources()
	parentSupportService.ProvideGuidance()
}