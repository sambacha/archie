/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package archie

import (
	"github.com/spf13/cobra"
)

var tagScope string
var tagTag string

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Generates a context diagram",
	Long: `Generates a diagram that shows the context of elements with a specified tag

[1] Main elements of interest
The 'eldest' element with the specified tag,

[2] Relevant associated elements
Those that are associated to one of the main elements of interest, where either:
- The parent is an ancestor of scope
- It is a root element.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Generate the diagram
		diagram, err = arch.TagView(tagScope, tagTag)
	},
}

func init() {
	generateCmd.AddCommand(tagCmd)

	fs := tagCmd.Flags()
	fs.StringVarP(&tagScope, "scope", "s", "", "scope for the tag view")
	fs.StringVarP(&tagTag, "tag", "t", "", "tag to filter by")

	cobra.MarkFlagRequired(fs, "scope")
	cobra.MarkFlagRequired(fs, "tag")
}
