/*
 * Charts API
 *
 * An API for creating, retrieving, updating, and deleting charts
 *
 * API version: 2.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package chart

// An object containing properties that specify the range of color values to use in a Heatmap chart, including the lowest and highest color to use. The `options.colorScale` property controls how SignalFx sets each intermediate chart color within the range.
type HeatmapColorRangeOptions struct {
	// The starting hex color value for data values in a heatmap chart. The `options.colorScale` property controls how SignalFx creates a range of colors to represent different data values; `options.colorRange.color` is the starting point of the range, and `options.colorScale` controls the number of different colors.<br> For best results, set `options.colorRange.color` to one of the 21 colors in the official SignalFx color palette.<br> Specify the value as a 6-character hexadecimal value preceded by the '#' character, for example \"#00b9ff\" (light blue)
	Color string `json:"color,omitempty"`
	// The color value associated with the largest incoming data value in the heatmap. <br> Specify the value as a 6-character hexadecimal value preceded by the '#' character, for example \"#ea1849\" (grass green).
	Max string `json:"max,omitempty"`
	// The color value associated with the smallest incoming data value in the heatmap. <br> Specify the value as a 6-character hexadecimal value preceded by the '#' character, for example \"#999999\".
	Min string `json:"min,omitempty"`
}
