package convert

func ConvertMapToSetStr(fields map[string]interface{}) (string, []interface{}) {
	var setStr string
	var args []interface{} = make([]interface{}, 0, len(fields))

	for k, v := range fields {
		setStr += k + "=" + "?,"
		args = append(args, v)
	}

	return setStr[:len(setStr)-1], args
}
