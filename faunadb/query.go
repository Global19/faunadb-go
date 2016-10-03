package faunadb

// Helper functions

func varargs(expr ...interface{}) interface{} {
	if len(expr) == 1 {
		return expr[0]
	}

	return expr
}

// Basic forms

func Do(exprs ...interface{}) Expr          { return fn{"do": varargs(exprs...)} }
func If(cond, then, elze interface{}) Expr  { return fn{"if": cond, "then": then, "else": elze} }
func Lambda(varName, expr interface{}) Expr { return fn{"lambda": varName, "expr": expr} }
func Let(bindings Obj, in interface{}) Expr { return fn{"let": fn(bindings), "in": in} }
func Var(name string) Expr                  { return fn{"var": name} }

func Ref(id string) Expr       { return RefV{id} }
func Null() Expr               { return NullV{} }
func Get(ref interface{}) Expr { return fn{"get": ref} }

func Create(ref, params interface{}) Expr  { return fn{"create": ref, "params": params} }
func Update(ref, params interface{}) Expr  { return fn{"update": ref, "params": params} }
func Replace(ref, params interface{}) Expr { return fn{"replace": ref, "params": params} }
func Delete(ref interface{}) Expr          { return fn{"delete": ref} }

func Exists(ref interface{}) Expr { return fn{"exists": ref} }