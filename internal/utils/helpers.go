// Package utils provides shared utility functions used across the SDK.
package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// Numbers joins multiple phone numbers with the "<" separator used by the Vobiz API.
func Numbers(numbers ...string) string {
	return strings.Join(numbers, "<")
}

// Headers encodes a map of header key/value pairs into a sorted, URL-escaped string.
func Headers(headers map[string]string) string {
	return headersWithSep(headers, "=", ",", true)
}

func headersWithSep(headers map[string]string, keyValSep, itemSep string, escape bool) string {
	v := url.Values{}
	for k, val := range headers {
		v.Set(k, val)
	}
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for _, k := range keys {
		if buf.Len() > 0 {
			buf.WriteString(itemSep)
		}
		if escape {
			buf.WriteString(url.QueryEscape(k) + keyValSep + url.QueryEscape(v.Get(k)))
		} else {
			buf.WriteString(k + keyValSep + v.Get(k))
		}
	}
	return buf.String()
}

// ComputeSignature computes the legacy (V1) HMAC-SHA1 signature.
func ComputeSignature(authToken, uri string, params map[string]string) string {
	raw := fmt.Sprintf("%s%s", uri, headersWithSep(params, "", "", false))
	mac := hmac.New(sha1.New, []byte(authToken))
	mac.Write([]byte(raw))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// ValidateSignature validates a V1 signature.
func ValidateSignature(authToken, uri string, params map[string]string, signature string) bool {
	return ComputeSignature(authToken, uri, params) == signature
}

// ComputeSignatureV2 computes a V2 HMAC-SHA256 signature using the nonce.
func ComputeSignatureV2(authToken, uri, nonce string) string {
	parsed, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	raw := parsed.Scheme + "://" + parsed.Host + parsed.Path + nonce
	mac := hmac.New(sha256.New, []byte(authToken))
	mac.Write([]byte(raw))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// ValidateSignatureV2 validates a V2 signature.
func ValidateSignatureV2(uri, nonce, signature, authToken string) bool {
	return ComputeSignatureV2(authToken, uri, nonce) == signature
}

// GenerateURL builds a canonical URL for V3 signature computation.
func GenerateURL(uri, method string, params map[string]string) string {
	parsed, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	base := parsed.Scheme + "://" + parsed.Host + parsed.Path
	if len(params) == 0 && len(parsed.RawQuery) == 0 {
		return base
	}
	base += "?"
	if len(parsed.RawQuery) > 0 {
		existing := getMapFromQuery(parsed.Query())
		if method == "GET" {
			for k, v := range params {
				existing[k] = v
			}
			return base + sortedQueryString(existing, true)
		}
		return base + sortedQueryString(existing, true) + "." + sortedQueryString(params, false)
	}
	if method == "GET" {
		return base + sortedQueryString(params, true)
	}
	return base + sortedQueryString(params, false)
}

// ComputeSignatureV3 computes a V3 HMAC-SHA256 signature.
func ComputeSignatureV3(authToken, uri, method, nonce string, params map[string]string) string {
	raw := GenerateURL(uri, method, params) + "." + nonce
	mac := hmac.New(sha256.New, []byte(authToken))
	mac.Write([]byte(raw))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// ValidateSignatureV3 validates a V3 signature (supports comma-separated multi-sig).
func ValidateSignatureV3(uri, nonce, method, signature, authToken string, params ...map[string]string) bool {
	p := map[string]string{}
	if len(params) > 0 {
		p = params[0]
	}
	computed := ComputeSignatureV3(authToken, uri, method, nonce, p)
	for _, s := range strings.Split(signature, ",") {
		if s == computed {
			return true
		}
	}
	return false
}

// Find returns true if val is present in slice.
func Find(val string, slice []string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

func getMapFromQuery(q url.Values) map[string]string {
	m := make(map[string]string, len(q))
	for k, v := range q {
		m[k] = v[0]
	}
	return m
}

func sortedQueryString(params map[string]string, asQuery bool) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sb strings.Builder
	for _, k := range keys {
		if asQuery {
			sb.WriteString(k + "=" + params[k] + "&")
		} else {
			sb.WriteString(k + params[k])
		}
	}
	return strings.TrimRight(sb.String(), "&")
}
