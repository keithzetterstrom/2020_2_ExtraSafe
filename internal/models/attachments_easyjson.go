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

func easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(in *jlexer.Lexer, out *AttachmentOutside) {
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
		case "attachmentID":
			out.AttachmentID = int64(in.Int64())
		case "attachmentFileName":
			out.Filename = string(in.String())
		case "attachmentFilePath":
			out.Filepath = string(in.String())
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
func easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(out *jwriter.Writer, in AttachmentOutside) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"attachmentID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.AttachmentID))
	}
	{
		const prefix string = ",\"attachmentFileName\":"
		out.RawString(prefix)
		out.String(string(in.Filename))
	}
	{
		const prefix string = ",\"attachmentFilePath\":"
		out.RawString(prefix)
		out.String(string(in.Filepath))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AttachmentOutside) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AttachmentOutside) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AttachmentOutside) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AttachmentOutside) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(l, v)
}
func easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(in *jlexer.Lexer, out *AttachmentInternal) {
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
		case "TaskID":
			out.TaskID = int64(in.Int64())
		case "AttachmentID":
			out.AttachmentID = int64(in.Int64())
		case "Filename":
			out.Filename = string(in.String())
		case "Filepath":
			out.Filepath = string(in.String())
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
func easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(out *jwriter.Writer, in AttachmentInternal) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"TaskID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.TaskID))
	}
	{
		const prefix string = ",\"AttachmentID\":"
		out.RawString(prefix)
		out.Int64(int64(in.AttachmentID))
	}
	{
		const prefix string = ",\"Filename\":"
		out.RawString(prefix)
		out.String(string(in.Filename))
	}
	{
		const prefix string = ",\"Filepath\":"
		out.RawString(prefix)
		out.String(string(in.Filepath))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AttachmentInternal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AttachmentInternal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AttachmentInternal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AttachmentInternal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(l, v)
}
func easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(in *jlexer.Lexer, out *AttachmentInput) {
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
		case "taskID":
			out.TaskID = int64(in.Int64())
		case "attachmentID":
			out.AttachmentID = int64(in.Int64())
		case "attachmentFileName":
			out.Filename = string(in.String())
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
func easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(out *jwriter.Writer, in AttachmentInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"taskID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.TaskID))
	}
	{
		const prefix string = ",\"attachmentID\":"
		out.RawString(prefix)
		out.Int64(int64(in.AttachmentID))
	}
	{
		const prefix string = ",\"attachmentFileName\":"
		out.RawString(prefix)
		out.String(string(in.Filename))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AttachmentInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AttachmentInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AttachmentInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AttachmentInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(l, v)
}
func easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(in *jlexer.Lexer, out *AttachmentFileInternal) {
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
		case "UserID":
			out.UserID = int64(in.Int64())
		case "Filename":
			out.Filename = string(in.String())
		case "File":
			if in.IsNull() {
				in.Skip()
				out.File = nil
			} else {
				out.File = in.Bytes()
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
func easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(out *jwriter.Writer, in AttachmentFileInternal) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"UserID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.UserID))
	}
	{
		const prefix string = ",\"Filename\":"
		out.RawString(prefix)
		out.String(string(in.Filename))
	}
	{
		const prefix string = ",\"File\":"
		out.RawString(prefix)
		out.Base64Bytes(in.File)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AttachmentFileInternal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AttachmentFileInternal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AttachmentFileInternal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AttachmentFileInternal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(l, v)
}
func easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(in *jlexer.Lexer, out *Attachment) {
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
		case "AttachmentID":
			out.AttachmentID = int64(in.Int64())
		case "TaskID":
			out.TaskID = int64(in.Int64())
		case "Filename":
			out.Filename = string(in.String())
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
func easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(out *jwriter.Writer, in Attachment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"AttachmentID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.AttachmentID))
	}
	{
		const prefix string = ",\"TaskID\":"
		out.RawString(prefix)
		out.Int64(int64(in.TaskID))
	}
	{
		const prefix string = ",\"Filename\":"
		out.RawString(prefix)
		out.String(string(in.Filename))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Attachment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Attachment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8a499b54EncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Attachment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Attachment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8a499b54DecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(l, v)
}