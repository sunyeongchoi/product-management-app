package common

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var charMap = map[string][]string{
	"ㄱ": {"가", "까"},
	"ㄲ": {"까", "나"},
	"ㄴ": {"나", "다"},
	"ㄷ": {"다", "따"},
	"ㄸ": {"따", "라"},
	"ㄹ": {"라", "마"},
	"ㅁ": {"마", "바"},
	"ㅂ": {"바", "빠"},
	"ㅃ": {"빠", "사"},
	"ㅅ": {"사", "싸"},
	"ㅆ": {"싸", "아"},
	"ㅇ": {"아", "자"},
	"ㅈ": {"자", "짜"},
	"ㅉ": {"짜", "차"},
	"ㅊ": {"차", "카"},
	"ㅋ": {"카", "타"},
	"ㅌ": {"타", "파"},
	"ㅍ": {"파", "하"},
	"ㅎ": {"하", "힣"},
}

// IsConsonants 초성으로만 이루어졌는지 확인
func IsConsonants(query string) bool {
	for _, c := range query {
		if _, ok := charMap[string(c)]; !ok {
			return false
		}
	}
	return true
}

// GetWhereClause 초성찾기의 WHERE 구문 생성
func GetWhereClause(query string, column string) string {
	var clause strings.Builder

	for i, c := range query {
		// TODO: index값 제대로 계산
		charSize := utf8.RuneLen(c)
		rangeValues := charMap[string(c)]
		clause.WriteString(fmt.Sprintf("SUBSTRING(%s, %d, 1) >= '%s' AND SUBSTRING(%s, %d, 1) < '%s' AND ", column, (i/charSize)+1, rangeValues[0], column, (i/charSize)+1, rangeValues[1]))
	}

	result := clause.String()
	return result[:len(result)-5]
}
