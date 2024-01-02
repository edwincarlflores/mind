// Code generated by templ@v0.2.316 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/edwincarlflores/mind/aggregate"
	"github.com/edwincarlflores/mind/entity"
	common "github.com/edwincarlflores/mind/templates/common"
)

func Thought(thought *entity.Thought) templ.Component {
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
		_, err = templBuffer.WriteString("<div class=\"flex flex-row space-x-3\"><p>")
		if err != nil {
			return err
		}
		var var_2 string = thought.Body
		_, err = templBuffer.WriteString(templ.EscapeString(var_2))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><input type=\"checkbox\"")
		if err != nil {
			return err
		}
		if thought.Public {
			_, err = templBuffer.WriteString(" checked")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func Thoughts(thoughts []*entity.Thought) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_3 := templ.GetChildren(ctx)
		if var_3 == nil {
			var_3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		for _, thought := range thoughts {
			err = Thought(thought).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func MindPage(mind *aggregate.Mind) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var_5 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div id=\"thoughts-container\">")
			if err != nil {
				return err
			}
			var_6 := `Hi `
			_, err = templBuffer.WriteString(var_6)
			if err != nil {
				return err
			}
			var var_7 string = mind.User.Name
			_, err = templBuffer.WriteString(templ.EscapeString(var_7))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			var_8 := `(`
			_, err = templBuffer.WriteString(var_8)
			if err != nil {
				return err
			}
			var var_9 string = mind.User.Username
			_, err = templBuffer.WriteString(templ.EscapeString(var_9))
			if err != nil {
				return err
			}
			var_10 := `)`
			_, err = templBuffer.WriteString(var_10)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(" ")
			if err != nil {
				return err
			}
			err = Thoughts(mind.Thoughts).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = common.Page("Mind", mind.User.Username).Render(templ.WithChildren(ctx, var_5), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}
