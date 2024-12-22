package pgext

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func (e *Extension) PrintInfo() {
	tmpl, err := template.New("extension").Funcs(template.FuncMap{
		"join": join,
	}).Parse(extensionInfoTmpl)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, e); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	fmt.Println(buf.String())
}

const extensionInfoTmpl = `
╭────────────────────────────────────────────────────────────────────────────╮
│ {{ printf "%-74s" .Name   }} │
├────────────────────────────────────────────────────────────────────────────┤
│ {{ printf "%-74s" .EnDesc }} │
├────────────────────────────────────────────────────────────────────────────┤
│ Extension : {{ printf "%-62s" .Name        }} │
│ Alias     : {{ printf "%-62s" .Alias       }} │
│ Category  : {{ printf "%-62s" .Category    }} │
│ Version   : {{ printf "%-62s" .Version     }} │
│ License   : {{ printf "%-62s" .License     }} │
│ Website   : {{ printf "%-62s" .URL         }} │
│ Details   : {{ printf "%-62s" .SummaryURL  }} │
├────────────────────────────────────────────────────────────────────────────┤
│ Extension Properties                                                       │
├────────────────────────────────────────────────────────────────────────────┤
│ PostgreSQL Ver │  Available on: {{ printf "%-42s" (join .PgVer ", ") }} │
│ CREATE  :  {{ if .NeedDDL  }}Yes{{ else }}No {{ end }} │  {{ printf "%-56s" .CreateSQL }} │
│ DYLOAD  :  {{ if .NeedLoad }}Yes{{ else }}No {{ end }} │  {{ printf "%-56s" .SharedLib }} │
│ {{ printf "%-74s" .SuperUser }} │
│ Reloc   :  {{ if eq .Relocatable "t" }}Yes{{ else }}No {{ end }} │  {{ printf "%-56s" .SchemaStr }} │
{{- if .Requires }}
│ Depend  :  Yes │  {{ printf "%-56s" (join .Requires ", ") }} │
{{- else }}
│ Depend  :  No  │                                                           │
{{- end }}
{{- if .NeedBy }}
├────────────────────────────────────────────────────────────────────────────┤
│ Required By                                                                │
├────────────────────────────────────────────────────────────────────────────┤
{{- range .NeedBy }}
│ - {{ printf "%-72s" . }} │
{{- end }}
{{- end }}

{{- if .RpmRepo }}
├────────────────────────────────────────────────────────────────────────────┤
│ RPM Package                                                                │
├────────────────────────────────────────────────────────────────────────────┤
│ Repository     │  {{ printf "%-56s" .RpmRepo }} │
│ Package        │  {{ printf "%-56s" .RpmPkg  }} │
│ Version        │  {{ printf "%-56s" .RpmVer  }} │
│ Availability   │  {{ printf "%-56s" (join .RpmPg ", ") }} │
{{- if .DebDeps }}
│ Dependencies   │  {{ printf "%-56s" (join .RpmDeps ", ") }} │
{{- end }}
{{- end }}

{{- if .DebRepo }}
├────────────────────────────────────────────────────────────────────────────┤
│ DEB Package                                                                │
├────────────────────────────────────────────────────────────────────────────┤
│ Repository     │  {{ printf "%-56s" .DebRepo }} │
│ Package        │  {{ printf "%-56s" .DebPkg  }} │
│ Version        │  {{ printf "%-56s" .DebVer  }} │
│ Availability   │  {{ printf "%-56s" (join .DebPg ", ") }} │
{{- if .DebDeps }}
│ Dependencies   │  {{ printf "%-56s" (join .DebDeps ", ") }} │
{{- end }}
{{- end }}

{{- if .BadCase }}
├────────────────────────────────────────────────────────────────────────────┤
│ Known Issues                                                               │
├────────────────────────────────────────────────────────────────────────────┤
{{- range .BadCase }}
│ {{ printf "%-74s" . }} │
{{- end }}
{{- end }}

{{- if .Comment }}
├────────────────────────────────────────────────────────────────────────────┤
│ Additional Comments                                                        │
├────────────────────────────────────────────────────────────────────────────┤
│ {{ printf "%-74s" .Comment }} │
{{- end }}
╰────────────────────────────────────────────────────────────────────────────╯
`

func join(strs []string, sep string) string {
	return strings.Join(strs, sep)
}
