package ext

import "strings"

type String string

func (_this String) _self() string {
	return string(_this)
}
func (_this String) String() string {
	return _this._self()
}
func (_this String) TrimSpace() String {
	return String(strings.TrimSpace(_this._self()))
}

func (_this String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(_this._self(), prefix)
}

func (_this String) HasPrefixIgnoreCase(prefix string) bool {
	return strings.HasPrefix(_this.ToLower(), strings.ToLower(prefix))
}

func (_this String) ISEmpty() bool {
	return len(_this) == 0
}
func (_this String) ToLower() string {
	return strings.ToLower(_this._self())
}
func (_this String) ToUpper() string {
	return strings.ToUpper(_this._self())
}
func (_this String) ReplaceAll(old, new string) string {
	return strings.ReplaceAll(_this._self(), old, new)
}

func (_this String) ReplaceAll_(old, new string) String {
	return String(strings.ReplaceAll(_this._self(), old, new))
}

func (_this String) ContainsIgnoreCase(searchStr string) bool {
	return strings.Contains(_this.ToLower(), strings.ToLower(searchStr))
}
func (_this String) Contains(searchStr string) bool {
	return strings.Contains(_this._self(), searchStr)
}
