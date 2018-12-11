package go_graphql_gqlgen

import (
	bytes "bytes"
	context "context"
	strconv "strconv"
	sync "sync"

	graphql "github.com/99designs/gqlgen/graphql"
	introspection "github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser"
	ast "github.com/vektah/gqlparser/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Employee struct {
		Id        func(childComplexity int) int
		Name      func(childComplexity int) int
		Email     func(childComplexity int) int
		Contactno func(childComplexity int) int
		Position  func(childComplexity int) int
	}

	Mutation struct {
		CreateNewEmployee func(childComplexity int, input NewEmployee) int
	}

	Query struct {
		Employees func(childComplexity int) int
	}
}

type MutationResolver interface {
	CreateNewEmployee(ctx context.Context, input NewEmployee) (Employee, error)
}
type QueryResolver interface {
	Employees(ctx context.Context) ([]Employee, error)
}

func field_Mutation_createNewEmployee_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 NewEmployee
	if tmp, ok := rawArgs["input"]; ok {
		var err error
		arg0, err = UnmarshalNewEmployee(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil

}

func field_Query___type_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		var err error
		arg0, err = graphql.UnmarshalString(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil

}

func field___Type_fields_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil

}

func field___Type_enumValues_args(rawArgs map[string]interface{}) (map[string]interface{}, error) {
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil

}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	switch typeName + "." + field {

	case "Employee.id":
		if e.complexity.Employee.Id == nil {
			break
		}

		return e.complexity.Employee.Id(childComplexity), true

	case "Employee.name":
		if e.complexity.Employee.Name == nil {
			break
		}

		return e.complexity.Employee.Name(childComplexity), true

	case "Employee.email":
		if e.complexity.Employee.Email == nil {
			break
		}

		return e.complexity.Employee.Email(childComplexity), true

	case "Employee.contactno":
		if e.complexity.Employee.Contactno == nil {
			break
		}

		return e.complexity.Employee.Contactno(childComplexity), true

	case "Employee.position":
		if e.complexity.Employee.Position == nil {
			break
		}

		return e.complexity.Employee.Position(childComplexity), true

	case "Mutation.createNewEmployee":
		if e.complexity.Mutation.CreateNewEmployee == nil {
			break
		}

		args, err := field_Mutation_createNewEmployee_args(rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateNewEmployee(childComplexity, args["input"].(NewEmployee)), true

	case "Query.employees":
		if e.complexity.Query.Employees == nil {
			break
		}

		return e.complexity.Query.Employees(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Query(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Query(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:       buf,
		Errors:     ec.Errors,
		Extensions: ec.Extensions}
}

func (e *executableSchema) Mutation(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Mutation(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:       buf,
		Errors:     ec.Errors,
		Extensions: ec.Extensions,
	}
}

func (e *executableSchema) Subscription(ctx context.Context, op *ast.OperationDefinition) func() *graphql.Response {
	return graphql.OneShot(graphql.ErrorResponse(ctx, "subscriptions are not supported"))
}

type executionContext struct {
	*graphql.RequestContext
	*executableSchema
}

var employeeImplementors = []string{"Employee"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Employee(ctx context.Context, sel ast.SelectionSet, obj *Employee) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, employeeImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Employee")
		case "id":
			out.Values[i] = ec._Employee_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "name":
			out.Values[i] = ec._Employee_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "email":
			out.Values[i] = ec._Employee_email(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "contactno":
			out.Values[i] = ec._Employee_contactno(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "position":
			out.Values[i] = ec._Employee_position(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _Employee_id(ctx context.Context, field graphql.CollectedField, obj *Employee) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "Employee",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalID(res)
}

// nolint: vetshadow
func (ec *executionContext) _Employee_name(ctx context.Context, field graphql.CollectedField, obj *Employee) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "Employee",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) _Employee_email(ctx context.Context, field graphql.CollectedField, obj *Employee) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "Employee",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Email, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) _Employee_contactno(ctx context.Context, field graphql.CollectedField, obj *Employee) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "Employee",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Contactno, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) _Employee_position(ctx context.Context, field graphql.CollectedField, obj *Employee) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "Employee",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Position, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

var mutationImplementors = []string{"Mutation"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, mutationImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Mutation",
	})

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createNewEmployee":
			out.Values[i] = ec._Mutation_createNewEmployee(ctx, field)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _Mutation_createNewEmployee(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field_Mutation_createNewEmployee_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "Mutation",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateNewEmployee(rctx, args["input"].(NewEmployee))
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(Employee)
	rctx.Result = res

	return ec._Employee(ctx, field.Selections, &res)
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, queryImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
	})

	var wg sync.WaitGroup
	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "employees":
			wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				out.Values[i] = ec._Query_employees(ctx, field)
				if out.Values[i] == graphql.Null {
					invalid = true
				}
				wg.Done()
			}(i, field)
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	wg.Wait()
	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) _Query_employees(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().Employees(rctx)
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]Employee)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec._Employee(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field_Query___type_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(args["name"].(string)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "Query",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, nil, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec.___Schema(ctx, field.Selections, res)
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __DirectiveImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Locations, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]string)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))

	for idx1 := range res {
		arr1[idx1] = func() graphql.Marshaler {
			return graphql.MarshalString(res[idx1])
		}()
	}

	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Directive",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___InputValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __EnumValueImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__EnumValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __FieldImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Args, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___InputValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.IsDeprecated, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	rctx.Result = res
	return graphql.MarshalBoolean(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Field",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DeprecationReason, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __InputValueImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Type, nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__InputValue",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DefaultValue, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __SchemaImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Types(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Type(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.QueryType(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.MutationType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SubscriptionType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

// nolint: vetshadow
func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Schema",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Directives(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Directive(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __TypeImplementors)

	out := graphql.NewOrderedMap(len(fields))
	invalid := false
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalid = true
			}
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	if invalid {
		return graphql.Null
	}
	return out
}

// nolint: vetshadow
func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Kind(), nil
	})
	if resTmp == nil {
		if !ec.HasError(rctx) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Description(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	rctx.Result = res
	return graphql.MarshalString(res)
}

// nolint: vetshadow
func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field___Type_fields_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Field(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Interfaces(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Type(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PossibleTypes(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___Type(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := field___Type_enumValues_args(rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   args,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___EnumValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.InputFields(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	rctx.Result = res

	arr1 := make(graphql.Array, len(res))
	var wg sync.WaitGroup

	isLen1 := len(res) == 1
	if !isLen1 {
		wg.Add(len(res))
	}

	for idx1 := range res {
		idx1 := idx1
		rctx := &graphql.ResolverContext{
			Index:  &idx1,
			Result: &res[idx1],
		}
		ctx := graphql.WithResolverContext(ctx, rctx)
		f := func(idx1 int) {
			if !isLen1 {
				defer wg.Done()
			}
			arr1[idx1] = func() graphql.Marshaler {

				return ec.___InputValue(ctx, field.Selections, &res[idx1])
			}()
		}
		if isLen1 {
			f(idx1)
		} else {
			go f(idx1)
		}

	}
	wg.Wait()
	return arr1
}

// nolint: vetshadow
func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := &graphql.ResolverContext{
		Object: "__Type",
		Args:   nil,
		Field:  field,
	}
	ctx = graphql.WithResolverContext(ctx, rctx)
	resTmp := ec.FieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.OfType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	rctx.Result = res

	if res == nil {
		return graphql.Null
	}

	return ec.___Type(ctx, field.Selections, res)
}

func UnmarshalNewEmployee(v interface{}) (NewEmployee, error) {
	var it NewEmployee
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "name":
			var err error
			it.Name, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "email":
			var err error
			it.Email, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "contactno":
			var err error
			it.Contactno, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "position":
			var err error
			it.Position, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) FieldMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver) (ret interface{}) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	res, err := ec.ResolverMiddleware(ctx, next)
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	return res
}

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name])
}

var parsedSchema = gqlparser.MustLoadSchema(
	&ast.Source{Name: "schema.graphql", Input: `type Employee{
    id: ID!
    name: String!
    email: String!
    contactno: String!
    position: String!
}
input NewEmployee{
    name: String!
    email: String!
    contactno: String!
    position: String!
}
type Mutation{
    createNewEmployee(input: NewEmployee!): Employee!
}
type Query{
    employees: [Employee!]!
}`},
)
