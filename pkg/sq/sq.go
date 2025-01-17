package sq

import (
	. "github.com/Masterminds/squirrel"
)

var Psql StatementBuilderType = StatementBuilder.PlaceholderFormat(Dollar)
