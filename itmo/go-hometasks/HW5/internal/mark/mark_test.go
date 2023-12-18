package mark

import "testing"

func Test_SummaryByStudentDifferentSymbols_Sum_Skip(t *testing.T) {
	t.Parallel()
	TestCase_summaryByStudent{
		input: `Примеров Пример Примерович1	9
Примеров Пример Примерович2	7
Примеров Пример Примерович1	5
Примеров Пример Примерович1	5
Примеров Пример Примерович1
a3f32ffd asf12fs	9
FSsf asffsa	12
FSsf asffsa	9
\qw12
عباس	6
FSsf asffsa12
FSsf asffsa	1
Примеров Пример Примерович1	6
Примеров Пример Примерович2	2`,
		studentName:   "FSsf asffsa",
		isWantSummary: true,
		wantSummary:   10,
	}.Run(t)
}

func Test_AverageByStudentDifferentSymbols_Sum_Skip(t *testing.T) {
	t.Parallel()
	TestCase_averageByStudent{
		input: `FSsf asffsa	9
Примеров Пример Примерович1	9
Примеров Пример Примерович2	7
Примеров Пример Примерович1	5
Примеров Пример Примерович1	5
Примеров Пример Примерович1
a3f32ffd asf12fs	9
FSsf asffsa	12
FSsf asffsa	9
\qw12
عباس	6
FSsf asffsa12
FSsf asffsa	1
Примеров Пример Примерович1	6
Примеров Пример Примерович2	2`,
		studentName:   "FSsf asffsa",
		isWantAverage: true,
		wantAverage:   6.33,
	}.Run(t)
}

func Test_AverageByStudentDifferentSymbols_Sum_Skip_NoStudent(t *testing.T) {
	t.Parallel()
	TestCase_averageByStudent{
		input: `FSsf asffsa	9
Примеров Пример Примерович1	9
Примеров Пример Примерович2	7
Примеров Пример Примерович1
a3f32ffd asf12fs	9
FSsf asffsa	12
FSsf asffsa	9
\qw12
عباس	6
FSsf asffsa12
FSsf asffsa	1
Примеров Пример Примерович1	6
Примеров Пример Примерович2	2`,
		studentName:   "\\qw12",
		isWantAverage: false,
		wantAverage:   0,
	}.Run(t)
}

func Test_StudentsDifferentSymbols_IgnoreBadLines(t *testing.T) {
	t.Parallel()
	TestCase_students{
		input: `FSsf asffsa	9
Примеров Пример Примерович1	9
Примеров Пример Примерович2	7
Примеров Пример Примерович1	5
Примеров Пример Примерович1	5
Примеров Пример Примерович1
a3f32ffd asf12fs	8
FSsf asffsa	12
FSsf asffsa	9
\qw12
عباس	6
FSsf asffsa12
FSsf asffsa	1
Примеров Пример Примерович1	6
Примеров Пример Примерович2	2`,
		wantStudents: []string{
			"Примеров Пример Примерович1",
			"FSsf asffsa",
			"Примеров Пример Примерович2",
			"a3f32ffd asf12fs",
			"عباس",
		},
	}.Run(t)
}

func Test_StudentsDifferentSymbols_AllSkipped(t *testing.T) {
	t.Parallel()
	TestCase_students{
		input: `FSsf asffsa9
Примеров Пример Примерович1
a3f32ffd asf12f9
FSsf asffsa12
FSsf asffsa9
\qw12
عباس6
FSsf asffsa12
Примеров Пример Примерович22`,
		wantStudents: []string{},
	}.Run(t)
}

func Test_Summary_SomeStudentsDifferentSymbols_AllSkipped(t *testing.T) {
	t.Parallel()
	TestCase_summary{
		input: `FSsf asffsa9
Примеров Пример Примерович1
a3f32ffd asf12f9
FSsf asffsa12
FSsf asffsa9
\qw12
عباس6
FSsf asffsa12
Примеров Пример Примерович22`,
		wantSummary: 0,
	}.Run(t)
}

func Test_Summary_SomeStudentsDifferentSymbols(t *testing.T) {
	t.Parallel()
	TestCase_summary{
		input: `FSsf asffsa	9
Примеров Пример Примерович1	9
Примеров Пример Примерович2	7
Примеров Пример Примерович1	5
Примеров Пример Примерович1	5
Примеров Пример Примерович1
a3f32ffd asf12fs	9
FSsf asffsa	12
FSsf asffsa	9
\qw12
عباس	6
FSsf asffsa12
FSsf asffsa	1
Примеров Пример Примерович1	6
Примеров Пример Примерович2	2`,
		wantSummary: 68,
	}.Run(t)
}

func Test_Median_SomeStudentsDifferentSymbols_AllSkipped(t *testing.T) {
	t.Parallel()
	TestCase_median{
		input: `FSsf asffsa9
Примеров Пример Примерович1
a3f32ffd asf12f9
FSsf asffsa12
FSsf asffsa9
\qw12
عباس6
FSsf asffsa12
Примеров Пример Примерович22`,
		wantMedian: 0,
	}.Run(t)
}

func Test_Median_SomeStudentsDifferentSymbols(t *testing.T) {
	t.Parallel()
	TestCase_median{
		input: `FSsf asffsa	9
Примеров Пример Примерович1	9
Примеров Пример Примерович2	7
Примеров Пример Примерович1	5
Примеров Пример Примерович1	5
Примеров Пример Примерович1
a3f32ffd asf12fs	9
FSsf asffsa	12
FSsf asffsa	9
\qw12
عباس	6
FSsf asffsa12
FSsf asffsa	1
Примеров Пример Примерович1	6
Примеров Пример Примерович2	2`,
		wantMedian: 6,
	}.Run(t)
}

func Test_MostFrequentDifferentSymbols_AllSkipped(t *testing.T) {
	t.Parallel()
	TestCase_mostFrequent{
		input: `FSsf asffsa9
Примеров Пример Примерович1
a3f32ffd asf12f9
FSsf asffsa12
FSsf asffsa9
\qw12
عباس6
FSsf asffsa12
Примеров Пример Примерович22`,
		wantMostFrequent: 0,
	}.Run(t)
}

func Test_MostFrequentDifferentSymbols_AllEqual(t *testing.T) {
	t.Parallel()
	TestCase_mostFrequent{
		input: `FSsf asffsa	9
Примеров Пример Примерович1	9
Примеров Пример Примерович2	9
Примеров Пример Примерович1	9
Примеров Пример Примерович1	9
Примеров Пример Примерович1
a3f32ffd asf12fs	9
FSsf asffsa	9
FSsf asffsa	9
\qw10
عباس	9
FSsf asffsa12
FSsf asffsa	9
Примеров Пример Примерович1	9
Примеров Пример Примерович2	9`,
		wantMostFrequent: 9,
	}.Run(t)
}
