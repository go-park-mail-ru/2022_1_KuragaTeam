// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson68b2ec0cDecodeMyappInternalModels(in *jlexer.Lexer, out *SearchCompilation) {
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
		case "movies":
			if in.IsNull() {
				in.Skip()
				out.Movies = nil
			} else {
				in.Delim('[')
				if out.Movies == nil {
					if !in.IsDelim(']') {
						out.Movies = make([]MovieInfo, 0, 1)
					} else {
						out.Movies = []MovieInfo{}
					}
				} else {
					out.Movies = (out.Movies)[:0]
				}
				for !in.IsDelim(']') {
					var v1 MovieInfo
					(v1).UnmarshalEasyJSON(in)
					out.Movies = append(out.Movies, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "series":
			if in.IsNull() {
				in.Skip()
				out.Series = nil
			} else {
				in.Delim('[')
				if out.Series == nil {
					if !in.IsDelim(']') {
						out.Series = make([]MovieInfo, 0, 1)
					} else {
						out.Series = []MovieInfo{}
					}
				} else {
					out.Series = (out.Series)[:0]
				}
				for !in.IsDelim(']') {
					var v2 MovieInfo
					(v2).UnmarshalEasyJSON(in)
					out.Series = append(out.Series, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "persons":
			if in.IsNull() {
				in.Skip()
				out.Persons = nil
			} else {
				in.Delim('[')
				if out.Persons == nil {
					if !in.IsDelim(']') {
						out.Persons = make([]PersonInfo, 0, 1)
					} else {
						out.Persons = []PersonInfo{}
					}
				} else {
					out.Persons = (out.Persons)[:0]
				}
				for !in.IsDelim(']') {
					var v3 PersonInfo
					(v3).UnmarshalEasyJSON(in)
					out.Persons = append(out.Persons, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson68b2ec0cEncodeMyappInternalModels(out *jwriter.Writer, in SearchCompilation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"movies\":"
		out.RawString(prefix[1:])
		if in.Movies == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Movies {
				if v4 > 0 {
					out.RawByte(',')
				}
				(v5).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"series\":"
		out.RawString(prefix)
		if in.Series == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Series {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"persons\":"
		out.RawString(prefix)
		if in.Persons == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Persons {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchCompilation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson68b2ec0cEncodeMyappInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchCompilation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson68b2ec0cEncodeMyappInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchCompilation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson68b2ec0cDecodeMyappInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchCompilation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson68b2ec0cDecodeMyappInternalModels(l, v)
}
func easyjson68b2ec0cDecodeMyappInternalModels1(in *jlexer.Lexer, out *PersonInfo) {
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
		case "id":
			out.ID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "photo":
			out.Photo = string(in.String())
		case "position":
			if in.IsNull() {
				in.Skip()
				out.Position = nil
			} else {
				in.Delim('[')
				if out.Position == nil {
					if !in.IsDelim(']') {
						out.Position = make([]string, 0, 4)
					} else {
						out.Position = []string{}
					}
				} else {
					out.Position = (out.Position)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Position = append(out.Position, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson68b2ec0cEncodeMyappInternalModels1(out *jwriter.Writer, in PersonInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"photo\":"
		out.RawString(prefix)
		out.String(string(in.Photo))
	}
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix)
		if in.Position == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Position {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PersonInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson68b2ec0cEncodeMyappInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PersonInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson68b2ec0cEncodeMyappInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PersonInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson68b2ec0cDecodeMyappInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PersonInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson68b2ec0cDecodeMyappInternalModels1(l, v)
}
func easyjson68b2ec0cDecodeMyappInternalModels2(in *jlexer.Lexer, out *MovieInfo) {
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
		case "id":
			out.ID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "genre":
			if in.IsNull() {
				in.Skip()
				out.Genre = nil
			} else {
				in.Delim('[')
				if out.Genre == nil {
					if !in.IsDelim(']') {
						out.Genre = make([]Genre, 0, 2)
					} else {
						out.Genre = []Genre{}
					}
				} else {
					out.Genre = (out.Genre)[:0]
				}
				for !in.IsDelim(']') {
					var v13 Genre
					(v13).UnmarshalEasyJSON(in)
					out.Genre = append(out.Genre, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "picture":
			out.Picture = string(in.String())
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
func easyjson68b2ec0cEncodeMyappInternalModels2(out *jwriter.Writer, in MovieInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"genre\":"
		out.RawString(prefix)
		if in.Genre == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Genre {
				if v14 > 0 {
					out.RawByte(',')
				}
				(v15).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"picture\":"
		out.RawString(prefix)
		out.String(string(in.Picture))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson68b2ec0cEncodeMyappInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson68b2ec0cEncodeMyappInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson68b2ec0cDecodeMyappInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson68b2ec0cDecodeMyappInternalModels2(l, v)
}
func easyjson68b2ec0cDecodeMyappInternalModels3(in *jlexer.Lexer, out *Genre) {
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
		case "id":
			out.ID = int(in.Int())
		case "name":
			out.Name = string(in.String())
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
func easyjson68b2ec0cEncodeMyappInternalModels3(out *jwriter.Writer, in Genre) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Genre) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson68b2ec0cEncodeMyappInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Genre) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson68b2ec0cEncodeMyappInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Genre) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson68b2ec0cDecodeMyappInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Genre) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson68b2ec0cDecodeMyappInternalModels3(l, v)
}
func easyjson68b2ec0cDecodeMyappInternalModels4(in *jlexer.Lexer, out *FindDTO) {
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
		case "find":
			out.Text = string(in.String())
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
func easyjson68b2ec0cEncodeMyappInternalModels4(out *jwriter.Writer, in FindDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"find\":"
		out.RawString(prefix[1:])
		out.String(string(in.Text))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FindDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson68b2ec0cEncodeMyappInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FindDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson68b2ec0cEncodeMyappInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FindDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson68b2ec0cDecodeMyappInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FindDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson68b2ec0cDecodeMyappInternalModels4(l, v)
}
