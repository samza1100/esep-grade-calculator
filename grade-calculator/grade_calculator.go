package esepunittests

type GradeCalculator struct {
	grades []Grade
	mode GradingMode
}

type GradeType int


const (
	Assignment GradeType = iota
	Exam
	Essay
)

type GradingMode int

const (
	ModeLetter GradingMode = iota
	ModePassFail
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
		mode:   ModeLetter,
	}
}

func NewGradeCalculatorPassFail() *GradeCalculator {
	gc := NewGradeCalculator()
	gc.mode = ModePassFail
	return gc
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if gc.mode == ModePassFail {
		if numericalGrade >= 70 {
			return "Pass"
		}
		return "Fail"
	}

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	var assignments, exams, essays []Grade

	for _, g := range gc.grades {
		switch g.Type {
		case Assignment:
			assignments = append(assignments, g)
		case Exam:
			exams = append(exams, g)
		case Essay:
			essays = append(essays, g)
		}
	}

	assignment_average := computeAverage(assignments)
	exam_average := computeAverage(exams)
	essay_average := computeAverage(essays)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
        return 0
    }
	sum := 0
	for _, g := range grades {
		sum += g.Grade
	}

	return sum / len(grades)
}
