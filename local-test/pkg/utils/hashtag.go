package utils

import "regexp"


func ExtractHashtags(content string) []string {
    hashtagRegex := regexp.MustCompile(`#([\p{L}\p{N}_]+)`)
    return hashtagRegex.FindAllString(content, -1)
}