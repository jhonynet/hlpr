/**
HTTP Processor.

USAGE:
type: http						Type of this processor.
content: json					Response Content-Type.
url: '{{ . }}/applications'		Url, can be a template.
threads: N						Number of threads.
headers:						map header:value, can be a template.
  x-auth-token: '{{ . }}'

*/

package http
