package processor_test

import (
	"fmt"

	"github.com/fatih/color"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JulzDiverse/resc/processor"
)

var _ = Describe("Processor", func() {

	var (
		processor       *Processor
		headerFormat    *color.Color
		subheaderFormat *color.Color
		boldFormat      *color.Color
		italicFormat    *color.Color
	)

	JustBeforeEach(func() {
		headerFormat = color.New(color.FgCyan, color.Bold, color.Underline)
		subheaderFormat = color.New(color.FgWhite, color.Bold)
		boldFormat = color.New(color.Bold)
		italicFormat = color.New(color.Italic)
		processor = New()
	})

	Context("When the input text is a markdown header", func() {
		It("formats the text into the header format", func() {
			result := processor.Process("# Header")
			Expect(result).To(Equal(headerFormat.Sprintf("%s", "HEADER")))
		})
	})

	Context("When the input text is a markdown subheader", func() {
		It("formats the text into the subheader format", func() {
			result := processor.Process("## SubHeader")
			Expect(result).To(Equal(subheaderFormat.Sprintf("%s", "SubHeader")))
		})
	})

	Context("When the input text is bold markdown text", func() {
		Context("When a single word is bolded", func() {
			It("formats the text into bold format", func() {
				result := processor.Process("**bold**")
				Expect(result).To(Equal(boldFormat.Sprintf("%s", "bold")))
			})
		})
		Context("When multiple words are bolded", func() {
			It("formats the text into bold format", func() {
				result := processor.Process("**bold text**")
				Expect(result).To(Equal(boldFormat.Sprintf("%s", "bold text")))
			})
		})
		Context("When a string contains bold markdown text", func() {
			It("formats only the bold text", func() {
				result := processor.Process("This **bold text** is bold")
				boldText := boldFormat.Sprintf("bold text")
				Expect(result).To(Equal(fmt.Sprintf("This %s is bold", boldText)))
			})
		})
		Context("When a string contains special characters", func() {
			It("still formats the bold text", func() {
				result := processor.Process("**Hello, World!**")
				Expect(result).To(Equal(boldFormat.Sprintf("%s", "Hello, World!")))
			})
		})
	})

	Context("When the input text is italic markdown text", func() {
		Context("When a single word is italic", func() {
			It("formats the text into italic format", func() {
				result := processor.Process("_italic_")
				Expect(result).To(Equal(italicFormat.Sprintf("%s", "italic")))
			})
		})
		Context("When multiple words are italic", func() {
			It("formats the text into italic format", func() {
				result := processor.Process("_italic text_")
				Expect(result).To(Equal(italicFormat.Sprintf("%s", "italic text")))
			})
		})
		Context("When a string contains italic markdown text", func() {
			It("formats only the italic text", func() {
				result := processor.Process("This _italic text_ is italic")
				boldText := italicFormat.Sprintf("italic text")
				Expect(result).To(Equal(fmt.Sprintf("This %s is italic", boldText)))
			})
		})

		Context("When a string contains special characters", func() {
			It("still formats the italic text", func() {
				result := processor.Process("_Hello, World!_")
				Expect(result).To(Equal(italicFormat.Sprintf("%s", "Hello, World!")))
			})
		})
	})

	Context("When the input text has no markdown format", func() {
		It("returns the string as it is", func() {
			result := processor.Process("hello")
			Expect(result).To(Equal("hello"))
		})
	})
})
