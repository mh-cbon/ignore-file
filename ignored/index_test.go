package ignored

import (
  "testing"
)

func TestParseEmptyFile(t *testing.T) {
	s := Ignored{}
	err := s.Parse("")

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}
	if len(s.Rules) > 0 {
		t.Errorf("should len(s.Rules)==0, got len(s.Rules)==%d\n", len(s.Rules))
	}
}

func TestCommentsAndEmptyLines(t *testing.T) {
	content := `# comment

# comment
`
	s := Ignored{}
	err := s.Parse(content)

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}
	if len(s.Rules) > 0 {
		t.Errorf("should len(s.Rules)==0, got len(s.Rules)==%d\n", len(s.Rules))
	}
}

func TestParseFile(t *testing.T) {
	content := `# comment
/some
other/
file
pattern*file
`
	s := Ignored{}
	err := s.Parse(content)

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}
	if len(s.Rules) != 4 {
		t.Errorf("should len(s.Rules)='4', got len(s.Rules)=%d\n", len(s.Rules))
	}
	r := s.Rules[0].String()
	if r != "^/some(/.*)?$" {
		t.Errorf("should Rules[0]='^/some(/.*)?$', got Rules[0]=%q\n", r)
	}
	r1 := s.Rules[1].String()
	if r1 != "/other/(/.*)?$" {
		t.Errorf("should Rules[1]='/other/(/.*)?$', got Rules[1]=%q\n", r1)
	}
  r2 := s.Rules[2].String()
	if r2 != "/file(/.*)?$" {
		t.Errorf("should Rules[2]='/file(/.*)?$', got Rules[2]=%q\n", r2)
	}
	r3 := s.Rules[3].String()
	if r3 != "/pattern[^/]*file(/.*)?$" {
		t.Errorf("should Rules[3]='/pattern[^/]*file(/.*)?$', got Rules[3]=%q\n", r3)
	}
}

func TestComputeDirectory1(t *testing.T) {
	content := ``
	s := Ignored{}
	err := s.Parse(content)

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}

  results := s.ComputeDirectory("../fixtures/")
	if len(results) != 13 {
		t.Errorf("should len(results)='13', got results=%q\n", results)
	}
  if contains(results, "/dira/fileb") == false {
		t.Errorf("should results should contain '/dira/fileb', got results=%q\n", results)
  }
  if contains(results, "/dira") {
		t.Errorf("should results should not contain '/dira', got results=%q\n", results)
  }
  if contains(results, "/dirb") {
		t.Errorf("should results should not contain '/dirb', got results=%q\n", results)
  }
  if contains(results, "/dirb/other") {
		t.Errorf("should results should not contain '/dirb/other', got results=%q\n", results)
  }
}

func TestComputeDirectory2(t *testing.T) {
	content := `*`
	s := Ignored{}
	err := s.Parse(content)

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}

  results := s.ComputeDirectory("../fixtures/")
	if len(results) != 0 {
		t.Errorf("should len(results)='0', got len(results)=%d\n", len(results))
	}
}

func TestComputeDirectory3(t *testing.T) {
	content := `some`
	s := Ignored{}
	err := s.Parse(content)

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}

  results := s.ComputeDirectory("../fixtures/")

	if len(results) != 10 {
		t.Errorf("should len(results)='10', got results=%q\n", results)
	}

  if contains(results, "/some") {
		t.Errorf("should results should not contain '/some', got results=%q\n", results)
  }

  if contains(results, "/dira/some") {
		t.Errorf("should results should not contain '/dira/some', got results=%q\n", results)
  }
  if contains(results, "/dirb/other/some") {
		t.Errorf("should results should not contain '/dirb/other/some', got results=%q\n", results)
  }
}

func TestComputeDirectory4(t *testing.T) {
	content := `/some`
	s := Ignored{}
	err := s.Parse(content)

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}

  results := s.ComputeDirectory("../fixtures/")

	if len(results) != 12 {
		t.Errorf("should len(results)='12', got results=%q\n", results)
	}

  if contains(results, "/some") {
		t.Errorf("should results should not contain '/some', got results=%q\n", results)
  }

  if contains(results, "/dira/some")==false {
		t.Errorf("should results should contain '/dira/some', got results=%q\n", results)
  }

  if contains(results, "/dirb/other/some")==false {
		t.Errorf("should results should contain '/dirb/other/some', got results=%q\n", results)
  }

  if contains(results, "/someother")==false {
		t.Errorf("should results should contain '/someother', got results=%q\n", results)
  }
}

