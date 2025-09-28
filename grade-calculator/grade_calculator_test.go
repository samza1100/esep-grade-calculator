package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("a1", 100, Assignment) 
	g.AddGrade("e1", 58, Exam)     
	// no essay                    
	if got := g.GetFinalGrade(); got != "C" {
		t.Fatalf("expected C, got %s", got)
	}
}

func TestGetGradeD(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("a1", 100, Assignment) 
	g.AddGrade("e1", 43, Exam) 
	// no essay     
	if got := g.GetFinalGrade(); got != "D" {
		t.Fatalf("expected D, got %s", got)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 30, Assignment)
	gradeCalculator.AddGrade("exam 1", 40, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 20, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment.String wrong: %q", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam.String wrong: %q", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay.String wrong: %q", Essay.String())
	}
}