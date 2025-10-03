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

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestBoundaries(t *testing.T) {
	cases := []struct {
		name   string
		assign int
		exam   int
		essay  int
		want   string
	}{
		{"A_at_90", 90, 90, 90, "A"},
		{"B_at_80", 80, 80, 80, "B"},
		{"C_at_70", 70, 70, 70, "C"},
		{"D_at_60", 60, 60, 60, "D"},
		{"F_below_60", 59, 59, 59, "F"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gc := NewGradeCalculator()
			gc.AddGrade("a1", tc.assign, Assignment)
			gc.AddGrade("e1", tc.exam, Exam)
			gc.AddGrade("s1", tc.essay, Essay)
			if got := gc.GetFinalGrade(); got != tc.want {
				t.Fatalf("want %s, got %s", tc.want, got)
			}
		})
	}
}

func TestBucketRouting_MixedAverages(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 70, Assignment)
	gc.AddGrade("a2", 80, Assignment)
	gc.AddGrade("x1", 75, Exam)
	gc.AddGrade("s1", 72, Essay)
	if got := gc.GetFinalGrade(); got != "C" {
		t.Fatalf("expected C, got %s", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatal("bad string; Assignment")
	}
	if Exam.String() != "exam" {
		t.Fatal("bad string; Exam")
	}
	if Essay.String() != "essay" {
		t.Fatal("bad string; Essay")
	}
}

func TestComputeAverage_EmptyCategoryIsZero(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 100, Assignment)
	if got := gc.GetFinalGrade(); got != "D" {
		t.Fatalf("expected D when two categories empty, got %s", got)
	}
}
