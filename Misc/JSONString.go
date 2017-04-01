package Misc

type JSONString string

func (j *JSONString) UnmarshalJSON(data []byte) error{
	dataAsString := string(data)
	tmp := JSONString(dataAsString)
	j = &tmp
	return nil
}
