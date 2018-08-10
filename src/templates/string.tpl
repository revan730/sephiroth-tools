metadata:
  description: {{ .Description }}
items:
  {{ range $key, $elem := .Items }}
  - {{ $key }}: "{{$elem}}"
  {{ end }}