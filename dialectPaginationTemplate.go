package sqlDialects

import "github.com/huhusen/sqlDialects/_internal/ext"

const (
	DISTINCT_TAG = "($DISTINCT)"
	SKIP_ROWS    = "$SKIP_ROWS"
	PAGESIZE     = "$PAGESIZE"
	TOTAL_ROWS   = "$TOTAL_ROWS"
)

func initPageTemplate(d *_Dialect) {
	d.sqlTemplate = initializePageSQLTemplate(d.kind)
	d.topLimitTemplate = initializeTopLimitSqlTemplate(d.kind)
}
func initializePageSQLTemplate(kind DialectKind) string {
	switch kind {
	case MysqlDialect, PostgreDialect:
		return "select $BODY limit $SKIP_ROWS, $PAGESIZE"
	case OracleDialect:
		return "select * from ( select row_.*, rownum rownum_ from ( select $BODY ) row_ where rownum <= $TOTAL_ROWS) where rownum_ > $SKIP_ROWS"
	default:
		return NotSupport.String()
	}
}
func initializeTopLimitSqlTemplate(kind DialectKind) string {
	switch kind {
	case MysqlDialect, PostgreDialect:
		return "select $BODY limit $PAGESIZE"
	case OracleDialect:
		return "select * from ( select $BODY ) where rownum <= $PAGESIZE"
	default:
		return NotSupport.String()
	}
}
func aTopLimitSqlExample(template ext.String) string {
	template = template.ReplaceAll_("$SQL", "select * from users order by userid")
	template = template.ReplaceAll_("$BODY", "* from users order by userid")
	template = template.ReplaceAll_(" "+DISTINCT_TAG, "")
	template = template.ReplaceAll_(SKIP_ROWS, "0")
	template = template.ReplaceAll_(PAGESIZE, "10")
	template = template.ReplaceAll_(TOTAL_ROWS, "10")
	return template.String()
}
