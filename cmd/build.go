/*
Copyright © 2020 Jeff MONNIER <kessler.dev@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kesslerdev/gitgen/internal/build"
	"github.com/kesslerdev/gitgen/pkg/generator"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build [generator-name]",
	Short: "build a generator",
	Long: `Genrate templates files using terms provided in generator config:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sublogger := log.With().
			Str("cmd", "build").
			Logger()
		file := args[0]
		if !strings.HasSuffix(file, generator.GeneratorFileSuffix) {
			file = fmt.Sprintf(generator.GeneratorFileFormat, file)
		}
		file = filepath.Join(dir, file)
		g := generator.Generator{}

		sublogger.Info().Msgf("Building generator at %s", file)

		if err := generator.FromYamlFile(file, &g); err != nil {
			sublogger.Fatal().Err(err).Send()
		}
		// indicate relative from generator file
		if err := build.BuildGenerator(filepath.Dir(file), &g); err != nil {
			sublogger.Fatal().Err(err).Send()
		}

		sublogger.Info().Msg("OK")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
