// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package lib

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9f2eff5fDecodeGithubComIsuconIsucon11QualifyIsuconditionLib(in *jlexer.Lexer, out *PostIsuConditionRequests) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostIsuConditionRequests, 0, 1)
			} else {
				*out = PostIsuConditionRequests{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostIsuConditionRequest
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComIsuconIsucon11QualifyIsuconditionLib(out *jwriter.Writer, in PostIsuConditionRequests) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v PostIsuConditionRequests) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComIsuconIsucon11QualifyIsuconditionLib(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostIsuConditionRequests) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComIsuconIsucon11QualifyIsuconditionLib(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostIsuConditionRequests) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComIsuconIsucon11QualifyIsuconditionLib(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostIsuConditionRequests) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComIsuconIsucon11QualifyIsuconditionLib(l, v)
}
func easyjson9f2eff5fDecodeGithubComIsuconIsucon11QualifyIsuconditionLib1(in *jlexer.Lexer, out *PostIsuConditionRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "is_sitting":
			out.IsSitting = bool(in.Bool())
		case "condition":
			out.Condition = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "timestamp":
			out.Timestamp = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComIsuconIsucon11QualifyIsuconditionLib1(out *jwriter.Writer, in PostIsuConditionRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"is_sitting\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.IsSitting))
	}
	{
		const prefix string = ",\"condition\":"
		out.RawString(prefix)
		out.String(string(in.Condition))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix)
		out.Int64(int64(in.Timestamp))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostIsuConditionRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComIsuconIsucon11QualifyIsuconditionLib1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostIsuConditionRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComIsuconIsucon11QualifyIsuconditionLib1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostIsuConditionRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComIsuconIsucon11QualifyIsuconditionLib1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostIsuConditionRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComIsuconIsucon11QualifyIsuconditionLib1(l, v)
}
