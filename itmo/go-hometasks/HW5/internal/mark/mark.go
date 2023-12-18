package mark

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
)

type Student struct {
	Name string
	Mark int
}

type StudentsStatistic interface {
	SummaryByStudent(student string) (int, bool)     // default_value, false - если студента нет
	AverageByStudent(student string) (float32, bool) // default_value, false - если студента нет
	Students() []string
	Summary() int
	Median() int
	MostFrequent() int
}

type StudentsStatisticImpl struct {
	marks map[string][]int
}

func average(summary int, count int) float32 {
	return float32(math.Round(float64(summary)/float64(count)*100) / 100)
}

func (a *StudentsStatisticImpl) SummaryByStudent(student string) (int, bool) {
	result := 0
	if a.marks[student] == nil {
		return 0, false
	}

	for _, v := range a.marks[student] {
		result += v
	}
	return result, true
}

func (a *StudentsStatisticImpl) AverageByStudent(student string) (float32, bool) {
	summary, hasInfo := a.SummaryByStudent(student)
	if !hasInfo {
		return 0, false
	}
	return average(summary, len(a.marks[student])), true
}

func (a *StudentsStatisticImpl) Students() []string {
	result, summaryMap := a.SummariesByStudents()
	sort.SliceStable(result, func(i, j int) bool {
		return summaryMap[result[i]] > summaryMap[result[j]]
	})
	return result
}

func (a *StudentsStatisticImpl) Summary() int {
	_, summaryMap := a.SummariesByStudents()
	result := 0
	for _, value := range summaryMap {
		result += value
	}
	return result
}

func (a *StudentsStatisticImpl) Median() int {
	marks := a.GetAllMarks()
	if len(marks) == 0 {
		return 0
	}
	sort.Ints(marks)
	return marks[len(marks)/2]
}

func (a *StudentsStatisticImpl) MostFrequent() int {
	marksCounterArray := make([]int, 10)
	marks := a.GetAllMarks()
	if len(marks) == 0 {
		return 0
	}
	for _, value := range marks {
		marksCounterArray[value]++
	}
	maxMarksCount := 0
	result := 0
	for i, value := range marksCounterArray {
		if value >= maxMarksCount {
			maxMarksCount = marksCounterArray[i]
			result = i
		}
	}
	return result
}

func (a *StudentsStatisticImpl) SummariesByStudents() ([]string, map[string]int) {
	keys := make([]string, 0)
	summaryMap := make(map[string]int)
	for key := range a.marks {
		keys = append(keys, key)
		summary, _ := a.SummaryByStudent(key)
		summaryMap[key] = summary
	}
	return keys, summaryMap
}

func (a *StudentsStatisticImpl) GetAllMarks() []int {
	result := make([]int, 0)
	for _, value := range a.marks {
		result = append(result, value...)
	}
	return result
}

func parseTokens(reader io.Reader) (map[string][]string, error) {
	result := map[string][]string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "\t")
		if len(tokens) != 2 {
			continue
		}
		result[tokens[0]] = append(result[tokens[0]], tokens[1])
	}
	if scanner.Err() != nil {
		return result, scanner.Err()
	}
	return result, nil
}

func ReadStudentsStatistic(reader io.Reader) (StudentsStatistic, error) {
	statistics := StudentsStatisticImpl{map[string][]int{}}
	tokens, err := parseTokens(reader)
	if err != nil {
		return &statistics, err
	}
	for student, marksStr := range tokens {
		for _, markStr := range marksStr {
			var mark int
			_, _ = fmt.Sscanf(markStr, "%d", &mark)
			if mark >= 0 && mark <= 10 {
				statistics.marks[student] = append(statistics.marks[student], mark)
			}
		}
	}
	return &statistics, nil
}

func WriteStudentsStatistic(writer io.Writer, statistic StudentsStatistic) error {
	summary := fmt.Sprintf("%d\t%d\t%d", statistic.Summary(), statistic.Median(), statistic.MostFrequent())
	_, err := writer.Write([]byte(summary))
	if err != nil {
		return err
	}
	students := statistic.Students()
	for _, value := range students {
		_, err := writer.Write([]byte("\n"))
		if err != nil {
			return err
		}
		studentSummary, hasSummary := statistic.SummaryByStudent(value)
		studentAverage, _ := statistic.AverageByStudent(value)
		formatted := value
		if hasSummary {
			if studentAverage == float32(int(studentAverage)) {
				formatted = fmt.Sprintf("%s\t%d\t%d", value, studentSummary, int(studentAverage))
			} else {
				formatted = fmt.Sprintf("%s\t%d\t%.2f", value, studentSummary, studentAverage)
			}
		}
		_, err = writer.Write([]byte(formatted))
		if err != nil {
			return err
		}
	}
	return nil
}
