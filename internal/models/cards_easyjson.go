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

func easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(in *jlexer.Lexer, out *CardsOrderInput) {
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
		case "cards":
			if in.IsNull() {
				in.Skip()
				out.Cards = nil
			} else {
				in.Delim('[')
				if out.Cards == nil {
					if !in.IsDelim(']') {
						out.Cards = make([]CardOrder, 0, 4)
					} else {
						out.Cards = []CardOrder{}
					}
				} else {
					out.Cards = (out.Cards)[:0]
				}
				for !in.IsDelim(']') {
					var v1 CardOrder
					(v1).UnmarshalEasyJSON(in)
					out.Cards = append(out.Cards, v1)
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
func easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(out *jwriter.Writer, in CardsOrderInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cards\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Cards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Cards {
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
func (v CardsOrderInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardsOrderInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardsOrderInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardsOrderInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(l, v)
}
func easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(in *jlexer.Lexer, out *CardOutsideShort) {
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
		case "cardID":
			out.CardID = int64(in.Int64())
		case "cardName":
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
func easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(out *jwriter.Writer, in CardOutsideShort) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cardID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CardID))
	}
	{
		const prefix string = ",\"cardName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardOutsideShort) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardOutsideShort) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardOutsideShort) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardOutsideShort) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(l, v)
}
func easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(in *jlexer.Lexer, out *CardOutside) {
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
		case "cardID":
			out.CardID = int64(in.Int64())
		case "cardName":
			out.Name = string(in.String())
		case "cardOrder":
			out.Order = int64(in.Int64())
		case "cardTasks":
			if in.IsNull() {
				in.Skip()
				out.Tasks = nil
			} else {
				in.Delim('[')
				if out.Tasks == nil {
					if !in.IsDelim(']') {
						out.Tasks = make([]TaskOutsideShort, 0, 0)
					} else {
						out.Tasks = []TaskOutsideShort{}
					}
				} else {
					out.Tasks = (out.Tasks)[:0]
				}
				for !in.IsDelim(']') {
					var v4 TaskOutsideShort
					(v4).UnmarshalEasyJSON(in)
					out.Tasks = append(out.Tasks, v4)
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
func easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(out *jwriter.Writer, in CardOutside) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cardID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CardID))
	}
	{
		const prefix string = ",\"cardName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"cardOrder\":"
		out.RawString(prefix)
		out.Int64(int64(in.Order))
	}
	{
		const prefix string = ",\"cardTasks\":"
		out.RawString(prefix)
		if in.Tasks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Tasks {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardOutside) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardOutside) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardOutside) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardOutside) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(l, v)
}
func easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(in *jlexer.Lexer, out *CardOrder) {
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
		case "cardID":
			out.CardID = int64(in.Int64())
		case "cardOrder":
			out.Order = int64(in.Int64())
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
func easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(out *jwriter.Writer, in CardOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cardID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CardID))
	}
	{
		const prefix string = ",\"cardOrder\":"
		out.RawString(prefix)
		out.Int64(int64(in.Order))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(l, v)
}
func easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(in *jlexer.Lexer, out *CardInternal) {
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
		case "CardID":
			out.CardID = int64(in.Int64())
		case "Name":
			out.Name = string(in.String())
		case "Order":
			out.Order = int64(in.Int64())
		case "Tasks":
			if in.IsNull() {
				in.Skip()
				out.Tasks = nil
			} else {
				in.Delim('[')
				if out.Tasks == nil {
					if !in.IsDelim(']') {
						out.Tasks = make([]TaskInternalShort, 0, 0)
					} else {
						out.Tasks = []TaskInternalShort{}
					}
				} else {
					out.Tasks = (out.Tasks)[:0]
				}
				for !in.IsDelim(']') {
					var v7 TaskInternalShort
					(v7).UnmarshalEasyJSON(in)
					out.Tasks = append(out.Tasks, v7)
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
func easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(out *jwriter.Writer, in CardInternal) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"CardID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CardID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Order\":"
		out.RawString(prefix)
		out.Int64(int64(in.Order))
	}
	{
		const prefix string = ",\"Tasks\":"
		out.RawString(prefix)
		if in.Tasks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Tasks {
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
func (v CardInternal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardInternal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardInternal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardInternal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(l, v)
}
func easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(in *jlexer.Lexer, out *CardInput) {
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
		case "cardID":
			out.CardID = int64(in.Int64())
		case "boardID":
			out.BoardID = int64(in.Int64())
		case "cardName":
			out.Name = string(in.String())
		case "cardOrder":
			out.Order = int64(in.Int64())
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
func easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(out *jwriter.Writer, in CardInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cardID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.CardID))
	}
	{
		const prefix string = ",\"boardID\":"
		out.RawString(prefix)
		out.Int64(int64(in.BoardID))
	}
	{
		const prefix string = ",\"cardName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"cardOrder\":"
		out.RawString(prefix)
		out.Int64(int64(in.Order))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(l, v)
}
func easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(in *jlexer.Lexer, out *Card) {
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
		case "columnID":
			out.CardID = int64(in.Int64())
		case "boardID":
			out.BoardID = int64(in.Int64())
		case "cardName":
			out.Name = string(in.String())
		case "order":
			out.Order = int64(in.Int64())
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
func easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(out *jwriter.Writer, in Card) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"columnID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CardID))
	}
	{
		const prefix string = ",\"boardID\":"
		out.RawString(prefix)
		out.Int64(int64(in.BoardID))
	}
	{
		const prefix string = ",\"cardName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"order\":"
		out.RawString(prefix)
		out.Int64(int64(in.Order))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Card) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Card) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1b92225fEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Card) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Card) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1b92225fDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(l, v)
}
