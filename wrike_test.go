package wrike

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var wrike = NewWrike(os.Getenv("BEARER"), nil)

// 특정 스페이스의 Task 리스트 조회
func TestTasks(t *testing.T) {
	tasks := wrike.Tasks()

	assert.NotEqual(t, tasks, nil)
}

// 프로젝트 리스트 조회
func TestProjects(t *testing.T) {
	projects := wrike.Projects(nil)

	fmt.Println(len(projects.Data))

	assert.NotEqual(t, projects, nil)
}

// 특정 프로젝트 조회 (링크)
func TestProjectsByLink(t *testing.T) {
	projects := wrike.ProjectsByLink("https://www.wrike.com/open.htm?id=865199939", nil)

	fmt.Println(len(projects.Data))

	assert.NotEqual(t, projects, nil)
}

// ID로 프로젝트 조회
func TestProjectsByIds(t *testing.T) {
	projects := wrike.ProjectsByLink("https://www.wrike.com/open.htm?id=865199939", nil)

	projectsSearch := wrike.ProjectsByIds(projects.Data[0].ChildIds)

	fmt.Println(len(projectsSearch.Data))
	fmt.Printf("%+v\n", projectsSearch.Data)

	assert.NotEqual(t, projectsSearch, nil)
}

// "2022.03.SP1"로 특정 스프린트 하위 폴더 조회
func TestSprints(t *testing.T) {
	sprints := wrike.Sprints("2022.03.SP1")

	fmt.Println(len(sprints))
	for _, v := range sprints {
		fmt.Printf("%+v\n\n", v)
	}

	assert.NotEqual(t, sprints, nil)
}
