/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/nycu-ucr/gonet/http"

	// "golang.org/x/net/http2"
	"github.com/nycu-ucr/net/http2"
	"golang.org/x/oauth2"
)

var (
	innerHTTP2Client = &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	innerHTTP2CleartextClient = &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	innerHTTP2OnvmClient = &http.Client{
		Transport: &http2.OnvmTransport{},
	}
)

type Configuration interface {
	BasePath() string
	Host() string
	UserAgent() string
	DefaultHeader() map[string]string
	HTTPClient() *http.Client
}

// SelectHeaderAccept join all accept types and return
func SelectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func TypeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// ParameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func ParameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func CallAPI(cfg Configuration, request *http.Request) (*http.Response, error) {
	if cfg.HTTPClient() != nil {
		return cfg.HTTPClient().Do(request)
	}
	if true {
		return innerHTTP2OnvmClient.Do(request)
	}
	if request.URL.Scheme == "https" {
		return innerHTTP2Client.Do(request)
	} else if request.URL.Scheme == "http" {
		return innerHTTP2CleartextClient.Do(request)
	}

	return nil, fmt.Errorf("unsupported scheme[%s]", request.URL.Scheme)
}

// // Change base path to allow switching to mocks
// func ChangeBasePath(cfg Configuration, path string) {
// 	cfg.BasePath() = path

// }
func parseMultipartFieldParameters(str string) (contentType string, ref string, class string) {
	for _, part := range strings.Split(str, ",") {
		switch {
		case strings.HasPrefix(part, "contentType:"):
			contentType = part[12:]
		case strings.HasPrefix(part, "ref:"):
			ref = part[4:]
		case strings.HasPrefix(part, "class:"):
			class = part[6:]
		}
	}
	return
}

func getContentID(v reflect.Value, ref string, class string) (contentID string, err error) {
	recursiveVal := v
	if ref[0] == '{' {
		contentID = ref[1 : len(ref)-1]
		return
	}
	if class != "" {
		var lastVal reflect.Value
		for _, part := range strings.Split(class, ".") {
			lastVal = recursiveVal
			recursiveVal = recursiveVal.FieldByName(part)
			if !recursiveVal.IsValid() {
				return "", fmt.Errorf("Do not have reference field %s in %s for multipart", part, lastVal.Type().String())
			}
			if recursiveVal.Kind() == reflect.Ptr {
				if recursiveVal.IsNil() {
					return "", nil
				}
				recursiveVal = recursiveVal.Elem()
			}
		}
		fieldName := recursiveVal.String()
		if i := strings.IndexRune(fieldName, '-'); i != -1 {
			fieldName = fieldName[:i]
		}
		fieldName = fieldName[:1] + strings.ToLower(fieldName[1:]) + "Info"
		recursiveVal = lastVal.FieldByName(fieldName)
		if recursiveVal.Kind() == reflect.Ptr {
			if recursiveVal.IsNil() {
				return "", nil
			}
			recursiveVal = recursiveVal.Elem()
		}
	}

	for _, part := range strings.Split(ref, ".") {
		lastValType := recursiveVal.Type()
		listIndex := -1

		if start := strings.IndexRune(part, '['); start != -1 {
			end := strings.IndexRune(part, ']')
			listIndex, err = strconv.Atoi(part[start+1 : end])
			if err != nil {
				return "", err
			}
			part = part[:start]
			recursiveVal = recursiveVal.FieldByName(part)
		} else if start = strings.IndexRune(part, '('); start != -1 {
			end := strings.IndexRune(part, ')')
			fieldTypeString := part[start+1 : end]
			var i int
			for i = 0; i < lastValType.NumField(); i++ {
				if fieldType := lastValType.Field(i).Type; strings.HasSuffix(fieldType.String(), fieldTypeString) {
					recursiveVal = recursiveVal.Field(i)
					break
				}
			}
			if i == lastValType.NumField() {
				return "", fmt.Errorf("Do not have reference field Type %s in %s for multipart", fieldTypeString, lastValType.String())
			}
		} else {
			recursiveVal = recursiveVal.FieldByName(part)
		}

		if !recursiveVal.IsValid() {
			return "", fmt.Errorf("Do not have reference field %s in %s for multipart", part, lastValType.String())
		}
		if recursiveVal.Kind() == reflect.Ptr {
			if recursiveVal.IsNil() {
				return "", nil
			}
			recursiveVal = recursiveVal.Elem()
		}
		if listIndex >= 0 {
			if listIndex >= recursiveVal.Len() {
				return "", nil
			}
			recursiveVal = recursiveVal.Index(listIndex)
		}
	}
	contentID = recursiveVal.String()
	return
}

func MultipartEncode(v interface{}, body io.Writer) (string, error) {
	val := reflect.Indirect(reflect.ValueOf(v))
	w := multipart.NewWriter(body)

	structType := val.Type()

	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).IsNil() {
			continue
		}
		contentType, ref, class := parseMultipartFieldParameters(structType.Field(i).Tag.Get("multipart"))
		h := make(textproto.MIMEHeader)

		if ref != "" {
			if contentID, err := getContentID(val, ref, class); err != nil {
				return "", err
			} else if contentID != "" {
				h.Set("Content-ID", contentID)
			} else {
				return "", fmt.Errorf("ContentID of multipart binaryData in JsonData is empty")
			}
		}
		h.Set("Content-Type", contentType)
		fieldData, err := w.CreatePart(h)
		if err != nil {
			return "", err
		}
		fieldBody, err := setBody(val.Field(i).Interface(), contentType)
		if err != nil {
			return "", err
		}
		_, err = fieldData.Write(fieldBody.Bytes())
		if err != nil {
			return "", err
		}
	}
	err := w.Close()
	if err != nil {
		return "", err
	}
	contentType := "multipart/related; boundary=\"" + w.Boundary() + "\""

	return contentType, nil
}

