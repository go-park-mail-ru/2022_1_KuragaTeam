// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package internal

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

func easyjsonD2b7633eDecodeMyappInternal(in *jlexer.Lexer, out *Season) {
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
		case "number":
			out.Number = int(in.Int())
		case "episodes":
			if in.IsNull() {
				in.Skip()
				out.Episodes = nil
			} else {
				in.Delim('[')
				if out.Episodes == nil {
					if !in.IsDelim(']') {
						out.Episodes = make([]Episode, 0, 0)
					} else {
						out.Episodes = []Episode{}
					}
				} else {
					out.Episodes = (out.Episodes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Episode
					(v1).UnmarshalEasyJSON(in)
					out.Episodes = append(out.Episodes, v1)
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
func easyjsonD2b7633eEncodeMyappInternal(out *jwriter.Writer, in Season) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"number\":"
		out.RawString(prefix)
		out.Int(int(in.Number))
	}
	{
		const prefix string = ",\"episodes\":"
		out.RawString(prefix)
		if in.Episodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Episodes {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Season) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Season) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Season) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Season) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal1(in *jlexer.Lexer, out *PersonInMovieDTO) {
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
			out.Position = string(in.String())
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
func easyjsonD2b7633eEncodeMyappInternal1(out *jwriter.Writer, in PersonInMovieDTO) {
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
		out.String(string(in.Position))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PersonInMovieDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PersonInMovieDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PersonInMovieDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PersonInMovieDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal1(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal2(in *jlexer.Lexer, out *Person) {
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
		case "addit_photo_1":
			out.AdditPhoto1 = string(in.String())
		case "addit_photo_2":
			out.AdditPhoto2 = string(in.String())
		case "description":
			out.Description = string(in.String())
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
					var v4 string
					v4 = string(in.String())
					out.Position = append(out.Position, v4)
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
func easyjsonD2b7633eEncodeMyappInternal2(out *jwriter.Writer, in Person) {
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
		const prefix string = ",\"addit_photo_1\":"
		out.RawString(prefix)
		out.String(string(in.AdditPhoto1))
	}
	{
		const prefix string = ",\"addit_photo_2\":"
		out.RawString(prefix)
		out.String(string(in.AdditPhoto2))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix)
		if in.Position == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Position {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Person) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Person) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Person) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Person) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal2(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal3(in *jlexer.Lexer, out *MovieRatingDTO) {
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
			out.MovieID = int(in.IntStr())
		case "rating":
			out.Rating = int(in.IntStr())
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
func easyjsonD2b7633eEncodeMyappInternal3(out *jwriter.Writer, in MovieRatingDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.IntStr(int(in.MovieID))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.IntStr(int(in.Rating))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieRatingDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieRatingDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieRatingDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieRatingDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal3(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal4(in *jlexer.Lexer, out *MovieInfo) {
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
					var v7 Genre
					(v7).UnmarshalEasyJSON(in)
					out.Genre = append(out.Genre, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "picture":
			out.Picture = string(in.String())
		case "rating":
			out.Rating = float32(in.Float32())
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
func easyjsonD2b7633eEncodeMyappInternal4(out *jwriter.Writer, in MovieInfo) {
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
			for v8, v9 := range in.Genre {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"picture\":"
		out.RawString(prefix)
		out.String(string(in.Picture))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal4(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal5(in *jlexer.Lexer, out *MovieCompilationResponse) {
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
		case "compilation_name":
			out.Name = string(in.String())
		case "movies":
			if in.IsNull() {
				in.Skip()
				out.Movies = nil
			} else {
				in.Delim('[')
				if out.Movies == nil {
					if !in.IsDelim(']') {
						out.Movies = make([]MovieInfo, 0, 0)
					} else {
						out.Movies = []MovieInfo{}
					}
				} else {
					out.Movies = (out.Movies)[:0]
				}
				for !in.IsDelim(']') {
					var v10 MovieInfo
					(v10).UnmarshalEasyJSON(in)
					out.Movies = append(out.Movies, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "has_next_page":
			out.HasNextPage = bool(in.Bool())
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
func easyjsonD2b7633eEncodeMyappInternal5(out *jwriter.Writer, in MovieCompilationResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"compilation_name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"movies\":"
		out.RawString(prefix)
		if in.Movies == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Movies {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"has_next_page\":"
		out.RawString(prefix)
		out.Bool(bool(in.HasNextPage))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieCompilationResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieCompilationResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieCompilationResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieCompilationResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal5(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal6(in *jlexer.Lexer, out *MovieCompilation) {
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
		case "compilation_name":
			out.Name = string(in.String())
		case "movies":
			if in.IsNull() {
				in.Skip()
				out.Movies = nil
			} else {
				in.Delim('[')
				if out.Movies == nil {
					if !in.IsDelim(']') {
						out.Movies = make([]MovieInfo, 0, 0)
					} else {
						out.Movies = []MovieInfo{}
					}
				} else {
					out.Movies = (out.Movies)[:0]
				}
				for !in.IsDelim(']') {
					var v13 MovieInfo
					(v13).UnmarshalEasyJSON(in)
					out.Movies = append(out.Movies, v13)
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
func easyjsonD2b7633eEncodeMyappInternal6(out *jwriter.Writer, in MovieCompilation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"compilation_name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"movies\":"
		out.RawString(prefix)
		if in.Movies == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Movies {
				if v14 > 0 {
					out.RawByte(',')
				}
				(v15).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MovieCompilation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieCompilation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MovieCompilation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieCompilation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal6(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal7(in *jlexer.Lexer, out *Movie) {
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
		case "is_movie":
			out.IsMovie = bool(in.Bool())
		case "name_picture":
			out.NamePicture = string(in.String())
		case "year":
			out.Year = int(in.Int())
		case "duration":
			out.Duration = string(in.String())
		case "age_limit":
			out.AgeLimit = int(in.Int())
		case "description":
			out.Description = string(in.String())
		case "kinopoisk_rating":
			out.KinopoiskRating = float32(in.Float32())
		case "rating":
			out.Rating = float32(in.Float32())
		case "tagline":
			out.Tagline = string(in.String())
		case "picture":
			out.Picture = string(in.String())
		case "video":
			out.Video = string(in.String())
		case "trailer":
			out.Trailer = string(in.String())
		case "season":
			if in.IsNull() {
				in.Skip()
				out.Season = nil
			} else {
				in.Delim('[')
				if out.Season == nil {
					if !in.IsDelim(']') {
						out.Season = make([]Season, 0, 1)
					} else {
						out.Season = []Season{}
					}
				} else {
					out.Season = (out.Season)[:0]
				}
				for !in.IsDelim(']') {
					var v16 Season
					(v16).UnmarshalEasyJSON(in)
					out.Season = append(out.Season, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "country":
			if in.IsNull() {
				in.Skip()
				out.Country = nil
			} else {
				in.Delim('[')
				if out.Country == nil {
					if !in.IsDelim(']') {
						out.Country = make([]string, 0, 4)
					} else {
						out.Country = []string{}
					}
				} else {
					out.Country = (out.Country)[:0]
				}
				for !in.IsDelim(']') {
					var v17 string
					v17 = string(in.String())
					out.Country = append(out.Country, v17)
					in.WantComma()
				}
				in.Delim(']')
			}
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
					var v18 Genre
					(v18).UnmarshalEasyJSON(in)
					out.Genre = append(out.Genre, v18)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "staff":
			if in.IsNull() {
				in.Skip()
				out.Staff = nil
			} else {
				in.Delim('[')
				if out.Staff == nil {
					if !in.IsDelim(']') {
						out.Staff = make([]PersonInMovieDTO, 0, 1)
					} else {
						out.Staff = []PersonInMovieDTO{}
					}
				} else {
					out.Staff = (out.Staff)[:0]
				}
				for !in.IsDelim(']') {
					var v19 PersonInMovieDTO
					(v19).UnmarshalEasyJSON(in)
					out.Staff = append(out.Staff, v19)
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
func easyjsonD2b7633eEncodeMyappInternal7(out *jwriter.Writer, in Movie) {
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
		const prefix string = ",\"is_movie\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsMovie))
	}
	{
		const prefix string = ",\"name_picture\":"
		out.RawString(prefix)
		out.String(string(in.NamePicture))
	}
	{
		const prefix string = ",\"year\":"
		out.RawString(prefix)
		out.Int(int(in.Year))
	}
	{
		const prefix string = ",\"duration\":"
		out.RawString(prefix)
		out.String(string(in.Duration))
	}
	{
		const prefix string = ",\"age_limit\":"
		out.RawString(prefix)
		out.Int(int(in.AgeLimit))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"kinopoisk_rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.KinopoiskRating))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	{
		const prefix string = ",\"tagline\":"
		out.RawString(prefix)
		out.String(string(in.Tagline))
	}
	{
		const prefix string = ",\"picture\":"
		out.RawString(prefix)
		out.String(string(in.Picture))
	}
	{
		const prefix string = ",\"video\":"
		out.RawString(prefix)
		out.String(string(in.Video))
	}
	{
		const prefix string = ",\"trailer\":"
		out.RawString(prefix)
		out.String(string(in.Trailer))
	}
	{
		const prefix string = ",\"season\":"
		out.RawString(prefix)
		if in.Season == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.Season {
				if v20 > 0 {
					out.RawByte(',')
				}
				(v21).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		if in.Country == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v22, v23 := range in.Country {
				if v22 > 0 {
					out.RawByte(',')
				}
				out.String(string(v23))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"genre\":"
		out.RawString(prefix)
		if in.Genre == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v24, v25 := range in.Genre {
				if v24 > 0 {
					out.RawByte(',')
				}
				(v25).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"staff\":"
		out.RawString(prefix)
		if in.Staff == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.Staff {
				if v26 > 0 {
					out.RawByte(',')
				}
				(v27).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Movie) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Movie) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Movie) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Movie) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal7(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal8(in *jlexer.Lexer, out *MainMovieInfoDTO) {
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
		case "name_picture":
			out.NamePicture = string(in.String())
		case "tagline":
			out.Tagline = string(in.String())
		case "picture":
			out.Picture = string(in.String())
		case "is_movie":
			out.IsMovie = bool(in.Bool())
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
func easyjsonD2b7633eEncodeMyappInternal8(out *jwriter.Writer, in MainMovieInfoDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"name_picture\":"
		out.RawString(prefix)
		out.String(string(in.NamePicture))
	}
	{
		const prefix string = ",\"tagline\":"
		out.RawString(prefix)
		out.String(string(in.Tagline))
	}
	{
		const prefix string = ",\"picture\":"
		out.RawString(prefix)
		out.String(string(in.Picture))
	}
	{
		const prefix string = ",\"is_movie\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsMovie))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MainMovieInfoDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MainMovieInfoDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MainMovieInfoDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MainMovieInfoDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal8(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal9(in *jlexer.Lexer, out *Genre) {
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
func easyjsonD2b7633eEncodeMyappInternal9(out *jwriter.Writer, in Genre) {
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
	easyjsonD2b7633eEncodeMyappInternal9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Genre) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Genre) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Genre) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal9(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal10(in *jlexer.Lexer, out *Episode) {
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
		case "number":
			out.Number = int(in.Int())
		case "description":
			out.Description = string(in.String())
		case "video":
			out.Video = string(in.String())
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
func easyjsonD2b7633eEncodeMyappInternal10(out *jwriter.Writer, in Episode) {
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
		const prefix string = ",\"number\":"
		out.RawString(prefix)
		out.Int(int(in.Number))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"video\":"
		out.RawString(prefix)
		out.String(string(in.Video))
	}
	{
		const prefix string = ",\"picture\":"
		out.RawString(prefix)
		out.String(string(in.Picture))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Episode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Episode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Episode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Episode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal10(l, v)
}
func easyjsonD2b7633eDecodeMyappInternal11(in *jlexer.Lexer, out *AllMoviesResponse) {
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
						out.Movies = make([]MovieInfo, 0, 0)
					} else {
						out.Movies = []MovieInfo{}
					}
				} else {
					out.Movies = (out.Movies)[:0]
				}
				for !in.IsDelim(']') {
					var v28 MovieInfo
					(v28).UnmarshalEasyJSON(in)
					out.Movies = append(out.Movies, v28)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "has_next_page":
			out.HasNextPage = bool(in.Bool())
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
func easyjsonD2b7633eEncodeMyappInternal11(out *jwriter.Writer, in AllMoviesResponse) {
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
			for v29, v30 := range in.Movies {
				if v29 > 0 {
					out.RawByte(',')
				}
				(v30).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"has_next_page\":"
		out.RawString(prefix)
		out.Bool(bool(in.HasNextPage))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AllMoviesResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeMyappInternal11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AllMoviesResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeMyappInternal11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AllMoviesResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeMyappInternal11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AllMoviesResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeMyappInternal11(l, v)
}
