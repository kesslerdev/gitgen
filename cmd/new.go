/*
Package cmd contains alls commands for gitgen
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

	"github.com/gosimple/slug"
	"github.com/kesslerdev/gitgen/internal/utils"
	"github.com/kesslerdev/gitgen/pkg/generator"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [generator name]",
	Short: "create a new generator in current path",
	Long: `Create a sample [generator name].gitgen.yaml:

In future you can pass several metadata to attach with it.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := slug.MakeLang(args[0], "en")
		sublogger := log.With().
			Str("cmd", "new").
			Logger()

		sublogger.Info().Msgf("Using %s as generator name\n", name)

		path := fmt.Sprintf(generator.GeneratorFileFormat, name)
		g := generator.NewGenerator(name)

		strategy, err := cmd.Flags().GetString("build-input-strategy")

		if err != nil {
			sublogger.Fatal().Err(err).Send()
		} else {
			g.Spec.Build.Input.Strategy = strategy
		}

		if err := utils.WriteNewYaml(path, &g); err != nil {
			sublogger.Fatal().Err(err).Send()
		} else {
			sublogger.Info().Msgf("Generator %s created at %s\n", name, path)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	newCmd.Flags().StringP("build-input-strategy", "b", "static", "Strategy used to gather files")
}
