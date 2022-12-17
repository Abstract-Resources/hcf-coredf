package chat

import "strings"

//goland:noinspection GoSnakeCaseUsage
var (
	YOU_ALREADY_IN_FACTION = TextReplacer{
		key:       "YOU_ALREADY_IN_FACTION",
		arguments: nil,
	}

	FACTION_ALREADY_EXISTS = TextReplacer{
		key:       "FACTION_ALREADY_EXISTS",
		arguments: []string{"faction"},
	}
)

type TextReplacer struct {
	key string
	arguments []string
}

func (t TextReplacer) Build(args ...string) string {
	replaced := map[string]string{}

	for i := 0; i < len(args); i++ {
		replaced[t.arguments[i]] = args[i]
	}

	return ReplacePlaceHolders(t.key, replaced)
}

func ReplacePlaceHolders(key string, args map[string]string) string {
	return ""
}

func Colour(text string) string {
	return strings.Replace(text, "&", "§", -1)
}