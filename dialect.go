package sqlDialects

import (
	"errors"
	"fmt"
	"github.com/huhusen/sqlDialects/_internal/ext"
)

type DialectKind string

func (_this DialectKind) String() string {
	return string(_this)
}

const (
	MysqlDialect   = DialectKind("Mysql")
	OracleDialect  = DialectKind("Oracle")
	PostgreDialect = DialectKind("Postgre")
	NotSupport     = DialectKind("NOT_SUPPORT")
)

type _Dialect struct {
	kind             DialectKind
	sqlTemplate      string
	topLimitTemplate string
}

func Mysql() *_Dialect {
	return NewDialect(MysqlDialect)
}
func Oracle() *_Dialect {
	return NewDialect(OracleDialect)
}
func Postgre() *_Dialect {
	return NewDialect(PostgreDialect)
}
func NewDialect(kind DialectKind) *_Dialect {
	d := &_Dialect{kind: kind}
	initPageTemplate(d)
	return d
}
func (_this *_Dialect) Page(pageNumber, pageSize int, sql ext.String) (string, error) {
	trimedSql := sql.TrimSpace()
	if trimedSql.ISEmpty() {
		return "", ErrorStringEmpty
	}
	if !trimedSql.HasPrefix("select ") {
		return "", errors.New("SQL should start with \"select \"")
	}
	body := trimedSql[7:].TrimSpace()
	if body.ISEmpty() {
		return "", ErrorStringEmpty
	}
	skipRows := (pageNumber - 1) * pageSize
	totalRows := pageNumber * pageSize
	var useTemplate ext.String
	var result ext.String
	if skipRows == 0 {
		useTemplate = ext.String(_this.topLimitTemplate)
	} else {
		useTemplate = ext.String(_this.sqlTemplate)
	}
	if NotSupport == DialectKind(useTemplate) {
		if NotSupport != DialectKind(_this.topLimitTemplate) {
			return "", errors.New("Dialect \"" + _this.kind.String() + "\" only support top limit SQL, for example: \"" +
				aTopLimitSqlExample(ext.String(_this.topLimitTemplate)) + "\"")
		}
		return "", errors.New("Dialect \"" + _this.kind.String() + "\" does not support physical pagination")
	}
	if useTemplate.Contains(DISTINCT_TAG) {
		if body.HasPrefixIgnoreCase("distinct ") {
			useTemplate = useTemplate.ReplaceAll_(DISTINCT_TAG, "")
		} else {
			useTemplate = useTemplate.ReplaceAll_(DISTINCT_TAG, "distinct")
			body = body[9:]
		}
	}
	result = useTemplate.ReplaceAll_(SKIP_ROWS, fmt.Sprintf("%v", skipRows))
	result = result.ReplaceAll_(PAGESIZE, fmt.Sprintf("%v", pageSize))
	result = result.ReplaceAll_(TOTAL_ROWS, fmt.Sprintf("%v", totalRows))

	// now insert the customer's real full SQL here
	result = result.ReplaceAll_("$SQL", trimedSql.String())

	result = result.ReplaceAll_("$BODY", body.String())
	return result.String(), nil
}