func TestComputeDirectory5(t *testing.T) {
	content := `/some*other`
	s := Ignored{}
	err := s.Parse(content)

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}

  results := s.ComputeDirectory("../fixtures/")

	if len(results) != 12 {
		t.Errorf("should len(results)='12', got results=%q\n", results)
	}

  if contains(results, "/someother") {
		t.Errorf("should results should not contain '/someother', got results=%q\n", results)
  }

  if contains(results, "/dira/some")==false {
		t.Errorf("should results should contain '/dira/some', got results=%q\n", results)
  }

  if contains(results, "/dirb/other/some")==false {
		t.Errorf("should results should contain '/dirb/other/some', got results=%q\n", results)
  }

  if contains(results, "/some")==false {
		t.Errorf("should results should contain '/some', got results=%q\n", results)
  }
}

func TestComputeDirectory6(t *testing.T) {
	content := `other`
	s := Ignored{}
	err := s.Parse(content)

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}

  results := s.ComputeDirectory("../fixtures/")

	if len(results) != 9 {
		t.Errorf("should len(results)='9', got results=%q\n", results)
	}

  if contains(results, "/other") {
		t.Errorf("should results should not contain '/other', got results=%q\n", results)
  }

  if contains(results, "/dira/other") {
		t.Errorf("should results should not contain '/dira/other', got results=%q\n", results)
  }

  if contains(results, "/dirb/other/some") {
		t.Errorf("should results should not contain '/dirb/other/some', got results=%q\n", results)
  }

  if contains(results, "/dirb/other/fileb") {
		t.Errorf("should results should not contain '/dirb/other/fileb', got results=%q\n", results)
  }

  if contains(results, "/othera")==false {
		t.Errorf("should results should contain '/othera', got results=%q\n", results)
  }

  if contains(results, "/otherfile")==false {
		t.Errorf("should results should contain '/otherfile', got results=%q\n", results)
  }

  if contains(results, "/someother")==false {
		t.Errorf("should results should contain '/someother', got results=%q\n", results)
  }

  if contains(results, "/dira/someother")==false {
		t.Errorf("should results should contain '/dira/someother', got results=%q\n", results)
  }
}

func TestComputeDirectory7(t *testing.T) {
	content := `*/*/some`
	s := Ignored{}
	err := s.Parse(content)

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}

  results := s.ComputeDirectory("../fixtures/")

	if len(results) != 12 {
		t.Errorf("should len(results)='12', got results=%q\n", results)
	}

  if contains(results, "/dirb/other/some") {
		t.Errorf("should results should not contain '/dirb/other/some', got results=%q\n", results)
  }

  if contains(results, "/some")==false {
		t.Errorf("should results should contain '/some', got results=%q\n", results)
  }
  if contains(results, "/someother")==false {
		t.Errorf("should results should contain '/someother', got results=%q\n", results)
  }
  if contains(results, "/dira/some")==false {
		t.Errorf("should results should contain '/dira/some', got results=%q\n", results)
  }
}

func TestComputeDirectory8(t *testing.T) {
	content := `/dirb`
	s := Ignored{}
	err := s.Parse(content)

	if err != nil {
		t.Errorf("should err=nil, got err=%q\n", err)
	}

  results := s.ComputeDirectory("../fixtures/")

	if len(results) != 10 {
		t.Errorf("should len(results)='10', got results=%q\n", results)
	}

  if contains(results, "/dirb/other/some") {
		t.Errorf("should results should not contain '/dirb/other/some', got results=%q\n", results)
  }

  if contains(results, "/dirb") {
		t.Errorf("should results should not contain '/dirb', got results=%q\n", results)
  }
  if contains(results, "/dirb/other/fileb") {
		t.Errorf("should results should not contain '/dirb/other/fileb', got results=%q\n", results)
  }
  if contains(results, "/dirb/fileb") {
		t.Errorf("should results should not contain '/dirb/fileb', got results=%q\n", results)
  }
}

func contains (s []string, l string) bool {
  for _, v := range s {
    if l==v {
      return true
    }
  }
  return false
}
