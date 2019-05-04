package util

//type Number struct {
//	Int     int
//	Float64 float64
//}
//
//func (a Number) Compare(b Number, operate string) bool {
//	if operate == ">" {
//		return a.Int > b.Int || a.Float64 > b.Float64
//	}
//	if operate == "=" {
//		return a.Int == b.Int && a.Float64 == b.Float64
//	}
//	if operate == "<" {
//		return a.Int < b.Int || a.Float64 < b.Float64
//	}
//	if operate == ">=" {
//		return a.Int >= b.Int && a.Float64 >= b.Float64
//	}
//	if operate == "<=" {
//		return a.Int <= b.Int && a.Float64 <= b.Float64
//	}
//	return false
//}
//
//func ParseNumber(num interface{}, typ string) Number {
//	a := Number{}
//	if typ == "int" {
//		switch num.(type) {
//		case string:
//			a.Int, _ = strconv.Atoi(num.(string))
//		case int:
//			a.Int = num.(int)
//		case float64:
//			a.Float64 = num.(float64)
//		case json.Number:
//			t, _ := num.(json.Number).Int64()
//			a.Int = int(t)
//		}
//	} else if typ == "float64" {
//		switch num.(type) {
//		case string:
//			a.Float64, _ = strconv.ParseFloat(num.(string), 64)
//		case int:
//			a.Float64 = num.(float64)
//		case float64:
//			a.Float64 = num.(float64)
//		case json.Number:
//			a.Float64, _ = num.(json.Number).Float64()
//		}
//	} else {
//		switch num.(type) {
//		case int:
//			a.Int = num.(int)
//		case float64:
//			a.Float64 = num.(float64)
//		case string:
//			a.Float64, _ = strconv.ParseFloat(num.(string), 64)
//		case json.Number:
//			a.Float64, _ = num.(json.Number).Float64()
//		}
//	}
//	return a
//}
//
//func ParserSearchSql(search dto.Search, args, comma string) string {
//	if search == nil {
//		return args
//	}
//	for _, row := range search{
//		switch row.Value.(type) {
//		case float64:
//			args = fmt.Sprintf(`%s %s %s %s %f`, args, comma, row.Name, row.Operate, row.Value.(float64))
//			comma = "AND"
//		case int:
//			args = fmt.Sprintf(`%s %s %s %s %d`, args, comma, row.Name, row.Operate, row.Value.(int))
//			comma = "AND"
//		case json.Number:
//			a, _ := row.Value.(json.Number).Float64()
//			args = fmt.Sprintf(`%s %s %s %s %f`, args, comma, row.Name, row.Operate, a)
//			comma = "AND"
//		case string:
//			if row.Operate == "=" {
//				args = fmt.Sprintf(`%s %s %s %s '%s'`, args, comma, row.Name, row.Operate, row.Value.(string))
//				comma = "AND"
//			} else {
//				a := "%" + row.Value.(string) + "%"
//				args = fmt.Sprintf(`%s %s %s %s '%s'`, args, comma, row.Name, row.Operate, a)
//				comma = "AND"
//			}
//		}
//	}
//	return args
//}
//
//func SplicingSql(args, comma string, object interface{}, ignores ...string) (string, string, error) {
//	buf, e := json.Marshal(object)
//	if e != nil {
//		return args, comma, e
//	}
//	fMap := make(map[string]interface{})
//	e = json.Unmarshal(buf, &fMap)
//	if e != nil {
//		return args, comma, e
//	}
//label:
//	for k, v := range fMap {
//		for _, ignore := range ignores {
//			if k == ignore {
//				continue label
//			}
//		}
//		args, comma = SplicingDebris(args, comma, k, v)
//	}
//	return args, comma, e
//}
//func SplicingDebris(args, comma string, k string, v interface{}) (string, string) {
//	if !IsValid(v) {
//		return args, comma
//	}
//	switch v.(type) {
//	case int:
//		args = fmt.Sprintf(`%s %s %s=%d`, args, comma, k, v.(int))
//		comma = parseComma(comma)
//	case bool:
//		args = fmt.Sprintf(`%s %s %s=%t`, args, comma, k, v.(bool))
//		comma = parseComma(comma)
//	case string:
//		args = fmt.Sprintf(`%s %s %s='%s'`, args, comma, k, v.(string))
//		comma = parseComma(comma)
//	case float64:
//		args = fmt.Sprintf(`%s %s %s=%f`, args, comma, k, v.(float64))
//		comma = parseComma(comma)
//	default:
//		log.Println("sql type is unknow")
//	}
//	return args, comma
//}
//
//func SplicingInsertSql(object interface{}, ignores ...string) (args1 string, args2 string, e error) {
//	buf, e := json.Marshal(object)
//	if e != nil {
//		return args1, args2, e
//	}
//	fMap := make(map[string]interface{})
//	e = json.Unmarshal(buf, &fMap)
//	if e != nil {
//		return args1, args2, e
//	}
//	comma := ""
//label:
//	for k, v := range fMap {
//		for _, ingore := range ignores {
//			if k == ingore {
//				continue label
//			}
//		}
//		args1, args2, comma = SplicInsertDebris(args1, args2, comma, k, v)
//	}
//	return args1, args2, nil
//}
//
//func SplicInsertDebris(args1, args2, comma string, k string, v interface{}) (string, string, string) {
//	args1 = fmt.Sprintf(`%s%s %s`, args1, comma, k)
//	switch v.(type) {
//	case int:
//		args2 = fmt.Sprintf(`%s%s %d`, args2, comma, v.(int))
//		comma = ","
//	case bool:
//		args2 = fmt.Sprintf(`%s%s %t`, args2, comma, v.(bool))
//		comma = ","
//	case string:
//		args2 = fmt.Sprintf(`%s%s '%s'`, args2, comma, v.(string))
//		comma = ","
//	case float64:
//		args2 = fmt.Sprintf(`%s%s %f`, args2, comma, v.(float64))
//		comma = ","
//	default:
//		log.Printf("v type is not in all type, k: %s, v:%v", k, v)
//	}
//	return args1, args2, comma
//}
//
//func parseComma(comma string) string {
//	comma = strings.ToUpper(comma)
//	if comma == "WHERE" || len(comma) == 0 || comma == "AND" {
//		return "AND"
//	} else if comma == "SET" || comma == "," {
//		return ","
//	}
//	return "AND"
//}
//func IsValid(object interface{}) bool {
//	switch object.(type) {
//	case string:
//		if len(object.(string)) != 0 && object.(string) != "[]" {
//			return true
//		}
//	case int:
//		if object.(int) > 0 {
//			return true
//		}
//	case float64:
//		if object.(float64) > 0 {
//			return true
//		}
//	case bool:
//		return object.(bool)
//	default:
//		log.Printf("object is not in all type")
//	}
//	return false
//}
