package utils

import "regexp"


func ExtractHashtags(content string) []string {
    hashtagRegex := regexp.MustCompile(`#([\p{L}\p{N}_]+)`)
    matches := hashtagRegex.FindAllStringSubmatch(content, -1)

    var hashtags []string
    for _, match := range matches {
        if len(match) > 1 {
            hashtags = append(hashtags, match[1])
        }
    }

    return hashtags
}
