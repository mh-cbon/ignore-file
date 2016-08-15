package ignored

import (
  "regexp"
  "strings"
  "os"
  "io/ioutil"
  "path/filepath"
)

type Ignored struct {
  Rules []IgnoreRule
}

type IgnoreRule struct {
  HasStartAnchor bool // true if the rule begin with /
  IsPattern bool // true if the rule contains *
  Pattern *regexp.Regexp // the regexp pattern of the string rule
}
func (p *IgnoreRule) String() string {
  return p.Pattern.String()
}

func (p *Ignored) Load (filepath string) error {
  data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

  return p.Parse(string(data))
}

func (p *Ignored) Parse (content string) error {
  lines := strings.Split(content, "\n")
  for _, line := range lines {
    if err := p.Append(line); err!=nil {
      return err
    }
  }
  return nil
}

// Append a rule to the Ignored instance.
// If the rule is starting with a #, it is not added, it returns nil.
// If the rule is empty, it is not added, it returns nil.
// If the rule cannot be computed to a regexp, it returns an error.
// When the rule was added, it returns nil.
func (p *Ignored) Append (ruleString string) error {
  ruleString = strings.TrimSpace(ruleString)
  if len(ruleString)>0 {
    if ruleString[0:1]!="#" {
      rule := IgnoreRule{}
      rule.HasStartAnchor = ruleString[0:1]=="/"
      rule.IsPattern = strings.Index(ruleString, "*")==-1

      ruleString = regexp.QuoteMeta(ruleString)
      ruleString = strings.Replace(ruleString, "\\*", "[^/]*", -1)

      if rule.HasStartAnchor {
        ruleString = "^"+ruleString+"(/.*)?$"
      } else {
        if ruleString[0:1] != "/" {
          ruleString = "/" + ruleString
        }
        ruleString = ruleString + "(/.*)?$"
      }

      r, err := regexp.Compile(ruleString)
      if err!=nil {
        return err
      }
      rule.Pattern = r
      p.Rules = append(p.Rules, rule)
    }
  }
  return nil
}

func (p *Ignored) Match(some string) bool {
  for _, rule := range p.Rules {
    if rule.Pattern.MatchString(some) {
      return true
    }
  }
  return false
}

func (p *Ignored) ComputeDirectory(dirpath string) []string {
  ret := make([]string, 0)
  abs, err := filepath.Abs(dirpath)
  if err!=nil {
    return ret
  }
  s, err := os.Stat(abs)
  if err!=nil || !s.IsDir() {
    return ret
  }
  err = filepath.Walk(abs, func(path string, stat os.FileInfo, _ error) error {
    if stat.IsDir() {
      return nil
    }
    path = path[len(abs):]
    if p.Match(path)==false {
      if len(path)>0 {
        ret = append(ret, path)
      }
    }
    return nil
  })
  return ret
}
