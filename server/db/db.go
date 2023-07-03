package db

import (
	"github.com/egublog/milestone/server/models"
)

var Milestones []models.Milestone

func init() {
	Milestones = []models.Milestone{
		{ID: 1, Title: "Milestone 1"},
		{ID: 2, Title: "Milestone 2"},
		{ID: 3, Title: "Milestone 3"},
	}
}

func GetDummyMilestones() []models.Milestone {
	return Milestones
}