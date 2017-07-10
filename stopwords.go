package sego

import (
    "sync"
    "io/ioutil"
    "strings"
    "errors"
)

type StopWordsMap map[string]interface{}

// DefaultStopWordMap contains some stop words.
var DefaultStopWordsMap = StopWordsMap {
	"the":   nil,
	"of":    nil,
	"is":    nil,
	"and":   nil,
	"to":    nil,
	"in":    nil,
	"that":  nil,
	"we":    nil,
	"for":   nil,
	"an":    nil,
	"are":   nil,
	"by":    nil,
	"be":    nil,
	"as":    nil,
	"on":    nil,
	"with":  nil,
	"can":   nil,
	"if":    nil,
	"from":  nil,
	"which": nil,
	"you":   nil,
	"it":    nil,
	"this":  nil,
	"then":  nil,
	"at":    nil,
	"have":  nil,
	"all":   nil,
	"not":   nil,
	"one":   nil,
	"has":   nil,
	"or":    nil,
}

// StopWord is a thread-safe dictionary for all stop words.
type StopWords struct {
	stopWordsMap StopWordsMap
	sync.RWMutex
}

// Add a stop word into StopWords dictionary.
func (s *StopWords) Add(word string) {
	s.Lock()
	s.stopWordsMap[word] = nil
	s.Unlock()
}

// NewStopWord create a new StopWord. If no parameter provided, default stop words will be used. At most one stop word dictionary could be loaded.
func NewStopWords(filepath ...string) (s *StopWords, err error) {
	s = new(StopWords)

    switch len(filepath) {
        case 0:
            s.stopWordsMap = DefaultStopWordsMap
        case 1:
            s.stopWordsMap, err = loadStopWordsDictionary(filepath[0])
            if err != nil {
                return
            }
        default:
            err = errors.New("Only one stop words dictionary could be loaded.")
            return
    }

    return
}

// IsStopWord checks if a given word is stop word.
func (s *StopWords) IsStopWord(word string) bool {
	s.RLock()
	_, ok := s.stopWordsMap[word]
	s.RUnlock()

    if ok {
        return true
    }

    // 空字符串，1字节长度的字符串都视为stop word
    if n := len(word); n == 0 || n == 1 {
        return true
    }

	return false
}

// 加载stop word字典，每行为一个stop word，加载后，stopWordMap 被新的字典内容替代
func (s *StopWords) LoadDictionary(filepath string) error {

    m, err := loadStopWordsDictionary(filepath)
    if err != nil {
        return err
    }

    s.Lock()
    s.stopWordsMap = m
    s.Unlock()

    return nil
}

func loadStopWordsDictionary(filepath string) (wm StopWordsMap, err error) {

    b, err := ioutil.ReadFile(filepath)
    if err != nil {
        return
    }

    wm = make(StopWordsMap)

    for _, line := range strings.Split(string(b), "\n") {
        line = strings.TrimSpace(line)
        if line != "" {
            wm[line] = nil
        }
    }

    return
}

// 从Segment里删除包含stop words的项目
func (s *StopWords) RemoveStopWords(in []Segment) (out []Segment) {

    for _, seg := range in {
        if !s.IsStopWord(seg.Token().Text()) {
            out = append(out, seg)
        }
    }

    return
}

