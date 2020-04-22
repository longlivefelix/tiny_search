/*
@Author: Felix <https://github.com/longlivefelix>
@Date:   2020-04-21 18:19
*/
package structs

import (
	"sort"
	"unicode/utf8"
)

const CONTENT_OFFSETS_SIZE=65530
type Offset uint16
type OffsetList struct {
	Count   uint16
	Offsets []Offset
}
type TestResult struct {
	c rune
	offset Offset
}
type FileIdx uint16                   // File serial number in block
type RuneField map[FileIdx]*OffsetList // Rune’s offsets
type RuneIndex map[rune]RuneField

func (index RuneIndex) BuildIndex(content []byte, fileIdx FileIdx) {
	wordIndex := 0
	for i := 0; i < len(content); {
		r, n := utf8.DecodeRune(content[i:])
		if r == 0 {
			break
		}
		if index[r] == nil{
			index[r] = make(RuneField,0)
		}
		if _,ok:= index[r][fileIdx];!ok{
			index[r][fileIdx] = &OffsetList{
				Count:0,
				Offsets:[]Offset{},
			}
		}
		if index[r][fileIdx].Count < CONTENT_OFFSETS_SIZE{
			field := index[r][fileIdx]
			field.Offsets = append(field.Offsets,Offset(wordIndex))
			field.Count++
		}
		wordIndex += 1
		i += n
	}
}

func (index RuneIndex)Search(content string) []FileIdx{
	// check all utf8char in this index
	for j:=0;j<len(content);{
		r,n := utf8.DecodeRuneInString(content[j:])
		j += n
		if _,ok:= index[r];!ok{
			return []FileIdx{}
		}
	}
	roughResult := map[rune]RuneField{}
	for j:=0;j<len(content);{
		r,n := utf8.DecodeRuneInString(content[j:])
		j += n
		if field,ok:= index[r];ok{
			if _,ok := roughResult[r]; !ok{
				roughResult[r] = field
			}
		}
	}
	startFileList := []FileIdx{}
	halfMatchFileList := map[FileIdx]bool{}
	fullMatchFileList := []FileIdx{}
	resultFileList := []FileIdx{}
	for _,field := range roughResult{
		for fileIdx,_:= range field{
			startFileList = append(startFileList, fileIdx)
		}
		break
	}
	for _,fileIdx := range startFileList{
		for _,field := range roughResult{
			if _,ok:=field[fileIdx]; !ok{
				halfMatchFileList[fileIdx] = true
			}
		}
	}
	for _,fileIdx := range startFileList{
		if _,ok := halfMatchFileList[fileIdx]; !ok{
			fullMatchFileList = append(fullMatchFileList,fileIdx)
		}
	}
	for _,fileIdx := range fullMatchFileList{
		testResults := []*TestResult{}
		for r,field := range roughResult{
			for _,offset := range field[fileIdx].Offsets{
				testResults = append(testResults, &TestResult{
					c:r,
					offset:offset,
				})
			}
		}
		sort.Slice(testResults, func(i, j int) bool {
			return testResults[i].offset < testResults[i].offset
		})
		preOffset := Offset(0)
		word := ""
		for _, result := range testResults{
			if preOffset==0 || Offset(preOffset+1) == result.offset{
				word+=string(result.c)
				preOffset = result.offset
				if word == content{
					resultFileList = append(resultFileList, fileIdx)
					break
				}
			}
		}
	}
	return resultFileList
}