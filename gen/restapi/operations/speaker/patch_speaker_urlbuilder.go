// Code generated by go-swagger; DO NOT EDIT.

package speaker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PatchSpeakerURL generates an URL for the patch speaker operation
type PatchSpeakerURL struct {
	MeetupID  string
	SpeakerID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PatchSpeakerURL) WithBasePath(bp string) *PatchSpeakerURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PatchSpeakerURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PatchSpeakerURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/meetups/{meetup_id}/speakers/{speaker_id}"

	meetupID := o.MeetupID
	if meetupID != "" {
		_path = strings.Replace(_path, "{meetup_id}", meetupID, -1)
	} else {
		return nil, errors.New("meetupId is required on PatchSpeakerURL")
	}

	speakerID := o.SpeakerID
	if speakerID != "" {
		_path = strings.Replace(_path, "{speaker_id}", speakerID, -1)
	} else {
		return nil, errors.New("speakerId is required on PatchSpeakerURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PatchSpeakerURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PatchSpeakerURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PatchSpeakerURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PatchSpeakerURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PatchSpeakerURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PatchSpeakerURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
