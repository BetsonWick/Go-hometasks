package learning

import (
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetTutorsIDPreferIndividual(t *testing.T) {
	t.Parallel()
	controller := minimock.NewController(t)

	tests := []struct {
		testName         string
		studentInfo      *studentInfo
		hasStudentInfo   bool
		individualTutors []int64
		groupTutors      []int64
		expected         []int64
		expectedOk       bool
		studentId        int64
		subject          string
	}{
		{
			testName:         "No student info",
			studentInfo:      nil,
			hasStudentInfo:   false,
			individualTutors: nil,
			groupTutors:      nil,
			expected:         nil,
			expectedOk:       false,
			studentId:        1,
			subject:          "Mathematics",
		},
		{
			testName:         "Student with preferred individual tutors",
			studentInfo:      &studentInfo{"Alex", 21, "Mathematics"},
			hasStudentInfo:   true,
			individualTutors: []int64{1},
			groupTutors:      nil,
			expected:         []int64{1},
			expectedOk:       true,
			studentId:        1,
			subject:          "Mathematics",
		},
		{
			testName:         "Student with preferred individual tutors having group tutors",
			studentInfo:      &studentInfo{"Alex", 21, "Mathematics"},
			hasStudentInfo:   true,
			individualTutors: []int64{1},
			groupTutors:      []int64{2},
			expected:         []int64{1},
			expectedOk:       true,
			studentId:        1,
			subject:          "Mathematics",
		},
		{
			testName:         "Student with preferred group tutors",
			studentInfo:      &studentInfo{"Alex", 21, "Mathematics"},
			hasStudentInfo:   true,
			individualTutors: nil,
			groupTutors:      []int64{2},
			expected:         []int64{2},
			expectedOk:       true,
			studentId:        1,
			subject:          "Mathematics",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			repositoryMock := NewRepositoryApiMock(controller)
			groupTutorServiceMock := NewTutorServiceApiMock(controller)
			individualTutorServiceMock := NewTutorServiceApiMock(controller)

			serviceMock := NewService(individualTutorServiceMock, groupTutorServiceMock, repositoryMock)

			repositoryMock.GetStudentInfoMock.Expect(test.studentId).Return(test.studentInfo, test.hasStudentInfo)

			if test.hasStudentInfo {
				individualTutorServiceMock.TutorsIDMock.Expect(test.subject).Return(test.individualTutors)

				if test.individualTutors == nil && len(test.individualTutors) == 0 {
					groupTutorServiceMock.TutorsIDMock.Expect(test.subject).Return(test.groupTutors)
				}
			}

			var actual, ok = serviceMock.GetTutorsIDPreferIndividual(test.studentId)

			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedOk, ok)
		})
	}

	t.Cleanup(controller.Finish)
}

func Test_GetTopSubjects(t *testing.T) {
	t.Parallel()
	controller := minimock.NewController(t)

	tests := []struct {
		testName       string
		subjectInfo    []subjectInfo
		hasSubjectInfo bool
		expected       []string
		expectedOk     bool
		topN           int
	}{
		{
			testName:       "Test empty subject info",
			subjectInfo:    []subjectInfo{},
			hasSubjectInfo: false,
			expected:       nil,
			expectedOk:     false,
			topN:           0,
		},
		{
			testName:       "Test one subject",
			subjectInfo:    []subjectInfo{{"Mathematics", 1}},
			hasSubjectInfo: true,
			expected:       []string{"Mathematics"},
			expectedOk:     true,
			topN:           1,
		},
		{
			testName:       "Test out of bounds",
			subjectInfo:    []subjectInfo{{"Mathematics", 1}},
			hasSubjectInfo: true,
			expected:       nil,
			expectedOk:     false,
			topN:           2,
		},
		{
			testName: "Test many subjects one not printed",
			subjectInfo: []subjectInfo{
				{"Physics", 10},
				{"Mathematics", 1},
				{"Programming", 2},
				{"Soft skills", 0},
			},
			hasSubjectInfo: true,
			expected:       []string{"Soft skills", "Mathematics", "Programming"},
			expectedOk:     true,
			topN:           3,
		},
		{
			testName: "Test many subjects all printed",
			subjectInfo: []subjectInfo{
				{"Physics", 10},
				{"Mathematics", 1},
				{"Programming", 2},
				{"Soft skills", 0},
			},
			hasSubjectInfo: true,
			expected:       []string{"Soft skills", "Mathematics", "Programming", "Physics"},
			expectedOk:     true,
			topN:           4,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			repositoryMock := NewRepositoryApiMock(controller)
			serviceMock := NewService(nil, nil, repositoryMock)

			repositoryMock.GetAllSubjectsInfoMock.Expect().Return(test.subjectInfo, test.hasSubjectInfo)

			var actual, ok = serviceMock.GetTopSubjects(test.topN)

			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedOk, ok)
		})
	}

	t.Cleanup(controller.Finish)
}
