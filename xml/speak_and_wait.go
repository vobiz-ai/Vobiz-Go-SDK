package xml

// SpeakAndWait appends a Speak + Wait pattern.
func (r *Response) SpeakAndWait(text string, waitLength int) *Response {
	if waitLength <= 0 {
		waitLength = 3
	}
	return r.Speak(text).Wait(waitLength)
}
