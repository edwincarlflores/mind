// Code generated by templ@v0.2.316 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"

func Navbar(mindID int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<nav class=\"fixed inset-x-0 z-50 bg-zinc-100 bg-opacity-70 shadow-lg backdrop-blur-lg\"><div class=\"mx-auto max-w-full px-4\"><div class=\"flex items-center sm:space-x-4 sm:px-3\"><div class=\"font-bold text-gray-700 hover:text-gray-900\"><a hx-get=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(fmt.Sprintf("%d/thoughts", mindID)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-target=\"#thoughts-container\" hx-swap=\"innerHTML\" class=\"flex items-center\"><img src=\"/static/images/mind-logo.png\" class=\"h-12 w-12\" alt=\"Mind\"><span class=\"hidden font-mono text-lg sm:block\">")
		if err != nil {
			return err
		}
		var_2 := `Mind`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span></a></div><div class=\"grow p-4\"><div class=\"w-full md:w-1/2\">")
		if err != nil {
			return err
		}
		var_3 := `Search Field Placeholder`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div></div></div></div></nav>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