func MultipartSerialize(v interface{}) ([]byte, string, error) {
	buffer := new(bytes.Buffer)
	val := reflect.Indirect(reflect.ValueOf(v))
	w := multipart.NewWriter(buffer)

	structType := val.Type()

	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).IsNil() {
			continue
		}
		contentType, ref, class := parseMultipartFieldParameters(structType.Field(i).Tag.Get("multipart"))
		h := make(textproto.MIMEHeader)

		if ref != "" {
			if contentID, err := getContentID(val, ref, class); err != nil {
				return nil, "", err
			} else if contentID != "" {
				h.Set("Content-ID", contentID)
			} else {
				return nil, "", fmt.Errorf("ContentID of multipart binaryData in JsonData is empty")
			}
		}
		h.Set("Content-Type", contentType)
		fieldData, err := w.CreatePart(h)
		if err != nil {
			return nil, "", err
		}
		fieldBody, err := setBody(val.Field(i).Interface(), contentType)
		if err != nil {
			return nil, "", err
		}
		_, err = fieldData.Write(fieldBody.Bytes())
		if err != nil {
			return nil, "", err
		}
	}
	err := w.Close()
	if err != nil {
		return nil, "", err
	}
	contentType := "multipart/related; boundary=\"" + w.Boundary() + "\""

	return buffer.Bytes(), contentType, nil
}

// prepareRequest build the request
func PrepareRequest(
	ctx context.Context,
	cfg Configuration,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFileName string,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {

		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		if strings.HasPrefix(headerParams["Content-Type"], "multipart/related") {

			body = &bytes.Buffer{}
			contentType, err = MultipartEncode(postBody, body)
			if err != nil {
				return nil, err
			}
			headerParams["Content-Type"] = contentType

		} else {
			body, err = setBody(postBody, contentType)
			if err != nil {
				return nil, err
			}
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile(formFileName, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if cfg.Host() != "" {
		localVarRequest.Host = cfg.Host()
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", cfg.UserAgent())

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			println("Error: latestToken.SetAuthHeader(localVarRequest), ", latestToken)
			// latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range cfg.DefaultHeader() {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func MultipartDeserialize(b []byte, v interface{}, boundary string) (err error) {
	body := bytes.NewReader(b)
	r := multipart.NewReader(body, boundary)
	val := reflect.Indirect(reflect.ValueOf(v))

	contentIDIndex := make(map[string]int)

	for {
		var part *multipart.Part
		multipartBody := make([]byte, 1000)

		// if no remian part, break this loop
		if nextPart, err := r.NextPart(); err == io.EOF {
			break
		} else {
			part = nextPart
		}

		contentType := part.Header.Get("Content-Type")
		var n int
		n, err = part.Read(multipartBody)
		if err == nil {
			return
		}
		multipartBody = multipartBody[:n]

		kind := KindOfMediaType(contentType)

		if kind == MediaKindJSON {
			value := val.Field(0)
			if value.Kind() == reflect.Ptr {
				ptr := reflect.New(value.Type().Elem())
				value.Set(ptr)
			}
			if err = json.Unmarshal(multipartBody, value.Interface()); err != nil {
				return err
			}
			structType := val.Type()
			for i := 1; i < structType.NumField(); i++ {
				_, ref, class := parseMultipartFieldParameters(structType.Field(i).Tag.Get("multipart"))
				if ref != "" {
					if contentID, err := getContentID(val, ref, class); err != nil {
						return err
					} else if contentID != "" {
						contentIDIndex[contentID] = i
					}
				}
			}
		} else {
			contentID := part.Header.Get("Content-ID")
			if index, ok := contentIDIndex[contentID]; ok {
				value := val.Field(index)
				value.SetBytes(multipartBody)
			} else {
				return fmt.Errorf("multipart binary data need Content-ID")
			}
		}
	}

	return nil
}

func Deserialize(v interface{}, b []byte, contentType string) (err error) {
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}

	switch KindOfMediaType(contentType) {
	case MediaKindJSON:
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	case MediaKindXML:
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	case MediaKindMultipartRelated:
		boundary := ""
		for _, part := range strings.Split(contentType, ";") {
			if strings.HasPrefix(part, " boundary=") {
				boundary = part[10:]
			}
		}
		if boundary == "" {
			return errors.New("multipart/related need boundary")
		}
		boundary = strings.Trim(boundary, "\" ")
		if err = MultipartDeserialize(b, v, boundary); err != nil {
			return err
		}
		return nil
	case MediaKindUnsupported:
		return errors.New("undefined response type")
	default:
		return errors.New("undefined response type")
	}
}

// add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func ReportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else {
		b, err = Serialize(body, contentType)
		if err != nil {
			return nil, err
		}
		_, err = bodyBuf.Write(b)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

func InterceptH2CClient() {
	println("Error: gock.InterceptClient(innerHTTP2CleartextClient)")
	// gock.InterceptClient(innerHTTP2CleartextClient)
}

func RestoreH2CClient() {
	println("Error: gock.RestoreClient(innerHTTP2CleartextClient)")
	// gock.RestoreClient(innerHTTP2CleartextClient)
}
