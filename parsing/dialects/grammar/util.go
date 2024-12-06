package grammar

import (
	"github.com/karlmoad/go_util_lib/generics/result"
	"github.com/karlmoad/go_util_lib/parsing/ast"
)

func ToRule(elem ast.Element) *result.Result[RuleExpr] {
	if elem.Elem().Kind() == RULE_ELEM {
		if conv, ok := elem.(RuleExpr); ok {
			return result.NewResultWithValue(conv)
		}
	}
	return nil
}

func ToBody(elem ast.Element) *result.Result[BodyStmt] {
	if elem.Elem().Kind() == BODY_ELEM {
		if conv, ok := elem.(BodyStmt); ok {
			return result.NewResultWithValue(conv)
		}
	}
	return nil
}

func ToStringOrIdent(elem ast.Element) *result.Result[StringOrIdentifierStmt] {
	kind := elem.Elem().Kind()
	if kind == STRING_ELEM || kind == IDENTIFIER_ELEM {
		if conv, ok := elem.(StringOrIdentifierStmt); ok {
			return result.NewResultWithValue(conv)
		}
	}
	return nil
}

func ToSet(elem ast.Element) *result.Result[SetExpr] {
	if elem.Elem().Kind() == SET_ELEM {
		if conv, ok := elem.(SetExpr); ok {
			return result.NewResultWithValue(conv)
		}
	}
	return nil
}

func ToUnknown(elem ast.Element) *result.Result[UnknownStmt] {
	if elem.Elem().Kind() == UNK_ELEM {
		if conv, ok := elem.(UnknownStmt); ok {
			return result.NewResultWithValue(conv)
		}
	}
	return nil
}

func ToAlt(elem ast.Element) *result.Result[AlternativeExpr] {
	if elem.Elem().Kind() == ALT_ELEM {
		if conv, ok := elem.(AlternativeExpr); ok {
			return result.NewResultWithValue(conv)
		}
	}
	return nil
}
